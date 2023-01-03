package services

import (
	"context"

	"github.com/docker/docker/api/types/container"
)

func (s *serviceManager) Wait(service string, condition container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error) {
	return s.dockerClient.ContainerWait(context.Background(), service, condition)
}
