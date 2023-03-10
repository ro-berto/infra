# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from recipe_engine import post_process
from recipe_engine.recipe_api import Property

from PB.go.chromium.org.luci.buildbucket.proto.build import Build
from PB.go.chromium.org.luci.buildbucket.proto.common import (FAILURE,
                                                              GerritChange)
from PB.recipe_engine.result import RawResult

PYTHON_VERSION_COMPATIBILITY = 'PY2+3'

DEPS = [
  'recipe_engine/buildbucket',
  'recipe_engine/context',
  'recipe_engine/file',
  'recipe_engine/json',
  'recipe_engine/path',
  'recipe_engine/properties',
  'recipe_engine/raw_io',
  'recipe_engine/step',

  'depot_tools/bot_update',
  'depot_tools/gclient',
  'depot_tools/git',
  'depot_tools/tryserver',
]


PROPERTIES = {
  'upstream_id': Property(
      kind=str,
      help='ID of the project to patch'),
  'upstream_url': Property(
      kind=str,
      help='URL of git repo of the upstream project'),

  'downstream_id': Property(
      kind=str,
      help=('ID of the project that includes |upstream_id| in its recipes.cfg '
            'to be tested with upstream patch')),
  'downstream_url': Property(
      kind=str,
      help='URL of the git repo of the downstream project'),
}


NONTRIVIAL_ROLL_FOOTER = 'Recipe-Nontrivial-Roll'
MANUAL_CHANGE_FOOTER = 'Recipe-Manual-Change'
BYPASS_FOOTER = 'Recipe-Tryjob-Bypass-Reason'
RELEVANT_CONFLICTING_FOOTERS = [
    NONTRIVIAL_ROLL_FOOTER, MANUAL_CHANGE_FOOTER, BYPASS_FOOTER
]

FOOTER_ADD_TEMPLATE = '''

Add

    {footer}: {down_id}

To your CL message.

'''

MANUAL_CHANGE_MSG = '''
This means that your upstream CL (this one) will require MANUAL CODE CHANGES
in the downstream repo {down_id!r}. Best practice is to prepare all downstream
changes before landing the upstream CL, using:

    {down_id}/{down_recipes} -O {up_id}=/path/to/local/{up_id} test train

When that CL has been reviewed, you can land this upstream change. Once the
upstream change lands, roll it into your downstream CL:

    {down_id}/recipes.py manual_roll   # may require running multiple times.

Re-train expectations and upload the expectations plus the roll to your
downstream CL. It's customary to copy the outputs of manual_roll to create
a changelog to attach to the downstream CL as well to help reviewers understand
what the roll contains.
'''.strip()

NONTRIVIAL_CHANGE_MSG = '''
This means that your upstream CL (this one) will change the EXPECTATION FILES
in the downstream repo {down_id!r}.

The recipe roller will automatically prepare the non-trivial CL and will upload
it with `git cl upload --r-owners` to the downstream repo. Best practice is to
review this non-trivial roll CL to ensure that the expectations you see there
are expected.
'''

EXTRA_MSG = {
  NONTRIVIAL_ROLL_FOOTER: NONTRIVIAL_CHANGE_MSG,
  MANUAL_CHANGE_FOOTER: MANUAL_CHANGE_MSG,
}

MAIN_REF = 'refs/remotes/origin/main'


class RecipeTrainingFailure(Exception):
  pass


