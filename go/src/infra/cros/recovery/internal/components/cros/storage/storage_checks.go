// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package storage

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.chromium.org/luci/common/errors"

	"infra/cros/dutstate"
	"infra/cros/recovery/internal/components"
	"infra/cros/recovery/internal/components/cros"
	"infra/cros/recovery/internal/log"
	"infra/cros/recovery/logger/metrics"
	"infra/cros/recovery/tlw"
)

// StorageState is a description of the DUT's storage state given the type of the DUT storage.
type StorageState string

const (
	// DUT storage state cannot be determined.
	StorageStateUndefined StorageState = "UNDEFINED"
	// DUT storage state is normal.
	StorageStateNormal StorageState = "NORMAL"
	// DUT storage state is warning.
	StorageStateWarning StorageState = "WARNING"
	// DUT storage state is critical.
	StorageStateCritical StorageState = "CRITICAL"
)

// storageSMART is used to store the processed information of both storage type and storage state
// after it reads from the storage-info-common.sh file on the DUT.
//
// supported storageType : MMC, NVME, SSD
// supported storageState: storageStateUndefined, storageStateNomral, storageStateWarning, storageStateCritical
type storageSMART struct {
	StorageType  tlw.Storage_Type
	StorageState StorageState
}

// ParseSMARTInfo reads the storage info from SMART.
// The info will be located as collection of lines
func ParseSMARTInfo(ctx context.Context, rawOutput string) (*storageSMART, error) {
	storageType, storageState, err := storageSMARTFieldValue(ctx, rawOutput)
	return &storageSMART{
		StorageType:  storageType,
		StorageState: storageState,
	}, errors.Annotate(err, "parse smart info").Err()
}

type storageStateFunc func(context.Context, []string) (StorageState, error)

var typeToStateFuncMap map[tlw.Storage_Type]storageStateFunc = map[tlw.Storage_Type]storageStateFunc{
	tlw.Storage_SSD:  detectSSDState,
	tlw.Storage_MMC:  detectMMCState,
	tlw.Storage_NVME: detectNVMEState,
	tlw.Storage_UFS:  detectUFSState,
}

// storageSMARTFieldValue takes the raw output from the command line and return the field value of the storageSMART struct.
func storageSMARTFieldValue(ctx context.Context, rawOutput string) (tlw.Storage_Type, StorageState, error) {
	rawOutput = strings.TrimSpace(rawOutput)
	if rawOutput == "" {
		return tlw.Storage_TYPE_UNSPECIFIED, StorageStateUndefined, errors.Reason("storageSMART field value: storage info is empty").Err()
	}
	storageInfoSlice := strings.Split(rawOutput, "\n")
	storageType, err := extractStorageType(ctx, storageInfoSlice)
	if err != nil {
		return tlw.Storage_TYPE_UNSPECIFIED, StorageStateUndefined, errors.Annotate(err, "storageSMART field value").Err()
	}
	funcToCall, typeInMap := typeToStateFuncMap[storageType]
	if !typeInMap {
		return storageType, StorageStateUndefined, nil
	}
	storageState, err := funcToCall(ctx, storageInfoSlice)
	if err != nil {
		return storageType, StorageStateUndefined, errors.Annotate(err, "storageSMART field value").Err()
	}
	return storageType, storageState, nil
}

const (
	// Example "SATA Version is: SATA 3.1, 6.0 Gb/s (current: 6.0 Gb/s)"
	ssdTypeStorageGlob = `SATA Version is:.*`
	// Example "   Extended CSD rev 1.7 (MMC 5.0)"
	mmcTypeStorageGlob = `\s*Extended CSD rev.*MMC (?P<version>\d+.\d+)`
	// Example "SMART/Health Information (NVMe Log 0x02, NSID 0xffffffff)"
	nvmeTypeStorageGlob = `.*NVMe Log .*`
	// Example "$ ufs-utils desc -a -p /dev/bsg/ufs-bsg0"
	ufsTypeStorageGlob = `\s*ufs-utils\s*`
)

