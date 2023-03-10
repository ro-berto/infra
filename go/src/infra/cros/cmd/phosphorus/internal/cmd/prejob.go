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
	"strings"
	"time"

	"github.com/maruel/subcommands"
	tlsapi "go.chromium.org/chromiumos/config/go/api/test/tls"
	"go.chromium.org/chromiumos/config/go/api/test/tls/dependencies/longrunning"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/phosphorus"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"infra/cros/cmd/phosphorus/internal/autotest/atutil"
	"infra/cros/cmd/phosphorus/internal/botcache"
	"infra/cros/cmd/phosphorus/internal/tls"
	"infra/libs/lro"
)

// If not set in the prejob requests we default to this bucket.
const ChromeOSImageBucketDefault = "gs://chromeos-image-archive"

type ChromeOSBuildLocation struct {
	ChromeOSBuild  string
	ChromeOSBucket string
}

// Prejob subcommand: Run a prejob (e.g. provision) against a DUT.
var Prejob = &subcommands.Command{
	UsageLine: "prejob -input_json /path/to/input.json",
	ShortDesc: "Run a prejob against a DUT.",
	LongDesc: `Run a prejob against a DUT.

Provision the DUT via 'autoserv --provision' if desired provisionable labels
do not match the existing ones. Otherwise, reset the DUT via
'autosev --reset'`,
	CommandRun: func() subcommands.CommandRun {
		c := &prejobRun{}
		c.Flags.StringVar(&c.InputPath, "input_json", "", "Path that contains JSON encoded test_platform.phosphorus.PrejobRequest")
		c.Flags.StringVar(&c.OutputPath, "output_json", "", "Path to write JSON encoded test_platform.phosphorus.PrejobResponse to")
		return c
	},
}

type prejobRun struct {
	CommonRun
}

func (c *prejobRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
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

func (c *prejobRun) innerRun(ctx context.Context, args []string, env subcommands.Env) error {
	var r phosphorus.PrejobRequest
	if err := ReadJSONPB(c.InputPath, &r); err != nil {
		return err
	}
	if err := validatePrejobRequest(&r); err != nil {
		return err
	}

	if r.Deadline.IsValid() {
		var c context.CancelFunc
		d := r.Deadline.AsTime()
		log.Printf("Running with deadline %s (current time: %s)", d, time.Now().UTC())
		ctx, c = context.WithDeadline(ctx, d)
		defer c()
	}

	var rl []*phosphorus.PrejobRequest
	rl = append(rl, &r)
	// Mimic a PrejobRequest for each additional provision targets.
	for _, provisionTarget := range r.GetAddtionalTargets() {
		peerRequest := phosphorus.PrejobRequest{
			Deadline:               r.GetDeadline(),
			SoftwareDependencies:   provisionTarget.GetSoftwareDependencies(),
			DutHostname:            provisionTarget.GetDutHostname(),
			UseTls:                 r.GetUseTls(),
			Config:                 r.GetConfig(),
			UpdateFirmware:         provisionTarget.GetUpdateFirmware(),
			ProvisionGooglerSshKey: provisionTarget.GetProvisionGooglerSshKey(),
		}
		rl = append(rl, &peerRequest)
	}
	// Run provision for every targets.
	var respList []*phosphorus.PrejobResponse
	for _, req := range rl {
		resp, err := runProvisionSteps(ctx, req)
		if err != nil {
			return err
		}
		respList = append(respList, resp)
	}
	return WriteJSONPB(c.OutputPath, aggregatePrejobResponse(respList))
}

func runProvisionSteps(ctx context.Context, r *phosphorus.PrejobRequest) (*phosphorus.PrejobResponse, error) {
	bt, err := tls.NewBackgroundTLS()
	if err != nil {
		return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_ABORTED}, err
	}
	defer bt.Close()

	resp, err := provisionChromeOSBuild(ctx, bt, r)
	if skipRemainingSteps(resp, err) {
		return resp, err
	}
	resp, err = provisionFirmwareLegacy(ctx, r)
	if skipRemainingSteps(resp, err) {
		return resp, err
	}
	return provisionLacros(ctx, bt, r)
}

