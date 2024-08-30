package monitoring

import "github.com/NethermindEth/sedge/internal/monitoring/common"

// DockerManager is an interface for managing Docker containers.
type DockerServiceManager interface {
	// ContainerStatus returns the status of a container.
	ContainerStatus(container string) (common.Status, error)

	// ContainerIP returns the IP address of the container.
	ContainerIP(container string) (string, error)

	// ContainerNetworks returns the networks of a container.
	ContainerNetworks(container string) ([]string, error)

	// NetworkConnect connects a container to a network.
	NetworkConnect(container, network string) error

	// NetworkDisconnect disconnects a container from a network.
	NetworkDisconnect(container, network string) error
}
