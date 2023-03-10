# Copyright (c) 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.


LUCICFG_ENTRY_SCRIPTS = ['main.star', 'dev.star']

USE_PYTHON3 = True


def CheckChangeOnUpload(input_api, output_api):
  tests = []
  for path in LUCICFG_ENTRY_SCRIPTS:
    tests += input_api.canned_checks.CheckLucicfgGenOutput(
        input_api, output_api, path)
  return input_api.RunTests(tests)


def CheckChangeOnCommit(input_api, output_api):
  return CheckChangeOnUpload(input_api, output_api)
