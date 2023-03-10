package controller

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/gae/service/datastore"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	ufspb "infra/unifiedfleet/api/v1/models"
	chromeosLab "infra/unifiedfleet/api/v1/models/chromeos/lab"
	ufsds "infra/unifiedfleet/app/model/datastore"
	"infra/unifiedfleet/app/model/inventory"
	"infra/unifiedfleet/app/model/registration"
	"infra/unifiedfleet/app/util"
)

// deployLabstationMaskPaths contains paths for which deploy task if required.
var deployLabstationMaskPaths = []string{
	"machines",
	"labstation.rpm.host",
	"labstation.rpm.outlet",
}

// CreateLabstation creates a new labstation entry in UFS.
func CreateLabstation(ctx context.Context, lse *ufspb.MachineLSE) (*ufspb.MachineLSE, error) {
	f := func(ctx context.Context) error {
		hc := getHostHistoryClient(lse)

		// Get machine to get zone and rack info for machinelse table indexing
		machine, err := GetMachine(ctx, lse.GetMachines()[0])
		if err != nil {
			return errors.Annotate(err, "unable to get machine %s", lse.GetMachines()[0]).Err()
		}

		// Validate input
		if err := validateCreateMachineLSE(ctx, lse, nil, machine); err != nil {
			return errors.Annotate(err, "Validation error - Failed to create labstation").Err()
		}

		//Copy for logging
		oldMachine := proto.Clone(machine).(*ufspb.Machine)

		machine.ResourceState = ufspb.State_STATE_SERVING
		// Fill the rack/zone OUTPUT only fields for indexing machinelse table/vm table
		setOutputField(ctx, machine, lse)

		// Create the machinelse
		if _, err := registration.BatchUpdateMachines(ctx, []*ufspb.Machine{machine}); err != nil {
			return errors.Annotate(err, "Fail to update machine %s", machine.GetName()).Err()
		}
		hc.LogMachineChanges(oldMachine, machine)
		lse.ResourceState = ufspb.State_STATE_REGISTERED
		if _, err := inventory.BatchUpdateMachineLSEs(ctx, []*ufspb.MachineLSE{lse}); err != nil {
			return errors.Annotate(err, "Failed to BatchUpdate MachineLSEs %s", lse.Name).Err()
		}

		if err := hc.stUdt.addLseStateHelper(ctx, lse, machine); err != nil {
			return errors.Annotate(err, "Fail to update host state").Err()
		}
		hc.LogMachineLSEChanges(nil, lse)
		return hc.SaveChangeEvents(ctx)
	}
	if err := datastore.RunInTransaction(ctx, f, nil); err != nil {
		logging.Errorf(ctx, "Failed to create machinelse in datastore: %s", err)
		return nil, err
	}
	return lse, nil
}

