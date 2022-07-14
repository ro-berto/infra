// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"context"
	"fmt"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"infra/appengine/crosskylabadmin/internal/app/frontend/routing"
	"infra/libs/skylab/common/heuristics"
)

// RouteTask routes a task for a given bot.
//
// The possible return values are:
// - "legacy"  (for legacy, which is the default)
// -       ""  (indicates an error, should be treated as equivalent to "legacy" by callers)
// -  "paris"  (for PARIS, which is new)
// -  "latest" (indicates the latest version of paris)
//
func RouteTask(ctx context.Context, taskType string, botID string, expectedState string, pools []string, randFloat float64) (heuristics.TaskType, error) {
	if taskType == "" {
		return heuristics.LegacyTaskType, errors.New("route task: task type cannot be empty")
	}
	switch taskType {
	case "repair":
		return routeRepairTask(ctx, botID, expectedState, pools, randFloat)
	case "audit_rpm":
		return routeAuditRPMTask()
	}
	return heuristics.LegacyTaskType, fmt.Errorf("route task: unrecognized task name %q", taskType)
}

// routeAuditRPMTask routes an audit RPM task to a specific implementation: legacy, paris, or latest.
func routeAuditRPMTask() (heuristics.TaskType, error) {
	return heuristics.LegacyTaskType, errors.New("route audit rpm task: not yet implemented")
}

// routeRepairTask routes a repair task for a given bot.
//
// The possible return values are:
// - "legacy"  (for legacy, which is the default)
// -       ""  (indicates an error, should be treated as equivalent to "legacy" by callers)
// -  "paris"  (for PARIS, which is new)
// -  "latest" (latest version of paris)
//
// routeRepairTask takes as an argument randFloat (which is a float64 in the closed interval [0, 1]).
// This argument is, by design, all the entropy that randFloat will need. Taking this as an argument allows
// routeRepairTask itself to be deterministic because the caller is responsible for generating the random
// value.
func routeRepairTask(ctx context.Context, botID string, expectedState string, pools []string, randFloat float64) (heuristics.TaskType, error) {
	if !(0.0 <= randFloat && randFloat <= 1.0) {
		return heuristics.LegacyTaskType, fmt.Errorf("Route repair task: randfloat %f is not in [0, 1]", randFloat)
	}
	isLabstation := heuristics.LooksLikeLabstation(botID)
	rolloutConfig, err := getRolloutConfig(ctx, "repair", isLabstation, expectedState)
	if err != nil {
		return heuristics.LegacyTaskType, errors.Annotate(err, "route repair task").Err()
	}
	out, r := routeRepairTaskImpl(
		ctx,
		rolloutConfig,
		&dutRoutingInfo{
			hostname:   heuristics.NormalizeBotNameToDeviceName(botID),
			labstation: isLabstation,
			pools:      pools,
		},
		randFloat,
	)
	reason, ok := routing.ReasonMessageMap[r]
	if !ok {
		logging.Infof(ctx, "Unrecognized reason %d", int64(r))
	}
	logging.Infof(ctx, "Sending device repair to %q because %q", out, reason)
	return out, nil
}