// extractStorageType extracts the storage type information from storageInfoSlice.
// return error if the regular expression cannot compile.
func extractStorageType(ctx context.Context, storageInfoSlice []string) (tlw.Storage_Type, error) {
	log.Debugf(ctx, "Extracting storage type")
	ssdTypeRegexp, err := regexp.Compile(ssdTypeStorageGlob)
	if err != nil {
		return tlw.Storage_TYPE_UNSPECIFIED, errors.Annotate(err, "extract storage type").Err()
	}
	mmcTypeRegexp, err := regexp.Compile(mmcTypeStorageGlob)
	if err != nil {
		return tlw.Storage_TYPE_UNSPECIFIED, errors.Annotate(err, "extract storage type").Err()
	}
	nvmeTypeRegexp, err := regexp.Compile(nvmeTypeStorageGlob)
	if err != nil {
		return tlw.Storage_TYPE_UNSPECIFIED, errors.Annotate(err, "extract storage type").Err()
	}
	ufsTypeRegexp, err := regexp.Compile(ufsTypeStorageGlob)
	if err != nil {
		return tlw.Storage_TYPE_UNSPECIFIED, errors.Annotate(err, "extract storage type").Err()
	}
	for _, line := range storageInfoSlice {
		// check if storage type is SSD
		if ssdTypeRegexp.MatchString(line) {
			return tlw.Storage_SSD, nil
		}
		// check if storage type is UFS
		if ufsTypeRegexp.MatchString(line) {
			return tlw.Storage_UFS, nil
		}
		// check if storage type is MMC
		mMMC, err := regexpSubmatchToMap(mmcTypeRegexp, line)
		if err == nil {
			log.Infof(ctx, "Found line => "+line)
			if version, ok := mMMC["version"]; ok {
				log.Infof(ctx, "Found eMMC device, version: %s", version)
			}
			return tlw.Storage_MMC, nil
		}
		// check if storage type is nvme
		if nvmeTypeRegexp.MatchString(line) {
			return tlw.Storage_NVME, nil
		}
	}
	return tlw.Storage_TYPE_UNSPECIFIED, nil
}

const (
	// Field meaning and example line that have failing attribute
	// https://en.wikipedia.org/wiki/S.M.A.R.T.
	// ID# ATTRIBUTE_NAME     FLAGS    VALUE WORST THRESH FAIL RAW_VALUE
	// 184 End-to-End_Error   PO--CK   001   001   097    NOW  135
	ssdFailGlob            = `\s*(?P<param>\S+\s\S+)\s+[P-][O-][S-][R-][C-][K-](\s+\d{3}){3}\s+NOW`
	ssdRelocateSectorsGlob = `\s*\d\sReallocated_Sector_Ct\s*[P-][O-][S-][R-][C-][K-]\s*(?P<value>\d{3})\s*(?P<worst>\d{3})\s*(?P<thresh>\d{3})`
)

// detectSSDState read the info to detect state for SSD storage.
// return error if the regular expression cannot compile.
func detectSSDState(ctx context.Context, storageInfoSlice []string) (StorageState, error) {
	log.Infof(ctx, "Extraction metrics for SSD storage")
	ssdFailRegexp, err := regexp.Compile(ssdFailGlob)
	if err != nil {
		return StorageStateUndefined, errors.Annotate(err, "detect ssd state").Err()
	}
	ssdRelocateSectorsRegexp, err := regexp.Compile(ssdRelocateSectorsGlob)
	if err != nil {
		return StorageStateUndefined, errors.Annotate(err, "detect ssd state").Err()
	}
	for _, line := range storageInfoSlice {
		_, err := regexpSubmatchToMap(ssdFailRegexp, line)
		if err == nil {
			log.Debugf(ctx, "Found critical line => %q", line)
			return StorageStateCritical, nil
		}
		mRelocate, err := regexpSubmatchToMap(ssdRelocateSectorsRegexp, line)
		if err == nil {
			log.Debugf(ctx, "Found warning line => %q", line)
			value, _ := strconv.ParseFloat(mRelocate["value"], 64)
			// manufacture set default value 100, if number started to grow then it is time to mark it.
			if value > 100 {
				return StorageStateWarning, nil
			}
		}
	}
	return StorageStateNormal, nil
}

const (
	// Ex:
	// Device life time type A [DEVICE_LIFE_TIME_EST_TYP_A: 0x01]
	// 0x00~9 means 0-90% band
	// 0x0a means 90-100% band
	// 0x0b means over 100% band
	mmcFailLifeGlob = `.*(?P<param>DEVICE_LIFE_TIME_EST_TYP_.)]?: 0x0(?P<val>\S)` // life time persentage
	// Ex "Pre EOL information [PRE_EOL_INFO: 0x01]"
	// 0x00 - not defined
	// 0x01 - Normal
	// 0x02 - Warning, consumed 80% of the reserved blocks
	// 0x03 - Urgent, consumed 90% of the reserved blocks
	mmcFailEolGlob = `.*(?P<param>PRE_EOL_INFO)]?: 0x0(?P<val>\d)`
)

