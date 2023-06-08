/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package actions

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	clientsimages "github.com/NethermindEth/sedge/configs/images"
	"github.com/NethermindEth/sedge/internal/images/validator-import/lighthouse"
	"github.com/NethermindEth/sedge/internal/images/validator-import/teku"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/otiai10/copy"
	log "github.com/sirupsen/logrus"
)

var ErrInterrupted = errors.New("interrupt")

type ImportValidatorKeysOptions struct {
	ValidatorClient string
	Network         string
	StopValidator   bool
	StartValidator  bool
	From            string
	GenerationPath  string
	ContainerTag    string
	CustomConfig    ImportValidatorKeysCustomOptions
}

type ImportValidatorKeysCustomOptions struct {
	NetworkConfigPath string
	GenesisPath       string
	DeployBlockPath   string
}

func (s *sedgeActions) ImportValidatorKeys(options ImportValidatorKeysOptions) error {
	// Ensure paths are absolute
	if absFrom, err := filepath.Abs(options.From); err != nil {
		return err
	} else {
		options.From = absFrom
	}
	if absGenerationPath, err := filepath.Abs(options.GenerationPath); err != nil {
		return err
	} else {
		options.GenerationPath = absGenerationPath
	}
	validatorCtName := services.ContainerNameWithTag(services.DefaultSedgeValidatorClient, options.ContainerTag)
	// Check validator container exists
	_, err := s.serviceManager.ContainerId(validatorCtName)
	if err != nil {
		return err
	}
	previouslyRunning, err := s.serviceManager.IsRunning(validatorCtName)
	if err != nil {
		return err
	}
	// Stop validator
	if previouslyRunning {
		log.Info("Stopping validator client")
		if err := s.serviceManager.Stop(validatorCtName); err != nil {
			return err
		}
	}

	absFrom, err := filepath.Abs(options.From)
	if err != nil {
		return err
	}
	options.From = absFrom
	absGenerationPath, err := filepath.Abs(options.GenerationPath)
	if err != nil {
		return err
	}
	options.GenerationPath = absGenerationPath

	if !isDefaultKeysPath(options.GenerationPath, options.From) {
		defaultKeystorePath := filepath.Join(options.GenerationPath, "keystore")
		log.Warnf("The keys path is not the default one, copying the keys to the default path %s", defaultKeystorePath)
		copy.Copy(options.From, defaultKeystorePath)
	}

	var ctID string
	switch options.ValidatorClient {
	case "prysm":
		prysmCtID, err := setupPrysmValidatorImportContainer(
			s.dockerClient,
			s.serviceManager,
			s.clientsImages,
			options,
		)
		if err != nil {
			return err
		}
		ctID = prysmCtID
	case "lodestar":
		lodestarCtID, err := setupLodestarValidatorImport(
			s.dockerClient,
			s.serviceManager,
			s.clientsImages,
			options,
		)
		if err != nil {
			return err
		}
		ctID = lodestarCtID
	case "lighthouse":
		lighthouseCtID, err := setupLighthouseValidatorImport(
			s.dockerClient,
			s.commandRunner,
			s.clientsImages,
			options,
		)
		if err != nil {
			return err
		}
		ctID = lighthouseCtID
	case "teku":
		tekuCtID, err := setupTekuValidatorImport(
			s.dockerClient,
			s.commandRunner,
			s.clientsImages,
			options,
		)
		if err != nil {
			return err
		}
		ctID = tekuCtID
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedValidatorClient, options.ValidatorClient)
	}
	log.Info("Importing validator keys")
	runErr := runAndWaitImportKeys(s.dockerClient, s.serviceManager, ctID)
	// Run validator again
	if (previouslyRunning && !options.StopValidator) || options.StartValidator {
		log.Info("The validator container is being restarted")
		if err := s.serviceManager.Start(validatorCtName); err != nil {
			return err
		}
	}
	if runErr == nil {
		log.Info("Validator keys imported successfully")
	}
	return runErr
}

func isDefaultKeysPath(generationPath, from string) bool {
	return from == filepath.Join(generationPath, "keystore")
}

