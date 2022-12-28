package actions

import (
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/docker/docker/client"
)

type SedgeActions interface {
	ImportSlashingInterchangeData(SlashingImportOptions) error
	ExportSlashingInterchangeData(SlashingExportOptions) error
	GenerateCompose(GenerateComposeOptions) error
}

type sedgeActions struct {
	dockerClient   client.APIClient
	serviceManager services.ServiceManager
}

func NewSedgeActions(dockerClient client.APIClient, serviceManager services.ServiceManager) SedgeActions {
	return &sedgeActions{
		dockerClient:   dockerClient,
		serviceManager: serviceManager,
	}
}