func skipRemainingSteps(r *phosphorus.PrejobResponse, err error) bool {
	return !prejobSucceeded(r, err)
}

func prejobSucceeded(r *phosphorus.PrejobResponse, err error) bool {
	return err == nil && r.GetState() == phosphorus.PrejobResponse_SUCCEEDED
}

func provisionChromeOSBuild(ctx context.Context, bt *tls.BackgroundTLS, r *phosphorus.PrejobRequest) (*phosphorus.PrejobResponse, error) {
	if shouldProvisionChromeOSViaTLS(r) {
		return provisionChromeOSBuildViaTLS(ctx, bt, r)
	}
	return provisionChromeOSBuildLegacy(ctx, r)
}

func shouldProvisionChromeOSViaTLS(r *phosphorus.PrejobRequest) bool {
	e := r.GetConfig().GetPrejobStep().GetProvisionDutExperiment()
	if e == nil {
		return false
	}
	if !e.Enabled {
		return false
	}

	v := chromeOSBuildDependencyOrEmpty(r.SoftwareDependencies)
	switch fs := e.GetCrosVersionSelector().(type) {
	case *phosphorus.ProvisionDutExperiment_CrosVersionAllowList:
		for _, p := range fs.CrosVersionAllowList.Prefixes {
			if strings.HasPrefix(v.ChromeOSBuild, p) {
				return true
			}
		}
		return false
	case *phosphorus.ProvisionDutExperiment_CrosVersionDisallowList:
		for _, p := range fs.CrosVersionDisallowList.Prefixes {
			if strings.HasPrefix(v.ChromeOSBuild, p) {
				return false
			}
		}
		return true
	default:
		// Arbitrary default. We don't actually want the config to ever leave
		// this oneof unset.
		return false
	}
}

func provisionChromeOSBuildLegacy(ctx context.Context, r *phosphorus.PrejobRequest) (*phosphorus.PrejobResponse, error) {
	bc := botCache(r)

	desired := chromeOSBuildDependencyOrEmpty(r.SoftwareDependencies).ChromeOSBuild
	found, err := bc.LoadProvisionableLabel(chromeOSBuildKey)
	if err != nil {
		return nil, errors.Annotate(err, "provision chromeos legacy").Err()
	}
	if labelAlreadySatisfied(desired, found) {
		ar, err := resetViaAutoserv(ctx, r)
		return toPrejobResponse(ar), err
	}

	if err := bc.ClearProvisionableLabel(chromeOSBuildKey); err != nil {
		return nil, errors.Annotate(err, "provision chromeos legacy").Err()
	}

	ar, err := provisionViaAutoserv(ctx, "os", r, []string{fmt.Sprintf("%s:%s", chromeOSBuildKey, desired)})
	resp := toPrejobResponse(ar)

	if desired != "" && prejobSucceeded(resp, err) {
		if err := bc.SetNonEmptyProvisionableLabel(chromeOSBuildKey, desired); err != nil {
			return nil, errors.Annotate(err, "provision chromeos legacy").Err()
		}
	}
	return resp, err
}

func labelAlreadySatisfied(desired, found string) bool {
	return desired == "" || desired == found
}

type FirmwareProvision int

const (
	FW_NONE FirmwareProvision = iota
	FW_RO_ONLY
	FW_RW_ONLY
	FW_RO_RW
)

// whichFirmwareToProvision helps determine which of the combination to apply when provisioning firmware:
func whichFirmwareToProvision(roDesired, rwDesired, roFound, rwFound string) FirmwareProvision {
	ro := labelAlreadySatisfied(roDesired, roFound)
	rw := labelAlreadySatisfied(rwDesired, rwFound)

	// Both satisfied, there is nothing to do.
	if ro && rw {
		return FW_NONE
	}

	// Flip flop to provision other firmware.
	if ro {
		return FW_RW_ONLY
	}
	if rw {
		return FW_RO_ONLY
	}

	// Special case to handle so we don't need to run `rw_only` control firmware provisioning.
	if roDesired == rwDesired {
		return FW_RO_ONLY
	}
	return FW_RO_RW
}

