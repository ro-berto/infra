// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import React from 'react';
import RefreshIcon from '@mui/icons-material/Refresh';
import DeleteIcon from '@mui/icons-material/Delete';
import {
  Select,
  TextField,
  Grid,
  Box,
  Typography,
  MenuItem,
  Button,
  Stack,
  InputLabel,
  FormControl,
} from '@mui/material';
import {
  clearSelectedRecord,
  createResourceAsync,
  setName,
  setType,
  setDescription,
  setMachineInfo,
  setDomainInfo,
} from './resourceSlice';

import { useAppSelector, useAppDispatch } from '../../app/hooks';

export const Resource = () => {
  const [activeResourceType, setActiveResourceType] = React.useState('machine');
  const name: string = useAppSelector((state) => state.resource.record.name);
  const type: string = useAppSelector((state) => state.resource.record.type);
  const description: string = useAppSelector(
    (state) => state.resource.record.description
  );
  const machineInfo: string = useAppSelector(
    (state) => state.resource.record.machineInfo
  );
  const domainInfo: string = useAppSelector(
    (state) => state.resource.record.domainInfo
  );
  const resourceId: string = useAppSelector(
    (state) => state.resource.record.resourceId
  );
  const dispatch = useAppDispatch();

  // Event Handlers
  const handleSaveClick = (
    name: string,
    type: string,
    description: string,
    machineInfo: string,
    domainInfo: string,
    resourceId: string
  ) => {
    if (resourceId === '') {
      dispatch(
        createResourceAsync({
          name,
          type,
          description,
          machineInfo,
          domainInfo,
        })
      );
    }
  };

  const handleCancelClick = () => {
    dispatch(clearSelectedRecord());
  };

  // Render functions

  // This function will be used once we give user the ability to select type of Resource
  const renderTypeDropdown = () => {
    return (
      <Grid container spacing={2} padding={1} paddingTop={6}>
        <Grid item xs={12}>
          <FormControl variant="standard" fullWidth>
            <InputLabel>Type</InputLabel>
            <Select
              label="Type"
              id="type"
              defaultValue="machine"
              value={type}
              onChange={(e) => {
                setActiveResourceType(e.target.value);
                dispatch(setType(e.target.value));
              }}
              fullWidth
              inputProps={{ fullWidth: true }}
              variant="standard"
              placeholder="Type"
            >
              <MenuItem value={'machine'}>Machine</MenuItem>
              <MenuItem value={'domain'}>Domain</MenuItem>
            </Select>
          </FormControl>
        </Grid>
      </Grid>
    );
  };

  const renderMachineMetaDropdown = () => {
    return (
      <Grid container spacing={2} padding={1} paddingTop={6}>
        <Grid item xs={12}>
          <FormControl variant="standard" fullWidth>
            <InputLabel>VM Images</InputLabel>
            <Select
              id="machininfo"
              value={machineInfo}
              onChange={(e) => dispatch(setMachineInfo(e.target.value))}
              fullWidth
              inputProps={{ fullWidth: true }}
              variant="standard"
              placeholder="Type"
            >
              <MenuItem value={'image-1'}>Image 1</MenuItem>
              <MenuItem value={'image-2'}>Image 2</MenuItem>
              <MenuItem value={'image-3'}>Image 3</MenuItem>
              <MenuItem value={'image-4'}>Image 4</MenuItem>
              <MenuItem value={'image-5'}>Image 5</MenuItem>
            </Select>
          </FormControl>
        </Grid>
      </Grid>
    );
  };

  const renderDomainMetaInput = () => {
    return (
      <Grid container spacing={2} padding={1} paddingTop={6}>
        <Grid item xs={12}>
          <TextField
            id="domainInfo"
            label="Domain Information"
            multiline
            rows={4}
            variant="standard"
            onChange={(e) => dispatch(setDomainInfo(e.target.value))}
            value={domainInfo}
            fullWidth
            InputProps={{ fullWidth: true }}
          />
        </Grid>
      </Grid>
    );
  };

  return (
    <Box
      sx={{
        width: 720,
        maxWidth: '100%',
        padding: 1,
      }}
    >
      <Grid container spacing={2} padding={1}>
        <Grid
          item
          style={{
            display: 'flex',
            justifyContent: 'flex-start',
            alignItems: 'center',
          }}
          xs={8}
        >
          <Typography variant="h5">Resource</Typography>
        </Grid>
      </Grid>
      <Grid container spacing={2} padding={1}>
        <Grid item xs={12}>
          <TextField
            label="Name"
            id="name"
            value={name}
            onChange={(e) => dispatch(setName(e.target.value))}
            fullWidth
            InputProps={{ fullWidth: true }}
            variant="standard"
          />
        </Grid>
      </Grid>
      <Grid container spacing={2} padding={1}>
        <Grid item xs={12}>
          <TextField
            id="description"
            label="Description"
            multiline
            rows={4}
            variant="standard"
            onChange={(e) => dispatch(setDescription(e.target.value))}
            value={description}
            fullWidth
            InputProps={{ fullWidth: true }}
          />
        </Grid>
      </Grid>

      {activeResourceType == 'machine'
        ? renderMachineMetaDropdown()
        : renderDomainMetaInput()}
      <Grid container spacing={2} padding={1} paddingTop={6}>
        <Grid item xs={12}>
          <TextField
            disabled
            label="Id"
            id="resourceid"
            variant="standard"
            value={resourceId}
            fullWidth
            InputProps={{ fullWidth: true }}
          />
        </Grid>
      </Grid>
      <Grid container spacing={2} padding={1}>
        <Grid
          item
          style={{
            display: 'flex',
            justifyContent: 'flex-end',
            alignItems: 'right',
          }}
          xs={12}
        >
          <Stack direction="row" spacing={2}>
            <Button
              variant="outlined"
              onClick={handleCancelClick}
              startIcon={<RefreshIcon />}
            >
              Cancel
            </Button>
            <Button
              variant="contained"
              onClick={() =>
                handleSaveClick(
                  name,
                  type,
                  description,
                  machineInfo,
                  domainInfo,
                  resourceId
                )
              }
              endIcon={<DeleteIcon />}
            >
              Save
            </Button>
          </Stack>
        </Grid>
      </Grid>
    </Box>
  );
};