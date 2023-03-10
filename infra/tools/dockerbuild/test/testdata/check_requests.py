#!/usr/bin/env vpython3
# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# This test is only run under python3, but this is needed to keep
# pylint 1.5 happy.
from __future__ import print_function

import os
import requests
import sys


# Cases to check when running locally and on a bot.
TEST_CASES = [
    # OK.
    ('https://clients3.google.com/generate_204', 204),
]


# These are checked only when running not on a bot. They are too flaky to run
# on a bot as presubmit or postsubmit tests.
NON_BOT_TEST_CASES = [
    # The cert for extended-validation.badssl.com (ironically?) expired. Disable
    # its test in the meantime.
    #('https://extended-validation.badssl.com', 200),

    # Bad certs.
    ('https://expired.badssl.com', requests.exceptions.SSLError),
    ('https://wrong.host.badssl.com', requests.exceptions.SSLError),
    ('https://self-signed.badssl.com', requests.exceptions.SSLError),
    ('https://untrusted-root.badssl.com', requests.exceptions.SSLError),

    # 'requests' is known to accept revoked certificates.
    # https://github.com/kennethreitz/requests/issues/3770
    ('https://revoked.badssl.com', requests.exceptions.SSLError),
]


def get_code_or_err(url):
  try:
    print('Trying %s' % url)
    return requests.get(url).status_code
  except requests.exceptions.SSLError as exc:
    return exc


def tests_succeed():
  cases = []
  cases.extend(TEST_CASES)
  if os.environ.get('SWARMING_HEADLESS') != '1':
    cases.extend(NON_BOT_TEST_CASES)

  ok = True
  for url, exp in cases:
    res = get_code_or_err(url)
    if isinstance(exp, int):
      if exp != res:
        print('For %s: expecting %d, got %s' % (url, exp, res), file=sys.stderr)
        ok = False
    elif not isinstance(res, exp):
      print('For %s: expecting %s, got %s' % (url, exp, res), file=sys.stderr)
      ok = False
  return ok


sys.exit(0 if tests_succeed() else 1)
