# Copyright 2016 The Chromium Authors
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Helper functions and classes used in issue search and sorting."""
from __future__ import print_function
from __future__ import division
from __future__ import absolute_import

import logging
import re

from features import savedqueries_helpers
from search import query2ast
from services import tracker_fulltext
from services import fulltext_helpers
from tracker import tracker_helpers


# Users can use "is:starred" in queries to limit
# search results to issues starred by that user.
IS_STARRED_RE = re.compile(r'\b(?![-@.:])is:starred\b(?![-@.:])', re.I)

# Users can use "me" in other fields to refer to the logged in user name.
KEYWORD_ME_RE = re.compile(r'\b[-_a-z0-9]+[=:]me\b(?![-@.:=])', re.I)
ME_RE = re.compile(r'(?<=[=:])me\b(?![-@.:=])', re.I)


def _AccumulateIssueProjectsAndConfigs(
    cnxn, project_dict, config_dict, services, issues):
  """Fetch any projects and configs that we need but haven't already loaded.

  Args:
    cnxn: connection to SQL database.
    project_dict: dict {project_id: project} of projects that we have
        already retrieved.
    config_dict: dict {project_id: project} of configs that we have
        already retrieved.
    services: connections to backends.
    issues: list of issues, which may be parts of different projects.

  Returns:
    Nothing, but projects_dict will be updated to include all the projects that
    contain the given issues, and config_dicts will be updated to incude all
    the corresponding configs.
  """
  new_ids = {issue.project_id for issue in issues}
  new_ids.difference_update(iter(project_dict.keys()))
  new_projects_dict = services.project.GetProjects(cnxn, new_ids)
  project_dict.update(new_projects_dict)
  new_configs_dict = services.config.GetProjectConfigs(cnxn, new_ids)
  config_dict.update(new_configs_dict)


def ReplaceKeywordsWithUserIDs(me_user_ids, query):
  """Substitutes User ID in terms such as is:starred and me.

  This is done on the query string before it is parsed because the query string
  is used as a key for cached search results in memcache.  A search for by one
  user for owner:me should not retrieve results stored for some other user.

  Args:
    me_user_ids: [] when no user is logged in, or user ID of the logged in
        user when doing an interactive search, or the viewed user ID when
        viewing someone else's dashboard, or the subscribing user's ID when
        evaluating subscriptions.  Also contains linked account IDs.
    query: The query string.

  Returns:
    A pair (query, warnings) where query is a string with "me" and "is:starred"
    removed or replaced by new terms that use the numeric user ID provided,
    and warnings is a list of warning strings to display to the user.
  """
  warnings = []
  if me_user_ids:
    me_user_ids_str = ','.join(str(uid) for uid in me_user_ids)
    star_term = 'starredby:%s' % me_user_ids_str
    query = IS_STARRED_RE.sub(star_term, query)
    if KEYWORD_ME_RE.search(query):
      query = ME_RE.sub(me_user_ids_str, query)
  else:
    if IS_STARRED_RE.search(query):
      warnings.append('"is:starred" ignored because you are not signed in.')
      query = IS_STARRED_RE.sub('', query)
    if KEYWORD_ME_RE.search(query):
      warnings.append('"me" keyword ignored because you are not signed in.')
      query = KEYWORD_ME_RE.sub('', query)

  return query, warnings
