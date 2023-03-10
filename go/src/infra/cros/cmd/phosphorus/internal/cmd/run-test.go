// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/maruel/subcommands"
	"github.com/pkg/errors"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/phosphorus"
	"go.chromium.org/luci/common/cli"

	"infra/cros/cmd/phosphorus/internal/autotest/atutil"
	"infra/cros/internal/cmd"
)

const (
	// Maximum build milestone until which SSP is not supported.
	// Context: b/248294114
	MaxBuildMilestoneNotSupportingSSP = 105
)

// RunTest subcommand: Run a test against one or multiple DUTs.
var RunTest = &subcommands.Command{
	UsageLine: "run-test -input_json /path/to/input.json",
	ShortDesc: "Run a test against one or multiple DUTs.",
	LongDesc: `Run a test against one or multiple DUTs.

A wrapper around 'autoserv'.`,
	CommandRun: func() subcommands.CommandRun {
		c := &runTestRun{}
		c.Flags.StringVar(&c.InputPath, "input_json", "", "Path that contains JSON encoded test_platform.phosphorus.RunTestRequest")
		c.Flags.StringVar(&c.OutputPath, "output_json", "", "Path to write JSON encoded test_platform.phosphorus.RunTestResponse to")
		return c
	},
}

type runTestRun struct {
	CommonRun
}

type autoservResult struct {
	CmdResult  *atutil.Result
	ResultsDir string
}

func (c *runTestRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.ValidateArgs(); err != nil {
		fmt.Fprintln(a.GetErr(), err.Error())
		c.Flags.Usage()
		return 1
	}

	ctx := cli.GetContext(a, c, env)
	if err := c.innerRun(ctx, args, env); err != nil {
		logApplicationError(ctx, a, err)
		return 1
	}
	return 0
}

func (c *runTestRun) innerRun(ctx context.Context, args []string, env subcommands.Env) error {
	r := &phosphorus.RunTestRequest{}
	if err := ReadJSONPB(c.InputPath, r); err != nil {
		return err
	}
	if err := validateRunTestRequest(r); err != nil {
		return err
	}

	if d := r.Deadline.AsTime(); !d.IsZero() {
		var c context.CancelFunc
		log.Printf("Running with deadline %s (current time: %s)", d, time.Now().UTC())
		log.Printf("Deadline length(seconds): %.0f", d.Sub(time.Now().UTC()).Seconds())
		ctx, c = context.WithDeadline(ctx, d)
		defer c()
	}
	ar, err := runTestStep(ctx, r)
	if err != nil {
		return err
	}
	return WriteJSONPB(c.OutputPath, runTestResponse(ar))
}

func runTestResponse(r *autoservResult) *phosphorus.RunTestResponse {
	return &phosphorus.RunTestResponse{
		ResultsDir: r.ResultsDir,
		State:      runTestState(r.CmdResult),
	}
}

func runTestState(r *atutil.Result) phosphorus.RunTestResponse_State {
	if r.Success() {
		return phosphorus.RunTestResponse_SUCCEEDED
	}
	if r.RunResult.Aborted {
		return phosphorus.RunTestResponse_ABORTED
	}
	return phosphorus.RunTestResponse_FAILED
}

func validateRunTestRequest(r *phosphorus.RunTestRequest) error {
	missingArgs := getCommonMissingArgs(r.Config)

	if len(r.DutHostnames) == 0 {
		missingArgs = append(missingArgs, "DUT hostname(s)")
	}

	if r.GetAutotest().GetName() == "" {
		missingArgs = append(missingArgs, "test name")
	}

	if len(missingArgs) > 0 {
		return fmt.Errorf("no %s provided", strings.Join(missingArgs, ", "))
	}

	return nil
}

// runTestStep runs an individual test. It is a wrapper around autoserv.
func runTestStep(ctx context.Context, r *phosphorus.RunTestRequest) (*autoservResult, error) {
	j := getMainJob(r.Config)

	dir := filepath.Join(r.Config.Task.ResultsDir, "autoserv_test")

	requireSSP := !r.GetAutotest().GetIsClientTest()

	// SSP is not supported for builds before a certain milestone.
	// Context: b/248294114
	if r.GetAutotest().GetName() == "autoupdate_EndToEndTest" {
		keyvals := r.GetAutotest().GetKeyvals()
		build, ok := keyvals["build"]
		if ok {
			buildVersionRegex := regexp.MustCompile(`^.*/R(\d{3,})-.*`)
			matches := buildVersionRegex.FindStringSubmatch(build)
			if len(matches) > 1 {
				buildMilestone, err := strconv.Atoi(matches[1])
				if err != nil {
					log.Printf("Failed to convert build milestone %s", matches[1])
				} else if buildMilestone <= MaxBuildMilestoneNotSupportingSSP {
					log.Printf("Setting requireSSP as false for build milestone %d", buildMilestone)
					requireSSP = false
				}
			} else {
				log.Printf("Failed to match to regex to acquire build milestone")
			}
		} else {
			log.Printf("`build` not found in keyvals")
		}
	}

	t := &atutil.Test{
		Args:               r.GetAutotest().GetTestArgs(),
		ClientTest:         r.GetAutotest().GetIsClientTest(),
		ControlName:        r.GetAutotest().GetName(),
		ImageStorageServer: r.GetAutotest().GetImageStorageServer(),
		Hosts:              r.DutHostnames,
		Keyvals:            r.GetAutotest().GetKeyvals(),
		Name:               r.GetAutotest().GetDisplayName(),
		PeerDuts:           r.GetAutotest().GetPeerDuts(),
		RequireSSP:         requireSSP,
		ResultsDir:         dir,
		SSPBaseImageName:   r.Config.GetTask().GetSspBaseImageName(),
	}

	var cmdRunner cmd.CommandRunner

	if r.ContainerImageInfo != nil {
		cmdRunner = cmd.RealCommandRunner{}
	}

	ar, err := atutil.RunAutoserv(ctx, j, t, os.Stdout, cmdRunner, r.ContainerImageInfo)
	if err != nil {
		return nil, errors.Wrap(err, "run test")
	}
	return &autoservResult{
		CmdResult:  ar,
		ResultsDir: dir,
	}, nil
}
