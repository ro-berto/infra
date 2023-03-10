// Copyright 2019 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build !windows
// +build !windows

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"golang.org/x/sys/unix"
)

// notifySIGTERM returns a context which is canceled when SIGTERM is
// received.
func notifySIGTERM(ctx context.Context) context.Context {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, unix.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case s := <-ch:
			log.Printf("Receiving signal %d", s)
			cancel()
		}
	}()
	return ctx
}
