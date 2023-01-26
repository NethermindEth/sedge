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
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/cli/prompts"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/NethermindEth/sedge/internal/utils"
	dockerct "github.com/docker/docker/api/types/container"
	"github.com/spf13/cobra"
)

const (
	slashingImportFile                        = "slashing-import.json"
	execution, consensus, validator, mevBoost = "execution", "consensus", "validator", "mevboost"
)

type CliCmdFlags struct {
	CustomFlags
	executionName      string
	consensusName      string
	validatorName      string
	generationPath     string
	checkpointSyncUrl  string
	network            string
	feeRecipient       string
	noMev              bool
	mevImage           string
	noValidator        bool
	jwtPath            string
	graffiti           string
	install            bool
	run                bool
	yes                bool
	mapAllPorts        bool
	services           *[]string
	fallbackEL         *[]string
	elExtraFlags       *[]string
	clExtraFlags       *[]string
	vlExtraFlags       *[]string
	logging            string
	slashingProtection string
}

type clientImages struct {
	execution string
	consensus string
	validator string
}

func CliCmd(cmdRunner commands.CommandRunner, prompt prompts.Prompt, serviceManager services.ServiceManager, sedgeActions actions.SedgeActions) *cobra.Command {
	var (
		flags  CliCmdFlags
		images clientImages
	)

	cmd := &cobra.Command{
		Use:   "cli [flags]",
		Short: "Quick start sedge",
		Long: `Run the setup tool on-premise in a quick way. Provide only the command line
	options and the tool will do all the work.
	
	First it will check if dependencies such as docker are installed on your machine
	and provide instructions for installing them if they are not installed.
	
	Second, it will generate docker-compose scripts to run the full setup according to your selection.
	
	Finally, it will run the generated docker-compose script. Only execution and consensus clients will be executed by default.`,
		Args: cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// notest
			clientImages, err := preRunCliCmd(cmd, args, &flags)
			if err != nil {
				return err
			}
			images = *clientImages
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// notest
			if errs := runCliCmd(cmd, args, &flags, &images, cmdRunner, prompt, serviceManager, sedgeActions); len(errs) > 0 {
				for _, err := range errs {
					log.Error(err)
				}
				os.Exit(1)
			}
		},
	}
	// Bind flags
	cmd.Flags().StringVarP(&flags.consensusName, "consensus", "c", "", "Consensus engine client, e.g. teku, lodestar, prysm, lighthouse, Nimbus. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVarP(&flags.executionName, "execution", "e", "", "Execution engine client, e.g. geth, nethermind, besu, erigon. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVarP(&flags.validatorName, "validator", "v", "", "Validator engine client, e.g. teku, lodestar, prysm, lighthouse, Nimbus. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVarP(&flags.generationPath, "path", "p", configs.DefaultAbsSedgeDataPath, "generation path for sedge data")
	cmd.Flags().StringVar(&flags.checkpointSyncUrl, "checkpoint-sync-url", "", "Initial state endpoint (trusted synced consensus endpoint) for the consensus client to sync from a finalized checkpoint. Provide faster sync process for the consensus client and protect it from long-range attacks affored by Weak Subjetivity")
	cmd.Flags().StringVarP(&flags.network, "network", "n", "mainnet", "Target network. e.g. mainnet, goerli, sepolia, etc.")
	cmd.Flags().StringVar(&flags.feeRecipient, "fee-recipient", "", "Suggested fee recipient. Is a 20-byte Ethereum address which the execution layer might choose to set as the coinbase and the recipient of other fees or rewards. There is no guarantee that an execution node will use the suggested fee recipient to collect fees, it may use any address it chooses. It is assumed that an honest execution node will use the suggested fee recipient, but users should note this trust assumption")
	cmd.Flags().BoolVar(&flags.noMev, "no-mev-boost", false, "Not use mev-boost if supported")
	cmd.Flags().StringVarP(&flags.mevImage, "mev-boost-image", "m", "", "Custom docker image to use for Mev Boost. Example: 'sedge cli --mev-boost-image flashbots/mev-boost:latest-portable'")
	cmd.Flags().BoolVar(&flags.noValidator, "no-validator", false, "Exclude the validator from the full node setup. Designed for execution and consensus nodes setup without a validator node. Exclude also the validator from other flags. If set, mev-boost will not be used.")
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().StringVar(&flags.graffiti, "graffiti", "", "Graffiti to be used by the validator")
	cmd.Flags().BoolVarP(&flags.install, "install", "i", false, "Install dependencies if not installed without asking")
	cmd.Flags().BoolVarP(&flags.run, "run", "r", false, "Run the generated docker-compose scripts without asking")
	cmd.Flags().BoolVarP(&flags.yes, "yes", "y", false, "Shortcut for 'sedge cli -r -i --run'. Run without prompts")
	cmd.Flags().BoolVar(&flags.mapAllPorts, "map-all", false, "Map all clients ports to host. Use with care. Useful to allow remote access to the clients")
	flags.services = cmd.Flags().StringSlice("run-clients", []string{execution, consensus}, "Run only the specified clients. Possible values: execution, consensus, validator, all, none. The 'all' and 'none' option must be used alone. Example: 'sedge cli -r --run-clients=consensus,validator'")
	flags.fallbackEL = cmd.Flags().StringSlice("fallback-execution-urls", []string{}, "Fallback/backup execution endpoints for the consensus client. Not supported by Teku. Example: 'sedge cli -r --fallback-execution=https://mainnet.infura.io/v3/YOUR-PROJECT-ID,https://eth-mainnet.alchemyapi.io/v2/YOUR-PROJECT-ID'")
	flags.elExtraFlags = cmd.Flags().StringArray("el-extra-flag", []string{}, "Additional flag to configure the execution client service in the generated docker-compose script. Example: 'sedge cli --el-extra-flag \"<flag1>=value1\" --el-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	flags.clExtraFlags = cmd.Flags().StringArray("cl-extra-flag", []string{}, "Additional flag to configure the consensus client service in the generated docker-compose script. Example: 'sedge cli --cl-extra-flag \"<flag1>=value1\" --cl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	flags.vlExtraFlags = cmd.Flags().StringArray("vl-extra-flag", []string{}, "Additional flag to configure the validator client service in the generated docker-compose script. Example: 'sedge cli --vl-extra-flag \"<flag1>=value1\" --vl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	cmd.Flags().StringVar(&flags.logging, "logging", "json", fmt.Sprintf("Docker logging driver used by all the services. Set 'none' to use the default docker logging driver. Possible values: %v", configs.ValidLoggingFlags()))
	cmd.Flags().StringVar(&flags.customTTD, "custom-ttd", "", "Custom Terminal Total Difficulty to use for the execution client")
	cmd.Flags().StringVar(&flags.customChainSpec, "custom-chainSpec", "", "File path or url to use as custom network chainSpec for execution client.")
	cmd.Flags().StringVar(&flags.customNetworkConfig, "custom-config", "", "File path or url to use as custom network config file for consensus client.")
	cmd.Flags().StringVar(&flags.customGenesis, "custom-genesis", "", "File path or url to use as custom network genesis for consensus client.")
	cmd.Flags().StringVar(&flags.customDeployBlock, "custom-deploy-block", "", "Custom network deploy block to use for consensus client.")
	flags.customEnodes = cmd.Flags().StringSlice("execution-bootnodes", []string{}, "List of comma separated enodes to use as custom network peers for execution client.")
	flags.customEnrs = cmd.Flags().StringSlice("consensus-bootnodes", []string{}, "List of comma separated enrs to use as custom network peers for consensus client.")
	cmd.Flags().StringVar(&flags.slashingProtection, "slashing-protection", "", "Path to the file with slashing protection interchange data (EIP-3076)")
	cmd.Flags().SortFlags = false
	return cmd
}

func preRunCliCmd(cmd *cobra.Command, args []string, flags *CliCmdFlags) (*clientImages, error) {
	// Quick run
	if flags.yes {
		// TODO: avoid flag edition
		flags.install, flags.run = true, true
	}

	// Validate run-clients flag
	if utils.Contains(*flags.services, "all") {
		if len(*flags.services) == 1 {
			// all used correctly
			// TODO: avoid edit flags
			flags.services = &[]string{execution, consensus, validator}
		} else {
			// Ambiguous value
			return nil, fmt.Errorf(configs.RunClientsFlagAmbiguousError, *flags.services)
		}
	} else if utils.Contains(*flags.services, "none") {
		if len(*flags.services) == 1 {
			// all used correctly
			// TODO: avoid edit flags
			flags.services = &[]string{}
		} else {
			// Ambiguous value
			return nil, fmt.Errorf(configs.RunClientsFlagAmbiguousError, *flags.services)
		}
	} else if !utils.ContainsOnly(*flags.services, []string{execution, consensus, validator}) {
		return nil, fmt.Errorf(configs.RunClientsError, strings.Join(*flags.services, ","), strings.Join([]string{execution, consensus, validator}, ","))
	}
	// Exclude validator from run-clients if no-validator flag is set
	if flags.noValidator && utils.Contains(*flags.services, validator) {
		*flags.services = utils.Filter(*flags.services, func(s string) bool {
			return s != validator
		})
	}

	// Validate network
	networks, err := utils.SupportedNetworks()
	if err != nil {
		return nil, fmt.Errorf(configs.NetworkValidationFailedError, err)
	}
	if !utils.Contains(networks, flags.network) {
		return nil, fmt.Errorf(configs.UnknownNetworkError, flags.network)
	}
	if flags.network == configs.CustomNetwork.Name {
		if flags.customChainSpec == "" || flags.customNetworkConfig == "" || flags.customGenesis == "" || flags.customTTD == "" || flags.customDeployBlock == "" {
			return nil, fmt.Errorf(configs.MissingCustomConfigs)
		}
	}

	// Validate custom ttd
	if flags.customTTD != "" &&
		!regexp.MustCompile(`^[1-9]\d*$`).Match([]byte(strings.TrimSpace(flags.customTTD))) {
		return nil, fmt.Errorf(configs.InvalidTTD)
	}

	// Validate custom deploy block
	if flags.customDeployBlock != "" &&
		!regexp.MustCompile(`^[0-9]\d*$`).Match([]byte(strings.TrimSpace(flags.customDeployBlock))) {
		return nil, fmt.Errorf(configs.InvalidDeployBLock)
	}

	// Validate fee recipient
	if flags.feeRecipient != "" && !utils.IsAddress(flags.feeRecipient) {
		return nil, fmt.Errorf(configs.InvalidFeeRecipientError)
	}

	var clientImages clientImages

	// Prepare custom images
	if flags.executionName != "" {
		executionParts := strings.Split(flags.executionName, ":")
		// TODO: avoid edit flag
		flags.executionName = executionParts[0]
		clientImages.execution = strings.Join(executionParts[1:], ":")
	}
	if flags.consensusName != "" {
		consensusParts := strings.Split(flags.consensusName, ":")
		// TODO: avoid edit flag
		flags.consensusName = consensusParts[0]
		clientImages.consensus = strings.Join(consensusParts[1:], ":")
	}
	if flags.validatorName != "" {
		validatorParts := strings.Split(flags.validatorName, ":")
		// TODO: avoid edit flag
		flags.validatorName = validatorParts[0]
		clientImages.validator = strings.Join(validatorParts[1:], ":")
	}

	if err := configs.ValidateLoggingFlag(flags.logging); err != nil {
		return nil, err
	}

	// validate custom network flags
	urlOrPaths := map[string]string{
		"ChainSpec":      flags.customChainSpec,
		"Network config": flags.customNetworkConfig,
		"Genesis":        flags.customGenesis,
	}
	for kind, value := range urlOrPaths {
		if value == "" {
			continue
		}
		if err := utils.CheckUrlOrPath(value); err != nil {
			return nil, fmt.Errorf("invalid %s: %w", kind, err)
		}
	}

	// ensure generation path exists
	_, err = os.Stat(flags.generationPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(flags.generationPath, os.ModePerm)
	}
	if err != nil {
		return nil, fmt.Errorf(configs.InvalidGenerationPath, flags.generationPath, err)
	}

	return &clientImages, nil
}

func runCliCmd(cmd *cobra.Command, args []string, flags *CliCmdFlags, clientImages *clientImages, cmdRunner commands.CommandRunner, prompt prompts.Prompt, serviceManager services.ServiceManager, sedgeActions actions.SedgeActions) []error {
	// Warnings
	// Warn if custom images are used
	if clientImages.execution != "" {
		log.Warn(configs.CustomExecutionImagesWarning)
	}
	if clientImages.consensus != "" {
		log.Warn(configs.CustomConsensusImagesWarning)
	}
	if clientImages.validator != "" {
		log.Warn(configs.CustomValidatorImagesWarning)
	}
	// Warn if exposed ports are used
	if flags.mapAllPorts {
		log.Warn(configs.MapAllPortsWarning)
	}

	// Warn if checkpoint url used
	if flags.checkpointSyncUrl != "" {
		log.Warnf(configs.CheckpointUrlUsedWarning, flags.checkpointSyncUrl)
	}

	// Get all clients: supported + configured
	c := clients.ClientInfo{Network: flags.network}
	clientsMap, clientsErrors := c.Clients([]string{execution, consensus, validator})
	if len(clientsErrors) > 0 {
		return clientsErrors
	}

	// Handle selection and validation of clients
	combinedClients, err := validateClients(clientsMap, cmd.OutOrStdout(), flags)
	if err != nil {
		return []error{err}
	}

	if err := sedgeActions.InstallDependencies(actions.InstallDependenciesOptions{
		Dependencies: configs.GetDependencies(),
		Install:      flags.install,
	}); err != nil {
		return []error{err}
	}

	// Generate JWT secret if necessary
	jwtPath := flags.jwtPath
	jwtPath, err = sedgeActions.CreateJWTSecrets(actions.CreateJWTSecretOptions{
		JWTPath:        jwtPath,
		Network:        flags.network,
		GenerationPath: flags.generationPath,
	})
	if err != nil {
		return []error{err}
	}

	// Get fee recipient
	if !flags.yes && flags.feeRecipient == "" {
		feeRecipient, err := prompt.FeeRecipient()
		if err != nil {
			return []error{err}
		}
		// TODO: avoid flag edition
		flags.feeRecipient = feeRecipient
	}

	// Get custom networks configs
	customNetworkConfigsData, err := LoadCustomNetworksConfig(&flags.CustomFlags, flags.network, flags.generationPath)
	if err != nil {
		return []error{err}
	}

	if combinedClients.Execution != nil {
		combinedClients.Execution.Image = clientImages.execution
	}
	if combinedClients.Consensus != nil {
		combinedClients.Consensus.Image = clientImages.consensus
	}

	if flags.noValidator {
		combinedClients.Validator = nil
	} else {
		if combinedClients.Validator != nil {
			combinedClients.Validator.Image = clientImages.validator
		}
	}

	vlStartGracePeriod := configs.NetworkEpochTime(flags.network)

	// Generate docker-compose scripts
	gd := &generate.GenData{
		Services:                *flags.services,
		ExecutionClient:         combinedClients.Execution,
		ConsensusClient:         combinedClients.Consensus,
		ValidatorClient:         combinedClients.Validator,
		Network:                 flags.network,
		CheckpointSyncUrl:       flags.checkpointSyncUrl,
		FeeRecipient:            flags.feeRecipient,
		JWTSecretPath:           jwtPath,
		Graffiti:                flags.graffiti,
		FallbackELUrls:          flags.fallbackEL,
		ElExtraFlags:            flags.elExtraFlags,
		ClExtraFlags:            flags.clExtraFlags,
		VlExtraFlags:            flags.vlExtraFlags,
		MapAllPorts:             flags.mapAllPorts,
		Mev:                     !flags.noMev && !flags.noValidator,
		MevImage:                flags.mevImage,
		LoggingDriver:           configs.GetLoggingDriver(flags.logging),
		VLStartGracePeriod:      uint(vlStartGracePeriod.Seconds()),
		ECBootnodes:             flags.customEnodes,
		CCBootnodes:             flags.customEnrs,
		CustomTTD:               flags.customTTD,
		CustomChainSpecPath:     customNetworkConfigsData.ChainSpecPath,
		CustomNetworkConfigPath: customNetworkConfigsData.NetworkConfigPath,
		CustomGenesisPath:       customNetworkConfigsData.NetworkGenesisPath,
		CustomDeployBlock:       flags.customDeployBlock,
		CustomDeployBlockPath:   customNetworkConfigsData.NetworkDeployBlockPath,
	}
	err = sedgeActions.Generate(actions.GenerateOptions{GenerationData: gd, GenerationPath: flags.generationPath})
	if err != nil {
		return []error{err}
	}

	// Print final files
	log.Infof(configs.CreatedFile, filepath.Join(generationPath, configs.DefaultEnvFileName))
	ui.PrintFileContent(cmd.OutOrStdout(), filepath.Join(generationPath, configs.DefaultEnvFileName))

	log.Infof(configs.CreatedFile, filepath.Join(generationPath, configs.DefaultDockerComposeScriptName))
	ui.PrintFileContent(cmd.OutOrStdout(), filepath.Join(generationPath, configs.DefaultDockerComposeScriptName))

	// If --run-clients=none was set then exit and don't run anything
	if len(*flags.services) == 0 {
		log.Info(configs.HappyStaking2)
		return nil
	}

	if flags.run {
		if utils.Contains(*flags.services, "validator") {
			*flags.services = append(*flags.services, "validator-import")
		}
		if err := buildContainers(cmdRunner, *flags.services, flags.generationPath); err != nil {
			return []error{err}
		}
		if flags.slashingProtection != "" {
			// Setup wait for validator import
			exitCh, errCh := serviceManager.Wait(services.ServiceCtValidatorImport, dockerct.WaitConditionNextExit)
			// Run validator-import
			if err := runAndShowContainers(cmdRunner, []string{"validator-import"}, flags); err != nil {
				return []error{err}
			}
			exitCode, err := func(exitCh <-chan dockerct.ContainerWaitOKBody, errCh <-chan error) (int64, error) {
				for {
					select {
					case exitOk := <-exitCh:
						return exitOk.StatusCode, nil
					case err := <-errCh:
						return -1, err
					}
				}
			}(exitCh, errCh)
			if err != nil {
				return []error{err}
			}
			if exitCode != 0 {
				return []error{fmt.Errorf("%s ends with unexpected status code %d", services.ServiceCtValidatorImport, exitCode)}
			}
			if err := sedgeActions.ImportSlashingInterchangeData(actions.SlashingImportOptions{
				ValidatorClient: flags.validatorName,
				Network:         flags.network,
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  flags.generationPath,
				From:            flags.slashingProtection,
			}); err != nil {
				return []error{err}
			}
		}
		if err = runAndShowContainers(cmdRunner, *flags.services, flags); err != nil {
			return []error{err}
		}
	} else {
		// Let the user decide to see the instructions for executing the scripts and exit or let the tool execute them
		if err = runScriptOrExit(cmdRunner, flags); err != nil {
			return []error{err}
		}
	}

	return nil
}
