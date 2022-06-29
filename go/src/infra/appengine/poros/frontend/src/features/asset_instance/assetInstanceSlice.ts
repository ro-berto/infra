// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import {
  IAssetInstanceService,
  AssetInstanceService,
  ListAssetInstancesRequest,
  AssetInstanceModel,
  UpdateAssetInstanceRequest,
} from '../../api/asset_instance_service';
import { AssetModel } from '../../api/asset_service';
import { RootState } from '../../app/store';
import { queryAssetAsync } from '../asset/assetSlice';
export interface AssetInstanceState {
  assetInstances: AssetInstanceModel[];
  pageToken: string | undefined;
  pageNumber: number;
  pageSize: number;
  fetchStatus: string;
  assets: AssetModel[];
  record: AssetInstanceModel;
  savingStatus: string;
  deleteAtBuffer: string;
}
const initialState: AssetInstanceState = {
  assetInstances: [],
  pageToken: undefined,
  pageNumber: 1,
  pageSize: 10,
  fetchStatus: 'idle',
  assets: [],
  record: AssetInstanceModel.defaultEntity(),
  savingStatus: 'idle',
  deleteAtBuffer: '',
};
// The function below is called a thunk and allows us to perform async logic. It
// can be dispatched like a regular action: `dispatch(queryAssetInstanceAsync())`. This
// will call the thunk with the `dispatch` function as the first argument. Async
// code can then be executed and other actions can be dispatched. Thunks are
// typically used to make async requests.
export const queryAssetInstanceAsync = createAsyncThunk(
  'assetInstance/queryAssetInstance',
  async ({ pageSize, pageToken }: { pageSize: number; pageToken: string }) => {
    const request: ListAssetInstancesRequest = {
      pageSize: pageSize,
      pageToken: pageToken,
      readMask: undefined,
    };
    const service: IAssetInstanceService = new AssetInstanceService();
    const response = await service.list(request);
    return response;
  }
);
export const updateAssetInstanceAsync = createAsyncThunk(
  'assetInstance/updateAssetInstance',
  async ({
    assetInstance,
    updateMask,
  }: {
    assetInstance: AssetInstanceModel;
    updateMask: string[];
  }) => {
    const request: UpdateAssetInstanceRequest = {
      assetInstance: assetInstance,
      updateMask: updateMask,
    };
    const service: IAssetInstanceService = new AssetInstanceService();
    const response = await service.update(request);
    return response;
  }
);
export const assetInstanceSlice = createSlice({
  name: 'assetInstance',
  initialState,
  reducers: {
    setDeleteTime: (state, action) => {
      state.deleteAtBuffer = action.payload;
      if (!isNaN(Date.parse(state.deleteAtBuffer))) {
        state.record.deleteAt = new Date(state.deleteAtBuffer);
      }
    },
    onSelectRecord: (state, action) => {
      state.record = state.assetInstances.filter(
        (s) => s.assetInstanceId == action.payload.assetInstanceId
      )[0];
      if (state.record.deleteAt === undefined) {
        state.deleteAtBuffer = '';
      } else {
        state.deleteAtBuffer = new Date(
          state.record.deleteAt.getTime() -
            60000 * new Date().getTimezoneOffset()
        )
          .toISOString()
          .slice(0, 16);
      }
    },
  },
  // The `extraReducers` field lets the slice handle actions generated by
  // createAsyncThunk or in other slices.
  extraReducers: (builder) => {
    builder
      .addCase(queryAssetInstanceAsync.pending, (state) => {
        state.fetchStatus = 'loading';
      })
      .addCase(queryAssetInstanceAsync.fulfilled, (state, action) => {
        state.fetchStatus = 'idle';
        state.assetInstances = action.payload.assetInstances;
        state.pageToken = action.payload.nextPageToken;
        state.assetInstances.forEach(function (assetInstance) {
          if (
            assetInstance.deleteAt?.toUTCString() ===
            'Mon, 01 Jan 0001 00:00:00 GMT'
          ) {
            assetInstance.deleteAt = undefined;
          }
        });
      })
      .addCase(queryAssetAsync.fulfilled, (state, action) => {
        state.assets = action.payload.assets;
      })
      .addCase(updateAssetInstanceAsync.pending, (state) => {
        state.savingStatus = 'loading';
      })
      .addCase(updateAssetInstanceAsync.fulfilled, (state, action) => {
        state.savingStatus = 'idle';
        state.record = action.payload;
      });
  },
});
// The function below is called a selector and allows us to select a value from
// the state. Selectors can also be defined inline where they're used instead of
// in the slice file. For example: `useSelector((state: RootState) => state.assetInstance)`
export const selectAssetInstanceState = (state: RootState) =>
  state.assetInstance;
export const { setDeleteTime, onSelectRecord } = assetInstanceSlice.actions;
export default assetInstanceSlice.reducer;