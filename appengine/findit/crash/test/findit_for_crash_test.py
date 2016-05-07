# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from collections import defaultdict

from common.blame import Region, Blame
from common.change_log import ChangeLog
from common.dependency import Dependency, DependencyRoll
from common.git_repository import GitRepository
from crash import findit_for_crash
from crash.callstack import StackFrame, CallStack
from crash.results import MatchResult
from crash.stacktrace import Stacktrace
from crash.test.crash_test_suite import CrashTestSuite


DUMMY_CHANGELOG1 = ChangeLog.FromDict({
    'author_name': 'r@chromium.org',
    'message': 'dummy',
    'committer_email': 'r@chromium.org',
    'commit_position': 175900,
    'author_email': 'r@chromium.org',
    'touched_files': [
        {
            'change_type': 'add',
            'new_path': 'a.cc',
            'old_path': None,
        },
    ],
    'author_time': 'Thu Mar 31 21:24:43 2016',
    'committer_time': 'Thu Mar 31 21:28:39 2016',
    'commit_url':
        'https://repo.test/+/1',
    'code_review_url': 'https://codereview.chromium.org/3281',
    'committer_name': 'example@chromium.org',
    'revision': '1',
    'reverted_revision': None
})

DUMMY_CHANGELOG2 = ChangeLog.FromDict({
    'author_name': 'example@chromium.org',
    'message': 'dummy',
    'committer_email': 'example@chromium.org',
    'commit_position': 175976,
    'author_email': 'example@chromium.org',
    'touched_files': [
        {
            'change_type': 'add',
            'new_path': 'f0.cc',
            'old_path': 'b/f0.cc'
        },
    ],
    'author_time': 'Thu Mar 31 21:24:43 2016',
    'committer_time': 'Thu Mar 31 21:28:39 2016',
    'commit_url':
        'https://repo.test/+/2',
    'code_review_url': 'https://codereview.chromium.org/3281',
    'committer_name': 'example@chromium.org',
    'revision': '2',
    'reverted_revision': '1'
})

DUMMY_CHANGELOG3 = ChangeLog.FromDict({
    'author_name': 'e@chromium.org',
    'message': 'dummy',
    'committer_email': 'e@chromium.org',
    'commit_position': 176000,
    'author_email': 'e@chromium.org',
    'touched_files': [
        {
            'change_type': 'modify',
            'new_path': 'f.cc',
            'old_path': 'f.cc'
        },
        {
            'change_type': 'delete',
            'new_path': None,
            'old_path': 'f1.cc'
        },
    ],
    'author_time': 'Thu Apr 1 21:24:43 2016',
    'committer_time': 'Thu Apr 1 21:28:39 2016',
    'commit_url':
        'https://repo.test/+/3',
    'code_review_url': 'https://codereview.chromium.org/3281',
    'committer_name': 'example@chromium.org',
    'revision': '3',
    'reverted_revision': None
})