func whichFirmwareLabelsToProvision(roDesired, rwDesired, roFound, rwFound string) []string {
	switch w := whichFirmwareToProvision(roDesired, rwDesired, roFound, rwFound); w {
	case FW_RO_ONLY:
		return []string{
			fmt.Sprintf("%s:%s", roFirmwareBuildKey, roDesired),
		}
	case FW_RW_ONLY:
		return []string{
			fmt.Sprintf("%s:%s", rwFirmwareBuildKey, rwDesired),
		}
	case FW_RO_RW:
		return []string{
			fmt.Sprintf("%s:%s", roFirmwareBuildKey, roDesired),
			fmt.Sprintf("%s:%s", rwFirmwareBuildKey, rwDesired),
		}
	default:
		return []string{}
	}
}

func provisionFirmwareLegacy(ctx context.Context, r *phosphorus.PrejobRequest) (*phosphorus.PrejobResponse, error) {
	bc := botCache(r)

	roDesired := roFirmwareBuildDependencyOrEmpty(r.SoftwareDependencies)
	roFound, err := bc.LoadProvisionableLabel(roFirmwareBuildKey)
	if err != nil {
		return nil, errors.Annotate(err, "provision firmware").Err()
	}

	rwDesired := rwFirmwareBuildDependencyOrEmpty(r.SoftwareDependencies)
	rwFound, err := bc.LoadProvisionableLabel(rwFirmwareBuildKey)
	if err != nil {
		return nil, errors.Annotate(err, "provision firmware").Err()
	}

	labels := whichFirmwareLabelsToProvision(roDesired, rwDesired, roFound, rwFound)
	if len(labels) == 0 {
		// Nothing to be done, so succeed trivially.
		return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_SUCCEEDED}, nil
	}

	if err := bc.ClearProvisionableLabel(roFirmwareBuildKey); err != nil {
		return nil, errors.Annotate(err, "provision firmware").Err()
	}
	if err := bc.ClearProvisionableLabel(rwFirmwareBuildKey); err != nil {
		return nil, errors.Annotate(err, "provision firmware").Err()
	}

	ar, err := provisionViaAutoserv(ctx, "firmware", r, labels)
	resp := toPrejobResponse(ar)

	if prejobSucceeded(resp, err) {
		if roDesired != "" {
			if err := bc.SetNonEmptyProvisionableLabel(roFirmwareBuildKey, roDesired); err != nil {
				return nil, errors.Annotate(err, "provision firmware").Err()
			}
		}
		if rwDesired != "" {
			if err := bc.SetNonEmptyProvisionableLabel(rwFirmwareBuildKey, rwDesired); err != nil {
				return nil, errors.Annotate(err, "provision firmware").Err()
			}
		}
	}
	return resp, err
}

func toPrejobResponse(r *atutil.Result) *phosphorus.PrejobResponse {
	switch {
	case r == nil:
		return nil
	case r.Success():
		return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_SUCCEEDED}
	case r.RunResult.Aborted:
		return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_ABORTED}
	default:
		return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_FAILED}
	}
}

const (
	chromeOSBuildKey   = "cros-version"
	roFirmwareBuildKey = "fwro-version"
	rwFirmwareBuildKey = "fwrw-version"
	lacrosPathKey      = "lacros-gcs-path"
)

func validatePrejobRequest(r *phosphorus.PrejobRequest) error {
	missingArgs := getCommonMissingArgs(r.Config)

	if r.DutHostname == "" {
		missingArgs = append(missingArgs, "DUT hostname")
	}

	if len(missingArgs) > 0 {
		return fmt.Errorf("no %s provided", strings.Join(missingArgs, ", "))
	}

	return nil
}

