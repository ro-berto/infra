# Copyright 2019 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from PB.recipes.infra.windows_image_builder import windows_image_builder as wib
from PB.recipes.infra.windows_image_builder import (offline_winpe_customization
                                                    as winpe)
from PB.recipes.infra.windows_image_builder import actions
from PB.recipes.infra.windows_image_builder import sources
from PB.recipes.infra.windows_image_builder import dest

from recipe_engine.post_process import DropExpectation, StatusFailure
from recipe_engine.post_process import StatusSuccess, StepCommandRE
from RECIPE_MODULES.infra.windows_scripts_executor import test_helper as t

DEPS = [
    'depot_tools/gitiles',
    'recipe_engine/properties',
    'recipe_engine/platform',
    'recipe_engine/json',
    'recipe_engine/path',
    'recipe_engine/raw_io',
    'windows_scripts_executor',
]

PYTHON_VERSION_COMPATIBILITY = 'PY3'

PROPERTIES = wib.Image

# constants for use in test
image = 'general_tests'
cust_name = 'generic_cust'
arch = 'x86'
key = '4fa9269b1b8ebc0cd8d2c1c2415374819838ffb0a4a541a601ec51749b555096'


def RunSteps(api, config):
  api.windows_scripts_executor.init()
  custs = api.windows_scripts_executor.init_customizations(config)
  custs = api.windows_scripts_executor.process_customizations(custs, {})
  api.windows_scripts_executor.download_all_packages(custs)
  # mock cipd packages to avoid spooking add_file execution
  api.path.mock_add_paths(
      '[CACHE]\\Pkgs\\CIPDPkgs\\' +
      'resolved-instance_id-of-latest----------\\' +
      'infra_internal\\labs\\drivers\\microsoft\\' +
      'windows_adk\\winpe\\winpe-dot3svc\\windows-amd64', 'DIRECTORY')
  # mock cipd packages to avoid spooking add_file execution
  api.path.mock_add_paths(
      '[CACHE]\\Pkgs\\CIPDPkgs\\' +
      'resolved-instance_id-of-latest----------\\' +
      'infra_internal\\labs\\drivers\\microsoft\\' +
      'windows_adk\\winpe\\winpe-wmi\\windows-amd64', 'DIRECTORY')
  api.windows_scripts_executor.execute_customizations(custs)
  # mock existence of customization output to trigger upload
  api.path.mock_add_paths('[CLEANUP]\\{}\\workdir\\'.format(cust_name) +
                          'media\\sources\\boot.wim')


