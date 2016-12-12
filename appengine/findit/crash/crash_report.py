# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
from collections import namedtuple

from crash.stacktrace import Stacktrace

class CrashReport(namedtuple('CrashReport',
    ['crashed_version', 'signature', 'platform', 'stacktrace',
     'regression_range'])):
  """A reported crash we want to analyze.

  This class comprises the inputs to the Predator library; as distinguished
  from the Culprit class, which comprises the outputs/results of Predator's
  analyses. N.B., the appengine clients conflate input and output into
  a single CrashAnalysis(ndb.Model) class, but that's up to them; in
  the library we keep inputs and outputs entirely distinct.

  Args:
    crashed_version (str): The version of Chrome in which the crash occurred.
    signature (str): The signature of the crash on the Chrome crash server.
    platform (str): The platform name; e.g., 'win', 'mac', 'linux', 'android',
      'ios', etc.
    stacktrace (Stacktrace): The stacktrace of the crash. N.B., this is
      an object generated by parsing the string containing the stack trace;
      we do not store the string itself.
    regression_range : a pair of the last-good and first-bad
      versions. N.B., because this is an input, it is up to clients
      to call DetectRegressionRange (or whatever else) in order to
      provide this information.
  """
  __slots__ = ()

  def __new__(cls, crashed_version, signature, platform, stacktrace,
      regression_range):
    assert isinstance(stacktrace, Stacktrace), TypeError(
        'In the fourth argument to CrashReport constructor, '
        'expected Stacktrace object, but got %s object instead.'
        % stacktrace.__class__.__name__)

    return super(cls, CrashReport).__new__(cls,
        crashed_version, signature, platform, stacktrace, regression_range)