// provisionViaAutoserv provisions a single host. It is a wrapper around
// `autoserv --provision`. It cannot modify its point arguments.
//
// tag is a human-readable name for the provision operation being attempted.
// labels are the list of Tauto labels to be provisioned.
//
// This function will be obsoleted once runTLSProvision is globally enabled.
func provisionViaAutoserv(ctx context.Context, tag string, r *phosphorus.PrejobRequest, labels []string) (*atutil.Result, error) {
	j := getMainJob(r.Config)
	subDir := fmt.Sprintf("provision_%s_%s", r.DutHostname, tag)
	fullPath := filepath.Join(r.Config.Task.ResultsDir, subDir)
	p := &atutil.Provision{
		Host:       r.DutHostname,
		Labels:     labels,
		ResultsDir: fullPath,
	}

	ar, err := atutil.RunAutoserv(ctx, j, p, os.Stdout, nil, nil)
	if err != nil {
		return nil, errors.Annotate(err, "run provision").Err()
	}
	return ar, nil
}

// resetViaAutoserv resets a single host. It is a wrapper around
// `autoserv --reset`.
//
// This function will be obsoleted once runTLSProvision is globally enabled.
func resetViaAutoserv(ctx context.Context, r *phosphorus.PrejobRequest) (*atutil.Result, error) {
	j := getMainJob(r.Config)
	subDir := fmt.Sprintf("reset_%s", r.DutHostname)
	fullPath := filepath.Join(r.Config.Task.ResultsDir, subDir)
	a := &atutil.AdminTask{
		Host:       r.DutHostname,
		ResultsDir: fullPath,
		Type:       atutil.Reset,
	}
	ar, err := atutil.RunAutoserv(ctx, j, a, os.Stdout, nil, nil)
	if err != nil {
		return nil, errors.Annotate(err, "run reset").Err()
	}
	return ar, nil
}

// provisionChromeOSBuildViaTLS provisions a DUT via the TLS API.
// See go/cros-tls go/cros-prover
//
// Errors returned from the actual provision operation are interpreted into
// the response. An error is returned for failure modes that can not be mapped
// to a response.
func provisionChromeOSBuildViaTLS(ctx context.Context, bt *tls.BackgroundTLS, r *phosphorus.PrejobRequest) (*phosphorus.PrejobResponse, error) {
	bc := botCache(r)
	desired := chromeOSBuildDependencyOrEmpty(r.SoftwareDependencies)
	if desired.ChromeOSBuild == "" {
		logging.Infof(ctx, "Skipping Chrome OS image provision because no specific version was requested.")
		return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_SUCCEEDED}, nil
	}

	logging.Infof(ctx, "Adding chromeos-version label to host file")
	hostInfoFileDir := r.Config.GetTask().GetResultsDir()
	err := atutil.AddProvisionDetailsToHostInfoFile(ctx, bt, hostInfoFileDir, r.DutHostname, desired.ChromeOSBucket, desired.ChromeOSBuild)
	if err != nil {
		return nil, errors.Annotate(err, "provision chromeos via tls").Err()
	}

	logging.Infof(ctx, "Copying host file")
	hostSubDir := fmt.Sprintf("provision_%s_os", r.DutHostname)
	fullHostPath := filepath.Join(r.Config.Task.ResultsDir, hostSubDir)
	if err := atutil.LinkHostInfoFile(hostInfoFileDir, fullHostPath, r.DutHostname); err != nil {
		return nil, errors.Annotate(err, "provision chromeos via tls").Err()
	}

	logging.Infof(ctx, "Starting to provision Chrome OS via TLS.")
	if err := bc.ClearProvisionableLabel(chromeOSBuildKey); err != nil {
		return nil, errors.Annotate(err, "provision chromeos via tls").Err()
	}

	c := tlsapi.NewCommonClient(bt.Client)
	op, err := c.ProvisionDut(
		ctx,
		&tlsapi.ProvisionDutRequest{
			Name: r.DutHostname,
			TargetBuild: &tlsapi.ChromeOsImage{
				PathOneof: &tlsapi.ChromeOsImage_GsPathPrefix{
					GsPathPrefix: fmt.Sprintf("%s/%s", desired.ChromeOSBucket, desired.ChromeOSBuild),
				},
			},
			UpdateFirmware:     r.UpdateFirmware,
			OnlyGooglerSshKeys: r.ProvisionGooglerSshKey,
		},
	)
	if err != nil {
		// Errors here indicate a failure in starting the operation, not failure
		// in the provision attempt.
		return nil, errors.Annotate(err, "provision chromeos via tls").Err()
	}

	resp, err := waitForOp(ctx, bt, op)
	if err != nil {
		return nil, errors.Annotate(err, "provision chromeos via tls").Err()
	}

	if desired.ChromeOSBuild != "" && prejobSucceeded(resp, err) {
		if err := bc.SetNonEmptyProvisionableLabel(chromeOSBuildKey, desired.ChromeOSBuild); err != nil {
			return nil, errors.Annotate(err, "provision chromeos via tls").Err()
		}
	}
	return resp, nil
}

