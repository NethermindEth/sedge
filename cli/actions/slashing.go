package actions

import (
	"context"
	"fmt"
	"path"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
)

const SlashingImportFile string = "slashing-import.json"

type SlashingImportOptions struct {
	ValidatorClient string
	Network         string
	StopValidator   bool
	StartValidator  bool
	GenerationPath  string
	From            string
}

func (s *sedgeActions) ImportSlashingInterchangeData(options SlashingImportOptions) error {
	// Check validator container exists
	_, err := s.serviceManager.ContainerId(services.ServiceCtValidator)
	if err != nil {
		return err
	}
	previouslyRunning, err := s.serviceManager.IsRunning(services.ServiceCtValidator)
	if err != nil {
		return err
	}
	// Stop validator
	log.Info("Stopping validator client")
	if err := s.serviceManager.Stop(services.ServiceCtValidator); err != nil {
		return err
	}

	// Copy slashing data to generation path
	slashingDataPath := path.Join(options.GenerationPath, configs.ValidatorDir, SlashingImportFile)
	log.Debugf("Copying slashing data file from %s to %s", options.From, slashingDataPath)
	if err := utils.CopyFile(options.From, slashingDataPath); err != nil {
		return err
	}

	log.Infof("Importing slashing data to client %s from %s", options.ValidatorClient, options.From)
	var cmd []string
	switch options.ValidatorClient {
	case "prysm":
		cmd = []string{
			"--", "slashing-protection-history",
			"import",
			"--accept-terms-of-use",
			"--" + options.Network,
			"--datadir=/data",
			"--slashing-protection-json-file=/data/slashing-import.json",
		}
	case "lighthouse":
		cmd = []string{
			"lighthouse", "account", "validator", "slashing-protection", "import",
			"--network", options.Network,
			"--datadir", "/data",
			"/data/slashing-import.json",
		}
	case "lodestar":
		cmd = []string{
			"validator", "slashing-protection", "import",
			"--network", options.Network,
			"--dataDir", "/data/validator",
			"--file", "/data/validator/slashing-import.json",
		}
	case "teku":
		cmd = []string{
			"slashing-protection",
			"import",
			"--data-path=/data",
			"--from=/data/slashing-import.json",
		}
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedValidatorClient, options.ValidatorClient)
	}
	if err := runSlashingContainer(s.dockerClient, s.serviceManager, cmd); err != nil {
		return err
	}

	// Run validator again
	if (previouslyRunning && !options.StopValidator) || options.StartValidator {
		log.Info("the validator container is being restarted")
		if err := s.serviceManager.Start(services.ServiceCtValidator); err != nil {
			return err
		}
	}
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
	// Check validator container exists
	_, err := s.serviceManager.ContainerId(services.ServiceCtValidator)
	if err != nil {
		return err
	}
	previouslyRunning, err := s.serviceManager.IsRunning(services.ServiceCtValidator)
	if err != nil {
		return err
	}
	// Stop validator client
	log.Info("Stopping validator client")
	if err := s.serviceManager.Stop(services.ServiceCtValidator); err != nil {
		return err
	}

	log.Infof("Exporting slashing data from client %s", options.ValidatorClient)
	var cmd []string
	switch options.ValidatorClient {
	case "prysm":
		cmd = []string{
			"--", "slashing-protection-history",
			"export",
			"--accept-terms-of-use",
			"--" + options.Network,
			"--datadir=/data",
			"--slashing-protection-export-dir=/data",
		}
	case "lighthouse":
		cmd = []string{
			"lighthouse", "account", "validator", "slashing-protection", "export",
			"--network", options.Network,
			"--datadir", "/data",
			"/data/slashing_protection.json",
		}
	case "lodestar":
		cmd = []string{
			"validator", "slashing-protection", "export",
			"--network", options.Network,
			"--dataDir", "/data/validator",
			"--file", "/data/validator/slashing_protection.json",
		}
	case "teku":
		cmd = []string{
			"slashing-protection",
			"export",
			"--data-path=/data",
			"--to=/data/slashing_protection.json",
		}
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedValidatorClient, options.ValidatorClient)
	}
	if err := runSlashingContainer(s.dockerClient, s.serviceManager, cmd); err != nil {
		return err
	}
	copyFrom := filepath.Join(options.GenerationPath, configs.ValidatorDir, "slashing_protection.json")
	log.Debug("copying slashing data file from %s to %s", copyFrom, options.Out)
	if err := utils.CopyFile(copyFrom, options.Out); err != nil {
		return err
	}

	// Run validator again
	if (previouslyRunning && !options.StopValidator) || options.StartValidator {
		log.Info("the validator container is being restarted")
		if err := s.serviceManager.Start(services.ServiceCtValidator); err != nil {
			return err
		}
	}
	return nil
}

func runSlashingContainer(dockerClient client.APIClient, serviceManager services.ServiceManager, cmd []string) error {
	validatorImage, err := serviceManager.Image(services.ServiceCtValidator)
	if err != nil {
		return err
	}
	log.Debugf("Creating %s container", services.ServiceCtSlashingData)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: validatorImage,
			Cmd:   cmd,
		},
		&container.HostConfig{
			VolumesFrom: []string{services.ServiceCtValidator},
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		services.ServiceCtSlashingData,
	)
	if err != nil {
		return err
	}
	log.Debugf("slashing container id: %s", ct.ID)
	ctExit, errChan := serviceManager.Wait(services.ServiceCtSlashingData, container.WaitConditionNextExit)
	if err := dockerClient.ContainerStart(context.Background(), ct.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	for {
		select {
		case exitResult := <-ctExit:
			if exitResult.StatusCode != 0 {
				return fmt.Errorf("slashing service ends with status code %d, check container %s logs for more details", exitResult.StatusCode, ct.ID)
			}
			return deleteContainer(dockerClient, ct.ID)
		case exitErr := <-errChan:
			return exitErr
		}
	}
}

func deleteContainer(dockerClient client.APIClient, container string) error {
	if err := dockerClient.ContainerRemove(context.Background(), container, types.ContainerRemoveOptions{}); err != nil {
		return fmt.Errorf("error removing slashing container %s: %w", container, err)
	}
	return nil
}
