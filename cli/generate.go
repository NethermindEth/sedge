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
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	log "github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"strings"

	"github.com/NethermindEth/sedge/cli/prompts"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/spf13/cobra"
)

// Global vars
var (
	install        bool
	generationPath string
	network        string
	logging        string
)

// GenCmdFlags is a struct that holds the flags of the generate command
type GenCmdFlags struct {
	executionName     string
	consensusName     string
	validatorName     string
	checkpointSyncUrl string
	feeRecipient      string
	noMev             bool
	mevImage          string
	noValidator       bool
	jwtPath           string
	graffiti          string
	mapAllPorts       bool
	fallbackEL        []string
	elExtraFlags      []string
	clExtraFlags      []string
	vlExtraFlags      []string
	relayURL          string
	mevBoostUrl       string
}

func GenerateCmd(prompt prompts.Prompt) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate [flags]",
		Short: "Generate new setups according to selected options",
		Long: `Run the setup tool on-premise in a quick way. Provide only the command line
	options and the tool will do all the work.
	
	First it will check if dependencies such as docker are installed on your machine
	and provide instructions for installing them if they are not installed.
	
	Second, it will generate docker-compose scripts to run the selected setup.`,
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(FullNodeSubCmd(prompt))
	cmd.AddCommand(ExecutionSubCmd(prompt))
	cmd.AddCommand(ConsensusSubCmd(prompt))
	cmd.AddCommand(ValidatorSubCmd(prompt))
	cmd.AddCommand(MevBoostSubCmd(prompt))

	cmd.PersistentFlags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")
	cmd.PersistentFlags().BoolVarP(&install, "install", "i", false, "Install dependencies if not installed without asking")
	cmd.PersistentFlags().StringVarP(&network, "network", "n", "mainnet", "Target network. e.g. mainnet, goerli, sepolia, etc.")
	cmd.PersistentFlags().StringVar(&logging, "logging", "json", fmt.Sprintf("Docker logging driver used by all the services. Set 'none' to use the default docker logging driver. Possible values: %v", configs.ValidLoggingFlags()))
	return cmd
}

func validateNetwork(network string) error {
	networks, err := utils.SupportedNetworks()
	if err != nil {
		return fmt.Errorf(configs.NetworkValidationFailedError, err)
	}
	if !utils.Contains(networks, network) {
		return fmt.Errorf(configs.UnknownNetworkError, network)
	}
	return nil
}

func preValidationGenerateCmd(flags *GenCmdFlags) error {
	// Validate network
	err := validateNetwork(network)
	if err != nil {
		return err
	}

	// Validate Logging flag
	if err = configs.ValidateLoggingFlag(logging); err != nil {
		return err
	}

	return nil
}

func runGenCmd(out io.Writer, flags *GenCmdFlags, prompt prompts.Prompt, services []string) error {

	// Warn if exposed ports are used
	if flags.mapAllPorts {
		log.Warn(configs.MapAllPortsWarning)
	}

	// Warn if checkpoint url used
	if flags.checkpointSyncUrl != "" {
		log.Warnf(configs.CheckpointUrlUsedWarning, flags.checkpointSyncUrl)
	}

	// Get all clients: supported + configured
	c := clients.ClientInfo{Network: network}
	clientsMap, errs := c.Clients(services)
	if len(errs) > 0 {
		return errs[0]
	}

	// Handle selection and validation of clients
	combinedClients, err := valClients(clientsMap, flags, services)
	if err != nil {
		return err
	}

	if err = checkDependencies(DependenciesOptions{Install: install}); err != nil {
		return err
	}

	if flags.jwtPath, err = generateJWTSecret(flags.jwtPath); err != nil {
		return err
	}

	// Get fee recipient
	feeRecipient := flags.feeRecipient
	if feeRecipient == "" {
		if feeRecipient, err = prompt.FeeRecipient(); err != nil {
			return err
		}
	}

	// Generate docker-compose scripts
	gd := generate.GenData{
		ExecutionClient:   &combinedClients.Execution,
		ConsensusClient:   &combinedClients.Consensus,
		ValidatorClient:   &combinedClients.Validator,
		Network:           network,
		CheckpointSyncUrl: flags.checkpointSyncUrl,
		FeeRecipient:      feeRecipient,
		JWTSecretPath:     flags.jwtPath,
		Graffiti:          flags.graffiti,
		FallbackELUrls:    flags.fallbackEL,
		ElExtraFlags:      flags.elExtraFlags,
		ClExtraFlags:      flags.clExtraFlags,
		VlExtraFlags:      flags.vlExtraFlags,
		MapAllPorts:       flags.mapAllPorts,
		Mev:               !flags.noMev && utils.Contains(services, validator) && !flags.noValidator,
		MevImage:          flags.mevImage,
		LoggingDriver:     configs.GetLoggingDriver(logging),
		RelayURL:          flags.relayURL,
		MevBoostService:   utils.Contains(services, mevBoost),
		MevBoostEndpoint:  flags.mevBoostUrl,
	}
	err = generate.GenerateDockerComposeAndEnvFile(&gd, generationPath)
	if err != nil {
		return err
	}

	// Clean generated .env and docker compose files
	err = generate.CleanGenerated(generationPath)
	if err != nil {
		return err
	}

	// Print Env File
	log.Infof(configs.CreatedFile, filepath.Join(generationPath, configs.DefaultEnvFileName))
	if err = ui.PrintFileContent(out, filepath.Join(generationPath, configs.DefaultEnvFileName)); err != nil {
		return err
	}

	// Print Docker Compose File
	log.Infof(configs.CreatedFile, filepath.Join(generationPath, configs.DefaultDockerComposeScriptName))
	if err = ui.PrintFileContent(out, filepath.Join(generationPath, configs.DefaultDockerComposeScriptName)); err != nil {
		return err
	}

	// If teku is chosen, then prepare datadir with 777 permissions
	if combinedClients.Consensus.Name == "teku" {
		if err = preRunTeku([]string{}, generationPath); err != nil {
			return err
		}
	}

	return nil
}

func generateJWTSecret(jwtPath string) (string, error) {
	// Generate JWT secret if necessary
	var err error
	if jwtPath == "" && configs.NetworksConfigs[network].RequireJWT {
		if jwtPath, err = handleJWTSecret(generationPath); err != nil {
			return jwtPath, err
		}
	} else if filepath.IsAbs(jwtPath) { // Ensure jwtPath is absolute
		if jwtPath, err = filepath.Abs(jwtPath); err != nil {
			return jwtPath, err
		}
	}
	return jwtPath, nil
}

type DependenciesOptions struct {
	Install bool
}

func checkDependencies(options DependenciesOptions) error {
	dependencies := configs.GetDependencies()
	log.Infof(configs.CheckingDependencies, strings.Join(dependencies, ", "))

	// Check if dependencies are installed. Keep checking dependencies until they are all installed
	for pending := utils.CheckDependencies(dependencies); len(pending) > 0; pending = utils.CheckDependencies(dependencies) {
		log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
		if options.Install {
			// Install dependencies directly
			if err := installDependencies(pending); err != nil {
				return err
			}
		} else {
			// Let the user decide to see the instructions for installing dependencies and exit or let the tool install them and continue
			if err := installOrShowInstructions(pending); err != nil {
				return err
			}
		}
	}
	log.Info(configs.DependenciesOK)
	return nil
}
