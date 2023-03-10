# Copyright 2021 The Chromium OS Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

GOTESTFLAGS ?= -count 20

.PHONY: default
default: build


# In order to keep the output readable, run tests with count 1 BEFORE running them
# with a non-1 count.
#
# We still want to run the tests with a non-1 count in a later step in order to catch
# failures due to nondeterminism (from iterating over maps).
.PHONY:
build:
	go vet ./...
	go test ./...
	go test -race ./...
	go test $(GOTESTFLAGS) ./...
	go test -race $(GOTESTFLAGS) ./...
	go build -o ./cmd/karte/karte             ./cmd/karte
	go build -o ./cmd/karteserver/karteserver ./cmd/karteserver

.PHONY: check
check:
	go vet ./...
	go test ./...

.PHONY: test
test:
	go test ./...

.PHONY: gen
gen:
	go generate ./...

# clean just the binary artifacts
.PHONY: clean
clean:
	$(RM) ./cmd/karte/karte ./cmd/karteserver/karteserver ./karte ./karteserver

# Add a target for removing the generated files.
# It is easier to keep the generated files in sync if there's
# a straightforward way to remove them.
.PHONY: realclean
realclean: clean
	find . -name '*.pb.go' -delete
	$(RM) ./api/pb.discovery.go

.PHONY: run
run:
	./cmd/karteserver/karteserver -auth-service-host chrome-infra-auth-dev.appspot.com

# For the rules that update the schema, those are part of every deploy action (for the relevant target)
# because forgetting to update the schema can cause loss of data. This loss is not permanent, but recovering
# from it is time-consuming.
.PHONY: update-schema-dev
update-schema-dev:
	./scripts/push_bq_schema -p dev

.PHONY: update-schema-prod
update-schema-prod:
	./scripts/push_bq_schema -p prod

# Intentionally use the gae.py version that is present on the path. If no such
# command is available on the path, you are not in the correct environment.
.PHONY: upload-dev
upload-dev: update-schema-dev
	@echo BEGIN UPLOAD DEV
	gae.py upload --app-id=chrome-fleet-karte-dev --app-dir=./cmd/karteserver
	@echo END UPLOAD DEV

.PHONY: switch-dev
switch-dev: update-schema-dev
	@echo BEGIN SWITCH DEV
	gae.py switch --app-id=chrome-fleet-karte-dev --app-dir=./cmd/karteserver
	@echo END SWITCH DEV

.PHONY: upload-switch-dev
upload-switch-dev: update-schema-dev
	gae.py upload --app-id=chrome-fleet-karte-dev --app-dir=./cmd/karteserver --force --verbose
	gae.py switch --app-id=chrome-fleet-karte-dev --app-dir=./cmd/karteserver --force --verbose

# For convenience, upload prod and switch to it.
# Do NOT use --force for prod.
.PHONY: upload-switch-prod
upload-switch-prod: update-schema-dev update-schema-prod
	gae.py upload --app-id=chrome-fleet-karte --app-dir=./cmd/karteserver --verbose
	gae.py switch --app-id=chrome-fleet-karte --app-dir=./cmd/karteserver --verbose

# Intentionally use the gae.py version that is present on the path. If no such
# command is available on the path, you are not in the correct environment.
.PHONY: upload-prod
upload-prod: update-schema-dev update-schema-prod
	@echo BEGIN UPLOAD PROD
	gae.py upload --app-id=chrome-fleet-karte --app-dir=./cmd/karteserver
	@echo END UPLOAD PROD
	@echo "If you did NOT intend to push prod, do NOT call \"switch prod\""

.PHONY: switch-prod
switch-prod: update-schema-dev update-schema-prod
	@echo BEGIN SWITCH PROD
	gae.py switch --app-id=chrome-fleet-karte --app-dir=./cmd/karteserver
	@echo END SWITCH PROD