def GenTests(api):

  # actions for adding files from git
  ACTION_ADD_STARTNET = actions.Action(
      add_file=actions.AddFile(
          name='add_startnet_file',
          src=sources.Src(
              git_src=sources.GITSrc(
                  repo='chromium.dev',
                  ref='HEAD',
                  src='windows/artifacts/startnet.cmd'),),
          dst='Windows\\System32',
      ))

  STARTNET_URL = 'chromium.dev/+/ef70cb069518e6dc3ff24bfae7f195de5099c377/' +\
                 'windows/artifacts/startnet.cmd'

  # action to copy file from cipd
  ACTION_ADD_DOT3SVC = actions.Action(
      add_file=actions.AddFile(
          name='add winpe-dot3svc',
          src=sources.Src(
              cipd_src=sources.CIPDSrc(
                  package='infra_internal/labs/drivers/' +
                  'microsoft/windows_adk/winpe/winpe-dot3svc',
                  refs='latest',
                  platform='windows-amd64',
              ),),
          dst='Windows\\System32\\',
      ))

  DOT3SVC_URL = 'https://chrome-infra-packages.appspot.com/p/infra_internal/' +\
                'labs/drivers/microsoft/windows_adk/winpe/winpe-dot3svc/' +\
                'windows-amd64/+/resolved-instance_id-of-latest----------'

  # actions for installing windows packages
  ACTION_INSTALL_WMI = actions.Action(
      add_windows_package=actions.AddWindowsPackage(
          name='install_winpe_wmi',
          src=sources.Src(
              cipd_src=sources.CIPDSrc(
                  package='infra_internal/labs/drivers/' +
                  'microsoft/windows_adk/winpe/winpe-wmi',
                  refs='latest',
                  platform='windows-amd64',
              ),),
          args=['-IgnoreCheck'],
      ))

  # add cygwin. [b/213240613] Test for space in dest location
  ACTION_ADD_CYGWIN = actions.Action(
      add_file=actions.AddFile(
          name='add_startnet_file',
          src=sources.Src(
              git_src=sources.GITSrc(
                  repo='chromium.dev',
                  ref='HEAD',
                  src='windows/artifacts/cygwin.exe'),),
          dst='Program Files\\Cygwin',
      ))

  CYGWIN_URL = 'chromium.dev/+/ef70cb069518e6dc3ff24bfae7f195de5099c377/' +\
               'windows/artifacts/cygwin.exe'

  # Fail the step to gen winpe media folder using copy-pe
  # https://docs.microsoft.com/en-us/windows-hardware/manufacture/desktop/copype-command-line-options
  yield (api.test('Fail win image folder creation', api.platform('win', 64)) +
         # generate image with add starnet action
         api.properties(
             t.WPE_IMAGE(image, wib.ARCH_X86, cust_name, 'network_setup',
                         [ACTION_ADD_STARTNET])) +
         # mock pinning the file to current refs
         t.GIT_PIN_FILE(api, cust_name, 'HEAD',
                        'windows/artifacts/startnet.cmd', 'HEAD') +
         # mock failure in gen winpe media step
         t.GEN_WPE_MEDIA(api, arch, image, cust_name, False) +
         api.post_process(StatusFailure) +  # recipe should fail
         api.post_process(DropExpectation))

  # Missing arch in config. Execution should fail if arch is ARCH_UNSPECIFIED
  yield (api.test('Missing arch in config', api.platform('win', 64)) +
         api.properties(
             wib.Image(
                 name=image,
                 customizations=[
                     wib.Customization(
                         offline_winpe_customization=winpe
                         .OfflineWinPECustomization(name=cust_name,))
                 ])) +
         # recipe execution should fail
         api.post_process(StatusFailure) + api.post_process(DropExpectation))

  # Failure in executing action add_step
  yield (api.test('Fail add file step', api.platform('win', 64)) +
         # generate image with add starnet action
         api.properties(
             t.WPE_IMAGE(image, wib.ARCH_X86, cust_name, 'network_setup',
                         [ACTION_ADD_STARTNET])) +
         # mock all the init and deinit steps
         t.MOCK_WPE_INIT_DEINIT_FAILURE(api, arch, image, cust_name) +
         # mock git pin execution
         t.GIT_PIN_FILE(api, cust_name, 'HEAD',
                        'windows/artifacts/startnet.cmd', 'HEAD') +
         # mock add file from git to image execution
         t.ADD_FILE(api, image, cust_name, STARTNET_URL, False)
         +  # Fail to add file
         api.post_process(StatusFailure) +  # recipe fails
         api.post_process(DropExpectation))

  # Add a file from cipd storage
  yield (api.test('Add file from cipd', api.platform('win', 64)) +
         # generate image with add dot3svc action
         api.properties(
             t.WPE_IMAGE(image, wib.ARCH_X86, cust_name, 'network_setup',
                         [ACTION_ADD_DOT3SVC])) +
         # mock init and deinit steps for offline customization
         t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, arch, image, cust_name) +
         # mock add cipd file step
         t.ADD_FILE(api, image, cust_name, DOT3SVC_URL) +
         # assert that recipe completed execution successfully
         api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  # Add a file from git
  yield (api.test('Add file from git', api.platform('win', 64)) +
         # generate image with add starnet action
         api.properties(
             t.WPE_IMAGE(image, wib.ARCH_X86, cust_name, 'network_setup',
                         [ACTION_ADD_STARTNET])) +
         # mock all the init and deinit steps for winpe
         t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, arch, image, cust_name) +
         # mock git pin file
         t.GIT_PIN_FILE(api, cust_name, 'HEAD',
                        'windows/artifacts/startnet.cmd', 'HEAD') +
         # mock add file to image mount dir step
         t.ADD_FILE(api, image, cust_name, STARTNET_URL) +
         # recipe execution should pass successfully
         api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  # install cab file from cipd
  yield (api.test('Install package from cipd', api.platform('win', 64)) +
         # generate image with install wmi action
         api.properties(
             t.WPE_IMAGE(image, wib.ARCH_X86, cust_name, 'wmi_setup',
                         [ACTION_INSTALL_WMI])) +
         # mock all the init and deinit steps for winpe
         t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, arch, image, cust_name) +
         # mock install file step
         t.INSTALL_FILE(api, 'install_winpe_wmi', image, cust_name) +
         # recipe should pass successfully
         api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  GCS_IMAGE = t.WPE_IMAGE(image, wib.ARCH_X86, cust_name, 'no_change', [])
  # Add custom image src to modify
  GCS_IMAGE.customizations[0].offline_winpe_customization.image_src.CopyFrom(
      sources.Src(
          gcs_src=sources.GCSSrc(
              bucket='chrome-gce-images',
              source='WIB-WIM/winpe_stable.zip',
          )))

  yield (api.test('Modify winpe image from GCS', api.platform('win', 64)) +
         # start recipe with the custom GCS image src
         api.properties(GCS_IMAGE) +
         # mock all the init and deinit steps
         t.MOCK_CUST_IMG_WPE_INIT_DEINIT_SUCCESS(api, key, arch, image,
                                                 cust_name) +
         api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  CIPD_IMAGE = t.WPE_IMAGE(image, wib.ARCH_X86, cust_name, 'no_change', [])
  # Add custom image src to modify
  CIPD_IMAGE.customizations[0].offline_winpe_customization.image_src.CopyFrom(
      sources.Src(
          cipd_src=sources.CIPDSrc(
              package='infra_internal/labs/windows/gce',
              refs='stable',
              platform='windows-amd64')))

  yield (api.test('Modify winpe image from CIPD', api.platform('win', 64)) +
         # start recipe with the custom GCS image src
         api.properties(CIPD_IMAGE) +
         # mock all the init and deinit steps
         t.MOCK_CUST_IMG_WPE_INIT_DEINIT_SUCCESS(api, key, arch, image,
                                                 cust_name) +
         api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  # [b/213240613] Test for spaces in Add_File destination
  pkg_path = '\[CACHE\]\\\\Pkgs\\\\GCSPkgs\\\\'
  zip_path = pkg_path + 'chrome-gce-images\\\\WIB-WIM\\\\{}.zip'
  yield (
      api.test('Escape spaces in path test', api.platform('win', 64)) +
      # generate image with install wmi action
      api.properties(
          t.WPE_IMAGE(image, wib.ARCH_X86, cust_name, 'wmi_setup',
                      [ACTION_ADD_CYGWIN])) +
      # mock all the init and deinit steps
      t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, arch, image, cust_name) +
      # mock git pin file
      t.GIT_PIN_FILE(api, cust_name, 'HEAD', 'windows/artifacts/cygwin.exe',
                     'HEAD') +
      # mock add file to image mount dir step
      t.ADD_FILE(api, image, cust_name, CYGWIN_URL) +
      # assert that the generated image was uploaded
      t.CHECK_GCS_UPLOAD(api, image, cust_name, zip_path.format(key),
                         'gs://chrome-gce-images/WIB-WIM/{}.zip'.format(key)) +
      t.CHECK_ADD_FILE(
          api, image, cust_name, CYGWIN_URL,
          "\"\[CLEANUP\]\\\\generic_cust\\\\workdir\\\\mount\\\\Program Files"
          "\\\\Cygwin\"") +
      # recipe should pass successfully
      api.post_process(StatusSuccess) + api.post_process(DropExpectation))

  # generate happy path image with custom destination for upload
  HAPPY_PATH_IMAGE = t.WPE_IMAGE(image, wib.ARCH_X86, cust_name,
                                 'network_setup',
                                 [ACTION_ADD_STARTNET, ACTION_ADD_DOT3SVC])
  HAPPY_PATH_IMAGE.customizations[
      0].offline_winpe_customization.image_dests.append(
          dest.Dest(
              gcs_src=sources.GCSSrc(
                  bucket='test-bucket',
                  source='out/gce_winpe_rel.zip',
              )))

  wpeg = pkg_path + 'test-bucket\\\\out\\\\gce_winpe_rel.zip'

  # Generic happy path for recipe execution
  yield (
      api.test('Happy path', api.platform('win', 64)) +
      # start recipe with happy path image
      api.properties(HAPPY_PATH_IMAGE) +
      # mock all the init and deinit steps
      t.MOCK_WPE_INIT_DEINIT_SUCCESS(api, key, arch, image, cust_name) +
      # mock git pin file
      t.GIT_PIN_FILE(api, cust_name, 'HEAD', 'windows/artifacts/startnet.cmd',
                     'HEAD') +
      # mock add file to image mount dir step
      t.ADD_FILE(api, image, cust_name, STARTNET_URL) +
      # mock add file to image mount dir step
      t.ADD_FILE(api, image, cust_name, DOT3SVC_URL) +
      # assert that the generated image was uploaded
      t.CHECK_GCS_UPLOAD(api, image, cust_name, zip_path.format(key),
                         'gs://chrome-gce-images/WIB-WIM/{}.zip'.format(key)) +
      # assert that the generated image was uploaded to custom dest
      t.CHECK_GCS_UPLOAD(
          api,
          image,
          cust_name,
          wpeg,
          'gs://test-bucket/out/gce_winpe_rel.zip',
          orig='gs://chrome-gce-images/WIB-WIM/{}.zip'.format(key)) +
      # recipe should pass successfully
      api.post_process(StatusSuccess) + api.post_process(DropExpectation))
