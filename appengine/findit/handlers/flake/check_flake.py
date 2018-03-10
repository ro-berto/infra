# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging

from google.appengine.ext import ndb

from gae_libs import token
from gae_libs.handlers.base_handler import BaseHandler
from gae_libs.handlers.base_handler import Permission
from gae_libs.http import auth_util
from libs import analysis_status
from libs import time_util
from model import triage_status
from model.flake.flake_analysis_request import FlakeAnalysisRequest
from model.flake.flake_try_job import FlakeTryJob
from model.flake.flake_try_job_data import FlakeTryJobData
from model.flake.master_flake_analysis import MasterFlakeAnalysis
from waterfall import buildbot
from waterfall.flake import flake_analysis_service
from waterfall.flake import triggering_sources
from waterfall.flake.recursive_flake_pipeline import RecursiveFlakePipeline
from waterfall.trigger_base_swarming_task_pipeline import NO_TASK
from waterfall.trigger_base_swarming_task_pipeline import NO_TASK_EXCEPTION


def _GetSuspectedFlakeInfo(analysis):
  """Returns a dict with information about the suspected flake build.

  Args:
    analysis (MasterFlakeAnalysis): The master flake analysis the suspected
      flake build is associated with.

  Returns:
    A dict in the format:
      {
          'confidence': float or None,
          'build_number': int,
          'commit_position': int,
          'git_hash': str,
          'lower_bound_commit_position': int,
          'lower_bound_git_hash': str,
          'triage_result': int (correct, incorrect, etc.)
      }
  """
  if not analysis or analysis.suspected_flake_build_number is None:
    return {}

  data_point = analysis.GetDataPointOfSuspectedBuild()
  assert data_point

  return {
      'confidence':
          analysis.confidence_in_suspected_build,
      'build_number':
          analysis.suspected_flake_build_number,
      'commit_position':
          data_point.commit_position,
      'git_hash':
          data_point.git_hash,
      'lower_bound_commit_position': (
          data_point.previous_build_commit_position),
      'lower_bound_git_hash':
          data_point.previous_build_git_hash,
      'triage_result': (analysis.triage_history[-1].triage_result
                        if analysis.triage_history else triage_status.UNTRIAGED)
  }


def _GetSuspectInfo(suspect_urlsafe_key):
  """Returns a dict with information about a suspect.

  Args:
    suspect_urlsaf_key (str): A urlsafe-key to a FlakeCulprit entity.

  Returns:
    A dict in the format:
      {
          'commit_position': int,
          'git_hash': str,
          'url': str,
      }
  """
  suspect_key = ndb.Key(urlsafe=suspect_urlsafe_key)
  # TODO(crbug.com/799308): Remove this hack when bug is fixed.
  assert suspect_key.pairs()[0]
  assert suspect_key.pairs()[0][0]  # Name of the model.
  assert suspect_key.pairs()[0][1]  # Id of the model.
  suspect = ndb.Key(suspect_key.pairs()[0][0], suspect_key.pairs()[0][1]).get()
  assert suspect

  return {
      'commit_position': suspect.commit_position,
      'git_hash': suspect.revision,
      'url': suspect.url,
  }


def _GetSuspectsInfoForAnalysis(analysis):
  """Returns a list of dicts with information about an analysis' suspected CLs.

  Args:
    analysis (MasterFlakeAnalysis): The master flake analysis the suspected
      flake build is associated with.

  Returns:
    A list of dicts in the format:
        [
            {
                'commit_position': int,
                'git_hash': str,
                'url': str,
            },
            ...
        ]
  """
  if not analysis or not analysis.suspect_urlsafe_keys:
    return []

  suspects_info = []
  for suspect_urlsafe_key in analysis.suspect_urlsafe_keys:
    suspects_info.append(_GetSuspectInfo(suspect_urlsafe_key))
  return suspects_info


