# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import collections
import glob
import os
import platform
import re
import shutil
import stat
import subprocess
import sys

from . import build_platform
from . import concurrency
from . import util

from .build_types import Wheel


class PlatformNotSupported(Exception):
  """Exception raised by Builder.build when the specified wheel's platform is
  not support."""


class Builder(object):

  def __init__(self, spec, arch_map=None, abi_map=None,
               only_plat=None, skip_plat=None):
    """Initializes a new wheel Builder.

    spec (Spec): The wheel specification.
    arch_map (dict or None): Naming map for architectures. If the current
        platform has an entry in this map, the generated wheel will use the
        value as the "platform" field.
    abi_map (dict or None): Naming map for ABI. If the current platform
        has an entry in this map, the generated wheel will use the
        value as the "abi" field.
    only_plat (iterable or None): If not None, this Builder will only declare
        that it can build for the named platforms.
    skip_plat (iterable or None): If not None, this Builder will avoid declaring
        that it can build for the named platforms.
    version_fn (callable or None): If not None, and spec.version is None, this
        function will be used to set the spec version at runtime.
    """
    if arch_map:
      assert not any(isinstance(v, str) for v in arch_map.values()), (
          'arch_map must map str->seq[str]')

    self._spec = spec
    self._arch_map = arch_map or {}
    self._abi_map = abi_map or {}
    self._only_plat = frozenset(only_plat or ())
    self._skip_plat = frozenset(skip_plat or ())

    def _valid_platform(p):
      return p in build_platform.NAMES

    # Validate the platforms.
    for p in self._arch_map.keys():
      assert _valid_platform(
          p), "Wheel %r has invalid platform %s in arch_map" % (self._spec, p)
    for p in self._abi_map.keys():
      assert _valid_platform(
          p), "Wheel %r has invalid platform %s in abi_map" % (self._spec, p)
    for p in self._only_plat:
      assert _valid_platform(
          p), "Wheel %r has invalid platform %s in only_plat" % (self._spec, p)
    for p in self._skip_plat:
      assert _valid_platform(
          p), "Wheel %r has invalid platform %s in skip_plat" % (self._spec, p)

  def build_fn(self, system, wheel, output_dir):
    """Must be overridden by the subclass.

    Args:
      system (runtime.System)
      wheel (types.Wheel) - The wheel we're attempting to build.
      output_dir (str) The output directory for the wheel.

    Returns:
      None - The `wheel` argument will be uploaded in the CIPD package; the
        build_fn was expected to produce the wheel indicated by
        `wheel.filename`.
      list[types.Wheel] - This list of wheels will be uploaded in the CIPD
        package, instead of `wheel`.
    """
    raise NotImplementedError('Subclasses must implement build_fn')

  def version_fn(self, system):  # pylint: disable=unused-argument
    """Optionally overridden by the subclass.

    Returns:
      str - The version of the wheel to build.  By default this is
      `spec.version`.
    """
    return self._spec.version

  def md_data_fn(self):
    """Optionally overridden by the subclass.

    Returns:
      list[str] - Extra details to include in generated markdown files.
    """
    return []

  @property
  def spec(self):
    return self._spec

  def wheel(self, system, plat):
    spec = self.spec._replace(version=self.version_fn(system))

    # Make sure universal wheels are only built on universal platforms. This
    # allows us to explicitly select which builders build universal wheels, so
    # we have a disjoint set of wheels for each builder and hence avoid races.
    if spec.universal and not plat.universal:
      raise PlatformNotSupported(
          "Wheel %s is universal, but platform [%s] is not" %
          (spec.name, plat.name))
    # Conversely, universal platforms can only build universal wheels.
    if not spec.universal and plat.universal:
      raise PlatformNotSupported(
          "Wheel %s is not universal, but platform [%s] is" %
          (spec.name, plat.name))

    if spec.universal:
      pyversion = '38'
    else:
      # e.g. cp311 -> 311
      pyversion = plat.wheel_abi[2:]

    wheel = Wheel(
        spec=spec,
        plat=plat,
        download_plat=None,
        pyversion=pyversion,
        filename=None,
        md_lines=self.md_data_fn())

    # Determine our package's wheel filename. This incorporates "abi" and "arch"
    # override maps, which are a priori knowledge of the package repository's
    # layout. This can differ from the local platform value if the package was
    # valid and built for multiple platforms, which seems to happen on Mac a
    # lot.
    download_plat = plat._replace(
        wheel_abi=self._abi_map.get(plat.name, plat.wheel_abi),
        wheel_plat=self._arch_map.get(plat.name, plat.wheel_plat),
    )
    plat_wheel = wheel._replace(plat=download_plat)
    return wheel._replace(
        download_plat=download_plat,
        filename=plat_wheel.default_filename(),
    )

  def supported(self, plat):
    if self._only_plat and plat.name not in self._only_plat:
      return False
    if plat.name in self._skip_plat:
      return False
    if self.spec.universal != plat.universal:
      return False

    return True

  def build(self, wheel, system, rebuild=False):
    if not self.supported(wheel.plat):
      raise PlatformNotSupported()

    pkg_path = os.path.join(system.pkg_dir, '%s.pkg' % (wheel.filename,))
    if not rebuild and os.path.isfile(pkg_path):
      util.LOGGER.info('Package is already built: %s', pkg_path)
      return pkg_path

    # Rebuild the wheel, if necessary. Get their ".whl" file paths.
    built_wheels = self.build_wheel(
        wheel, system, system.wheel_dir, rebuild=rebuild)
    wheel_paths = [w.path(system.wheel_dir) for w in built_wheels]

    # Create a CIPD package for the wheel. Give the wheel a universal filename
    # within the CIPD package.
    #
    # See "A Note on Universality" at the top.
    util.LOGGER.info('Creating CIPD package: %r => %r', wheel_paths, pkg_path)
    with system.temp_subdir('cipd_%s_%s' % wheel.spec.tuple) as tdir:
      for w in built_wheels:
        universal_wheel_path = os.path.join(tdir, w.universal_filename())
        with concurrency.PROCESS_SPAWN_LOCK.shared():
          shutil.copy(w.path(system.wheel_dir), universal_wheel_path)
      _, git_revision = system.check_run(['git', 'rev-parse', 'HEAD'])
      system.cipd.create_package(wheel.cipd_package(git_revision),
                                 tdir, pkg_path)

    return pkg_path

  def build_wheel(self, wheel, system, output_dir, rebuild=False):
    built_wheels = [wheel]
    wheel_path = wheel.path(output_dir)
    if rebuild or not os.path.isfile(wheel_path):
      # The build_fn may return an alternate list of wheels.
      built_wheels = self.build_fn(system, wheel, output_dir) or built_wheels
    else:
      util.LOGGER.info('Wheel is already built: %s', wheel_path)
    return built_wheels


