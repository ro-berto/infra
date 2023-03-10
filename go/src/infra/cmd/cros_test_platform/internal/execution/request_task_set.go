// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package execution

import (
	"context"
	"fmt"
	"time"

	"infra/cmd/cros_test_platform/internal/execution/args"
	"infra/cmd/cros_test_platform/internal/execution/build"
	"infra/cmd/cros_test_platform/internal/execution/response"
	"infra/cmd/cros_test_platform/internal/execution/retry"
	"infra/cmd/cros_test_platform/internal/execution/testrunner"
	trservice "infra/cmd/cros_test_platform/internal/execution/testrunner/service"
	"infra/cmd/cros_test_platform/internal/execution/types"

	"go.chromium.org/chromiumos/infra/proto/go/test_platform"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/config"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/steps"
	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	luci_retry "go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/grpc/grpcutil"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Retry count on transient errors
const RetryCountOnTransientError = 5

// RequestTaskSet encapsulates the running state of the set of tasks for one
// cros_test_platform request.
type RequestTaskSet struct {
	// Unique names for invocations preserving the order of incoming arguments.
	// This is used to preserve the order in the response.
	invocationIDs []types.InvocationID

	argsGenerators      map[types.InvocationID]*args.Generator
	invocationResponses map[types.InvocationID]*response.Invocation
	activeTasks         map[types.InvocationID]*testrunner.Build
	retryCounter        retry.Counter

	step            *build.RequestStepUpdater
	invocationSteps map[types.InvocationID]*build.InvocationStepUpdater

	launched bool
}

// TaskSetConfig is a wrapper for the parameters common to the testTaskSets.
type TaskSetConfig struct {
	ParentTaskID        string
	ParentBuildID       int64
	RequestUID          string
	Deadline            time.Time
	StatusUpdateChannel *config.Config_PubSub
}

// NewRequestTaskSet creates a new RequestTaskSet.
func NewRequestTaskSet(
	name string,
	buildInstance *bbpb.Build,
	workerConfig *config.Config_SkylabWorker,
	tc *TaskSetConfig,
	params *test_platform.Request_Params,
	tests []*steps.EnumerationResponse_AutotestInvocation) (*RequestTaskSet, error) {

	step := build.NewRequestStep(name, buildInstance)

	invocationIDs := make([]types.InvocationID, len(tests))
	invocationResponses := make(map[types.InvocationID]*response.Invocation)
	argsGenerators := make(map[types.InvocationID]*args.Generator)
	invocationSteps := make(map[types.InvocationID]*build.InvocationStepUpdater)
	tm := make(map[types.InvocationID]*steps.EnumerationResponse_AutotestInvocation)
	for i, test := range tests {
		iid := types.NewInvocationID(i, test)
		invocationIDs[i] = iid
		argsGenerators[iid] = &args.Generator{
			Invocation:          test,
			Params:              params,
			WorkerConfig:        workerConfig,
			ParentTaskID:        tc.ParentTaskID,
			ParentBuildID:       tc.ParentBuildID,
			ParentRequestUID:    tc.RequestUID,
			Deadline:            tc.Deadline,
			StatusUpdateChannel: tc.StatusUpdateChannel,
			Experiments:         buildInstance.GetInput().GetExperiments(),
			GerritChanges:       buildInstance.GetInput().GetGerritChanges(),
		}
		// test, params, workerConfig, tc.ParentTaskID, tc.RequestUID, tc.Deadline)
		invocationResponses[iid] = response.NewInvocation(test.GetTest().GetName())
		invocationSteps[iid] = step.NewInvocationStep(test.GetTest().GetName())
		tm[iid] = test
	}
	return &RequestTaskSet{
		argsGenerators:      argsGenerators,
		invocationIDs:       invocationIDs,
		invocationResponses: invocationResponses,
		activeTasks:         make(map[types.InvocationID]*testrunner.Build),
		retryCounter:        retry.NewCounter(params, tm),
		invocationSteps:     invocationSteps,
		step:                step,
	}, nil
}

// completed returns true if all tasks for this request have completed.
func (r *RequestTaskSet) completed() bool {
	return r.launched && len(r.activeTasks) == 0
}

// LaunchTasks launches initial tasks for all the tests in this request.
func (r *RequestTaskSet) LaunchTasks(ctx context.Context, c trservice.Client) error {
	r.launched = true
	for _, iid := range r.invocationIDs {
		ts := r.getInvocationResponse(iid)
		ag := r.getArgsGenerator(iid)

		if rejected, err := testrunner.ValidateDependencies(ctx, c, ag); err != nil {
			if !testrunner.InvalidDependencies.In(err) {
				return err
			}
			logging.Warningf(ctx, "Dependency validation failed for %s: %s.", ts.Name, err)
			ts.MarkNotRunnable(rejected)
			continue
		}

		task, err := r.createNewBuildWithRetry(ctx, &c, ag, ts.Name)
		if err != nil {
			return errors.Annotate(err, "Error during new test_runner build creation for %s", ts.Name).Err()
		}
		ts.NotifyTask(task)
		r.getInvocationStep(iid).NotifyNewTask(task)
		r.activeTasks[iid] = task
	}
	return nil
}

// retryParams defines retry strategy for handling transient errors
func (r *RequestTaskSet) retryParams() luci_retry.Iterator {
	return &luci_retry.ExponentialBackoff{
		Limited: luci_retry.Limited{
			Delay:    10 * time.Second,
			Retries:  RetryCountOnTransientError,
			MaxTotal: 2 * time.Minute,
		},
		Multiplier: 2,
	}
}

// createNewBuildWithRetry attempts to create new test_runner build. It retries if transient error occurs.
func (r *RequestTaskSet) createNewBuildWithRetry(ctx context.Context, c *trservice.Client, ag *args.Generator, taskName string) (task *testrunner.Build, err error) {
	err = luci_retry.Retry(ctx, transient.Only(r.retryParams), func() error {
		task, err = testrunner.NewBuild(ctx, *c, ag, nil)
		if err != nil {
			if r.isTransientError(ctx, err) {
				logging.Infof(ctx, "Transient error occured for %s: %s", taskName, err.Error())
				return transient.Tag.Apply(err)
			} else {
				logging.Infof(ctx, "Found a non-transient error for %s: %s", taskName, err.Error())
			}
		}
		return err
	}, luci_retry.LogCallback(ctx, "create-new-test_runner-build"))
	return
}

// isTransientError returns if provided error is transient or not
func (r *RequestTaskSet) isTransientError(ctx context.Context, err error) bool {
	if s, ok := status.FromError(err); ok {
		//return grpcutil.IsTransientCode(s.Code()) || s.Code() == codes.DeadlineExceeded
		if grpcutil.IsTransientCode(s.Code()) {
			logging.Infof(ctx, "Found status in grpc transient errors")
			return true
		}
		if s.Code() == codes.DeadlineExceeded {
			logging.Infof(ctx, "Found DeadlineExceeded")
		}
	} else {
		logging.Infof(ctx, "Failed to find error status")
	}
	return false
}

func (r *RequestTaskSet) getInvocationResponse(iid types.InvocationID) *response.Invocation {
	ir, ok := r.invocationResponses[iid]
	if !ok {
		panic(fmt.Sprintf("No test task set for invocation %s", iid))
	}
	return ir
}

func (r *RequestTaskSet) getArgsGenerator(iid types.InvocationID) *args.Generator {
	ag, ok := r.argsGenerators[iid]
	if !ok {
		panic(fmt.Sprintf("No args.Generator for invocation %s", iid))
	}
	return ag
}

func (r *RequestTaskSet) getInvocationStep(iid types.InvocationID) *build.InvocationStepUpdater {
	s, ok := r.invocationSteps[iid]
	if !ok {
		panic(fmt.Sprintf("No step for invocation %s", iid))
	}
	return s
}

// CheckTasksAndRetry checks the status of currently running tasks for this
// request and retries failed tasks when allowed.
//
// Returns whether all tasks are complete (so future calls to this function are
// unnecessary)
func (r *RequestTaskSet) CheckTasksAndRetry(ctx context.Context, c trservice.Client) (bool, error) {
	completedTests := make([]types.InvocationID, len(r.activeTasks))
	newTasks := make(map[types.InvocationID]*testrunner.Build)
	for iid, task := range r.activeTasks {
		rerr := task.Refresh(ctx, c)
		tr := task.Result()
		if rerr != nil {
			return false, errors.Annotate(rerr, "tick for task %s", tr.LogUrl).Err()
		}
		if !task.Completed() {
			continue
		}

		ts := r.getInvocationResponse(iid)
		logging.Infof(ctx, "Task %s (%s) completed with verdict %s", tr.LogUrl, ts.Name, tr.GetState().GetVerdict())

		// At this point, we've determined that latestTask finished, and we've
		// updated the testTaskSet with its result. We can remove it from our
		// attention set... as long as we don't have to retry.
		shouldRetry := retry.IsNeeded(task.Result()) && r.retryCounter.CanRetry(ctx, iid)
		if !shouldRetry {
			completedTests = append(completedTests, iid)
			continue
		}

		logging.Infof(ctx, "Retrying %s", ts.Name)
		for _, tc := range tr.TestCases {
			if tc.Verdict != test_platform.TaskState_VERDICT_PASSED {
				logging.Infof(ctx, "Test (%s) did not pass with Verdict %s: %s", tc.Name, tc.Verdict, tc.HumanReadableSummary)
			}
		}
		nt, err := task.Retry(ctx, c, int32(r.retryCounter.RetryCount(iid)+1))
		if err != nil {
			return false, errors.Annotate(err, "tick for task %s: retry test", tr.LogUrl).Err()
		}
		newTasks[iid] = nt
		ts.NotifyTask(nt)
		r.getInvocationStep(iid).NotifyNewTask(nt)
		r.retryCounter.NotifyRetry(iid)
	}

	for _, iid := range completedTests {
		delete(r.activeTasks, iid)
	}
	for iid, task := range newTasks {
		r.activeTasks[iid] = task
	}
	return r.completed(), nil
}

// Close notifies that all execution for this request has completed.
//
// Finalize must be called exactly once to clean up state.
// It is an error to call any methods except Response() on a Close()ed instance.
func (r *RequestTaskSet) Close() {
	r.step.Close()
}

// Response returns the current response for this request.
func (r *RequestTaskSet) Response() *steps.ExecuteResponse {
	tss := make([]*response.Invocation, len(r.invocationIDs))
	for i, iid := range r.invocationIDs {
		tss[i] = r.invocationResponses[iid]
	}
	return response.Summarize(tss)
}