func setupPrysmValidatorImportContainer(
	dockerClient client.APIClient,
	serviceManager services.ServiceManager,
	clientsImages clientsimages.ClientsImages,
	options ImportValidatorKeysOptions,
) (string, error) {
	var (
		validatorCtName       = services.ContainerNameWithTag(services.DefaultSedgeValidatorClient, options.ContainerTag)
		validatorImportCtName = services.ContainerNameWithTag(services.ServiceCtValidatorImport, options.ContainerTag)
	)
	validatorImage, err := serviceManager.Image(validatorCtName)
	if err != nil {
		return "", err
	}
	// Mounts
	mounts := []mount.Mount{
		{
			Type:   mount.TypeBind,
			Source: options.From,
			Target: "/keystore",
		},
	}
	// CMD
	cmd := []string{
		"--",
		"accounts", "import",
		"--accept-terms-of-use",
		"--keys-dir=/keystore/validator_keys",
		"--wallet-dir=/data/wallet",
		"--wallet-password-file=/keystore/keystore_password.txt",
		"--account-password-file=/keystore/keystore_password.txt",
	}
	// Custom options
	if options.CustomConfig.NetworkConfigPath != "" {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: options.CustomConfig.NetworkConfigPath,
			Target: "/network_config/config.yml",
		})
		cmd = append(cmd, "--chain-config-file=/network_config/config.yml")
	} else {
		cmd = append(cmd, "--"+options.Network)
	}
	log.Debugf("Creating %s container", validatorImportCtName)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: validatorImage,
			Cmd:   cmd,
		},
		&container.HostConfig{
			Mounts:      mounts,
			VolumesFrom: []string{validatorCtName},
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		validatorImportCtName,
	)
	if err != nil {
		return "", err
	}
	return ct.ID, nil
}

func setupLodestarValidatorImport(
	dockerClient client.APIClient,
	serviceManager services.ServiceManager,
	clientsImages clientsimages.ClientsImages,
	options ImportValidatorKeysOptions,
) (string, error) {
	var (
		validatorCtName       = services.ContainerNameWithTag(services.DefaultSedgeValidatorClient, options.ContainerTag)
		validatorImportCtName = services.ContainerNameWithTag(services.ServiceCtValidatorImport, options.ContainerTag)
	)
	validatorImage, err := serviceManager.Image(validatorCtName)
	if err != nil {
		return "", err
	}
	// Mounts
	mounts := []mount.Mount{
		{
			Type:   mount.TypeBind,
			Source: options.From,
			Target: "/keystore",
		},
	}
	// CMD
	var preset string
	switch options.Network {
	case "mainnet", "goerli", "sepolia":
		preset = "mainnet"
	case "gnosis", "chiado":
		preset = "gnosis"
	default:
		return "", newUnknownLodestarPresetError(options.Network)
	}
	cmd := []string{
		"validator", "import",
		"--dataDir", "/data",
		"--importKeystores=/keystore/validator_keys",
		"--importKeystoresPassword=/keystore/keystore_password.txt",
	}
	if options.CustomConfig.NetworkConfigPath != "" {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: options.CustomConfig.NetworkConfigPath,
			Target: "/network_config/config.yaml",
		})
		cmd = append(cmd, "--paramsFile=/network_config/config.yaml")
	} else {
		cmd = append(cmd, "--network", options.Network)
		cmd = append(cmd, "--preset", preset)
	}
	log.Debugf("Creating %s container", validatorImportCtName)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: validatorImage,
			Cmd:   cmd,
		},
		&container.HostConfig{
			Mounts:      mounts,
			VolumesFrom: []string{validatorCtName},
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		validatorImportCtName,
	)
	if err != nil {
		return "", err
	}
	return ct.ID, nil
}

