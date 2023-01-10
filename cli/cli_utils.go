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
package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/crypto"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/manifoldco/promptui"
)

func randomizeClients(allClients clients.OrderedClients) (clients.Clients, error) {
	var executionClient, consensusClient *clients.Client
	var combinedClients clients.Clients

	executionClient, err := clients.RandomChoice(allClients[execution])
	if err != nil {
		return combinedClients, err
	}
	consensusClient, err = clients.RandomChoice(allClients[consensus])
	if err != nil {
		return combinedClients, err
	}

	combinedClients = clients.Clients{
		Execution: executionClient,
		Consensus: consensusClient,
		Validator: consensusClient,
	}
	return combinedClients, nil
}

func valClients(allClients clients.OrderedClients, flags *GenCmdFlags, services []string) (*clients.Clients, error) {
	var executionClient, consensusClient, validatorClient *clients.Client
	var err error

	// execution client
	if utils.Contains(services, execution) {
		executionParts := strings.Split(flags.executionName, ":")
		executionClient, err = clients.RandomChoice(allClients[execution])
		if err != nil {
			return nil, err
		}
		if flags.executionName != "" {
			log.Warn(configs.CustomImagesWarning)
			executionClient.Name = executionParts[0]
			executionClient.Image = strings.Join(executionParts[1:], ":")
		}
		if err = clients.ValidateClient(executionClient, execution); err != nil {
			return nil, err
		}
	} else {
		executionClient = nil
	}
	// consensus client
	if utils.Contains(services, consensus) {
		consensusParts := strings.Split(flags.consensusName, ":")
		consensusClient, err = clients.RandomChoice(allClients[consensus])
		if err != nil {
			return nil, err
		}
		if flags.consensusName != "" {
			log.Warn(configs.CustomImagesWarning)
			consensusClient.Name = consensusParts[0]
			consensusClient.Image = strings.Join(consensusParts[1:], ":")
		}
		if err = clients.ValidateClient(consensusClient, consensus); err != nil {
			return nil, err
		}
	} else {
		consensusClient = nil
	}
	// validator client
	if utils.Contains(services, validator) && !flags.noValidator {
		validatorParts := strings.Split(flags.validatorName, ":")
		validatorClient, err = clients.RandomChoice(allClients[validator])
		if err != nil {
			return nil, err
		}
		if flags.validatorName != "" {
			log.Warn(configs.CustomImagesWarning)
			validatorClient.Name = validatorParts[0]
			validatorClient.Image = strings.Join(validatorParts[1:], ":")
		}
		if err = clients.ValidateClient(validatorClient, validator); err != nil {
			return nil, err
		}
	} else {
		validatorClient = nil
	}

	return &clients.Clients{
		Execution: executionClient,
		Consensus: consensusClient,
		Validator: validatorClient,
	}, err
}

func validateClients(allClients clients.OrderedClients, w io.Writer, flags *CliCmdFlags) (clients.Clients, error) {
	var combinedClients clients.Clients
	var err error

	// Select a random execution client, consensus client and validator client
	randomizedClients, err := randomizeClients(allClients)
	if err != nil {
		return combinedClients, err
	}

	// Randomize missing clients, and choose same pair of client for consensus and validator if at least one of them is missing
	if flags.executionName == "" {
		log.Warnf(configs.ExecutionClientNotSpecifiedWarn, randomizedClients.Execution.Name)
		// TODO: avoid flag edition
		flags.executionName = randomizedClients.Execution.Name
	}
	if flags.consensusName == "" && flags.validatorName == "" {
		log.Warnf(configs.CLNotSpecifiedWarn, randomizedClients.Consensus.Name)
		// TODO: avoid edit flags
		flags.consensusName = randomizedClients.Consensus.Name
		flags.validatorName = randomizedClients.Validator.Name
	} else if flags.consensusName == "" {
		log.Warn(configs.ConsensusClientNotSpecifiedWarn)
		// TODO: avoid flag edition
		flags.consensusName = flags.validatorName
	} else if flags.validatorName == "" {
		log.Warn(configs.ValidatorClientNotSpecifiedWarn)
		// TODO: avoid flag edition
		flags.validatorName = flags.consensusName
	}

	exec, ok := allClients[execution][flags.executionName]
	if !ok {
		exec.Name = flags.executionName
	}
	cons, ok := allClients[consensus][flags.consensusName]
	if !ok {
		cons.Name = flags.consensusName
	}
	val, ok := allClients[validator][flags.validatorName]
	if !ok {
		val.Name = flags.validatorName
	}
	if !utils.Contains(*flags.services, execution) && len(*flags.services) > 0 && (*flags.services)[0] != "all" {
		exec = nil
	}
	if !utils.Contains(*flags.services, consensus) && len(*flags.services) > 0 && (*flags.services)[0] != "all" {
		cons = nil
	}
	if !utils.Contains(*flags.services, validator) && len(*flags.services) > 0 && (*flags.services)[0] != "all" ||
		flags.noValidator {
		val = nil
	}

	combinedClients = clients.Clients{
		Execution: exec,
		Consensus: cons,
		Validator: val,
	}

	if err = clients.ValidateClient(combinedClients.Execution, execution); err != nil {
		return combinedClients, err
	}
	if err = clients.ValidateClient(combinedClients.Consensus, consensus); err != nil {
		return combinedClients, err
	}
	if err = clients.ValidateClient(combinedClients.Validator, validator); err != nil {
		return combinedClients, err
	}

	return combinedClients, nil
}

