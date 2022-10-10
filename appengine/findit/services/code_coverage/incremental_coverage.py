# Copyright 2022 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import logging
from datetime import datetime
from datetime import timedelta

from google.appengine.ext import ndb

from libs import time_util
from model.code_coverage import PresubmitCoverageData
from model.test_location import TestLocation
from services import bigquery_helper
from services import test_tag_util

_PAGE_SIZE = 100

# Time period for which coverage report is to fetched and processed
_NUM_REPORT_DAYS = 30


def ExportIncrementalCoverage(run_id):
  """Exports incremental coverage metrics to Bigquery for last _NUM_REPORT_DAYS.

  Reads presubmit coverage data from Datastore, add few other dimensions to it
  and exports it to a Bigquery table.

  """
  query = PresubmitCoverageData.query(
      PresubmitCoverageData.cl_patchset.server_host == \
        'chromium-review.googlesource.com',
       PresubmitCoverageData.update_timestamp >= datetime.now() -
      timedelta(days=_NUM_REPORT_DAYS))
  total_rows = 0
  more = True
  cursor = None
  while more:
    results, cursor, more = query.fetch_page(_PAGE_SIZE, start_cursor=cursor)
    for result in results:
      bqrows = _CreateBigqueryRows(result, run_id)
      if bqrows:
        bigquery_helper.ReportRowsToBigquery(bqrows, 'findit-for-me',
                                             'code_coverage_summaries',
                                             'incremental_coverage')
        total_rows += 1
  logging.info('Total patchsets processed = %d', total_rows)


def _CreateBigqueryRows(presubmit_coverage, run_id):
  """Create a bigquery row for incremental coverage.

  Returns a list of dict whose keys are column names and values are column
  values corresponding to the schema of the bigquery table.

  Args:
    presubmit_coverage (PresubmitCoverageData): The PresubmitCoverageData
    fetched from Datastore
  """
  if not presubmit_coverage.incremental_percentages:
    return None
  coverage = []
  for inc_coverage in presubmit_coverage.incremental_percentages:
    coverage.append({
        'run_id': run_id,
        'total_lines': inc_coverage.total_lines,
        'covered_lines': inc_coverage.covered_lines,
        # ignore the leading double slash(//)
        'path': inc_coverage.path[2:],
        'cl_number': presubmit_coverage.cl_patchset.change,
        'cl_patchset': presubmit_coverage.cl_patchset.patchset,
        'server_host': presubmit_coverage.cl_patchset.server_host,
        'insert_timestamp': time_util.GetUTCNow().isoformat()
    })
  return coverage
