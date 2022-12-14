package slashing

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
)

type SlashingDataManager interface {
	Import(clientName, network string) error
	Export(clientName, network, generationPath, out string) error
}

type slashingDataManager struct {
	dockerClient   client.APIClient
	serviceManager services.ServiceManager
}

func NewSlashingDataManager(dockerClient client.APIClient, serviceManager services.ServiceManager) SlashingDataManager {
	return &slashingDataManager{
		dockerClient:   dockerClient,
		serviceManager: serviceManager,
	}
}

func (s *slashingDataManager) Import(clientName, network string) error {
	log.Infof("Importing slashing data to client %s", clientName)
	var cmd []string
	switch clientName {
	case "prysm":
		cmd = []string{
			"--", "slashing-protection-history",
			"import",
			"--accept-terms-of-use",
			"--" + network,
			"--datadir=/data",
			"--slashing-protection-json-file=/data/slashing-import.json",
		}
	case "lighthouse":
		cmd = []string{
			"lighthouse", "account", "validator", "slashing-protection", "import",
			"--network", network,
			"--datadir", "/data",
			"/data/slashing-import.json",
		}
	case "lodestar":
		cmd = []string{
			"validator", "slashing-protection", "import",
			"--network", network,
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
		return fmt.Errorf("slashing import nos supported for client %s", clientName)
	}
	return runSlashingContainer(s.dockerClient, s.serviceManager, cmd)
}

func (s *slashingDataManager) Export(clientName, network, generationPath, out string) error {
	log.Infof("Exporting slashing data from client %s", clientName)
	var cmd []string
	switch clientName {
	case "prysm":
		cmd = []string{
			"--", "slashing-protection-history",
			"export",
			"--accept-terms-of-use",
			"--" + network,
			"--datadir=/data",
			"--slashing-protection-export-dir=/data",
		}
	case "lighthouse":
		cmd = []string{
			"lighthouse", "account", "validator", "slashing-protection", "export",
			"--network", network,
			"--datadir", "/data",
			"/data/slashing_protection.json",
		}
	case "lodestar":
		cmd = []string{
			"validator", "slashing-protection", "export",
			"--network", network,
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
		return fmt.Errorf("slashing export not supported for client %s", clientName)
	}
	if err := runSlashingContainer(s.dockerClient, s.serviceManager, cmd); err != nil {
		return err
	}
	return utils.CopyFile(filepath.Join(generationPath, "validator-data", "slashing_protection.json"), out)
}

func runSlashingContainer(dockerClient client.APIClient, serviceManager services.ServiceManager, cmd []string) error {
	validatorImage, err := serviceManager.Image(services.ServiceValidator)
	if err != nil {
		return err
	}
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: validatorImage,
			Cmd:   cmd,
		},
		&container.HostConfig{
			VolumesFrom: []string{services.ServiceContainer[services.ServiceValidator]},
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		"",
	)
	if err != nil {
		return err
	}
	log.Infof("slashing container id: %s", ct.ID)
	ctExit, errChan := dockerClient.ContainerWait(context.Background(), ct.ID, container.WaitConditionNextExit)
	if err := dockerClient.ContainerStart(context.Background(), ct.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	log.Infof("Waiting for slashing service (container id: %s)", ct.ID)
	for {
		select {
		case exitResult := <-ctExit:
			if exitResult.StatusCode != 0 {
				return fmt.Errorf("slashing service ends with status code %d, check container %s logs for more details", exitResult.StatusCode, ct.ID)
			}
			return deleteContainer(dockerClient, ct.ID)
		case exitErr := <-errChan:
			fmt.Printf("Exit err: %v\n", exitErr)
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
