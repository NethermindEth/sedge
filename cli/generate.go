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
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/NethermindEth/sedge/internal/crypto"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/NethermindEth/sedge/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Global vars
var (
	generationPath string
	network        string
	logging        string
	containerTag   string
)

const (
	execution, consensus, validator, mevBoost = "execution", "consensus", "validator", "mevboost"
)

type CustomFlags struct {
	customTTD           string
	customChainSpec     string
	customNetworkConfig string
	customGenesis       string
	customDeployBlock   string
	customEnodes        *[]string
	customEnrs          *[]string
}

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
	waitEpoch         int
}

func GenerateCmd(sedgeAction actions.SedgeActions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate [flags]",
		Short: "Generate new setups according to selected options",
		Long: `Generate new setups according to selected options and flags.

It will create a 'docker-compose.yml' and a '.env', which you will need later to run the docker-compose script. You can use 'sedge run' to run the script and start the setup.

You can generate:
- Full Node (execution + consensus + validator)
- Full Node without Validator (execution + consensus)
- Execution Node
- Consensus Node
- Validator Node
- Mev-Boost Instance
`,
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(FullNodeSubCmd(sedgeAction))
	cmd.AddCommand(ExecutionSubCmd(sedgeAction))
	cmd.AddCommand(ConsensusSubCmd(sedgeAction))
	cmd.AddCommand(ValidatorSubCmd(sedgeAction))
	cmd.AddCommand(MevBoostSubCmd(sedgeAction))

	cmd.PersistentFlags().StringVarP(&generationPath, "path", "p", configs.DefaultSedgeDataFolderName, "generation path for sedge data. Default is sedge-data")
	cmd.PersistentFlags().StringVarP(&network, "network", "n", "mainnet", "Target network. e.g. mainnet, goerli, sepolia, etc.")
	cmd.PersistentFlags().StringVar(&logging, "logging", "json", fmt.Sprintf("Docker logging driver used by all the services. Set 'none' to use the default docker logging driver. Possible values: %v", configs.ValidLoggingFlags()))
	cmd.PersistentFlags().StringVar(&containerTag, "container-tag", "", "Container tag to use. If defined, sedge will add to each container and the network, a suffix with the tag. e.g. sedge-validator-client -> sedge-validator-client-<tag>.")
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
	// Warn if ports are being exposed
	if flags.mapAllPorts {
		log.Warn(configs.MapAllPortsWarning)
	}

	// Warn if checkpoint url used
	if flags.checkpointSyncUrl != "" {
		// Check checkpoint url is a valid URL
		_, err := url.ParseRequestURI(flags.checkpointSyncUrl)
		if err != nil {
			return fmt.Errorf(configs.InvalidCkptSyncURL, flags.checkpointSyncUrl)
		}
		log.Warnf(configs.CheckpointUrlUsedWarning, flags.checkpointSyncUrl)
	}

	// Get all supported clients
	c := clients.ClientInfo{Network: network}
	clientsMap, errs := c.Clients(onlyClients(services))
	if len(errs) > 0 {
		return errs[0]
	}

	// Handle selection and validation of clients
	combinedClients, err := valClients(clientsMap, flags, services)
	if err != nil {
		return err
	}

	// Mkdir generation path if doesn't exist
	err = initGenPath(generationPath)
	if err != nil {
		return err
	}

	// Generate jwt secrets if needed
	if flags.jwtPath, err = generateJWTSecret(flags.jwtPath); err != nil {
		return err
	}

	// Warning if no fee recipient is set
	if flags.feeRecipient == "" {
		log.Warn(configs.EmptyFeeRecipientError)
	}

	// Get custom networks configs
	customNetworkConfigsData, err := loadCustomNetworksConfig(&flags.CustomFlags, network, generationPath)
	if err != nil {
		return err
	}

	vlStartGracePeriod := configs.NetworkEpochTime(network) * time.Duration(flags.waitEpoch)

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
		ContainerTag:            containerTag,
	}
	err = sedgeAction.Generate(actions.GenerateOptions{
		GenerationData: gd,
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

	log.Info(configs.GenerationEnd)

	return nil
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
			executionClient.Name = executionParts[0]
			if len(executionParts) > 1 {
				log.Warn(configs.CustomExecutionImagesWarning)
				executionClient.Image = strings.Join(executionParts[1:], ":")
			}
		}
		executionClient.SetImageOrDefault(strings.Join(executionParts[1:], ":"))
		// Patch Geth image if network needs TTD to be set
		if executionClient.Name == "geth" && network != "mainnet" {
			executionClient.Image = "ethereum/client-go:v1.10.26"
		}
		// Patch Erigon image if network needs TTD to be set
		if executionClient.Name == "erigon" && network != "mainnet" {
			executionClient.Image = "thorax/erigon:v2.29.0"
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
			consensusClient.Name = consensusParts[0]
			if len(consensusParts) > 1 {
				log.Warn(configs.CustomConsensusImagesWarning)
				consensusClient.Image = strings.Join(consensusParts[1:], ":")
			}
		}
		consensusClient.SetImageOrDefault(strings.Join(consensusParts[1:], ":"))
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
			validatorClient.Name = validatorParts[0]
			if len(validatorParts) > 1 {
				log.Warn(configs.CustomValidatorImagesWarning)
				validatorClient.Image = strings.Join(validatorParts[1:], ":")

			}
		}
		validatorClient.SetImageOrDefault(strings.Join(validatorParts[1:], ":"))
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

func onlyClients(services []string) []string {
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
		if f, err := os.Stat(jwtPath); os.IsNotExist(err) || !f.Mode().IsRegular() {
			return jwtPath, fmt.Errorf(configs.InvalidJWTSecret, jwtPath)
		}
		if jwtPath, err = filepath.Abs(jwtPath); err != nil {
			return jwtPath, err
		}
	}
	return jwtPath, nil
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

type customNetworkConfigsData struct {
	ChainSpecPath          string
	NetworkConfigPath      string
	NetworkGenesisPath     string
	NetworkDeployBlockPath string
}

func loadCustomNetworksConfig(flags *CustomFlags, network, generationPath string) (customNetworkConfigsData, error) {
	var customNetworkConfigsData customNetworkConfigsData
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
