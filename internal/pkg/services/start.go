package services

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
)

func (s *serviceManager) Start(service string) error {
	if err := s.dockerClient.ContainerStart(context.Background(), service, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("%w %s: %s", ErrStartingContainer, service, err)
	}
	return nil
}