class FinditForCrashTest(CrashTestSuite):

  def testGetDepsInCrashStack(self):
    crash_stack = CallStack(0)
    crash_stack.extend([
        StackFrame(0, 'src/', '', 'func0', 'f0.cc', [1]),
        StackFrame(1, 'src/', '', 'func1', 'f1.cc', [2, 3]),
    ])
    crash_deps = {'src/': Dependency('src/', 'https://chromium_repo', '1'),
                  'src/v8/': Dependency('src/v8/', 'https://v8_repo', '2')}

    expected_stack_deps = {'src/': crash_deps['src/']}

    self.assertEqual(
        findit_for_crash.GetDepsInCrashStack(crash_stack, crash_deps),
        expected_stack_deps)

  def testGetChangeLogsForFilesGroupedByDeps(self):
    regression_deps_rolls = {
        'src/dep1': {
            'path': 'src/dep1',
            'repo_url': 'https://url_dep1',
            'old_revision': '7',
            'new_revision': '9',
        },
        'src/dep2': {
            'path': 'src/dep2',
            'repo_url': 'https://url_dep2',
            'old_revision': '3',
            'new_revision': None,
        },
        'src/': {
            'path': 'src/',
            'repo_url': ('https://chromium.googlesource.com/chromium/'
                         'src.git'),
            'old_revision': '4',
            'new_revision': '5',
        },
    }

    stack_deps = {
        'src/': Dependency('src/', 'https://url_src', 'rev1', 'DEPS'),
    }

    def _MockGetChangeLogs(*_):
      return [DUMMY_CHANGELOG1, DUMMY_CHANGELOG2, DUMMY_CHANGELOG3]

    self.mock(GitRepository, 'GetChangeLogs', _MockGetChangeLogs)

    dep_file_to_changelogs, ignore_cls = (
        findit_for_crash.GetChangeLogsForFilesGroupedByDeps(
            regression_deps_rolls, stack_deps))
    dep_file_to_changelogs_json = defaultdict(lambda: defaultdict(list))
    for dep, file_to_changelogs in dep_file_to_changelogs.iteritems():
      for file_path, changelogs in file_to_changelogs.iteritems():
        for changelog in changelogs:
          dep_file_to_changelogs_json[dep][file_path].append(changelog.ToDict())

    expected_dep_file_to_changelogs_json = {
        'src/': {
            'a.cc': [DUMMY_CHANGELOG1.ToDict()],
            'f.cc': [DUMMY_CHANGELOG3.ToDict()]
        }
    }
    self.assertEqual(dep_file_to_changelogs_json,
                     expected_dep_file_to_changelogs_json)
    self.assertEqual(ignore_cls, set(['1']))

  def testGetStackInfosForFilesGroupedByDeps(self):

    main_stack = CallStack(0)
    main_stack.extend(
        [StackFrame(0, 'src/', '', 'c(p* &d)', 'a.cc', [177]),
         StackFrame(1, 'src/', '', 'd(a* c)', 'a.cc', [227, 228, 229]),
         StackFrame(2, 'src/v8/', '', 'e(int)', 'b.cc', [87, 88, 89, 90])])

    low_priority_stack = CallStack(1)
    low_priority_stack.append(
        StackFrame(0, 'src/dummy/', '', 'c(p* &d)', 'd.cc', [17]))

    stacktrace = Stacktrace()
    stacktrace.extend([main_stack, low_priority_stack])

    crashed_deps = {'src/': Dependency('src/', 'https//repo', '2'),
                    'src/v8/': Dependency('src/v8', 'https//repo', '1')}

    expected_dep_file_to_stack_infos = {
        'src/': {
            'a.cc': [
                (main_stack[0], 0),
                (main_stack[1], 0),
            ],
        },
        'src/v8/': {
            'b.cc': [
                (main_stack[2], 0),
            ]
        }
    }

    dep_file_to_stack_infos = (
        findit_for_crash.GetStackInfosForFilesGroupedByDeps(
            stacktrace, crashed_deps))

    self.assertEqual(len(dep_file_to_stack_infos),
                     len(expected_dep_file_to_stack_infos))

    for dep, file_to_stack_infos in dep_file_to_stack_infos.iteritems():
      self.assertTrue(dep in expected_dep_file_to_stack_infos)
      expected_file_to_stack_infos = expected_dep_file_to_stack_infos[dep]

      for file_path, stack_infos in file_to_stack_infos.iteritems():
        self.assertTrue(file_path in expected_file_to_stack_infos)
        expected_stack_infos = expected_file_to_stack_infos[file_path]

        self._VerifyTwoStackInfosEqual(stack_infos, expected_stack_infos)

  def testGetBlameForFilesGroupedByDeps(self):

    dummy_blame = Blame('1', 'a.cc')
    dummy_blame.AddRegion(
        Region(1, 5, '0', 'a', 'a@email.com', 'Thu Mar 25 21:24:43 2016'))
    dummy_blame.AddRegion(
        Region(6, 10, '1', 'b', 'b@email.com', 'Thu Mar 31 21:24:43 2016'))

    def _MockGetBlame(*_):
      return dummy_blame

    self.mock(GitRepository, 'GetBlame', _MockGetBlame)

    main_stack = CallStack(0)
    main_stack.extend(
        [StackFrame(0, 'src/', '', 'c(p* &d)', 'a.cc', [177]),
         StackFrame(1, 'src/', '', 'd(a* c)', 'a.cc', [227, 228, 229])])

    low_priority_stack = CallStack(1)
    low_priority_stack.append(
        StackFrame(0, 'dummy_dep/', '', 'k(p* &d)', 'h.cc', [17]))

    stacktrace = Stacktrace()
    stacktrace.extend([main_stack, low_priority_stack])

    crash_deps = {'src/': Dependency('src/', 'https://chromium_repo', '1')}

    dep_file_to_blame = findit_for_crash.GetBlameForFilesGroupedByDeps(
        stacktrace,
        findit_for_crash.GetDepsInCrashStack(main_stack, crash_deps))

    expected_dep_file_to_blame = {
        'src/': {
            'a.cc': dummy_blame
        }
    }

    for dep, file_to_blame in dep_file_to_blame.iteritems():

      self.assertTrue(dep in expected_dep_file_to_blame)

      expected_file_to_blame = expected_dep_file_to_blame[dep]
      for file_path, blame in file_to_blame.iteritems():
        self.assertTrue(file_path in expected_file_to_blame)
        expected_blame = expected_file_to_blame[file_path]

        self._VerifyTwoBlamesEqual(blame, expected_blame)

  def testFindMatchResults(self):
    dep_file_to_changelogs = {
        'src/': {
            'a.cc': [
                DUMMY_CHANGELOG1,
            ]
        }
    }

    dep_file_to_stack_infos = {
        'src/': {
            'a.cc': [
                (StackFrame(0, 'src/', '', 'func', 'a.cc', [1]), 0),
                (StackFrame(1, 'src/', '', 'func', 'a.cc', [7]), 0),
            ],
            'b.cc': [
                (StackFrame(2, 'src/', '', 'func', 'b.cc', [36]), 0),
            ]
        }
    }

    dummy_blame = Blame('9', 'a.cc')
    dummy_blame.AddRegion(
        Region(1, 5, '6', 'a', 'a@chromium.org', 'Thu Mar 31 21:24:43 2016'))
    dummy_blame.AddRegion(
        Region(6, 10, '1', 'b', 'b@chromium.org', 'Thu Jun 19 12:11:40 2015'))

    dep_file_to_blame = {
        'src/': {
            'a.cc': dummy_blame
        }
    }

    expected_match_results = [{
        'url': 'https://repo.test/+/1',
        'revision': '1',
        'dep_path': 'src/',
        'component': '',
        'author': 'r@chromium.org',
        'time': 'Thu Mar 31 21:24:43 2016',
        'reason': None,
        'confidence': None,
    }]

    match_results = findit_for_crash.FindMatchResults(dep_file_to_changelogs,
                                                      dep_file_to_stack_infos,
                                                      dep_file_to_blame)
    self.assertEqual([result.ToDict() for result in match_results],
                     expected_match_results)

  def testFindItForCrashNoRegressionRange(self):
    self.assertEqual(
        findit_for_crash.FindItForCrash(Stacktrace(), {}, {}),
        [])

  def testFindItForCrashNoMatchFound(self):

    def _MockFindMatchResults(*_):
      return []

    self.mock(findit_for_crash, 'FindMatchResults', _MockFindMatchResults)

    regression_deps_rolls = {'src/': DependencyRoll('src/', 'https://repo',
                                                    '1', '2')}
    self.assertEqual(findit_for_crash.FindItForCrash(
        Stacktrace(), regression_deps_rolls, {}), [])

  def testFindItForCrash(self):

    def _MockFindMatchResults(*_):
      match_result1 = MatchResult(DUMMY_CHANGELOG1, 'src/', '')
      match_result1.file_to_stack_infos = {
          'a.cc': [
              (StackFrame(0, 'src/', '', 'func', 'a.cc', [1]), 0),
              (StackFrame(1, 'src/', '', 'func', 'a.cc', [7]), 0),
          ]
      }
      match_result1.min_distance = 0

      match_result2 = MatchResult(DUMMY_CHANGELOG3, 'src/', '')
      match_result2.file_to_stack_infos = {
          'f.cc': [
              (StackFrame(5, 'src/', '', 'func', 'f.cc', [1]), 0),
          ]
      }
      match_result2.min_distance = 20

      return [match_result1, match_result2]

    self.mock(findit_for_crash, 'FindMatchResults', _MockFindMatchResults)

    expected_match_results = [
        {
            'url': 'https://repo.test/+/1',
            'revision': '1',
            'dep_path': 'src/',
            'component': '',
            'author': 'r@chromium.org',
            'time': 'Thu Mar 31 21:24:43 2016',
            'reason': ('1. Top frame changed is frame #0 (score: 1)\n'
                       '2. Minimum distance to crashed line is 0 (score: 1)\n'
                       '\nChanged file a.cc crashed in func (#0)'
                       ', func (#1)'),
            'confidence': 1,
        },
        {
            'url': 'https://repo.test/+/3',
            'revision': '3',
            'dep_path': 'src/',
            'component': '',
            'author': 'e@chromium.org',
            'time': 'Thu Apr 1 21:24:43 2016',
            'reason': ('1. Top frame changed is frame #5 (score: 0)\n'
                       '2. Minimum distance to crashed line is 20 (score: 0)\n'
                       '\nChanged file f.cc crashed in func (#5)'),
            'confidence': 0.22857142857142856,
        },
    ]

    regression_deps_rolls = {'src/': DependencyRoll('src/', 'https://repo',
                                                    '1', '2')}

    self.assertEqual(findit_for_crash.FindItForCrash(
        Stacktrace(), regression_deps_rolls, {}), expected_match_results)
