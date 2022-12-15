package actions

import (
	"github.com/NethermindEth/sedge/internal/pkg/services"
)

func (s *sedgeActions) ImportSlashingInterchangeData() error {
	return nil
}

type SlashingExportOptions struct {
	ValidatorClient string
	Network         string
	StopValidator   bool
	StartValidator  bool
	GenerationPath  string
	Out             string
}

func (s *sedgeActions) ExportSlashingInterchangeData(options SlashingExportOptions) error {
	previouslyRunning, err := s.serviceManager.IsRunning(services.ServiceValidator)
	if err != nil {
		return err
	}
	// Stop validator client
	if err := s.serviceManager.Stop(services.ServiceValidator); err != nil {
		return err
	}
	// Export slashing data
	if err := s.slashingManager.Export(options.ValidatorClient, options.Network, options.GenerationPath, options.Out); err != nil {
		return err
	}
	// Run validator again
	if (previouslyRunning && !options.StopValidator) || options.StartValidator {
		if err := s.serviceManager.Start(services.ServiceValidator); err != nil {
			return err
		}
	}
	return nil
}
