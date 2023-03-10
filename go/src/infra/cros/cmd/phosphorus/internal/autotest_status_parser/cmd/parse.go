// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/maruel/subcommands"

	"go.chromium.org/chromiumos/infra/proto/go/test_platform/skylab_test_runner"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/logging"
)

const (
	testSubdir         = "autoserv_test"
	resultsSummaryFile = "status.log"
	exitStatusFile     = ".autoserv_execute"

	verdictStringPrefix  = "END "
	undefinedName        = "----"
	rebootTestName       = "reboot"
	rebootVerifyTestName = "reboot.verify"
	suspendTestName      = "suspend"

	dutStateReady       = "ready"
	dutStateNeedsRepair = "needs_repair"
)

var prejobPrefixes = []string{"provision", "prejob"}

// Parse subcommand: Extract test case results from status.log.
var Parse = &subcommands.Command{
	UsageLine: "parse DIR",
	ShortDesc: "Extract test case results from an autotest results directory.",
	LongDesc: `Extract test case results from an autotest results directory.

Parse the test result summary file (status.log) and the exit code file
(.autoserv_execute) created by the autotest harness inside DIR. The parsing is
a simplified version of the one done by tko/parse.

Write the parsed test case results as a JSON-pb
test_platform/skylab_test_runner/result.proto to stdout.`,
	CommandRun: func() subcommands.CommandRun {
		c := &parseRun{}
		return c
	},
}

type parseRun struct {
	subcommands.CommandRunBase
}

func (c *parseRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	ctx := cli.GetContext(a, c, env)
	if err := c.innerRun(ctx, a, args, env); err != nil {
		fmt.Fprintf(a.GetErr(), "%s\n", err)
		return 1
	}
	return 0
}

func (c *parseRun) innerRun(ctx context.Context, a subcommands.Application, args []string, env subcommands.Env) error {
	if err := c.validateArgs(); err != nil {
		return err
	}
	dir := c.Flags.Args()[0]

	testDir := filepath.Join(dir, testSubdir)
	autotestResult := getTestResults(ctx, testDir)

	prejobResult := getPrejobResults(ctx, dir)

	result := skylab_test_runner.Result{
		Harness: &skylab_test_runner.Result_AutotestResult{
			AutotestResult: &autotestResult,
		},
		Prejob: prejobResult,
		StateUpdate: &skylab_test_runner.Result_StateUpdate{
			DutState: getDutState(prejobResult, &autotestResult),
		},
	}

	return printProtoJSON(a.GetOut(), &result)
}

func (c *parseRun) validateArgs() error {
	if c.Flags.NArg() != 1 {
		return errors.New("must specify exactly 1 results directory to parse")
	}
	return nil
}

// getTestResults extracts all test case results from the status.log file
// inside the given results directory.
func getTestResults(ctx context.Context, dir string) skylab_test_runner.Result_Autotest {
	resultsSummaryPath := filepath.Join(dir, resultsSummaryFile)
	resultsSummaryContent, err := ioutil.ReadFile(resultsSummaryPath)

	if err != nil {
		logging.Errorf(ctx, err.Error())
		// Errors in reading status.log are expected when the server
		// job is aborted.
		return skylab_test_runner.Result_Autotest{
			Incomplete: true,
		}
	}

	testCases := parseResultsFile(string(resultsSummaryContent))

	exitStatusFilePath := filepath.Join(dir, exitStatusFile)
	exitStatusContent, err := ioutil.ReadFile(exitStatusFilePath)

	if err != nil {
		logging.Errorf(ctx, err.Error())
		return skylab_test_runner.Result_Autotest{
			TestCases:  testCases,
			Incomplete: true,
		}
	}

	incomplete := exitedWithErrors(string(exitStatusContent)) || len(testCases) == 0

	return skylab_test_runner.Result_Autotest{
		TestCases:  testCases,
		Incomplete: incomplete,
	}
}

