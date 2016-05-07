# Copyright 2016 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from crash.callstack import StackFrame
from crash.results import MatchResult
from crash.scorers.aggregator import Aggregator
from crash.scorers.min_distance import MinDistance
from crash.scorers.test.scorer_test_suite import ScorerTestSuite
from crash.scorers.top_frame_index import TopFrameIndex


class AggregatorTest(ScorerTestSuite):

  def testScoreAndReason(self):
    result = MatchResult(self._GetDummyChangeLog(), 'src/', '')
    result.file_to_stack_infos = {
        'a.cc': [(StackFrame(0, 'src/', '', 'func', 'a.cc', [7]), 0)]
    }
    result.min_distance = 0

    aggregator = Aggregator([TopFrameIndex(), MinDistance()])
    aggregator.ScoreAndReason(result)

    self.assertEqual(result.confidence, 1)
    self.assertEqual(result.reason,
                     ('1. Top frame changed is frame #0 (score: 1)\n'
                      '2. Minimum distance to crashed line is 0 (score: 1)\n'
                      '\nChanged file a.cc crashed in func (#0)'))