func runScriptOrExit(cmdRunner commands.CommandRunner, flags *CliCmdFlags) (err error) {
	// notest
	log.Infof(configs.InstructionsFor, "running docker-compose script")
	upCMD := cmdRunner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     filepath.Join(flags.generationPath, configs.DefaultDockerComposeScriptName),
		Services: *flags.services,
	})
	fmt.Printf("\n%s\n\n", upCMD.Cmd)

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Run the script with the selected services %s", strings.Join(*flags.services, ", ")),
		IsConfirm: true,
		Default:   "Y",
	}
	_, err = prompt.Run()
	if err != nil {
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	if err := buildContainers(cmdRunner, *flags.services, flags.generationPath); err != nil {
		return err
	}
	if err = runAndShowContainers(cmdRunner, *flags.services, flags); err != nil {
		return err
	}

	return nil
}

func checkRunDependencies(cmdRunner commands.CommandRunner, generationPath string) error {
	// TODO: (refac) Put this check to checks.go and call it from there
	// Check if docker engine is on
	log.Info(configs.CheckingDockerEngine)
	psCMD := cmdRunner.BuildDockerPSCMD(commands.DockerPSOptions{
		All: true,
	})
	psCMD.GetOutput = true
	log.Infof(configs.RunningCommand, psCMD.Cmd)
	if _, err := cmdRunner.RunCMD(psCMD); err != nil {
		return fmt.Errorf(configs.DockerEngineOffError, err)
	}
	// Check that compose plugin is installed with docker running 'docker compose ps'
	dockerComposePsCMD := cmdRunner.BuildDockerComposePSCMD(commands.DockerComposePsOptions{
		Path: filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
	})
	log.Debugf(configs.RunningCommand, dockerComposePsCMD.Cmd)
	dockerComposePsCMD.GetOutput = true
	_, err := cmdRunner.RunCMD(dockerComposePsCMD)
	if err != nil {
		return fmt.Errorf(configs.DockerComposeOffError, err)
	}
	return nil
}

func buildImages(cmdRunner commands.CommandRunner, services []string, generationPath string) error {
	// Build images
	buildCmd := cmdRunner.BuildDockerComposeBuildCMD(commands.DockerComposeBuildOptions{
		Path:     filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services: services,
	})
	log.Infof(configs.RunningCommand, buildCmd.Cmd)
	if _, err := cmdRunner.RunCMD(buildCmd); err != nil {
		return fmt.Errorf(configs.CommandError, buildCmd.Cmd, err)
	}
	return nil
}

func downloadImages(cmdRunner commands.CommandRunner, services []string, generationPath string) error {
	// Download images
	pullCmd := cmdRunner.BuildDockerComposePullCMD(commands.DockerComposePullOptions{
		Path:     filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services: services,
	})
	log.Infof(configs.RunningCommand, pullCmd.Cmd)
	if _, err := cmdRunner.RunCMD(pullCmd); err != nil {
		return fmt.Errorf(configs.CommandError, pullCmd.Cmd, err)
	}
	return nil
}

