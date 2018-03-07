# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Implements export of builds from datastore to BigQuery."""

import datetime
import json
import logging
import re

from google.appengine.api import app_identity
from google.appengine.api import taskqueue
from google.appengine.ext import ndb
import webapp2

from components import decorators
from components import net
from components import utils
import bqh

import config
import model
import v2


# Mocked in tests.
@ndb.tasklet
def enqueue_pull_task_async(queue, payload):  # pragma: no cover
  task = taskqueue.Task(
      payload=payload,
      method='PULL')
  # Cannot just return add_async's return value because it is
  # a non-Future object and does not play nice with `yield fut1, fut2` construct
  yield task.add_async(queue_name=queue, transactional=True)


@ndb.tasklet
def enqueue_bq_export_async(build):  # pragma: no cover
  """Enqueues a pull task to export a completed build to BigQuery."""
  assert ndb.in_transaction()
  assert build
  assert build.status == model.BuildStatus.COMPLETED

  settings = yield config.get_settings_async()
  bucket_re = _compile_bucket_re(settings.bq_export)
  if bucket_re.match(build.bucket):
    yield enqueue_pull_task_async(
        'bq-export-experimental' if build.experimental else 'bq-export-prod',
        json.dumps({'id': build.key.id()}))


class CronExportBuilds(webapp2.RequestHandler):  # pragma: no cover
  """Exports builds to a BigQuery table."""

  queue_name = None
  dataset = None

  @decorators.require_cronjob
  def get(self):
    assert self.queue_name
    assert self.dataset

    settings = config.get_settings_async().get_result()
    _process_pull_task_batch(self.queue_name, self.dataset, settings.bq_export)


class CronExportBuildsProd(CronExportBuilds):
  queue_name = 'bq-export-prod'
  dataset = 'builds'


class CronExportBuildsExperimental(CronExportBuilds):
  queue_name = 'bq-export-experimental'
  dataset = 'builds_experimental'


def _process_pull_task_batch(queue_name, dataset, bq_settings):
  """Exports up to 1000 builds to BigQuery.

  Leases pull tasks, fetches build entities, tries to convert them to v2 format
  and insert into BigQuery in v2 format.

  Ignores a build if its bucket does not match any of bq_settings.buckets_re.

  If v2 conversion raises v2.errors.UnsupportedBuild, skips the build.
  If v2 conversion raises any other exception, including
  v2.errors.MalformedBuild, logs the exception and does not remove the task from
  the queue. Such a task will be retried later.

  If v2 conversion indicates that the build is not finalized and it has been
  20m or more since the build was completed, the following strategies apply:
  - if the build infra-failed with BOT_DIED or TIMED_OUT task status,
    saves build as is.
  - if the build infra-failed with BOOTSTRAPPER_ERROR and there are no steps,
    assumes the build failed to register LogDog prefix and saves it as is.
  - otherwise logs a warning/error, does not save to BigQuery and retries the
    task later.
  """
  now = utils.utcnow()

  # Parse settings.
  bucket_re = _compile_bucket_re(bq_settings)
  allowed_logdog_hosts = set(bq_settings.allowed_logdog_hosts)

  # Lease tasks.
  lease_duration = datetime.timedelta(minutes=5)
  lease_deadline = now + lease_duration
  q = taskqueue.Queue(queue_name)
  tasks = q.lease_tasks(lease_duration.total_seconds(), 1000)
  if not tasks:
    return

  build_ids = [json.loads(t.payload)['id'] for t in tasks]
  # IDs of builds that we could not save and want to retry later.
  ids_to_retry = set()

  # Fetch builds for the tasks and convert to v2 format.
  builds = ndb.get_multi([ndb.Key(model.Build, bid) for bid in build_ids])
  v2_builds_futs = [
    (bid, _build_to_v2_async(bid, b, bucket_re, allowed_logdog_hosts))
    for bid, b in zip(build_ids, builds)
  ]
  v2_builds = []
  for bid, fut in v2_builds_futs:
    v2_build, retry = fut.get_result()
    if retry:
      ids_to_retry.add(bid)
    elif v2_build:  # pragma: no branch
      v2_builds.append(v2_build)

  row_count = 0
  if v2_builds:
    not_inserted_ids = _export_builds(dataset, v2_builds, lease_deadline)
    row_count = len(v2_builds) - len(not_inserted_ids)
    ids_to_retry.update(not_inserted_ids)

  if ids_to_retry:
    logging.warning('will retry builds %r later', sorted(ids_to_retry))

  done_tasks = [
    t
    for bid, t in zip(build_ids, tasks)
    if bid not in ids_to_retry
  ]
  q.delete_tasks(done_tasks)
  logging.info(
      'inserted %d rows, processed %d tasks', row_count, len(done_tasks))


