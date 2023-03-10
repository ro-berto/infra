# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""API for calling 'cloudbuildhelper' tool.

See https://chromium.googlesource.com/infra/infra/+/main/build/images/.
"""

import contextlib
import re

from collections import namedtuple

from recipe_engine import recipe_api


# Version of `cloudbuildhelper` to use by default.
CBH_VERSION = 'git_revision:4693beef9e565b07b62beaaa3a1a72bf65479f1c'


class CloudBuildHelperApi(recipe_api.RecipeApi):
  """API for calling 'cloudbuildhelper' tool."""

  # Describes (potentially) multi-repo checkout on disk.
  CheckoutMetadata = namedtuple('CheckoutMetadata', [
      'root',  # a Path to the checkout root
      'repos', # a dict {"rel/path": {"repository": ..., "revision": ..."}}
  ])

  # Reference to a docker image uploaded to a registry.
  Image = namedtuple('Image', [
      'image',          # <registry>/<name>
      'digest',         # sha256:...
      'tag',            # its canonical tag, if any
      'context_dir',    # absolute path to a context directory passed to Docker
      'view_image_url', # link to GCR page, for humans, optional
      'view_build_url', # link to GCB page, for humans, optional
      'notify',         # a list of NotifyConfig tuples
      'sources',        # a list of dicts, see _process_sources
  ])

  # NotifyConfig identifies a system to notify about the image.
  NotifyConfig = namedtuple('NotifyConfig', [
      'kind',   # only "git" is supported currently
      'repo',   # a repository URL to checkout and update
      'script', # a script inside the repository to call to update the checkout
  ])

  # Reference to a tarball in GS produced by `upload`.
  Tarball = namedtuple('Tarball', [
      'name',      # name from the manifest
      'bucket',    # name of the GS bucket with the tarball
      'path',      # path within the bucket
      'sha256',    # hex digest
      'version',   # canonical tag
      'sources',   # a list of dicts, see _process_sources
  ])

  # Returned by a callback passed to do_roll.
  RollCL = namedtuple('RollCL', [
      'message',  # commit message
      'cc',       # a list of emails to CC on the CL
      'tbr',      # a list of emails to put in TBR= line
      'commit',   # if True, submit to CQ, if False, trigger CQ dry run only
  ])

  # A set of restrictions on what the manifest can target.
  #
  # All fields are lists of strings. Empty lists means "do not restrict".
  Restrictions = namedtuple('Restrictions', [
      'targets',             # prefixes of allowed target names
      'build_steps',         # allowed kinds of build steps
      'storage',             # prefixes of allowed tarball destinations
      'container_registry',  # allowed Container Registries
      'cloud_build',         # allowed Cloud Build projects
      'notifications',       # prefixes of allowed notification destinations
  ])

  # Used in place of Image to indicate that the image builds successfully, but
  # it wasn't uploaded anywhere.
  #
  # This happens if the target manifest doesn't specify a registry to upload
  # the image to. This is rare.
  NotUploadedImage = Image(None, None, None, None, None, None, None, None)

  def __init__(self, **kwargs):
    super(CloudBuildHelperApi, self).__init__(**kwargs)
    self._cbh_bin = None

  @property
  def command(self):
    """Path to the 'cloudbuildhelper' binary to use.

    If unset on first access, will bootstrap the binary from CIPD. Otherwise
    will return whatever was set previously (useful if 'cloudbuildhelper' is
    part of the checkout already).
    """
    if self._cbh_bin is None:
      self._cbh_bin = self.m.cipd.ensure_tool(
          'infra/tools/cloudbuildhelper/${platform}', CBH_VERSION)
    return self._cbh_bin

  @command.setter
  def command(self, val):
    """Can be used to tell the module to use an existing binary."""
    self._cbh_bin = val

  def report_version(self):
    """Reports the version of cloudbuildhelper tool via the step text.

    Returns:
      None.
    """
    res = self.m.step(
        name='cloudbuildhelper version',
        cmd=[self.command, 'version'],
        stdout=self.m.raw_io.output_text(),
        step_test_data=lambda: self.m.raw_io.test_api.stream_output_text(
            '\n'.join([
                'cloudbuildhelper v6.6.6',
                '',
                'CIPD package name: infra/tools/cloudbuildhelper/...',
                'CIPD instance ID:  lTJD7x...',
            ])),
    )
    res.presentation.step_text += '\n' + res.stdout

  def build(self,
            manifest,
            canonical_tag=None,
            build_id=None,
            infra=None,
            restrictions=None,
            labels=None,
            tags=None,
            checkout_metadata=None,
            step_test_image=None):
    """Calls `cloudbuildhelper build <manifest>` interpreting the result.

    Args:
      * manifest (Path) - path to YAML file with definition of what to build.
      * canonical_tag (str) - tag to push the image to if we built a new image.
      * build_id (str) - identifier of the CI build to put into metadata.
      * infra (str) - what section to pick from 'infra' field in the YAML.
      * restrictions (Restrictions) - restrictions to apply to manifests.
      * labels ({str: str}) - labels to attach to the docker image.
      * tags ([str]) - tags to unconditionally push the image to.
      * checkout_metadata (CheckoutMetadata) - to get revisions.
      * step_test_image (Image) - image to produce in training mode.

    Returns:
      Image instance or NotUploadedImage if the YAML doesn't specify a registry.

    Raises:
      StepFailure on failures.
    """
    name, _ = self.m.path.splitext(self.m.path.basename(manifest))

    cmd = [self.command, 'build', manifest]
    if canonical_tag:
      cmd += ['-canonical-tag', canonical_tag]
    if build_id:
      cmd += ['-build-id', build_id]
    if infra:
      cmd += ['-infra', infra]
    if restrictions:
      cmd += self._restriction_args(restrictions)
    for k in sorted(labels or {}):
      cmd += ['-label', '%s=%s' % (k, labels[k])]
    for t in (tags or []):
      cmd += ['-tag', t]
    cmd += ['-json-output', self.m.json.output()]

    # Expected JSON output (may be produced even on failures).
    #
    # {
    #   "error": "...",  # error message on errors
    #   "image": {
    #     "image": "registry/name",
    #     "digest": "sha256:...",
    #     "tag": "its-canonical-tag",
    #   },
    #   "view_image_url": "https://...",  # for humans
    #   "view_build_url": "https://...",  # for humans
    # }
    try:
      res = self.m.step(
          name='cloudbuildhelper build %s' % name,
          cmd=cmd,
          step_test_data=lambda: self.test_api.build_success_output(
              step_test_image, name, canonical_tag,
              checkout_metadata=checkout_metadata,
          ),
      )
      if not res.json.output:  # pragma: no cover
        res.presentation.status = self.m.step.FAILURE
        raise recipe_api.InfraFailure(
            'Call succeeded, but didn\'t produce -json-output')
      out = res.json.output
      if not out.get('image'):
        return self.NotUploadedImage
      return self.Image(
          image=out['image']['image'],
          digest=out['image']['digest'],
          tag=out['image'].get('tag'),
          context_dir=out.get('context_dir'),
          view_image_url=out.get('view_image_url'),
          view_build_url=out.get('view_build_url'),
          notify=[self._parse_notify_config(n) for n in out.get('notify', [])],
          sources=self._process_sources(out.get('sources'), checkout_metadata),
      )
    finally:
      self._make_build_step_pretty(self.m.step.active_result, tags)

  @staticmethod
  def _parse_notify_config(cfg):
    kind = cfg.get('kind')
    if kind != 'git':  # pragma: no cover
      raise recipe_api.InfraFailure('Unrecognized `notify` kind %r' % (kind,))
    return CloudBuildHelperApi.NotifyConfig(
        kind=kind,
        repo=cfg['repo'],
        script=cfg['script'],
    )

  @staticmethod
  def _restriction_args(r):
    out = []
    for s in sorted(r.targets or []):
      out.extend(['-restrict-targets', s])
    for s in sorted(r.build_steps or []):
      out.extend(['-restrict-build-steps', s])
    for s in sorted(r.storage or []):
      out.extend(['-restrict-storage', s])
    for s in sorted(r.container_registry or []):
      out.extend(['-restrict-container-registry', s])
    for s in sorted(r.cloud_build or []):
      out.extend(['-restrict-cloud-build', s])
    for s in sorted(r.notifications or []):
      out.extend(['-restrict-notifications', s])
    return out

  def _process_sources(self, sources, checkout_metadata):
    """Joins `sources` from -json-output with gclient checkout spec.

    `sources` look like: ["/root/a/b/c/d", "/root/a/b/c/e", "/root/a"].

    `checkout_metadata` looks like:
       root: Path("/root")
       repos:
          {
              "a": {
                  "repository": "https://example.com/a",
                  "revision": "aaaaa"
              },
              "a/b/c": {
                  "repository": "https://example.com/b",
                  "revision": "aaaaa"
              }
          }

    Returns something like:
       [
          {
              "repository": "https://example.com/b",
              "revision": "aaaaa",
              "sources": ["d", "e"]
          },
          {
              "repository": "https://example.com/a",
              "revision": "aaaaa",
              "sources": ["."]
          },
       ]
    """
    if not sources or not checkout_metadata:
      return []

    # Helper to convert Path to str. Note that join('.') is not allowed, so
    # we need to special-case it.
    render_path = lambda root, p: str(root.join(p) if p != '.' else root)

    # Build full paths to repos in the gclient checkout, ordering them by
    # "most nested first".
    checkout_paths = [
        (render_path(checkout_metadata.root, p), e['repository'], e['revision'])
        for p, e in checkout_metadata.repos.items()
    ]
    checkout_paths.sort(key=lambda e: len(e[0]), reverse=True)

    # Outputs are ordered based on `sources`.
    output_entries = []

    # For each emitted source path, find the corresponding gclient repo.
    for source_path in sources:
      for checkout_path, repository, revision in checkout_paths:
        is_under = (
            source_path == checkout_path or
            source_path.startswith(checkout_path + self.m.path.sep))
        if not is_under:
          continue

        # The path to `source_path` relative to the repository root.
        rel = source_path[len(checkout_path):]
        rel = rel.replace(self.m.path.sep, '/').strip('/')
        if rel == "":
          rel = "."

        # Find the existing repo entry in `output_entries` or create a new one.
        for entry in output_entries:
          if entry['repository'] == repository:
            entry['sources'].append(rel)
            break
        else:
          output_entries.append({
              'repository': repository,
              'revision': revision,
              'sources': [rel],
          })

        # Found the repo checkout, don't visit any parent directories which also
        # may be repo checkouts.
        break

    return output_entries

  @staticmethod
  def _make_build_step_pretty(r, tags):
    js = r.json.output
    if not js or not isinstance(js, dict):  # pragma: no cover
      return

    if js.get('view_image_url'):
      r.presentation.links['image'] = js['view_image_url']
    if js.get('view_build_url'):
      r.presentation.links['build'] = js['view_build_url']

    if js.get('error'):
      r.presentation.step_text += '\nERROR: %s' % js['error']
    elif js.get('image'):
      img = js['image']
      tag = img.get('tag') or (tags[0] if tags else None)
      if tag:
        ref = '%s:%s' % (img['image'], tag)
      else:
        ref = '%s@%s' % (img['image'], img['digest'])
      lines = [
          '',
          'Image: %s' % ref,
          'Digest: %s' % img['digest'],
      ]
      # Display all added tags (including the canonical one we got via `img`).
      for t in sorted(
          set((tags or []) + ([img['tag']] if img.get('tag') else []))):
        lines.append('Tag: %s' % t)
      r.presentation.step_text += '\n'.join(lines)
    else:
      r.presentation.step_text += '\nImage builds successfully'

  def upload(self,
             manifest,
             canonical_tag,
             build_id=None,
             infra=None,
             restrictions=None,
             checkout_metadata=None,
             step_test_tarball=None):
    """Calls `cloudbuildhelper upload <manifest>` interpreting the result.

    Args:
      * manifest (Path) - path to YAML file with definition of what to build.
      * canonical_tag (str) - tag to apply to a tarball if we built a new one.
      * build_id (str) - identifier of the CI build to put into metadata.
      * infra (str) - what section to pick from 'infra' field in the YAML.
      * restrictions (Restrictions) - restrictions to apply to manifests.
      * checkout_metadata (CheckoutMetadata) - to get revisions.
      * step_test_tarball (Tarball) - tarball to produce in training mode.

    Returns:
      Tarball instance.

    Raises:
      StepFailure on failures.
    """
    name = self.m.path.basename(self.m.path.splitext(manifest)[0])

    cmd = [self.command, 'upload', manifest, '-canonical-tag', canonical_tag]
    if build_id:
      cmd += ['-build-id', build_id]
    if infra:
      cmd += ['-infra', infra]
    if restrictions:
      cmd += self._restriction_args(restrictions)
    cmd += ['-json-output', self.m.json.output()]

    # Expected JSON output (may be produced even on failures).
    #
    # {
    #   "name": "<name from the manifest>",
    #   "error": "...",  # error message on errors
    #   "gs": {
    #     "bucket": "something",
    #     "name": "a/b/c/abcdef...tar.gz",
    #   },
    #   "sha256": "abcdef...",
    #   "canonical_tag": "<oldest tag>"
    # }
    try:
      res = self.m.step(
          name='cloudbuildhelper upload %s' % name,
          cmd=cmd,
          step_test_data=lambda: self.test_api.upload_success_output(
              step_test_tarball, name, canonical_tag,
              checkout_metadata=checkout_metadata,
          ),
      )
      if not res.json.output:  # pragma: no cover
        res.presentation.status = self.m.step.FAILURE
        raise recipe_api.InfraFailure(
            'Call succeeded, but didn\'t produce -json-output')
      out = res.json.output
      if 'gs' not in out:  # pragma: no cover
        res.presentation.status = self.m.step.FAILURE
        raise recipe_api.InfraFailure('No "gs" section in -json-output')
      return self.Tarball(
          name=out['name'],
          bucket=out['gs']['bucket'],
          path=out['gs']['name'],
          sha256=out['sha256'],
          version=out['canonical_tag'],
          sources=self._process_sources(out.get('sources'), checkout_metadata),
      )
    finally:
      self._make_upload_step_pretty(self.m.step.active_result)

  @staticmethod
  def _make_upload_step_pretty(r):
    js = r.json.output
    if not js or not isinstance(js, dict):  # pragma: no cover
      return

    if js.get('error'):
      r.presentation.step_text += '\nERROR: %s' % js['error']
      return

    # Note: '???' should never appear during normal operation. They are used
    # here defensively in case _make_upload_step_pretty is called due to
    # malformed JSON output.
    r.presentation.step_text += '\n'.join([
        '',
        'Name: %s' % js.get('name', '???'),
        'Version: %s' % js.get('canonical_tag', '???'),
        'SHA256: %s' % js.get('sha256', '???'),
    ])

  def update_pins(self, path):
    """Calls `cloudbuildhelper pins-update <path>`.

    Updates the file at `path` in place if some docker tags mentioned there have
    moved since the last pins update.

    Args:
      * path (Path) - path to a `pins.yaml` file to update.

    Returns:
      List of strings with updated "<image>:<tag>" pairs, if any.
    """
    res = self.m.step(
        'cloudbuildhelper pins-update',
        [
            self.command, 'pins-update', path,
            '-json-output', self.m.json.output(),
        ],
        step_test_data=lambda: self.test_api.update_pins_output(
            updated=['some_image:tag'],
        ),
    )
    return res.json.output.get('updated') or []

  def discover_manifests(self, root, entries, test_data=None):
    """Returns a list with paths to all manifests we need to build.

    Each entry is either a directory to scan recursively, or a reference to
    a concrete manifest (if ends with ".yaml").

    Args:
      * root (Path) - gclient solution root.
      * entries ([str]) - paths relative to the solution root to scan.
      * test_data ([str]) - paths to put into each `dirs` in training mode.

    Returns:
      [Path].
    """
    paths = []
    for entry in entries:
      if entry.endswith('.yaml'):
        paths.append(root.join(*entry.split('/')))
      else:
        paths.extend(self.m.file.glob_paths(
            'list %s' % entry,
            root.join(entry),
            '**/*.yaml',
            test_data=test_data if test_data is not None else ['target.yaml']))
    return paths

  def do_roll(self, repo_url, root, callback, ref='main'):
    """Checks out a repo, calls the callback to modify it, uploads the result.

    Args:
      * repo_url (str) - repo to checkout.
      * root (Path) - where to check it out too (can be a cache).
      * callback (func(Path)) - will be called as `callback(root)` with cwd also
          set to `root`. It can modify files there and either return None to
          skip the roll or RollCL to attempt the roll. If no files are modified,
          the roll will be skipped regardless of the return value.
      * ref (str) - a ref to update (e.g. "main").

    Returns:
      * (None, None) if didn't create a CL (because nothing has changed).
      * (Issue number, Issue URL) if created a CL.
    """
    self.m.git.checkout(repo_url, ref=ref, dir_path=root, submodules=False)

    with self.m.context(cwd=root):
      self.m.git('branch', '-D', 'roll-attempt', ok_ret=(0, 1))
      self.m.git('checkout', '-t', 'origin/'+ref, '-b', 'roll-attempt')

      # Let the caller modify files in root.
      verdict = callback(root)
      if not verdict:  # pragma: no cover
        return None, None

      # Stage all added and deleted files to be able to `git diff` them.
      self.m.git('add', '.')

      # Check if we actually updated something.
      diff_check = self.m.git('diff', '--cached', '--exit-code', ok_ret='any')
      if diff_check.retcode == 0:  # pragma: no cover
        return None, None

      self.m.git('commit', '-m', verdict.message)

      # Skip uploading if in the dry run mode.
      if not verdict.commit:
        return None, None

      # Upload as a CL.
      self.m.git_cl.upload(verdict.message, name='git cl upload', upload_args=[
          '--force',        # skip asking for description, we already set it
          '--bypass-hooks', # we may not have the full checkout
      ] + [
          '--reviewers=%s' % tbr for tbr in sorted(set(verdict.tbr or []))
      ] + [
          '--cc=%s' % cc for cc in sorted(set(verdict.cc or []))
      ] +(['--set-bot-commit', '--use-commit-queue'] if verdict.commit else []))

      # Put a link to the uploaded CL.
      step = self.m.git_cl(
          'issue', ['--json', self.m.json.output()],
          name='git cl issue',
          step_test_data=lambda: self.m.json.test_api.output({
              'issue': 123456789,
              'issue_url': 'https://chromium-review.googlesource.com/c/1234567',
          }),
      )
      out = step.json.output
      step.presentation.links['Issue %s' % out['issue']] = out['issue_url']

      # TODO(vadimsh): Babysit the CL until it lands or until timeout happens.
      # Without this if images_builder runs again while the previous roll is
      # still in-flight, it will produce a duplicate roll (which will eventually
      # fail due to merge conflicts).

      return out['issue'], out['issue_url']

  def get_version_label(self,
                        path,
                        revision,
                        ref=None,
                        commit_position=None,
                        template=None):
    """Computes a version string identifying a commit.

    To calculate a commit position it either uses `Cr-Commit-Position` footer,
    if available, or falls back to `git number <rev>`.

    Uses the given `template` string to format the label. It is a Python format
    string, with following placeholders available:

      {rev}: a string with the short git revision hash.
      {ref}: a string with the last component of the commit ref (e.g. 'main').
      {cp}: an integer with the commit position number.
      {date}: "YYYY.MM.DD" UTC date when the build started.
      {build}: an integer build number or 0 if not available.

    Defaults to `{cp}-{rev}`.

    This label is used as a version name for artifacts produced based on this
    checked out commit.

    Args:
      * path (Path) - path to the git checkout root.
      * revision (str) - checked out revision.
      * ref (str) - checked out git ref if known.
      * commit_position (str) - `Cr-Commit-Position` footer value if available.
      * template (str) - a Python format string to use to format the version.

    Returns:
      A version string.
    """
    if commit_position:
      cp_ref, cp = self.m.commit_position.parse(commit_position)
      if cp_ref and not ref:
        ref = cp_ref
    else:
      with self.m.context(cwd=path, env={'CHROME_HEADLESS': '1'}):
        with self.m.depot_tools.on_path():
          result = self.m.git(
              'number', revision,
              name='get commit label',
              stdout=self.m.raw_io.output(),
              step_test_data=(
                  lambda: self.m.raw_io.test_api.stream_output('11112\n')))
      cp = int(result.stdout.strip())

    return (template or '{cp}-{rev}').format(
        rev=revision[:7],
        ref='unknown' if not ref else ref[ref.rfind('/')+1:],
        cp=cp,
        date=self.m.time.utcnow().strftime('%Y.%m.%d'),
        build=self.m.buildbucket.build.number,
    )

  @contextlib.contextmanager
  def build_environment(self,
                        root,
                        go_version_file=None,
                        nodejs_version_file=None):
    """A context manager that activates the build environment.

    Used to build code in a standalone git repositories that don't have Go
    or Node.js available via some other mechanism (like via gclient DEPS).

    It reads the Golang version from `<root>/<go_version_file>` and Node.js
    version from `<root>/<nodejs_version_file>`, bootstraps the corresponding
    versions in cache directories, adjusts PATH and other environment variables
    and yields to the user code.

    Let's assume the requested Go version is '1.16.10' and Node.js version
    is '16.13.0', then this call will use following cache directories:
      * `go1_16_10`: to install Go under.
      * `gocache`: for Go cache directories.
      * `nodejs16_13_0`: to install Node.js under.
      * `npmcache`: for NPM cache directories.

    For best performance the builder must map these directories as named
    caches using e.g.

        luci.builder(
            ...
            caches = [
                swarming.cache("go1_16_10"),
                swarming.cache("gocache"),
                swarming.cache("nodejs16_13_0"),
                swarming.cache("npmcache"),
            ],
        )

    Args:
      root (Path) - path to the checkout root.
      go_version_file (str) - path within the checkout to a text file with Go
          version to bootstrap or None to skip bootstrapping Go.
      nodejs_version_file (str) - path within the checkout to a text file with
          Node.js version to bootstrap or None to skip bootstrapping Node.js.
    """
    @contextlib.contextmanager
    def bootstrap(tool, module, version_file):
      if version_file:
        version = self.m.file.read_text(
            'read %s' % version_file,
            root.join(version_file),
            test_data='6.6.6\n',
        ).strip().lower()
        if not re.match(r'^\d+\.\d+\.\d+$', version):  # pragma: no cover
          raise ValueError('Bad %s version number %r' % (tool, version))
        install_path = self.m.path['cache'].join(
            '%s%s' % (tool, version.replace('.', '_')))
        with module(version, path=install_path):
          yield
      else:
        yield

    with bootstrap('go', self.m.golang, go_version_file):
      with bootstrap('nodejs', self.m.nodejs, nodejs_version_file):
        yield
