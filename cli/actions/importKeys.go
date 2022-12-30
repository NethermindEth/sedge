package actions

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
)

type ImportValidatorKeysOptions struct {
	ValidatorClient string
	Network         string
	StopValidator   bool
	StartValidator  bool
	From            string
}

func (s *sedgeActions) ImportValidatorKeys(options ImportValidatorKeysOptions) error {
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

	var cmd []string
	switch options.ValidatorClient {
	case "prysm":
		cmd = []string{
			"--",
			"accounts", "import",
			"--accept-terms-of-use",
			"--" + options.Network,
			"--keys-dir=/keystore/validator_keys",
			"--wallet-dir=/data/wallet",
			"--wallet-password-file=/keystore/keystore_password.txt",
			"--account-password-file=/keystore/keystore_password.txt",
		}
	case "lodestar":
		var preset string
		switch options.Network {
		case "mainnet", "goerli", "sepolia":
			preset = "mainnet"
		case "gnosis", "chiado":
			preset = "gnosis"
		default:
			return fmt.Errorf("unknown lodestar preset for network %s", options.Network)
		}
		cmd = []string{
			"validator", "import",
			"--preset", preset,
			"--network", options.Network,
			"--dataDir", "/data",
			"--importKeystores=/keystore/validator_keys",
			"--importKeystoresPassword=/keystore/keystore_password.txt",
		}
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedValidatorClient, options.ValidatorClient)
	}
	log.Info("Importing validator keys")
	if err := runValidatorImportContainer(s.dockerClient, s.serviceManager, cmd, options.From); err != nil {
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

func runValidatorImportContainer(dockerClient client.APIClient, serviceManager services.ServiceManager, cmd []string, from string) error {
	from, err := filepath.Abs(from)
	if err != nil {
		return err
	}
	validatorImage, err := serviceManager.Image(services.ServiceCtValidator)
	if err != nil {
		return err
	}
	log.Debugf("Creating %s container", services.ServiceCtValidatorImport)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: validatorImage,
			Cmd:   cmd,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: from,
					Target: "/keystore",
				},
			},
			VolumesFrom: []string{services.ServiceCtValidator},
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		services.ServiceCtValidatorImport,
	)
	if err != nil {
		return err
	}
	log.Debugf("import keys container id: %s", ct.ID)
	ctExit, errChan := serviceManager.Wait(services.ServiceCtValidatorImport, container.WaitConditionNextExit)
	log.Info("The keys import container is starting")
	if err := dockerClient.ContainerStart(context.Background(), ct.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	for {
		select {
		case exitResult := <-ctExit:
			if exitResult.StatusCode != 0 {
				return fmt.Errorf("validator-import service ends with status code %d, check container %s logs for more details", exitResult.StatusCode, ct.ID)
			}
			return deleteContainer(dockerClient, ct.ID)
		case exitErr := <-errChan:
			return exitErr
		}
	}
}
