# Copyright 2022 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from recipe_engine.post_process import DropExpectation, StatusSuccess
from recipe_engine.recipe_api import Property

DEPS = [
    'qemu', 'recipe_engine/raw_io', 'recipe_engine/json', 'recipe_engine/path'
]

PYTHON_VERSION_COMPATIBILITY = 'PY3'


def RunSteps(api):
  api.qemu.powerdown_vm(name='test')


def GenTests(api):

  yield (api.test('Test powerdown_vm') + api.post_process(StatusSuccess) +
         api.step_data(
             'Powerdown test', api.json.output({
                 'return': {},
             }), retcode=0) + api.post_process(DropExpectation))

  yield (api.test('Test qemu fail') + api.post_process(StatusSuccess) +
         api.step_data(
             'Powerdown test',
             api.raw_io.output("""
              [No write since last change]
              Traceback (most recent call last):
              File \"/something/qemu/resources/qmp.py\", line 74, in <module>
                  main()
              File \"/something/qemu/resources/qmp.py\", line 58, in main
                  sock.connect(args.sock)
              FileNotFoundError: [Errno 2] No such file or directory
            """),
             retcode=1) + api.post_process(DropExpectation))