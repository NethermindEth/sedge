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
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/NethermindEth/sedge/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"time"
)

// Global vars
var (
	generationPath string
	network        string
	logging        string
)

// GenCmdFlags is a struct that holds the flags of the generate command
type GenCmdFlags struct {
	CustomFlags
	executionName     string
	consensusName     string
	validatorName     string
	checkpointSyncUrl string
	feeRecipient      string
	noMev             bool
	mevImage          string
	mevBoostOnVal     bool
	noValidator       bool
	jwtPath           string
	graffiti          string
	mapAllPorts       bool
	fallbackEL        *[]string
	elExtraFlags      *[]string
	clExtraFlags      *[]string
	vlExtraFlags      *[]string
	relayURL          string
	mevBoostUrl       string
	executionApiUrl   string
	executionAuthUrl  string
	consensusApiUrl   string
}

func GenerateCmd(sedgeAction actions.SedgeActions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate [flags]",
		Short: "Generate new setups according to selected options",
		Long: `Generate new setups according to selected options and flags.

It will only create a docker-compose script, you will need later to run that script.

You can generate:
- Full Nodes (consensus+execution+validator)
- Full Nodes without Validator (consensus+execution)
- Consensus Node
- Execution Node
- Validator Node
- MevBoost Node
`,
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(FullNodeSubCmd(sedgeAction))
	cmd.AddCommand(ExecutionSubCmd(sedgeAction))
	cmd.AddCommand(ConsensusSubCmd(sedgeAction))
	cmd.AddCommand(ValidatorSubCmd(sedgeAction))
	cmd.AddCommand(MevBoostSubCmd(sedgeAction))

	cmd.PersistentFlags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")
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

func preValidationGenerateCmd(network, logging string) error {
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

func runGenCmd(out io.Writer, flags *GenCmdFlags, sedgeAction actions.SedgeActions, services []string) error {

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
	clientsMap, errs := c.Clients(lessMevBoost(services))
	if len(errs) > 0 {
		return errs[0]
	}

	// Handle selection and validation of clients
	combinedClients, err := valClients(clientsMap, flags, services)
	if err != nil {
		return err
	}

	err = initGenPath(generationPath)
	if err != nil {
		return err
	}

	if flags.jwtPath, err = generateJWTSecret(flags.jwtPath); err != nil {
		return err
	}

	// Get fee recipient
	if flags.feeRecipient == "" {
		log.Warn(configs.EmptyFeeRecipientError)
	}

	// Get custom networks configs
	customNetworkConfigsData, err := LoadCustomNetworksConfig(&flags.CustomFlags, network, generationPath)
	if err != nil {
		return err
	}

	var vlStartGracePeriod time.Duration
	switch network {
	case "mainnet", "goerli", "sepolia":
		vlStartGracePeriod = 2 * configs.EpochTimeETH
	case "gnosis", "chiado":
		vlStartGracePeriod = 2 * configs.EpochTimeGNO
	default:
		vlStartGracePeriod = 2 * configs.EpochTimeETH
	}
	// Generate docker-compose scripts
	gd := generate.GenData{
		ExecutionClient:         combinedClients.Execution,
		ConsensusClient:         combinedClients.Consensus,
		ValidatorClient:         combinedClients.Validator,
		Network:                 network,
		CheckpointSyncUrl:       flags.checkpointSyncUrl,
		FeeRecipient:            flags.feeRecipient,
		JWTSecretPath:           flags.jwtPath,
		Graffiti:                flags.graffiti,
		FallbackELUrls:          flags.fallbackEL,
		ElExtraFlags:            flags.elExtraFlags,
		ClExtraFlags:            flags.clExtraFlags,
		VlExtraFlags:            flags.vlExtraFlags,
		MapAllPorts:             flags.mapAllPorts,
		Mev:                     !flags.noMev && utils.Contains(services, validator) && !flags.noValidator,
		MevImage:                flags.mevImage,
		LoggingDriver:           configs.GetLoggingDriver(logging),
		RelayURL:                flags.relayURL,
		MevBoostService:         utils.Contains(services, mevBoost),
		MevBoostEndpoint:        flags.mevBoostUrl,
		Services:                services,
		VLStartGracePeriod:      uint(vlStartGracePeriod.Seconds()),
		ExecutionApiUrl:         flags.executionApiUrl,
		ExecutionAuthUrl:        flags.executionAuthUrl,
		ConsensusApiUrl:         flags.consensusApiUrl,
		ECBootnodes:             flags.customEnodes,
		CCBootnodes:             flags.customEnrs,
		CustomTTD:               flags.customTTD,
		CustomChainSpecPath:     customNetworkConfigsData.ChainSpecPath,
		CustomNetworkConfigPath: customNetworkConfigsData.NetworkConfigPath,
		CustomGenesisPath:       customNetworkConfigsData.NetworkGenesisPath,
		CustomDeployBlock:       flags.customDeployBlock,
		CustomDeployBlockPath:   customNetworkConfigsData.NetworkDeployBlockPath,
		MevBoostOnValidator:     flags.mevBoostOnVal,
	}
	err = sedgeAction.GenerateCompose(actions.GenerateComposeOptions{
		GenerationData: &gd,
		GenerationPath: generationPath,
	})
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
	if combinedClients.Consensus != nil && combinedClients.Consensus.Name == "teku" {
		if err = preRunTeku(services, generationPath); err != nil {
			return err
		}
	}
	log.Info(configs.GenerationEnd)

	return nil
}

func lessMevBoost(services []string) []string {
	newServices := make([]string, 0)
	for _, service := range services {
		if service != mevBoost {
			newServices = append(newServices, service)
		}
	}
	return newServices
}

func initGenPath(path string) error {
	// Create directory if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0o755)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateJWTSecret(jwtPath string) (string, error) {
	// Generate JWT secret if necessary
	var err error
	if jwtPath == "" && configs.NetworksConfigs()[network].RequireJWT {
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
