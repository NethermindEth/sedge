package services

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (s *serviceManager) ContainerId(service string) (string, error) {
	containers, err := s.dockerClient.ContainerList(context.Background(), types.ContainerListOptions{
		All:     true,
		Filters: filters.NewArgs(filters.Arg("name", service)),
	})
	if err != nil {
		return "", err
	}
	if len(containers) == 0 {
		return "", fmt.Errorf("%w: %s", ErrContainerNotFound, service)
	}
	if len(containers) > 1 {
		return "", fmt.Errorf("%w: %s", ErrMultipleContainers, service)
	}
	return containers[0].ID, nil
}
