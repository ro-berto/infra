// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package testfinder find tests by using cros-test-finder.
package testfinder

import (
	"context"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/golang/protobuf/jsonpb"

	build_api "go.chromium.org/chromiumos/config/go/build/api"
	"go.chromium.org/chromiumos/config/go/test/api"
	"go.chromium.org/luci/common/errors"

	"infra/cros/cmd/cros-tool-runner/internal/common"
	"infra/cros/cmd/cros-tool-runner/internal/services"
)

const CrosTestFinderName = "cros-test-finder"

// Run find tests by using cros-test-finder.
func Run(ctx context.Context, req *api.CrosToolRunnerTestFinderRequest, crosTestFinderContainer *build_api.ContainerImageInfo, tokenFile string) (res *api.CrosToolRunnerTestFinderResponse, err error) {
	// Use host network for dev environment which DUT address is in the form localhost:<port>
	const (
		networkName = "host"
	)

	artifactDir, err := filepath.Abs(req.ArtifactDir)
	if err != nil {
		return nil, errors.Annotate(err, "prepare to run test finder: failed to resolve artifact directory %v", req.ArtifactDir).Err()
	}
	// All artifacts will be in <artifact_dir>/cros-test-finder.
	crosTestFinderDir := path.Join(artifactDir, CrosTestFinderName)
	// The input file name.
	inputFileName := path.Join(crosTestFinderDir, "request.json")

	// Setting up directories.
	if err := os.MkdirAll(crosTestFinderDir, 0755); err != nil {
		return nil, errors.Annotate(err, "prepare to run test finder: failed to create directory %s", crosTestFinderDir).Err()
	}
	log.Printf("Run test: created the %s directory %s", CrosTestFinderName, crosTestFinderDir)

	testReq := &api.CrosTestFinderRequest{
		TestSuites: req.GetTestSuites(),
	}
	if err := writeTestFinderInput(inputFileName, testReq); err != nil {
		return nil, errors.Annotate(err, "prepare to run test finder: failed to create input file %s", inputFileName).Err()
	}
	if err = services.RunTestFinderCLI(ctx, crosTestFinderContainer, networkName, crosTestFinderDir, tokenFile); err != nil {
		return nil, errors.Annotate(err, "run test finder: failed to run %s CLI", CrosTestFinderName).Err()
	}

	resultFileName := path.Join(crosTestFinderDir, "result.json")
	if _, err := os.Stat(resultFileName); os.IsNotExist(err) {
		return nil, errors.Reason("process test finder result: result not found").Err()
	}
	out, err := readTestFinderOutput(resultFileName)
	if err != nil {
		return nil, errors.Annotate(err, "process test finder result: failed to read test finder output").Err()
	}

	return &api.CrosToolRunnerTestFinderResponse{
		TestSuites: out.TestSuites,
	}, err
}

// writeTestFinderInput writes a CrosTestFinderRequest json.
func writeTestFinderInput(file string, req *api.CrosTestFinderRequest) error {
	f, err := os.Create(file)
	if err != nil {
		return errors.Annotate(err, "fail to create file %v", file).Err()
	}
	m := jsonpb.Marshaler{}
	if err := m.Marshal(f, req); err != nil {
		return errors.Annotate(err, "fail to marshal request to file %v", file).Err()
	}
	return nil
}

// readTestFinderOutput reads output file generated by cros-test-finder.
func readTestFinderOutput(filePath string) (*api.CrosTestFinderResponse, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Annotate(err, "read output").Err()
	}
	out := &api.CrosTestFinderResponse{}

	umrsh := common.JsonPbUnmarshaler()
	err = umrsh.Unmarshal(r, out)
	return out, errors.Annotate(err, "read output").Err()
}
