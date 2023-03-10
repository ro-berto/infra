# Copyright 2021 The Chromium OS Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# NOTE: because the intended destination is a docker container running an
# alpine-based distribution, we must disable cgo when building.

.PHONY: satlab
satlab:
	@if test -f ./devscripts/pre-satlab 2>/dev/null; then ./devscripts/pre-satlab; fi
	env CGO_ENABLED=0 go build -o satlab -- .
	@if test -f ./devscripts/post-satlab 2>/dev/null; then ./devscripts/post-satlab; fi

.PHONY: test
test:
	env CGO_ENABLED=0 go test ./...

.PHONY: check
check: test

.PHONY: clean
clean:
	$(RM) ./satlab

.PHONY: vet
vet:
	env CGO_ENABLED=0 go vet ./...

.PHONY: fmt
fmt:
	env CGO_ENABLED=0 go fmt ./...

# This command works by marching forward the "prod" label to the version that the "latest" label
# points to. The "latest" label is automatically marched forward as part of the build process that
# happens post-submit.
.PHONY: deploy-to-prod
deploy-to-prod:
	@echo 'WARNING: This deploys to prod. Deploying in ten seconds. Press Ctrl-C to Cancel'
	@sleep 10.
	cipd set-ref 'infra/cros/cmd/satlab/' -ref prod -version latest