const (
	// Ex:
	// Device Health Descriptor: [Byte offset 0x3]: bDeviceLifeTimeEstA = 0x1
	// 0x0~9 means 0-90% band
	// 0xa means 90-100% band
	// 0xb means over 100% band
	ufsFailLifeGlob = `.*(?P<param>bDeviceLifeTimeEst.) = 0x(?P<val>\S)`
	// Ex:
	// Device Health Descriptor: [Byte offset 0x2]: bPreEOLInfo = 0x3
	// 0x0 - not defined
	// 0x1 - Normal
	// 0x2 - Warning, consumed 80% of the reserved blocks
	// 0x3 - Urgent, consumed 90% of the reserved blocks
	ufsFailEolGlob = `.*(?P<param>bPreEOLInfo) = 0x(?P<val>\d)`
)

// detectMMCState read the info to detect state for MMC storage.
// return error if the regular expression cannot compile.
func detectMMCState(ctx context.Context, storageInfoSlice []string) (StorageState, error) {
	return detectJedecState(ctx, "MMC", mmcFailLifeGlob, mmcFailEolGlob, storageInfoSlice)
}

// detectUFSState read the info to detect state for UFS storage.
// return error if the regular expression cannot compile.
func detectUFSState(ctx context.Context, storageInfoSlice []string) (StorageState, error) {
	return detectJedecState(ctx, "UFS", ufsFailLifeGlob, ufsFailEolGlob, storageInfoSlice)
}

// detectJedecState read the info to detect state for UFS or MMC storage.
// return error if the regular expression cannot compile.
func detectJedecState(ctx context.Context, ifaceName, jedecFailLifeGlob, jedecFailEolGlob string, storageInfoSlice []string) (StorageState, error) {
	log.Infof(ctx, "Extraction metrics for "+ifaceName+" storage")
	jedecFailLiveRegexp, err := regexp.Compile(jedecFailLifeGlob)
	if err != nil {
		return StorageStateUndefined, errors.Annotate(err, "detect "+ifaceName+" state").Err()
	}
	jedecFailEolRegexp, err := regexp.Compile(jedecFailEolGlob)
	if err != nil {
		return StorageStateUndefined, errors.Annotate(err, "detect "+ifaceName+" state").Err()
	}
	eolValue := 0
	lifeValue := -1
	for _, line := range storageInfoSlice {
		mLife, err := regexpSubmatchToMap(jedecFailLiveRegexp, line)
		if err == nil {
			param := mLife["val"]
			log.Debugf(ctx, "Found line for lifetime estimate => %q", line)
			var val int
			if param == "a" {
				val = 100
			} else if param == "b" {
				val = 101
			} else {
				parsedVal, parseIntErr := strconv.ParseInt(param, 10, 64)
				if parseIntErr != nil {
					log.Errorf(ctx, parseIntErr.Error())
				}
				val = int(parsedVal * 10)
			}
			if val > lifeValue {
				lifeValue = val
			}
			continue
		}
		mEol, err := regexpSubmatchToMap(jedecFailEolRegexp, line)
		if err == nil {
			param := mEol["val"]
			log.Debugf(ctx, "Found line for end-of-life => %q", line)
			parsedVal, parseIntErr := strconv.ParseInt(param, 10, 64)
			if parseIntErr != nil {
				log.Errorf(ctx, parseIntErr.Error())
			}
			eolValue = int(parsedVal)
			break
		}
	}
	// set state based on end-of-life
	if eolValue == 3 {
		return StorageStateCritical, nil
	} else if eolValue == 2 {
		return StorageStateWarning, nil
	} else if eolValue == 1 {
		return StorageStateNormal, nil
	}
	// set state based on life of estimates
	if lifeValue < 90 {
		return StorageStateNormal, nil
	} else if lifeValue < 100 {
		return StorageStateWarning, nil
	}
	return StorageStateCritical, nil
}

const (
	// Ex "Percentage Used:         100%"
	nvmeFailGlob = `Percentage Used:\s+(?P<param>(\d{1,3}))%`
)

