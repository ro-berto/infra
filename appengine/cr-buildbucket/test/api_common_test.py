# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import datetime
import json

from components import utils

from testing_utils import testing

import api_common
import model


class ApiCommonTests(testing.AppengineTestCase):

  def setUp(self):
    super(ApiCommonTests, self).setUp()
    self.patch(
        'components.utils.utcnow', return_value=datetime.datetime(2017, 1, 1)
    )
    self.test_build = model.Build(
        id=1,
        project='chromium',
        bucket='luci.chromium.try',
        create_time=datetime.datetime(2017, 1, 1),
        parameters={
            'buildername': 'linux_rel',
        },
        canary_preference=model.CanaryPreference.AUTO,
        swarming_hostname='swarming.example.com',
    )

  def test_expired_build_to_message(self):
    yesterday = utils.utcnow() - datetime.timedelta(days=1)
    yesterday_timestamp = utils.datetime_to_timestamp(yesterday)
    self.test_build.lease_key = 1
    self.test_build.lease_expiration_date = yesterday
    msg = api_common.build_to_message(self.test_build)
    self.assertEqual(msg.lease_expiration_ts, yesterday_timestamp)

  def test_build_to_dict_empty(self):
    expected = {
        'project': 'chromium',
        'bucket': 'luci.chromium.try',
        'created_ts': '1483228800000000',
        'id': '1',
        'parameters_json': json.dumps({'buildername': 'linux_rel'}),
        'result_details_json': 'null',
        'status': 'SCHEDULED',
        'tags': [],
        'utcnow_ts': '1483228800000000',
        'canary_preference': 'AUTO',
    }
    self.assertEqual(expected, api_common.build_to_dict(self.test_build))

  def test_build_to_dict_non_luci(self):
    self.test_build.bucket = 'master.chromium'
    self.test_build.swarming_hostname = None

    actual = api_common.build_to_dict(self.test_build)
    self.assertEqual(actual['project'], 'chromium')
    self.assertEqual(actual['bucket'], 'master.chromium')

  def test_build_to_dict_full(self):
    self.test_build.start_time = datetime.datetime(2017, 1, 2)
    self.test_build.complete_time = datetime.datetime(2017, 1, 2)
    self.test_build.status = model.BuildStatus.COMPLETED
    self.test_build.result = model.BuildResult.SUCCESS
    self.test_build.result_details = {'result': 'nice'}
    self.test_build.service_account = 'robot@example.com'
    expected = {
        'project': 'chromium',
        'bucket': 'luci.chromium.try',
        'completed_ts': '1483315200000000',
        'created_ts': '1483228800000000',
        'id': '1',
        'parameters_json': json.dumps({'buildername': 'linux_rel'}),
        'result': 'SUCCESS',
        'result_details_json': json.dumps({'result': 'nice'}),
        'started_ts': '1483315200000000',
        'status': 'COMPLETED',
        'tags': [],
        'utcnow_ts': '1483228800000000',
        'canary_preference': 'AUTO',
        'service_account': 'robot@example.com',
    }
    self.assertEqual(expected, api_common.build_to_dict(self.test_build))

  def test_format_luci_bucket(self):
    self.assertEqual(
        api_common.format_luci_bucket('chromium/try'), 'luci.chromium.try'
    )

  def test_parse_luci_bucket(self):
    self.assertEqual(
        api_common.parse_luci_bucket('luci.chromium.try'), 'chromium/try'
    )
    self.assertEqual(api_common.parse_luci_bucket('master.x'), '')
