# Copyright 2015 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import base64
from datetime import timedelta
import hashlib
import json
import logging
import random
import time
import urllib
from urlparse import urlparse
import zlib

from google.appengine.api.urlfetch_errors import DeadlineExceededError
from google.appengine.api.urlfetch_errors import DownloadError
from google.appengine.api.urlfetch_errors import ConnectionClosedError

from common.findit_http_client import FinditHttpClient
from common import monitoring
from common.waterfall import buildbucket_client
from gae_libs.gitiles.cached_gitiles_repository import CachedGitilesRepository
from gae_libs.http import auth_util
from infra_api_clients import logdog_util
from libs import time_util
from model.wf_try_bot_cache import WfTryBot
from model.wf_try_bot_cache import WfTryBotCache
from waterfall import waterfall_config
from waterfall.flake import flake_constants
from waterfall.swarming_task_request import SwarmingTaskRequest

# Swarming task states.
STATES_RUNNING = ('RUNNING', 'PENDING')
STATE_COMPLETED = 'COMPLETED'
STATES_NOT_RUNNING = ('BOT_DIED', 'CANCELED', 'COMPLETED', 'EXPIRED',
                      'TIMED_OUT')

# TODO(crbug.com/785463): Use enum for error codes.

# Swarming task stopped error codes.
BOT_DIED = 30
CANCELED = 40
EXPIRED = 50
TIMED_OUT = 60

STATES_NOT_RUNNING_TO_ERROR_CODES = {
    'BOT_DIED': BOT_DIED,
    'CANCELED': CANCELED,
    'EXPIRED': EXPIRED,
    'TIMED_OUT': TIMED_OUT,
}

# Urlfetch error codes.
URLFETCH_DOWNLOAD_ERROR = 100
URLFETCH_DEADLINE_EXCEEDED_ERROR = 110
URLFETCH_CONNECTION_CLOSED_ERROR = 120

# Outputs_ref is None.
NO_TASK_OUTPUTS = 300

# Unable to retrieve output json.
NO_OUTPUT_JSON = 320

# Other/miscellaneous error codes.
UNKNOWN = 1000

# Swarming URL templates.
BOT_LIST_URL = 'https://%s/_ah/api/swarming/v1/bots/list%s'
BOT_COUNT_URL = 'https://%s/_ah/api/swarming/v1/bots/count%s'
NEW_TASK_URL = 'https://%s/_ah/api/swarming/v1/tasks/new'
TASK_ID_URL = 'https://%s/_ah/api/swarming/v1/task/%s/request'
TASK_RESULT_URL = 'https://%s/_ah/api/swarming/v1/task/%s/result'

DEFAULT_MINIMUM_NUMBER_AVAILABLE_BOTS = 5
DEFAULT_MINIMUM_PERCENTAGE_AVAILABLE_BOTS = 0.1


def SwarmingHost():
  return waterfall_config.GetSwarmingSettings().get('server_host')


def _GetBackoffSeconds(retry_backoff, tries, maximum_retry_interval):
  """Returns how many seconds to wait before next retry.

  Params:
    retry_backoff (int): The base backoff in seconds.
    tries (int): Indicates how many tries have been done.
    maximum_retry_interval (int): The upper limit in seconds of how long to wait
      between retries.
  """
  return min(retry_backoff * (2**(tries - 1)), maximum_retry_interval)


def _OnConnectionFailed(url, exception_type):
  host = urlparse(url).hostname
  assert host
  monitoring.outgoing_http_errors.increment({
      'host': host,
      'exception': exception_type
  })


