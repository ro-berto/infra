# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

.PHONY: build-for-deployment test clean

build-for-deployment:
	GOOS=linux GOARCH=amd64 go build -o isolator

isolator:
	go build -o isolator

test: isolator
	./isolator -input=test -output=out

clean:
	rm -rf out isolator
