package monitoring

import (
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/monitoring/compose"
)

// ComposeManager is an interface that defines methods for managing Docker Compose operations.
type ComposeManager interface {
	// Up starts the Docker Compose services defined in the Docker Compose file specified in the options.
	Up(opts commands.DockerComposeUpOptions) error

	// Down stops and removes the Docker Compose services defined in the Docker Compose file specified in the options.
	Down(opts commands.DockerComposeDownOptions) error

	// Create creates the Docker Compose services defined in the Docker Compose file specified in the options, but does not start them.
	Create(opts commands.DockerComposeCreateOptions) error

	/// PS runs the Docker Compose 'ps' command for the specified options and returns the list of services.
	PS(opts commands.DockerComposePsOptions) ([]compose.ComposeService, error)
}
