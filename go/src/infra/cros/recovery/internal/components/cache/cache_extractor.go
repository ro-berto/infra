// Copyright 2022 The Chromium OS Authors.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cache

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.chromium.org/luci/common/errors"

	"infra/cros/recovery/internal/components"
	"infra/cros/recovery/internal/log"
)

// ExtractRequest holds all data required to extract file from file on cache service.
type ExtractRequest struct {
	// URL to download the file from cache service.
	CacheFileURL string
	// Name of the file we wantt o extract from file.
	ExtractFileName string
	// Filepath of destination file.
	DestintionFilePath string
	// Download timeout.
	Timeout time.Duration
}

// Extract extract file from cache service by modifying URL to download the file.
func Extract(ctx context.Context, req *ExtractRequest, run components.Runner) error {
	// Path provided by TLS cannot be used for downloading and/or extracting the image file.
	// But we can utilize the address of caching service and apply some string manipulation to construct the URL that can be used for this.
	// Example: `http://Addr:8082/extract/chromeos-image-archive/board-release/R99-XXXXX.XX.0/chromiumos_test_image.tar.xz?file=chromiumos_test_image.bin`
	extractPath := strings.Replace(req.CacheFileURL, "/download/", "/extract/", 1)
	sourcePath := fmt.Sprintf("%s/chromiumos_test_image.tar.xz?file=%s", extractPath, req.ExtractFileName)
	if err := CurlFile(ctx, run, sourcePath, req.DestintionFilePath, req.Timeout); err != nil {
		return errors.Annotate(err, "extract from cache").Err()
	}
	return nil
}

// CurlFile downloads file by using curl util.
func CurlFile(ctx context.Context, run components.Runner, sourcePath, destinationPath string, timeout time.Duration) error {
	out, err := run(ctx, timeout, "curl", sourcePath, "--output", destinationPath)
	if err == nil {
		log.Debugf(ctx, "Successfully download %q from %q", destinationPath, sourcePath)
		return nil
	}
	log.Debugf(ctx, "Fail to download %q from %q", destinationPath, sourcePath)
	log.Debugf(ctx, "Fail to download %q: output %s", destinationPath, out)
	return errors.Annotate(err, "install firmware image").Err()
}