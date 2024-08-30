package monitoring

import (
	"net"

	"github.com/NethermindEth/sedge/internal/monitoring/monitoring/services/types"
)

// ServiceAPI defines the interface for managing a monitoring service. It provides methods for
// adding and removing targets, retrieving environment variables, setting up the service, and initializing the service.
type ServiceAPI interface {
	// AddTarget adds a new target to the service's configuration given the endpoint of the new node.
	// The instanceID, network, and container name are used to identify the node as jobName in the service's configuration.
	// The labels are added to the service's metrics.
	AddTarget(target types.MonitoringTarget, labels map[string]string, jobName string) error

	// RemoveTarget removes a target from the service's configuration given the instanceID of the node to be removed.
	// It returns the network of the removed node.
	RemoveTarget(instanceID string) (string, error)

	// DotEnv returns a map of the service's environment variables and their default values.
	DotEnv() map[string]string

	// Setup configures the service given a map of options. The options should include the values for the environment variables.
	Setup(options map[string]string) error

	// Init initializes the service with the given ServiceOptions.
	Init(types.ServiceOptions) error

	// SetContainerIP sets the container IP of the service.
	SetContainerIP(ip net.IP)

	// ContainerName returns the name of the service's container.
	ContainerName() string

	// Endpoint returns the endpoint of the service.
	Endpoint() string
}
