package services

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

const (
	ServiceCtSlashingData    = "validator-slashing-data"
	ServiceCtValidator       = "validator-client"
	ServiceCtValidatorImport = "validator-import-client"
)

type ServiceManager interface {
	Image(service string) (string, error)
	Stop(service string) error
	Start(service string) error
	IsRunning(service string) (bool, error)
	Wait(service string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error)
	ContainerId(service string) (string, error)
}

func NewServiceManager(dockerClient client.APIClient) ServiceManager {
	return &serviceManager{dockerClient: dockerClient}
}

type serviceManager struct {
	dockerClient client.APIClient
}

func (s *serviceManager) Stop(service string) error {
	ctInfo, err := s.dockerClient.ContainerInspect(context.Background(), service)
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil
		}
		return err
	}
	if ctInfo.State.Running {
		log.Infof("stopping service: %s, currently on %s status", service, ctInfo.State.Status)
		timeout := 5 * time.Minute
		if err := s.dockerClient.ContainerStop(context.Background(), service, &timeout); err != nil {
			return fmt.Errorf("error stopping service %s: %w", service, err)
		}
	}
	return nil
}

func (s *serviceManager) Start(service string) error {
	if err := s.dockerClient.ContainerStart(context.Background(), service, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("error starting service %s (with container %s): %w", service, service, err)
	}
	return nil
}

func (s *serviceManager) IsRunning(ct string) (bool, error) {
	info, err := s.dockerClient.ContainerInspect(context.Background(), ct)
	return info.State.Running, err
}

func (s *serviceManager) Wait(service string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error) {
	return s.dockerClient.ContainerWait(context.Background(), service, condition)
}

func (s *serviceManager) ContainerId(service string) (string, error) {
	containers, err := s.dockerClient.ContainerList(context.Background(), types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", service)),
	})
	if err != nil {
		return "", err
	}
	if len(containers) == 0 {
		return "", fmt.Errorf("container %s not found", service)
	}
	if len(containers) > 1 {
		return "", fmt.Errorf("multiple containers with name %s", service)
	}
	return containers[0].ID, nil
}