def _GetCulpritInfo(analysis):
  """Returns a dict with information about a suspected culprit.

  Args:
    analysis (MasterFlakeAnalysis): The master flake analysis the suspected
      flake build is associated with.

  Returns:
    A dict in the format:
      {
          'commit_position': int,
          'git_hash': str,
          'url': str,
      }
  """
  if analysis.culprit_urlsafe_key is None:
    return {}

  suspect_info = _GetSuspectInfo(analysis.culprit_urlsafe_key)
  suspect_info['confidence'] = analysis.confidence_in_culprit
  return suspect_info


def _GetCoordinatesData(analysis):

  def _GetBasicData(point):
    return {
        'commit_position': point.commit_position,
        'pass_rate': point.pass_rate,
        'task_ids': point.task_ids,
        'build_number': point.build_number,
        'git_hash': point.git_hash,
        'try_job_url': point.try_job_url
    }

  if not analysis or not analysis.data_points:
    return []

  # Order by commit position from earliest to latest.
  data_points = sorted(analysis.data_points, key=lambda x: x.commit_position)
  coordinates = []

  previous_data_point = data_points[0]
  data = _GetBasicData(previous_data_point)
  coordinates.append(data)

  for i in range(1, len(data_points)):
    data_point = data_points[i]
    data = _GetBasicData(data_point)
    data['lower_bound_commit_position'] = previous_data_point.commit_position
    data['lower_bound_git_hash'] = previous_data_point.git_hash
    previous_data_point = data_point
    coordinates.append(data)

  return coordinates


def _GetNumbersOfDataPointGroups(data_points):
  build_level_number = 0
  revision_level_number = 0

  for data_point in data_points:
    if data_point.try_job_url:
      revision_level_number += 1
    else:
      build_level_number += 1

  return build_level_number, revision_level_number


def _GetLastAttemptedSwarmingTaskDetails(analysis):
  swarming_task_id = analysis.last_attempted_swarming_task_id
  build_number = analysis.last_attempted_build_number

  task_id = (
      swarming_task_id if swarming_task_id and
      swarming_task_id.lower() not in (NO_TASK, NO_TASK_EXCEPTION) else None)

  return {'task_id': task_id, 'build_number': build_number}


def _GetLastAttemptedTryJobDetails(analysis):
  last_attempted_revision = analysis.last_attempted_revision
  if not last_attempted_revision:
    return {}

  try_job = FlakeTryJob.Get(analysis.master_name, analysis.builder_name,
                            analysis.step_name, analysis.test_name,
                            last_attempted_revision)

  if not try_job or not try_job.try_job_ids:
    return {}

  try_job_id = try_job.try_job_ids[-1]
  try_job_data = FlakeTryJobData.Get(try_job_id)
  if not try_job_data:
    return {}

  return {
      'status': analysis_status.STATUS_TO_DESCRIPTION.get(try_job.status),
      'url': try_job_data.try_job_url
  }


def _GetDurationForAnalysis(analysis):
  """Returns the duration of the given analysis."""
  if analysis.status == analysis_status.PENDING:
    return None
  return time_util.FormatDuration(analysis.start_time, analysis.end_time or
                                  time_util.GetUTCNow())


