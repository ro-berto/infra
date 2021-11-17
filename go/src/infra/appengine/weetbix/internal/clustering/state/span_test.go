// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package state

import (
	"context"
	"sort"
	"strings"
	"testing"
	"time"

	"go.chromium.org/luci/server/span"

	"infra/appengine/weetbix/internal/testutil"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestSpanner(t *testing.T) {
	Convey(`With Spanner Test Database`, t, func() {
		ctx := testutil.SpannerTestContext(t)
		Convey(`Create`, func() {
			testCreate := func(e *Entry) error {
				_, err := span.ReadWriteTransaction(ctx, func(ctx context.Context) error {
					return Create(ctx, e)
				})
				return err
			}
			e := NewEntry(100).Build()
			Convey(`Valid`, func() {
				err := testCreate(e)
				So(err, ShouldBeNil)

				txn := span.Single(ctx)
				actual, err := Read(txn, e.Project, e.ChunkID)
				So(err, ShouldBeNil)

				// Check the LastUpdated time is set, but ignore it for
				// further comparisons.
				clearLastUpdatedTimestamps(actual)

				So(err, ShouldBeNil)
				So(actual, ShouldResemble, e)
			})
			Convey(`Invalid`, func() {
				Convey(`Project missing`, func() {
					e.Project = ""
					err := testCreate(e)
					So(err, ShouldErrLike, `project "" is not valid`)
				})
				Convey(`Chunk ID missing`, func() {
					e.ChunkID = ""
					err := testCreate(e)
					So(err, ShouldErrLike, `chunk ID "" is not valid`)
				})
				Convey(`Partition Time missing`, func() {
					var t time.Time
					e.PartitionTime = t
					err := testCreate(e)
					So(err, ShouldErrLike, "partition time must be specified")
				})
				Convey(`Object ID missing`, func() {
					e.ObjectID = ""
					err := testCreate(e)
					So(err, ShouldErrLike, "object ID must be specified")
				})
				Convey(`Rules Version missing`, func() {
					var t time.Time
					e.Clustering.RulesVersion = t
					err := testCreate(e)
					So(err, ShouldErrLike, "rules version must be specified")
				})
				Convey(`Algorithms Version missing`, func() {
					e.Clustering.AlgorithmsVersion = 0
					err := testCreate(e)
					So(err, ShouldErrLike, "algorithms version must be specified")
				})
				Convey(`Clusters missing`, func() {
					e.Clustering.Clusters = nil
					err := testCreate(e)
					So(err, ShouldErrLike, "there must be clustered test results in the chunk")
				})
				Convey(`Algorithms invalid`, func() {
					Convey(`Empty algorithm`, func() {
						e.Clustering.Algorithms[""] = struct{}{}
						err := testCreate(e)
						So(err, ShouldErrLike, `algorithm "" is not valid`)
					})
					Convey("Algorithm invalid", func() {
						e.Clustering.Algorithms["!!!"] = struct{}{}
						err := testCreate(e)
						So(err, ShouldErrLike, `algorithm "!!!" is not valid`)
					})
				})
				Convey(`Clusters invalid`, func() {
					Convey(`Algorithm missing`, func() {
						e.Clustering.Clusters[1][1].Algorithm = ""
						err := testCreate(e)
						So(err, ShouldErrLike, `clusters: test result 1: cluster 1: cluster ID is not valid: algorithm not valid`)
					})
					Convey("Algorithm invalid", func() {
						e.Clustering.Clusters[1][1].Algorithm = "!!!"
						err := testCreate(e)
						So(err, ShouldErrLike, `clusters: test result 1: cluster 1: cluster ID is not valid: algorithm not valid`)
					})
					Convey("Algorithm not in algorithms set", func() {
						e.Clustering.Algorithms = map[string]struct{}{
							"alg-extra": {},
						}
						err := testCreate(e)
						So(err, ShouldErrLike, `a test result was clustered with an unregistered algorithm`)
					})
					Convey("ID missing", func() {
						e.Clustering.Clusters[1][1].ID = ""
						err := testCreate(e)
						So(err, ShouldErrLike, `clusters: test result 1: cluster 1: cluster ID is not valid: ID is empty`)
					})
				})
			})
		})
		Convey(`ReadNextN`, func() {
			targetRulesVersion := time.Date(2024, 1, 1, 1, 1, 1, 0, time.UTC)
			targetAlgorithmsVersion := 10
			entries := []*Entry{
				// Should not be read.
				NewEntry(0).WithChunkIDPrefix("11").WithAlgorithmsVersion(10).WithRulesVersion(targetRulesVersion).Build(),

				// Should be read (rulesVersion < targetRulesVersion).
				NewEntry(1).WithChunkIDPrefix("11").WithAlgorithmsVersion(10).WithRulesVersion(targetRulesVersion.Add(-1 * time.Hour)).Build(), // Should be read.
				NewEntry(3).WithChunkIDPrefix("11").WithRulesVersion(targetRulesVersion.Add(-1 * time.Hour)).Build(),

				// Should be read (algorithmsVersion < targetAlgorithmsVersion).
				NewEntry(2).WithChunkIDPrefix("11").WithAlgorithmsVersion(9).WithRulesVersion(targetRulesVersion).Build(),
				NewEntry(4).WithChunkIDPrefix("11").WithAlgorithmsVersion(2).Build(),

				// Should not be read (other project).
				NewEntry(5).WithChunkIDPrefix("11").WithAlgorithmsVersion(2).WithProject("other").Build(),

				// Check handling of EndChunkID as an inclusive upper-bound.
				NewEntry(6).WithChunkIDPrefix("11" + strings.Repeat("ff", 15)).WithAlgorithmsVersion(2).Build(), // Should be read.
				NewEntry(7).WithChunkIDPrefix("12" + strings.Repeat("00", 15)).WithAlgorithmsVersion(2).Build(), // Should not be read.
			}

			err := CreateEntriesForTesting(ctx, entries)
			So(err, ShouldBeNil)

			expectedEntries := []*Entry{
				entries[1],
				entries[2],
				entries[3],
				entries[4],
				entries[6],
			}
			sort.Slice(expectedEntries, func(i, j int) bool {
				return expectedEntries[i].ChunkID < expectedEntries[j].ChunkID
			})

			readOpts := ReadNextOptions{
				StartChunkID:      "11" + strings.Repeat("00", 15),
				EndChunkID:        "11" + strings.Repeat("ff", 15),
				AlgorithmsVersion: int64(targetAlgorithmsVersion),
				RulesVersion:      targetRulesVersion,
			}
			// Reads first page.
			rows, err := ReadNextN(span.Single(ctx), testProject, readOpts, 3)
			So(err, ShouldBeNil)
			clearLastUpdatedTimestamps(rows...)
			So(rows, ShouldResemble, expectedEntries[0:3])

			// Read second page.
			readOpts.StartChunkID = rows[2].ChunkID
			rows, err = ReadNextN(span.Single(ctx), testProject, readOpts, 3)
			So(err, ShouldBeNil)
			clearLastUpdatedTimestamps(rows...)
			So(rows, ShouldResemble, expectedEntries[3:])

			// Read empty last page.
			readOpts.StartChunkID = rows[1].ChunkID
			rows, err = ReadNextN(span.Single(ctx), testProject, readOpts, 3)
			So(err, ShouldBeNil)
			So(rows, ShouldBeEmpty)
		})
		Convey(`EstimateChunks`, func() {
			Convey(`Less than 100 chunks`, func() {
				est, err := EstimateChunks(span.Single(ctx), testProject)
				So(err, ShouldBeNil)
				So(est, ShouldBeLessThan, 100)
			})
			Convey(`At least 100 chunks`, func() {
				var entries []*Entry
				for i := 0; i < 200; i++ {
					entries = append(entries, NewEntry(i).Build())
				}
				err := CreateEntriesForTesting(ctx, entries)
				So(err, ShouldBeNil)

				count, err := EstimateChunks(span.Single(ctx), testProject)
				So(err, ShouldBeNil)
				So(count, ShouldBeGreaterThan, 190)
				So(count, ShouldBeLessThan, 210)
			})
		})
	})
	Convey(`estimateChunksFromID`, t, func() {
		// Extremely full table. This is the minimum that the 100th ID
		// could be (considering 0x63 = 99).
		count, err := estimateChunksFromID("00000000000000000000000000000063")
		So(err, ShouldBeNil)
		// The maximum estimate.
		So(count, ShouldEqual, 1000*1000*1000)

		// The 100th ID is right in the middle of the keyspace.
		count, err = estimateChunksFromID("7fffffffffffffffffffffffffffffff")
		So(err, ShouldBeNil)
		So(count, ShouldEqual, 200)

		// The 100th ID is right at the end of the keyspace.
		count, err = estimateChunksFromID("ffffffffffffffffffffffffffffffff")
		So(err, ShouldBeNil)
		So(count, ShouldEqual, 100)
	})
}

func clearLastUpdatedTimestamps(entries ...*Entry) {
	for _, e := range entries {
		// Check the LastUpdated time is set, but ignore it for
		// further comparisons.
		So(e.LastUpdated, ShouldNotBeZeroValue)
		e.LastUpdated = time.Time{}
	}
}