// UpdateLabstation validates and updates the given labstation machine LSE.
func UpdateLabstation(ctx context.Context, machinelse *ufspb.MachineLSE, mask *field_mask.FieldMask) (*ufspb.MachineLSE, error) {
	f := func(ctx context.Context) error {

		hc := getHostHistoryClient(machinelse)

		// Get the existing MachineLSE(Labstation).
		oldMachinelse, err := inventory.GetMachineLSE(ctx, machinelse.GetName())
		if err != nil {
			return errors.Annotate(err, "Failed to get existing Labstation").Err()
		}
		// Validate the input. Not passing the update mask as there is a different validation for labstation.
		if err := validateUpdateMachineLSE(ctx, oldMachinelse, machinelse, nil); err != nil {
			return errors.Annotate(err, "Validation error - Failed to update ChromeOSMachineLSEDUT").Err()
		}

		// Validate the input from labstation perspective
		if err := validateUpdateLabstation(ctx, oldMachinelse, machinelse, mask); err != nil {
			return errors.Annotate(err, "Validation error - Failed to update labstation").Err()
		}

		// Assign hostname to the labstation.
		machinelse.GetChromeosMachineLse().GetDeviceLse().GetLabstation().Hostname = machinelse.GetHostname()
		// Copy output only fields.
		machinelse.Zone = oldMachinelse.GetZone()
		machinelse.Rack = oldMachinelse.GetRack()

		var machine *ufspb.Machine

		// Validate the update mask and process it.
		if mask != nil && len(mask.Paths) > 0 {
			// Partial update with mask
			if err := validateUpdateLabstationMask(mask, machinelse); err != nil {
				return errors.Annotate(err, "UpdateLabstation - Failed update mask validation").Err()
			}
			if machinelse, err = processUpdateLabstationMask(ctx, proto.Clone(oldMachinelse).(*ufspb.MachineLSE), machinelse, mask); err != nil {
				return errors.Annotate(err, "UpdateLabstation - Failed to process update mask").Err()
			}
		} else {
			// Full update, Machines cannot be empty.
			if len(machinelse.GetMachines()) > 0 {
				if machinelse.GetMachines()[0] != oldMachinelse.GetMachines()[0] {
					if machine, err = GetMachine(ctx, machinelse.GetMachines()[0]); err != nil {
						return err
					}
					// Check if we have permission for the new machine.
					if err := util.CheckPermission(ctx, util.InventoriesUpdate, machine.GetRealm()); err != nil {
						return err
					}
					// Machine was changed. Copy zone and rack info into machinelse.
					setOutputField(ctx, machine, machinelse)
				}
			} else {
				// Empty machines field, Invalid update.
				return status.Error(codes.InvalidArgument, "UpdateLabstation - machines field cannot be empty/nil.")
			}
			// Copy old state if state was not updated.
			if machinelse.GetResourceState() == ufspb.State_STATE_UNSPECIFIED {
				machinelse.ResourceState = oldMachinelse.GetResourceState()
			}
		}

		_, err = inventory.BatchUpdateMachineLSEs(ctx, []*ufspb.MachineLSE{machinelse})
		if err != nil {
			logging.Errorf(ctx, "Failed to BatchUpdate ChromeOSMachineLSEDUTs %s", err)
			return err
		}
		hc.LogMachineLSEChanges(oldMachinelse, machinelse)

		// Update state for the labstation.
		if err := hc.stUdt.updateStateHelper(ctx, machinelse.GetResourceState()); err != nil {
			return err
		}
		return hc.SaveChangeEvents(ctx)
	}
	if err := datastore.RunInTransaction(ctx, f, nil); err != nil {
		logging.Errorf(ctx, "Failed to update MachineLSE DUT in datastore: %s", err)
		return nil, errors.Annotate(err, "UpdateLabstation - Failed to update").Err()
	}
	return machinelse, nil
}

// validateUpdateLabstationMask validates the labstation update mask.
func validateUpdateLabstationMask(mask *field_mask.FieldMask, machinelse *ufspb.MachineLSE) error {
	// GetLabstation should return an object. Otherwise UpdateLabstation isn't called
	labstation := machinelse.GetChromeosMachineLse().GetDeviceLse().GetLabstation()
	rpm := labstation.GetRpm()
	if rpm == nil {
		// Assign an empty rpm to avoid segfaults.
		rpm = &chromeosLab.OSRPM{}
	}

	maskSet := make(map[string]struct{}) // Set of all the masks
	for _, path := range mask.Paths {
		maskSet[path] = struct{}{}
	}
	// validate the give field mask
	for _, path := range mask.Paths {
		switch path {
		case "name":
			return status.Error(codes.InvalidArgument, "validateUpdateMachineLSELabstationUpdateMask - name cannot be updated, delete and create a new machinelse instead.")
		case "update_time":
			return status.Error(codes.InvalidArgument, "validateUpdateMachineLSELabstationUpdateMask - update_time cannot be updated, it is a output only field.")
		case "machines":
			if machinelse.GetMachines() == nil || len(machinelse.GetMachines()) == 0 || machinelse.GetMachines()[0] == "" {
				return status.Error(codes.InvalidArgument, "machines field cannot be empty/nil.")
			}
		case "labstation.hostname":
			return status.Error(codes.InvalidArgument, "validateUpdateMachineLSELabstationUpdateMask - hostname cannot be updated, delete and create a new dut.")
		case "labstation.rpm.host":
			// Check for deletion of the host. Outlet cannot be updated if host is deleted.
			if _, ok := maskSet["labstation.rpm.outlet"]; ok && rpm.GetPowerunitName() == "" && rpm.GetPowerunitOutlet() != "" {
				return status.Error(codes.InvalidArgument, "validateUpdateMachineLSELabstationUpdateMask - Deleting rpm host deletes everything. Cannot update outlet.")
			}
		case "labstation.rpm.outlet":
			// Check for deletion of rpm outlet. This should not be possible without deleting the host.
			if _, ok := maskSet["labstation.rpm.host"]; rpm.GetPowerunitOutlet() == "" && (!ok || (ok && rpm.GetPowerunitName() != "")) {
				return status.Error(codes.InvalidArgument, "validateUpdateMachineLSELabstationUpdateMask - Cannot remove rpm outlet. Please delete rpm.")
			}
		case "deploymentTicket":
		case "tags":
		case "description":
		case "resourceState":
		case "labstation.pools":
			// valid fields, nothing to validate.
		default:
			return status.Errorf(codes.InvalidArgument, "validateUpdateMachineLSELabstationUpdateMask - unsupported update mask path %q", path)
		}
	}
	return nil
}