class RecipesRepo(object):
  """An abstraction of a recipes project to encapsulate common interactions."""

  def __init__(self, api, workdir_base, name, url):
    """
    Args:
      api (RecipeApi): The recipe api for this build.
      workdir_base (Path): The global directory for all recipe repo checkouts.
      name (str): See `name` property.
      url (str): The remote URL for this repo.
    """
    self._api = api
    self._workdir = workdir_base.join(name)
    self._name = name
    self._url = url

    self._root = None
    self._cl_revision = None
    self._recipes_py = None

  @property
  def name(self):
    """The name of this recipes project, e.g. 'recipe_engine'."""
    return self._name

  @property
  def root(self):
    """The absolute path to the root of the checkout for this repo.

    Will be None until `clone()` is called.
    """
    return self._root

  @property
  def recipes_py(self):
    """The path to the recipes.py file for this repo, relative to the root."""
    if self._recipes_py is None:
      recipes_cfg = self._api.file.read_json(
          'parse recipes.cfg',
          self._root.join('infra', 'config', 'recipes.cfg'),
          test_data={
            'recipes_path': 'some/path',
          })
      self._recipes_py = self._api.path.join(
          recipes_cfg.get('recipes_path', ''), 'recipes.py')
    return self._recipes_py

  def checkout_cl(self, base):
    """Sync the repo the CL that triggered this build.

    Assumes this repo is the repo for the CL.
    """
    assert self._cl_revision
    self.checkout(base, 'checkout base for cherry-pick')
    with self._api.context(cwd=self.root):
      try:
        self._api.git(
            'cherry-pick',
            self._cl_revision,
            name='cherry-pick CL onto %s' % base)
      except self._api.step.StepFailure:
        self._api.git('cherry-pick', '--abort')
        raise

  def checkout_default_ref(self):
    """Sync the repo to default ref (main)."""
    self.checkout(MAIN_REF, 'sync %s to %s' % (self.name, MAIN_REF))

  def checkout(self, checkout_ref, step_name):
    """Check out the specified ref."""
    with self._api.context(cwd=self.root):
      # Clean out those stale pyc's!
      self._api.git('clean', '-xf')
      return self._api.git('checkout', '-f', checkout_ref, name=step_name)

  def clear_diffs(self, step_name):
    """Throw away all unstaged changes to this repo."""
    with self._api.context(cwd=self.root):
      return self._api.git('checkout', '--', '.', name=step_name)

  def clone(self):
    """Clones the repo into a subdirectory of _workdir.

    Sets `root` to the root of the checkout, and `_cl_revision` to the
    revision for the CL if this repo is the one that triggered the current
    build.
    """
    assert not self._root, 'checkout already initialized'

    is_triggering_repo = self._url == self._api.tryserver.gerrit_change_repo_url

    self._api.file.ensure_directory(
        '%s checkout' % self._name, self._workdir)

    gclient_config = self._api.gclient.make_config()
    gclient_config.got_revision_reverse_mapping['got_revision'] = self._name
    soln = gclient_config.solutions.add()
    soln.name = self._name
    soln.url = self._url

    with self._api.context(cwd=self._workdir):
      ret = self._api.bot_update.ensure_checkout(
          gclient_config=gclient_config,
          assert_one_gerrit_change=False,
          # Only try to checkout the CL if this repo is the one that triggered
          # the current build.
          patch=is_triggering_repo)
      self._root = self._workdir.join(ret.json.output['root'])

      if is_triggering_repo:
        with self._api.context(cwd=self._root):
          rev_parse_step = self._api.git(
              'rev-parse',
              'FETCH_HEAD',
              name='read CL revision',
              stdout=self._api.raw_io.output_text(),
              step_test_data=lambda: self._api.raw_io.test_api.
              stream_output_text('deadbeef'))
          self._cl_revision = rev_parse_step.stdout.strip()

  def is_dirty(self, name):
    """Check whether the repo has any unstaged changes.

    Specifically, this can be used after calling `train()` to determine
    whether the training caused a change in the expectation files.
    """
    with self._api.context(cwd=self._root):
      # This has the benefit of showing the expectation diff to the user.
      diff_step = self._api.git('diff', '--exit-code', name=name, ok_ret='any')
      dirty = diff_step.retcode != 0
      if dirty:
        diff_step.presentation.status = 'FAILURE'
      return dirty

  def train(self, upstream_repo, step_name):
    """Re-trains the expectation files for this repo.

    Args:
      upstream_repo (RecipeRepo): A locally checked-out recipes repo that's
        among the dependencies of `self`. The training will be run using the
        local version of the dependency rather than the version pinned in
        recipes.cfg.
      step_name (str): The name to use for the training step.

    Raises:
      `RecipeTrainingFailure` if the training produces an uncaught exception.
    """
    try:
      cmd = [
          'python3',
          self._root.join(self.recipes_py),
          '-O',
          '%s=%s' % (upstream_repo.name, upstream_repo.root),
          'test',
          'train',
          '--no-docs',
      ]
      return self._api.step(step_name, cmd)
    except self._api.step.StepFailure:
      # Train recipes again as py3 tests only compare the on-disk expectation
      # files for PY2+3 recipes.
      # TODO(crbug.com/1147793): Remove it after Py3 migration is fully done.
      try:
        return self._api.step(step_name + ' (py3 retrain)', cmd)
      except self._api.step.StepFailure:
        raise RecipeTrainingFailure('failed to train recipes')


