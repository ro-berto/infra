// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import {
  CreateResourceRequest,
  DeleteResourceRequest,
  GetResourceRequest,
  IResourceService,
  ResourceService,
  ListResourcesRequest,
  ResourceModel,
} from '../../api/resource_service';
import { RootState } from '../../app/store';

export interface ResourceState {
  resources: ResourceModel[];
  pageToken: string | undefined;
  record: ResourceModel;
  fetchStatus: string;
  savingStatus: string;
  deletingStatus: string;
  pageNumber: number;
  pageSize: number;
}

const initialState: ResourceState = {
  resources: [],
  pageToken: undefined,
  pageNumber: 1,
  pageSize: 25,
  fetchStatus: 'idle',
  record: ResourceModel.defaultEntity(),
  savingStatus: 'idle',
  deletingStatus: 'idle',
};

// The function below is called a thunk and allows us to perform async logic. It
// can be dispatched like a regular action: `dispatch(fetchResourceAsync(10))`. This
// will call the thunk with the `dispatch` function as the first argument. Async
// code can then be executed and other actions can be dispatched. Thunks are
// typically used to make async requests.
export const fetchResourceAsync = createAsyncThunk(
  'resource/fetchResource',
  async (resourceId: string) => {
    const request: GetResourceRequest = {
      id: resourceId,
    };
    const service: IResourceService = new ResourceService();
    const response = await service.get(request);
    // The value we return becomes the `fulfilled` action payload
    return response;
  }
);

export const createResourceAsync = createAsyncThunk(
  'resource/createResource',
  async ({
    name,
    type,
    description,
    machineInfo,
    domainInfo,
  }: {
    name: string;
    type: string;
    description: string;
    machineInfo: string;
    domainInfo: string;
  }) => {
    const request: CreateResourceRequest = {
      name,
      type,
      description,
      machineInfo,
      domainInfo,
    };
    const service: IResourceService = new ResourceService();
    const response = await service.create(request);
    return response;
  }
);

export const queryResourceAsync = createAsyncThunk(
  'resource/queryResource',
  async ({ pageSize, pageToken }: { pageSize: number; pageToken: string }) => {
    const request: ListResourcesRequest = {
      pageSize: pageSize,
      pageToken: pageToken,
      readMask: undefined,
    };
    const service: IResourceService = new ResourceService();
    const response = await service.list(request);
    return response;
  }
);

export const deleteResourceAsync = createAsyncThunk(
  'resource/deleteResource',
  async (resourceId: string) => {
    const request: DeleteResourceRequest = {
      id: resourceId,
    };
    const service: IResourceService = new ResourceService();
    const response = await service.get(request);
    return response;
  }
);

export const resourceSlice = createSlice({
  name: 'resource',
  initialState,
  reducers: {
    setPageSize: (state, action) => {
      state.pageSize = action.payload.pageSize;
    },
    setName: (state, action) => {
      state.record.name = action.payload;
    },
    setType: (state, action) => {
      state.record.type = action.payload;
    },
    setDescription: (state, action) => {
      state.record.description = action.payload;
    },
    setMachineInfo: (state, action) => {
      state.record.machineInfo = action.payload;
    },
    setDomainInfo: (state, action) => {
      state.record.domainInfo = action.payload;
    },
    onSelectRecord: (state, action) => {
      state.record = state.resources.filter(
        (s) => s.resourceId == action.payload.resourceId
      )[0];
    },
    clearSelectedRecord: (state) => {
      state.record = ResourceModel.defaultEntity();
    },
  },

  // The `extraReducers` field lets the slice handle actions generated by
  // createAsyncThunk or in other slices.
  extraReducers: (builder) => {
    builder
      .addCase(fetchResourceAsync.pending, (state) => {
        state.fetchStatus = 'loading';
      })
      .addCase(fetchResourceAsync.fulfilled, (state, action) => {
        state.fetchStatus = 'idle';
        state.record = action.payload;
      })
      .addCase(createResourceAsync.pending, (state) => {
        state.savingStatus = 'loading';
      })
      .addCase(createResourceAsync.fulfilled, (state, action) => {
        state.savingStatus = 'idle';
        state.record = action.payload;
      })
      .addCase(queryResourceAsync.pending, (state) => {
        state.fetchStatus = 'loading';
      })
      .addCase(queryResourceAsync.fulfilled, (state, action) => {
        state.fetchStatus = 'idle';
        state.resources = action.payload.resources;
        state.pageToken = action.payload.nextPageToken;
      })
      .addCase(deleteResourceAsync.pending, (state) => {
        state.deletingStatus = 'loading';
      })
      .addCase(deleteResourceAsync.fulfilled, (state) => {
        state.deletingStatus = 'idle';
        state.record = ResourceModel.defaultEntity();
      });
  },
});

// The function below is called a selector and allows us to select a value from
// the state. Selectors can also be defined inline where they're used instead of
// in the slice file. For example: `useSelector((state: RootState) => state.resource)`
export const selectResourceState = (state: RootState) => state.resource;

export const {
  setPageSize,
  onSelectRecord,
  clearSelectedRecord,
  setName,
  setType,
  setDescription,
  setMachineInfo,
  setDomainInfo,
} = resourceSlice.actions;

export default resourceSlice.reducer;