// processUpdateLabstationMask processes the update mask provided for the labstation and returns oldMachineLSE updated.
func processUpdateLabstationMask(ctx context.Context, oldMachineLSE, newMachineLSE *ufspb.MachineLSE, mask *field_mask.FieldMask) (*ufspb.MachineLSE, error) {
	oldLabstation := oldMachineLSE.GetChromeosMachineLse().GetDeviceLse().GetLabstation()
	newLabstation := newMachineLSE.GetChromeosMachineLse().GetDeviceLse().GetLabstation()
	for _, path := range mask.Paths {
		switch path {
		case "machines":
			if len(newMachineLSE.GetMachines()) == 0 || newMachineLSE.GetMachines()[0] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "Cannot delete asset connected to %s", oldMachineLSE.GetName())
			}
			// Get machine to get zone and rack info for machinelse table indexing
			machine, err := GetMachine(ctx, newMachineLSE.GetMachines()[0])
			if err != nil {
				return oldMachineLSE, errors.Annotate(err, "Unable to get machine %s", newMachineLSE.GetMachines()[0]).Err()
			}
			// Check if we have permission for the new machine.
			if err := util.CheckPermission(ctx, util.InventoriesUpdate, machine.GetRealm()); err != nil {
				return nil, err
			}
			oldMachineLSE.Machines = newMachineLSE.GetMachines()
			setOutputField(ctx, machine, oldMachineLSE)
		case "resourceState":
			// Avoid setting state to unspecified.
			if newMachineLSE.GetResourceState() != ufspb.State_STATE_UNSPECIFIED {
				oldMachineLSE.ResourceState = newMachineLSE.GetResourceState()
			}
		case "tags":
			if tags := newMachineLSE.GetTags(); len(tags) > 0 {
				// Regular tag updates append to the existing tags.
				oldMachineLSE.Tags = mergeTags(oldMachineLSE.GetTags(), newMachineLSE.GetTags())
			} else {
				// Updating tags without any input clears the tags.
				oldMachineLSE.Tags = nil
			}
		case "description":
			oldMachineLSE.Description = newMachineLSE.Description
		case "deploymentTicket":
			oldMachineLSE.DeploymentTicket = newMachineLSE.GetDeploymentTicket()
		case "labstation.pools":
			// Append/Clear the pools given.
			if len(newLabstation.GetPools()) > 0 {
				oldLabstation.Pools = util.AppendUniqueStrings(oldLabstation.GetPools(), newLabstation.GetPools()...)
			} else {
				// Clear all the pools assigned if nothing is given.
				oldLabstation.Pools = nil
			}
		case "labstation.rpm.host":
			if newLabstation.GetRpm() == nil || newLabstation.GetRpm().GetPowerunitName() == "" {
				// Ensure that outlet is not being updated when deleting RPM.
				if util.ContainsAnyStrings(mask.Paths, "labstation.rpm.outlet") && newLabstation.GetRpm() != nil && newLabstation.GetRpm().GetPowerunitOutlet() != "" {
					return nil, status.Errorf(codes.InvalidArgument, "Cannot delete RPM and update outlet to %s", newLabstation.GetRpm().GetPowerunitOutlet())
				}
				oldLabstation.Rpm = nil
			} else {
				if oldLabstation.Rpm == nil {
					oldLabstation.Rpm = &chromeosLab.OSRPM{}
				}
				oldLabstation.GetRpm().PowerunitName = newLabstation.GetRpm().GetPowerunitName()
			}
		case "labstation.rpm.outlet":
			if newLabstation.GetRpm() == nil || newLabstation.GetRpm().GetPowerunitOutlet() == "" {
				// Ensure host is being cleared if the outlet is cleared
				if util.ContainsAnyStrings(mask.Paths, "labstation.rpm.host") && newLabstation.GetRpm() != nil && newLabstation.GetRpm().GetPowerunitName() != "" {
					return nil, status.Errorf(codes.InvalidArgument, "Cannot update RPM to %s and delete outlet", newLabstation.GetRpm().GetPowerunitName())
				}
				// Delete rpm
				oldLabstation.Rpm = nil
			} else {
				if oldLabstation.Rpm == nil {
					oldLabstation.Rpm = &chromeosLab.OSRPM{}
				}
				// Copy the outlet for update
				oldLabstation.GetRpm().PowerunitOutlet = newLabstation.GetRpm().GetPowerunitOutlet()
			}
		default:
			// Ideally, this piece of code should never execute unless validation is wrong.
			return nil, status.Errorf(codes.Internal, "Unable to process update mask for %s", path)
		}
	}
	return oldMachineLSE, nil
}

