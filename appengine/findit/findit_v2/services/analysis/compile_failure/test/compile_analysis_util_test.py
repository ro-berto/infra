# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from datetime import datetime
import mock

from buildbucket_proto import common_pb2
from buildbucket_proto.build_pb2 import Build
from buildbucket_proto.build_pb2 import BuilderID
from buildbucket_proto.rpc_pb2 import SearchBuildsResponse
from buildbucket_proto.step_pb2 import Step

from common.waterfall import buildbucket_client
from findit_v2.model.compile_failure import CompileFailure
from findit_v2.model.compile_failure import CompileFailureAnalysis
from findit_v2.model.compile_failure import CompileRerunBuild
from findit_v2.model.gitiles_commit import GitilesCommit
from findit_v2.model.luci_build import LuciFailedBuild
from findit_v2.services.analysis.compile_failure import compile_analysis_util
from findit_v2.services.chromium_api import ChromiumProjectAPI
from findit_v2.services.context import Context
from findit_v2.services.failure_type import StepTypeEnum
from services import git
from waterfall.test import wf_testcase


class CompileUtilTest(wf_testcase.TestCase):

  def setUp(self):
    super(CompileUtilTest, self).setUp()
    self.build_id = 8000000000123
    self.build_number = 123
    self.builder = BuilderID(
        project='chromium', bucket='try', builder='linux-rel')
    self.build = Build(
        id=self.build_id,
        builder=self.builder,
        number=self.build_number,
        status=common_pb2.FAILURE)
    self.build.input.gitiles_commit.host = 'gitiles.host.com'
    self.build.input.gitiles_commit.project = 'project/name'
    self.build.input.gitiles_commit.ref = 'ref/heads/master'
    self.build.input.gitiles_commit.id = 'git_sha_123'
    self.build.create_time.FromDatetime(datetime(2019, 4, 9))
    self.build.start_time.FromDatetime(datetime(2019, 4, 9, 0, 1))
    self.build.end_time.FromDatetime(datetime(2019, 4, 9, 1))

    self.context = Context(
        luci_project_name='chromium',
        gitiles_host='gitiles.host.com',
        gitiles_project='project/name',
        gitiles_ref='ref/heads/master',
        gitiles_id='git_sha')

    self.build_info = {
        'id': 8000000000123,
        'number': self.build_number,
        'commit_id': 'git_sha_123'
    }

  @mock.patch.object(git, 'GetCommitPositionFromRevision', return_value=67890)
  def testSaveCompileFailures(self, _):
    detailed_compile_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': {
                        'id': 8000000000121,
                        'number': 121,
                        'commit_id': 'git_sha'
                    },
                    'last_passed_build': {
                        'id': 8000000000120,
                        'number': 120,
                        'commit_id': 'git_sha'
                    },
                },
            },
            'first_failed_build': {
                'id': 8000000000121,
                'number': 121,
                'commit_id': 'git_sha'
            },
            'last_passed_build': {
                'id': 8000000000120,
                'number': 120,
                'commit_id': 'git_sha'
            },
        },
    }

    compile_analysis_util.SaveCompileFailures(self.context, self.build,
                                              detailed_compile_failures)

    build = LuciFailedBuild.get_by_id(self.build_id)
    self.assertIsNotNone(build)

    compile_failures = CompileFailure.query(ancestor=build.key).fetch()
    self.assertEqual(1, len(compile_failures))
    self.assertEqual(8000000000121, compile_failures[0].first_failed_build_id)

  @mock.patch.object(git, 'GetCommitPositionFromRevision', return_value=67890)
  def testSaveCompileFailuresOnlyStepLevelFailures(self, _):
    detailed_compile_failures = {
        'compile': {
            'failures': {},
            'first_failed_build': {
                'id': 8000000000121,
                'number': 121,
                'commit_id': 'git_sha'
            },
            'last_passed_build': {
                'id': 8000000000120,
                'number': 120,
                'commit_id': 'git_sha'
            },
        },
    }

    compile_analysis_util.SaveCompileFailures(self.context, self.build,
                                              detailed_compile_failures)

    build_entity = LuciFailedBuild.get_by_id(self.build_id)
    self.assertIsNotNone(build_entity)

    compile_failures = CompileFailure.query(ancestor=build_entity.key).fetch()
    self.assertEqual(1, len(compile_failures))
    self.assertEqual(8000000000121, compile_failures[0].first_failed_build_id)
    self.assertEqual([], compile_failures[0].output_targets)

  @mock.patch.object(ChromiumProjectAPI, 'GetCompileFailures')
  @mock.patch.object(buildbucket_client, 'GetV2Build')
  @mock.patch.object(buildbucket_client, 'SearchV2BuildsOnBuilder')
  def testDetectFirstFailures(self, mock_prev_builds, mock_get_build,
                              mock_prev_failures):
    """Test for the most common case: found both first_failed_build_id and
      last_passed_build_id."""
    mock_step = Step()
    mock_step.name = 'compile'
    mock_step.status = common_pb2.FAILURE
    build_122_id = 8000000000122
    build_122 = Build(
        id=build_122_id,
        builder=self.builder,
        number=self.build_number - 1,
        status=common_pb2.FAILURE)
    build_122.steps.extend([mock_step])
    build_122.input.gitiles_commit.id = 'git_sha_122'
    build_122_info = {
        'id': build_122_id,
        'number': self.build_number - 1,
        'commit_id': 'git_sha_122'
    }

    build_121_id = 8000000000121
    build_121 = Build(
        id=build_121_id,
        builder=self.builder,
        number=self.build_number - 2,
        status=common_pb2.SUCCESS)
    build_121.input.gitiles_commit.id = 'git_sha_121'
    build_121_info = {
        'id': build_121_id,
        'number': self.build_number - 2,
        'commit_id': 'git_sha_121'
    }

    mock_prev_builds.return_value = SearchBuildsResponse(
        builds=[build_122, build_121])
    mock_get_build.return_value = build_122

    mock_prev_failures.return_value = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': build_122_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': build_122_info,
            'last_passed_build': None,
        },
    }

    detailed_compile_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': None,
        },
    }

    compile_analysis_util.DetectFirstFailures(self.context, self.build,
                                              detailed_compile_failures)

    expected_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': build_122_info,
                    'last_passed_build': build_121_info,
                },
            },
            'first_failed_build': build_122_info,
            'last_passed_build': build_121_info,
        },
    }

    self.assertEqual(expected_failures, detailed_compile_failures)

  @mock.patch.object(buildbucket_client, 'GetV2Build')
  @mock.patch.object(buildbucket_client, 'SearchV2BuildsOnBuilder')
  def testDetectFirstFailuresPrevBuildDifferentStep(self, mock_prev_builds,
                                                    mock_get_build):
    """Test for previous build failed with different steps."""
    mock_step = Step()
    mock_step.name = 'test'
    mock_step.status = common_pb2.FAILURE
    mock_step1 = Step()
    mock_step1.name = 'compile'
    mock_step1.status = common_pb2.SUCCESS
    build_122_id = 8000000000122
    build_122 = Build(
        id=build_122_id,
        builder=self.builder,
        number=self.build_number - 1,
        status=common_pb2.FAILURE)
    build_122.steps.extend([mock_step, mock_step1])
    build_122.input.gitiles_commit.id = 'git_sha_122'
    build_122_info = {
        'id': build_122_id,
        'number': self.build_number - 1,
        'commit_id': 'git_sha_122'
    }

    build_121_id = 8000000000121
    build_121 = Build(
        id=build_121_id,
        builder=self.builder,
        number=self.build_number - 2,
        status=common_pb2.SUCCESS)

    mock_prev_builds.return_value = SearchBuildsResponse(
        builds=[build_122, build_121])
    mock_get_build.return_value = build_122

    detailed_compile_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': None,
        },
    }

    compile_analysis_util.DetectFirstFailures(self.context, self.build,
                                              detailed_compile_failures)

    expected_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': build_122_info,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': build_122_info,
        },
    }
    self.assertEqual(expected_failures, detailed_compile_failures)

  @mock.patch.object(buildbucket_client, 'GetV2Build')
  @mock.patch.object(buildbucket_client, 'SearchV2BuildsOnBuilder')
  def testDetectFirstFailuresPrevBuildNoCompile(self, mock_prev_builds,
                                                mock_get_build):
    """Test for previous build didn't run compile."""
    mock_step = Step()
    mock_step.name = 'test'
    mock_step.status = common_pb2.FAILURE
    build_122_id = 8000000000122
    build_122 = Build(
        id=build_122_id,
        builder=self.builder,
        number=self.build_number - 1,
        status=common_pb2.FAILURE)
    build_122.steps.extend([mock_step])

    build_121_id = 8000000000121
    build_121 = Build(
        id=build_121_id,
        builder=self.builder,
        number=self.build_number - 2,
        status=common_pb2.SUCCESS)
    build_121.input.gitiles_commit.id = 'git_sha_121'
    build_121_info = {
        'id': build_121_id,
        'number': self.build_number - 2,
        'commit_id': 'git_sha_121'
    }

    mock_prev_builds.return_value = SearchBuildsResponse(
        builds=[build_122, build_121])
    mock_get_build.return_value = build_122

    detailed_compile_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': None,
        },
    }

    compile_analysis_util.DetectFirstFailures(self.context, self.build,
                                              detailed_compile_failures)

    expected_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': build_121_info,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': build_121_info,
        },
    }
    self.assertEqual(expected_failures, detailed_compile_failures)

  @mock.patch.object(ChromiumProjectAPI, 'GetCompileFailures')
  @mock.patch.object(buildbucket_client, 'GetV2Build')
  @mock.patch.object(buildbucket_client, 'SearchV2BuildsOnBuilder')
  def testDetectFirstFailuresDifferentFirstFailure(
      self, mock_prev_builds, mock_get_build, mock_prev_failures):
    """Test for targets in current build failed from different builds."""
    mock_step = Step()
    mock_step.name = 'compile'
    mock_step.status = common_pb2.FAILURE
    build_122_id = 8000000000122
    build_122 = Build(
        id=build_122_id,
        builder=self.builder,
        number=self.build_number - 1,
        status=common_pb2.FAILURE)
    build_122.steps.extend([mock_step])
    build_122.input.gitiles_commit.id = 'git_sha_122'
    build_122_info = {
        'id': build_122_id,
        'number': self.build_number - 1,
        'commit_id': 'git_sha_122'
    }

    mock_step1 = Step()
    mock_step1.name = 'compile'
    mock_step1.status = common_pb2.FAILURE
    build_121_id = 8000000000121
    build_121 = Build(
        id=build_121_id,
        builder=self.builder,
        number=self.build_number - 2,
        status=common_pb2.FAILURE)
    build_121.steps.extend([mock_step1])
    build_121.input.gitiles_commit.id = 'git_sha_121'
    build_121_info = {
        'id': build_121_id,
        'number': self.build_number - 2,
        'commit_id': 'git_sha_121'
    }

    mock_step2 = Step()
    mock_step2.name = 'compile'
    mock_step2.status = common_pb2.FAILURE
    build_120_id = 8000000000121
    build_120 = Build(
        id=build_120_id,
        builder=self.builder,
        number=self.build_number - 3,
        status=common_pb2.FAILURE)
    build_120.steps.extend([mock_step2])
    build_120.input.gitiles_commit.id = 'git_sha_120'
    build_120_info = {
        'id': build_120_id,
        'number': self.build_number - 3,
        'commit_id': 'git_sha_120'
    }

    mock_prev_builds.return_value = SearchBuildsResponse(
        builds=[build_122, build_121, build_120])
    mock_get_build.side_effect = [build_122, build_121, build_120]

    # Failed compiling target3 but successfully compiled target1&2.
    failures_122 = {
        'compile': {
            'failures': {
                'target3': {
                    'output_targets': ['target3'],
                    'rule': 'ACTION',
                    'first_failed_build': build_122_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': build_122_info,
            'last_passed_build': None,
        },
    }
    # Has the same failed targets as current build.
    failures_121 = {
        'compile': {
            'failures': {
                'target3': {
                    'output_targets': ['target3'],
                    'rule': 'ACTION',
                    'first_failed_build': build_121_info,
                    'last_passed_build': None,
                },
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': build_121_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': build_121_info,
            'last_passed_build': None,
        },
    }
    # Failed compile step, but only different targets.
    failures_120 = {
        'compile': {
            'failures': {
                'target4': {
                    'output_targets': ['target4'],
                    'rule': 'CC',
                    'first_failed_build': build_120_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': build_120_info,
            'last_passed_build': None,
        },
    }
    mock_prev_failures.side_effect = [failures_122, failures_121, failures_120]

    detailed_compile_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': None,
                },
                'target3': {
                    'output_targets': ['target3'],
                    'rule': 'ACTION',
                    'first_failed_build': self.build_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': None,
        },
    }

    compile_analysis_util.DetectFirstFailures(self.context, self.build,
                                              detailed_compile_failures)

    expected_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': build_122_info,
                },
                'target3': {
                    'output_targets': ['target3'],
                    'rule': 'ACTION',
                    'first_failed_build': build_121_info,
                    'last_passed_build': build_120_info,
                },
            },
            'first_failed_build': build_121_info,
            'last_passed_build': build_120_info,
        },
    }

    self.assertEqual(expected_failures, detailed_compile_failures)

  @mock.patch.object(buildbucket_client, 'GetV2Build')
  @mock.patch.object(buildbucket_client, 'SearchV2BuildsOnBuilder')
  def testDetectFirstFailuresPrevBuildInfraFailure(self, mock_prev_builds,
                                                   mock_get_build):
    """Test for previous build failed with different steps."""
    mock_step1 = Step()
    mock_step1.name = 'compile'
    mock_step1.status = common_pb2.INFRA_FAILURE
    build_122_id = 8000000000122
    build_122 = Build(
        id=build_122_id,
        builder=self.builder,
        number=self.build_number - 1,
        status=common_pb2.FAILURE)
    build_122.steps.extend([mock_step1])

    build_121_id = 8000000000121
    build_121 = Build(
        id=build_121_id,
        builder=self.builder,
        number=self.build_number - 2,
        status=common_pb2.SUCCESS)
    build_121.input.gitiles_commit.id = 'git_sha_121'
    build_121_info = {
        'id': build_121_id,
        'number': self.build_number - 2,
        'commit_id': 'git_sha_121'
    }

    mock_prev_builds.return_value = SearchBuildsResponse(
        builds=[build_122, build_121])
    mock_get_build.return_value = build_122

    detailed_compile_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': None,
        },
    }

    compile_analysis_util.DetectFirstFailures(self.context, self.build,
                                              detailed_compile_failures)

    expected_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': build_121_info,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': build_121_info,
        },
    }
    self.assertEqual(expected_failures, detailed_compile_failures)

  def testGetFirstFailuresInCurrentBuild(self):
    build_122_info = {
        'id': 8000000000122,
        'number': self.build_number - 1,
        'commit_id': 'git_sha_122'
    }

    failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': build_122_info,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': build_122_info,
        },
    }

    expected_res = {
        'failures': {
            'compile': {
                'output_targets': ['target1', 'target2'],
                'last_passed_build': build_122_info,
            },
        },
        'last_passed_build': build_122_info
    }

    self.assertEqual(
        expected_res,
        compile_analysis_util.GetFirstFailuresInCurrentBuild(
            self.context, self.build, failures))

  def testGetFirstFailuresInCurrentBuildNoFirstFailures(self):
    build_122_info = {
        'id': 8000000000122,
        'number': self.build_number - 1,
        'commit_id': 'git_sha_122'
    }

    build_121_info = {
        'id': 8000000000121,
        'number': self.build_number - 2,
        'commit_id': 'git_sha_121'
    }

    failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': build_122_info,
                    'last_passed_build': build_121_info,
                },
            },
            'first_failed_build': build_122_info,
            'last_passed_build': build_121_info,
        },
    }

    expected_res = {'failures': {}, 'last_passed_build': None}

    self.assertEqual(
        expected_res,
        compile_analysis_util.GetFirstFailuresInCurrentBuild(
            self.context, self.build, failures))

  def testGetFirstFailuresInCurrentBuildNoLastPass(self):

    failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': self.build_info,
            'last_passed_build': None,
        },
    }

    expected_res = {'failures': {}, 'last_passed_build': None}

    self.assertEqual(
        expected_res,
        compile_analysis_util.GetFirstFailuresInCurrentBuild(
            self.context, self.build, failures))

  def testGetFirstFailuresInCurrentBuildOnlyStep(self):
    build_122_info = {
        'id': 8000000000122,
        'number': self.build_number - 1,
        'commit_id': 'git_sha_122'
    }

    failures = {
        'compile': {
            'failures': {},
            'first_failed_build': self.build_info,
            'last_passed_build': build_122_info,
        },
    }

    expected_res = {
        'failures': {
            'compile': {
                'output_targets': [],
                'last_passed_build': build_122_info,
            },
        },
        'last_passed_build': build_122_info
    }

    self.assertEqual(
        expected_res,
        compile_analysis_util.GetFirstFailuresInCurrentBuild(
            self.context, self.build, failures))

  def testGetFirstFailuresInCurrentBuildOnlyStepFailedBefore(self):
    build_122_info = {
        'id': 8000000000122,
        'number': self.build_number - 1,
        'commit_id': 'git_sha_122'
    }
    build_121_info = {
        'id': 8000000000121,
        'number': self.build_number - 2,
        'commit_id': 'git_sha_121'
    }

    failures = {
        'compile': {
            'failures': {},
            'first_failed_build': build_122_info,
            'last_passed_build': build_121_info,
        },
    }

    expected_res = {'failures': {}, 'last_passed_build': None}

    self.assertEqual(
        expected_res,
        compile_analysis_util.GetFirstFailuresInCurrentBuild(
            self.context, self.build, failures))

  def testGetFirstFailuresInCurrentBuildFailureStartedInDifferentBuild(self):
    build_122_info = {
        'id': 8000000000122,
        'number': self.build_number - 1,
        'commit_id': 'git_sha_122'
    }
    build_121_info = {
        'id': 8000000000121,
        'number': self.build_number - 2,
        'commit_id': 'git_sha_121'
    }

    failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': self.build_info,
                    'last_passed_build': build_122_info,
                },
                'target3': {
                    'output_targets': ['target3'],
                    'rule': 'ACTION',
                    'first_failed_build': build_122_info,
                    'last_passed_build': None,
                },
                'target4': {
                    'output_targets': ['target4'],
                    'rule': 'ACTION',
                    'first_failed_build': self.build_info,
                    'last_passed_build': build_121_info,
                },
            },
            'first_failed_build': build_122_info,
            'last_passed_build': None,
        },
    }

    expected_res = {
        'failures': {
            'compile': {
                'output_targets': ['target4', 'target1', 'target2'],
                'last_passed_build': build_121_info,
            },
        },
        'last_passed_build': build_121_info
    }

    self.assertEqual(
        expected_res,
        compile_analysis_util.GetFirstFailuresInCurrentBuild(
            self.context, self.build, failures))

  @mock.patch.object(
      ChromiumProjectAPI,
      'GetRerunBuilderId',
      return_value='chromium/findit/findit_variables')
  @mock.patch.object(
      git, 'GetCommitPositionFromRevision', side_effect=[66680, 66666, 66680])
  def testSaveCompileAnalysis(self, *_):
    build_121_info = {
        'id': 8000000000121,
        'number': self.build_number - 2,
        'commit_id': 'git_sha_121'
    }

    detailed_compile_failures = {
        'compile': {
            'failures': {
                'target1 target2': {
                    'output_targets': ['target1', 'target2'],
                    'rule': 'CXX',
                    'first_failed_build': {
                        'id': 8000000000121,
                        'number': 121,
                        'commit_id': 'git_sha'
                    },
                    'last_passed_build': {
                        'id': 8000000000120,
                        'number': 120,
                        'commit_id': 'git_sha'
                    },
                },
                'target3': {
                    'output_targets': ['target3'],
                    'rule': 'ACTION',
                    'first_failed_build': build_121_info,
                    'last_passed_build': None,
                },
            },
            'first_failed_build': {
                'id': 8000000000121,
                'number': 121,
                'commit_id': 'git_sha'
            },
            'last_passed_build': {
                'id': 8000000000120,
                'number': 120,
                'commit_id': 'git_sha'
            },
        },
    }

    compile_analysis_util.SaveCompileFailures(self.context, self.build,
                                              detailed_compile_failures)

    first_failures_in_current_build = {
        'failures': {
            'compile': {
                'output_targets': ['target1', 'target2'],
                'last_passed_build': build_121_info,
            },
        },
        'last_passed_build': build_121_info
    }
    compile_analysis_util.SaveCompileAnalysis(self.context, self.build,
                                              first_failures_in_current_build)

    analysis = CompileFailureAnalysis.GetVersion(self.build_id)
    self.assertIsNotNone(analysis)
    self.assertEqual('git_sha_121', analysis.last_passed_commit.gitiles_id)
    self.assertEqual(66666, analysis.last_passed_commit.commit_position)
    self.assertEqual('chromium/findit/findit_variables',
                     analysis.rerun_builder_id)
    self.assertEqual(1, len(analysis.compile_failure_keys))
    self.assertEqual(['target1', 'target2'],
                     analysis.compile_failure_keys[0].get().output_targets)

  @mock.patch.object(
      ChromiumProjectAPI,
      'GetCompileRerunBuildInputProperties',
      return_value={'recipe': 'compile'})
  @mock.patch.object(buildbucket_client, 'TriggerV2Build')
  def testTriggerRerunBuild(self, mock_trigger_build, _):
    build_entity = LuciFailedBuild.Create(
        luci_project=self.context.luci_project_name,
        luci_bucket=self.build.builder.bucket,
        luci_builder=self.build.builder.builder,
        build_id=self.build_id,
        legacy_build_number=self.build_number,
        gitiles_host=self.context.gitiles_host,
        gitiles_project=self.context.gitiles_project,
        gitiles_ref=self.context.gitiles_ref,
        gitiles_id=self.context.gitiles_id,
        commit_position=6000005,
        status=20,
        create_time=datetime(2019, 3, 28),
        start_time=datetime(2019, 3, 28, 0, 1),
        end_time=datetime(2019, 3, 28, 1),
        build_failure_type=StepTypeEnum.COMPILE)
    build_entity.put()

    compile_failure = CompileFailure.Create(
        failed_build_key=build_entity.key,
        step_ui_name='compile',
        output_targets=['a.o'],
        first_failed_build_id=self.build_id,
        failure_group_build_id=None)
    compile_failure.put()

    analysis = CompileFailureAnalysis.Create(
        luci_project=self.context.luci_project_name,
        luci_bucket=self.build.builder.bucket,
        luci_builder=self.build.builder.builder,
        build_id=self.build_id,
        gitiles_host=self.context.gitiles_host,
        gitiles_project=self.context.gitiles_project,
        gitiles_ref=self.context.gitiles_ref,
        last_passed_gitiles_id='left_sha',
        last_passed_cp=6000000,
        first_failed_gitiles_id=self.context.gitiles_id,
        first_failed_cp=6000005,
        rerun_builder_id='chromium/findit/findit-variables',
        compile_failure_keys=[compile_failure.key])
    analysis.Save()

    new_build_id = 800000024324
    new_build = Build(id=new_build_id, number=300)
    new_build.status = common_pb2.SCHEDULED
    new_build.create_time.FromDatetime(datetime(2019, 4, 20))
    rerun_builder = BuilderID(
        project='chromium', bucket='findit', builder='findit-variables')
    rerun_commit = GitilesCommit(
        gitiles_host=self.context.gitiles_host,
        gitiles_project=self.context.gitiles_project,
        gitiles_ref=self.context.gitiles_ref,
        gitiles_id='6000002',
        commit_position=6000002)
    output_targets = {'compile': ['a.o']}

    mock_trigger_build.return_value = new_build

    compile_analysis_util.TriggerRerunBuild(
        self.context, self.build_id, self.build, analysis.key, rerun_builder,
        rerun_commit, output_targets)

    rerun_build = CompileRerunBuild.get_by_id(new_build_id, parent=analysis.key)
    self.assertIsNotNone(rerun_build)

  @mock.patch.object(
      ChromiumProjectAPI,
      'GetCompileRerunBuildInputProperties',
      return_value={'recipe': 'compile'})
  @mock.patch.object(buildbucket_client, 'TriggerV2Build')
  def testTriggerRerunBuildFoundRunningBuild(self, mock_trigger_build, _):
    """This test is for the case where there's already an existing rerun build,
      so no new rerun-build should be scheduled."""
    build_entity = LuciFailedBuild.Create(
        luci_project=self.context.luci_project_name,
        luci_bucket=self.build.builder.bucket,
        luci_builder=self.build.builder.builder,
        build_id=self.build_id,
        legacy_build_number=self.build_number,
        gitiles_host=self.context.gitiles_host,
        gitiles_project=self.context.gitiles_project,
        gitiles_ref=self.context.gitiles_ref,
        gitiles_id=self.context.gitiles_id,
        commit_position=6000005,
        status=20,
        create_time=datetime(2019, 3, 28),
        start_time=datetime(2019, 3, 28, 0, 1),
        end_time=datetime(2019, 3, 28, 1),
        build_failure_type=StepTypeEnum.COMPILE)
    build_entity.put()

    compile_failure = CompileFailure.Create(
        failed_build_key=build_entity.key,
        step_ui_name='compile',
        output_targets=['a.o'],
        first_failed_build_id=self.build_id,
        failure_group_build_id=None)
    compile_failure.put()

    analysis = CompileFailureAnalysis.Create(
        luci_project=self.context.luci_project_name,
        luci_bucket=self.build.builder.bucket,
        luci_builder=self.build.builder.builder,
        build_id=self.build_id,
        gitiles_host=self.context.gitiles_host,
        gitiles_project=self.context.gitiles_project,
        gitiles_ref=self.context.gitiles_ref,
        last_passed_gitiles_id='left_sha',
        last_passed_cp=6000000,
        first_failed_gitiles_id=self.context.gitiles_id,
        first_failed_cp=6000005,
        rerun_builder_id='chromium/findit/findit-variables',
        compile_failure_keys=[compile_failure.key])
    analysis.Save()

    rerun_commit = GitilesCommit(
        gitiles_host=self.context.gitiles_host,
        gitiles_project=self.context.gitiles_project,
        gitiles_ref=self.context.gitiles_ref,
        gitiles_id='6000002',
        commit_position=6000002)

    rerun_builder = BuilderID(
        project='chromium', bucket='findit', builder='findit-variables')
    output_targets = {'compile': ['a.o']}

    CompileRerunBuild.Create(
        luci_project=rerun_builder.project,
        luci_bucket=rerun_builder.bucket,
        luci_builder=rerun_builder.builder,
        build_id=8000000000789,
        legacy_build_number=60789,
        gitiles_host=rerun_commit.gitiles_host,
        gitiles_project=rerun_commit.gitiles_project,
        gitiles_ref=rerun_commit.gitiles_ref,
        gitiles_id=rerun_commit.gitiles_id,
        commit_position=rerun_commit.commit_position,
        status=1,
        create_time=datetime(2019, 3, 28),
        parent_key=analysis.key).put()

    compile_analysis_util.TriggerRerunBuild(
        self.context, self.build_id, self.build, analysis.key, rerun_builder,
        rerun_commit, output_targets)

    self.assertFalse(mock_trigger_build.called)
