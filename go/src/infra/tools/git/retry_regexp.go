// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"regexp"
	"strings"
)

// defaultGitRetry is the set of RE2-formatted Regular Expressions to add
// to the DefaultGitRetryRegexp.
//
// defaultGitRetryPOSIX was originally translated from "chromite":
// https://chromium.googlesource.com/chromiumos/chromite/+/07d4626c40a501866d7c01954f8cabef7b50f482/lib/git.py#29
var defaultGitRetryRegexpSource = []string{
	// crbug.com/285832
	`!.*\[remote rejected\].*\(error in hook\)`,

	// crbug.com/289932
	`!.*\[remote rejected\].*\(failed to lock\)`,

	// crbug.com/307156
	`!.*\[remote rejected\].*\(error in Gerrit backend\)`,

	// crbug.com/285832
	`remote error: Internal Server Error`,

	// crbug.com/294449
	`fatal: Couldn't find remote ref `,

	// crbug.com/220543
	`git fetch_pack: expected ACK/NAK, got`,

	// crbug.com/189455
	`protocol error: bad pack header`,

	// crbug.com/202807
	`The remote end hung up unexpectedly`,
	`The remote end hung up upon initial contact`,

	// crbug.com/298189
	`TLS packet with unexpected length was received`,

	// crbug.com/187444
	`RPC failed; result=\d+, HTTP code = \d+`,

	// crbug.com/388876
	`Connection timed out`,

	// crbug.com/451458, b/19202011
	`repository cannot accept new pushes; contact support`,

	// crbug.com/535306
	`Service Temporarily Unavailable`,

	// fxbug.dev/78001
	`The service is currently unavailable`,

	// crbug.com/675262
	`Connection refused`,

	// crbug.com/430343
	`The requested URL returned error: 5\d+`,

	// crbug.com/725233
	`Operation too slow`,

	`Connection reset by peer`,

	`Unable to look up`,
	`Couldn't resolve host`,
	`Unknown SSL protocol error`,

	// b/110032771
	`Revision .* of patch set \d+ does not match refs/changes`,

	// crbug.com/850130 & crbug.com/898208
	`Git repository not found`,
	`Couldn't connect to server`,
	`transfer closed with outstanding read data remaining`,
	`Access denied to`,

	// crbug.com/1061473
	`The requested URL returned error: 429`,

	// b/155578399
	`RESOURCE_EXHAUSTED`,
	`Resource has been exhausted`,
	`check quota`,

	// b/158498614
	`fetch-pack: protocol error: bad band #\d+`,

	// b/170222146
	`The requested URL returned error: 400`,

	// fxbug.dev/63100
	`fetch-pack: fetch failed`,
	`fetch-pack: unable to spawn http-fetch`,
	`fetch-pack: expected keep then TAB at start of http-fetch output`,
	`fetch-pack: expected hash then LF at end of http-fetch output`,
	`fetch-pack: unable to finish http-fetch`,
	`fetch-pack: pack downloaded from .* does not match expected hash .*`,
	`fetch-pack: invalid index-pack output`,
	`fetch-pack: unexpected disconnect while reading sideband packet`,

	// fxbug.dev/87312
	`fatal: expected flush after ref listing`,
	`fatal: expected response end packet after ref listing`,
	`error: [^/].+ did not send all necessary objects`,
	`fatal: .*: OpenSSL SSL_\w+: SSL_ERROR_SYSCALL`,
	`fatal: .*: SSL_\w+ returned SYSCALL`,
	`fatal: .*: Empty reply from server`,
	`fatal: early EOF`,
	`fatal: error processing wanted refs`,
	`fatal: .*: Server aborted the SSL handshake`,

	// fxbug.dev/91419
	`fatal: .* trying to write ref .* with nonexistent object`,
}

// DefaultGitRetryRegexp is the set of default transient regular expressions to
// retry on.
var DefaultGitRetryRegexp *regexp.Regexp

func init() {
	if len(defaultGitRetryRegexpSource) > 0 {
		DefaultGitRetryRegexp = regexp.MustCompile(mergeRegex(defaultGitRetryRegexpSource))
	}
}

// mergeRegex merges multiple regular expression strings together into a single
// "|"-delimited regular expression group. No capture groups are introduced in
// this merge.
func mergeRegex(regexps []string) string {
	// Merge all of the regex into a single regex.
	allRE := make([]string, len(regexps))
	for i, re := range regexps {
		allRE[i] = "(?:" + re + ")"
	}
	return "(?i)(?:" + strings.Join(allRE, "|") + ")"
}
