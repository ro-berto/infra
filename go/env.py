#!/usr/bin/env python
# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Can be used to point environment variable to hermetic Go toolset.

Usage (on linux and mac):
$ eval `./env.py`
$ go version

Or it can be used to wrap a command:

$ ./env.py go version
"""

assert __name__ == '__main__'

import imp
import os
import subprocess
import sys

# Do not want to mess with sys.path, load the module directly.
bootstrap = imp.load_source(
    'bootstrap', os.path.join(os.path.dirname(__file__), 'bootstrap.py'))

old = os.environ.copy()
new = bootstrap.prepare_go_environ()

if len(sys.argv) == 1:
  for key, value in sorted(new.iteritems()):
    if old.get(key) != value:
      print 'export %s="%s"' % (key, value)
else:
  if sys.argv[1] == 'python':
    sys.argv[1] = sys.executable
  sys.exit(subprocess.call(sys.argv[1:], env=new))
