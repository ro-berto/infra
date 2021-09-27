// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package insert implements functions to insert rows for testing purposes.
package insert

import (
	"cloud.google.com/go/spanner"

	pb "infra/appengine/weetbix/proto/v1"
)

func updateDict(dest, source map[string]interface{}) {
	for k, v := range source {
		dest[k] = v
	}
}

// AnalyzedTestVariant returns a spanner mutation that inserts an analyzed test variant.
func AnalyzedTestVariant(realm, tId, vHash string, status pb.AnalyzedTestVariantStatus, extraValues map[string]interface{}) *spanner.Mutation {
	values := map[string]interface{}{
		"Realm":            realm,
		"TestId":           tId,
		"VariantHash":      vHash,
		"Status":           int64(status),
		"CreateTime":       spanner.CommitTimestamp,
		"StatusUpdateTime": spanner.CommitTimestamp,
	}
	updateDict(values, extraValues)
	return spanner.InsertMap("AnalyzedTestVariants", values)
}