// detectNVMEState read the info to detect state for NVMe storage.
// return error if the regular expression cannot compile
func detectNVMEState(ctx context.Context, storageInfoSlice []string) (StorageState, error) {
	log.Infof(ctx, "Extraction metrics for NVMe storage")
	nvmeFailRegexp, err := regexp.Compile(nvmeFailGlob)
	if err != nil {
		return StorageStateUndefined, errors.Annotate(err, "detect nvme state").Err()
	}
	var usedValue int = -1
	for _, line := range storageInfoSlice {
		m, err := regexpSubmatchToMap(nvmeFailRegexp, line)
		if err == nil {
			log.Debugf(ctx, "Found line for usage => %q", line)
			val, convertErr := strconv.ParseInt(m["param"], 10, 64)
			if convertErr == nil {
				usedValue = int(val)
			} else {
				log.Debugf(ctx, "Could not cast: %s to int", m["param"])
			}
			break
		}
	}
	if usedValue < 91 {
		log.Infof(ctx, "NVME storage usage value: %v", usedValue)
		return StorageStateNormal, nil
	}
	return StorageStateWarning, nil
}

// regexpSubmatchToMap takes pattern of regex and the source string and returns
// the map containing the groups defined in the regex expression.
// Assumes the pattern can compile.
// return error if it doesn't find any match
func regexpSubmatchToMap(r *regexp.Regexp, source string) (map[string]string, error) {
	m := make(map[string]string)
	matches := r.FindStringSubmatch(source)
	if len(matches) < 1 {
		return m, errors.Reason("regexp submatch to map: no match found").Err()
	}
	// there is at least 1 match found
	names := r.SubexpNames()
	for i := range names {
		if i != 0 {
			m[names[i]] = matches[i]
		}
	}
	return m, nil
}

const (
	readStorageInfoCMD = ". /usr/share/misc/storage-info-common.sh; get_storage_info"
)

// storageStateMap maps state from storageState type to tlw.HardwareState type
var storageStateMap = map[StorageState]tlw.HardwareState{
	StorageStateNormal:    tlw.HardwareState_HARDWARE_NORMAL,
	StorageStateWarning:   tlw.HardwareState_HARDWARE_ACCEPTABLE,
	StorageStateCritical:  tlw.HardwareState_HARDWARE_NEED_REPLACEMENT,
	StorageStateUndefined: tlw.HardwareState_HARDWARE_UNSPECIFIED,
}

// AuditStorageSMART checks the storage using it SMART capabilities,
// and mark the DUT for replacement if needed.
//
// This is a helper function to encapsulate the logic for SMART-check
// and intended to be called from the exec, as well as any other place
// where such a check is required.
func AuditStorageSMART(ctx context.Context, r components.Runner, storage *tlw.Storage, dut *tlw.Dut) error {
	if storage == nil {
		return errors.Reason("audit storage smart: data is not present in dut info").Err()
	}
	log.Debugf(ctx, "audit storage smart: initial storage state: %s, initial dut state: %s", storage.State, dut.State)
	defer func() {
		log.Debugf(ctx, "audit storage smart: final storage state: %s, final dut state: %s", storage.State, dut.State)
	}()
	rawOutput, err := r(ctx, time.Minute, readStorageInfoCMD)
	if err != nil {
		return errors.Annotate(err, "audit storage smart").Err()
	}
	ss, err := ParseSMARTInfo(ctx, rawOutput)
	if err != nil {
		return errors.Annotate(err, "audit storage smart").Err()
	}
	log.Debugf(ctx, "Detected storage type: %q", ss.StorageType)
	log.Debugf(ctx, "Detected storage state: %q", ss.StorageState)
	convertedHardwareState, ok := storageStateMap[ss.StorageState]
	if !ok {
		return errors.Reason("audit storage smart: cannot find corresponding hardware state match in the map").Err()
	}
	switch convertedHardwareState {
	case tlw.HardwareState_HARDWARE_UNSPECIFIED:
		return errors.Reason("audit storage smart: DUT storage did not detected or state cannot extracted").Err()
	case tlw.HardwareState_HARDWARE_NEED_REPLACEMENT:
		log.Debugf(ctx, "Detected issue with storage on the DUT")
		storage.State = tlw.HardwareState_HARDWARE_NEED_REPLACEMENT
		log.Debugf(ctx, "Setting the DUT state: %q", string(dutstate.NeedsReplacement))
		dut.State = dutstate.NeedsReplacement
		return errors.Reason("audit storage smart: hardware state need replacement").Err()
	default:
		log.Debugf(ctx, "New storage state: %q", convertedHardwareState)
		return nil
	}
}

