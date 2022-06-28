// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"context"
	"time"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	kartepb "infra/cros/karte/api"
	"infra/cros/karte/internal/errors"
	"infra/cros/karte/internal/idstrategy"
	"infra/cros/karte/internal/scalars"
)

// CreateAction creates an action, stores it in datastore, and then returns the just-created action.
func (k *karteFrontend) CreateAction(ctx context.Context, req *kartepb.CreateActionRequest) (*kartepb.Action, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "create action: request is nil")
	}
	if req.GetAction() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "create action: action is nil")
	}
	if req.GetAction().GetName() != "" {
		return nil, status.Errorf(codes.InvalidArgument, "create action: custom names not supported; names will be generated by the service")
	}

	// If no timestamp is provided default to the current time.
	// TODO(gregorynisbet): There are multiple kinds of timestamps floating around with different ideas about what "zero" is. Make this more consistent.
	//
	// TODO(gregorynisbet): Log the name of the action that created is responsible for this message.
	if req.GetAction().GetCreateTime() == nil {
		logging.Infof(ctx, "(msgid: 68361e64-46fb-4881-b4a3-d6b40e8ffd90) Applying default timestamp to request")
		req.GetAction().CreateTime = scalars.ConvertTimeToTimestampPtr(clock.Now(ctx))
	}

	// Add a seal time. Seal times prohibit modification after the seal time has passed.
	if req.GetAction().GetSealTime() == nil {
		// TODO(gregorynisbet): Move seal time to a configuration setting.
		// If we don't have a seal time, create a default one in the future.
		req.GetAction().SealTime = scalars.ConvertTimeToTimestampPtr(
			scalars.ConvertTimestampPtrToTime(req.GetAction().GetCreateTime()).Add(time.Duration(12) * time.Hour))
	}

	// Here we use the action create time given to us in the request proto instead of time.Now() so that
	// It is possible to backfill Karte with data from past tasks.
	// We don't trust these timestamps completely (after all, backfilled timestamps are lies), but the UUID suffix
	// should do a good job of guaranteeing uniqueness.
	// Additionally, Karte queries depend on the end_time of the event *as reported by the event*.
	// Events also have an a priori maximum duration,  which means that we can perform a semantically correct query based on the
	// end time using IDs whose lexicographic sort order takes the current timestamp into account.
	name, err := idstrategy.Get(ctx).IDForAction(ctx, req.GetAction())
	if err != nil {
		return nil, errors.Annotate(err, "create action").Err()
	}
	req.Action.Name = name

	logging.Infof(ctx, "Creating action of kind %q", req.GetAction().GetKind())
	actionEntity, err := convertActionToActionEntity(req.GetAction())
	if err != nil {
		logging.Errorf(ctx, "Error converting action: %s", err)
		return nil, errors.Annotate(err, "create action").Err()
	}
	if err := PutActionEntities(ctx, actionEntity); err != nil {
		logging.Errorf(ctx, "error writing action: %s", err)
		return nil, errors.Annotate(err, "writing action to datastore").Err()
	}
	return req.GetAction(), nil
}

// CreateObservation creates an observation and then returns the just-created observation.
func (k *karteFrontend) CreateObservation(ctx context.Context, req *kartepb.CreateObservationRequest) (*kartepb.Observation, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "create observation: request is nil")
	}
	if req.GetObservation() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "create observation: observation is nil")
	}
	if req.GetObservation().GetName() != "" {
		return nil, status.Errorf(codes.InvalidArgument, "create observation: custom names not supported; names will be generated by the service")
	}
	name, err := idstrategy.Get(ctx).IDForObservation(ctx, req.GetObservation())
	if err != nil {
		return nil, errors.Annotate(err, "create-action").Err()
	}
	req.Observation.Name = name

	observationEntity, err := convertObservationToObservationEntity(req.GetObservation())
	if err != nil {
		return nil, errors.Annotate(err, "create observation").Err()
	}
	if err := PutObservationEntities(ctx, observationEntity); err != nil {
		return nil, errors.Annotate(err, "writing action to datastore").Err()
	}
	return req.GetObservation(), nil
}