// validateUpdateLabstation checks if servos are being updated on full update or if the MLSE is a labstation
func validateUpdateLabstation(ctx context.Context, oldLabstation, labstation *ufspb.MachineLSE, mask *field_mask.FieldMask) error {
	// Check if labstation MachineLSE is updating any servo information
	// It is also not allowed to update the servo Hostname and servo Port of any servo.
	// Servo info is added/updated into Labstation only when a DUT(MachineLSE) is added/updated
	if labstation.GetChromeosMachineLse().GetDeviceLse().GetLabstation() != nil {
		if mask == nil || len(mask.Paths) == 0 {
			// Full update. Avoid changing labstation servo data.
			newServos := labstation.GetChromeosMachineLse().GetDeviceLse().GetLabstation().GetServos()
			if oldLabstation != nil {
				existingServos := oldLabstation.GetChromeosMachineLse().GetDeviceLse().GetLabstation().GetServos()
				if !testServoEq(newServos, existingServos) {
					return status.Errorf(codes.FailedPrecondition, "Servos are not allowed to be updated in redeploying labstations")
				}
			}
		}
		return nil
	}
	return status.Errorf(codes.FailedPrecondition, "Not a valid labstation")
}

// renameLabstation renames the labstation with the given name. Use inside a transaction
func renameLabstation(ctx context.Context, oldName, newName string, lse *ufspb.MachineLSE, machine *ufspb.Machine) (*ufspb.MachineLSE, error) {
	//Get all the duts referencing the old labstation
	dutMachinelses, err := getDUTsConnectedToLabstation(ctx, oldName)
	if err != nil {
		return nil, err
	}
	if err := validateRenameLabstation(ctx, lse, dutMachinelses); err != nil {
		return nil, err
	}
	oldLse := proto.Clone(lse).(*ufspb.MachineLSE)
	hc := getHostHistoryClient(lse)
	// Delete the old host record
	if err := inventory.DeleteMachineLSE(ctx, oldName); err != nil {
		return nil, err
	}
	// Delete old state record for host. Avoid deleting machine state.
	if err := hc.stUdt.deleteLseStateHelper(ctx, lse, nil); err != nil {
		return nil, errors.Annotate(err, "Fail to delete lse-related states").Err()
	}
	// Update the host name in all labstation refs
	lse.Name = newName
	lse.Hostname = newName
	lse.GetChromeosMachineLse().GetDeviceLse().GetLabstation().Hostname = newName
	for _, servo := range lse.GetChromeosMachineLse().GetDeviceLse().GetLabstation().GetServos() {
		servo.ServoHostname = newName
	}
	// Update the duts connected to the servos
	if err := updateServoHostForDUTs(ctx, lse, dutMachinelses); err != nil {
		return nil, err
	}
	// Update states
	if err := hc.stUdt.addLseStateHelper(ctx, lse, machine); err != nil {
		return nil, err
	}
	hc.LogMachineLSEChanges(oldLse, lse)
	// Update all the MLSEs affected.
	_, err = inventory.BatchUpdateMachineLSEs(ctx, []*ufspb.MachineLSE{lse})
	if err != nil {
		return nil, errors.Annotate(err, "Failed to add MachineLSE").Err()
	}
	// Save all changes to history
	if err := hc.SaveChangeEvents(ctx); err != nil {
		return nil, errors.Annotate(err, "Failed to save history").Err()
	}
	return lse, nil
}

