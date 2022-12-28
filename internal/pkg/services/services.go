package services

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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
