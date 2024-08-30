package actions

import "github.com/docker/docker/api/types/container"

// DockerServiceManager is an interface for managing Docker containers.
type DockerServiceManager interface {
	Image(service string) (string, error)
	Stop(service string) error
	Start(service string) error
	IsRunning(service string) (bool, error)
	Wait(service string, condition container.WaitCondition) (<-chan container.WaitResponse, <-chan error)
	ContainerId(service string) (string, error)
	ContainerLogs(ctID, service string) (string, error)
}