def _FindWheelInDirectory(wheel_dir):
  """Finds the single wheel in wheel_dir and returns its full path."""
  # Find the wheel in "wheel_dir". We scan expecting exactly one wheel.
  wheels = glob.glob(os.path.join(wheel_dir, '*.whl'))
  assert len(wheels) == 1, 'Unexpected wheels: %s' % (wheels,)
  return wheels[0]


def StageWheelForPackage(system, wheel_dir, wheel, output_dir):
  """Finds the single wheel in wheel_dir and copies it to the filename indicated
  by wheel.filename. Returns the name of the wheel we found.
  """
  dst = wheel.path(output_dir)

  source_path = _FindWheelInDirectory(wheel_dir)
  util.LOGGER.debug('Identified source wheel: %s', source_path)
  with concurrency.PROCESS_SPAWN_LOCK.shared():
    shutil.copy(source_path, dst)

  return os.path.basename(source_path)


def BuildPackageFromPyPiWheel(system, wheel, output_dir):
  """Builds a wheel by obtaining a matching wheel from PyPi."""
  with system.temp_subdir('%s_%s' % wheel.spec.tuple) as tdir:
    # TODO: This can copy the Python package unnecessarily: even if the platform
    # uses docker, we don't run 'pip download' using docker, but
    # SetupPythonPackages doesn't know this, and infers that we need it for
    # docker based on the platform, and hence copies it to the temporary
    # directory.
    #
    # Ignore the environment returned by SetupPythonPackages.
    # Setting PYTHONHOME interferes with pip invoking lsb_release_info,
    # and we don't need this anyway since we aren't compiling.
    interpreter, _ = SetupPythonPackages(system, wheel, tdir, tdir)
    util.check_run(
        system,
        None,
        tdir,
        [
            interpreter,
            '-m',
            'pip',
            'download',
            '--no-deps',
            '--only-binary=:all:',
            '--abi=%s' % (wheel.abi,),
            '--python-version=%s' % (wheel.pyversion,),
        ] + ['--platform=%s' % p for p in wheel.download_platform] + [
            '%s==%s' % (wheel.spec.name, wheel.spec.version),
        ],
        cwd=tdir)

    wheel_name = StageWheelForPackage(system, tdir, wheel, output_dir)

    # Make sure the prebuilt wheel supports the declared pyversions.
    if wheel.spec.universal:
      supported_versions = set(re.findall(r'py\d', wheel_name))
      required_versions = set(wheel.spec.pyversions or ['py3'])
      if not required_versions.issubset(supported_versions):
        raise Exception(
            'Wheel %s requires [%s], but wheel %s only supports [%s]' % (
                wheel.spec.name,
                ', '.join(sorted(required_versions)),
                wheel_name,
                ', '.join(sorted(supported_versions)),
            ))