def _find_footer(api, repo_id):
  all_footers = api.tryserver.get_footers()

  if BYPASS_FOOTER in all_footers:
    api.step.empty(
        'BYPASS ENABLED',
        step_text='Roll tryjob bypassed for %r' % (
            # It's unlikely that there's more than one value, but just in case.
            ', '.join(all_footers[BYPASS_FOOTER]),))
    return None, True

  found_set = set()
  for name in RELEVANT_CONFLICTING_FOOTERS:
    values = all_footers.get(name, ())
    if repo_id in values:
      found_set.add(name)

  if len(found_set) > 1:
    api.step.empty(
        'Too many footers for %r' % (repo_id,),
        status=api.step.FAILURE,
        step_text='Found incompatible footers in CL message:\n' +
        ('\n'.join(' * ' + f for f in sorted(found_set))))

  return found_set.pop() if found_set else None, False


def _get_upstream_change(api, upstream_url):
  gs_suffix = '-review.googlesource.com'
  for cl in api.buildbucket.build.input.gerrit_changes:
    if not cl.host.endswith(gs_suffix):
      continue
    git_host = '%s.googlesource.com' % cl.host[:-len(gs_suffix)]
    if upstream_url == 'https://%s/%s' % (git_host, cl.project):
      return cl
  return None


