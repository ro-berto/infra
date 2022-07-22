// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package utils contains various helper functions.
package utils

import (
	"context"
	"fmt"
	"strings"

	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/realms"
	"google.golang.org/grpc/codes"
)

var (
	ErrInvalidRealm     = errors.New("realm must be in the format <project>:<realm>")
	ErrMultipleProjects = errors.New("all realms must be from the same projects")
)

// SplitRealm splits the realm into the LUCI project name and the (sub)realm.
// Returns ErrInvalidRealm if the provided realm doesn't have a valid format.
func SplitRealm(realm string) (proj string, subRealm string, err error) {
	parts := strings.SplitN(realm, ":", 2)
	if len(parts) != 2 {
		return "", "", ErrInvalidRealm
	}
	return parts[0], parts[1], nil
}

// SplitRealms splits the realms into the LUCI project name and the (sub)realms.
// All realms must belong to the same project.
//
// Returns ErrInvalidRealm if any of the realm doesn't have a valid format.
// Returns ErrMultipleProjects if not all realms are from the same project.
func SplitRealms(realms []string) (proj string, subRealms []string, err error) {
	if len(realms) == 0 {
		return "", nil, nil
	}

	subRealms = make([]string, 0, len(realms))
	proj, subRealm, err := SplitRealm(realms[0])
	if err != nil {
		return "", nil, ErrInvalidRealm
	}
	subRealms = append(subRealms, subRealm)
	for _, realm := range realms {
		currentProj, subRealm, err := SplitRealm(realm)
		if err != nil {
			return "", nil, ErrInvalidRealm
		}
		if currentProj != proj {
			return "", nil, ErrMultipleProjects
		}
		subRealms = append(subRealms, subRealm)

	}
	return proj, subRealms, nil
}

// HasPermissions is a wrapper around luci/server/auth.HasPermission that checks
// whether the user has all the listed permissions and return an appstatus
// annotated error if users have no permission.
func HasPermissions(ctx context.Context, permissions []realms.Permission, realm string, attrs realms.Attrs) error {
	for _, perm := range permissions {
		allowed, err := auth.HasPermission(ctx, perm, realm, attrs)
		if err != nil {
			return err
		}
		if !allowed {
			return appstatus.Errorf(codes.PermissionDenied, `caller does not have permission %s in realm %q`, perm, realm)
		}
	}
	return nil
}

// QueryRealms is a wrapper around luci/server/auth.QueryRealms that returns a
// list of realms where the current caller has all the listed permissions.
//
// If `project` is specified, only returns realms in the project.
// If both `project` and `subRealm` are specified, only returns the
// `<project>:<subRealm>` realm.
//
// Returns an appstatus annotated error if users have no permission in any
// matching realm.
//
// The permissions should be flagged in the process with UsedInQueryRealms
// flag, which lets the runtime know it must prepare indexes for the
// corresponding QueryRealms call.
func QueryRealms(ctx context.Context, permissions []realms.Permission, project, subRealm string, attrs realms.Attrs) ([]string, error) {
	if len(permissions) == 0 {
		return nil, errors.New("at least one permission must be provided")
	}

	if subRealm != "" {
		if project == "" {
			return nil, errors.New("project must be specified when the subRealm is specified")
		}
		realm := project + ":" + subRealm
		if err := HasPermissions(ctx, permissions, realm, nil); err != nil {
			return nil, err
		}
		return []string{realm}, nil
	}

	allowedRealms, err := auth.QueryRealms(ctx, permissions[0], project, attrs)
	if err != nil {
		return nil, err
	}
	allowedRealmSet := stringset.NewFromSlice(allowedRealms...)

	for _, perm := range permissions[1:] {
		allowedRealms, err := auth.QueryRealms(ctx, perm, project, attrs)
		if err != nil {
			return nil, err
		}
		allowedRealmSet = allowedRealmSet.Intersect(stringset.NewFromSlice(allowedRealms...))
	}

	if len(allowedRealmSet) == 0 {
		projectLabel := "any projects"
		if project != "" {
			projectLabel = fmt.Sprintf("project %q", project)
		}
		return nil, appstatus.Errorf(codes.PermissionDenied, `caller does not have permissions %v in %s`, permissions, projectLabel)
	}

	return allowedRealmSet.ToSortedSlice(), nil
}

// QuerySubRealms is similar to QueryRealms with the following differences:
//  1. project is required.
//  2. a list of subRealms is returned instead of a list of realms
//   (e.g. ["realm1", "realm2"] instead of ["project:realm1", "project:realm2"])
func QuerySubRealms(ctx context.Context, permissions []realms.Permission, project, subRealm string, attrs realms.Attrs) ([]string, error) {
	if project == "" {
		return nil, errors.New("project must be provided")
	}

	realms, err := QueryRealms(ctx, permissions, project, subRealm, nil)
	if err != nil {
		return nil, err
	}
	_, subRealms, err := SplitRealms(realms)
	if err != nil {
		// Realms from `QueryRealms` should always be valid. This should never happen.
		panic(err)
	}
	return subRealms, nil
}
