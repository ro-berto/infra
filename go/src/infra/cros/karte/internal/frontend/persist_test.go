// Copyright 2022 The ChromiumOS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/appengine/gaetesting"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/gae/service/datastore"

	kartepb "infra/cros/karte/api"
	"infra/cros/karte/internal/idstrategy"
	"infra/cros/karte/internal/scalars"
)

// TestPersistObservations tests persisting observations.
func TestPersistObservations(t *testing.T) {
	t.Parallel()

	const kind = "c98f39d2-592b-4700-b6ee-874ce8f6edc2"
	const metricKind = "abf5fa64-69e5-4983-83be-0366c3d4a4f8"

	Convey("test persisting observation", t, func() {
		ctx := gaetesting.TestingContext()
		ctx = idstrategy.Use(ctx, idstrategy.NewNaive())
		testClock := testclock.New(time.Unix(10, 0))
		ctx = clock.Set(ctx, testClock)
		datastore.GetTestable(ctx).Consistent(true)
		k := NewKarteFrontend().(*karteFrontend)
		fake := &fakeClient{}
		a, err := k.CreateAction(ctx, &kartepb.CreateActionRequest{
			Action: &kartepb.Action{
				Kind:      kind,
				StartTime: scalars.ConvertTimeToTimestampPtr(time.Unix(1, 0)),
				StopTime:  scalars.ConvertTimeToTimestampPtr(time.Unix(1, 0)),
				SealTime:  scalars.ConvertTimeToTimestampPtr(time.Unix(1, 0)),
			},
		})
		So(err, ShouldBeNil)
		So(a.Kind, ShouldEqual, kind)
		So(a.SealTime, ShouldResemble, scalars.ConvertTimeToTimestampPtr(time.Unix(1, 0)))
		o, err := k.CreateObservation(ctx, &kartepb.CreateObservationRequest{
			Observation: &kartepb.Observation{
				ActionName: a.Name,
				MetricKind: metricKind,
			},
		})
		So(err, ShouldBeNil)
		So(o.MetricKind, ShouldEqual, metricKind)
		So(o.ActionName, ShouldEqual, a.Name)
		_, err = k.persistActionRangeImpl(ctx, fake, &kartepb.PersistActionRangeRequest{
			StartTime: scalars.ConvertTimeToTimestampPtr(time.Unix(0, 0)),
			StopTime:  scalars.ConvertTimeToTimestampPtr(time.Unix(100, 0)),
		})
		So(err, ShouldBeNil)
		So(fake.size(), ShouldEqual, 2)
	})

	Convey("test persisting multiple observations associated with single observation", t, func() {
		ctx := gaetesting.TestingContext()
		ctx = idstrategy.Use(ctx, idstrategy.NewNaive())
		testClock := testclock.New(time.Unix(10, 0))
		ctx = clock.Set(ctx, testClock)
		datastore.GetTestable(ctx).Consistent(true)
		k := NewKarteFrontend().(*karteFrontend)
		const times = 10
		fake := &fakeClient{}
		a, err := k.CreateAction(ctx, &kartepb.CreateActionRequest{
			Action: &kartepb.Action{
				Kind:      kind,
				StartTime: scalars.ConvertTimeToTimestampPtr(time.Unix(1, 0)),
				StopTime:  scalars.ConvertTimeToTimestampPtr(time.Unix(1, 0)),
				SealTime:  scalars.ConvertTimeToTimestampPtr(time.Unix(1, 0)),
			},
		})
		So(err, ShouldBeNil)
		So(a.Kind, ShouldEqual, kind)
		So(a.SealTime, ShouldResemble, scalars.ConvertTimeToTimestampPtr(time.Unix(1, 0)))
		for i := 0; i < times; i++ {
			o, err := k.CreateObservation(ctx, &kartepb.CreateObservationRequest{
				Observation: &kartepb.Observation{
					ActionName: a.Name,
					MetricKind: metricKind,
				},
			})
			So(err, ShouldBeNil)
			So(o.MetricKind, ShouldEqual, metricKind)
			So(o.ActionName, ShouldEqual, a.Name)
		}
		_, err = k.persistActionRangeImpl(ctx, fake, &kartepb.PersistActionRangeRequest{
			StartTime: scalars.ConvertTimeToTimestampPtr(time.Unix(0, 0)),
			StopTime:  scalars.ConvertTimeToTimestampPtr(time.Unix(100, 0)),
		})
		So(err, ShouldBeNil)
		So(fake.size(), ShouldEqual, 1+times)
	})
}