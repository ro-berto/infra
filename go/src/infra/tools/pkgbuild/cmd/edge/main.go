// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"infra/libs/cipkg"
	"infra/libs/cipkg/builtins"
	"infra/libs/cipkg/utilities"
	"infra/tools/pkgbuild/pkg/spec"
	"infra/tools/pkgbuild/pkg/stdenv"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
)

func main() {
	ctx := gologger.StdConfig.Use(context.Background())
	ctx = logging.SetLevel(ctx, logging.Error)
	storageDir, specDir := os.Args[1], os.Args[2]
	if err := stdenv.Init(); err != nil {
		log.Fatal(err)
	}
	s := spec.NewSpecLoader(specDir, "linux-amd64")
	g, err := s.FromSpec("curl")
	if err != nil {
		log.Fatal(err)
	}
	if err := build(ctx, storageDir, g); err != nil {
		log.Fatal(err)
	}
}
func build(ctx context.Context, path string, out cipkg.Generator) error {
	s, err := utilities.NewLocalStorage(path)
	if err != nil {
		return errors.Annotate(err, "failed to load storage").Err()
	}
	// Generate derivations
	bctx := &cipkg.BuildContext{
		Platform: cipkg.Platform{
			Build:  fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH),
			Host:   fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH),
			Target: fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH),
		},
		Storage: s,
		Context: ctx,
	}
	drv, meta, err := out.Generate(bctx)
	if err != nil {
		return errors.Annotate(err, "failed to generate venv derivation").Err()
	}
	pkg := s.Add(drv, meta)
	// Build derivations
	b := utilities.NewBuilder(s)
	if err := b.Add(pkg); err != nil {
		return errors.Annotate(err, "failed to add package to builder").Err()
	}
	var temp = filepath.Join(path, "temp")
	if err := os.RemoveAll(temp); err != nil {
		return err
	}
	if err := os.Mkdir(temp, os.ModePerm); err != nil {
		return err
	}
	if err := b.BuildAll(func(p cipkg.Package) error {
		id := p.Derivation().ID()
		d, err := os.MkdirTemp(temp, fmt.Sprintf("%s-", id))
		if err != nil {
			return err
		}
		var out strings.Builder
		cmd := utilities.CommandFromPackage(p)
		cmd.Stdout = &out
		cmd.Stderr = &out
		cmd.Dir = d
		if err := builtins.Execute(ctx, cmd); err != nil {
			logging.Errorf(ctx, "%s", out.String())
			return err
		}
		return nil
	}); err != nil {
		return errors.Annotate(err, "failed to build package").Err()
	}
	s.Prune(ctx, time.Hour*24, 256)
	return nil
}