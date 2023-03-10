// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package site contains site local constants for the satlab
package site

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/common/gcloud/gs"
	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/hardcoded/chromeinfra"
)

// AppPrefix is the prefix to use the satlab CLI.
var AppPrefix = "satlab"

// DevCrosAdmService is the dev CrOSSkylabAdmin service.
const DevCrosAdmService = "staging-skylab-bot-fleet.appspot.com"

// ProdCrosAdmService is the prod CrOSSkylabAdmin service.
const ProdCrosAdmService = "chromeos-skylab-bot-fleet.appspot.com"

// DevUFSService is the dev UFS service
const DevUFSService = "staging.ufs.api.cr.dev"

// ProdUFSService is the prod UFS service
const ProdUFSService = "ufs.api.cr.dev"

// Satlab is just the string "satlab". The literal string "satlab" is used in many
// places to create resource names.
const Satlab = "satlab"

// CommonFlags controls some commonly-used CLI flags.
type CommonFlags struct {
	Verbose  bool
	SatlabID string
}

// Register sets up the common flags.
func (f *CommonFlags) Register(fl *flag.FlagSet) {
	fl.BoolVar(&f.Verbose, "verbose", false, "whether to log verbosely")
	fl.StringVar(&f.SatlabID, "satlab-id", "", "the ID for the satlab in question")
}

// OutputFlags controls output-related CLI flags.
type OutputFlags struct {
	json   bool
	tsv    bool
	full   bool
	noemit bool
}

// Register sets up the output flags.
func (f *OutputFlags) Register(fl *flag.FlagSet) {
	fl.BoolVar(&f.json, "json", false, "log output in json format")
	fl.BoolVar(&f.tsv, "tsv", false, "log output in tsv format (without title)")
	fl.BoolVar(&f.full, "full", false, "log full output in specified format")
	fl.BoolVar(&f.noemit, "noemit", false, "specifies NOT to emit/print unpopulated fields in json format.")
}

// JSON returns if the output is logged in json format
func (f *OutputFlags) JSON() bool {
	return f.json
}

// Tsv returns if the output is logged in tsv format (without title)
func (f *OutputFlags) Tsv() bool {
	return f.tsv
}

// Full returns if the full format of output is logged in tsv format (without title)
func (f *OutputFlags) Full() bool {
	return f.full
}

// NoEmit returns if output json should NOT print/emit unpopulated fields
func (f *OutputFlags) NoEmit() bool {
	return f.noemit
}

// EnvFlags controls selection of the environment: either prod (default) or dev.
type EnvFlags struct {
	dev       bool
	Namespace string
}

// DefaultNamespace is the default namespace to use for all operations.
const DefaultNamespace = "os"

// Register sets up the -dev argument.
func (f *EnvFlags) Register(fl *flag.FlagSet) {
	fl.BoolVar(&f.dev, "dev", false, "Run in dev environment.")
	fl.StringVar(&f.Namespace, "namespace", DefaultNamespace, "namespace where data resides.")
}

// GetCrosAdmService returns the hostname of the CrOSSkylabAdmin service that is
// appropriate for the given environment.
func (f *EnvFlags) GetCrosAdmService() string {
	if f.dev {
		return DevCrosAdmService
	}
	return ProdCrosAdmService
}

// GetUFSService returns the hostname of the UFS service appropriate for an environment
func (f *EnvFlags) GetUFSService() string {
	if f.dev {
		return DevUFSService
	}
	return ProdUFSService
}

// DefaultZone is the default value for the zone command line flag.
const DefaultZone = "satlab"

// MaybePrepend adds a prefix with a leading dash unless the string already
// begins with the prefix in question.
func MaybePrepend(prefix string, satlabID string, content string) string {
	if prefix == "" || satlabID == "" {
		return content
	}
	satlabPrefix := fmt.Sprintf("%s-%s", prefix, satlabID)
	if strings.HasPrefix(content, satlabPrefix) {
		return content
	}
	return fmt.Sprintf("%s-%s", satlabPrefix, content)
}

// GetFullyQualifiedHostname takes in primitives used to create the fully qualified hostname and produces that hostname
// specifiedSatlabID refers to a user specified ID which should take precedence over the fetchedSatlabID which should be fetched in an automated fashion
func GetFullyQualifiedHostname(specifiedSatlabID string, fetchedSatlabID, prefix string, content string) string {
	//uses either a user provided or automatically fetched Satlab ID to build the fully qualified host name
	// no-op if the hostname is already in expected format
	satlabIDToUse := specifiedSatlabID
	if satlabIDToUse == "" {
		satlabIDToUse = fetchedSatlabID
	}
	return MaybePrepend(prefix, satlabIDToUse, content)
}

// DefaultAuthOptions is an auth.Options struct prefilled with chrome-infra
// defaults.
var DefaultAuthOptions = chromeinfra.SetDefaultAuthOptions(auth.Options{
	Scopes:     append(gs.ReadWriteScopes, auth.OAuthScopeEmail),
	SecretsDir: SecretsDir(),
})

// SecretsDir customizes the location for auth-related secrets.
func SecretsDir() string {
	configDir := os.Getenv("XDG_CACHE_HOME")
	if configDir == "" {
		configDir = filepath.Join(os.Getenv("HOME"), ".cache")
	}
	return filepath.Join(configDir, "satlab", "auth")
}

// // VersionNumber is the version number for the tool. It follows the Semantic
// // Versioning Specification (http://semver.org) and the format is:
// // "MAJOR.MINOR.0+BUILD_TIME".
// // We can ignore the PATCH part (i.e. it's always 0) to make the maintenance
// // work easier.
// // We can also print out the build time (e.g. 20060102150405) as the METADATA
// // when show version to users.
var VersionNumber = fmt.Sprintf("%d.%d.%d", Major, Minor, Patch)

// // Major is the Major version number
const Major = 0

// // Minor is the Minor version number
const Minor = 1

// // Patch is the PAtch version number
const Patch = 0

// DefaultPRPCOptions is used for PRPC clients.  If it is nil, the
// default value is used.  See prpc.Options for details.
//
// This is provided so it can be overridden for testing.
var DefaultPRPCOptions = prpcOptionWithUserAgent(fmt.Sprintf("satlab/%s", VersionNumber))

// CipdInstalledPath is the installed path for satlab package.
// This is the path to the directory containing main.go relative to the repo root.
var CipdInstalledPath = "infra/cros/cmd/satlab/"

// prpcOptionWithUserAgent create prpc option with custom UserAgent.
//
// DefaultOptions provides Retry ability in case we have issue with service.
func prpcOptionWithUserAgent(userAgent string) *prpc.Options {
	options := *prpc.DefaultOptions()
	options.UserAgent = userAgent
	return &options
}
