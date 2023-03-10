// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build !windows
// +build !windows

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sys/unix"
)

func cancelOnSignals(ctx context.Context, idleConns chan struct{}, svr *http.Server, gracePeriod time.Duration) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, unix.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("Caught signal: %s. Gracefully shutting down archive-server", sig)
		ctx, cancel := context.WithTimeout(ctx, gracePeriod)
		defer cancel()
		if err := svr.Shutdown(ctx); err != nil {
			log.Printf("archive-server shutdown unsuccesfully: %v", err)
		} else {
			log.Printf("archive-server shutdown successfully!")
		}
		close(idleConns)
	}()
	return ctx
}