def _get_expected_footer(api, upstream_repo, downstream_repo):
  # Run a 'train' on the downstream repo, first using the upstream repo at
  # main (more accurately, the latest non-crashing ancestor of main) and
  # then at the CL revision. We compare these two runs to avoid taking unrolled
  # CLs into account in the resulting diff.
  #
  # If the CL train fails, we require a Manual-Change footer
  # If there's a diff between the two trains, we require a Nontrivial-Roll
  # footer
  # If there's no diff between the two trains, we require no footers

  with api.step.nest('initialize checkouts'):
    upstream_repo.clone()
    downstream_repo.clone()

  upstream_repo.checkout_default_ref()

  last_non_crashing_revision = None

  # Starting from the tip of the upstream main branch, go back in history
  # until we find a commit that we can train the downstream repo against
  # without a crash. We'll use the downstream diff caused by this commit to the
  # diff caused by the current CL to determine whether the CL is trivial.
  with api.step.nest('find last non-crashing upstream revision'):
    # TODO(olivernewman): Only go back as far as the upstream revision
    # currently pinned in the downstream repo, instead of going back an
    # arbitrary number of commits. This should always be a valid git revision,
    # so we can also get rid of the "checkout HEAD~" error handling.
    for ancestor_index in range(10):
      ref = 'main~%d' % ancestor_index if ancestor_index else 'main'
      # Check out the parent commit to train against (except for the first
      # iteration of the for loop).
      if ancestor_index:
        try:
          upstream_repo.checkout('HEAD~', 'checkout upstream at %s' % ref)
        except api.step.StepFailure:
          # We've reached the beginning of the upstream repo's git history, so
          # give up trying to find a non-crashing revision.
          raise api.step.InfraFailure(
              'no non-crashing revision in upstream git history')
      try:
        downstream_repo.train(
            upstream_repo, 'train recipes at upstream %s' % ref)
      except RecipeTrainingFailure:
        # Clear diffs and move onto the previous upstream commit.
        downstream_repo.clear_diffs('clear diffs from upstream %s' % ref)
      else:
        api.step('using upstream %s as base' % ref, None)
        with api.context(cwd=upstream_repo.root):
          rev_parse_step = api.git(
              'rev-parse',
              'HEAD',
              name='get upstream base revision',
              stdout=api.raw_io.output_text(),
              step_test_data=lambda: api.raw_io.test_api.stream_output_text(
                  'abcd'))
          last_non_crashing_revision = rev_parse_step.stdout.strip()
        break

  with api.context(cwd=downstream_repo.root):
    api.git('add', '--all', name='save post-train downstream diff')

  # If we failed to find a non-crashing revision, just cherry-pick onto main.
  base = last_non_crashing_revision or MAIN_REF

  try:
    upstream_repo.checkout_cl(base=base)
  except api.step.StepFailure:
    # If this CL doesn't cherry-pick cleanly on top of the last non-crashing
    # revision, we'll fall back to cherry-picking on top of main. Training
    # after this will probably fail unless the CL actually fixes the crash.
    upstream_repo.checkout_cl(base=MAIN_REF)

  try:
    downstream_repo.train(upstream_repo, 'train recipes at upstream CL')
  except RecipeTrainingFailure:
    return MANUAL_CHANGE_FOOTER

  # If the training passes, we can use the diff between training against
  # upstream HEAD and training against the upstream CL to more accurately
  # determine if this CL introduces diffs, while disregarding diffs introduced
  # by older unrolled changes.
  trivial = not downstream_repo.is_dirty('post-train diff at upstream CL')
  return NONTRIVIAL_ROLL_FOOTER if not trivial else None


def _failing_step(api, name, message):
  result = api.step(name, [])
  result.presentation.status = api.step.FAILURE
  result.presentation.step_text = message

  raw_result = RawResult()
  raw_result.status = FAILURE
  raw_result.summary_markdown = message

  return raw_result


def RunSteps(api, upstream_id, upstream_url, downstream_id, downstream_url):
  # NOTE: this recipe is only useful as a tryjob with patch applied against
  # upstream repo, which means upstream_url must at least match a change from
  # api.buildbucket.build.input.gerrit_changes. upstream_url remains as a
  # required parameter for symmetric input for upstream/downstream.
  change = _get_upstream_change(api, upstream_url)
  assert change is not None, 'no change found from upstream repo'
  api.tryserver.set_change(change)

  # TODO: figure out upstream_id from downstream's repo recipes.cfg file using
  # patch and deprecated both upstream_id and upstream_url parameters.
  workdir_base = api.path['cache'].join('builder')

  upstream_repo = RecipesRepo(
    api, workdir_base, upstream_id, upstream_url)
  downstream_repo = RecipesRepo(
    api, workdir_base, downstream_id, downstream_url)

  # First, check to see if the user has bypassed this tryjob's analysis
  # entirely.
  actual_footer, bypass = _find_footer(api, downstream_id)
  if bypass:
    return

  expected_footer = _get_expected_footer(api, upstream_repo, downstream_repo)

  # Either expected_footer and actual_footer are both None or both matching
  # footers.
  if expected_footer == actual_footer:
    if expected_footer:
      msg = (
        'CL message contains correct footer (%r) for this repo.'
      ) % expected_footer
    else:
      msg = 'CL is trivial and message contains no footers for this repo.'
    api.step.empty('Roll OK', step_text=msg)
    return

  # trivial roll, but user has footer in CL message.
  if expected_footer is None and actual_footer is not None:
    return _failing_step(
        api, 'UNEXPECTED FOOTER IN CL MESSAGE',
        'Change is trivial, but found %r footer' % (actual_footer,))

  # nontrivial/manual roll, but user has wrong footer in CL message.
  if expected_footer is not None and actual_footer is not None:
    return _failing_step(
        api, 'WRONG FOOTER IN CL MESSAGE',
        'Change requires %r, but found %r footer' % (
            expected_footer,
            actual_footer,
        ))

  # expected != None at this point, so actual_footer must be None
  msg = FOOTER_ADD_TEMPLATE + EXTRA_MSG[expected_footer]
  return _failing_step(
      api, 'MISSING FOOTER IN CL MESSAGE',
      msg.format(
          footer=expected_footer,
          up_id=upstream_id,
          down_id=downstream_id,
          down_recipes=downstream_repo.recipes_py,
      ))