def HostCipdPlatform():
  """Return the CIPD platform for the host system.

  We need this to determine which Python CIPD package to download. We can't just
  use cipd_platform from the platform, because some platforms are
  cross-compiled.
  """

  ARM64_MACHINES = {'aarch64_be', 'aarch64', 'armv8b', 'armv8l', 'arm64'}
  system, machine = platform.system(), build_platform.NativeMachine()

  # Note that we don't need to match all possible values or combinations of
  # 'system' and 'machine', only the ones we actually have Python interpreters
  # for in CIPD.
  if system == 'Linux':
    if machine == 'x86_64':
      return 'linux-amd64'
    if machine in ARM64_MACHINES:
      return 'linux-arm64'
  elif system == 'Windows':
    if machine == 'AMD64':
      return 'windows-amd64'
    if machine in {'i386', 'i686'}:
      return 'windows-386'
  elif system == 'Darwin':
    if machine == 'x86_64':
      return 'mac-amd64'
    if machine in ARM64_MACHINES:
      return 'mac-arm64'

  raise Exception("No CIPD platform for %s, %s" % (system, machine))


CIPD_PYTHON_LOCK = concurrency.KeyedLock()


def _FixSysconfigPaths(python_dir, wheel_platform, work_root):
  # For POSIX platforms (everything except Windows), there is a
  # _sysconfigdata module generated at build time that can be queried
  # for the various include and lib paths for the Python installation.
  #
  # In our 3pp cpython3 packages, we replace these with the string
  # "[INSTALL_PREFIX]" for easy substitution with the real install
  # location.
  if wheel_platform.name.startswith('windows'):
    # Nothing needs to be done for Windows.
    return

  # When using docker, translate the Python directory so that it
  # works inside the container.
  if wheel_platform.dockcross_base:
    python_dir = '/work/' + os.path.relpath(python_dir, work_root)

  for config_data in glob.iglob('%s/lib/python*/_sysconfigdata*.py' %
                                python_dir):
    output_lines = []
    with open(config_data, 'r') as f:
      for line in f:
        output_lines.append(line.replace('[INSTALL_PREFIX]', python_dir))
    st = os.stat(config_data)
    os.chmod(config_data, st.st_mode | stat.S_IWUSR)
    with open(config_data, 'w') as f:
      f.writelines(output_lines)


# Track which python packages we have installed during this invocation.
# Each package is reinstalled the first time it is used during each
# dockerbuild run, to pick up version changes.
_initialized_cipd_python = set()


