// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package run

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"infra/cmd/crosfleet/internal/buildbucket"
	"infra/cmd/crosfleet/internal/common"
	"infra/cmd/crosfleet/internal/flagx"
	"infra/cmd/crosfleet/internal/site"
	"infra/cmdsupport/cmdlib"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/maruel/subcommands"
	"go.chromium.org/luci/auth/client/authcli"
	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/cli"
	"google.golang.org/genproto/protobuf/field_mask"
)

const backfillCmd = "backfill"

var backfill = &subcommands.Command{
	UsageLine: fmt.Sprintf("%s [FLAGS...]", backfillCmd),
	ShortDesc: "backfill unsuccessful results for a previous request",
	LongDesc: `Backfill unsuccessful results for a previous request.

This command creates a new cros_test_platform request to backfill results from
a (finished) previous build.

The backfill request uses the same parameters as the original request (model,
pool, build etc.). The backfill request attempts to minimize unnecessary task
execution by skipping tasks that have succeeded previously when possible.

This command does not wait for the backfill to start running.

This command's behavior is subject to change without notice.
Do not build automation around this subcommand.`,
	CommandRun: func() subcommands.CommandRun {
		c := &backfillRun{}
		c.authFlags.Register(&c.Flags, site.DefaultAuthOptions)
		c.envFlags.Register(&c.Flags)
		c.Flags.Int64Var(&c.buildID, "id", 0, "ID of CTP build to backfill. Mutually exclusive with -tag(s).")
		c.Flags.Var(flagx.KeyVals(&c.buildTags), "tag", `Tag to identify build(s) to backfill, in format key=val or key:val; may be specified multiple times.
Mutually exclusive with -id.`)
		c.Flags.Var(flagx.KeyVals(&c.buildTags), "tags", "Comma-separated build tags in same format as -tag. Mutually exclusive with -id.")
		c.Flags.BoolVar(&c.skipConfirmation, "skip-confirmation", false, "Skip confirmation when backfilling multiple runs.")
		c.Flags.BoolVar(&c.allowDupes, "allow-duplicates", false, "For development purposes only: allow duplicate backfills for the given id/tag(s).")
		// -------------------------------------------------------------------------
		// NOTE: This is not a public feature. Only un-comment this section for
		// locally-built crosfleet executions by the Test Scheduling team.
		//c.Flags.StringVar(&c.qsAccount, "qs-account", "", "For use by the ChromeOS Test Scheduling Team only: override the quota account used for backfills.")
		// -------------------------------------------------------------------------
		return c
	},
}

type backfillRun struct {
	subcommands.CommandRunBase
	authFlags        authcli.Flags
	envFlags         common.EnvFlags
	buildID          int64
	buildTags        map[string]string
	skipConfirmation bool
	allowDupes       bool
	qsAccount        string
}

func (args *backfillRun) Run(a subcommands.Application, _ []string, env subcommands.Env) int {
	if err := args.innerRun(a, env); err != nil {
		cmdlib.PrintError(a, err)
		return 1
	}
	return 0
}

func (args *backfillRun) innerRun(a subcommands.Application, env subcommands.Env) error {
	ctx := cli.GetContext(a, args, env)
	bbService := args.envFlags.Env().BuildbucketService
	ctpBBClient, err := buildbucket.NewClient(ctx, args.envFlags.Env().DefaultCTPBuilder, bbService, args.authFlags)
	if err != nil {
		return err
	}
	originalBuilds, err := args.findOriginalBuilds(ctx, ctpBBClient)
	if err != nil {
		return err
	}
	backfillCount := len(originalBuilds)
	if backfillCount == 0 {
		return fmt.Errorf("no matching, finished build(s) found")
	}
	if backfillCount > 1 && !args.skipConfirmation {
		userPromptReason := fmt.Sprintf("Found %d builds to backfill", backfillCount)
		confirmMultipleBackfills, err := common.CLIPrompt(userPromptReason, false)
		if err != nil {
			return err
		}
		if !confirmMultipleBackfills {
			return nil
		}
	}

	for _, original := range originalBuilds {
		backfillTags := backfillTags(original)
		if !args.allowDupes {
			backfillAlreadyRunning, runningBackfillID, err := ctpBBClient.AnyIncompleteBuildsWithTags(ctx, backfillTags)
			if err != nil {
				return err
			}
			if backfillAlreadyRunning {
				runningBackfillURL := ctpBBClient.BuildURL(runningBackfillID)
				fmt.Fprintf(os.Stdout, "Backfill already running at %s\nfor original build %d\n", runningBackfillURL, original.Id)
				continue
			}
		}
		requests := original.Input.Properties.GetFields()["requests"]
		properties := map[string]interface{}{"requests": requests}
		// -------------------------------------------------------------------------
		// NOTE: This is not a public feature. Only un-comment this section for
		// locally-built crosfleet executions by the Test Scheduling team.
		//if args.qsAccount != "" {
		//	newRequests := changeQuotaAccount(requests.GetStructValue().GetFields(), args.qsAccount)
		//	properties["requests"] = newRequests
		//}
		// -------------------------------------------------------------------------
		newBackfill, err := ctpBBClient.ScheduleBuild(ctx, properties, nil, backfillTags, 0)
		if err != nil {
			return err
		}
		newBackfillURL := ctpBBClient.BuildURL(newBackfill.Id)
		fmt.Fprintf(os.Stdout, "Scheduled backfill at %s\nfor original build %d\n", newBackfillURL, original.Id)
	}
	return nil
}