def GenTests(api):
  REPO_URLS = {
      'build':
          'https://chromium.googlesource.com/chromium/tools/build',
      'depot_tools':
          'https://chromium.googlesource.com/chromium/tools/depot_tools',
      'recipe_engine':
          'https://chromium.googlesource.com/infra/luci/recipes-py',
  }

  CL_TANGENTIAL = GerritChange(
      host='chromium-review.googlesource.com',
      project='chromium/src',
      change=111111,
      patchset=1,
  )

  CL_CHANGE_ACTUAL = GerritChange(
      host='chromium-review.googlesource.com',
      project='infra/luci/recipes-py',
      change=456789,
      patchset=12,
  )

  # GerritChange with bad host name.
  CL_UNSUPPORTED = GerritChange(
      host='chromium-review.foo-source.com',
      project='chromium/src',
      change=0,
      patchset=0,
  )

  def test_with_changes(name, upstream_id, downstream_id, changes, *footers):
    # Depending on the size of api.buildbucket.build.input.gerrit_changes,
    # the source for initialize step may differ.
    initialize = api.buildbucket.build(
        Build(id=1234567, input=Build.Input(gerrit_changes=changes)))
    if len(changes) == 1:
      initialize = api.buildbucket.try_build(
          git_repo=REPO_URLS[upstream_id],
          change_number=changes[0].change,
          patch_set=changes[0].patchset)
    return (api.test(name) + api.properties(
        upstream_id=upstream_id,
        upstream_url=REPO_URLS[upstream_id],
        downstream_id=downstream_id,
        downstream_url=REPO_URLS[downstream_id]) + initialize +
            api.override_step_data(
                'gerrit changes',
                api.json.output([{
                    'revisions': {
                        'deadbeef': {
                            '_number': 12,
                            'commit': {
                                'message': ''
                            }
                        },
                    }
                }])) + api.step_data(
                    'parse description',
                    api.json.output({
                        k: ['Reasons' if k == BYPASS_FOOTER else downstream_id]
                        for k in footers
                    })))

  def test(name, *footers):
    return test_with_changes(name, 'recipe_engine', 'depot_tools',
                             [CL_CHANGE_ACTUAL], *footers)

  yield (
    test('find_trivial_roll')
  )

  yield (
    test('bypass', BYPASS_FOOTER)
    + api.post_check(lambda check, steps: check('BYPASS ENABLED' in steps))
  )

  yield (
    test('too_many_footers', MANUAL_CHANGE_FOOTER, NONTRIVIAL_ROLL_FOOTER)
    + api.post_check(lambda check, steps: check(
        "Too many footers for 'depot_tools'" in steps
    ))
  )

  yield (
    test('find_trivial_roll_unexpected', MANUAL_CHANGE_FOOTER)
    + api.post_check(lambda check, steps: check(
        'UNEXPECTED FOOTER IN CL MESSAGE' in steps
    ))
  )

  yield (
      test('find_manual_roll_missing') +
      api.step_data('train recipes at upstream CL', retcode=1) +
      api.step_data('train recipes at upstream CL (py3 retrain)', retcode=1) +
      api.post_check(lambda check, steps: check(MANUAL_CHANGE_FOOTER in steps[
          'MISSING FOOTER IN CL MESSAGE'].step_text)))

  yield (
      test('find_manual_roll_wrong', NONTRIVIAL_ROLL_FOOTER) +
      api.step_data('train recipes at upstream CL', retcode=1) +
      api.step_data('train recipes at upstream CL (py3 retrain)', retcode=1) +
      api.post_check(lambda check, steps: check(MANUAL_CHANGE_FOOTER in steps[
          'WRONG FOOTER IN CL MESSAGE'].step_text)))

  yield (
    test('find_non_trivial_roll')
    + api.step_data('post-train diff at upstream CL', retcode=1)
    + api.post_check(lambda check, steps: check(
      NONTRIVIAL_ROLL_FOOTER in steps['MISSING FOOTER IN CL MESSAGE'].step_text
    ))
  )

  yield (
    test('trivial_roll_unrolled_changes')
    + api.step_data('post-train diff at upstream CL', retcode=1)
  )

  yield (
    test('nontrivial_roll_match', NONTRIVIAL_ROLL_FOOTER)
    + api.step_data('post-train diff at upstream CL', retcode=1)
  )

  # The current upstream main causes a crash in the downstream repo, but the
  # trivial CL fixes that.
  yield (test('trivial_roll_upstream_main_broken') + api.step_data(
      'find last non-crashing upstream revision'
      '.train recipes at upstream main',
      retcode=1) + api.step_data(
          'find last non-crashing upstream revision'
          '.train recipes at upstream main (py3 retrain)',
          retcode=1))

  # The current upstream main causes a crash in the downstream repo, but the
  # trivial CL fixes that.
  yield (test('trivial_roll_upstream_main_broken_rebase_fails') + api.step_data(
      'find last non-crashing upstream revision'
      '.train recipes at upstream main',
      retcode=1) + api.step_data(
          'find last non-crashing upstream revision'
          '.train recipes at upstream main (py3 retrain)',
          retcode=1) + api.step_data(
              'find last non-crashing upstream revision'
              '.get upstream base revision',
              stdout=api.raw_io.output_text('deadbeef')) +
         api.step_data('cherry-pick CL onto deadbeef', retcode=1))

  # None of the ancestor commits of the upstream repo's main branch are
  # compatible with the downstream repo (they all cause crashes).
  yield (
      test('no_good_upstream_main_commits') + api.step_data(
          'find last non-crashing upstream revision'
          '.train recipes at upstream main',
          retcode=1) + api.step_data(
              'find last non-crashing upstream revision'
              '.train recipes at upstream main (py3 retrain)',
              retcode=1)
      # Checkout fails -> we've hit the beginning of the upstream repo's git
      # history without finding a working commit.
      + api.step_data(
          'find last non-crashing upstream revision.'
          'checkout upstream at main~1',
          retcode=1) + api.post_process(post_process.StatusException))

  yield (test_with_changes('multiple_patchsets', 'recipe_engine', 'depot_tools',
                           [CL_TANGENTIAL, CL_CHANGE_ACTUAL], BYPASS_FOOTER) +
         api.step_data('parse description',
                       api.json.output({BYPASS_FOOTER: ['Reasons']})))

  yield (api.test('crash_for_no_upstream_patchset') + api.buildbucket.build(
      Build(id=1234567, input=Build.Input(gerrit_changes=[CL_UNSUPPORTED]))) +
         api.properties(
             upstream_id='recipe_engine',
             upstream_url=REPO_URLS['recipe_engine'],
             downstream_id='depot_tools',
             downstream_url=REPO_URLS['depot_tools'],
         ) + api.expect_exception('AssertionError'))