// isItTimeToRunBadBlocksRO determines if it is time to run the RO
// badblock again on this device.
func isItTimeToRunBadBlocksRO(ctx context.Context, metrics metrics.Metrics) error {
	// TODO(b/242584223): complete the implementation of this function using the
	// information about when the last time RO bad-blocks was run for
	// this device. This is captured by bug b/242584223. The default
	// value of 'nil' returned here will cause the badblocks check to
	// always be run. This does not have any affect on the correctness
	// of badblocks audit, it just will execute the audit even when it
	// is not required.
	return nil
}

// isItTimeToRunBadBlocksRW determines if it is time to run the RW
// badblock again on this device.
func isItTimeToRunBadBlocksRW(ctx context.Context, metrics metrics.Metrics) error {
	// TODO(b/242584223): complete the implementation of this function using the
	// information about when the last time RW bad-blocks was run for
	// this device. This is captured by bug b/242584223. The default
	// value of the "not-implemented" error returned here will cause
	// the badblocks check to not be run. This will be changed once
	// the logic to check whether it is time to execute the bad blocks
	// is ready.
	return errors.Reason("is it time to run bad blocks: not implemented").Err()
}

// AuditMode represents the type of audit mode.
type AuditMode string

const (
	// "auto" represents that the type of bad-blocks check will be
	// determined automatically by the task.
	auditModeAuto AuditMode = "auto"
	// "not" means that the bad-blocks check will not be run at all. This
	// option is intended to make the functionality of recovery-lib (a.k.a
	// Paris) feature-complete w.r.t. legacy repair. However, unlike
	// legacy repair, this is not exercised in Paris because we do not run
	// badblocks check from repair at all.
	auditModeNot AuditMode = "not"
	// "rw" represents the read-write mode of bad-blocks check.
	auditModeRW AuditMode = "rw"
	// "ro" represents the read-only mode of bad-blocks check.
	auditModeRO AuditMode = "ro"
)

// Commands for RO and RW badblock execution.
var badBlockCommands = map[AuditMode]string{
	auditModeRW: "badblocks -e 100 -nsv -b 4096 %s",
	auditModeRO: "badblocks -e 100 -s -b 512 %s",
}

// runBadBlocksCheck executes the badblocks check on device.
func runBadBlocksCheck(ctx context.Context, run components.Runner, timeout time.Duration, mainStorage string, badBlocksMode AuditMode) (string, error) {
	if badBlockCommands[badBlocksMode] == "" {
		return "", errors.Reason("run bad blocks check: unknown badblocks mode %q", badBlocksMode).Err()
	}
	badBlocksCmd := fmt.Sprintf(badBlockCommands[badBlocksMode], mainStorage)
	log.Debugf(ctx, "Run Bad Blocks Check: executing command %q", badBlocksCmd)
	cmdResult, err := run(ctx, timeout, badBlocksCmd)
	if err != nil {
		return "", errors.Annotate(err, "run bad blocks check").Err()
	}
	if cmdResult != "" {
		// Following the logic from legacy repair, a non-empty result
		// from execution of badblocks is not good.
		log.Debugf(ctx, "Run Bad Blocks Check: non-empty result of badblocks command is %q", badBlocksCmd)
		return cmdResult, errors.Reason("run bad blocks check: badblocks output is non empty: %q", cmdResult).Err()
	}
	return "", nil
}

// auditModes represents the possible modes for badblocks execution on
// the DUT.
var auditModes = map[AuditMode]bool{
	auditModeAuto: true,
	auditModeNot:  true,
	auditModeRW:   true,
	auditModeRO:   true,
}

// BadBlocksArgs collects together all the parameters that are
// applicable for bad blocks execution.
type BadBlocksArgs struct {
	AuditMode AuditMode
	Run       components.Runner
	Storage   *tlw.Storage
	Dut       *tlw.Dut
	Metrics   metrics.Metrics
	TimeoutRW time.Duration
	TimeoutRO time.Duration
}

