# Copyright (c) 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from datetime import datetime
import docker
import logging
import os
import pwd
import requests
import socket
import sys
import time

import six

_DOCKER_VOLUMES = {
    # The following four mounts are needed to add the host's chrome-bot user in
    # the container.
    '/home/chrome-bot': {'bind': '/home/chrome-bot', 'mode': 'rw'},
    '/etc/shadow': {'bind': '/etc/shadow', 'mode': 'ro'},
    '/etc/passwd': {'bind': '/etc/passwd', 'mode': 'ro'},
    '/etc/group': {'bind': '/etc/group', 'mode': 'ro'},
    # Needed to give chrome-bot the same root access inside the container.
    '/etc/sudoers': {'bind': '/etc/sudoers', 'mode': 'ro'},
    '/etc/sudoers.d': {'bind': '/etc/sudoers.d', 'mode': 'ro'},
    # Needed by swarming bot to auth with server.
    '/var/lib/luci_machine_tokend': {
        'bind': '/var/lib/luci_machine_tokend',
        'mode': 'ro',
    },
    # Needed for authenticating with monitoring endpoints.
    '/creds/service_accounts': {
        'bind': '/creds/service_accounts',
        'mode': 'ro'
    },
    '/etc/chrome-infra/ts-mon.json': {
        'bind': '/etc/chrome-infra/ts-mon.json',
        'mode': 'ro'
    },
    # Needed to access mmutex locks shared with puppet running outside the
    # container (see http://crbug.com/808060).
    '/mmutex': {
      'bind': '/mmutex',
      'mode': 'rw',
    },
}

_SWARMING_URL_ENV_VAR = 'SWARM_URL'
_HOST_HOSTNAME_ENV_VAR = 'DOCKER_HOST_HOSTNAME'
_KVM_DEVICE = '/dev/kvm'
_TUN_DEVICE = '/dev/net/tun'


class FrozenEngineError(Exception):
  """Raised when the docker engine is unresponsive."""


class FrozenContainerError(Exception):
  """Raised when a container is unresponsive."""


class ContainerDescriptorBase(object):
  @property
  def name(self):
    """Returns name to be used for the container."""
    raise NotImplementedError()

  @property
  def shutdown_file(self):
    """Returns the name of the file to drain the swarm bot in the container."""
    raise NotImplementedError()

  @property
  def lock_file(self):
    """Returns the name of the file to flock on when managing the container."""
    raise NotImplementedError()

  @property
  def hostname(self):
    """Returns hostname to be used for the container."""
    raise NotImplementedError()

  def log_started(self):
    """Logs a debug message that the container has been started."""
    raise NotImplementedError()

  def should_create_container(self):
    """Returns true if the container should be created for this descriptor."""
    return True


class ContainerDescriptor(ContainerDescriptorBase):
  def __init__(self, name):
    self._name = name

  @property
  def name(self):
    return self._name

  @property
  def shutdown_file(self):
    return '/b/%s.shutdown.stamp' % self._name

  @property
  def lock_file(self):
    return '/var/lock/swarm_docker.%s.lock' % self._name

  @property
  def hostname(self):
    this_host = socket.gethostname().split('.')[0]
    return '%s--%s' % (this_host, self._name)

  def log_started(self):
    logging.debug('Launched new container %s.', self._name)


