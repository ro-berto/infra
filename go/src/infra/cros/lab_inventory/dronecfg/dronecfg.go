// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package dronecfg implements datastore access for storing drone
// configs.
package dronecfg

import (
	"context"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/gae/service/datastore"

	ds "infra/cros/lab_inventory/datastore"
)

// DUT describes a DUT for the purpose of a drone config.
type DUT struct {
	ID       string
	Hostname string
}

// Entity is a drone config datastore entity.
type Entity struct {
	_kind    string `gae:"$kind,droneConfig"`
	Hostname string `gae:"$id"`
	DUTs     []DUT  `gae:",noindex"`
}

func dutsUnion(self, other []DUT) []DUT {
	idToDut := map[string]DUT{}
	for _, d := range self {
		idToDut[d.ID] = d
	}
	for _, d := range other {
		idToDut[d.ID] = d
	}
	result := make([]DUT, len(idToDut))
	idx := 0
	for _, v := range idToDut {
		result[idx] = v
		idx++
	}
	return result
}

func dutsDifference(self, other []DUT) []DUT {
	dutToRemove := map[string]bool{}
	for _, d := range other {
		dutToRemove[d.ID] = true
	}
	result := make([]DUT, 0, len(self))
	for _, d := range self {
		if _, found := dutToRemove[d.ID]; !found {
			result = append(result, d)
		}
	}
	return result
}

// Get gets a drone config from datastore by hostname.
func Get(ctx context.Context, hostname string) (Entity, error) {
	e := Entity{Hostname: hostname}
	if err := datastore.Get(ctx, &e); err != nil {
		return e, errors.Annotate(err, "get drone config").Err()
	}
	return e, nil
}

const queenDronePrefix = "drone-queen-"

// QueenDroneName returns the name of the fake drone whose DUTs should
// be pushed to the drone queen service.
func QueenDroneName(env string) string {
	return queenDronePrefix + env
}

// SyncDeviceList sync the device list of the drone config with the device data
// from the inventory.
func SyncDeviceList(ctx context.Context, droneQueenName string) error {
	var devs []ds.DeviceEntity
	if err := datastore.GetAll(ctx, datastore.NewQuery(ds.DeviceKind), &devs); err != nil {
		return errors.Annotate(err, "sync dev list to drone config").Err()
	}
	logging.Infof(ctx, "Syncing %s devices to drone config", len(devs))
	duts := make([]DUT, len(devs))
	for i := range devs {
		duts[i] = DUT{ID: string(devs[i].ID), Hostname: devs[i].Hostname}
	}
	e, err := Get(ctx, droneQueenName)
	if err != nil {
		return errors.Annotate(err, "get drone config").Err()
	}
	e.DUTs = duts
	if err := datastore.Put(ctx, &e); err != nil {
		return errors.Annotate(err, "save new drone config").Err()
	}
	return nil
}
