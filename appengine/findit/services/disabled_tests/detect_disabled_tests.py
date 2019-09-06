# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
import os
import re

from google.appengine.ext import ndb

from common.swarmbucket import swarmbucket
from gae_libs import appengine_util
from libs import time_util
from model.flake.flake import Flake
from model.flake.flake import FlakeIssue
from model.test_inventory import LuciTest
from services import bigquery_helper
from services import step_util

_DEFAULT_LUCI_PROJECT = 'chromium'

_DEFAULT_CONFIG = 'Unknown'

_ISSUE_LINK_REGEX = [
    re.compile(r'^(?:https?://)?crbug.com/([0-9]+)$'),
    re.compile(
      r'^(?:https?://)?bugs.chromium.org/p/chromium/issues/detail\?id=([0-9]+)$'
    )
]

_MEMORY_FLAGS_REGEX = [
    (re.compile('ASan', re.I), 'ASan:True'),
    (re.compile('LSan', re.I), 'LSan:True'),
    (re.compile('MSan', re.I), 'MSan:True'),
    (re.compile('TSan', re.I), 'TSan:True'),
    (re.compile('UBSan', re.I), 'UBSan:True'),
]


def _GetQueryParameters():
  return [
      bigquery_helper.GenerateArrayQueryParameter(
          'supported_masters', 'STRING',
          swarmbucket.GetMasters('luci.chromium.ci') +
          swarmbucket.GetMasters('luci.chromium.try'))
  ]


def _ExecuteQuery(parameters=None):

  def GetQuery():
    path = os.path.realpath(
        os.path.join(__file__, os.path.pardir, 'disabled_tests.sql'))
    with open(path) as f:
      query = f.read()
    return query

  query = GetQuery()
  local_tests = {}
  total_rows = 0
  for row in bigquery_helper.QueryResultIterator(
      appengine_util.GetApplicationId(), query, parameters=parameters):
    total_rows += 1
    _CreateLocalTests(row, local_tests)

  assert total_rows > 0, '0 rows fetched for disabled tests from BigQuery.'

  logging.info('Total fetched %d rows for disabled tests from BigQuery.',
               total_rows)
  return local_tests


def _GetMemoryFlags(builder_name):
  """Parses the builder_name for memory flags."""
  memory_flags = []
  for pattern, flag in _MEMORY_FLAGS_REGEX:
    if re.search(pattern, builder_name):
      memory_flags.append(flag)
  return memory_flags


def _CreateDisabledVariant(build_id, builder_name, step_name):
  """Creates a test variant for which a test is disabled.

  Args:
    build_id (int): Build id of the build.
    builder_name (str): Builder name of the build.
    step_name (str): The name of the step.

  Returns:
    variant_configurations (tuple): Alphabetically sorted tuple of key-value
      pairs defining the test variant or 'Unknown' if no configurations found.
  """
  variant_configurations = []
  os_name = step_util.GetOS(
      build_id, builder_name, step_name, partial_match=True)
  if os_name:
    variant_configurations.append('os:%s' % os_name)
  else:
    logging.info('Failed to obtain os for build_id: %s', build_id)

  variant_configurations.extend(_GetMemoryFlags(builder_name))
  variant_configurations.sort()

  if not variant_configurations:
    logging.info(
        'Failed to define test variant for build_id: %s, step_name: %s',
        build_id, step_name)
    variant_configurations = (_DEFAULT_CONFIG,)

  return tuple(variant_configurations)


def _CreateIssueKeys(bugs):
  """Creates a list of FlakeIssue keys from a list of bugs.

  Args:
    bugs (list): list of crbug.com links.

  Returns:
    issue_keys (list): List of FlakeIssue keys for each valid bug link.
  """
  issue_keys = set()
  for bug in bugs:
    if bug and isinstance(bug, basestring):
      match = _ISSUE_LINK_REGEX[0].match(bug) or _ISSUE_LINK_REGEX[1].match(bug)
      if not match:
        continue
      issue_id = int(match.groups()[0])
      issue_key = ndb.Key('FlakeIssue',
                          '%s@%d' % (_DEFAULT_LUCI_PROJECT, issue_id))
      issue_keys.add(issue_key)
  return issue_keys