class DockerClient(object):
  def __init__(self):
    self._client = docker.from_env()
    self._num_configured_containers = None
    self.volumes = _DOCKER_VOLUMES.copy()

  @staticmethod
  def image_matches_url(image, image_url):
    """Checks if the passed in image matches the given URL."""
    return any(t == image_url for t in image.tags)

  def ping(self, retries=5):
    """Checks if the engine is responsive.

    Will sleep with in between retries with exponential backoff.
    Returns True if engine is responding, else False.
    """
    sleep_time = 1
    for i in range(retries):
      try:
        self._client.ping()
        return True
      except (docker.errors.APIError, requests.exceptions.ConnectionError):
        pass
      if i < retries - 1:
        time.sleep(sleep_time)
        sleep_time *= 2
    return False

  def pull(self, image_url):
    self._client.images.pull(image_url)

  def images(self):
    """Lists all local images.

    Note: this returns a list of Image objects with 'id' and 'tags', the former
    being a checksum of the image contents and the latter a list of the
    human-readable URLs. Various APIs expect either one or the other, so we need
    to provide both here unfortunately.
    """
    return self._client.images.list()

  def has_image(self, image_url):
    try:
      self._client.images.get(image_url)
      return True
    except docker.errors.ImageNotFound:
      return False

  def remove_image(self, image_id):
    self._client.images.remove(image_id)

  def remove_outdated_images(self, most_recent_image_url):
    """Removes all local images but for most_recent_image_url."""
    for image in self.images():
      if not self.image_matches_url(image, most_recent_image_url):
        logging.info(
            'Current image (%s) not found in the tags of %s: [%s]. Deleting.',
            most_recent_image_url, image.id, ' '.join(image.tags))
        self.remove_image(image.id)

  def _get_containers_by_status(self, status):
    return [
        Container(c) for c in self._client.containers.list(
            filters={'status': status})
    ]

  def get_created_containers(self):
    return self._get_containers_by_status('created')

  def get_exited_containers(self):
    return self._get_containers_by_status('exited')

  def get_paused_containers(self):
    return self._get_containers_by_status('paused')

  def get_running_containers(self):
    return self._get_containers_by_status('running')

  def get_container(self, container_desc):
    try:
      return Container(self._client.containers.get(container_desc.name))
    except docker.errors.NotFound:
      logging.error('No running container %s.', container_desc.name)
      return None

  def stop_old_containers(self, running_containers, max_uptime):
    now = datetime.utcnow()
    frozen_containers = 0
    for container in running_containers:
      uptime = container.get_container_uptime(now)
      logging.debug(
          'Container %s has uptime of %s minutes.', container.name, str(uptime))
      if uptime is not None and uptime > max_uptime:
        try:
          container.kill_swarming_bot(max_uptime=max_uptime)
        except FrozenContainerError:
          frozen_containers += 1
    if running_containers and frozen_containers == len(running_containers):
      logging.error('All containers frozen. Docker engine most likely hosed.')
      raise FrozenEngineError()

  def delete_stopped_containers(self):
    for c in self.get_exited_containers():
      logging.debug('Found stopped container %s. Removing it.', c.name)
      c.remove()

    # Occasionally containers will fail to enter the "run" state after
    # they have been "created". This will remove any containers in this
    # state. See the issue below for more details:
    # https://github.com/moby/moby/issues/8294
    for c in self.get_created_containers():
      logging.error(
          'Container %s failed to enter a running state and is currently '
          'stopped in a "Created" state with exit code %s. Removing it.',
          c.name, str(c.exit_code))
      # It's already stopped, so removal is the only option to fix this.
      c.remove(force=True)

  def _get_env(self, swarming_url):
    env = {
        _SWARMING_URL_ENV_VAR: swarming_url + '/bot_code',
        _HOST_HOSTNAME_ENV_VAR: socket.getfqdn(),
    }
    if self._num_configured_containers:
      env['NUM_CONFIGURED_CONTAINERS'] = self._num_configured_containers
    return env

  def set_num_configured_containers(self, num_configured_containers):
    self._num_configured_containers = num_configured_containers

  def create_container(self, container_desc, image_name, swarming_url, labels,
                       additional_env=None, **kwargs):
    container_workdir = '/b/%s' % container_desc.name
    container_volumes = self.volumes.copy()
    container_volumes[container_workdir] = {'bind': '/b/', 'mode': 'rw'}
    pw = pwd.getpwnam('chrome-bot')
    uid, gid = pw.pw_uid, pw.pw_gid
    if not os.path.exists(container_workdir):
      os.mkdir(container_workdir)
      os.chown(container_workdir, uid, gid)
    else: # pragma: no cover
      # TODO(bpastene): Remove this once existing workdirs everywhere have been
      # chown'ed.
      os.chown(container_workdir, uid, gid)
    env = self._get_env(swarming_url)
    if additional_env:
      env.update(additional_env)
    if sys.platform.startswith('linux') and os.path.exists(_KVM_DEVICE):
      # Allow the container access to the KVM device so it can run qemu-kvm.
      devices = ['{0}:{0}'.format(_KVM_DEVICE)]
      # Forward the tun device to the docker instance if it exists on the host.
      # By default, the docker provides Read-Write-Makenode permissions.
      if os.path.exists(_TUN_DEVICE):
        devices.append('{0}:{0}'.format(_TUN_DEVICE))
    else:
      devices = None
    new_container = self._client.containers.create(
        image=image_name,
        hostname=container_desc.hostname,
        volumes=container_volumes,
        environment=env,
        devices=devices,
        name=container_desc.name,
        detach=True,  # Don't block until it exits.
        labels=labels,
        **kwargs)
    new_container.start()
    container_desc.log_started()
    return new_container


