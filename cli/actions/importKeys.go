package actions

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/images/validator-import/lighthouse"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
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

	var ctID string
	switch options.ValidatorClient {
	case "prysm":
		cmd := []string{
			"--",
			"accounts", "import",
			"--accept-terms-of-use",
			"--" + options.Network,
			"--keys-dir=/keystore/validator_keys",
			"--wallet-dir=/data/wallet",
			"--wallet-password-file=/keystore/keystore_password.txt",
			"--account-password-file=/keystore/keystore_password.txt",
		}
		prysmCtID, err := setupValidatorImportContainer(s.dockerClient, s.serviceManager, cmd, options.From)
		if err != nil {
			return err
		}
		ctID = prysmCtID
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
		cmd := []string{
			"validator", "import",
			"--preset", preset,
			"--network", options.Network,
			"--dataDir", "/data",
			"--importKeystores=/keystore/validator_keys",
			"--importKeystoresPassword=/keystore/keystore_password.txt",
		}
		lodestarCtID, err := setupValidatorImportContainer(s.dockerClient, s.serviceManager, cmd, options.From)
		if err != nil {
			return err
		}
		ctID = lodestarCtID
	case "lighthouse":
		// Init build context
		contextDir, err := lighthouse.InitContext()
		if err != nil {
			return fmt.Errorf("error creating context dir: %s", err.Error())
		}
		// Build image
		buildCmd := s.commandRunner.BuildDockerBuildCMD(commands.DockerBuildOptions{
			Path: contextDir,
			Tag:  "sedge/validator-import-lighthouse",
			Args: map[string]string{
				"NETWORK":    options.Network,
				"LH_VERSION": "sigp/lighthouse:v3.3.0",
			},
		})
		log.Infof(configs.RunningCommand, buildCmd.Cmd)
		if _, err = s.commandRunner.RunCMD(buildCmd); err != nil {
			return err
		}
		// Setup container
		lighthouseCtID, err := setupLighthouseValidatorImport(s.dockerClient, s.serviceManager, options.From)
		if err != nil {
			return err
		}
		ctID = lighthouseCtID
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedValidatorClient, options.ValidatorClient)
	}
	log.Info("Importing validator keys")
	if err := runAndWait(s.dockerClient, s.serviceManager, ctID); err != nil {
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

func setupValidatorImportContainer(dockerClient client.APIClient, serviceManager services.ServiceManager, cmd []string, from string) (string, error) {
	from, err := filepath.Abs(from)
	if err != nil {
		return "", err
	}
	validatorImage, err := serviceManager.Image(services.ServiceCtValidator)
	if err != nil {
		return "", err
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
		return "", err
	}
	return ct.ID, nil
}

func setupLighthouseValidatorImport(dockerClient client.APIClient, serviceManager services.ServiceManager, from string) (string, error) {
	log.Debugf("Creating %s container", services.ServiceCtValidatorImport)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: "sedge/validator-import-lighthouse",
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
		return "", err
	}
	return ct.ID, nil
}

func runAndWait(dockerClient client.APIClient, serviceManager services.ServiceManager, ctID string) error {
	log.Debugf("import keys container id: %s", ctID)
	ctExit, errChan := serviceManager.Wait(services.ServiceCtValidatorImport, container.WaitConditionNextExit)
	log.Info("The keys import container is starting")
	if err := dockerClient.ContainerStart(context.Background(), ctID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	for {
		select {
		case exitResult := <-ctExit:
			if exitResult.StatusCode != 0 {
				return fmt.Errorf("validator-import service ends with status code %d, check container %s logs for more details", exitResult.StatusCode, ctID)
			}
			return deleteContainer(dockerClient, ctID)
		case exitErr := <-errChan:
			return exitErr
		}
	}
}
