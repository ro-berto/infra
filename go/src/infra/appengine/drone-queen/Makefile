# Copyright 2019 The LUCI Authors. All rights reserved.
# Use of this source code is governed under the Apache License, Version 2.0
# that can be found in the LICENSE file.

.PHONY: default
default: vet test

.PHONY: vet
vet:
	go vet -all ./...

.PHONY: test
test:
	go test ./...

.PHONY: generate
generate:
	go generate ./...

APP_DIR = appengine

# Uploading

GCP_PROJECT_DEV = drone-queen-dev

.PHONY: up-dev
up-dev:
	gae.py upload -A $(GCP_PROJECT_DEV) --app-dir $(APP_DIR) -f
	gae.py switch -A $(GCP_PROJECT_DEV) --app-dir $(APP_DIR) -f

.PHONY: cleanup-dev
cleanup-dev:
	gae.py cleanup -A $(GCP_PROJECT_DEV) --app-dir $(APP_DIR)

GCP_PROJECT_PROD = drone-queen-prod

.PHONY: up-prod
up-prod:
	gae.py upload -A $(GCP_PROJECT_PROD) --app-dir $(APP_DIR) -f
	gae.py switch -A $(GCP_PROJECT_PROD) --app-dir $(APP_DIR) -f

.PHONY: cleanup-prod
cleanup-prod:
	gae.py cleanup -A $(GCP_PROJECT_PROD) --app-dir $(APP_DIR)
