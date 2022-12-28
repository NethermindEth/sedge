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
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/crypto"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/manifoldco/promptui"
)

func installOrShowInstructions(cmdRunner commands.CommandRunner, pending []string) (err error) {
	// notest
	optInstall, optExit := "Install dependencies", "Exit. You will manage this dependencies on your own"
	prompt := promptui.Select{
		Label: "Select how to proceed with the pending dependencies",
		Items: []string{optInstall, optExit},
	}

	if err = utils.HandleInstructions(cmdRunner, pending, utils.ShowInstructions); err != nil {
		return fmt.Errorf(configs.ShowingInstructionsError, err)
	}
	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf(configs.PromptFailedError, err)
	}

	switch result {
	case optInstall:
		return installDependencies(cmdRunner, pending)
	default:
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	return nil
}

func installDependencies(cmdRunner commands.CommandRunner, pending []string) error {
	if runtime.GOOS != "windows" { // Windows doesn't support docker installation through scripts
		if err := utils.HandleInstructions(cmdRunner, pending, utils.InstallDependency); err != nil {
			return fmt.Errorf(configs.InstallingDependenciesError, err)
		}
	}
	return nil
}

func randomizeClients(allClients clients.OrderedClients) (clients.Clients, error) {
	var executionClient, consensusClient clients.Client
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

func handleJWTSecret(flags *CliCmdFlags) error {
	log.Info(configs.GeneratingJWTSecret)

	jwtscret, err := crypto.GenerateJWTSecret()
	if err != nil {
		return fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	flags.jwtPath, err = filepath.Abs(filepath.Join(flags.generationPath, "jwtsecret"))
	if err != nil {
		return fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	if err := os.MkdirAll(filepath.Dir(flags.jwtPath), os.ModePerm); err != nil {
		return fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	err = os.WriteFile(flags.jwtPath, []byte(jwtscret), os.ModePerm)
	if err != nil {
		return fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	log.Info(configs.JWTSecretGenerated)
	return nil
}

func preRunTeku(flags *CliCmdFlags) error {
	log.Info(configs.PreparingTekuDatadir)
	// Change umask to avoid OS from changing the permissions
	utils.SetUmask(0)
	for _, s := range *flags.services {
		if s == "all" || s == consensus {
			// Prepare consensus datadir
			path := filepath.Join(flags.generationPath, configs.ConsensusDir)
			if err := os.MkdirAll(path, 0o777); err != nil {
				return fmt.Errorf(configs.TekuDatadirError, consensus, err)
			}
		}
		if s == "all" || s == validator {
			// Prepare validator datadir
			path := filepath.Join(flags.generationPath, configs.ValidatorDir)
			if err := os.MkdirAll(path, 0o777); err != nil {
				return fmt.Errorf(configs.TekuDatadirError, validator, err)
			}
		}
	}
	return nil
}
