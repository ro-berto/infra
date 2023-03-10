package containers

import (
	"infra/cros/cmd/cros_test_runner/internal/interfaces"
	"infra/cros/cmd/cros_test_runner/internal/tools/crostoolrunner"
)

func NewCrosDutTemplatedContainer(
	containerImage string,
	ctr *crostoolrunner.CrosToolRunner) interfaces.ContainerInterface {

	return NewContainer(CrosDutTemplatedContainerType, "cros-dut", containerImage, ctr, true)
}

func NewCrosProvisionTemplatedContainer(
	containerImage string,
	ctr *crostoolrunner.CrosToolRunner) interfaces.ContainerInterface {

	return NewContainer(CrosProvisionTemplatedContainerType, "cros-provision", containerImage, ctr, true)
}

func NewCrosTestTemplatedContainer(
	containerImage string,
	ctr *crostoolrunner.CrosToolRunner) interfaces.ContainerInterface {

	return NewContainer(CrosTestTemplatedContainerType, "cros-test", containerImage, ctr, true)
}

func NewCrosPublishTemplatedContainer(
	contType interfaces.ContainerType,
	containerImage string,
	ctr *crostoolrunner.CrosToolRunner) interfaces.ContainerInterface {

	if contType != CrosGcsPublishTemplatedContainerType && contType != CrosTkoPublishTemplatedContainerType && contType != CrosRdbPublishTemplatedContainerType {
		return nil
	}
	return NewContainer(contType, "cros-publish", containerImage, ctr, true)
}