def _InstallCipdPythonPackage(system, cipd_platform, wheel, base_dir,
                              work_root):
  PY_CIPD_VERSION_MAP = {
      '38': 'version:2@3.8.10.chromium.25',
      '311': 'version:2@3.11.0.chromium.23',
  }

  package_ident = 'py_%s_%s' % (wheel.pyversion, cipd_platform)
  common_dir = os.path.join(system.root, 'cipd_python', package_ident)

  # Lock to prevent races with other threads trying to install the same CIPD
  # package. The cipd_python CIPD root is shared between all threads, so we need
  # to ensure we don't try to install the same package to it multiple times
  # concurrently.
  with CIPD_PYTHON_LOCK.get(package_ident):
    if package_ident not in _initialized_cipd_python:
      if os.path.exists(common_dir):
        shutil.rmtree(common_dir)
      os.makedirs(common_dir)
      system.cipd.init(common_dir)

      cipd_pkg = 'infra/3pp/tools/cpython3/%s' % cipd_platform
      version = PY_CIPD_VERSION_MAP[wheel.pyversion]
      system.cipd.install(cipd_pkg, version, common_dir)
      # Note: this assumes we use CIPD install_mode 'copy' for the cpython3
      # package, or else we might be modifying a shared cache copy.
      _FixSysconfigPaths(common_dir, wheel.plat, work_root)
      _initialized_cipd_python.add(package_ident)

  # For docker builds, we need the Python interpreter in the base working dir
  # for this particular build. Ideally we'd just mount the cipd_python directory
  # and avoid this copy, but that would require a lot of plumbing. We'll live
  # with it for now.
  if wheel.plat.dockcross_base:
    pkg_dir = os.path.join(
        base_dir, 'cipd_python_%s_%s' % (wheel.pyversion, cipd_platform))
    # Grab a shared lock on the system process spawn RWLock while copying
    # binaries. We copy to a subdirectory of `base_dir', which is a unique,
    # per-build temporary directory. Since each directory is unique, we can copy
    # to multiple directories in parallel without a problem.
    #
    # This one is particularly important as we're copying Python binaries, which
    # are large enough to give a high chance of in-flight copies when a process
    # is spawned, and are executed, which can't be done when they're open in
    # write mode.
    with concurrency.PROCESS_SPAWN_LOCK.shared():
      # Don't copy the '__pycache__' bytecode cache directory, to avoid races
      # between Python processes running out of the common dir and us.
      shutil.copytree(common_dir, pkg_dir, ignore=lambda *_: {'__pycache__'})
      # We must fix up paths every time the installation is copied.
      _FixSysconfigPaths(pkg_dir, wheel.plat, work_root)
  else:
    pkg_dir = common_dir

  interpreter = os.path.join(pkg_dir, 'bin', 'python3')
  return pkg_dir, interpreter


def SetupPythonPackages(system, wheel, base_dir, work_root):
  """Installs python package(s) from CIPD and sets up the build environment.

  Args:
     system (System): A System object.
     wheel (Wheel): The Wheel object to install a build environment for.
     base_dir (str): The top-level build directory for the wheel.
     work_root (str): The root working directory, used for dockcross.

  Returns: A tuple (path to the python interpreter to run,
                    dict of environment variables to be set).
  """
  host_platform = HostCipdPlatform()

  # Building Windows x86 on Windows x64 is a special case. In this situation,
  # we want to directly install and run the windows-x86 python package. This
  # is because some wheels use properties of the Python interpreter (e.g.
  # sys.maxsize) to detect whether to build for 32-bit or 64-bit CPUs.
  if (host_platform == 'windows-amd64' and
      wheel.plat.cipd_platform == 'windows-386'):
    host_platform = 'windows-386'

  _, interpreter = _InstallCipdPythonPackage(system, host_platform, wheel,
                                             base_dir, work_root)
  env = wheel.plat.env.copy()

  # If we are cross-compiling, also install the target-platform python and set
  # PYTHONHOME to point to it. This will ensure that we use the correct
  # compiler and linker command lines which are generated at build time in the
  # sysconfigdata module.
  if not wheel.spec.universal and host_platform != wheel.plat.cipd_platform:
    pkg_dir, _ = _InstallCipdPythonPackage(system, wheel.plat.cipd_platform,
                                           wheel, base_dir, work_root)
    env['PYTHONHOME'] = pkg_dir
    # Set _PYTHON_SYSCONFIGDATA_NAME to point to the target-architecture
    # sysconfig module.
    sysconfigdata_modules = glob.glob('%s/lib/python%s/_sysconfigdata_*.py' %
                                      (pkg_dir, '.'.join(wheel.pyversion)))
    if len(sysconfigdata_modules) != 1:
      raise Exception(
          'Expected 1 sysconfigdata module in python package ' +
          'for %s, got: [%s]',
          (wheel.plat.cipd_platform, ','.join(sysconfigdata_modules)))
    env['_PYTHON_SYSCONFIGDATA_NAME'] = (os.path.basename(
        sysconfigdata_modules[0])[:-3])  # remove .py

  # Make sure not to pick up any extra host python modules.
  env['PYTHONPATH'] = ''

  return interpreter, env


