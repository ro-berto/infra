// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/golang/protobuf/jsonpb"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/luciexe/build"
	"google.golang.org/grpc/metadata"

	steps "infra/cros/cmd/labpack/internal/steps"
	"infra/cros/cmd/labpack/logger"
	ufsUtil "infra/unifiedfleet/app/util"
)

// createLogger creates a logger for recovery lib.
func createLogger(ctx context.Context, logDir string, level logging.Level) (context.Context, logger.Logger, error) {
	const callDepth = 2
	newCtx, log, err := logger.NewLogger(ctx, callDepth, logDir, level, logger.DefaultFormat, true)
	return newCtx, log, errors.Annotate(err, "init logger").Err()
}

// printInputs prints input params.
func printInputs(ctx context.Context, input *steps.LabpackInput) (err error) {
	step, ctx := build.StartStep(ctx, "Input params")
	defer func() { step.End(err) }()
	req := step.Log("input proto")
	marsh := jsonpb.Marshaler{Indent: "  "}
	if err = marsh.Marshal(req, input); err != nil {
		return errors.Annotate(err, "failed to marshal proto").Err()
	}
	return nil
}

// describeEnvironment describes the environment where labpack is being run.
// TODO(gregorynisbet): Remove this thing.
func describeEnvironment(stderr io.Writer) error {
	command := exec.Command("/bin/sh", "-c", DescriptionCommand)
	// DescriptionCommand writes its contents to stdout, so wire it up to stderr.
	command.Stdout = stderr
	err := command.Run()
	return errors.Annotate(err, "describe environment").Err()
}

// setupContextNamespace sets namespace to the context for UFS client.
func setupContextNamespace(ctx context.Context, namespace string) context.Context {
	md := metadata.Pairs(ufsUtil.Namespace, namespace)
	return metadata.NewOutgoingContext(ctx, md)
}

// getTaskDir return directory for the executed task.
func getTaskDir() (string, error) {
	wDir, err := os.Getwd()
	if err != nil {
		log.Printf("Cannot get task dir: %q", wDir)
		return wDir, errors.Annotate(err, "get task dir").Err()
	}
	absPath, err := filepath.Abs(wDir)
	return absPath, errors.Annotate(err, "get task dir").Err()
}
