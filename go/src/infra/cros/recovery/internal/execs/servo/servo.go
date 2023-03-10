// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package servo

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"go.chromium.org/luci/common/errors"

	"infra/cros/recovery/internal/components"
	"infra/cros/recovery/internal/components/servo"
	"infra/cros/recovery/internal/execs"
	"infra/cros/recovery/internal/execs/servo/topology"
	"infra/cros/recovery/internal/log"
	"infra/cros/recovery/tlw"
)

const (
	// This is the servod control for obtaining ppdut5 bus voltage in
	// millivolts.
	servodPPDut5Cmd = "ppdut5_mv"
)

// GetUSBDrivePathOnDut finds and returns the path of USB drive on a DUT.
func GetUSBDrivePathOnDut(ctx context.Context, run components.Runner, s components.Servod) (string, error) {
	// switch USB on servo multiplexer to the DUT-side
	if err := s.Set(ctx, "image_usbkey_direction", "dut_sees_usbkey"); err != nil {
		return "", errors.Annotate(err, "get usb drive path on dut: could not switch USB to DUT").Err()
	}
	// A detection delay is required when attaching this USB drive to DUT
	time.Sleep(usbDetectionDelay * time.Second)
	if out, err := run(ctx, time.Minute, "ls /dev/sd[a-z]"); err != nil {
		return "", errors.Annotate(err, "get usb drive path on dut").Err()
	} else {
		for _, p := range strings.Split(out, "\n") {
			dtOut, dtErr := run(ctx, time.Minute, fmt.Sprintf(". /usr/share/misc/chromeos-common.sh; get_device_type %s", p))
			if dtErr != nil {
				return "", errors.Annotate(dtErr, "get usb drive path on dut: could not check %q", p).Err()
			}
			if dtOut == "USB" {
				if _, fErr := run(ctx, time.Minute, fmt.Sprintf("fdisk -l %s", p)); fErr == nil {
					return p, nil
				} else {
					log.Debugf(ctx, "Get USB-drive path on dut: checked candidate usb drive path %q and found it incorrect.", p)
				}
			}
		}
		log.Debugf(ctx, "Get USB-drive path on dut: did not find any valid USB drive path on the DUT.")
	}
	return "", errors.Reason("get usb drive path on dut: did not find any USB Drive connected to the DUT as we checked that DUT is up").Err()
}

// IsContainerizedServoHost checks if the servohost is using servod container.
func IsContainerizedServoHost(ctx context.Context, servoHost *tlw.ServoHost) bool {
	if servoHost == nil || servoHost.ContainerName == "" {
		return false
	}
	log.Debugf(ctx, "Servo uses servod container with the name: %s", servoHost.ContainerName)
	return true
}

// WrappedServoType returns the type of servo device.
//
// This function first looks up the servo type using the servod
// control. If that does not work, it looks up the dut information for
// the servo host.
func WrappedServoType(ctx context.Context, info *execs.ExecInfo) (*servo.ServoType, error) {
	servoType, err := servo.GetServoType(ctx, info.NewServod())
	if err != nil {
		log.Debugf(ctx, "Wrapped Servo Type: Could not read the servo type from servod.")
		if st := info.GetChromeos().GetServo().GetServodType(); st != "" {
			servoType = servo.NewServoType(st)
		} else {
			return nil, errors.Reason("wrapped servo type: could not determine the servo type from servod control as well DUT Info.").Err()
		}
	}
	return servoType, nil
}

// ResetUsbkeyAuthorized resets usb-key detected under labstation.
//
// This is work around to address issue found for servo_v4p1.
// TODO(b/197647872): Remove as soon issue will be addressed.
func ResetUsbkeyAuthorized(ctx context.Context, run execs.Runner, servoSerial string, servoType string) error {
	if !strings.HasPrefix(servoSerial, "SERVOV4P1") {
		log.Debugf(ctx, "Authorized flag reset only for servo_v4p1.")
		return nil
	}
	log.Debugf(ctx, "Start reset authorized flag for servo_v4p1.")
	rootServoPath, err := topology.GetRootServoPath(ctx, run, servoSerial)
	if err != nil {
		return errors.Annotate(err, "reset usbkey authorized").Err()
	}
	pathDir := filepath.Dir(rootServoPath)
	pathTail := filepath.Base(rootServoPath)
	// For usb-path path looks like '/sys/bus/usb/devices/1-4.2.5' we need
	// remove last number, to make it as path to the servo-hub.
	pathTailElements := strings.Split(pathTail, ".")
	pathTail = strings.Join(pathTailElements[:(len(pathTailElements)-1)], ".")
	// Replace the first number '1' to '2 (usb3). Finally it will look like
	// '/sys/bus/usb/devices/2-4.2'
	pathTail = strings.Replace(pathTail, "1-", "2-", 1)
	const authorizedFlagName = "authorized"
	authorizedPath := filepath.Join(pathDir, pathTail, authorizedFlagName)
	log.Infof(ctx, "Authorized flag file path: %s", authorizedPath)
	// Setting flag to 0.
	if _, err := run(ctx, 30*time.Second, fmt.Sprintf("echo 0 > %s", authorizedPath)); err != nil {
		log.Debugf(ctx, `Attempt to reset %q flag to 0 for servo-hub failed`, authorizedFlagName)
		return errors.Annotate(err, "reset usbkey authorized: set to 0").Err()
	}
	time.Sleep(time.Second)
	// Setting flag to 1.
	if _, err := run(ctx, 30*time.Second, fmt.Sprintf("echo 1 > %s", authorizedPath)); err != nil {
		log.Debugf(ctx, `Attempt to reset %q flag to 1 for servo-hub failed`, authorizedFlagName)
		return errors.Annotate(err, "reset usbkey authorized: set to 1").Err()
	}
	time.Sleep(time.Second)
	log.Infof(ctx, "Attempt to reset %q succeed", authorizedFlagName)
	return nil
}
