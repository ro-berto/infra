# Copyright 2016 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Classes for banning spammer users"""
from __future__ import print_function
from __future__ import division
from __future__ import absolute_import

import json
import time

from framework import cloud_tasks_helpers
from framework import framework_helpers
from framework import permissions
from framework import jsonfeed
from framework import servlet
from framework import urls


class BanSpammer(servlet.Servlet):
  """Ban a user and mark their content as spam"""

  def AssertBasePermission(self, mr):
    """Check whether the user has any permission to ban users.

    Args:
      mr: commonly used info parsed from the request.
    """
    super(BanSpammer, self).AssertBasePermission(mr)
    if not permissions.CanBan(mr, self.services):
      raise permissions.PermissionException(
          'User is not allowed to ban users.')

  def ProcessFormData(self, mr, post_data):
    self.AssertBasePermission(mr)
    viewed_user_id = mr.viewed_user_auth.user_pb.user_id
    reporter_id = mr.auth.user_id

    # First ban or un-ban the user as a spammer.
    framework_helpers.UserSettings.ProcessBanForm(
        mr.cnxn, self.services.user, post_data, mr.viewed_user_auth.user_id,
        mr.viewed_user_auth.user_pb)

    params = {
        'spammer_id': viewed_user_id,
        'reporter_id': reporter_id,
        'is_spammer': 'banned' in post_data
    }
    task = cloud_tasks_helpers.generate_simple_task(
        urls.BAN_SPAMMER_TASK + '.do', params)
    cloud_tasks_helpers.create_task(task)

    return framework_helpers.FormatAbsoluteURL(
        mr, mr.viewed_user_auth.user_view.profile_url, include_project=False,
        saved=1, ts=int(time.time()))

  def PostBanSpammerPage(self, **kwargs):
    return self.handler(**kwargs)


class BanSpammerTask(jsonfeed.InternalTask):
  """This task will update all of the comments and issues created by the
     target user with is_spam=True, and also add a manual verdict attached
     to the user who originated the ban request. This is a potentially long
     running operation, so it is implemented as an async task.
  """

  def HandleRequest(self, mr):
    spammer_id = mr.GetPositiveIntParam('spammer_id')
    reporter_id = mr.GetPositiveIntParam('reporter_id')
    is_spammer = mr.GetBoolParam('is_spammer')

    # Get all of the issues reported by the spammer.
    issue_ids = self.services.issue.GetIssueIDsReportedByUser(mr.cnxn,
        spammer_id)

    issues = []

    if len(issue_ids) > 0:
      issues = self.services.issue.GetIssues(
          mr.cnxn, issue_ids, use_cache=False)

      # Mark them as spam/ham in bulk.
      self.services.spam.RecordManualIssueVerdicts(mr.cnxn, self.services.issue,
          issues, reporter_id, is_spammer)

    # Get all of the comments
    comments = self.services.issue.GetCommentsByUser(mr.cnxn, spammer_id)

    for comment in comments:
      self.services.spam.RecordManualCommentVerdict(mr.cnxn,
            self.services.issue, self.services.user, comment.id,
            reporter_id, is_spammer)

    return json.dumps({
        'comments': len(comments),
        'issues': len(issues),
    })

  def PostBanSpammer(self, **kwargs):
    return self.handler(**kwargs)