@ndb.tasklet
def _build_to_v2_async(bid, build, bucket_re, allowed_logdog_hosts):
  """Returns (v2_build, should_retry) tuple.

  Logs reasons for returning v2_build=None or retry=True.
  """
  if not build:
    logging.error('skipping build %d: not found', bid)
    raise ndb.Return(None, False)

  if build.status != model.BuildStatus.COMPLETED:
    logging.error('skipping build %d: not complete', bid)
    raise ndb.Return(None, False)

  if not bucket_re.match(build.bucket):
    logging.warning(
        'skipping build %d: bucket %r does not match %r',
        bid, build.bucket, bucket_re.pattern)
    raise ndb.Return(None, False)

  try:
    v2_build, finalized = yield v2.build_to_v2_async(
        build, allowed_logdog_hosts)
  except v2.errors.UnsupportedBuild as ex:
    logging.warning(
        'skipping build %d: not supported by v2 conversion: %s',
        bid, ex)
    raise ndb.Return(None, False)
  except Exception:
    logging.exception(
        'failed to convert build to v2\nBuild id: %d', bid)
    raise ndb.Return(None, True)

  if finalized:
    raise ndb.Return(v2_build, False)

  build_age = utils.utcnow() - build.complete_time
  if build_age >= datetime.timedelta(minutes=20):  # pragma: no branch
    task_state = _dict_get(
        build.result_details, 'swarming', 'task_result', 'state')
    if task_state in ('BOT_DIED', 'TIMED_OUT'):
      logging.warning(
          'build %d is not finalized and its task state is %r >=20m ago. '
          'Will not wait longer.', bid, task_state)
      raise ndb.Return(v2_build, False)

    infra_failure_type = _dict_get(
        build.result_details, 'build_run_result', 'infraFailure', 'type')
    if infra_failure_type == 'BOOTSTRAPPER_ERROR' and not v2_build.steps:
      logging.warning(
          'build %d is not finalized, has no steps and failed with '
          'BOOTSTRAPPER_ERROR error >=20m ago. Will not wait longer.', bid)
      raise ndb.Return(v2_build, False)

  logging.warning(
      'build %d is not finalized. Build age: %s', build.key.id(), build_age)
  if build_age > datetime.timedelta(hours=2):  # pragma: no branch
    # Poor man's monitoring.
    logging.error('build is not finalized for >=2h')
  raise ndb.Return(None, True)


def _export_builds(dataset, v2_builds, deadline):
  """Saves v2 builds to BigQuery.

  Logs insert errors and returns a list of ids of builds that could not be
  inserted.
  """
  table_name = 'completed_BETA'  # TODO(nodir): remove beta suffix.
  # BigQuery API doc:
  # https://cloud.google.com/bigquery/docs/reference/rest/v2/tabledata/insertAll
  logging.info('sending %d rows', len(v2_builds))
  res = net.json_request(
      url=(
        ('https://www.googleapis.com/bigquery/v2/'
          'projects/%s/datasets/%s/tables/%s/insertAll') % (
          app_identity.get_application_id(), dataset, table_name)
        ),
      method='POST',
      payload={
        'kind': 'bigquery#tableDataInsertAllRequest',
        'skipInvalidRows': False,
        'ignoreUnknownValues': False,
        'rows': [
          {
            'insertId': str(b.id),
            'json': bqh.message_to_dict(b),
          }
          for b in v2_builds
        ],
      },
      scopes=bqh.INSERT_ROWS_SCOPE,
      # deadline parameter here is duration in seconds.
      deadline=(deadline - utils.utcnow()).total_seconds(),
  )

  failed_ids = []
  for err in res.get('insertErrors', []):
    b = v2_builds[err['index']]
    failed_ids.append(b.id)
    logging.error(
        'failed to insert row for build %d: %r',
        b.id, err['errors'])
  return failed_ids


def _compile_bucket_re(bq_settings):
  return re.compile('|'.join(
      '(%s)' % r
      for r in bq_settings.buckets_re or ['^$']
  ))


def _dict_get(d, *keys):
  """Returns a value from nested dicts, e.g. _dict_get({a:{b:c}}, a, b) == c."""
  v = d
  for k in keys:
    v = (v or {}).get(k)
  return v
