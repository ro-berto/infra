# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from PB.recipes.infra.windows_image_builder import (offline_winpe_customization
                                                    as winpe)
from PB.recipes.infra.windows_image_builder import windows_image_builder as wib
from PB.recipes.infra.windows_image_builder import actions
from PB.recipes.infra.windows_image_builder import sources
from PB.recipes.infra.windows_image_builder import dest

from recipe_engine.post_process import DropExpectation, StatusFailure
from recipe_engine.post_process import StatusSuccess, StepCommandRE
from RECIPE_MODULES.infra.windows_scripts_executor import test_helper as t

DEPS = [
    'windows_scripts_executor',
    'recipe_engine/path',
    'recipe_engine/properties',
    'recipe_engine/platform',
    'recipe_engine/json',
    'recipe_engine/raw_io'
]

PYTHON_VERSION_COMPATIBILITY = 'PY3'

PROPERTIES = wib.Image

image = 'cipd_test'
customization = 'cipd_add_file'
key = '835663538df204d1d6ba072b185850ba502e4520fddbfe2262562596511368af'


def RunSteps(api, config):
  api.windows_scripts_executor.init()
  custs = api.windows_scripts_executor.init_customizations(config)
  custs = api.windows_scripts_executor.process_customizations(custs, {})
  api.windows_scripts_executor.download_all_packages(custs)
  # mock existence of cipd files to avoid failures
  api.path.mock_add_paths(
      '[CACHE]\\Pkgs\\CIPDPkgs\\resolved-instance_id-of-latest----------' +
      '\\infra\\files\\cipd-1\\windows-amd64', 'DIRECTORY')
  # mock existence of cipd files to avoid failures
  api.path.mock_add_paths(
      '[CACHE]\\Pkgs\\CIPDPkgs\\resolved-instance_id-of-latest----------' +
      '\\infra\\files\\cipd-2\\windows-amd64', 'DIRECTORY')
  # mock existence of cipd files to avoid failures
  api.path.mock_add_paths(
      '[CACHE]\\Pkgs\\CIPDPkgs\\resolved-instance_id-of-latest----------' +
      '\\infra\\files\\cipd-3\\windows-amd64', 'DIRECTORY')
  api.windows_scripts_executor.execute_customizations(custs)
  # mock existence of customization output to trigger upload
  api.path.mock_add_paths('[CLEANUP]\\{}\\workdir\\'.format(customization) +
                          'media\\sources\\boot.wim')