def SendRequestToServer(url, http_client, post_data=None):
  """Sends GET/POST request to arbitrary url and returns response content.

  Because the Swarming and Isolated servers that SendRequestToServer tries to
  contact are prone to outages, exceptions trying to reach them may occur thus
  this method should retry. We want to monitor and document these occurrences
  even if the request eventually succeeds after retrying, with the last error
  encountered being the one that is reported.

  Args:
    url (str): The url to send the request to.
    http_client (HttpClient): The httpclient object with which to make the
      server calls.
    post_data (dict): Data/params to send with the request, if any.

  Returns:
    content (dict), error (dict): The content from the server and the last error
    encountered trying to retrieve it.
  """
  headers = {'Authorization': 'Bearer ' + auth_util.GetAuthToken()}
  swarming_settings = waterfall_config.GetSwarmingSettings()
  should_retry = swarming_settings.get('should_retry_server')
  timeout_seconds = (
      swarming_settings.get('server_retry_timeout_hours') * 60 * 60)
  maximum_retry_interval = swarming_settings.get(
      'maximum_server_contact_retry_interval_seconds')
  deadline = time.time() + timeout_seconds
  retry_backoff = 60
  tries = 1
  error = None

  if post_data:
    post_data = json.dumps(post_data, sort_keys=True, separators=(',', ':'))
    headers['Content-Type'] = 'application/json; charset=UTF-8'
    headers['Content-Length'] = len(post_data)

  while True:
    try:
      if post_data:
        status_code, content = http_client.Post(url, post_data, headers=headers)
      else:
        status_code, content = http_client.Get(url, headers=headers)
      if status_code == 200:
        # Also return the last error encountered to be handled in the calling
        # code.
        return content, error
      else:
        # The retry upon 50x (501 excluded) is automatically handled in the
        # underlying http_client, which by default retries 5 times with
        # exponential backoff.
        return None, {
            'code': status_code,
            'message': 'Unexpected status code from http request'
        }
    except ConnectionClosedError as e:
      error = {'code': URLFETCH_CONNECTION_CLOSED_ERROR, 'message': e.message}
      _OnConnectionFailed(url, 'ConnectionClosedError')
    except DeadlineExceededError as e:
      error = {'code': URLFETCH_DEADLINE_EXCEEDED_ERROR, 'message': e.message}
      _OnConnectionFailed(url, 'DeadlineExceededError')
    except DownloadError as e:
      error = {'code': URLFETCH_DOWNLOAD_ERROR, 'message': e.message}
      _OnConnectionFailed(url, 'DownloadError')
    except Exception as e:  # pragma: no cover
      logging.error(
          'An unknown exception occurred that need to be monitored: %s',
          e.message)
      error = {'code': UNKNOWN, 'message': e.message}
      _OnConnectionFailed(url, 'Unknown Exception')

    if should_retry and time.time() < deadline:  # pragma: no cover
      # Wait, then retry if applicable.
      wait_time = _GetBackoffSeconds(retry_backoff, tries,
                                     maximum_retry_interval)
      logging.info('Retrying connection to %s in %d seconds', url, wait_time)
      time.sleep(wait_time)
      tries += 1
    else:
      if should_retry:
        # Indicate in the error that the retry timeout was reached.
        error['retry_timeout'] = True
      break

  logging.error('Failed to get an adequate response from %s. No data could be '
                'retrieved', url)
  return None, error


def GetSwarmingTaskRequest(task_id, http_client):
  """Returns an instance of SwarmingTaskRequest representing the given task."""
  url = TASK_ID_URL % (SwarmingHost(), task_id)
  content, error = SendRequestToServer(url, http_client)

  # TODO(lijeffrey): Handle/report error in calling functions.
  if not error:
    json_data = json.loads(content)
    return SwarmingTaskRequest.Deserialize(json_data)
  return None


def TriggerSwarmingTask(request, http_client):
  """Triggers a new Swarming task for the given request.

  The Swarming task priority will be overwritten, and extra tags might be added.
  Args:
    request (SwarmingTaskRequest): A Swarming task request.
    http_client (RetryHttpClient): An http client with automatic retry.
  """
  # Use a priority much lower than CQ for now (CQ's priority is 30).
  # Later we might use a higher priority -- a lower value here.
  # Note: the smaller value, the higher priority.
  swarming_settings = waterfall_config.GetSwarmingSettings()
  request_expiration_hours = swarming_settings.get('request_expiration_hours')
  request.priority = max(100, swarming_settings.get('default_request_priority'))
  request.expiration_secs = request_expiration_hours * 60 * 60

  request.tags.extend(['findit:1', 'project:Chromium', 'purpose:post-commit'])

  response_data, error = SendRequestToServer(NEW_TASK_URL % SwarmingHost(),
                                             http_client, request.Serialize())

  if not error:
    return json.loads(response_data)['task_id'], None

  return None, error