// CheckBadblocks executes the bad-blocks check on the storage device.
//
// It will also mark the DUT for replacement if required.
func CheckBadblocks(ctx context.Context, bbArgs *BadBlocksArgs) error {
	if !auditModes[bbArgs.AuditMode] {
		errors.Reason("check bad blocks: unknown audit mode %q", bbArgs.AuditMode).Err()
	}
	switch bbArgs.AuditMode {
	case auditModeNot:
		// As also mentioned in the comment that introduced this flag,
		// this option is intended to make the functionality of
		// recovery-lib (a.k.a Paris) feature-complete w.r.t. legacy
		// repair. However, unlike legacy repair, this is not
		// exercised in Paris because we do not run badblocks check
		// from repair at all.
		log.Debugf(ctx, "Check Bad Blocks: audit mode : %q: skipping badblocks.", bbArgs.AuditMode)
		return nil
	case auditModeAuto:
		// The mode "auto" means that we will determine the
		// appropriate mode for badblocks ourselves.
		if err := isItTimeToRunBadBlocksRO(ctx, bbArgs.Metrics); err == nil {
			bbArgs.AuditMode = auditModeRO
		} else {
			log.Debugf(ctx, "Check Bad Blocks: error while checking if it is time to run RO badblocks (non-critical): %s.", err)
		}
		// Here we might end up overwriting the AuditMode after also
		// setting it to "ro" above. This is intentional. The "rw"
		// badblocks is a stronger check, and if it is the right time
		// for it, it will supercede any "ro" checks.
		if err := isItTimeToRunBadBlocksRW(ctx, bbArgs.Metrics); err == nil {
			bbArgs.AuditMode = auditModeRW
		} else {
			log.Debugf(ctx, "Check Bad Blocks: error while checking if it is time to run RW badblocks (non-critical): %s.", err)
		}
	}
	log.Debugf(ctx, "Check Bad Blocks: the finalized audit mode is :%q", bbArgs.AuditMode)
	mainStorage, err := cros.DeviceMainStoragePath(ctx, bbArgs.Run)
	if err != nil {
		return errors.Annotate(err, "check bad blocks").Err()
	}
	if mainStorage == "" {
		log.Debugf(ctx, "Check Bad Blocks: path to main storage is empty, hence cannot run any type of badblocks check (non-critical).")
		// We return without error if the path of main storage device
		// is empty. This following the logic in legacy repair.
		return nil
	}
	// TODO(b/242570493): We shouldn't have to create a new logger to
	// pass to IsBootedFromExternalStorage. The method should be able
	// to obtain the logger from ctx. This will be addressed in
	// b/242570493. This function currently does not make use of the
	// logger. Hence we are passing a value of nil. When this bug is
	// resolve, we will not pass the logger at all.
	usbBootErr := cros.IsBootedFromExternalStorage(ctx, bbArgs.Run)
	if usbBootErr != nil {
		log.Debugf(ctx, "Check Bad Blocks: not booted from USB device, RW badblocks on main storage cannot be done even if selected (non-critical).")
	}
	if usbBootErr == nil && bbArgs.AuditMode == auditModeRW {
		if out, err := runBadBlocksCheck(ctx, bbArgs.Run, bbArgs.TimeoutRW, mainStorage, bbArgs.AuditMode); err != nil {
			if out != "" {
				bbArgs.Storage.State = tlw.HardwareState_HARDWARE_NEED_REPLACEMENT
				bbArgs.Dut.State = dutstate.NeedsReplacement
			}
			return errors.Annotate(err, "audit storage badblocks").Err()
		}
		// TODO(b/242584223): record this execution of badblocks for RW
		// as well as RO, since RO badblocks check is subset of RW
		// badblocks check. This is being tracked in b/242584223. This
		// does not have any affect on the correctness of badblocks
		// audit, it will only avoid any unnecessary execution.
	}
	if bbArgs.AuditMode == auditModeRO {
		if out, err := runBadBlocksCheck(ctx, bbArgs.Run, bbArgs.TimeoutRO, mainStorage, bbArgs.AuditMode); err != nil {
			if out != "" {
				bbArgs.Storage.State = tlw.HardwareState_HARDWARE_NEED_REPLACEMENT
				bbArgs.Dut.State = dutstate.NeedsReplacement
			}
			return errors.Annotate(err, "audit storage badblocks").Err()
		}
		// TODO(b/242584223): record this execution of badblocks for
		// RO. This is being tracked in b/242584223. This does not
		// have any affect on the correctness of badblocks audit, it
		// will only avoid any unnecessary execution.
	}
	return nil
}
