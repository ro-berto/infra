#!/bin/bash
# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

set -e
set -x
set -o pipefail
shopt -s dotglob

PREFIX="$1"

# By default config will be written to home directory.
export CLOUDSDK_CONFIG="$(pwd)/.config"

# Install additional components. This will also install their dependencies.
#
# We assume here that "overall" gcloud SDK version is bumped whenever some of
# the dependencies change in a significant way. If a dependency changes without
# gcloud SDK version bump, 3pp won't notice this.
LINUX_EXTRAS=
if [ "$(uname -s)" == "Linux" ]; then
      LINUX_EXTRAS=cloud-spanner-emulator
fi
./google-cloud-sdk/bin/gcloud components install -q \
    alpha \
    beta \
    app-engine-go \
    app-engine-python \
    app-engine-python-extras \
    docker-credential-gcr \
    kubectl \
    gke-gcloud-auth-plugin \
    $LINUX_EXTRAS

# This is just a dead weight in the package, we won't rollback.
rm -rf ./google-cloud-sdk/.install/.backup

# Disable update checks, we deploy updates through CIPD.
./google-cloud-sdk/bin/gcloud config set --installation \
    component_manager/disable_update_check true

# No need to report usage from bots.
./google-cloud-sdk/bin/gcloud config set --installation \
    core/disable_usage_reporting true

# No need to survey bots.
./google-cloud-sdk/bin/gcloud config set --installation \
    survey/disable_prompts true

# Copy CHECKSUM to mitigate crbug/1365718#c14
cp ./google-cloud-sdk/platform/gsutil/CHECKSUM ./google-cloud-sdk/platform/gsutil/gslib/CHECKSUM

# No need to ~= double number of files in the package.
find ./google-cloud-sdk -name "*.pyc" -delete

# Put gcloud SDK root (including hidden files) at the root of the package.
mv ./google-cloud-sdk/* "$PREFIX"/
