// Copyright 2019 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"context"
	"fmt"
	"sort"
	"strings"

	proto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"go.chromium.org/chromiumos/infra/proto/go/device"
	"go.chromium.org/chromiumos/infra/proto/go/lab"
	"go.chromium.org/chromiumos/infra/proto/go/manufacturing"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	ds "go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/grpcutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "infra/appengine/cros/lab_inventory/api/v1"
	"infra/appengine/cros/lab_inventory/app/config"
	"infra/appengine/cros/lab_inventory/app/external/ufs"
	"infra/cros/lab_inventory/changehistory"
	"infra/cros/lab_inventory/datastore"
	"infra/cros/lab_inventory/deviceconfig"
	"infra/cros/lab_inventory/hwid"
	"infra/cros/lab_inventory/manufacturingconfig"
	invlibs "infra/cros/lab_inventory/protos"
	"infra/cros/lab_inventory/utils"
)

// InventoryServerImpl implements service interfaces.
type InventoryServerImpl struct {
}

var (
	getHwidDataFunc            = hwid.GetHwidData
	getDeviceConfigFunc        = deviceconfig.GetCachedConfig
	getManufacturingConfigFunc = manufacturingconfig.GetCachedConfig
)

func getPassedResults(ctx context.Context, results []datastore.DeviceOpResult) []*api.DeviceOpResult {
	passedDevices := make([]*api.DeviceOpResult, 0, len(results))
	for _, res := range datastore.DeviceOpResults(results).Passed() {
		r := new(api.DeviceOpResult)
		r.Id = string(res.Entity.ID)
		r.Hostname = res.Entity.Hostname
		passedDevices = append(passedDevices, r)
		logging.Debugf(ctx, "Passed: %s: %s", r.Hostname, r.Id)
	}
	logging.Infof(ctx, "%d device(s) passed", len(passedDevices))

	return passedDevices
}

func getFailedResults(ctx context.Context, results []datastore.DeviceOpResult, hideUUID bool) []*api.DeviceOpResult {
	failedDevices := make([]*api.DeviceOpResult, 0, len(results))
	for _, res := range datastore.DeviceOpResults(results).Failed() {
		r := new(api.DeviceOpResult)
		r.Hostname = res.Entity.Hostname
		r.ErrorMsg = res.Err.Error()
		id := string(res.Entity.ID)
		if !(hideUUID && strings.HasPrefix(id, datastore.UUIDPrefix)) {
			r.Id = id
		}
		failedDevices = append(failedDevices, r)
		logging.Errorf(ctx, "Failed: %s: %s: %s", r.Hostname, r.Id, r.ErrorMsg)
	}
	if failedCount := len(failedDevices); failedCount > 0 {
		logging.Errorf(ctx, "%d device(s) failed", failedCount)
	} else {
		logging.Infof(ctx, "0 devices failed")
	}

	return failedDevices
}