def GenTests(api):
  # add file from cipd to winpe image action
  ACTION_ADD_CIPD_1 = actions.Action(
      add_file=actions.AddFile(
          name='add cipd-1',
          src=sources.Src(
              cipd_src=sources.CIPDSrc(
                  package='infra/files/cipd-1',
                  refs='latest',
                  platform='windows-amd64',
              ),),
          dst='Windows\\Users\\',
      ))

  CIPD_1_URL = 'https://chrome-infra-packages.appspot.com/p/infra/files/' +\
               'cipd-1/windows-amd64/+/' +\
               'resolved-instance_id-of-latest----------'

  # add file from cipd to winpe image action
  ACTION_ADD_CIPD_2 = actions.Action(
      add_file=actions.AddFile(
          name='add cipd-2',
          src=sources.Src(
              cipd_src=sources.CIPDSrc(
                  package='infra/files/cipd-2',
                  refs='latest',
                  platform='windows-amd64',
                  filename='cipd-2.py'),),
          dst='Windows\\Users\\',
      ))

  CIPD_2_URL = 'https://chrome-infra-packages.appspot.com/p/infra/files/' +\
              'cipd-2/windows-amd64/+/' +\
              'resolved-instance_id-of-latest----------'

  # add file from cipd to winpe image action
  ACTION_ADD_CIPD_3 = actions.Action(
      add_file=actions.AddFile(
          name='add cipd-3',
          src=sources.Src(
              cipd_src=sources.CIPDSrc(
                  package='infra/files/cipd-3',
                  refs='latest',
                  platform='windows-amd64',
              ),),
          dst='Windows\\Users\\',
      ))

  CIPD_3_URL = 'https://chrome-infra-packages.appspot.com/p/infra/files/' +\
               'cipd-3/windows-amd64/+/' +\
               'resolved-instance_id-of-latest----------'

  UPLOAD_TO_CIPD_1 = dest.Dest(
      cipd_src=sources.CIPDSrc(
          package='experimental/mock/wib/test-1',
          refs='latest',
          platform='windows-amd64',
      ),
      tags={
          'version': '0.0.1',
          'type': 'vanilla'
      })
  UPLOAD_TO_CIPD_2 = dest.Dest(
      cipd_src=sources.CIPDSrc(
          package='experimental/mock/wib/test-2',
          refs='latest',
          platform='windows-amd64',
      ),
      tags={
          'version': '0.0.1',
          'type': 'vanilla'
      })

  yield (
      api.test('Test cipd pin and download package', api.platform('win', 64)) +
      # image with an action to add file from cipd
      api.properties(
          t.WPE_IMAGE(image, wib.ARCH_X86, customization, 'add pkg from cipd',
                      [ACTION_ADD_CIPD_1])) +
      # mock init and deinit steps
      t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, 'x86', image, customization) +
      # mock add cipd file step
      t.ADD_FILE(api, image, customization, CIPD_1_URL) +
      # assert that the recipe was executed successfully
      api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  yield (
      api.test('Test cipd pin and download packages in single action',
               api.platform('win', 64)) +
      # image with two different actions to add files from cipd
      api.properties(
          t.WPE_IMAGE(image, wib.ARCH_X86, customization, 'add pkg from cipd',
                      [ACTION_ADD_CIPD_1, ACTION_ADD_CIPD_2])) +
      # mock init and deinit steps
      t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, 'x86', image, customization) +
      # mock add cipd file step
      t.ADD_FILE(api, image, customization, CIPD_1_URL) +
      # mock add cipd file step
      t.ADD_FILE(api, image, customization, CIPD_2_URL) +
      # assert that the recipe execution was a success
      api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  # image with multiple sub customization
  CIPD_PACKAGE_MULTIPLE_ACTIONS = t.WPE_IMAGE(
      image, wib.ARCH_X86, customization, 'add cipd pkgs',
      [ACTION_ADD_CIPD_1, ACTION_ADD_CIPD_2])
  cust = CIPD_PACKAGE_MULTIPLE_ACTIONS.customizations[0]
  cust.offline_winpe_customization.offline_customization.append(
      actions.OfflineAction(name='add cipd pkg', actions=[ACTION_ADD_CIPD_3]))

  yield (api.test('Test cipd pin and download packages in multiple actions',
                  api.platform('win', 64)) +
         # use the image with multiple actions
         api.properties(CIPD_PACKAGE_MULTIPLE_ACTIONS) +
         # mock all the init and deinit steps
         t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, 'x86', image, customization) +
         # mock add cipd file step
         t.ADD_FILE(api, image, customization, CIPD_1_URL) +
         # mock add cipd file step
         t.ADD_FILE(api, image, customization, CIPD_2_URL) +
         # mock add cipd file step
         t.ADD_FILE(api, image, customization, CIPD_3_URL) +
         # assert that the recipe executed successfully
         api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  pkg_path = '\[CACHE\]\\\\Pkgs\\\\GCSPkgs\\\\chrome-gce-images\\\\'
  zip_path = pkg_path + 'WIB-WIM\\\\{}.zip'
  yield (
      api.test('Test cipd upload package', api.platform('win', 64)) +
      # image with an action to add file from cipd
      api.properties(
          t.WPE_IMAGE(image, wib.ARCH_X86, customization, 'add pkg from cipd',
                      [ACTION_ADD_CIPD_1], [UPLOAD_TO_CIPD_1, UPLOAD_TO_CIPD_2])
      ) +
      # mock init and deinit steps
      t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, 'x86', image, customization) +
      # mock add cipd file step
      t.ADD_FILE(api, image, customization, CIPD_1_URL) +
      # assert that the generated image was uploaded
      t.CHECK_GCS_UPLOAD(api, image, customization, zip_path.format(key),
                         'gs://chrome-gce-images/WIB-WIM/{}.zip'.format(key)) +
      t.CHECK_CIPD_UPLOAD(api, image, customization, UPLOAD_TO_CIPD_1) +
      t.CHECK_CIPD_UPLOAD(api, image, customization, UPLOAD_TO_CIPD_2) +
      # assert that the recipe was executed successfully
      api.post_process(StatusSuccess) + api.post_process(DropExpectation))