func createContainers(cmdRunner commands.CommandRunner, services []string, generationPath string) error {
	if _, err := cmdRunner.RunCMD(cmdRunner.BuildDockerComposeCreateCMD(commands.DockerComposeCreateOptions{
		Path:     filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services: services,
	})); err != nil {
		return fmt.Errorf("error creating containers: %w", err)
	}
	return nil
}

func buildContainers(cmdRunner commands.CommandRunner, services []string, generationPath string) error {
	if err := checkRunDependencies(cmdRunner, generationPath); err != nil {
		return err
	}
	if err := buildImages(cmdRunner, services, generationPath); err != nil {
		return err
	}
	if err := downloadImages(cmdRunner, services, generationPath); err != nil {
		return err
	}
	if err := createContainers(cmdRunner, services, generationPath); err != nil {
		return err
	}
	return nil
}

// TODO: use flags.services instead a separated arg
func runAndShowContainers(cmdRunner commands.CommandRunner, services []string, flags *CliCmdFlags) error {
	// Run docker-compose script
	upCMD := cmdRunner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     filepath.Join(flags.generationPath, configs.DefaultDockerComposeScriptName),
		Services: services,
	})
	log.Infof(configs.RunningCommand, upCMD.Cmd)
	if _, err := cmdRunner.RunCMD(upCMD); err != nil {
		return fmt.Errorf(configs.CommandError, upCMD.Cmd, err)
	}

	// Run docker compose ps --filter status=running to show script running containers
	dcpsCMD := cmdRunner.BuildDockerComposePSCMD(commands.DockerComposePsOptions{
		Path:          filepath.Join(flags.generationPath, configs.DefaultDockerComposeScriptName),
		FilterRunning: true,
	})
	log.Infof(configs.RunningCommand, dcpsCMD.Cmd)
	if _, err := cmdRunner.RunCMD(dcpsCMD); err != nil {
		return fmt.Errorf(configs.CommandError, dcpsCMD.Cmd, err)
	}

	return nil
}

func RunValidatorOrExit(cmdRunner commands.CommandRunner, flags *CliCmdFlags) error {
	// notest
	log.Infof(configs.InstructionsFor, "running validator service of docker-compose script")
	upCMD := cmdRunner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     filepath.Join(flags.generationPath, configs.DefaultDockerComposeScriptName),
		Services: []string{validator},
	})
	fmt.Printf("\n%s\n\n", upCMD.Cmd)

	prompt := promptui.Prompt{
		Label:     "Run validator service",
		IsConfirm: true,
		Default:   "Y",
	}
	_, err := prompt.Run()
	if err != nil {
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	if err = runAndShowContainers(cmdRunner, []string{validator}, flags); err != nil {
		return err
	}

	return nil
}

func handleJWTSecret(generationPath string) (string, error) {
	log.Info(configs.GeneratingJWTSecret)

	jwtscret, err := crypto.GenerateJWTSecret()
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	jwtPath, err := filepath.Abs(filepath.Join(generationPath, "jwtsecret"))
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	if err := os.MkdirAll(filepath.Dir(jwtPath), 0o755); err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	err = os.WriteFile(jwtPath, []byte(jwtscret), 0o755)
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	log.Info(configs.JWTSecretGenerated)
	return jwtPath, nil
}

func preRunTeku(services []string, generationPath string) error {
	log.Info(configs.PreparingTekuDatadir)
	// Change umask to avoid OS from changing the permissions
	utils.SetUmask(0)
	for _, s := range services {
		if s == "all" || s == consensus {
			// Prepare consensus datadir
			path := filepath.Join(generationPath, configs.ConsensusDir)
			if err := os.MkdirAll(path, 0o777); err != nil {
				return fmt.Errorf(configs.TekuDatadirError, consensus, err)
			}
		}
		if s == "all" || s == validator {
			// Prepare validator datadir
			path := filepath.Join(generationPath, configs.ValidatorDir)
			if err := os.MkdirAll(path, 0o777); err != nil {
				return fmt.Errorf(configs.TekuDatadirError, validator, err)
			}
		}
	}
	return nil
}

type CustomNetworkConfigsData struct {
	ChainSpecPath          string
	NetworkConfigPath      string
	NetworkGenesisPath     string
	NetworkDeployBlockPath string
}

