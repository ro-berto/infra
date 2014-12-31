# Copyright 2014 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# TODO(nodir): write tests for ACL

# pylint: disable=W0613

from components import auth


def can_add_build_to_namespace(namespace, identity):  # pragma: no cover
  return auth.is_admin(identity)

def can_peek_namespace(namespace, identity):  # pragma: no cover
  return auth.is_admin(identity)

def can_lease_build(build, identity):  # pragma: no cover
  return auth.is_admin(identity)

def can_cancel_build(build, identity):  # pragma: no cover
  return auth.is_admin(identity)

def can_view_build(build, identity):  # pragma: no cover
  return auth.is_admin(identity)