def _CreateLocalTests(row, local_tests):
  """Creates a LuciTest key-test variant pair for a row fetched from BigQuery.

  Args:
    row: A row of query result.
    local_tests (dict): LuciTest entities in local memory in the format
      {LuciTest.key: {'disabled_test_variants : set(), issue_keys: set()},
      mutated by this function.
  """
  build_id = row['build_id']
  builder_name = row['builder_name']
  step_name = row['step_name']
  test_name = row['test_name']
  bugs = row['bugs']

  if int(build_id) == 1:
    # To filter out tests results with invalid build_id.
    # TODO (crbug.com/999215): Remove this check after test-results is fixed.
    logging.info('Failed to define test variant for build_id: %s, row is %r',
                 build_id, row)
    return

  normalized_step_name = Flake.NormalizeStepName(build_id, step_name)
  normalized_test_name = Flake.NormalizeTestName(test_name, step_name)
  test_key = LuciTest.CreateKey(_DEFAULT_LUCI_PROJECT, normalized_step_name,
                                normalized_test_name)
  if not local_tests.get(test_key):
    local_tests[test_key] = {
        'disabled_test_variants': set(),
        'issue_keys': set()
    }

  disabled_variant = _CreateDisabledVariant(build_id, builder_name, step_name)
  local_tests[test_key]['disabled_test_variants'].add(disabled_variant)
  local_tests[test_key]['issue_keys'].update(_CreateIssueKeys(bugs))


@ndb.tasklet
def _UpdateDatastore(test_key, query_time, disabled_test_variants, issue_keys):
  """Updates disabled_test_variants, issue_keys for a LuciTest in the datastore.

  Args:
    test_key (ndb.Key): Key of LuciTest entities.
    query_time (datetime): The time of the latest query.
    disabled_test_variants (set): Disabled test variants to write to datastore.
    issue_keys (ndb.Key): FlakeIssue keys for bugs associated with the test.
  """
  test = yield test_key.get_async()
  if not test:
    test = LuciTest(key=test_key)
    test.issue_keys = []
  new_issue_keys = issue_keys.difference(test.issue_keys)
  for new_issue_key in new_issue_keys:
    _CreateIssue(new_issue_key)
  test.issue_keys.extend(new_issue_keys)
  test.issue_keys.sort()

  test.disabled_test_variants = disabled_test_variants
  test.last_updated_time = query_time
  yield test.put_async()


@ndb.tasklet
def _CreateIssue(issue_key):
  """Creates an issue in the datastore for the given issue_key.

  Creates an issue if one does not already exist. Does not overwrite existing
  entities.

  Args:
    issue_key (ndb.Key): FlakeIssue key for which to create a FlakeIssue entity.
  """
  issue = yield issue_key.get_async()
  if not issue:
    monorail_project, issue_id = issue_key.id().split('@')
    issue_id = int(issue_id)
    issue = FlakeIssue.Create(monorail_project, issue_id)
    yield issue.put_async()


@ndb.toplevel
def _UpdateCurrentlyDisabledTests(local_tests, query_time):
  """Updates currently disabled tests.

  Overwrites existing disabled_test_variants and adds to issue_keys if new
  monorail issues are associated with the test.

  Args:
    local_tests (dict): LuciTest entities in local memory in the format
      {LuciTest.key: {'disabled_test_variants : set(), issue_keys: set()},
      mutated by this function.
    query_time(datetime): The time of the latest query.
  """
  remote_tests = ndb.get_multi(local_tests.keys())

  # (LuciTest key, set of disabled test variants)
  updated_test_keys = []

  for remote_test, local_test in zip(remote_tests, local_tests.items()):
    if not remote_test:
      updated_test_keys.append(local_test[0])
    elif local_test[1]['disabled_test_variants'].symmetric_difference(
        remote_test.disabled_test_variants):
      updated_test_keys.append(local_test[0])
    elif local_test[1]['issue_keys'].difference(remote_test.issue_keys):
      updated_test_keys.append(local_test[0])

  logging.info('Updating or Creating %d LuciTests: ', len(updated_test_keys))
  for updated_test_key in updated_test_keys:
    _UpdateDatastore(updated_test_key, query_time,
                     local_tests[updated_test_key]['disabled_test_variants'],
                     local_tests[updated_test_key]['issue_keys'])


@ndb.toplevel
def _UpdateNoLongerDisabledTests(currently_disabled_test_keys, query_time):
  """Removes test variants from LuciTest entities which are no longer disabled.

  Does not overwrite other attributes from the LuciTest entity.

  Args:
    currently_disabled_test_keys (list): Keys of currently disabled LuciTest
      entities.
    query_time (datetime): The time of the latest query.
  """
  # pylint: disable=singleton-comparison
  disabled_test_keys = LuciTest.query(LuciTest.disabled == True).fetch(
      keys_only=True)

  no_longer_disabled_test_keys = set(disabled_test_keys) - set(
      currently_disabled_test_keys)

  logging.info('%d tests are no longer disabled: ',
               len(no_longer_disabled_test_keys))
  for no_longer_disabled_test_key in no_longer_disabled_test_keys:
    _UpdateDatastore(no_longer_disabled_test_key, query_time, set(), set())


def ProcessQueryForDisabledTests():
  query_time = time_util.GetUTCNow()
  local_tests = _ExecuteQuery(parameters=_GetQueryParameters())
  _UpdateCurrentlyDisabledTests(local_tests, query_time)
  _UpdateNoLongerDisabledTests(local_tests.keys(), query_time)