class BuildDependencies(
    collections.namedtuple(
        'BuildDependencies',
        (
            # remote (List[str]): Wheels fetched from pip repository
            'remote',

            # local (List[Builder]): Wheels installed from local path
            'local',
        ))):
  pass


def BuildPackageFromSource(system,
                           wheel,
                           src,
                           output_dir,
                           src_filter=None,
                           deps=None,
                           tpp_libs=None,
                           tpp_tools=None,
                           env=None,
                           skip_auditwheel=False):
  """Creates Python wheel from src.

  Args:
    system (dockerbuild.runtime.System): Represents the local system.
    wheel (dockerbuild.wheel.Wheel): The wheel to build.
    src (dockerbuild.source.Source): The source to build the wheel from.
    output_dir (str): The wheel output directory.
    deps (dockerbuild.builder.BuildDependencies|None): Dependencies required
      to build the wheel.
    tpp_libs (List[(str, str)]|None): 3pp libraries to install in the
      build environment. The list items are (package name, version).
    tpp_tools (List[(str, str)]|None): 3pp tools (for the build platform)
      which are needed to build the wheel. The list items are
      (package name, version).
    env (Dict[str, str]|None): Additional envvars to set while building the
      wheel.
    skip_auditwheel (bool): See SourceOrPrebuilt documentation.
  """
  dx = system.dockcross_image(wheel.plat)
  with system.temp_subdir('%s_%s' % wheel.spec.tuple) as tdir:
    build_dir = system.repo.ensure(src, tdir, unpack_file_filter=src_filter)

    for patch in src.get_patches():
      util.LOGGER.info('Applying patch %s', os.path.basename(patch))
      cmd = ['git', 'apply', '-p1', patch]
      subprocess.check_call(cmd, cwd=build_dir)

    interpreter, extra_env = SetupPythonPackages(system, wheel, build_dir, tdir)

    if dx.platform:
      extra_env.update({
          'host_alias': dx.platform.cross_triple,
      })
    if env:
      extra_env.update(env)

    with util.tempdir(tdir, 'deps') as tdeps:
      env_prefix = []

      # Install the given package and return the path to it.
      def install_pkg(pkg, plat, build_dir=build_dir):
        pkg_name = '%s/%s' % (pkg, plat)
        pkg_dir = os.path.join(build_dir, '%s_%s_cipd' % (pkg, plat))
        system.cipd.init(pkg_dir)
        system.cipd.install(pkg_name, version, pkg_dir)
        return pkg_dir

      if tpp_libs:
        cflags = extra_env.get('CFLAGS', '')
        cxxflags = extra_env.get('CXXFLAGS', '')
        ldflags = extra_env.get('LDFLAGS', '')
        for pkg, version in tpp_libs:
          pkg_dir = install_pkg(pkg, wheel.plat.cipd_platform)
          cflags += ' -I' + os.path.join(pkg_dir, 'include')
          cxxflags += ' -I' + os.path.join(pkg_dir, 'include')
          ldflags += ' -L' + os.path.join(pkg_dir, 'lib')
          # Prepend the bin/ directory of each package to PATH.
          env_prefix.append(('PATH', os.path.join(pkg_dir, 'bin')))
          # Make sure the lib/ directory is also in LD_LIBRARY PATH in case of
          # shared libraries. This can be needed below when running auditwheel.
          env_prefix.append(('LD_LIBRARY_PATH', os.path.join(pkg_dir, 'lib')))

        extra_env['CFLAGS'] = cflags.lstrip()
        extra_env['CXXFLAGS'] = cxxflags.lstrip()
        extra_env['LDFLAGS'] = ldflags.lstrip()

      if tpp_tools:
        host_plat = HostCipdPlatform()
        for pkg, version in tpp_tools:
          pkg_dir = install_pkg(pkg, host_plat)
          # Prepend the bin/ directory of each package to PATH.
          env_prefix.append(('PATH', os.path.join(pkg_dir, 'bin')))

      if deps:
        util.copy_to(util.resource_path('generate_pyproject.py'), build_dir)
        cmd = PrepareBuildDependenciesCmd(system, wheel, build_dir, tdeps, deps)
        util.check_run(
            system, dx, tdir, [interpreter] + cmd, cwd=build_dir, env=extra_env)

      util.copy_to(util.resource_path('constraints.txt'), build_dir)
      # Use the environment variable, rather than the --constraint flag,
      # otherwise the constraints are not passed through when pip installs
      # dependencies (https://github.com/pypa/pip/issues/8439).
      extra_env['PIP_CONSTRAINT'] = 'constraints.txt'

      # Use the per-thread pip cache, see comments in runtime.py
      extra_env['PIP_CACHE_DIR'] = system.pip_cache_dir

      # TODO: Perhaps we should pass --build-option --universal for universal
      # py2.py3 wheels, to ensure the filename is right even if the wheel
      # setup.cfg isn't configured correctly.
      cmd = [
          interpreter,
          '-m',
          'pip',
          'wheel',
          '--no-deps',
          '--only-binary=:all:',
          '--wheel-dir',
          tdir,
          '.',
      ]

      util.check_run(
          system,
          dx,
          tdir,
          cmd,
          cwd=build_dir,
          env=extra_env,
          env_prefix=env_prefix)

    wheel_dir = tdir
    if not skip_auditwheel and wheel.plat.wheel_plat[0].startswith('manylinux'):
      wheel_dir = os.path.join(tdir, 'auditwheel-out')
      util.check_run(
          system,
          dx,
          tdir, [
              'auditwheel',
              'repair',
              '--no-update-tags',
              '--plat',
              wheel.plat.wheel_plat[0],
              '-w',
              wheel_dir,
              _FindWheelInDirectory(tdir),
          ],
          cwd=tdir,
          env=extra_env,
          env_prefix=env_prefix)

    StageWheelForPackage(system, wheel_dir, wheel, output_dir)