type prejobInfo struct {
	Name string
	Dir  string
}

func getPrejobResults(ctx context.Context, dir string) *skylab_test_runner.Result_Prejob {
	var steps []*skylab_test_runner.Result_Prejob_Step

	prejobs, err := getPrejobs(dir)

	if err != nil {
		logging.Errorf(ctx, err.Error())
		return &skylab_test_runner.Result_Prejob{
			Step: []*skylab_test_runner.Result_Prejob_Step{
				{
					Name:    "unknown",
					Verdict: skylab_test_runner.Result_Prejob_Step_VERDICT_FAIL,
				},
			},
		}
	}

	for _, prejob := range prejobs {
		steps = append(steps, &skylab_test_runner.Result_Prejob_Step{
			Name:    prejob.Name,
			Verdict: getPrejobVerdict(prejob.Name, prejob.Dir),
		})
	}
	return &skylab_test_runner.Result_Prejob{
		Step: steps,
	}
}

func getPrejobs(parentDir string) ([]prejobInfo, error) {
	prejobs := []prejobInfo{}
	for _, prefix := range prejobPrefixes {
		fullPrefix := filepath.Join(parentDir, prefix)
		regex := fullPrefix + "*"
		prejobDirList, err := filepath.Glob(regex)
		if err != nil {
			return nil, fmt.Errorf("bad regex: %s", regex)
		}
		for _, dir := range prejobDirList {
			prejobs = append(prejobs, prejobInfo{
				Name: prefix,
				Dir:  dir,
			})
		}
	}
	return prejobs, nil
}

func getPrejobVerdict(prejobName, prejobDir string) skylab_test_runner.Result_Prejob_Step_Verdict {
	exitStatusFilePath := filepath.Join(prejobDir, exitStatusFile)

	if prejobName == "provision" {
		if _, err := os.Stat(exitStatusFilePath); err != nil {
			// If the prejob is "provision" and the .autoserv_execute file
			// does not exist, we assume that provisioning was done via TLS,
			// not Autoserv (see b/182579728, b/182416536).
			// We assume success, relying on an earlier step to indicate failure.
			return skylab_test_runner.Result_Prejob_Step_VERDICT_PASS
		}
	}
	exitStatusContent, err := ioutil.ReadFile(exitStatusFilePath)

	if err != nil {
		return skylab_test_runner.Result_Prejob_Step_VERDICT_FAIL
	}

	if exitedWithErrors(string(exitStatusContent)) {
		return skylab_test_runner.Result_Prejob_Step_VERDICT_FAIL
	}

	return skylab_test_runner.Result_Prejob_Step_VERDICT_PASS
}

// parseResultsFile extracts all test case results from contents of a
// status.log file.
func parseResultsFile(contents string) []*skylab_test_runner.Result_Autotest_TestCase {
	var stack testCaseStack
	lines := strings.Split(contents, "\n")
	testCases := []*skylab_test_runner.Result_Autotest_TestCase{}

	// Stores nested aborted test cases because they might not be legal test
	// case (e.g. reboot) but they imply the outer test case is aborted.
	nestedAbortTestCases := []*skylab_test_runner.Result_Autotest_TestCase{}
	for _, line := range lines {
		stack.ParseLine(line)

		testCase := stack.PopFullyParsedTestCase()

		if testCase == nil {
			continue
		}

		if isLegalTestCaseName(testCase.Name) {
			processNestedAbortTestCase(testCase, nestedAbortTestCases)
			testCases = append(testCases, testCase)
		} else {
			if testCase.Name == rebootTestName && testCase.Verdict == skylab_test_runner.Result_Autotest_TestCase_VERDICT_ABORT {
				nestedAbortTestCases = append(nestedAbortTestCases, testCase)
			}
		}
	}

	testCases = append(testCases, stack.FlushLegalTestCaseName()...)

	// Stay consistent with the default value which is nil.
	if len(testCases) == 0 {
		return nil
	}
	return testCases
}

