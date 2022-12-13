package slashing

import (
	"context"
	"fmt"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
)

type SlashingDataManager interface {
	Import(clientName, network string) error
	Export(clientName string) error
}

type slashingDataManager struct {
	dockerClient client.APIClient
}

func NewSlashingDataManager(dockerClient client.APIClient) SlashingDataManager {
	return &slashingDataManager{dockerClient: dockerClient}
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
	return importSlashing(s.dockerClient, cmd)
}

func importSlashing(dockerClient client.APIClient, cmd []string) error {
	validatorImage, err := services.Image(dockerClient, services.ServiceValidator)
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
	log.Infof("slashing import container id: %s", ct.ID)
	ctExit, errChan := dockerClient.ContainerWait(context.Background(), ct.ID, container.WaitConditionNextExit)
	if err := dockerClient.ContainerStart(context.Background(), ct.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	log.Infof("Waiting for slashing import service (container id: %s)", ct.ID)
	for {
		select {
		case exitResult := <-ctExit:
			fmt.Printf("Exit result: %v\n", exitResult.StatusCode)
			if exitResult.StatusCode != 0 {
				return fmt.Errorf("slashing import ends with status code %d", exitResult.StatusCode)
			}
			return nil
		case exitErr := <-errChan:
			fmt.Printf("Exit err: %v\n", exitErr)
			return exitErr
		}
	}
}

func (s *slashingDataManager) Export(clientName string) error {
	return nil
}
