// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"

	"cloud.google.com/go/bigquery"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/gaeemulation"
	"go.chromium.org/luci/server/limiter"
	"go.chromium.org/luci/server/module"

	"infra/unifiedfleet/app/config"
	"infra/unifiedfleet/app/dumper"
	"infra/unifiedfleet/app/external"
	"infra/unifiedfleet/app/util"
)

func main() {
	// skip the realms check for dumper service
	util.SkipRealmsCheck = true
	modules := []module.Module{
		gaeemulation.NewModuleFromFlags(),
		limiter.NewModuleFromFlags(),
	}

	cfgLoader := config.Loader{}
	cfgLoader.RegisterFlags(flag.CommandLine)

	server.Main(nil, modules, func(srv *server.Server) error {
		// Load service config form a local file (deployed via GKE),
		// periodically reread it to pick up changes without full restart.
		if _, err := cfgLoader.Load(); err != nil {
			return err
		}
		srv.RunInBackground("ufs.config", cfgLoader.ReloadLoop)
		srv.Context = config.Use(srv.Context, cfgLoader.Config())
		srv.Context = external.WithServerInterface(srv.Context)

		client, err := bigquery.NewClient(srv.Context, srv.Options.CloudProject)
		if err != nil {
			return err
		}
		srv.Context = dumper.Use(srv.Context, client)
		srv.Context = dumper.UseProject(srv.Context, srv.Options.CloudProject)
		dumper.InstallCronServices(srv)
		dumper.InitServer(srv)
		return nil
	})
}