func setupLighthouseValidatorImport(
	dockerClient client.APIClient,
	commandRunner commands.CommandRunner,
	clientsImages clientsimages.ClientsImages,
	options ImportValidatorKeysOptions,
) (string, error) {
	var (
		validatorCtName       = services.ContainerNameWithTag(services.DefaultSedgeValidatorClient, options.ContainerTag)
		validatorImportCtName = services.ContainerNameWithTag(services.ServiceCtValidatorImport, options.ContainerTag)
	)
	if options.Network == "chiado" {
		options.Network = "custom"
	}
	// Init build context
	contextDir, err := lighthouse.InitContext()
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrCreatingContextDir, err.Error())
	}
	// Build image
	buildCmd := commandRunner.BuildDockerBuildCMD(commands.DockerBuildOptions{
		Path: contextDir,
		Tag:  "sedge/validator-import-lighthouse",
		Args: map[string]string{
			"NETWORK":    options.Network,
			"LH_VERSION": clientsImages.Validator().Lighthouse().String(),
		},
	})
	log.Infof(configs.RunningCommand, buildCmd.Cmd)
	if _, _, err = commandRunner.RunCMD(buildCmd); err != nil {
		return "", err
	}
	mounts := []mount.Mount{
		{
			Type:   mount.TypeBind,
			Source: options.From,
			Target: "/keystore",
		},
	}
	if options.CustomConfig.NetworkConfigPath != "" {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: options.CustomConfig.NetworkConfigPath,
			Target: "/network_config/config.yaml",
		})
	}
	if options.CustomConfig.GenesisPath != "" {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: options.CustomConfig.GenesisPath,
			Target: "/network_config/genesis.ssz",
		})
	}
	if options.CustomConfig.DeployBlockPath != "" {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: options.CustomConfig.DeployBlockPath,
			Target: "/network_config/deploy_block.txt",
		})
	}
	log.Debugf("Creating %s container", validatorImportCtName)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: "sedge/validator-import-lighthouse",
		},
		&container.HostConfig{
			Mounts:      mounts,
			VolumesFrom: []string{validatorCtName},
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		validatorImportCtName,
	)
	if err != nil {
		return "", err
	}
	return ct.ID, nil
}

func setupTekuValidatorImport(
	dockerClient client.APIClient,
	commandRunner commands.CommandRunner,
	clientsImages clientsimages.ClientsImages,
	options ImportValidatorKeysOptions,
) (string, error) {
	var (
		validatorCtName       = services.ContainerNameWithTag(services.DefaultSedgeValidatorClient, options.ContainerTag)
		validatorImportCtName = services.ContainerNameWithTag(services.ServiceCtValidatorImport, options.ContainerTag)
	)
	// Init build context
	contextDir, err := teku.InitContext()
	if err != nil {
		return "", err
	}
	// Build image
	buildCmd := commandRunner.BuildDockerBuildCMD(commands.DockerBuildOptions{
		Path: contextDir,
		Tag:  "sedge/validator-import-teku",
	})
	log.Infof(configs.RunningCommand, buildCmd.Cmd)
	if _, _, err := commandRunner.RunCMD(buildCmd); err != nil {
		return "", err
	}
	// Mounts
	mounts := []mount.Mount{
		{
			Type:   mount.TypeBind,
			Source: options.From,
			Target: "/keystore",
		},
	}
	if options.CustomConfig.NetworkConfigPath != "" {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: options.CustomConfig.NetworkConfigPath,
			Target: "/network_config/config.yml",
		})
	}
	log.Debugf("Creating %s container", validatorImportCtName)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: "sedge/validator-import-teku",
		},
		&container.HostConfig{
			Mounts:      mounts,
			VolumesFrom: []string{validatorCtName},
			LogConfig: container.LogConfig{
				Type: "json-file",
			},
		},
		&network.NetworkingConfig{},
		&v1.Platform{},
		validatorImportCtName,
	)
	if err != nil {
		return "", err
	}
	return ct.ID, nil
}

func runAndWaitImportKeys(dockerClient client.APIClient, serviceManager services.ServiceManager, ctID string) error {
	log.Debugf("import keys container id: %s", ctID)
	ctExit, errChan := serviceManager.Wait(ctID, container.WaitConditionNextExit)
	log.Info("The keys import container is starting")
	if err := dockerClient.ContainerStart(context.Background(), ctID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)
	for {
		select {
		case exitResult := <-ctExit:
			logs, err := serviceManager.ContainerLogs(ctID, "Import keys")
			if err != nil {
				return err
			}
			if err = deleteContainer(dockerClient, ctID); err != nil {
				return err
			}
			if exitResult.StatusCode != 0 {
				return newValidatorImportCtBadExitCodeError(ctID, exitResult.StatusCode, logs)
			}
			return nil
		case <-osSignals:
			if err := stopContainer(dockerClient, ctID); err != nil {
				return err
			}
			if err := deleteContainer(dockerClient, ctID); err != nil {
				return err
			}
			return ErrInterrupted
		case exitErr := <-errChan:
			return exitErr
		}
	}
}
