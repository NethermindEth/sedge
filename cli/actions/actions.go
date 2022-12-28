package actions

import (
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/docker/docker/client"
)

type SedgeActions interface {
	ImportSlashingInterchangeData(SlashingImportOptions) error
	ExportSlashingInterchangeData(SlashingExportOptions) error
	SetupContainers(SetupContainersOptions) error
}

type sedgeActions struct {
	dockerClient   client.APIClient
	serviceManager services.ServiceManager
	commandRunner  commands.CommandRunner
}

func NewSedgeActions(dockerClient client.APIClient, serviceManager services.ServiceManager, commandRunner commands.CommandRunner) SedgeActions {
	return &sedgeActions{
		dockerClient:   dockerClient,
		serviceManager: serviceManager,
		commandRunner:  commandRunner,
	}
}