func waitForOp(ctx context.Context, bt *tls.BackgroundTLS, op *longrunning.Operation) (*phosphorus.PrejobResponse, error) {
	op, err := lro.Wait(ctx, longrunning.NewOperationsClient(bt.Client), op.GetName())
	if err != nil {
		// TODO(pprabhu) Cancel operation.
		// - Create 60 second headroom before deadline for cancellation.
		// - Cancel operation and wait up to deadline for cancellation to complete.
		// - Return multi-error with failure to cancel, if cancellation fails.
		s, isGRPCErr := status.FromError(err)
		if err == context.DeadlineExceeded || (isGRPCErr && s.Code() == codes.InvalidArgument) {
			return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_ABORTED}, nil
		}
		return nil, err
	}
	if s := op.GetError(); s != nil {
		// Error here is a failure in the provision attempt.
		// TODO(pprabhu) Surface detailed errors up.
		// See https://docs.google.com/document/d/12w5pPntorUY1cgDHHxT3Nu6wdhVox288g5_BnyKCPOE/edit#heading=h.fj6zbs6kop08
		err := fmt.Errorf("Operation failed: (code: %d, message: %+v, details: %s", s.Code, s.Message, s.Details)
		logging.Errorf(ctx, err.Error())
		return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_FAILED}, err
	}
	return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_SUCCEEDED}, nil
}

func provisionLacros(ctx context.Context, bt *tls.BackgroundTLS, r *phosphorus.PrejobRequest) (*phosphorus.PrejobResponse, error) {
	bc := botCache(r)
	desired := lacrosGCSPathOrEmpty(r.SoftwareDependencies)
	if desired == "" {
		logging.Infof(ctx, "Skipping LaCrOS provision because no specific version was requested.")
		return &phosphorus.PrejobResponse{State: phosphorus.PrejobResponse_SUCCEEDED}, nil
	}

	logging.Infof(ctx, "Starting to provision LaCrOS.")
	if err := bc.ClearProvisionableLabel(lacrosPathKey); err != nil {
		return nil, errors.Annotate(err, "provision lacros").Err()
	}

	c := tlsapi.NewCommonClient(bt.Client)
	op, err := c.ProvisionLacros(
		ctx,
		&tlsapi.ProvisionLacrosRequest{
			Name: r.DutHostname,
			Image: &tlsapi.ProvisionLacrosRequest_LacrosImage{
				PathOneof: &tlsapi.ProvisionLacrosRequest_LacrosImage_GsPathPrefix{
					GsPathPrefix: desired,
				},
			},
		},
	)
	if err != nil {
		// Errors here indicate a failure in starting the operation, not failure
		// in the provision attempt.
		return nil, errors.Annotate(err, "provision lacros").Err()
	}

	resp, err := waitForOp(ctx, bt, op)
	if err != nil {
		return nil, errors.Annotate(err, "provision lacros").Err()
	}

	if desired != "" && prejobSucceeded(resp, err) {
		if err := bc.SetNonEmptyProvisionableLabel(lacrosPathKey, desired); err != nil {
			return nil, errors.Annotate(err, "provision lacros").Err()
		}
	}
	return resp, nil
}

