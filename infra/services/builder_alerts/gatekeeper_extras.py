# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# FIXME: Everything in this file belongs in gatekeeper_ng_config.py

import logging

from infra.services.builder_alerts.buildbot import master_name_from_url

def is_excluded_builder(master_config, builder):
  cfg = master_config[0]
  if builder in cfg:
    cfg = cfg[builder]
  else:
    cfg = cfg.get('*', {})
  return builder in cfg.get('excluded_builders', set())

def builder_is_excluded(builder, config, master_config):
  if is_excluded_builder(config, builder):
    return True

  if '*' in master_config:
    return False

  return builder not in master_config

def trees_for_master(master_url, gatekeeper_trees_config):
  """Get the name of the tree for a given master url, or the master's name."""
  trees = []
  for tree_name, tree_config in gatekeeper_trees_config.iteritems():
    if master_url in tree_config['masters']:
      trees.append(tree_name)

  if trees:
    return trees

  return [master_name_from_url(master_url)]


def apply_gatekeeper_rules(alerts, gatekeeper, gatekeeper_trees):
  filtered_alerts = []
  for alert in alerts:
    master_url = alert['master_url']
    config = gatekeeper.get(master_url)
    if not config:
      # Unclear if this should be set or not?
      # alert['would_close_tree'] = False
      filtered_alerts.append(alert)
      continue

    builder = alert.get('builder_name')
    alert_trees = trees_for_master(master_url, gatekeeper_trees)
    for tree in alert_trees:
      alert_cpy = alert.copy()
      alert_cpy['tree'] = tree
      if builder:
        masters = gatekeeper_trees.get(tree, {}).get('masters', {})
        if builder_is_excluded(builder, config, masters.get(master_url, [])):
          continue
        # Only apply tree closer logic for step failures
        if 'step_name' in alert_cpy:
          alert_cpy['would_close_tree'] = would_close_tree(
              config, builder, alert_cpy['step_name'])

      filtered_alerts.append(alert_cpy)
  return filtered_alerts


def fetch_master_urls(gatekeeper, args):
  # Currently using gatekeeper.json, but could use:
  # https://chrome-infra-stats.appspot.com/_ah/api#p/stats/v1/stats.masters.list
  master_urls = gatekeeper.keys()
  if args.master_filter:
    master_urls = [url for url in master_urls if args.master_filter not in url]
  return master_urls


def would_close_tree(master_config, builder_name, step_name):
  # FIXME: Section support should be removed:
  master_config = master_config[0]
  builder_config = master_config.get(builder_name, {})
  if not builder_config:
    builder_config = master_config.get('*', {})

  # close_tree is currently unused in gatekeeper.json but planned to be.
  close_tree = builder_config.get('close_tree', True)
  if not close_tree:
    logging.debug('close_tree is false')
    return False

  # Excluded steps never close.
  excluded_steps = set(builder_config.get('excluded_steps', []))
  if step_name in excluded_steps:
    logging.debug('%s is an excluded_step', step_name)
    return False

  # See gatekeeper_ng_config.py for documentation of
  # the config format.
  # forgiving/closing controls if mails are sent on close.
  # steps/optional controls if step-absence indicates failure.
  # this function assumes the step is present and failing
  # and thus doesn't care between these 4 types:
  closing_steps = (builder_config.get('forgiving_steps', set()) |
                   builder_config.get('forgiving_optional', set()) |
                   builder_config.get('closing_steps', set()) |
                   builder_config.get('closing_optional', set()))

  # A '*' in any of the above types means it applies to all steps.
  if '*' in closing_steps:
    return True

  if step_name in closing_steps:
    return True

  logging.debug('%s not in closing_steps: %s', step_name, closing_steps)
  return False
