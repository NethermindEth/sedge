package services

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ctInfo(dockerClient client.APIClient, serviceName string) (types.ContainerJSON, error) {
	containerName, ok := ServiceContainer[serviceName]
	if !ok {
		return types.ContainerJSON{}, fmt.Errorf("unknown container name for service: %s", serviceName)
	}
	return dockerClient.ContainerInspect(context.Background(), containerName)
}