class CheckFlake(BaseHandler):
  PERMISSION_LEVEL = Permission.ANYONE

  def _ShowCustomRunOptions(self, analysis):
    # TODO(lijeffrey): Remove checks for admin and debug flag once analyze
    # manual input for a regression range is implemented.
    return (auth_util.IsCurrentUserAdmin() and
            self.request.get('debug') == '1' and
            analysis.status != analysis_status.RUNNING and
            analysis.try_job_status != analysis_status.RUNNING)

  def _ValidateInput(self, step_name, test_name, bug_id):
    """Ensures the input is valid and generates an error otherwise.

    Args:
      step_name (str): The step the flaky test was found on.
      test_name (str): The name of the flaky test.
      bug_id (str): The bug number associated with the flaky test.

    Returns:
      None if all input fields are valid, or an error dict otherwise.
    """
    if not step_name:
      return self.CreateError('Step name must be specified', 400)

    if not test_name:
      return self.CreateError('Test name must be specified', 400)

    if bug_id and not bug_id.isdigit():
      return self.CreateError('Bug id must be an int', 400)

    return None

  @staticmethod
  def _CreateAndScheduleFlakeAnalysis(master_name,
                                      builder_name,
                                      build_number,
                                      step_name,
                                      test_name,
                                      bug_id,
                                      rerun=False):
    # pylint: disable=unused-argument
    """Create and schedule a flake analysis.

    Args:
      master_name (string): The name of the master.
      builder_name (string): The name of the builder.
      build_number (int): Build number to run against.
      step_name (string): The name of the step.
      test_name (string): The name of the test.
      bug_id (int): The bug id.
      rerun (boolean): Is this analysis a rerun.
    Returns:
      (analysis, scheduled) analysis is the new analysis created.
      scheduled is returned from flake analysis service.
    """
    user_email = auth_util.GetUserEmail()
    is_admin = auth_util.IsCurrentUserAdmin()

    request = FlakeAnalysisRequest.Create(test_name, False, bug_id)
    request.AddBuildStep(master_name, builder_name, build_number, step_name,
                         time_util.GetUTCNow())
    scheduled = flake_analysis_service.ScheduleAnalysisForFlake(
        request,
        user_email,
        is_admin,
        triggering_sources.FINDIT_UI,
        rerun=rerun)

    analysis = MasterFlakeAnalysis.GetVersion(
        master_name, builder_name, build_number, step_name, test_name)

    return analysis, scheduled

  @staticmethod
  def _CanRerunAnalysis(analysis):
    return not (analysis.status == analysis_status.RUNNING or
                analysis.status == analysis_status.PENDING or
                analysis.try_job_status == analysis_status.RUNNING or
                analysis.try_job_status == analysis_status.PENDING)

  def _HandleRerunAnalysis(self):
    """Rerun an analysis as a response to a user request."""
    # If the key has been specified, we can derive the above information
    # from the analysis itself.
    if not auth_util.IsCurrentUserAdmin():
      return self.CreateError('Only admin is allowed to rerun.', 403)

    key = self.request.get('key')
    if not key:
      return self.CreateError('No key was provided.', 404)

    analysis = ndb.Key(urlsafe=key).get()
    if not analysis:
      return self.CreateError('Analysis of flake is not found.', 404)

    if not self._CanRerunAnalysis(analysis):
      return self.CreateError(
          'Cannot rerun analysis if one is currently running or pending.', 400)

    master_name = analysis.master_name
    builder_name = analysis.builder_name
    build_number = analysis.build_number
    step_name = analysis.step_name
    test_name = analysis.test_name
    bug_id = analysis.bug_id

    logging.info('Rerun button pushed, analysis will be reset and triggered.\n'
                 'Analysis key: %s', key)

    analysis, _ = self._CreateAndScheduleFlakeAnalysis(
        master_name, builder_name, build_number, step_name, test_name, bug_id,
        True)
    return self.CreateRedirect(
        '/waterfall/flake?redirect=1&key=%s' % analysis.key.urlsafe())

  def _HandleCancelAnalysis(self):
    """Cancel analysis as a response to a user request."""
    if not auth_util.IsCurrentUserAdmin():
      return self.CreateError('Only admin is allowed to cancel.', 403)

    key = self.request.get('key')
    if not key:
      return self.CreateError('No key was provided.', 404)

    analysis = ndb.Key(urlsafe=key).get()
    if not analysis:
      return self.CreateError('Analysis of flake is not found.', 404)

    if (analysis.status != analysis_status.RUNNING and
        analysis.try_job_status != analysis_status.RUNNING):
      return self.CreateError('Can\'t cancel an analysis that\'s complete', 400)

    if not analysis.root_pipeline_id:
      return self.CreateError('No root pipeline found for analysis.', 404)
    root_pipeline = RecursiveFlakePipeline.from_id(analysis.root_pipeline_id)

    if not root_pipeline:
      return self.CreateError('Root pipeline couldn\'t be found.', 404)

    # If we can find the pipeline, cancel it.
    root_pipeline.abort('Pipeline was cancelled manually.')
    error = {
        'error': 'The pipeline was aborted manually.',
        'message': 'The pipeline was aborted manually.'
    }

    # If culprit analysis is running, set the status to error.
    if analysis.try_job_status == analysis_status.RUNNING:
      try_job_status = analysis_status.ERROR
    # if culprit analysis is pending, set it to skipped.
    elif analysis.try_job_status == analysis_status.PENDING:
      try_job_status = analysis_status.SKIPPED
    else:
      try_job_status = analysis.try_job_status

    analysis.Update(
        status=analysis_status.ERROR,
        try_job_status=try_job_status,
        error=error,
        end_time=time_util.GetUTCNow())

    return self.CreateRedirect(
        '/waterfall/flake?redirect=1&key=%s' % analysis.key.urlsafe())

  @token.VerifyXSRFToken()
  def HandlePost(self):
    # Information needed to execute this endpoint, will be populated
    # by the branches below.
    rerun = self.request.get('rerun', '0').strip() == '1'
    cancel = self.request.get('cancel', '0').strip() == '1'
    if rerun:  # Rerun an analysis.
      return self._HandleRerunAnalysis()
    elif cancel:  # Force an analysis to be cancelled.
      return self._HandleCancelAnalysis()
    else:  # Regular POST requests to start an analysis.
      # If the key hasn't been specified, then we get the information from
      # other URL parameters.
      build_url = self.request.get('url', '').strip()
      build_info = buildbot.ParseBuildUrl(build_url)
      if not build_info:
        return self.CreateError('Unknown build info!', 400)
      master_name, builder_name, build_number = build_info

      step_name = self.request.get('step_name', '').strip()
      test_name = self.request.get('test_name', '').strip()
      bug_id = self.request.get('bug_id', '').strip()

      error = self._ValidateInput(step_name, test_name, bug_id)
      if error:
        return error

      build_number = int(build_number)
      bug_id = int(bug_id) if bug_id else None

      analysis, scheduled = self._CreateAndScheduleFlakeAnalysis(
          master_name, builder_name, build_number, step_name, test_name, bug_id,
          False)

      if not analysis:
        if scheduled is None:
          # User does not have permission to trigger, nor was any previous
          # analysis triggered to view.
          return {
              'template': 'error.html',
              'data': {
                  'error_message': (
                      'No permission to schedule an analysis for flaky test. '
                      'Please log in with your @google.com account first.'),
              },
              'return_code': 403,
          }

        # Check if a previous request has already covered this analysis so use
        # the results from that analysis.
        request = FlakeAnalysisRequest.GetVersion(key=test_name)

        if not (request and request.analyses):
          return {
              'template': 'error.html',
              'data': {
                  'error_message': (
                      'Flake analysis is not supported for "%s/%s". Either '
                      'the test type is not supported or the test is not '
                      'swarmed yet.' % (step_name, test_name)),
              },
              'return_code': 400,
          }

        analysis = request.FindMatchingAnalysisForConfiguration(
            master_name, builder_name)

        if not analysis:
          logging.error('Flake analysis was deleted unexpectedly!')
          return {
              'template': 'error.html',
              'data': {
                  'error_message': 'Flake analysis was deleted unexpectedly!',
              },
              'return_code': 404,
          }

      logging.info('Analysis: %s has a scheduled status of: %r', analysis.key,
                   scheduled)
      return self.CreateRedirect(
          '/waterfall/flake?redirect=1&key=%s' % analysis.key.urlsafe())

  def HandleGet(self):
    key = self.request.get('key')
    if not key:
      return self.CreateError('No key was provided.', 404)

    analysis = ndb.Key(urlsafe=key).get()
    if not analysis:
      return self.CreateError('Analysis of flake is not found.', 404)

    suspected_flake = _GetSuspectedFlakeInfo(analysis)
    culprit = _GetCulpritInfo(analysis)
    build_level_number, revision_level_number = _GetNumbersOfDataPointGroups(
        analysis.data_points)
    regression_range_confidence = suspected_flake.get('confidence', 0)
    culprit_confidence = culprit.get('confidence', 0)

    def AsPercentString(val):
      """0-1 as a percent, rounded and returned as a string"""
      return "{0:d}".format(int(round(val * 100.0))) if val else ''

    regression_range_confidence = AsPercentString(regression_range_confidence)
    culprit_confidence = AsPercentString(culprit_confidence)

    # TODO(crbug.com/789289): Culprit status value should be correct.
    culprit_status = (
        analysis_status.STATUS_TO_DESCRIPTION.get(analysis.try_job_status) or
        '')

    if not culprit_status:
      if analysis.status == analysis_status.RUNNING:
        culprit_status = 'pending'
      elif analysis.status == analysis_status.ERROR:
        culprit_status = 'skipped'
      else:
        culprit_status = 'error'

    data = {
        'key':
            analysis.key.urlsafe(),
        'build_number':
            analysis.build_number,
        'pass_rates': [],
        'try_job_status':
            analysis_status.STATUS_TO_DESCRIPTION.get(analysis.try_job_status),
        'last_attempted_swarming_task':
            _GetLastAttemptedSwarmingTaskDetails(analysis),
        'last_attempted_try_job':
            _GetLastAttemptedTryJobDetails(analysis),
        'version_number':
            analysis.version_number,
        'suspected_flake':
            suspected_flake,
        'suspected_culprits':
            _GetSuspectsInfoForAnalysis(analysis),
        'culprit':
            culprit,
        'request_time':
            time_util.FormatDatetime(analysis.request_time),
        'build_level_number':
            build_level_number,
        'revision_level_number':
            revision_level_number,
        'error':
            analysis.error_message,
        'iterations_to_rerun':
            analysis.iterations_to_rerun,
        'show_admin_options':
            self._ShowCustomRunOptions(analysis),
        'show_debug_options':
            self._ShowDebugInfo(),
        'pipeline_status_path':
            analysis.pipeline_status_path,

        # new ui stuff
        'master_name':
            analysis.master_name,
        'builder_name':
            analysis.builder_name,
        'step_name':
            analysis.step_name,
        'test_name':
            analysis.test_name,
        'analysis_status':
            analysis.status_description,
        'regression_range_upper': (suspected_flake.get('commit_position', '') or
                                   suspected_flake.get('git_hash', '') or ''),
        'regression_range_lower': (
            suspected_flake.get('lower_bound_commit_position', '') or
            suspected_flake.get('lower_bound_git_hash', '') or ''),
        'regression_range_confidence':
            regression_range_confidence,
        'culprit_analysis_status':
            culprit_status,
        'culprit_url':
            culprit.get('url', ''),
        'culprit_text':
            culprit.get('commit_position', '') or culprit.get('revision', ''),
        'culprit_confidence':
            culprit_confidence,
        'bug_id':
            str(analysis.bug_id) if analysis.bug_id else ''
    }

    if (auth_util.IsCurrentUserAdmin() and analysis.completed and
        analysis.triage_history):
      data['triage_history'] = analysis.GetTriageHistory()

    data['pending_time'] = time_util.FormatDuration(
        analysis.request_time, analysis.start_time or time_util.GetUTCNow())
    data['duration'] = _GetDurationForAnalysis(analysis)

    data['pass_rates'] = _GetCoordinatesData(analysis)

    return {'template': 'flake/result.html', 'data': data}