type testCaseStack struct {
	testCases []*skylab_test_runner.Result_Autotest_TestCase
}

// ParseLine parses a line from status.log in the context of the current test
// case stack.
func (s *testCaseStack) ParseLine(line string) {
	parts := strings.Split(strings.TrimLeft(line, "\t"), "\t")
	if len(parts) < 3 {
		return
	}

	// The last part of a line is a failure summary for final events and status
	// update events.
	failureSummary := parts[len(parts)-1]
	switch {
	case isStartingEvent(parts[0]):
		testCaseName := parts[2] // Declared test name if any.
		if testCaseName == undefinedName {
			// Use subdir name if declared name not available.
			testCaseName = parts[1]
		}
		s.push(testCaseName)

	case isFinalEvent(parts[0]):
		// ABORT status in the nested job overrides FAIL status in the
		// outer final event statement.
		if len(s.testCases) > 0 && s.testCases[len(s.testCases)-1].Verdict == skylab_test_runner.Result_Autotest_TestCase_VERDICT_ABORT {
			return
		}

		// Sets the failure summary if it exists in the ABORT final event.
		if strings.Contains(parts[0], "ABORT") && len(failureSummary) > 0 {
			s.addSummary(failureSummary)
		}

		s.setVerdict(parseVerdict(parts[0]))

	case isStatusUpdateEvent(parts[0]):
		s.addSummary(failureSummary)

		// ABORT status might be recorded in the nested reboot.verify job.
		if parts[0] == "ABORT" && parts[2] == rebootVerifyTestName {
			s.setVerdict(skylab_test_runner.Result_Autotest_TestCase_VERDICT_ABORT)
		}
	}
}

// PopFullyParsedTestCase pops and returns the top test case of the stack, if
// it is fully parsed. If the top of the stack is only partially parsed, this
// function returns nil.
func (s *testCaseStack) PopFullyParsedTestCase() *skylab_test_runner.Result_Autotest_TestCase {
	if len(s.testCases) == 0 {
		return nil
	}

	r := s.testCases[len(s.testCases)-1]

	// The test case is not fully parsed.
	if r.Verdict == skylab_test_runner.Result_Autotest_TestCase_VERDICT_UNDEFINED {
		return nil
	}

	s.testCases = s.testCases[:len(s.testCases)-1]
	return r
}

// Flush pops all the legal test cases currently in the stack, declares them
// aborted and returns them.
func (s *testCaseStack) FlushLegalTestCaseName() []*skylab_test_runner.Result_Autotest_TestCase {
	r := []*skylab_test_runner.Result_Autotest_TestCase{}

	for {
		// Set the incomplete test case as ABORT
		s.setVerdict(skylab_test_runner.Result_Autotest_TestCase_VERDICT_ABORT)

		tc := s.PopFullyParsedTestCase()

		if tc == nil {
			break
		}

		if isLegalTestCaseName(tc.Name) {
			r = append(r, tc)
		}
	}

	return r
}

// push adds a test case with a given name to the stack.
func (s *testCaseStack) push(name string) {
	if s == nil {
		panic("the test case stack is nil")
	}

	tc := skylab_test_runner.Result_Autotest_TestCase{
		Name: name,
	}
	s.testCases = append(s.testCases, &tc)
}

// addSummary appends a string to the human_readable_summary of the top
// test case in the stack.
func (s *testCaseStack) addSummary(summary string) {
	// Ignore comments that do not correspond to a specific test case.
	if len(s.testCases) == 0 {
		return
	}
	if summary == "" {
		return
	}

	s.testCases[len(s.testCases)-1].HumanReadableSummary += summary + "\n"
}

// setVerdict sets the verdict of the top test case in the stack.
func (s *testCaseStack) setVerdict(verdict skylab_test_runner.Result_Autotest_TestCase_Verdict) {
	if len(s.testCases) == 0 {
		return
	}

	s.testCases[len(s.testCases)-1].Verdict = verdict
}