def ListSwarmingTasksDataByTags(master_name,
                                builder_name,
                                build_number,
                                http_client,
                                additional_tag_filters=None):
  """Downloads tasks data from swarming server.

  Args:
    master_name(str): Value of the master tag.
    builder_name(str): Value of the buildername tag.
    build_number(int): Value of the buildnumber tag.
    http_client(RetryHttpClient): The http client to send HTTPs requests.
    additional_tag_filters(dict): More tag filters to be added.

  Returns:
    ([{
      'swarming_task_id': ...,
      'completed_ts': ...,
      'started_ts': ...,
      ...
    }, ...])
  """
  base_url = ('https://%s/_ah/api/swarming/v1/tasks/'
              'list?tags=%s&tags=%s&tags=%s') % (
                  waterfall_config.GetSwarmingSettings().get('server_host'),
                  urllib.quote('master:%s' % master_name),
                  urllib.quote('buildername:%s' % builder_name),
                  urllib.quote('buildnumber:%d' % build_number))
  additional_tag_filters = additional_tag_filters or {}
  for tag_name, tag_value in additional_tag_filters.iteritems():
    base_url += '&tags=%s' % urllib.quote('%s:%s' % (tag_name, tag_value))

  items = []
  cursor = None

  while True:
    if not cursor:
      url = base_url
    else:
      url = base_url + '&cursor=%s' % urllib.quote(cursor)
    new_data, _ = SendRequestToServer(url, http_client)

    # TODO(lijeffrey): handle error in calling functions.
    if not new_data:
      break

    new_data_json = json.loads(new_data)
    if new_data_json.get('items'):
      items.extend(new_data_json['items'])

    if new_data_json.get('cursor'):
      cursor = new_data_json['cursor']
    else:
      break

  return items


def GenerateIsolatedData(outputs_ref):
  if not outputs_ref:
    return {}
  return {
      'digest': outputs_ref['isolated'],
      'namespace': outputs_ref['namespace'],
      'isolatedserver': outputs_ref['isolatedserver']
  }


def GetSwarmingTaskResultById(task_id, http_client):
  """Gets swarming result, checks state and returns outputs ref if needed."""
  base_url = TASK_RESULT_URL % (SwarmingHost(), task_id)
  json_data = {}

  data, error = SendRequestToServer(base_url, http_client)

  if not error:
    json_data = json.loads(data)

  return json_data, error


def GetSwarmingTaskFailureLog(outputs_ref, http_client):
  """Downloads failure log from isolated server."""
  isolated_data = GenerateIsolatedData(outputs_ref)
  return _DownloadTestResults(isolated_data, http_client)


def GetTagValue(tags, tag_name):
  """Returns the content for a specific tag."""
  tag_prefix = tag_name + ':'
  content = None
  for tag in tags:
    if tag.startswith(tag_prefix):
      content = tag[len(tag_prefix):]
      break
  return content


def GetIsolatedDataForStep(master_name,
                           builder_name,
                           build_number,
                           step_name,
                           http_client,
                           only_failure=True):
  """Returns the isolated data for a specific step.

  Args:
    master_name(str): Value of the master tag.
    builder_name(str): Value of the buildername tag.
    build_number(int): Value of the buildnumber tag.
    step_name(str): Value of the stepname tag.
    http_client(RetryHttpClient): The http client to send HTTPs requests.
    only_failure(bool): A flag to determine if only failure info is needed.
  """
  step_isolated_data = []
  data = ListSwarmingTasksDataByTags(master_name, builder_name, build_number,
                                     http_client, {'stepname': step_name})
  if not data:
    return step_isolated_data

  for item in data:
    if not item.get('outputs_ref'):
      # Task might time out and no outputs_ref was saved.
      continue

    if only_failure:
      if item['failure'] and not item['internal_failure']:
        # Only retrieves test results from tasks which have failures and
        # the failure should not be internal infrastructure failure.
        isolated_data = GenerateIsolatedData(item['outputs_ref'])
        step_isolated_data.append(isolated_data)
    else:
      isolated_data = GenerateIsolatedData(item['outputs_ref'])
      step_isolated_data.append(isolated_data)

  return step_isolated_data


