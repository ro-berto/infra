// Copyright 2023 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package data

import (
	lab_api "go.chromium.org/chromiumos/config/go/test/lab/api"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/skylab_test_runner"

	"infra/cros/cmd/cros_test_runner/internal/interfaces"
	"infra/cros/cmd/cros_test_runner/internal/tools/crostoolrunner"
)

// HwTestStateKeeper represents all the data hw test execution flow requires.
type HwTestStateKeeper struct {
	interfaces.StateKeeper

	// Set from input
	CftTestRequest *skylab_test_runner.CFTTestRequest

	// Dut related
	HostName    string
	DutTopology *lab_api.DutTopology

	// Tools and their related dependencies
	Ctr                   *crostoolrunner.CrosToolRunner
	DockerKeyFileLocation string
}
