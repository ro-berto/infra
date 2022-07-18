package templates

import (
	"strings"
	"testing"

	"infra/cros/cmd/cros-tool-runner/api"
)

func TestCrosTestPopulate(t *testing.T) {
	processor := newCrosTestProcessor()
	request := &api.StartTemplatedContainerRequest{
		Name:           "my-container",
		ContainerImage: "gcr.io/image:123",
		Template: &api.Template{
			Container: &api.Template_CrosTest{
				CrosTest: &api.CrosTestTemplate{
					Network:     "mynet",
					ArtifactDir: "/tmp",
				}}}}

	convertedRequest, err := processor.Process(request)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	check(t, convertedRequest.Name, request.Name)
	check(t, convertedRequest.ContainerImage, request.ContainerImage)
	check(t, convertedRequest.AdditionalOptions.Network, "mynet")
	check(t, convertedRequest.AdditionalOptions.Expose[0], "8001")
	check(t, len(convertedRequest.AdditionalOptions.Volume), 2)
	if !strings.Contains(strings.Join(convertedRequest.StartCommand, " "), "cros-test") {
		t.Fatalf("cros-test is not part of start command")
	}
}