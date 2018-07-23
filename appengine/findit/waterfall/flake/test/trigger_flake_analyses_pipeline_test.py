# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import mock

from gae_libs.pipelines import pipeline_handlers
from services import ci_failure
from model.wf_analysis import WfAnalysis
from waterfall import waterfall_config
from waterfall.flake import flake_analysis_service
from waterfall.flake.trigger_flake_analyses_pipeline import (
    TriggerFlakeAnalysesPipeline)
from waterfall.test import wf_testcase


class TriggerFlakeAnalysesPipelineTest(wf_testcase.WaterfallTestCase):
  app_module = pipeline_handlers._APP

  @mock.patch.object(
      waterfall_config,
      'GetCheckFlakeSettings',
      return_value={'throttle_flake_analyses': True})
  @mock.patch.object(
      ci_failure,
      'GetStepMetadata',
      return_value={
          'canonical_step_name': 'a_tests',
          'isolate_target_name': 'a_tests'
      })
  @mock.patch('services.monitoring.OnFlakeIdentified')
  def testTriggerFlakeAnalysesPipeline(self, mock_monitoring, *_):
    master_name = 'm'
    builder_name = 'b'
    build_number = 2
    step_name = 'a_tests'
    test_name = 'Unittest1.Subtest1'

    analysis = WfAnalysis.Create(master_name, builder_name, build_number)
    analysis.flaky_tests = {step_name: [test_name, 'Unittest1.Subtest2']}
    analysis.put()

    with mock.patch.object(
        flake_analysis_service,
        'ScheduleAnalysisForFlake') as mocked_ScheduleAnalysisForFlake:
      pipeline = TriggerFlakeAnalysesPipeline()
      pipeline.run(master_name, builder_name, build_number)
      self.assertTrue(mocked_ScheduleAnalysisForFlake.called)
      mock_monitoring.assert_has_calls([
          mock.call('a_tests', 'a_tests', 'analyzed', 1),
          mock.call('a_tests', 'a_tests', 'throttled', 1)
      ])

  @mock.patch.object(
      waterfall_config,
      'GetCheckFlakeSettings',
      return_value={'throttle_flake_analyses': False})
  @mock.patch.object(
      ci_failure,
      'GetStepMetadata',
      return_value={
          'canonical_step_name': 'a_tests',
          'isolate_target_name': 'a_tests'
      })
  @mock.patch('services.monitoring.OnFlakeIdentified')
  def testTriggerFlakeAnalysesPipelineUnthrottled(self, mock_monitoring, *_):
    master_name = 'm'
    builder_name = 'b'
    build_number = 2
    step_name = 'a_tests'
    test_name = 'Unittest1.Subtest1'

    analysis = WfAnalysis.Create(master_name, builder_name, build_number)
    analysis.flaky_tests = {step_name: [test_name]}
    analysis.put()

    with mock.patch.object(
        flake_analysis_service,
        'ScheduleAnalysisForFlake') as mocked_ScheduleAnalysisForFlake:
      pipeline = TriggerFlakeAnalysesPipeline()
      pipeline.run(master_name, builder_name, build_number)
      self.assertTrue(mocked_ScheduleAnalysisForFlake.called)
      mock_monitoring.assert_has_calls(
          [mock.call('a_tests', 'a_tests', 'analyzed', 1)])

  @mock.patch.object(
      waterfall_config,
      'GetCheckFlakeSettings',
      return_value={'throttle_flake_analyses': True})
  def testTriggerFlakeAnalysesPipelineScheduledReturnsFalse(self, _):
    master_name = 'm'
    builder_name = 'b'
    build_number = 2

    analysis = WfAnalysis.Create(master_name, builder_name, build_number)
    analysis.put()

    with mock.patch.object(
        flake_analysis_service, 'ScheduleAnalysisForFlake',
        return_value=False) as mocked_ScheduleAnalysisForFlake:
      pipeline = TriggerFlakeAnalysesPipeline()
      pipeline.run(master_name, builder_name, build_number)
      self.assertFalse(mocked_ScheduleAnalysisForFlake.called)