// processNestedAbortTestCase processes if there is any nested aborted test
// case. If so, update the top test case with the abort status and error
// summary.
func processNestedAbortTestCase(testCase *skylab_test_runner.Result_Autotest_TestCase, nestedAbortTestCases []*skylab_test_runner.Result_Autotest_TestCase) {
	nestedAbortTestCaseSize := len(nestedAbortTestCases)
	if nestedAbortTestCaseSize == 0 {
		return
	}

	// Updates the verdict and error summary of the current test case.
	topAbortTestCase := nestedAbortTestCases[nestedAbortTestCaseSize-1]
	testCase.Verdict = topAbortTestCase.Verdict
	testCase.HumanReadableSummary = topAbortTestCase.HumanReadableSummary

	// Pops the top nested aborted test case.
	nestedAbortTestCases = nestedAbortTestCases[:nestedAbortTestCaseSize-1]
}

// True for the event string from the first line of a test case.
func isStartingEvent(event string) bool {
	return event == "START"
}

// True for the event string from the final line of a test case.
func isFinalEvent(event string) bool {
	return strings.HasPrefix(event, verdictStringPrefix)
}

// True for an event string that may contain a failure/warning reason.
func isStatusUpdateEvent(event string) bool {
	switch event {
	case "FAIL", "WARN", "ERROR", "ABORT", "TEST_NA":
		return true
	}
	return false
}

// isLegalTestCaseName filters out uninformative execution steps.
func isLegalTestCaseName(name string) bool {
	switch name {
	case rebootTestName, suspendTestName, undefinedName, "":
		return false
	}
	return true
}

// parseVerdict converts a verdict string from status.log (e.g. "END GOOD",
// "END FAIL" etc) into a Verdict proto.
func parseVerdict(verdict string) skylab_test_runner.Result_Autotest_TestCase_Verdict {
	// Remove "END " prefix.
	switch verdict[len(verdictStringPrefix):] {
	case "GOOD", "WARN":
		return skylab_test_runner.Result_Autotest_TestCase_VERDICT_PASS
	case "ERROR":
		return skylab_test_runner.Result_Autotest_TestCase_VERDICT_ERROR
	case "ABORT":
		return skylab_test_runner.Result_Autotest_TestCase_VERDICT_ABORT
	case "TEST_NA":
		return skylab_test_runner.Result_Autotest_TestCase_VERDICT_NO_VERDICT
	}
	return skylab_test_runner.Result_Autotest_TestCase_VERDICT_FAIL
}

func requiresDutRepair(v skylab_test_runner.Result_Autotest_TestCase_Verdict) bool {
	switch v {
	case skylab_test_runner.Result_Autotest_TestCase_VERDICT_PASS:
		return false
	case skylab_test_runner.Result_Autotest_TestCase_VERDICT_NO_VERDICT:
		return false
	default:
		return true
	}
}

// getDutState returns the state of the DUT after the prejob and test run.
func getDutState(prejob *skylab_test_runner.Result_Prejob, tests *skylab_test_runner.Result_Autotest) string {
	for _, s := range prejob.Step {
		if s.GetVerdict() != skylab_test_runner.Result_Prejob_Step_VERDICT_PASS {
			return dutStateNeedsRepair
		}
	}
	if tests.Incomplete {
		return dutStateNeedsRepair
	}
	for _, tc := range tests.TestCases {
		if requiresDutRepair(tc.GetVerdict()) {
			return dutStateNeedsRepair
		}
	}
	return dutStateReady
}

// printProtoJSON prints the parsed test cases as a JSON-pb to stdout.
func printProtoJSON(w io.Writer, result *skylab_test_runner.Result) error {
	m := jsonpb.Marshaler{
		EnumsAsInts: false,
		Indent:      "\t",
	}
	return m.Marshal(w, result)
}
