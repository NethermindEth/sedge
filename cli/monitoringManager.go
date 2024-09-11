package cli

//go:generate mockgen -package=sedge_mocks -destination=../mocks/monitoringManager.go github.com/NethermindEth/sedge/cli MonitoringManager

import (
	"github.com/NethermindEth/sedge/internal/common"
	"github.com/NethermindEth/sedge/internal/monitoring/services/types"
)

type MonitoringManager interface {
	// Init initializes the monitoring stack. Assumes that the stack is already installed.
	Init() error

	// InstallStack installs the monitoring stack.
	InstallStack() error

	// AddTarget adds a new target to all services in the monitoring stack.
	// It also connects the target to the docker network of the monitoring stack if it isn't already connected.
	// The labels are added to the service's metrics.
	AddTarget(target types.MonitoringTarget, labels map[string]string, dockerNetwork string) error

	// RemoveTarget removes a target from the monitoring stack.
	// The dockerNetwork is the name of the network the node is connected to.
	RemoveTarget(endpoint string) error

	// Status returns the status of the monitoring stack.
	Status() (common.Status, error)

	// InstallationStatus returns the installation status of the monitoring stack.
	InstallationStatus() (common.Status, error)

	// Run runs the monitoring stack.
	Run() error

	// Stop stops the monitoring stack.
	Stop() error

	// Cleanup removes the monitoring stack. If force is true, it will remove the stack directly bypassing any checks.
	Cleanup() error

	// ServiceEndpoints returns the endpoints of the monitoring services.
	ServiceEndpoints() map[string]string
}