def PrepareBuildDependenciesCmd(system, wheel, build_dir, deps_dir, deps):
  local_wheels = []
  # For build requirements, we should install wheels corresponding to the host
  # environment. Currently this only matters on Linux, where we cross-compile.
  if sys.platform.startswith('linux'):
    host_plat = build_platform.ALL.get({
        'cp38': 'manylinux-x64-py3.8',
        'cp311': 'manylinux-x64-py3.11',
    }[wheel.plat.wheel_abi])
  else:
    host_plat = wheel.plat
  for dep in deps.local:
    dep_wheel = dep.wheel(system, host_plat)
    # build_wheel may return sub wheels different from the original wheel.
    # e.g. the result of a MultiWheel's build is a list of all included wheels.
    subs = dep.build_wheel(dep_wheel, system, deps_dir, rebuild=True)
    for sub in subs:
      sub_path = sub.path(deps_dir)
      sub_rel = os.path.relpath(sub_path, build_dir)
      local_wheels.append('{}@{}'.format(sub.spec.name, sub_rel))

  # Generate pyproject.toml to control build dependencies
  cmd = ['generate_pyproject.py']
  if deps.remote:
    cmd.append('--remotes')
    cmd.extend(deps.remote)
  if deps.local:
    cmd.append('--locals')
    cmd.extend(local_wheels)
  return cmd