class Container(object):
  def __init__(self, container):
    self._container = container
    self.name = container.name

  @property
  def labels(self):
    return self._container.attrs.get('Config', {}).get('Labels', {})

  @property
  def state(self):
    return self._container.attrs.get('State', {}).get('Status', 'unknown')

  @property
  def exit_code(self):
    return self._container.attrs.get('State', {}).get('ExitCode', 'unknown')

  @property
  def attrs(self):
    return self._container.attrs

  @property
  def image(self):
    return self._container.image

  def exec_run(self, cmd):
    res = self._container.exec_run(cmd)
    return six.ensure_text(res.output)

  def get_container_uptime(self, now):
    """Returns the containers uptime in minutes."""
    # Docker returns start time in format "%Y-%m-%dT%H:%M:%S.%f\d\d\dZ", so chop
    # off the last 4 digits to convert from nanoseconds to micoseconds
    start_time_string = self._container.attrs['State']['StartedAt'][:-4]
    start_time = datetime.strptime(start_time_string, '%Y-%m-%dT%H:%M:%S.%f')
    return ((now - start_time).total_seconds())/60

  def get_swarming_bot_pid(self):
    try:
      output = self.exec_run(
          'su chrome-bot -c "lsof -t /b/swarming/swarming.lck"').strip()
    except docker.errors.NotFound:
      logging.error('Docker engine returned 404 for container %s', self.name)
      return None
    if 'rpc error:' in output:
      logging.error(
          'Unable to get bot pid of %s: %s', self._container.name, output)
      return None
    try:
      return int(output)
    except ValueError:
      logging.error(
          'Unable to get bot pid of %s. Output of lsof: "%s"',
          self._container.name, output)
      return None

  def kill_swarming_bot(self, now=None, max_uptime=None):
    """Shuts down the swarming bot process in the container.

    If the bot can't safely be killed, the container will be stopped instead
    if its uptime exceeds max_uptime.

    Args:
      now: datetime.datetime representing 'now'. Passed in mostly for tests.
      max_uptime: Max uptime to allow for a container.
    """
    pid = self.get_swarming_bot_pid()
    if pid is not None:
      # The swarming bot process will capture this signal and shut itself
      # down at the next opportunity. Once the process exits, its container
      # will exit as well.
      try:
        self.exec_run('kill -15 %d' % pid)
      except docker.errors.APIError:  # pragma: no cover
        logging.exception('Unable to send SIGTERM to swarming bot.')
      else:
        logging.info('Sent SIGTERM to swarming bot of %s.', self.name)
    else:
      now = now or datetime.utcnow()
      uptime = self.get_container_uptime(now)
      if max_uptime and uptime <= 2 * max_uptime:
        logging.warning(
            'Unknown bot pid. Quitting early since container uptime (%d) is '
            'within grace period of %d.', uptime, 2 * max_uptime)
        return
      logging.error('Unknown bot pid. Stopping container.')
      try:
        self.stop()
      except requests.exceptions.ReadTimeout:
        logging.error('Timeout when stopping %s, force removing...', self.name)
        try:
          self.remove(force=True)
        except docker.errors.APIError:
          logging.exception(
              'Unable to remove %s. The docker engine is most likely stuck '
              'and will need a reboot.', self.name)
          raise FrozenContainerError()

  def pause(self):
    self._container.pause()

  def unpause(self):
    self._container.unpause()

  def stop(self, timeout=10):
    self._container.stop(timeout=timeout)

  def remove(self, force=False):
    self._container.remove(force=force)
