# Copyright 2017 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""A class to handle the initial warmup request from AppEngine."""
from __future__ import print_function
from __future__ import division
from __future__ import absolute_import

import logging


def Warmup():
  """Placeholder for warmup work.  Used only to enable min_idle_instances."""
  # Don't do anything that could cause a jam when many instances start.
  logging.info('/_ah/startup does nothing in Monorail.')
  logging.info('However it is needed for min_idle_instances in app.yaml.')
  return ''


def Start():
  """Placeholder for start work.  Used only to enable manual_scaling."""
  # Don't do anything that could cause a jam when many instances start.
  logging.info('/_ah/start does nothing in Monorail.')
  logging.info('However it is needed for manual_scaling in app.yaml.')
  return ''


def Stop():
  """Placeholder for stop work.  Used only to enable manual_scaling."""
  # Don't do anything that could cause a jam when many instances start."""
  logging.info('/_ah/stop does nothing in Monorail.')
  logging.info('However it is needed for manual_scaling in app.yaml.')
  return ''