def GetIsolatedShaForStep(master_name, builder_name, build_number, step_name,
                          http_client):
  """Gets the isolated sha of a master/builder/build/step to send to swarming.

  Args:
    master_name (str): The name of the main waterall master.
    builder_name (str): The name of the main waterfall builder.
    build_number (int): The build number to retrieve the isolated sha of.
    step_name (str): The step name to retrieve the isolated sha of.

  Returns:
    (str): The isolated sha pointing to the compiled binaries at the requested
        configuration.
  """
  data = ListSwarmingTasksDataByTags(master_name, builder_name, build_number,
                                     http_client, {'stepname': step_name})
  if not data:
    logging.error('Failed to get swarming task data for %s/%s/%s/%s',
                  master_name, builder_name, build_number, step_name)
    return None

  # Each task should have the same sha, so only need to read from the first one.
  tags = data[0]['tags']
  sha = GetTagValue(tags, 'data')
  if sha:
    return sha

  logging.error('Isolated sha not found for %s/%s/%s/%s', master_name,
                builder_name, build_number, step_name)
  return None


def _FetchOutputJsonInfoFromIsolatedServer(isolated_data, http_client):
  """Sends POST request to isolated server and returns response content.

  This function is used for fetching
    1. hash code for the output.json file,
    2. the redirect url.
  """
  if not isolated_data:
    return None

  post_data = {
      'digest': isolated_data['digest'],
      'namespace': {
          'namespace': isolated_data['namespace']
      }
  }
  url = '%s/_ah/api/isolateservice/v1/retrieve' % (
      isolated_data['isolatedserver'])

  return SendRequestToServer(url, http_client, post_data)


def _ProcessRetrievedContent(output_json_content, http_client):
  """Downloads output.json file from isolated server or process it directly.

  If there is a url provided, send get request to that url to download log;
  else the log would be in content so use it directly.
  """
  json_content = json.loads(output_json_content)
  output_json_url = json_content.get('url')
  if output_json_url:
    get_content, _ = SendRequestToServer(output_json_url, http_client)
    # TODO(lijeffrey): handle error in calling function.
  elif json_content.get('content'):
    get_content = base64.b64decode(json_content['content'])
  else:  # pragma: no cover
    get_content = None  # Just for precausion.
  try:
    return json.loads(zlib.decompress(get_content)) if get_content else None
  except ValueError:  # pragma: no cover
    logging.info(
        'swarming result is invalid: %s' % zlib.decompress(get_content))
    return None


def _DownloadTestResults(isolated_data, http_client):
  """Downloads the output.json file and returns the json object.

  The basic steps to get test results are:
  1. Use isolated_data to get hash to output.json,
  2. Use hash from step 1 to get test results.

  But in each step, if the returned content is too big, we may not directly get
  the content, instead we get a url and we need to send an extra request to the
  url. This is handled in _ProcessRetrievedContent.
  """
  # First POST request to get hash for the output.json file.
  content, error = _FetchOutputJsonInfoFromIsolatedServer(
      isolated_data, http_client)
  if error:
    return None, error

  processed_content = _ProcessRetrievedContent(content, http_client)
  output_json_hash = processed_content.get('files', {}).get(
      'output.json', {}).get('h') if processed_content else None
  if not output_json_hash:
    return None, None

  # Second POST request to get the redirect url for the output.json file.
  data_for_output_json = {
      'digest': output_json_hash,
      'namespace': isolated_data['namespace'],
      'isolatedserver': isolated_data['isolatedserver']
  }

  output_json_content, error = _FetchOutputJsonInfoFromIsolatedServer(
      data_for_output_json, http_client)
  if error:
    return None, error

  # GET Request to get output.json file.
  return _ProcessRetrievedContent(output_json_content, http_client), None


def _MergeListsOfDicts(merged, shard):
  output = []
  for i in xrange(max(len(merged), len(shard))):
    merged_dict = merged[i] if i < len(merged) else {}
    shard_dict = shard[i] if i < len(shard) else {}
    output_dict = merged_dict.copy()
    output_dict.update(shard_dict)
    output.append(output_dict)
  return output


