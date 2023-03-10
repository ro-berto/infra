# Copyright 2016 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Tests for notify.py."""
from __future__ import print_function
from __future__ import division
from __future__ import absolute_import

import json
import mock
import unittest
import flask

from google.appengine.ext import testbed

from features import notify
from features import notify_reasons
from framework import emailfmt
from framework import urls
from proto import tracker_pb2
from services import service_manager
from testing import fake
from testing import testing_helpers
from tracker import attachment_helpers
from tracker import tracker_bizobj


def MakeTestIssue(project_id, local_id, owner_id, reporter_id, is_spam=False):
  issue = tracker_pb2.Issue()
  issue.project_id = project_id
  issue.local_id = local_id
  issue.issue_id = 1000 * project_id + local_id
  issue.owner_id = owner_id
  issue.reporter_id = reporter_id
  issue.is_spam = is_spam
  return issue


class NotifyTaskHandleRequestTest(unittest.TestCase):

  def setUp(self):
    self.services = service_manager.Services(
        user=fake.UserService(),
        usergroup=fake.UserGroupService(),
        project=fake.ProjectService(),
        config=fake.ConfigService(),
        issue=fake.IssueService(),
        issue_star=fake.IssueStarService(),
        features=fake.FeaturesService())
    self.requester = self.services.user.TestAddUser('requester@example.com', 1)
    self.nonmember = self.services.user.TestAddUser('user@example.com', 2)
    self.member = self.services.user.TestAddUser('member@example.com', 3)
    self.project = self.services.project.TestAddProject(
        'test-project', owner_ids=[1, 3], project_id=12345)
    self.issue1 = MakeTestIssue(
        project_id=12345, local_id=1, owner_id=2, reporter_id=1)
    self.issue2 = MakeTestIssue(
        project_id=12345, local_id=2, owner_id=2, reporter_id=1)
    self.services.issue.TestAddIssue(self.issue1)

    self.orig_sign_attachment_id = attachment_helpers.SignAttachmentID
    attachment_helpers.SignAttachmentID = (
        lambda aid: 'signed_%d' % aid)
    self.servlet = notify.OutboundEmailTask(services=self.services)
    self.app = flask.Flask('test_app')
    self.app.config['TESTING'] = True
    self.app.add_url_rule(
        '/_task/outboundEmail.do',
        view_func=self.servlet.PostOutboundEmailTask,
        methods=['POST'])
    self.testbed = testbed.Testbed()
    self.testbed.activate()
    self.testbed.init_memcache_stub()
    self.testbed.init_datastore_v3_stub()

  def tearDown(self):
    attachment_helpers.SignAttachmentID = self.orig_sign_attachment_id

  def get_filtered_task_call_args(self, create_task_mock, relative_uri):
    return [
        (args, kwargs)
        for (args, kwargs) in create_task_mock.call_args_list
        if args[0]['app_engine_http_request']['relative_uri'] == relative_uri
    ]

  def VerifyParams(self, result, params):
    self.assertEqual(
        bool(params['send_email']), result['params']['send_email'])
    if 'issue_id' in params:
      self.assertEqual(params['issue_id'], result['params']['issue_id'])
    if 'issue_ids' in params:
      self.assertEqual([int(p) for p in params['issue_ids'].split(',')],
                       result['params']['issue_ids'])

  def testNotifyIssueChangeTask_Normal(self):
    task = notify.NotifyIssueChangeTask(services=self.services)
    params = {'send_email': 1, 'issue_id': 12345001, 'seq': 0,
              'commenter_id': 2}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.VerifyParams(result, params)

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyIssueChangeTask_Spam(self, _create_task_mock):
    issue = MakeTestIssue(
        project_id=12345, local_id=1, owner_id=1, reporter_id=1,
        is_spam=True)
    self.services.issue.TestAddIssue(issue)
    task = notify.NotifyIssueChangeTask(services=self.services)
    params = {'send_email': 0, 'issue_id': issue.issue_id, 'seq': 0,
              'commenter_id': 2}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.assertEqual(0, len(result['notified']))

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyBlockingChangeTask_Normal(self, _create_task_mock):
    issue2 = MakeTestIssue(
        project_id=12345, local_id=2, owner_id=2, reporter_id=1)
    self.services.issue.TestAddIssue(issue2)
    task = notify.NotifyBlockingChangeTask(services=self.services)
    params = {
        'send_email': 1, 'issue_id': issue2.issue_id, 'seq': 0,
        'delta_blocker_iids': str(self.issue1.issue_id), 'commenter_id': 1,
        'hostport': 'bugs.chromium.org'}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.VerifyParams(result, params)

  def testNotifyBlockingChangeTask_Spam(self):
    issue2 = MakeTestIssue(
        project_id=12345, local_id=2, owner_id=2, reporter_id=1,
        is_spam=True)
    self.services.issue.TestAddIssue(issue2)
    task = notify.NotifyBlockingChangeTask(services=self.services)
    params = {
        'send_email': 1, 'issue_id': issue2.issue_id, 'seq': 0,
        'delta_blocker_iids': str(self.issue1.issue_id), 'commenter_id': 1}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.assertEqual(0, len(result['notified']))

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyBulkChangeTask_Normal(self, create_task_mock):
    """We generate email tasks for each user involved in the issues."""
    issue2 = MakeTestIssue(
        project_id=12345, local_id=2, owner_id=2, reporter_id=1)
    issue2.cc_ids = [3]
    self.services.issue.TestAddIssue(issue2)
    task = notify.NotifyBulkChangeTask(services=self.services)
    params = {
        'send_email': 1, 'seq': 0,
        'issue_ids': '%d,%d' % (self.issue1.issue_id, issue2.issue_id),
        'old_owner_ids': '1,1', 'commenter_id': 1}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.VerifyParams(result, params)

    call_args_list = self.get_filtered_task_call_args(
        create_task_mock, urls.OUTBOUND_EMAIL_TASK + '.do')
    self.assertEqual(2, len(call_args_list))

    for (args, _kwargs) in call_args_list:
      task = args[0]
      body = json.loads(task['app_engine_http_request']['body'].decode())
      if 'user' in body['to']:
        self.assertIn(u'\u2026', body['from_addr'])
      # Full email for members
      if 'member' in body['to']:
        self.assertNotIn(u'\u2026', body['from_addr'])

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyBulkChangeTask_AlsoNotify(self, create_task_mock):
    """We generate email tasks for also-notify addresses."""
    self.issue1.derived_notify_addrs = [
        'mailing-list@example.com', 'member@example.com']
    task = notify.NotifyBulkChangeTask(services=self.services)
    params = {
        'send_email': 1, 'seq': 0,
        'issue_ids': '%d' % (self.issue1.issue_id),
        'old_owner_ids': '1', 'commenter_id': 1}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.VerifyParams(result, params)

    call_args_list = self.get_filtered_task_call_args(
        create_task_mock, urls.OUTBOUND_EMAIL_TASK + '.do')
    self.assertEqual(3, len(call_args_list))

    self.assertItemsEqual(
        ['user@example.com', 'mailing-list@example.com', 'member@example.com'],
        result['notified'])
    for (args, _kwargs) in call_args_list:
      task = args[0]
      body = json.loads(task['app_engine_http_request']['body'].decode())
      # obfuscated email for non-members
      if 'user' in body['to']:
        self.assertIn(u'\u2026', body['from_addr'])
      # Full email for members
      if 'member' in body['to']:
        self.assertNotIn(u'\u2026', body['from_addr'])

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyBulkChangeTask_ProjectNotify(self, create_task_mock):
    """We generate email tasks for project.issue_notify_address."""
    self.project.issue_notify_address = 'mailing-list@example.com'
    task = notify.NotifyBulkChangeTask(services=self.services)
    params = {
        'send_email': 1, 'seq': 0,
        'issue_ids': '%d' % (self.issue1.issue_id),
        'old_owner_ids': '1', 'commenter_id': 1}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.VerifyParams(result, params)

    call_args_list = self.get_filtered_task_call_args(
        create_task_mock, urls.OUTBOUND_EMAIL_TASK + '.do')
    self.assertEqual(2, len(call_args_list))

    self.assertItemsEqual(
        ['user@example.com', 'mailing-list@example.com'],
        result['notified'])

    for (args, _kwargs) in call_args_list:
      task = args[0]
      body = json.loads(task['app_engine_http_request']['body'].decode())
      # obfuscated email for non-members
      if 'user' in body['to']:
        self.assertIn(u'\u2026', body['from_addr'])
      # Full email for members
      if 'member' in body['to']:
        self.assertNotIn(u'\u2026', body['from_addr'])

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyBulkChangeTask_SubscriberGetsEmail(self, create_task_mock):
    """If a user subscription matches the issue, notify that user."""
    task = notify.NotifyBulkChangeTask(services=self.services)
    params = {
        'send_email': 1,
        'issue_ids': '%d' % (self.issue1.issue_id),
        'seq': 0,
        'old_owner_ids': '1', 'commenter_id': 1}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    self.services.user.TestAddUser('subscriber@example.com', 4)
    sq = tracker_bizobj.MakeSavedQuery(
        1, 'all open issues', 2, '', subscription_mode='immediate',
        executes_in_project_ids=[self.issue1.project_id])
    self.services.features.UpdateUserSavedQueries('cnxn', 4, [sq])
    result = task.HandleRequest(mr)
    self.VerifyParams(result, params)

    call_args_list = self.get_filtered_task_call_args(
        create_task_mock, urls.OUTBOUND_EMAIL_TASK + '.do')
    self.assertEqual(2, len(call_args_list))

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyBulkChangeTask_CCAndSubscriberListsIssueOnce(
      self, create_task_mock):
    """If a user both CCs and subscribes, include issue only once."""
    task = notify.NotifyBulkChangeTask(services=self.services)
    params = {
        'send_email': 1,
        'issue_ids': '%d' % (self.issue1.issue_id),
        'seq': 0,
        'old_owner_ids': '1', 'commenter_id': 1}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    self.services.user.TestAddUser('subscriber@example.com', 4)
    self.issue1.cc_ids = [4]
    sq = tracker_bizobj.MakeSavedQuery(
        1, 'all open issues', 2, '', subscription_mode='immediate',
        executes_in_project_ids=[self.issue1.project_id])
    self.services.features.UpdateUserSavedQueries('cnxn', 4, [sq])
    result = task.HandleRequest(mr)
    self.VerifyParams(result, params)

    call_args_list = self.get_filtered_task_call_args(
        create_task_mock, urls.OUTBOUND_EMAIL_TASK + '.do')
    self.assertEqual(2, len(call_args_list))

    found = False
    for (args, _kwargs) in call_args_list:
      task = args[0]
      body = json.loads(task['app_engine_http_request']['body'].decode())
      if body['to'] == 'subscriber@example.com':
        found = True
        task_body = body['body']
        self.assertEqual(1, task_body.count('Issue %d' % self.issue1.local_id))
    self.assertTrue(found)

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyBulkChangeTask_Spam(self, _create_task_mock):
    """A spam issue is excluded from notification emails."""
    issue2 = MakeTestIssue(
        project_id=12345, local_id=2, owner_id=2, reporter_id=1,
        is_spam=True)
    self.services.issue.TestAddIssue(issue2)
    task = notify.NotifyBulkChangeTask(services=self.services)
    params = {
        'send_email': 1,
        'issue_ids': '%d,%d' % (self.issue1.issue_id, issue2.issue_id),
        'seq': 0,
        'old_owner_ids': '1,1', 'commenter_id': 1}
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.assertEqual(1, len(result['notified']))

  def testFormatBulkIssues_Normal_Single(self):
    """A user may see full notification details for all changed issues."""
    self.issue1.summary = 'one summary'
    task = notify.NotifyBulkChangeTask(services=self.services)
    users_by_id = {}
    commenter_view = None
    config = self.services.config.GetProjectConfig('cnxn', 12345)
    addrperm = notify_reasons.AddrPerm(
        False, 'nonmember@example.com', self.nonmember,
        notify_reasons.REPLY_NOT_ALLOWED, None)

    subject, body = task._FormatBulkIssues(
      [self.issue1], users_by_id, commenter_view, 'localhost:8080',
      'test comment', [], config, addrperm)

    self.assertIn('one summary', subject)
    self.assertIn('one summary', body)
    self.assertIn('test comment', body)

  def testFormatBulkIssues_Normal_Multiple(self):
    """A user may see full notification details for all changed issues."""
    self.issue1.summary = 'one summary'
    self.issue2.summary = 'two summary'
    task = notify.NotifyBulkChangeTask(services=self.services)
    users_by_id = {}
    commenter_view = None
    config = self.services.config.GetProjectConfig('cnxn', 12345)
    addrperm = notify_reasons.AddrPerm(
        False, 'nonmember@example.com', self.nonmember,
        notify_reasons.REPLY_NOT_ALLOWED, None)

    subject, body = task._FormatBulkIssues(
      [self.issue1, self.issue2], users_by_id, commenter_view, 'localhost:8080',
      'test comment', [], config, addrperm)

    self.assertIn('2 issues changed', subject)
    self.assertIn('one summary', body)
    self.assertIn('two summary', body)
    self.assertIn('test comment', body)

  def testFormatBulkIssues_LinkOnly_Single(self):
    """A user may not see full notification details for some changed issue."""
    self.issue1.summary = 'one summary'
    self.issue1.labels = ['Restrict-View-Google']
    task = notify.NotifyBulkChangeTask(services=self.services)
    users_by_id = {}
    commenter_view = None
    config = self.services.config.GetProjectConfig('cnxn', 12345)
    addrperm = notify_reasons.AddrPerm(
        False, 'nonmember@example.com', self.nonmember,
        notify_reasons.REPLY_NOT_ALLOWED, None)

    subject, body = task._FormatBulkIssues(
      [self.issue1], users_by_id, commenter_view, 'localhost:8080',
      'test comment', [], config, addrperm)

    self.assertIn('issue 1', subject)
    self.assertNotIn('one summary', subject)
    self.assertNotIn('one summary', body)
    self.assertNotIn('test comment', body)

  def testFormatBulkIssues_LinkOnly_Multiple(self):
    """A user may not see full notification details for some changed issue."""
    self.issue1.summary = 'one summary'
    self.issue1.labels = ['Restrict-View-Google']
    self.issue2.summary = 'two summary'
    task = notify.NotifyBulkChangeTask(services=self.services)
    users_by_id = {}
    commenter_view = None
    config = self.services.config.GetProjectConfig('cnxn', 12345)
    addrperm = notify_reasons.AddrPerm(
        False, 'nonmember@example.com', self.nonmember,
        notify_reasons.REPLY_NOT_ALLOWED, None)

    subject, body = task._FormatBulkIssues(
      [self.issue1, self.issue2], users_by_id, commenter_view, 'localhost:8080',
      'test comment', [], config, addrperm)

    self.assertIn('2 issues', subject)
    self.assertNotIn('summary', subject)
    self.assertNotIn('one summary', body)
    self.assertIn('two summary', body)
    self.assertNotIn('test comment', body)

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyApprovalChangeTask_Normal(self, _create_task_mock):
    config = self.services.config.GetProjectConfig('cnxn', 12345)
    config.field_defs = [
        # issue's User field with any_comment is notified.
        tracker_bizobj.MakeFieldDef(
            121, 12345, 'TL', tracker_pb2.FieldTypes.USER_TYPE,
            '', '', False, False, False, None, None, None, False, '',
            None, tracker_pb2.NotifyTriggers.ANY_COMMENT, 'no_action',
            'TL, notified on everything', False),
        # issue's User field with never is not notified.
        tracker_bizobj.MakeFieldDef(
            122, 12345, 'silentTL', tracker_pb2.FieldTypes.USER_TYPE,
            '', '', False, False, False, None, None, None, False, '',
            None, tracker_pb2.NotifyTriggers.NEVER, 'no_action',
            'TL, notified on nothing', False),
        # approval's User field with any_comment is notified.
        tracker_bizobj.MakeFieldDef(
            123, 12345, 'otherapprovalTL', tracker_pb2.FieldTypes.USER_TYPE,
            '', '', False, False, False, None, None, None, False, '',
            None, tracker_pb2.NotifyTriggers.ANY_COMMENT, 'no_action',
            'TL on the approvers team', False, approval_id=3),
        # another approval's User field with any_comment is not notified.
        tracker_bizobj.MakeFieldDef(
            124, 12345, 'otherapprovalTL', tracker_pb2.FieldTypes.USER_TYPE,
            '', '', False, False, False, None, None, None, False, '',
            None, tracker_pb2.NotifyTriggers.ANY_COMMENT, 'no_action',
            'TL on another approvers team', False, approval_id=4),
        tracker_bizobj.MakeFieldDef(
            3, 12345, 'Goat-Approval', tracker_pb2.FieldTypes.APPROVAL_TYPE,
            '', '', False, False, False, None, None, None, False, '',
            None, tracker_pb2.NotifyTriggers.NEVER, 'no_action',
            'Get Approval from Goats', False)
    ]
    self.services.config.StoreConfig('cnxn', config)

    # Custom user_type field TLs
    self.services.user.TestAddUser('TL@example.com', 111)
    self.services.user.TestAddUser('silentTL@example.com', 222)
    self.services.user.TestAddUser('approvalTL@example.com', 333)
    self.services.user.TestAddUser('otherapprovalTL@example.com', 444)

    # Approvers
    self.services.user.TestAddUser('approver_old@example.com', 777)
    self.services.user.TestAddUser('approver_new@example.com', 888)
    self.services.user.TestAddUser('approver_still@example.com', 999)
    self.services.user.TestAddUser('approver_group@example.com', 666)
    self.services.user.TestAddUser('group_mem1@example.com', 661)
    self.services.user.TestAddUser('group_mem2@example.com', 662)
    self.services.user.TestAddUser('group_mem3@example.com', 663)
    self.services.usergroup.TestAddGroupSettings(
        666, 'approver_group@example.com')
    self.services.usergroup.TestAddMembers(666, [661, 662, 663])
    canary_phase = tracker_pb2.Phase(
        name='Canary', phase_id=1, rank=1)
    approval_values = [
        tracker_pb2.ApprovalValue(approval_id=3,
                                  approver_ids=[888, 999, 666, 661])]
    approval_issue = MakeTestIssue(
        project_id=12345, local_id=2, owner_id=2, reporter_id=1,
        is_spam=True)
    approval_issue.phases = [canary_phase]
    approval_issue.approval_values = approval_values
    approval_issue.field_values = [
        tracker_bizobj.MakeFieldValue(121, None, None, 111, None, None, False),
        tracker_bizobj.MakeFieldValue(122, None, None, 222, None, None, False),
        tracker_bizobj.MakeFieldValue(123, None, None, 333, None, None, False),
        tracker_bizobj.MakeFieldValue(124, None, None, 444, None, None, False),
    ]
    self.services.issue.TestAddIssue(approval_issue)

    amend = tracker_bizobj.MakeApprovalApproversAmendment([888], [777])

    comment = tracker_pb2.IssueComment(
        project_id=12345, user_id=999, issue_id=approval_issue.issue_id,
        amendments=[amend], timestamp=1234567890, content='just a comment.')
    attach = tracker_pb2.Attachment(
        attachment_id=4567, filename='sploot.jpg', mimetype='image/png',
        gcs_object_id='/pid/attachments/abcd', filesize=(1024 * 1023))
    comment.attachments.append(attach)
    self.services.issue.TestAddComment(comment, approval_issue.local_id)
    self.services.issue.TestAddAttachment(
        attach, comment.id, approval_issue.issue_id)

    task = notify.NotifyApprovalChangeTask(services=self.services)
    params = {
        'send_email': 1,
        'issue_id': approval_issue.issue_id,
        'approval_id': 3,
        'comment_id': comment.id,
    }
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.assertTrue('just a comment' in result['tasks'][0]['body'])
    self.assertTrue('Approvers: -appro...' in result['tasks'][0]['body'])
    self.assertTrue('sploot.jpg' in result['tasks'][0]['body'])
    self.assertTrue(
        '/issues/attachment?aid=4567' in result['tasks'][0]['body'])
    self.assertItemsEqual(
        ['user@example.com', 'approver_old@example.com',
         'approver_new@example.com', 'TL@example.com',
         'approvalTL@example.com', 'group_mem1@example.com',
         'group_mem2@example.com', 'group_mem3@example.com'],
        result['notified'])

    # Test no approvers/groups notified
    # Status change to NEED_INFO does not email approvers.
    amend2 = tracker_bizobj.MakeApprovalStatusAmendment(
        tracker_pb2.ApprovalStatus.NEED_INFO)
    comment2 = tracker_pb2.IssueComment(
        project_id=12345, user_id=999, issue_id=approval_issue.issue_id,
        amendments=[amend2], timestamp=1234567891, content='')
    self.services.issue.TestAddComment(comment2, approval_issue.local_id)
    task = notify.NotifyApprovalChangeTask(services=self.services)
    params = {
        'send_email': 1,
        'issue_id': approval_issue.issue_id,
        'approval_id': 3,
        'comment_id': comment2.id,
    }
    mr = testing_helpers.MakeMonorailRequest(
        user_info={'user_id': 1},
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)

    self.assertIsNotNone(result['tasks'][0].get('references'))
    self.assertEqual(result['tasks'][0]['reply_to'], emailfmt.NoReplyAddress())
    self.assertTrue('Status: need_info' in result['tasks'][0]['body'])
    self.assertItemsEqual(
        ['user@example.com', 'TL@example.com', 'approvalTL@example.com'],
        result['notified'])

  def testNotifyApprovalChangeTask_GetApprovalEmailRecipients(self):
    task = notify.NotifyApprovalChangeTask(services=self.services)
    issue = fake.MakeTestIssue(789, 1, 'summary', 'New', 111)
    approval_value = tracker_pb2.ApprovalValue(
        approver_ids=[222, 333],
        status=tracker_pb2.ApprovalStatus.APPROVED)
    comment = tracker_pb2.IssueComment(
        project_id=789, user_id=1, issue_id=78901)

    # Comment with not amendments notifies everyone.
    rids = task._GetApprovalEmailRecipients(
        approval_value, comment, issue, [777, 888])
    self.assertItemsEqual(rids, [111, 222, 333, 777, 888])

    # New APPROVED status notifies owners and any_comment users.
    amendment = tracker_bizobj.MakeApprovalStatusAmendment(
        tracker_pb2.ApprovalStatus.APPROVED)
    comment.amendments = [amendment]
    rids = task._GetApprovalEmailRecipients(
        approval_value, comment, issue, [777, 888])
    self.assertItemsEqual(rids, [111, 777, 888])

    # New REVIEW_REQUESTED status notifies approvers.
    approval_value.status = tracker_pb2.ApprovalStatus.REVIEW_REQUESTED
    amendment = tracker_bizobj.MakeApprovalStatusAmendment(
        tracker_pb2.ApprovalStatus.REVIEW_REQUESTED)
    comment.amendments = [amendment]
    rids = task._GetApprovalEmailRecipients(
        approval_value, comment, issue, [777, 888])
    self.assertItemsEqual(rids, [222, 333])

    # Approvers change notifies everyone.
    amendment = tracker_bizobj.MakeApprovalApproversAmendment(
        [222], [555])
    comment.amendments = [amendment]
    approval_value.approver_ids = [222]
    rids = task._GetApprovalEmailRecipients(
        approval_value, comment, issue, [777], omit_ids=[444, 333])
    self.assertItemsEqual(rids, [111, 222, 555, 777])

  @mock.patch('framework.cloud_tasks_helpers.create_task')
  def testNotifyRulesDeletedTask(self, _create_task_mock):
    self.services.project.TestAddProject(
        'proj', owner_ids=[777, 888], project_id=789)
    self.services.user.TestAddUser('owner1@test.com', 777)
    self.services.user.TestAddUser('cow@test.com', 888)
    task = notify.NotifyRulesDeletedTask(services=self.services)
    params = {'project_id': 789,
              'filter_rules': 'if green make yellow,if orange make blue'}
    mr = testing_helpers.MakeMonorailRequest(
        params=params,
        method='POST',
        services=self.services)
    result = task.HandleRequest(mr)
    self.assertEqual(len(result['tasks']), 2)
    body = result['tasks'][0]['body']
    self.assertTrue('if green make yellow' in body)
    self.assertTrue('if green make yellow' in body)
    self.assertTrue('/p/proj/adminRules' in body)
    self.assertItemsEqual(
        ['cow@test.com', 'owner1@test.com'], result['notified'])

  def testOutboundEmailTask_Normal(self):
    """We can send an email."""
    params = {
        'from_addr': 'requester@example.com',
        'reply_to': 'user@example.com',
        'to': 'user@example.com',
        'subject': 'Test subject'}
    data = json.dumps(params)
    res = self.app.test_client().post('/_task/outboundEmail.do', data=data)
    res_string = res.get_data()[5:]
    res_json = json.loads(res_string)
    self.assertEqual(params['from_addr'], res_json['sender'])
    self.assertEqual(params['subject'], res_json['subject'])

  def testOutboundEmailTask_MissingTo(self):
    """We skip emails that don't specify the To-line."""
    params = {
        'from_addr': 'requester@example.com',
        'reply_to': 'user@example.com',
        'subject': 'Test subject'}
    data = json.dumps(params)
    res = self.app.test_client().post('/_task/outboundEmail.do', data=data)
    res_string = res.get_data()[5:]
    res_json = json.loads(res_string)
    self.assertEqual(
        'Skipping because no "to" address found.', res_json['note'])
    self.assertNotIn('from_addr', res_string)

  def testOutboundEmailTask_BannedUser(self):
    """We don't send emails to banned users.."""
    self.servlet.services.user.TestAddUser(
        'banned@example.com', 404, banned=True)
    params = {
        'from_addr': 'requester@example.com',
        'reply_to': 'user@example.com',
        'to': 'banned@example.com',
        'subject': 'Test subject'}
    data = json.dumps(params)
    res = self.app.test_client().post('/_task/outboundEmail.do', data=data)
    res_string = res.get_data()[5:]
    res_json = json.loads(res_string)
    self.assertEqual('Skipping because user is banned.', res_json['note'])
    self.assertNotIn('from_addr', res_string)