// AddCrosDevices adds new Chrome OS devices to the inventory.
func (is *InventoryServerImpl) AddCrosDevices(ctx context.Context, req *api.AddCrosDevicesRequest) (resp *api.AddCrosDevicesResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	AddCrosDevices is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

func addFailedDevice(ctx context.Context, failedDevices *[]*api.DeviceOpResult, dev *lab.ChromeOSDevice, err error, operation string) {
	hostname := utils.GetHostname(dev)
	logging.Errorf(ctx, "failed to %s for %s: %s", operation, hostname, err.Error())
	*failedDevices = append(*failedDevices, &api.DeviceOpResult{
		Id:       dev.GetId().GetValue(),
		Hostname: hostname,
		ErrorMsg: err.Error(),
	})

}

func getHwidDataInBatch(ctx context.Context, extendedData []*api.ExtendedDeviceData) ([]*api.ExtendedDeviceData, []*api.DeviceOpResult) {
	// Deduplicate the HWIDs in devices to improve the query performance.
	secret := config.Get(ctx).HwidSecret
	hwids := make([]string, 0, len(extendedData))
	idToHwidData := map[string]*hwid.Data{}
	for _, d := range extendedData {
		hwid := d.LabConfig.GetManufacturingId().GetValue()
		if hwid == "" {
			logging.Warningf(ctx, "%v has empty HWID.", utils.GetHostname(d.LabConfig))
		}
		if _, found := idToHwidData[hwid]; found {
			continue
		}
		hwids = append(hwids, hwid)
		idToHwidData[hwid] = nil
	}

	for _, hwid := range hwids {
		if hwid == "" {
			continue
		}
		if hwidData, err := getHwidDataFunc(ctx, hwid, secret); err == nil {
			idToHwidData[hwid] = hwidData
		} else {
			// HWID server may cannot find records for the HWID. Ignore the
			// error for now.
			logging.Warningf(ctx, "Ignored error: failed to get response from HWID server for %s", hwid)
		}
	}
	newExtendedData := make([]*api.ExtendedDeviceData, 0, len(extendedData))
	for i := range extendedData {
		hwid := extendedData[i].LabConfig.GetManufacturingId().GetValue()
		if hwidData := idToHwidData[hwid]; hwidData != nil {
			extendedData[i].HwidData = &api.HwidData{
				Sku:     hwidData.Sku,
				Variant: hwidData.Variant,
			}
		}
		newExtendedData = append(newExtendedData, extendedData[i])
	}
	return newExtendedData, nil
}

func getDeviceConfigData(ctx context.Context, extendedData []*api.ExtendedDeviceData) ([]*api.ExtendedDeviceData, []*api.DeviceOpResult) {
	// Deduplicate the device config ids to improve the query performance.
	devCfgIds := make([]*device.ConfigId, 0, len(extendedData))
	idToDevCfg := map[string]*device.Config{}
	for _, d := range extendedData {
		convertedID := deviceconfig.ConvertValidDeviceConfigID(d.LabConfig.GetDeviceConfigId())
		fallbackID := getFallbackDeviceConfigID(convertedID)
		logging.Debugf(ctx, "before convert: %s", d.LabConfig.DeviceConfigId.String())
		logging.Debugf(ctx, "real device config ID: %s", convertedID.String())
		logging.Debugf(ctx, "fallback device config ID: %s", fallbackID.String())

		for _, cID := range []*device.ConfigId{convertedID, fallbackID} {
			if _, found := idToDevCfg[cID.String()]; found {
				continue
			}
			devCfgIds = append(devCfgIds, cID)
			idToDevCfg[cID.String()] = nil
		}
	}

	devCfgs, err := getDeviceConfigFunc(ctx, devCfgIds)
	for i := range devCfgs {
		if err == nil || err.(errors.MultiError)[i] == nil {
			idToDevCfg[devCfgIds[i].String()] = devCfgs[i].(*device.Config)
		} else {
			logging.Warningf(ctx, "Ignored error: cannot get device config for %v: %v", devCfgIds[i], err.(errors.MultiError)[i])
		}
	}
	newExtendedData := make([]*api.ExtendedDeviceData, 0, len(extendedData))
	failedDevices := make([]*api.DeviceOpResult, 0, len(extendedData))
	for i := range extendedData {
		convertedID := deviceconfig.ConvertValidDeviceConfigID(extendedData[i].LabConfig.GetDeviceConfigId())
		extendedData[i].DeviceConfig = idToDevCfg[convertedID.String()]
		if extendedData[i].DeviceConfig == nil || extendedData[i].DeviceConfig.GetId() == nil {
			fallbackID := getFallbackDeviceConfigID(convertedID)
			extendedData[i].DeviceConfig = idToDevCfg[fallbackID.String()]
		}
		newExtendedData = append(newExtendedData, extendedData[i])
	}
	return newExtendedData, failedDevices
}

func getFallbackDeviceConfigID(oldConfigID *device.ConfigId) *device.ConfigId {
	if oldConfigID.GetVariantId().GetValue() != "" {
		fallbackID := proto.Clone(oldConfigID).(*device.ConfigId)
		fallbackID.VariantId = nil
		return fallbackID
	}
	return oldConfigID
}

func getManufacturingConfigData(ctx context.Context, extendedData []*api.ExtendedDeviceData) ([]*api.ExtendedDeviceData, []*api.DeviceOpResult) {
	// Start to retrieve manufacturing config data.
	cfgIds := make([]*manufacturing.ConfigID, 0, len(extendedData))
	idToCfg := map[string]*manufacturing.Config{}
	for _, d := range extendedData {
		manufacturingID := d.LabConfig.GetManufacturingId()
		if manufacturingID.GetValue() == "" {
			// We use manufacturingID as Key to query datastore. When it's
			// empty, datastore.Get will fail due to incomplete key and all
			// entities queried in same request will be <nil>.
			continue
		}
		if _, found := idToCfg[manufacturingID.GetValue()]; found {
			continue
		}
		cfgIds = append(cfgIds, manufacturingID)
		idToCfg[manufacturingID.GetValue()] = nil
	}
	mCfgs, err := getManufacturingConfigFunc(ctx, cfgIds)
	for i, d := range mCfgs {
		if err == nil || err.(errors.MultiError)[i] == nil {
			idToCfg[cfgIds[i].GetValue()] = d.(*manufacturing.Config)
		} else {
			logging.Warningf(ctx, "Ignored error: cannot get manufacturing config for %v: %v", cfgIds[i], err.(errors.MultiError)[i])
		}
	}
	newExtendedData := make([]*api.ExtendedDeviceData, 0, len(extendedData))
	failedDevices := make([]*api.DeviceOpResult, 0, len(extendedData))
	for i := range extendedData {
		if manufacturingID := extendedData[i].LabConfig.GetManufacturingId().GetValue(); manufacturingID != "" {
			extendedData[i].ManufacturingConfig = idToCfg[manufacturingID]
		}
		newExtendedData = append(newExtendedData, extendedData[i])
	}
	return newExtendedData, failedDevices
}

// GetExtendedDeviceData gets the lab data joined with device config,
// manufacturing config, etc.
func GetExtendedDeviceData(ctx context.Context, devices []datastore.DeviceOpResult) ([]*api.ExtendedDeviceData, []*api.DeviceOpResult) {
	logging.Debugf(ctx, "Get exteneded data for %d devcies", len(devices))
	extendedData := make([]*api.ExtendedDeviceData, 0, len(devices))
	failedDevices := make([]*api.DeviceOpResult, 0, len(devices))
	for _, r := range devices {
		var labData lab.ChromeOSDevice
		logging.Debugf(ctx, "get ext data for %v", r.Entity.Hostname)
		if err := r.Entity.GetCrosDeviceProto(&labData); err != nil {
			logging.Errorf(ctx, "Wrong lab config data of device entity %s", r.Entity)
			failedDevices = append(failedDevices, &api.DeviceOpResult{
				Id:       string(r.Entity.ID),
				Hostname: r.Entity.Hostname,
				ErrorMsg: err.Error(),
			})
			continue
		}
		var dutState lab.DutState
		if err := r.Entity.GetDutStateProto(&dutState); err != nil {
			addFailedDevice(ctx, &failedDevices, &labData, err, "unmarshal dut state data")
			continue
		}

		data := api.ExtendedDeviceData{
			LabConfig: &labData,
			DutState:  &dutState,
		}
		extendedData = append(extendedData, &data)
	}
	// Get HWID data in a batch.
	extendedData, moreFailedDevices := getHwidDataInBatch(ctx, extendedData)
	failedDevices = append(failedDevices, moreFailedDevices...)

	// Get device config in a batch.
	extendedData, moreFailedDevices = getDeviceConfigData(ctx, extendedData)
	failedDevices = append(failedDevices, moreFailedDevices...)

	extendedData, moreFailedDevices = getManufacturingConfigData(ctx, extendedData)
	failedDevices = append(failedDevices, moreFailedDevices...)
	logging.Debugf(ctx, "Got extended data for %d device(s)", len(extendedData))
	return extendedData, failedDevices
}

// GetExtendedDeviceDataForUFSRouting gets the lab data joined with device config,
// manufacturing config, etc.
func GetExtendedDeviceDataForUFSRouting(ctx context.Context, extendedData []*api.ExtendedDeviceData) ([]*api.ExtendedDeviceData, []*api.DeviceOpResult) {
	failedDevices := make([]*api.DeviceOpResult, 0, len(extendedData))
	// Get HWID data in a batch.
	extendedData, moreFailedDevices := getHwidDataInBatch(ctx, extendedData)
	failedDevices = append(failedDevices, moreFailedDevices...)

	// Get device config in a batch.
	extendedData, moreFailedDevices = getDeviceConfigData(ctx, extendedData)
	failedDevices = append(failedDevices, moreFailedDevices...)

	extendedData, moreFailedDevices = getManufacturingConfigData(ctx, extendedData)
	failedDevices = append(failedDevices, moreFailedDevices...)
	logging.Debugf(ctx, "Got extended data for %d device(s)", len(extendedData))
	return extendedData, failedDevices
}

type requestWithIds interface {
	GetIds() []*api.DeviceID
}

// extractHostnamesAndDeviceIDs extracts hostnames and lab.ChromeOSDeviceIDs
// from the input request.
func extractHostnamesAndDeviceIDs(ctx context.Context, req requestWithIds) ([]string, []string) {
	reqIds := req.GetIds()
	maxLen := len(reqIds)
	hostnames := make([]string, 0, maxLen)
	devIds := make([]string, 0, maxLen)
	for _, id := range reqIds {
		if _, ok := id.GetId().(*api.DeviceID_Hostname); ok {
			hostnames = append(hostnames, id.GetHostname())
		} else {
			devIds = append(devIds, id.GetChromeosDeviceId())
		}
	}
	logging.Debugf(ctx, "There are %d hostnames and %d Chrome OS Device IDs in the request", len(hostnames), len(devIds))
	return hostnames, devIds
}

// GetCrosDevices retrieves requested Chrome OS devices from the inventory.
func (is *InventoryServerImpl) GetCrosDevices(ctx context.Context, req *api.GetCrosDevicesRequest) (resp *api.GetCrosDevicesResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	if err = req.Validate(); err != nil {
		return nil, err
	}

	hostnames, devIds := extractHostnamesAndDeviceIDs(ctx, req)

	// Route the call to UFS
	if config.Get(ctx).GetRouting().GetGetCrosDevices() {
		logging.Infof(ctx, "Routing GetCrosDevices to UFS: %+v", req)
		ufsClient, err := ufs.GetUFSClient(ctx)
		if err != nil {
			return nil, err
		}
		var failedDevices []*api.DeviceOpResult
		var devices []*lab.ChromeOSDevice
		crosDevices, failed := ufs.GetUFSDevicesByIds(ctx, ufsClient, devIds)
		logging.Debugf(ctx, "Get %d devices by ID(UFS)", len(devices))
		devices = append(devices, crosDevices...)
		failedDevices = append(failedDevices, failed...)

		crosDevices, failed = ufs.GetUFSDevicesByHostnames(ctx, ufsClient, hostnames)
		logging.Debugf(ctx, "Get %d more devices by hostname(UFS)", len(devices))
		devices = append(devices, crosDevices...)
		failedDevices = append(failedDevices, failed...)

		crosDevices, failed = ufs.GetUFSDevicesByModels(ctx, ufsClient, req.GetModels())
		logging.Debugf(ctx, "Get %d more devices by model(UFS)", len(devices))
		devices = append(devices, crosDevices...)
		failedDevices = append(failedDevices, failed...)

		extendedData, moreFailedDevices := ufs.GetUFSDutStateForDevices(ctx, ufsClient, devices)
		failedDevices = append(failedDevices, moreFailedDevices...)

		extendedData, moreFailedDevices = GetExtendedDeviceDataForUFSRouting(ctx, extendedData)

		failedDevices = append(failedDevices, moreFailedDevices...)
		resp = &api.GetCrosDevicesResponse{
			Data:          extendedData,
			FailedDevices: failedDevices,
		}
		return resp, nil
	}

	result := ([]datastore.DeviceOpResult)(datastore.GetDevicesByIds(ctx, devIds))
	logging.Debugf(ctx, "Get %d devices by ID", len(result))
	result = append(result, datastore.GetDevicesByHostnames(ctx, hostnames)...)
	logging.Debugf(ctx, "Get %d more devices by hostname", len(result))
	byModels, err := datastore.GetDevicesByModels(ctx, req.GetModels())
	if err != nil {
		return nil, errors.Annotate(err, "get devices by models").Err()
	}
	result = append(result, byModels...)
	logging.Debugf(ctx, "Get %d more devices by models", len(result))

	extendedData, moreFailedDevices := GetExtendedDeviceData(ctx, datastore.DeviceOpResults(result).Passed())
	failedDevices := getFailedResults(ctx, result, false)
	failedDevices = append(failedDevices, moreFailedDevices...)

	resp = &api.GetCrosDevicesResponse{
		Data:          extendedData,
		FailedDevices: failedDevices,
	}
	return resp, nil
}

func logDeviceOpResults(ctx context.Context, res datastore.DeviceOpResults) {
	for _, r := range res {
		if r.Err == nil {
			logging.Debugf(ctx, "Device ID %s: succeed", r.Entity.ID)
		} else {
			logging.Debugf(ctx, "Device ID %s: %s", r.Entity.ID, r.Err)
		}
	}
}

// UpdateDutsStatus updates selected Duts' status labels, metas related to testing.
func (is *InventoryServerImpl) UpdateDutsStatus(ctx context.Context, req *api.UpdateDutsStatusRequest) (resp *api.UpdateDutsStatusResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	if err = req.Validate(); err != nil {
		return nil, err
	}

	// Route the call to UFS
	if config.Get(ctx).GetRouting().GetUpdateDutsStatus() {
		logging.Infof(ctx, "Routing UpdateDutsStatus to UFS: %+v", req)
		passed, failed, err := ufs.UpdateUFSDutState(ctx, req)
		if err != nil {
			logging.Errorf(ctx, "fail to update dutmeta, labmeta and dutstate in UFS: %s", err.Error())
			return nil, err
		}
		return &api.UpdateDutsStatusResponse{
			UpdatedDevices: passed,
			FailedDevices:  failed,
		}, nil
	}

	meta := make(map[string]datastore.DutMeta, len(req.GetDutMetas()))
	for _, d := range req.GetDutMetas() {
		meta[d.GetChromeosDeviceId()] = datastore.DutMeta{
			SerialNumber: d.GetSerialNumber(),
			HwID:         d.GetHwID(),
			DeviceSku:    d.GetDeviceSku(),
		}
	}
	metaUpdateResults, err := datastore.UpdateDutMeta(ctx, meta)
	logging.Debugf(ctx, "Meta update results")
	logDeviceOpResults(ctx, metaUpdateResults)
	if err != nil {
		logging.Errorf(ctx, "fail to update dut meta: %s", err.Error())
		return nil, err
	}

	labMeta := make(map[string]datastore.LabMeta, len(req.GetLabMetas()))
	for _, d := range req.GetLabMetas() {
		labMeta[d.GetChromeosDeviceId()] = datastore.LabMeta{
			ServoType:     d.GetServoType(),
			SmartUsbhub:   d.GetSmartUsbhub(),
			ServoTopology: d.GetServoTopology(),
		}
	}
	metaUpdateResults, err = datastore.UpdateLabMeta(ctx, labMeta)
	logging.Debugf(ctx, "Lab meta update results")
	logDeviceOpResults(ctx, metaUpdateResults)
	if err != nil {
		logging.Errorf(ctx, "fail to update lab meta: %s", err.Error())
		return nil, err
	}

	updatingResults, err := datastore.UpdateDutsStatus(changehistory.Use(ctx, req.Reason), req.States)
	if err != nil {
		return nil, err
	}
	logging.Debugf(ctx, "State update results")
	logDeviceOpResults(ctx, updatingResults)

	updatedDevices := getPassedResults(ctx, updatingResults)
	failedDevices := getFailedResults(ctx, updatingResults, false)
	resp = &api.UpdateDutsStatusResponse{
		UpdatedDevices: updatedDevices,
		FailedDevices:  failedDevices,
	}
	return resp, nil
}

// UpdateLabstations updates the given labstations.
func (is *InventoryServerImpl) UpdateLabstations(ctx context.Context, req *api.UpdateLabstationsRequest) (resp *api.UpdateLabstationsResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	 UpdateLabstations is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

// UpdateCrosDevicesSetup updates the selected Chrome OS devices setup data in
// the inventory.
func (is *InventoryServerImpl) UpdateCrosDevicesSetup(ctx context.Context, req *api.UpdateCrosDevicesSetupRequest) (resp *api.UpdateCrosDevicesSetupResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	UpdateCrosDevicesSetup is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

func getRemovalReason(req *api.DeleteCrosDevicesRequest) string {
	if r := req.GetReason(); r.GetBug() != "" || r.GetComment() != "" {
		return fmt.Sprintf("%s: %s", r.GetBug(), r.GetComment())
	}
	return ""
}

// DeleteCrosDevices delete the selelcted devices from the inventory.
func (is *InventoryServerImpl) DeleteCrosDevices(ctx context.Context, req *api.DeleteCrosDevicesRequest) (resp *api.DeleteCrosDevicesResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	DeleteCrosDevices is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

// BatchUpdateDevices updates some specific devices properties in batch.
func (is *InventoryServerImpl) BatchUpdateDevices(ctx context.Context, req *api.BatchUpdateDevicesRequest) (resp *api.BatchUpdateDevicesResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	BatchUpdateDevices is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

// AddAssets adds a record of the given asset to datastore
func (is *InventoryServerImpl) AddAssets(ctx context.Context, req *api.AssetList) (response *api.AssetResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	AddAssets is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

// UpdateAssets updates a record of the given asset to datastore
func (is *InventoryServerImpl) UpdateAssets(ctx context.Context, req *api.AssetList) (response *api.AssetResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	UpdateAssets is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

// GetAssets retrieves the asset information given its asset ID
func (is *InventoryServerImpl) GetAssets(ctx context.Context, req *api.AssetIDList) (response *api.AssetResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	GetAssets is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

// ListCrosDevicesLabConfig retrieves all lab configs
func (is *InventoryServerImpl) ListCrosDevicesLabConfig(ctx context.Context, req *api.ListCrosDevicesLabConfigRequest) (response *api.ListCrosDevicesLabConfigResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	allDevices, err := datastore.GetAllDevices(ctx)
	logging.Debugf(ctx, "got devices (%d)", len(allDevices))
	if err != nil {
		return nil, errors.Annotate(err, "get all devices").Err()
	}
	labConfigs := make([]*api.ListCrosDevicesLabConfigResponse_LabConfig, 0, len(allDevices))
	for _, d := range allDevices {
		if d.Entity == nil || d.Err != nil || d.Entity.ID == "" {
			continue
		}
		dev := &lab.ChromeOSDevice{}
		if err := d.Entity.GetCrosDeviceProto(dev); err != nil {
			logging.Debugf(ctx, "fail to get lab config proto for %s (%s)", d.Entity.ID, d.Entity.Hostname)
		}
		dutState := &lab.DutState{}
		if err := d.Entity.GetDutStateProto(dutState); err != nil {
			logging.Debugf(ctx, "fail to get dut state proto for %s (%s)", d.Entity.ID, d.Entity.Hostname)
		}
		utime, _ := ptypes.TimestampProto(d.Entity.Updated)
		labConfigs = append(labConfigs, &api.ListCrosDevicesLabConfigResponse_LabConfig{
			Config:      dev,
			State:       dutState,
			UpdatedTime: utime,
		})
	}
	return &api.ListCrosDevicesLabConfigResponse{
		LabConfigs: labConfigs,
	}, nil
}

// DeleteAssets deletes the asset information from datastore
func (is *InventoryServerImpl) DeleteAssets(ctx context.Context, req *api.AssetIDList) (response *api.AssetIDResponse, err error) {
	// TODO(crbug.com/1184794): Remove all the dependencies
	return nil, grpcutil.GRPCifyAndLogErr(ctx, errors.Reason(`
	DeleteAssets is deprecated.
	Check go/shivas, go/shivas-manual-os for alternative tools.
	Unified-Fleet-System [UFS] maintains an alternative API.
	`).Err())
}

func seperateAssetIDResults(a []*datastore.AssetOpResult) (pAssetIDs, fAssetIDs []*api.AssetIDResult) {
	passed, failed := seperateAssetResults(a)
	pAssetIDs = make([]*api.AssetIDResult, 0, len(passed))
	fAssetIDs = make([]*api.AssetIDResult, 0, len(failed))
	toAssetIDResult := func(b *api.AssetResult) *api.AssetIDResult {
		return &api.AssetIDResult{
			Id:       b.Asset.GetId(),
			ErrorMsg: b.ErrorMsg,
		}
	}
	for _, res := range passed {
		pAssetIDs = append(pAssetIDs, toAssetIDResult(res))
	}
	for _, res := range failed {
		fAssetIDs = append(fAssetIDs, toAssetIDResult(res))
	}
	return pAssetIDs, fAssetIDs
}

func seperateAssetResults(results []*datastore.AssetOpResult) (success, failure []*api.AssetResult) {
	successResults := make([]*api.AssetResult, 0, len(results))
	failureResults := make([]*api.AssetResult, 0, len(results))
	for _, res := range results {
		if res.Err != nil {
			var failedResult api.AssetResult
			failedResult.Asset = res.ToAsset()
			failedResult.ErrorMsg = res.Err.Error()
			failureResults = append(failureResults, &failedResult)
		} else {
			var successResult api.AssetResult
			successResult.Asset = res.ToAsset()
			successResults = append(successResults, &successResult)
		}
	}
	return successResults, failureResults
}

// DeviceConfigsExists checks if the device_configs for the given configIds exists in the datastore
func (is *InventoryServerImpl) DeviceConfigsExists(ctx context.Context, req *api.DeviceConfigsExistsRequest) (rsp *api.DeviceConfigsExistsResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	devCfgIds := make([]*device.ConfigId, 0, len(req.ConfigIds))
	for _, d := range req.ConfigIds {
		convertedID := deviceconfig.ConvertValidDeviceConfigID(d)
		devCfgIds = append(devCfgIds, convertedID)
	}
	res, err := deviceconfig.DeviceConfigsExists(ctx, devCfgIds)
	if err != nil {
		return nil, err
	}
	response := &api.DeviceConfigsExistsResponse{
		Exists: res,
	}
	return response, err
}

// GetDeviceManualRepairRecord checks and returns a manual repair record for
// a given device hostname if it exists.
func (is *InventoryServerImpl) GetDeviceManualRepairRecord(ctx context.Context, req *api.GetDeviceManualRepairRecordRequest) (rsp *api.GetDeviceManualRepairRecordResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	propFilter := map[string]string{
		"hostname":     req.Hostname,
		"repair_state": "STATE_IN_PROGRESS",
	}
	getRes, err := datastore.GetRepairRecordByPropertyName(ctx, propFilter, -1, 0, []string{})
	if err != nil {
		return nil, errors.Annotate(err, "Error encountered for get request %s", req.Hostname).Tag(grpcutil.InvalidArgumentTag).Err()
	}

	// There should only be one record in progress per hostname at a time. User
	// should complete the record returned.
	if len(getRes) == 0 {
		return nil, errors.Reason("No record found").Tag(grpcutil.InvalidArgumentTag).Err()
	} else if len(getRes) > 1 {
		err = errors.Reason("More than one active record found; returning first record").Tag(grpcutil.InvalidArgumentTag).Err()
	}

	// Return first active record found.
	r := getRes[0]
	if e := r.Err; e != nil {
		err = errors.Annotate(e, "Error encountered for record %s", r.Entity.ID).Tag(grpcutil.InvalidArgumentTag).Err()
	}

	rsp = &api.GetDeviceManualRepairRecordResponse{
		DeviceRepairRecord: r.Record,
		Id:                 r.Entity.ID,
	}
	return rsp, err
}

// CreateDeviceManualRepairRecord adds a new submitted manual repair record for
// a given device.
func (is *InventoryServerImpl) CreateDeviceManualRepairRecord(ctx context.Context, req *api.CreateDeviceManualRepairRecordRequest) (rsp *api.CreateDeviceManualRepairRecordResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	// Query asset info using hostname and create records
	record := req.DeviceRepairRecord
	hostname := record.Hostname

	_, err = queryToCreateHost(ctx, hostname)
	if err != nil {
		return nil, err
	}

	record = parseToCreateRequest(ctx, record, timestamppb.Now())

	resp, err := datastore.AddDeviceManualRepairRecords(ctx, []*invlibs.DeviceManualRepairRecord{record})
	if err != nil {
		return nil, err
	}

	if len(resp) > 0 && resp[0].Err != nil {
		return nil, resp[0].Err
	}

	return &api.CreateDeviceManualRepairRecordResponse{}, nil
}

// UpdateDeviceManualRepairRecord updates an existing manual repair record with
// new submitted info for a given device.
func (is *InventoryServerImpl) UpdateDeviceManualRepairRecord(ctx context.Context, req *api.UpdateDeviceManualRepairRecordRequest) (rsp *api.UpdateDeviceManualRepairRecordResponse, err error) {
	id := req.Id
	record := req.DeviceRepairRecord
	hostname := record.Hostname

	if id == "" {
		return nil, errors.Reason("ID cannot be empty").Tag(grpcutil.InvalidArgumentTag).Err()
	}

	// Check if an open record exists for this hostname. If not, throw error
	// stating no active record to be updated.
	getRes, err := queryInProgressMRHost(ctx, hostname)
	if err != nil {
		return nil, err
	}

	if len(getRes) == 0 {
		return nil, errors.Reason("No open record exists for host %s; Please create a new record", hostname).Tag(grpcutil.InvalidArgumentTag).Err()
	}

	// Construct updated record
	record.UpdatedTime = timestamppb.Now()
	if record.RepairState == invlibs.DeviceManualRepairRecord_STATE_COMPLETED {
		record.CompletedTime = record.UpdatedTime
	}

	resp, err := datastore.UpdateDeviceManualRepairRecords(ctx, map[string]*invlibs.DeviceManualRepairRecord{id: record})
	if err != nil {
		return nil, err
	}

	if len(resp) > 0 && resp[0].Err != nil {
		return nil, resp[0].Err
	}

	return &api.UpdateDeviceManualRepairRecordResponse{}, nil
}

// queryInProgressMRHost takes a hostname and queries for any in progress
// Manual Repair records associated with it.
func queryInProgressMRHost(ctx context.Context, hostname string) ([]*datastore.DeviceManualRepairRecordsOpRes, error) {
	getRes, err := datastore.GetRepairRecordByPropertyName(ctx, map[string]string{
		"hostname":     hostname,
		"repair_state": "STATE_IN_PROGRESS",
	}, -1, 0, []string{})
	if err != nil {
		return nil, errors.Annotate(err, "Failed to query STATE_IN_PROGRESS host %s", hostname).Tag(grpcutil.InvalidArgumentTag).Err()
	}

	return getRes, err
}

// ListManualRepairRecords takes filtering parameters and returns a list of
// repair records that match the filters.
//
// Currently supports filtering on:
// - hostname
// - asset tag
// - user ldap
// - repair state
// - limit (number of records)
// - offset - used for pagination
func (is *InventoryServerImpl) ListManualRepairRecords(ctx context.Context, req *api.ListManualRepairRecordsRequest) (rsp *api.ListManualRepairRecordsResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	var propFilter = map[string]string{}

	if req.GetHostname() != "" {
		propFilter["hostname"] = req.GetHostname()
	}

	if req.GetAssetTag() != "" {
		propFilter["asset_tag"] = req.GetAssetTag()
	}

	if req.GetUserLdap() != "" {
		propFilter["user_ldap"] = req.GetUserLdap()
	}

	if req.GetRepairState() != "" {
		propFilter["repair_state"] = req.GetRepairState()
	}

	limit := req.GetLimit()
	if limit <= 0 {
		limit = -1
	}

	offset := req.GetOffset()
	if offset < 0 {
		offset = 0
	}

	getRes, err := datastore.GetRepairRecordByPropertyName(ctx, propFilter, limit, offset, []string{"-updated_time"})
	if err != nil {
		return nil, errors.Annotate(err, "Error encountered for get request %s", req.Hostname).Err()
	}

	var repairRecords []*invlibs.DeviceManualRepairRecord
	for _, r := range getRes {
		repairRecords = append(repairRecords, r.Record)
	}

	return &api.ListManualRepairRecordsResponse{
		RepairRecords: repairRecords,
	}, err
}

// GetManufacturingConfig retrieves requested Chrome OS device manufacturing config from the inventory.
func (is *InventoryServerImpl) GetManufacturingConfig(ctx context.Context, req *api.GetManufacturingConfigRequest) (resp *manufacturing.Config, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	if err = req.Validate(); err != nil {
		return nil, err
	}
	mfgCfgID := &manufacturing.ConfigID{Value: req.GetName()}
	cfgIds := []*manufacturing.ConfigID{mfgCfgID}
	mCfgs, err := getManufacturingConfigFunc(ctx, cfgIds)
	if err != nil {
		if err.Error() == ds.ErrNoSuchEntity.Error() {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("manufacturing config not found for %s", req.GetName()))
		}
		return nil, err.(errors.MultiError)[0]
	}
	if len(mCfgs) == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("manufacturing config not found for %s", req.GetName()))
	}
	return mCfgs[0].(*manufacturing.Config), nil
}

// GetDeviceConfig retrieves requested Chrome OS device device config from the inventory.
func (is *InventoryServerImpl) GetDeviceConfig(ctx context.Context, req *api.GetDeviceConfigRequest) (resp *device.Config, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	if err = req.Validate(); err != nil {
		return nil, err
	}
	convertedID := deviceconfig.ConvertValidDeviceConfigID(req.GetConfigId())
	fallbackID := getFallbackDeviceConfigID(convertedID)
	logging.Debugf(ctx, "before convert: %s", req.GetConfigId().String())
	logging.Debugf(ctx, "real device config ID: %s", convertedID.String())
	logging.Debugf(ctx, "fallback device config ID: %s", fallbackID.String())

	devCfgIds := make([]*device.ConfigId, 0, 1)
	idToDevCfg := map[string]*device.Config{}
	for _, cID := range []*device.ConfigId{convertedID, fallbackID} {
		if _, found := idToDevCfg[cID.String()]; found {
			continue
		}
		devCfgIds = append(devCfgIds, cID)
		idToDevCfg[cID.String()] = nil
	}

	devCfgs, err := getDeviceConfigFunc(ctx, devCfgIds)
	for i := range devCfgs {
		if err == nil || err.(errors.MultiError)[i] == nil {
			idToDevCfg[devCfgIds[i].String()] = devCfgs[i].(*device.Config)
		} else {
			logging.Warningf(ctx, "Ignored error: cannot get device config for %v: %v", devCfgIds[i], err.(errors.MultiError)[i])
		}
	}

	res := idToDevCfg[convertedID.String()]
	if res == nil || res.GetId() == nil {
		res = idToDevCfg[fallbackID.String()]
		if res == nil || res.GetId() == nil {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("device config not found for %+v", req.GetConfigId()))
		}
	}
	return res, nil
}

// GetHwidData retrieves requested Chrome OS device Hwid Data from the inventory.
func (is *InventoryServerImpl) GetHwidData(ctx context.Context, req *api.GetHwidDataRequest) (resp *api.HwidData, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	if err = req.Validate(); err != nil {
		return nil, err
	}
	secret := config.Get(ctx).HwidSecret
	hwidData, err := getHwidDataFunc(ctx, req.GetName(), secret)
	if err != nil {
		if strings.Contains(err.Error(), "\"code\": 5") {
			// NotFound status code is 5
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Hwid data not found for %s", req.GetName()))
		}
		return nil, err
	}
	return &api.HwidData{
		Sku:     hwidData.Sku,
		Variant: hwidData.Variant,
	}, nil
}

// BatchGetManualRepairRecords gets the open record corresponding to each host
// in the list of given hostnames. If no open record is found, an empty object
// will be returned for that hostname.
func (is *InventoryServerImpl) BatchGetManualRepairRecords(ctx context.Context, req *api.BatchGetManualRepairRecordsRequest) (rsp *api.BatchGetManualRepairRecordsResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	hostnames := req.Hostnames
	repairRecords := make([]*api.ManualRepairRecordResult, 0, len(hostnames))

	for _, h := range hostnames {
		propFilter := map[string]string{
			"hostname":     h,
			"repair_state": "STATE_IN_PROGRESS",
		}
		getRes, getErr := datastore.GetRepairRecordByPropertyName(ctx, propFilter, -1, 0, []string{})
		repairRecords = append(repairRecords, parseManualRepairRecordResult(ctx, h, getRes, getErr))
	}

	return &api.BatchGetManualRepairRecordsResponse{
		RepairRecords: repairRecords,
	}, nil
}

// BatchCreateManualRepairRecords creates new submitted manual repair records
// for a batch of given devices. All records will have the same CreatedTime.
func (is *InventoryServerImpl) BatchCreateManualRepairRecords(ctx context.Context, req *api.BatchCreateManualRepairRecordsRequest) (rsp *api.BatchCreateManualRepairRecordsResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()

	requestRecords := req.RepairRecords
	createdTime := timestamppb.Now()
	allRecordsMap := make(map[string]*api.ManualRepairRecordResult)
	toCreateRecords := make([]*invlibs.DeviceManualRepairRecord, 0, len(requestRecords))

	for _, r := range requestRecords {
		rec := &api.ManualRepairRecordResult{
			Hostname: r.Hostname,
		}

		_, err := queryToCreateHost(ctx, r.Hostname)
		if err != nil {
			rec.ErrorMsg = err.Error()
		} else {
			toCreateRec := parseToCreateRequest(ctx, r, createdTime)
			rec.RepairRecord = toCreateRec
			toCreateRecords = append(toCreateRecords, toCreateRec)
		}

		allRecordsMap[r.Hostname] = rec
	}

	createRsp, err := datastore.AddDeviceManualRepairRecords(ctx, toCreateRecords)
	if err != nil {
		return nil, err
	}

	for _, r := range createRsp {
		if r.Err != nil {
			allRecordsMap[r.Record.Hostname].ErrorMsg = r.Err.Error()
		} else {
			allRecordsMap[r.Record.Hostname].Id = r.Entity.ID
			allRecordsMap[r.Record.Hostname].RepairRecord = r.Record
		}
	}

	// Gather keys and sort result by hostname.
	keys := make([]string, 0)
	for k := range allRecordsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	allRecords := make([]*api.ManualRepairRecordResult, 0, len(requestRecords))
	for _, k := range keys {
		allRecords = append(allRecords, allRecordsMap[k])
	}

	return &api.BatchCreateManualRepairRecordsResponse{
		RepairRecords: allRecords,
	}, nil
}

// parseManualRepairRecordResult parses the repair records found in the
// datastore and converts it into a ManualRepairRecordResult. If an error
// occurs, the error is attached to the result object and returned.
func parseManualRepairRecordResult(ctx context.Context, hostname string, getRes []*datastore.DeviceManualRepairRecordsOpRes, getErr error) *api.ManualRepairRecordResult {
	recordRes := &api.ManualRepairRecordResult{
		Hostname: hostname,
	}

	if getErr != nil {
		recordRes.ErrorMsg = "Error encountered for get request " + hostname
		return recordRes
	}

	// There should only be one record in progress per hostname at a time. User
	// should complete the record returned.
	if len(getRes) == 0 {
		recordRes.ErrorMsg = "No record found"
		return recordRes
	} else if len(getRes) > 1 {
		logging.Warningf(ctx, "More than one active record found; returning first record")
	}

	// Return first active record found.
	r := getRes[0]
	if e := r.Err; e != nil {
		recordRes.ErrorMsg = "Error encountered for record " + r.Entity.ID
		return recordRes
	}

	recordRes.Id = r.Entity.ID
	recordRes.RepairRecord = r.Record
	return recordRes
}

// queryToCreateHost queries datastore based on hostname and checks if an
// open record already exists. Returns query result and error if any.
func queryToCreateHost(ctx context.Context, hostname string) ([]*datastore.DeviceManualRepairRecordsOpRes, error) {
	getRes, err := queryInProgressMRHost(ctx, hostname)
	if err != nil {
		return nil, err
	}

	if len(getRes) >= 1 {
		return nil, errors.Reason("A record already exists for host %s; Please complete the existing record", hostname).Tag(grpcutil.InvalidArgumentTag).Err()
	}

	return getRes, nil
}

// parseToCreateRequest takes a repair record to be created and fills in the
// queried asset tag from datastore and timestamps.
func parseToCreateRequest(ctx context.Context, r *invlibs.DeviceManualRepairRecord, createdTime *timestamppb.Timestamp) *invlibs.DeviceManualRepairRecord {
	// Try to get asset tag from DeviceEntity id. In most cases (current coverage
	// is ~71%), DeviceEntity will use asset tag as its ID. UUID is used if not
	// asset tag. We will set its asset tag to "n/a" in the interim if the current
	// id is the same as the hostname entered by the user.
	//
	// The ID should either be an asset tag or a uuid. Checking if it equals
	// hostname is to prevent datastore errors.
	var assetTag string
	// Route the call to UFS
	if config.Get(ctx).GetRouting().GetGetCrosDevices() {
		logging.Infof(ctx, "Routing GetCrosDevices/ManualRepair to UFS: %s", r.Hostname)
		assetTag = "n/a"
		ufsClient, err := ufs.GetUFSClient(ctx)
		if err == nil {
			devices, _ := ufs.GetUFSDevicesByHostnames(ctx, ufsClient, []string{r.Hostname})
			if len(devices) > 0 {
				assetTag = devices[0].GetId().GetValue()
			} else {
				logging.Warningf(ctx, "Device not found; setting asset tag to n/a")
			}
		} else {
			logging.Warningf(ctx, "Failed to get UFSClient; setting asset tag to n/a")
		}
	} else {
		devices := datastore.GetDevicesByHostnames(ctx, []string{r.Hostname})
		if err := devices[0].Err; err != nil {
			logging.Warningf(ctx, "DeviceEntity not queryable; setting asset tag to n/a")
			assetTag = "n/a"
		} else {
			assetTag = string(devices[0].Entity.ID)
		}
	}

	// Fill in the asset tag field for the records and write to datastore.
	r.AssetTag = assetTag
	r.CreatedTime = createdTime
	r.UpdatedTime = r.CreatedTime

	if r.RepairState == invlibs.DeviceManualRepairRecord_STATE_COMPLETED {
		r.CompletedTime = r.CreatedTime
	}

	return r
}