def _MergeSwarmingTestShards(shard_results):
  """Merges the shards into one.

  Args:
    shard_results (list): A list of dicts with individual shard results.

  Returns:
    A dict with the following form:
    {
      'all_tests':[
        'AllForms/FormStructureBrowserTest.DataDrivenHeuristics/0',
        'AllForms/FormStructureBrowserTest.DataDrivenHeuristics/1',
        'AllForms/FormStructureBrowserTest.DataDrivenHeuristics/10',
        ...
      ]
      'per_iteration_data':[
        {
          'AllForms/FormStructureBrowserTest.DataDrivenHeuristics/109': [
            {
              'elapsed_time_ms': 4719,
              'losless_snippet': true,
              'output_snippet': '[ RUN      ] run outputs\\n',
              'output_snippet_base64': 'WyBSVU4gICAgICBdIEFsbEZvcm1zL0Zvcm1T'
              'status': 'SUCCESS'
            }
          ],
        },
        ...
      ]
    }
  """
  merged_results = {'all_tests': set(), 'per_iteration_data': []}
  for shard_result in shard_results:
    merged_results['all_tests'].update(shard_result.get('all_tests', []))
    merged_results['per_iteration_data'] = _MergeListsOfDicts(
        merged_results['per_iteration_data'],
        shard_result.get('per_iteration_data', []))
  merged_results['all_tests'] = sorted(merged_results['all_tests'])
  return merged_results


def RetrieveShardedTestResultsFromIsolatedServer(list_isolated_data,
                                                 http_client):
  """Gets test results from isolated server and merge the results."""
  shard_results = []
  for isolated_data in list_isolated_data:
    output_json, _ = _DownloadTestResults(isolated_data, http_client)
    if not output_json:
      # TODO(lijeffrey): Report/handle error returned from _DownloadTestResults.
      return None
    shard_results.append(output_json)

  if len(list_isolated_data) == 1:
    return shard_results[0]
  return _MergeSwarmingTestShards(shard_results)


def GetIsolatedOutputForTask(task_id, http_client):
  """Get isolated output for a swarming task based on it's id."""
  json_data, error = GetSwarmingTaskResultById(task_id, http_client)

  if error or not json_data:
    return None

  outputs_ref = json_data.get('outputs_ref')
  if not outputs_ref:
    return None

  output_json, error = GetSwarmingTaskFailureLog(outputs_ref, http_client)

  if error:
    return None
  return output_json


def DimensionsToQueryString(dimensions):
  if isinstance(dimensions, dict):
    dimension_list = ['%s:%s' % (k, v) for k, v in dimensions.iteritems()]
  else:
    dimension_list = dimensions
  dimension_qs = '&dimensions='.join(dimension_list)
  # Url looks like 'https://chromium-swarm.appspot.com/_ah/api/swarming/v1/bots
  # /count?dimensions=os:Windows-7-SP1&dimensions=cpu:x86-64'
  return '?dimensions=%s' % dimension_qs


def BotsAvailableForTask(step_metadata):
  """Check if there are available bots for a swarming task's dimensions."""
  if not step_metadata:
    return False

  minimum_number_of_available_bots = (
      waterfall_config.GetSwarmingSettings().get(
          'minimum_number_of_available_bots',
          DEFAULT_MINIMUM_NUMBER_AVAILABLE_BOTS))
  minimum_percentage_of_available_bots = (
      waterfall_config.GetSwarmingSettings().get(
          'minimum_percentage_of_available_bots',
          DEFAULT_MINIMUM_PERCENTAGE_AVAILABLE_BOTS))
  dimensions = step_metadata.get('dimensions')
  bot_counts = GetSwarmingBotCounts(dimensions, FinditHttpClient())

  total_count = bot_counts.get('count') or -1
  available_count = bot_counts.get('available', 0)
  available_rate = float(available_count) / total_count

  return (available_count > minimum_number_of_available_bots and
          available_rate > minimum_percentage_of_available_bots)


