// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tasks

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/maruel/subcommands"
	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/errors"
	"google.golang.org/api/option"

	"infra/cmd/mallet/internal/site"
	"infra/cmdsupport/cmdlib"
)

// EthernetHook does nothing, but eventually it will print all the ethernet events for a single
// day as a convenient list of protos.
var EthernetHook = &subcommands.Command{
	UsageLine: "ethernet-hook",
	ShortDesc: "ethernet-hook",
	LongDesc:  `ethernet-hook prints ethernet hook events for a single day.`,
	CommandRun: func() subcommands.CommandRun {
		c := &ethernetHookRun{}
		c.authFlags.Register(&c.Flags, site.DefaultAuthOptions)
		c.Flags.StringVar(&c.date, "date", "", "the date to process")
		c.Flags.StringVar(&c.bucket, "bucket", "chromeos-test-logs", "the base GS bucket to check for logs")
		return c
	},
}

type ethernetHookRun struct {
	subcommands.CommandRunBase
	authFlags authcli.Flags
	envFlags  site.EnvFlags

	// Dates are given in YYYY-MM-DD format, the only correct format.
	date string

	// Add a default bucket.
	bucket string
}

func (c *ethernetHookRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.innerRun(cli.GetContext(a, c, env), a, args, env); err != nil {
		cmdlib.PrintError(a, err)
		return 1
	}
	return 0
}

// Just print the projectID of one bucket.
func (c *ethernetHookRun) innerRun(ctx context.Context, a subcommands.Application, args []string, env subcommands.Env) error {
	options, err := c.authFlags.Options()
	if err != nil {
		return errors.Annotate(err, "failed to get auth options").Err()
	}
	authenticator := auth.NewAuthenticator(ctx, auth.SilentLogin, options)
	httpClient, err := authenticator.Client()
	if err != nil {
		return errors.Annotate(err, "failed to set up http client").Err()
	}
	email, err := authenticator.GetEmail()
	if err != nil {
		return errors.Annotate(err, "failed to get email").Err()
	}
	tokenSource, err := authenticator.TokenSource()
	if err != nil {
		return errors.Annotate(err, "failed to get token source").Err()
	}
	storageClient, err := storage.NewClient(
		ctx,
		option.WithHTTPClient(httpClient),
		option.WithTokenSource(tokenSource),
		option.WithScopes(options.Scopes...),
	)
	if err != nil {
		return errors.Annotate(err, "failed to set up storage client as %q", email).Err()
	}

	bucketAttrs, err := storageClient.Bucket(c.bucket).Attrs(ctx)
	if err != nil {
		return errors.Annotate(err, "get bucket properties for bucket %q as %q", c.bucket, email).Err()
	}

	_, err = fmt.Fprintf(a.GetErr(), "%d\n", bucketAttrs.ProjectNumber)
	if err != nil {
		return errors.Annotate(err, "printing").Err()
	}

	return nil
}