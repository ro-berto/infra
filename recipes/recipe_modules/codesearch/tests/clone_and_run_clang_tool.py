# Copyright 2017 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

DEPS = [
    'codesearch',
    'recipe_engine/path',
]


def RunSteps(api):
  api.codesearch.set_config('chromium', PROJECT='chromium')
  api.codesearch.clone_clang_tools(api.path['cache'])
  api.codesearch.run_clang_tool(clang_dir=None, run_dirs=[api.path['cache']])


def GenTests(api):
  yield api.test('basic')

  yield api.test(
      'run_translation_unit_clang_tool_failed',
      api.step_data('run translation_unit clang tool', retcode=1),
  )
