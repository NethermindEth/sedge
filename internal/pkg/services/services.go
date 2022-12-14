package services

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
)

const (
	ServiceValidator = "validator"
)

var ServiceContainer map[string]string = map[string]string{
	ServiceValidator: "validator-client",
}

type ServiceManager interface {
	Image(service string) (string, error)
	Stop(service string) error
	Start(service string) error
	IsRunning(service string) (bool, error)
}

func NewServiceManager(dockerClient client.APIClient) ServiceManager {
	return &serviceManager{dockerClient: dockerClient}
}

type serviceManager struct {
	dockerClient client.APIClient
}

func (s *serviceManager) Image(serviceName string) (string, error) {
	info, err := ctInfo(s.dockerClient, serviceName)
	return info.Image, err
}

func (s *serviceManager) Stop(serviceName string) error {
	containerName, ok := ServiceContainer[serviceName]
	if !ok {
		return fmt.Errorf("unknown container name for service: %s", serviceName)
	}
	ctInfo, err := s.dockerClient.ContainerInspect(context.Background(), containerName)
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil
		}
		return err
	}
	if ctInfo.State.Running {
		log.Infof("stopping service: %s, currently on %s status", serviceName, ctInfo.State.Status)
		timeout := 5 * time.Minute
		if err := s.dockerClient.ContainerStop(context.Background(), containerName, &timeout); err != nil {
			return fmt.Errorf("error stopping service %s (with container %s): %w", serviceName, containerName, err)
		}
	}
	return nil
}

func (s *serviceManager) Start(serviceName string) error {
	containerName, ok := ServiceContainer[serviceName]
	if !ok {
		return fmt.Errorf("unknown container name for service: %s", serviceName)
	}
	if err := s.dockerClient.ContainerStart(context.Background(), containerName, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("error starting service %s (with container %s): %w", serviceName, containerName, err)
	}
	return nil
}

func (s *serviceManager) IsRunning(serviceName string) (bool, error) {
	info, err := ctInfo(s.dockerClient, serviceName)
	return info.State.Running, err
}
