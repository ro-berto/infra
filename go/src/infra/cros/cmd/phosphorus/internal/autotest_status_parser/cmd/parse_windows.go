// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build windows
// +build windows

package cmd

func exitedWithErrors(content string) bool {
	panic("not supported on windows")
}
