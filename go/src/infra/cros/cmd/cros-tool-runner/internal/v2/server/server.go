// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc/codes"
	"infra/cros/cmd/cros-tool-runner/api"
	"infra/cros/cmd/cros-tool-runner/internal/v2/commands"
)

// ContainerServerImpl implements the gRPC services by running commands and
// mapping errors to proper gRPC status codes.
type ContainerServerImpl struct {
	api.UnimplementedCrosToolRunnerContainerServiceServer
	networks   []string // TODO(mingkong) use a map
	containers []string
}

// CreateNetwork creates a new docker network with the given name.
func (s *ContainerServerImpl) CreateNetwork(ctx context.Context, request *api.CreateNetworkRequest) (*api.CreateNetworkResponse, error) {
	cmd := commands.NetworkCreate{Name: request.Name}
	stdout, stderr, err := cmd.Execute(ctx)
	if stderr != "" {
		return nil, utils.toStatusError(stderr)
	}
	if err != nil {
		return nil, err
	}
	id := utils.firstLine(stdout)
	log.Println("success: created network", id)
	s.networks = append(s.networks, request.Name)
	return &api.CreateNetworkResponse{Network: &api.Network{Name: request.Name, Id: id, Owned: true}}, nil
}

// GetNetwork retrieves information of given docker network.
func (s *ContainerServerImpl) GetNetwork(ctx context.Context, request *api.GetNetworkRequest) (*api.GetNetworkResponse, error) {
	cmd := commands.NetworkInspect{Names: []string{request.Name}, Format: "{{.Id}}"}
	stdout, stderr, err := cmd.Execute(ctx)
	if stderr != "" {
		return nil, utils.toStatusError(stderr)
	}
	if err != nil {
		return nil, err
	}
	id := utils.firstLine(stdout)
	log.Println("success: found network", id)
	return &api.GetNetworkResponse{Network: &api.Network{Name: request.Name, Id: id, Owned: utils.contains(s.networks, request.Name)}}, nil
}

// Shutdown signals to shut down the CTRv2 gRPC server.
func (s *ContainerServerImpl) Shutdown(context.Context, *api.ShutdownRequest) (*api.ShutdownResponse, error) {
	log.Println("processing shutdown request")
	p, err := os.FindProcess(os.Getpid())
	if err == nil {
		p.Signal(os.Interrupt)
	}
	log.Println("interrupt signal sent")
	return &api.ShutdownResponse{}, err
}

// StartContainer pulls image and then calls docker run to start a container.
func (s *ContainerServerImpl) StartContainer(ctx context.Context, request *api.StartContainerRequest) (*api.StartContainerResponse, error) {
	if request.Name == "" {
		return nil, utils.invalidArgument("A unique name must be specified")
	}
	if request.ContainerImage == "" {
		return nil, utils.invalidArgument("A container image must be specified")
	}
	// TODO(mingkong): define behavior of existing name in containers[]
	if request.StartCommand == nil || len(request.StartCommand) == 0 {
		return nil, utils.invalidArgument("A start command must be specified")
	}
	if request.AdditionalOptions != nil {
		options := request.AdditionalOptions
		if options.Expose != nil && (len(options.Expose) > 1 || strings.Contains(options.Expose[0], "-")) {
			return nil, utils.unimplemented("Exposing multiple ports are not supported")
		}
	}
	pullErr := s.pullImage(ctx, request.ContainerImage)
	if pullErr != nil {
		return nil, pullErr
	}

	cmd := commands.DockerRun{StartContainerRequest: request}
	stdout, stderr, err := cmd.Execute(ctx)
	if stderr != "" {
		return nil, utils.toStatusErrorWithMapper(stderr, func(s string) codes.Code {
			switch {
			case strings.Contains(s, fmt.Sprintf("The container name \"/%s\" is already in use", request.Name)):
				return codes.AlreadyExists
			default:
				return codes.Unknown
			}
		})
	}
	if err != nil {
		return nil, err
	}
	// TODO(mingkong): handle edge case where id is returned but container has immediately stopped: e.g. cros-dut cannot connect to dut
	id := utils.firstLine(stdout)
	log.Println("success: started container", id)
	s.containers = append(s.containers, request.Name)
	return &api.StartContainerResponse{Container: &api.Container{Name: request.Name, Id: id, Owned: true}}, nil
}

// pullImage pulls docker image and handles error mapping specifically
func (s *ContainerServerImpl) pullImage(ctx context.Context, image string) error {
	pullCmd := commands.DockerPull{ContainerImage: image}
	_, stderr, _ := pullCmd.Execute(ctx)
	if stderr != "" {
		return utils.toStatusErrorWithMapper(stderr, func(s string) codes.Code {
			switch {
			case strings.Contains(s, "Permission \"artifactregistry.repositories.downloadArtifacts\" denied on resource"):
				return codes.PermissionDenied
			case strings.Contains(s, fmt.Sprintf("manifest for %s not found", image)):
				return codes.NotFound
			default:
				return codes.Unknown
			}
		})
	}
	return nil
}

// stopContainers removes containers that are owned by current CTRv2 service in the reverse order of how they are started.
func (s *ContainerServerImpl) stopContainers() {
	if len(s.containers) == 0 {
		log.Println("no containers to clean up")
		return
	}
	log.Println("stopping containers")
	cmd := commands.ContainerStop{Names: utils.reverse(s.containers)}
	stdout, stderr, _ := cmd.Execute(context.Background())
	if stdout != "" {
		log.Println("received stdout:", stdout)
	}
	if stderr != "" {
		log.Println("received stderr", stderr)
	}
	// TODO(mingkong) define the behavior of stop container error.
	s.containers = make([]string, 0)
}

// removeNetworks removes networks that were created by current CTRv2 service.
func (s *ContainerServerImpl) removeNetworks() {
	if len(s.networks) == 0 {
		log.Println("no networks to clean up")
		return
	}
	log.Println("removing networks")
	cmd := commands.NetworkRemove{Names: s.networks}
	stdout, stderr, _ := cmd.Execute(context.Background())
	if stdout != "" {
		log.Println("received stdout:", stdout)
	}
	if stderr != "" {
		log.Println("received stderr", stderr)
	}
	// TODO(mingkong) define the behavior of remove network error.
	s.networks = make([]string, 0)
}

// cleanup removes containers and networks in order to allow graceful shutdown of the CTRv2 service.
func (s *ContainerServerImpl) cleanup() {
	s.stopContainers()
	s.removeNetworks()
}