func chromeOSBuildDependencyOrEmpty(deps []*test_platform.Request_Params_SoftwareDependency) *ChromeOSBuildLocation {
	bLoc := &ChromeOSBuildLocation{}

	for _, d := range deps {
		switch b := d.Dep.(type) {
		case *test_platform.Request_Params_SoftwareDependency_ChromeosBuild:
			bLoc.ChromeOSBuild = b.ChromeosBuild
		case *test_platform.Request_Params_SoftwareDependency_ChromeosBuildGcsBucket:
			bucket := b.ChromeosBuildGcsBucket
			if !strings.HasPrefix(bucket, "gs://") {
				bucket = "gs://" + bucket
			}
			bLoc.ChromeOSBucket = bucket
		}
	}

	// If we haven't found a bucket, but have a build set the default.
	if bLoc.ChromeOSBucket == "" && bLoc.ChromeOSBuild != "" {
		bLoc.ChromeOSBucket = ChromeOSImageBucketDefault
	}
	return bLoc
}

func roFirmwareBuildDependencyOrEmpty(deps []*test_platform.Request_Params_SoftwareDependency) string {
	for _, d := range deps {
		if b, ok := d.Dep.(*test_platform.Request_Params_SoftwareDependency_RoFirmwareBuild); ok {
			return b.RoFirmwareBuild
		}
	}
	return ""
}

func rwFirmwareBuildDependencyOrEmpty(deps []*test_platform.Request_Params_SoftwareDependency) string {
	for _, d := range deps {
		if b, ok := d.Dep.(*test_platform.Request_Params_SoftwareDependency_RwFirmwareBuild); ok {
			return b.RwFirmwareBuild
		}
	}
	return ""
}

func lacrosGCSPathOrEmpty(deps []*test_platform.Request_Params_SoftwareDependency) string {
	for _, d := range deps {
		if b, ok := d.Dep.(*test_platform.Request_Params_SoftwareDependency_LacrosGcsPath); ok {
			return b.LacrosGcsPath
		}
	}
	return ""
}

func botCache(r *phosphorus.PrejobRequest) *botcache.Store {
	return &botcache.Store{
		CacheDir: r.GetConfig().GetBot().GetAutotestDir(),
		Name:     r.GetDutHostname(),
	}
}

func aggregatePrejobResponse(respList []*phosphorus.PrejobResponse) *phosphorus.PrejobResponse {
	// The State in aggregated response should be set to highest severity level.
	// We consider Severity level as: UNSPECIFIED > FAILED > ABORTED > SUCCEEDED.
	severityMap := map[phosphorus.PrejobResponse_State]int{
		phosphorus.PrejobResponse_STATE_UNSPECIFIED: 4,
		phosphorus.PrejobResponse_FAILED:            3,
		phosphorus.PrejobResponse_ABORTED:           2,
		phosphorus.PrejobResponse_SUCCEEDED:         1,
	}
	state := phosphorus.PrejobResponse_SUCCEEDED
	for _, resp := range respList {
		if s, ok := severityMap[resp.State]; ok {
			if s > severityMap[state] {
				state = resp.State
			}
		} else {
			// We don't expected to hit this point, but if we do we set state to UNSPECIFIED
			// which is the highest severity.
			state = phosphorus.PrejobResponse_STATE_UNSPECIFIED
		}
	}
	return &phosphorus.PrejobResponse{State: state}
}
