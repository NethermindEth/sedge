package actions

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
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

type ImportValidatorKeysOptions struct {
	ValidatorClient string
	Network         string
	StopValidator   bool
	StartValidator  bool
	From            string
	GenerationPath  string
	CustomConfig    ImportValidatorKeysCustomOptions
}
type ImportValidatorKeysCustomOptions struct {
	NetworkConfigPath *string
	GenesisPath       *string
	DeployBlockPath   *string
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
	if previouslyRunning {
		log.Info("Stopping validator client")
		if err := s.serviceManager.Stop(services.ServiceCtValidator); err != nil {
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
		prysmCtID, err := setupPrysmValidatorImportContainer(s.dockerClient, s.serviceManager, options)
		if err != nil {
			return err
		}
		ctID = prysmCtID
	case "lodestar":
		lodestarCtID, err := setupLodestarValidatorImport(s.dockerClient, s.serviceManager, options)
		if err != nil {
			return err
		}
		ctID = lodestarCtID
	case "lighthouse":
		lighthouseCtID, err := setupLighthouseValidatorImport(s.dockerClient, s.serviceManager, s.commandRunner, options)
		if err != nil {
			return err
		}
		ctID = lighthouseCtID
	case "teku":
		tekuCtID, err := setupTekuValidatorImport(s.dockerClient, s.serviceManager, s.commandRunner, options)
		if err != nil {
			return err
		}
		ctID = tekuCtID
	default:
		return fmt.Errorf("%w: %s", ErrUnsupportedValidatorClient, options.ValidatorClient)
	}
	log.Info("Importing validator keys")
	if err := runAndWait(s.dockerClient, s.serviceManager, ctID); err != nil {
		return err
	}
	// Run validator again
	if (previouslyRunning && !options.StopValidator) || options.StartValidator {
		log.Info("The validator container is being restarted")
		if err := s.serviceManager.Start(services.ServiceCtValidator); err != nil {
			return err
		}
	}
	log.Info("Validator keys imported successfully")
	return nil
}

func isDefaultKeysPath(generationPath, from string) bool {
	return from == filepath.Join(generationPath, "keystore")
}

func setupPrysmValidatorImportContainer(dockerClient client.APIClient, serviceManager services.ServiceManager, options ImportValidatorKeysOptions) (string, error) {
	validatorImage, err := serviceManager.Image(services.ServiceCtValidator)
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
	if options.CustomConfig.NetworkConfigPath != nil {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: *options.CustomConfig.NetworkConfigPath,
			Target: "/network_config/config.yml",
		})
		cmd = append(cmd, "--chain-config-file=/network_config/config.yml")
	} else {
		cmd = append(cmd, "--"+options.Network)
	}
	log.Debugf("Creating %s container", services.ServiceCtValidatorImport)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: validatorImage,
			Cmd:   cmd,
		},
		&container.HostConfig{
			Mounts:      mounts,
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

func setupLodestarValidatorImport(dockerClient client.APIClient, serviceManager services.ServiceManager, options ImportValidatorKeysOptions) (string, error) {
	validatorImage, err := serviceManager.Image(services.ServiceCtValidator)
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
		return "", newErrUnknownLodestarPreset(options.Network)
	}
	cmd := []string{
		"validator", "import",
		"--preset", preset,
		"--dataDir", "/data",
		"--importKeystores=/keystore/validator_keys",
		"--importKeystoresPassword=/keystore/keystore_password.txt",
	}
	if options.CustomConfig.NetworkConfigPath != nil {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: *options.CustomConfig.NetworkConfigPath,
			Target: "/network_config/config.yaml",
		})
		cmd = append(cmd, "--paramsFile=/network_config/config.yaml")
	} else {
		cmd = append(cmd, "--network", options.Network)
	}
	log.Debugf("Creating %s container", services.ServiceCtValidatorImport)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: validatorImage,
			Cmd:   cmd,
		},
		&container.HostConfig{
			Mounts:      mounts,
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

func setupLighthouseValidatorImport(dockerClient client.APIClient, serviceManager services.ServiceManager, commandRunner commands.CommandRunner, options ImportValidatorKeysOptions) (string, error) {
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
			"LH_VERSION": configs.ClientImages.Validator.Lighthouse.String(),
		},
	})
	log.Infof(configs.RunningCommand, buildCmd.Cmd)
	if _, err = commandRunner.RunCMD(buildCmd); err != nil {
		return "", err
	}
	mounts := []mount.Mount{
		{
			Type:   mount.TypeBind,
			Source: options.From,
			Target: "/keystore",
		},
	}
	if options.CustomConfig.NetworkConfigPath != nil {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: *options.CustomConfig.NetworkConfigPath,
			Target: "/network_config/config.yaml",
		})
	}
	if options.CustomConfig.GenesisPath != nil {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: *options.CustomConfig.GenesisPath,
			Target: "/network_config/genesis.ssz",
		})
	}
	if options.CustomConfig.DeployBlockPath != nil {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: *options.CustomConfig.DeployBlockPath,
			Target: "/network_config/deploy_block.txt",
		})
	}
	log.Debugf("Creating %s container", services.ServiceCtValidatorImport)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: "sedge/validator-import-lighthouse",
		},
		&container.HostConfig{
			Mounts:      mounts,
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

func setupTekuValidatorImport(dockerClient client.APIClient, serviceManager services.ServiceManager, commandRunner commands.CommandRunner, options ImportValidatorKeysOptions) (string, error) {
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
	if _, err := commandRunner.RunCMD(buildCmd); err != nil {
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
	if options.CustomConfig.NetworkConfigPath != nil {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: *options.CustomConfig.NetworkConfigPath,
			Target: "/network_config/config.yml",
		})
	}
	log.Debugf("Creating %s container", services.ServiceCtValidatorImport)
	ct, err := dockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Image: "sedge/validator-import-teku",
		},
		&container.HostConfig{
			Mounts:      mounts,
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
				logs, err := containerLogs(dockerClient, ctID)
				if err != nil {
					return err
				}
				return newErrValidatorImportCtBadExitCode(ctID, exitResult.StatusCode, logs)
			}
			return deleteContainer(dockerClient, ctID)
		case exitErr := <-errChan:
			return exitErr
		}
	}
}

func containerLogs(dockerClient client.APIClient, ctID string) (string, error) {
	logReader, err := dockerClient.ContainerLogs(context.Background(), ctID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     false,
	})
	if err != nil {
		return "", err
	}
	defer logReader.Close()
	logs, err := ioutil.ReadAll(logReader)
	return string(logs), err
}