// getDUTsConnectedToLabstation returns a list of DUTs whose servo hostname is the given labstation
func getDUTsConnectedToLabstation(ctx context.Context, labstation string) ([]*ufspb.MachineLSE, error) {
	// As the DUTs are indexed by servo_id which is a concatenation of
	// <host><port> strings we do a ranged query for everything greater
	// than or equal to <labstation>9000 (min port) and less than or equal
	// to <labstation>9999 (max port)
	dutMachinelses, err := inventory.RangedQueryMachineLSEByPropertyName(ctx,
		"servo_id", ufsds.GetServoID(labstation, servoPortMin),
		ufsds.GetServoID(labstation, servoPortMax), false)
	if err != nil {
		return nil, errors.Annotate(err, "getDUTsConnectedToLabstation - %s", labstation).Err()
	}
	duts := make([]*ufspb.MachineLSE, 0, len(dutMachinelses))
	for _, d := range dutMachinelses {
		// Filter out the DUTs connected to the labstation as the ranged query might include unexpected duts.
		shost := d.GetChromeosMachineLse().GetDeviceLse().GetDut().GetPeripherals().GetServo().ServoHostname
		if shost == labstation {
			duts = append(duts, d)
		}
	}
	return duts, nil
}

// validateRenameLabstation checks if we have the correct permissions to update the duts connected to the labstation.
func validateRenameLabstation(ctx context.Context, labstation *ufspb.MachineLSE, duts []*ufspb.MachineLSE) error {
	for _, dut := range duts {
		machine, err := registration.GetMachine(ctx, dut.GetMachines()[0])
		if err != nil {
			return errors.Annotate(err, "unable to get machine %s. Misconfigured host?", dut.GetMachines()[0]).Err()
		}
		if err := util.CheckPermission(ctx, util.InventoriesUpdate, machine.GetRealm()); err != nil {
			return status.Errorf(codes.PermissionDenied, fmt.Sprintf("Need update permission for %s. It's connected to %s", dut.GetName(), labstation.GetName()))
		}
	}
	return nil
}

// updateServohostForDUTs changes the DUTs servoHostname to the given labstation. Run inside a transaction
func updateServoHostForDUTs(ctx context.Context, labstation *ufspb.MachineLSE, duts []*ufspb.MachineLSE) error {
	// Write to all the associated duts history.
	var hcs []*HistoryClient
	// Update the servo hostname for connected duts
	for _, dut := range duts {
		dutHc := getHostHistoryClient(dut)
		// Make a copy of the dut for logging
		dutCpy := proto.Clone(dut).(*ufspb.MachineLSE)
		// Update the servo host for the connected hosts.
		dut.GetChromeosMachineLse().GetDeviceLse().GetDut().GetPeripherals().GetServo().ServoHostname = labstation.GetName()
		// Validate the update to dut
		if err := validateUpdateMachineLSE(ctx, dutCpy, dut, nil); err != nil {
			return errors.Annotate(err, "Failed to rename labstation").Err()
		}
		dutHc.LogMachineLSEChanges(dutCpy, dut)
		hcs = append(hcs, dutHc)
	}
	if _, err := inventory.BatchUpdateMachineLSEs(ctx, duts); err != nil {
		return errors.Annotate(err, "Failed to update duts.").Err()
	}
	for _, hhc := range hcs {
		if err := hhc.SaveChangeEvents(ctx); err != nil {
			return errors.Annotate(err, "Failed to save history").Err()
		}
	}
	return nil
}