def GetSwarmingBotCounts(dimensions, http_client):
  """Gets number of swarming bots for certain dimensions.

  Args:
    dimensions (dict): A dict of dimensions.
    http_client (HttpClient): The httpclient object with which to make the
      server calls.
  Returns:
    bot_counts (dict): Dict of numbers of available swarming bots.
  """
  if not dimensions:
    return {}

  url = BOT_COUNT_URL % (SwarmingHost(), DimensionsToQueryString(dimensions))

  content, error = SendRequestToServer(url, http_client)
  if error or not content:
    return {}

  content_data = json.loads(content)

  bot_counts = {
      k: int(content_data.get(k, 0))
      for k in ('busy', 'count', 'dead', 'quarantined')
  }
  bot_counts['available'] = (bot_counts['count'] - bot_counts['busy'] -
                             bot_counts['dead'] - bot_counts['quarantined'])

  return bot_counts


# TODO(robertocn): Move this outside of swarming_util, as this applies to all
# tryjobs now.
def GetStepLog(try_job_id, full_step_name, http_client, log_type='stdout'):
  """Returns specific log of the specified step."""

  error, build = buildbucket_client.GetTryJobs([try_job_id])[0]
  if error:
    logging.exception('Error retrieving buildbucket build id: %s' % try_job_id)
    return None

  # 1. Get log.
  data = logdog_util.GetStepLogForBuild(build.response, full_step_name,
                                        log_type, http_client)

  if log_type.lower() == 'json.output[ninja_info]':
    # Check if data is malformatted.
    try:
      json.loads(data)
    except ValueError:
      logging.error('json.output[ninja_info] is malformatted')
      return None

  if log_type.lower() not in ['stdout', 'json.output[ninja_info]']:
    try:
      return json.loads(data) if data else None
    except ValueError:
      logging.error('Failed to json load data for %s. Data is: %s.' % (log_type,
                                                                       data))

  return data


def UpdateAnalysisResult(analysis_result, flaky_failures):
  """Updates WfAnalysis' result and result_analysis on flaky failures.

  If found flaky tests from swarming reruns, or flaky tests or compile from
  try jobs, updates WfAnalysis.
  """
  all_flaked = True
  for failure in analysis_result.get('failures', {}):
    step_name = failure.get('step_name')
    if step_name in flaky_failures:
      failure['flaky'] = True
      for test in failure.get('tests', []):
        if test.get('test_name') in flaky_failures[step_name]:
          test['flaky'] = True
        else:
          all_flaked = False
          failure['flaky'] = False
    else:
      # Checks all other steps to see if all failed steps/ tests are flaky.
      if not failure.get('flaky'):
        all_flaked = False

  return all_flaked


def GetETAToStartAnalysis(manually_triggered):
  """Returns an ETA as of a UTC datetime.datetime to start the analysis.

  If not urgent, Swarming tasks should be run off PST peak hours from 11am to
  6pm on workdays.

  Args:
    manually_triggered (bool): True if the analysis is from manual request, like
        by a Chromium sheriff.

  Returns:
    The ETA as of a UTC datetime.datetime to start the analysis.
  """
  if manually_triggered:
    # If the analysis is manually triggered, run it right away.
    return time_util.GetUTCNow()

  now_at_pst = time_util.GetPSTNow()
  if now_at_pst.weekday() >= 5:  # PST Saturday or Sunday.
    return time_util.GetUTCNow()

  if now_at_pst.hour < 11 or now_at_pst.hour >= 18:  # Before 11am or after 6pm.
    return time_util.GetUTCNow()

  # Set ETA time to 6pm, and also with a random latency within 30 minutes to
  # avoid sudden burst traffic to Swarming.
  diff = timedelta(
      hours=18 - now_at_pst.hour,
      minutes=-now_at_pst.minute,
      seconds=-now_at_pst.second + random.randint(0, 30 * 60),
      microseconds=-now_at_pst.microsecond)
  eta = now_at_pst + diff

  # Convert back to UTC.
  return time_util.ConvertPSTToUTC(eta)


def IsTestEnabled(test_name, task_id):
  """Returns True if the test is enabled, False otherwise."""
  # Get the isolated outputs from the test that was just run.
  isolate_output = GetIsolatedOutputForTask(task_id, FinditHttpClient())
  isolate_output = isolate_output if isolate_output else {}
  all_tests = isolate_output.get('all_tests', [])
  disabled_tests = isolate_output.get('disabled_tests', [])

  # Check if the analysis' test is disabled in the current build.
  # If the disabled tests array is empty, we assume the test is enabled.
  return test_name in all_tests and test_name not in disabled_tests
