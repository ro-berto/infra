// Copyright 202 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import { createAsyncThunk, createSlice, PayloadAction } from '@reduxjs/toolkit'
import { RootState } from '../../app/store';

export interface LabState {
  status : string,
  labId : string,
  name : string
}

const initialState : LabState = {
  status: 'idle',
  labId: '',
  name: ''
}

// The function below is called a thunk and allows us to perform async logic. It
// can be dispatched like a regular action: `dispatch(fetchLabAsync(10))`. This
// will call the thunk with the `dispatch` function as the first argument. Async
// code can then be executed and other actions can be dispatched. Thunks are
// typically used to make async requests.
export const fetchLabAsync = createAsyncThunk(
  'lab/fetchLab',
  async (labId: string) => {
    const response = await fetchLab(labId);
    // The value we return becomes the `fulfilled` action payload
    return response.data;
  }
);

export const labSlice = createSlice({
  name: 'lab',
  initialState,
  reducers: {
    queryLab: (state, action: PayloadAction<string>) => {

    },
  },

  // The `extraReducers` field lets the slice handle actions generated by
  // createAsyncThunk or in other slices.
  extraReducers: (builder) => {
    builder
      .addCase(fetchLabAsync.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchLabAsync.fulfilled, (state, action) => {
        state.status = 'idle';
        state.name += action.payload;
      });
  },
})

// The function below is called a selector and allows us to select a value from
// the state. Selectors can also be defined inline where they're used instead of
// in the slice file. For example: `useSelector((state: RootState) => state.lab)`
export const selectLab = (state: RootState) => state.lab;

// A mock function to mimic making an async request for data, will be replaced
// with actual server call
function fetchLab(labId = '') {
  return new Promise<{ data: string }>((resolve) =>
    setTimeout(() => resolve({ data: labId}), 500)
  );
}

export const { queryLab } = labSlice.actions

export default labSlice.reducer