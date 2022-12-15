package actions

import (
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/pkg/slashing"
)

type SedgeActions interface {
	ImportSlashingInterchangeData() error
	ExportSlashingInterchangeData(SlashingExportOptions) error
}

type sedgeActions struct {
	serviceManager  services.ServiceManager
	slashingManager slashing.SlashingDataManager
}

func NewSedgeActions(serviceManager services.ServiceManager, slashingManager slashing.SlashingDataManager) SedgeActions {
	return &sedgeActions{
		serviceManager:  serviceManager,
		slashingManager: slashingManager,
	}
}