func (args *backfillRun) findOriginalBuilds(ctx context.Context, bbClient *buildbucket.Client) ([]*buildbucketpb.Build, error) {
	searchByTags := len(args.buildTags) > 0
	searchByID := args.buildID > 0
	if searchByTags == searchByID {
		return nil, fmt.Errorf("must search by -id or -tag(s), but not both")
	}

	if searchByID {
		build, err := bbClient.GetBuild(ctx, args.buildID, "id", "input.properties", "tags")
		if build != nil && (build.Status == buildbucketpb.Status_SCHEDULED || build.Status == buildbucketpb.Status_STARTED) {
			err = fmt.Errorf("can't backfill an unfinished build")
		}
		if err != nil {
			return nil, err
		}
		return []*buildbucketpb.Build{build}, nil
	}

	allBuildsWithTags, err := bbClient.GetAllBuildsWithTags(ctx, args.buildTags, &buildbucketpb.SearchBuildsRequest{
		Predicate: &buildbucketpb.BuildPredicate{
			Status: buildbucketpb.Status_ENDED_MASK,
		},
		Fields: &field_mask.FieldMask{Paths: []string{
			"builds.*.id",
			"builds.*.input.properties",
			"builds.*.tags",
		}},
	})
	if err != nil {
		return nil, err
	}
	return removeBackfills(allBuildsWithTags), nil
}

// removeBackfills removes any backfills from the given list of builds.
func removeBackfills(builds []*buildbucketpb.Build) []*buildbucketpb.Build {
	var filtered []*buildbucketpb.Build
	for _, build := range builds {
		isBackfill := buildbucket.FindTagVal(common.CrosfleetToolTag, build) == backfillCmd
		if !isBackfill {
			filtered = append(filtered, build)
		}
	}
	return filtered
}

// backfillTags constructs backfill-specific tags for a backfill of the given
// build.
func backfillTags(build *buildbucketpb.Build) map[string]string {
	tags := map[string]string{}
	for _, originalTag := range build.Tags {
		tags[originalTag.Key] = originalTag.Value
	}
	tags[common.CrosfleetToolTag] = backfillCmd
	tags["backfill"] = strconv.FormatInt(build.Id, 10)
	return tags
}

// changeQuotaAccount returns a copy of the given map of CTP requests with a
// new quota account set on each request.
func changeQuotaAccount(requests map[string]*structpb.Value, quotaAccount string) map[string]interface{} {
	newRequests := map[string]interface{}{}
	for key, req := range requests {
		reqMap := req.GetStructValue().GetFields()
		newReqMap := map[string]interface{}{}
		for key, val := range reqMap {
			newReqMap[key] = val
		}
		paramsMap := reqMap["params"].GetStructValue().GetFields()
		newParamsMap := map[string]interface{}{}
		for key, val := range paramsMap {
			newParamsMap[key] = val
		}

		// Scheduling
		schedulingMap := paramsMap["scheduling"].GetStructValue().GetFields()
		newSchedulingMap := map[string]interface{}{}
		for key, val := range schedulingMap {
			if key == "priority" {
				continue
			}
			newSchedulingMap[key] = val.GetStringValue()
		}
		newSchedulingMap["qsAccount"] = quotaAccount
		newParamsMap["scheduling"] = newSchedulingMap

		// Software Dependencies
		rawDependencies := paramsMap["softwareDependencies"].GetListValue().GetValues()[0].GetStructValue().GetFields()
		santizedDependencies := map[string]interface{}{}
		for key, val := range rawDependencies {
			santizedDependencies[key] = val.GetStringValue()
		}
		newParamsMap["softwareDependencies"] = []interface{}{santizedDependencies}

		newReqMap["params"] = newParamsMap
		newRequests[key] = newReqMap
	}
	return newRequests
}
