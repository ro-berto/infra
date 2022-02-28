// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package config

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
	"go.chromium.org/luci/gae/impl/memory"
)

func TestConfig(t *testing.T) {
	Convey("SetTestConfig updates context config", t, func() {
		sampleCfg, err := CreatePlaceholderConfig()
		So(err, ShouldBeNil)

		ctx := memory.Use(context.Background())
		SetTestConfig(ctx, sampleCfg)

		cfg, err := Get(ctx)
		So(err, ShouldBeNil)
		So(cfg, ShouldResembleProto, sampleCfg)
	})
}
