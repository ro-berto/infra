// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package crauditcommits implements cr-audit-commits.appspot.com services.
package crauditcommits

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/context"

	"go.chromium.org/luci/common/api/gerrit"
	"go.chromium.org/luci/common/api/gitiles"
	"go.chromium.org/luci/server/auth"

	buildbot "infra/monitoring/messages"
)

const (
	// TODO(robertocn): Move this to the gitiles library.
	gitilesScope      = "https://www.googleapis.com/auth/gerritcodereview"
	emailScope        = "https://www.googleapis.com/auth/userinfo.email"
	failedBuildPrefix = "Sample Failed Build:"
)

type gerritClientInterface interface {
	GetChangeDetails(context.Context, string, []string) (*gerrit.Change, error)
	ChangeQuery(context.Context, gerrit.ChangeQueryRequest) ([]*gerrit.Change, bool, error)
}

type gitilesClientInterface interface {
	LogForward(context.Context, string, string, string) ([]gitiles.Commit, error)
	Log(context.Context, string, string, int) ([]gitiles.Commit, error)
}

type miloClientInterface interface {
	GetBuildInfo(context.Context, string) (*buildbot.Build, error)
}

// getGitilesClient creates a new gitiles client bound to a new http client
// that is bound to an authenticated transport.
func getGitilesClient(ctx context.Context) (*gitiles.Client, error) {
	httpClient, err := getAuthenticatedHTTPClient(ctx)
	if err != nil {
		return nil, err
	}
	return &gitiles.Client{Client: httpClient}, nil
}

// TODO(robertocn): move this into a dedicated file for authentication, and
// accept a list of scopes to make this function usable for communicating for
// different systems.
func getAuthenticatedHTTPClient(ctx context.Context) (*http.Client, error) {
	t, err := auth.GetRPCTransport(ctx, auth.AsSelf, auth.WithScopes(gitilesScope, emailScope))
	if err != nil {
		return nil, err
	}
	return &http.Client{Transport: t}, nil
}

func failedBuildFromCommitMessage(m string) (string, error) {
	s := bufio.NewScanner(strings.NewReader(m))
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, failedBuildPrefix) {
			return strings.TrimSpace(strings.TrimPrefix(line, failedBuildPrefix)), nil
		}
	}
	return "", fmt.Errorf(
		"commit message does not contain url to failed build prefixed with %q",
		failedBuildPrefix)
}

func getFailedBuild(ctx context.Context, miloClient miloClientInterface, rc *RelevantCommit) (string, *buildbot.Build) {
	buildURL, err := failedBuildFromCommitMessage(rc.CommitMessage)
	if err != nil {
		panic(err)
	}

	failedBuildInfo, err := miloClient.GetBuildInfo(ctx, buildURL)
	if err != nil {
		panic(err)
	}
	return buildURL, failedBuildInfo
}
