// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package configuration

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.chromium.org/luci/appengine/gaetesting"
	"go.chromium.org/luci/gae/service/datastore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddPublicBoardData(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContextWithAppID("go-test")
	datastore.GetTestable(ctx).Consistent(true)

	t.Run("add non-existent Public Board", func(t *testing.T) {
		const expectedBoardName = "board1"
		expectedModels := []string{"model1", "model2"}
		got, err := AddPublicBoardModelData(ctx, expectedBoardName, expectedModels, true)
		if err != nil {
			t.Fatalf("AddPublicBoardModelData failed: %s", err)
		}
		if got.Board != expectedBoardName {
			t.Errorf("AddPublicBoardModelData returned unexpected Board:\n%s", got.Board)
		}
		if got.BoardHasPrivateModels != true {
			t.Errorf("AddPublicBoardModelData returned unexpected result for BoardHasPrivateModels:\n%v", got.BoardHasPrivateModels)
		}
		if diff := cmp.Diff(expectedModels, got.Models); diff != "" {
			t.Errorf("AddPublicBoardModelData returned unexpected models (-want +got):\n%s", diff)
		}
	})

	t.Run("add existing Board", func(t *testing.T) {
		const expectedBoardName = "board1"
		expectedModels := []string{"model1", "model2"}

		// Insert board1 into datastore
		AddPublicBoardModelData(ctx, expectedBoardName, expectedModels, false)

		// Update board1
		got, err := AddPublicBoardModelData(ctx, expectedBoardName, expectedModels, false)
		if err != nil {
			t.Fatalf("AddPublicBoardModelData failed: %s", err)
		}
		if got.Board != expectedBoardName {
			t.Errorf("AddPublicBoardModelData returned unexpected Board:\n%s", got.Board)
		}
		if diff := cmp.Diff(expectedModels, got.Models); diff != "" {
			t.Errorf("AddPublicBoardModelData returned unexpected models (-want +got):\n%s", diff)
		}
	})

	t.Run("add empty board", func(t *testing.T) {
		_, err := AddPublicBoardModelData(ctx, "", []string{}, false)
		if err == nil {
			t.Errorf("AddPublicBoardModelData succeeded with empty Board name")
		}
		if c := status.Code(err); c != codes.Internal {
			t.Errorf("Unexpected error when calling AddPublicBoardModelData: %s", err)
		}
	})
}

func TestGetPublicBoardModelData(t *testing.T) {
	t.Parallel()
	ctx := gaetesting.TestingContextWithAppID("go-test")
	datastore.GetTestable(ctx).Consistent(true)

	t.Run("get PublicBoardModelData by existing ID", func(t *testing.T) {
		const expectedBoardName = "board1"
		expectedModels := []string{"model1", "model2"}
		_, err := AddPublicBoardModelData(ctx, expectedBoardName, expectedModels, true)
		if err != nil {
			t.Fatalf("AddPublicBoardModelData failed: %s", err)
		}

		got, err := GetPublicBoardModelData(ctx, expectedBoardName)
		if err != nil {
			t.Fatalf("GetPublicBoardModelData failed: %s", err)
		}
		if got.Board != expectedBoardName {
			t.Errorf("GetPublicBoardModelData returned unexpected Board:\n%s", got.Board)
		}
		if got.BoardHasPrivateModels != true {
			t.Errorf("GetPublicBoardModelData returned unexpected result for BoardHasPrivateModels:\n%v", got.BoardHasPrivateModels)
		}
		if diff := cmp.Diff(expectedModels, got.Models); diff != "" {
			t.Errorf("GetPublicBoardModelData returned unexpected models (-want +got):\n%s", diff)
		}
	})

	t.Run("get PublicBoardModelData for previous data model with no boolean for tracking private models", func(t *testing.T) {
		const expectedBoardName = "board1"
		expectedModels := []string{"model1", "model2"}
		entity := &PublicBoardModelDataEntity{
			Board:  expectedBoardName,
			Models: expectedModels,
		}
		err := datastore.Put(ctx, entity)
		if err != nil {
			t.Fatalf("adding datastore entity PublicBoardModelDataEntity failed : %s", err)
		}

		got, err := GetPublicBoardModelData(ctx, expectedBoardName)
		if err != nil {
			t.Fatalf("GetPublicBoardModelData failed: %s", err)
		}
		if got.Board != expectedBoardName {
			t.Errorf("GetPublicBoardModelData returned unexpected Board:\n%s", got.Board)
		}
		if got.BoardHasPrivateModels != false {
			t.Errorf("GetPublicBoardModelData returned unexpected result for BoardHasPrivateModels:\n%v", got.BoardHasPrivateModels)
		}
		if diff := cmp.Diff(expectedModels, got.Models); diff != "" {
			t.Errorf("GetPublicBoardModelData returned unexpected models (-want +got):\n%s", diff)
		}
	})

	t.Run("get PublicBoardModelData by non-existent ID", func(t *testing.T) {
		const expectedBoardName = "board2"
		_, err := GetPublicBoardModelData(ctx, expectedBoardName)
		if err == nil {
			t.Errorf("GetPublicBoardModelData succeeded with non-existent ID: %s", expectedBoardName)
		}
		if c := status.Code(err); c != codes.NotFound {
			t.Errorf("Unexpected error when calling GetPublicBoardModelData: %s", err)
		}
	})
}
