# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import os
import sys
import shutil
import subprocess
import tempfile
import unittest


PYTHON_TEST_CIPD_PACKAGE = None


# A "vpython" spec for a VirtualEnv containing "psutil".
VPYTHON_SPEC_PSUTIL = """\
wheel: <
  name: "infra/python/wheels/psutil/${platform}_${py_python}_${py_abi}"
  version: "version:5.2.2"
>
"""

VPYTHON_CIPD_PACKAGE = b'infra/tools/luci/vpython/${platform} git_revision:788e09418ced22000a9774e4549b9ad42084b814'

class TestPython(unittest.TestCase):

  # Public repository that uses HTTPS.
  HTTPS_REPO_URL = 'https://chromium.googlesource.com/infra/infra'

  @classmethod
  def setUpClass(cls):
    cls._is_windows = os.name == 'nt'
    cls._exe_suffix = '.exe' if cls._is_windows else ''

    cls.tdir = tempfile.mkdtemp(dir=os.getcwd(), suffix='test_python')

    cpython_dir = os.path.join(cls.tdir, '_cpython')
    subprocess.check_call([
      'cipd', 'pkg-deploy', PYTHON_TEST_CIPD_PACKAGE, '-root', cpython_dir],
      shell=cls._is_windows)

    vpython_dir = os.path.join(cls.tdir, '_vpython')
    subprocess.check_output([
      'cipd', 'ensure', '--ensure-file', '-', '-root', vpython_dir],
      input=VPYTHON_CIPD_PACKAGE, shell=cls._is_windows)

    cls.python = os.path.join(cpython_dir, 'bin', 'python' + cls._exe_suffix)
    cls.vpython = os.path.join(vpython_dir, 'vpython' + cls._exe_suffix)

  @classmethod
  def tearDownClass(cls):
    # If we fail to delete, try and make the path user-writable. This can happen
    # notably with "vpython" paths.
    def _make_writable(base, path):
      path = os.path.join(base, path)
      st = os.lstat(path)
      os.chmod(path, st.st_mode | 0o660)

    for dirpath, dirnames, filenames in os.walk(cls.tdir):
      for name in dirnames:
        _make_writable(dirpath, name)
      for name in filenames:
        _make_writable(dirpath, name)

    shutil.rmtree(cls.tdir, ignore_errors=True)

  def setUp(self):
    self.test_tdir = tempfile.mkdtemp(dir=self.tdir)
    self.env = os.environ.copy()

  def _write_file(self, content):
    fd = None
    try:
      fileno, path = tempfile.mkstemp(dir=self.test_tdir)
      fd = os.fdopen(fileno, 'w')
      fd.write(content)
      return path
    finally:
      if fd:
        fd.close()

  def test_version(self):
    output = subprocess.check_output([self.python, '--version'],
                                     stderr=subprocess.STDOUT).decode('utf-8')
    self.assertTrue(output.startswith('Python '))

    expected_version = os.environ['_3PP_VERSION']
    # On windows we don't append the patch version because we actually bundle
    # the official python release, so don't have an opportunity to change the
    # version string.
    if '_3PP_PATCH_VERSION' in os.environ and not self._is_windows:
      expected_version += '+' + os.environ['_3PP_PATCH_VERSION']

    self.assertEqual(output.lstrip('Python ').strip(), expected_version)

  def test_package_import(self):
    for pkg in (
        'ctypes', 'ssl', 'cStringIO', 'binascii', 'hashlib', 'sqlite3'):
      script = 'import %s; print %s' % (pkg, pkg)
      rv = subprocess.call([self.python, '-c', script])
      self.assertEqual(rv, 0, 'Could not import %r.' % (pkg,))

  def test_use_https(self):
    script = 'import urllib; print urllib.urlretrieve("%s")' % (
        self.HTTPS_REPO_URL)
    rv = subprocess.call([self.python, '-c', script])
    self.assertEqual(rv, 0)

  @unittest.skipIf(
    sys.platform == 'win32',
    'sqlite in windows distro is older'
  )
  def test_sqlite_version(self):
    script = (
        'import sqlite3; '
        'print ".".join(str(x) for x in sqlite3.sqlite_version_info)')
    proc = subprocess.Popen([self.python, '-c', script],
                            stdout=subprocess.PIPE,
                            stderr=subprocess.STDOUT)
    stdout, _ = proc.communicate()
    self.assertEqual(proc.returncode, 0)
    self.assertEqual(stdout.decode('utf-8').strip(),
                     '3.19.3')  # Matches sqlite3 CIPD package.

  def test_psutil_import(self):
    vpython_spec = self._write_file(VPYTHON_SPEC_PSUTIL)

    script = 'import psutil; print psutil'
    cmd = [
        self.vpython,
        '-vpython-interpreter',
        self.python,
        '-vpython-spec',
        vpython_spec,
        '-vpython-root',
        os.path.join(self.test_tdir, 'vpython'),
        '-c',
        script,
    ]
    rv = subprocess.call(cmd)
    self.assertEqual(rv, 0)

  def test_cryptography_import(self):
    self.env.update({
      # Cause abort() on heap error.
      'MALLOC_CHECK_': '2',
    })

    # This sequence of events can cause Python to freeze or break if linking is
    # messed up and "cryptography"'s OpenSSL is cross-linking with Python's
    # statically built-in OpenSSL.
    script = (
      'import OpenSSL; '
      'import cryptography.hazmat.bindings.openssl.binding')
    cmd = [
        self.vpython,
        '-vpython-interpreter',
        self.python,
        '-vpython-spec',
        os.path.join(os.path.dirname(__file__), 'example.vpython'),
        '-vpython-root',
        os.path.join(self.test_tdir, 'vpython'),
        '-c',
        script,
    ]
    proc = subprocess.Popen(
        cmd, stdout=subprocess.PIPE, stderr=subprocess.STDOUT, env=self.env)
    stdout, _ = proc.communicate()
    self.assertEqual(
        proc.returncode, 0,
        'Failed to import openssl libraries (return code %d). output:\n%s' %
        (proc.returncode, stdout.decode('utf-8')))

  def test_no_version_script_in_sysconfig(self):
    # On Linux, we use a linker version script to restrict the exported
    # symbols. Verify that this has not leaked into the build flags that
    # will be used by Python wheels.
    script = ('import sysconfig\n'
              'for k, v in sysconfig.get_config_vars().iteritems():\n'
              '  if isinstance(v, str):\n'
              '    assert "version-script" not in v, (\n'
              '      "Found unexpected version-script in %s: %s" % (k, v))')
    proc = subprocess.Popen([self.python, '-c', script],
                            stdout=subprocess.PIPE,
                            stderr=subprocess.STDOUT)
    stdout, _ = proc.communicate()
    self.assertEqual(proc.returncode, 0, stdout.decode('utf-8'))


if __name__ == '__main__':
  platform = os.environ['_3PP_PLATFORM']
  tool_platform = os.environ['_3PP_TOOL_PLATFORM']
  if 'windows' not in platform and platform != tool_platform:
    print('SKIPPING TESTS')
    print('  platform:', platform)
    print('  tool_platform:', tool_platform)
    sys.exit(0)

  PYTHON_TEST_CIPD_PACKAGE = sys.argv[1]
  sys.argv.pop(1)
  unittest.main()