type CustomFlags struct {
	customTTD           string
	customChainSpec     string
	customNetworkConfig string
	customGenesis       string
	customDeployBlock   string
	customEnodes        []string
	customEnrs          []string
}

func LoadCustomNetworksConfig(flags *CustomFlags, network, generationPath string) (CustomNetworkConfigsData, error) {
	var customNetworkConfigsData CustomNetworkConfigsData
	var chainSpecSrc, networkConfigSrc, genesisSrc, deployBlock string

	networkData, ok := configs.NetworksConfigs()[network]
	if !ok {
		return customNetworkConfigsData, fmt.Errorf(configs.UnknownNetworkError, network)
	}

	eval := func(value, def string) string {
		if value != "" {
			return value
		}
		return def
	}
	chainSpecSrc = eval(flags.customChainSpec, networkData.DefaultCustomChainSpecSrc)
	networkConfigSrc = eval(flags.customNetworkConfig, networkData.DefaultCustomConfigSrc)
	genesisSrc = eval(flags.customGenesis, networkData.DefaultCustomGenesisSrc)
	deployBlock = eval(flags.customDeployBlock, networkData.DefaultCustomDeployBlock)

	// Check if any custom config is needed
	if chainSpecSrc == "" && networkConfigSrc == "" && genesisSrc == "" && deployBlock == "" {
		return customNetworkConfigsData, nil
	}

	destFolder := filepath.Join(generationPath, configs.CustomNetworkConfigsFolder)
	if _, err := os.Stat(destFolder); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(destFolder, os.ModePerm)
			if err != nil {
				return customNetworkConfigsData, err
			}
		} else {
			return customNetworkConfigsData, err
		}
	}

	if chainSpecSrc != "" {
		customNetworkConfigsData.ChainSpecPath = filepath.Join(destFolder, configs.ExecutionNetworkConfigFileName)
		log.Info(configs.GettingCustomChainSpec)
		err := utils.DownloadOrCopy(chainSpecSrc, customNetworkConfigsData.ChainSpecPath, true)
		if err != nil {
			return customNetworkConfigsData, err
		}
		customNetworkConfigsData.ChainSpecPath, err = filepath.Abs(customNetworkConfigsData.ChainSpecPath)
		if err != nil {
			return customNetworkConfigsData, err
		}
	}

	if networkConfigSrc != "" {
		customNetworkConfigsData.NetworkConfigPath = filepath.Join(destFolder, configs.ConsensusNetworkConfigFileName)
		log.Info(configs.GettingCustomNetworkConfig)
		err := utils.DownloadOrCopy(networkConfigSrc, customNetworkConfigsData.NetworkConfigPath, true)
		if err != nil {
			return customNetworkConfigsData, err
		}
		customNetworkConfigsData.NetworkConfigPath, err = filepath.Abs(customNetworkConfigsData.NetworkConfigPath)
		if err != nil {
			return customNetworkConfigsData, err
		}
	}

	if genesisSrc != "" {
		customNetworkConfigsData.NetworkGenesisPath = filepath.Join(destFolder, configs.ConsensusNetworkGenesisFileName)
		log.Info(configs.GettingCustomGenesis)
		err := utils.DownloadOrCopy(genesisSrc, customNetworkConfigsData.NetworkGenesisPath, true)
		if err != nil {
			return customNetworkConfigsData, err
		}
		customNetworkConfigsData.NetworkGenesisPath, err = filepath.Abs(customNetworkConfigsData.NetworkGenesisPath)
		if err != nil {
			return customNetworkConfigsData, err
		}
	}

	if deployBlock != "" {
		customNetworkConfigsData.NetworkDeployBlockPath = filepath.Join(destFolder, configs.ConsensusNetworkDeployBlockFileName)
		log.Info(configs.WritingCustomDeployBlock)
		err := os.WriteFile(customNetworkConfigsData.NetworkDeployBlockPath, []byte(deployBlock), os.ModePerm)
		if err != nil {
			return customNetworkConfigsData, fmt.Errorf(configs.ErrorWritingDeployBlockFile, customNetworkConfigsData.NetworkDeployBlockPath, err)
		}
		customNetworkConfigsData.NetworkDeployBlockPath, err = filepath.Abs(customNetworkConfigsData.NetworkDeployBlockPath)
		if err != nil {
			return customNetworkConfigsData, err
		}
	}

	return customNetworkConfigsData, nil
}
