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
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/NethermindEth/sedge/internal/crypto"
	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/mevboostrelaylist"

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
	lidoNode       bool
)

const (
	execution, consensus, validator, mevBoost = "execution", "consensus", "validator", "mev-boost"
)

type CustomFlags struct {
	customChainSpec     string
	customNetworkConfig string
	customGenesis       string
	customDeployBlock   string
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
	fallbackEL        []string
	elExtraFlags      []string
	clExtraFlags      []string
	vlExtraFlags      []string
	relayURLs         []string
	mevBoostUrl       string
	executionApiUrl   string
	executionAuthUrl  string
	consensusApiUrl   string
	waitEpoch         int
	customEnodes      []string
	customEnrs        []string
	latestVersion     bool
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
- Lido CSM node
`,
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(FullNodeSubCmd(sedgeAction))
	cmd.AddCommand(ExecutionSubCmd(sedgeAction))
	cmd.AddCommand(ConsensusSubCmd(sedgeAction))
	cmd.AddCommand(ValidatorSubCmd(sedgeAction))
	cmd.AddCommand(MevBoostSubCmd(sedgeAction))

	cmd.PersistentFlags().BoolVar(&lidoNode, "lido", false, "generate Lido CSM node")
	cmd.PersistentFlags().StringVarP(&generationPath, "path", "p", configs.DefaultAbsSedgeDataPath, "generation path for sedge data. Default is sedge-data")
	cmd.PersistentFlags().StringVarP(&network, "network", "n", "mainnet", "Target network. e.g. mainnet,sepolia, holesky, gnosis, chiado, etc.")
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

func preValidationGenerateCmd(network, logging string, flags *GenCmdFlags) error {
	// Validate network
	err := validateNetwork(network)
	if err != nil {
		return err
	}

	// Validate Logging flag
	if err = configs.ValidateLoggingFlag(logging); err != nil {
		return err
	}

	// Validate url flags
	singleUriValidator := func(uriType string, validator func([]string) (string, bool)) func([]string) error {
		return func(value []string) error {
			badUri, valid := validator(value)
			if !valid {
				return fmt.Errorf(configs.InvalidUrlFlagError, uriType, badUri)
			}
			return nil
		}
	}

	type uriData struct {
		value     []string
		check     bool
		validator func([]string) error
	}
	toValidate := []uriData{
		{
			value:     []string{flags.executionApiUrl},
			check:     flags.executionApiUrl != "",
			validator: singleUriValidator("execution api", utils.UriValidator),
		},
		{
			value:     []string{flags.executionAuthUrl},
			check:     flags.executionAuthUrl != "",
			validator: singleUriValidator("execution auth", utils.UriValidator),
		},
		{
			value:     []string{flags.consensusApiUrl},
			check:     flags.consensusApiUrl != "",
			validator: singleUriValidator("consensus api", utils.UriValidator),
		},
		{
			value:     []string{flags.mevBoostUrl},
			check:     flags.mevBoostUrl != "",
			validator: singleUriValidator("mev-boost endpoint", utils.UriValidator),
		},
		{
			value:     []string{flags.checkpointSyncUrl},
			check:     flags.checkpointSyncUrl != "",
			validator: singleUriValidator("checkpoint sync", utils.UriValidator),
		},
		{
			value:     flags.relayURLs,
			check:     len(flags.relayURLs) > 0,
			validator: singleUriValidator("relay", utils.UriValidator),
		},
		{
			value:     flags.fallbackEL,
			check:     len(flags.fallbackEL) > 0,
			validator: singleUriValidator("fallback execution", utils.UriValidator),
		},
		{
			value:     flags.customEnodes,
			check:     len(flags.customEnodes) > 0,
			validator: utils.ENodesValidator,
		},
		{
			value:     flags.customEnrs,
			check:     len(flags.customEnrs) > 0,
			validator: utils.ENRValidator,
		},
	}
	for _, uri := range toValidate {
		if uri.check {
			if err := uri.validator(uri.value); err != nil {
				return err
			}
		}
	}

	// Validate Graffiti flag
	if len(flags.graffiti) > 16 {
		return fmt.Errorf(configs.ErrGraffitiLength, flags.graffiti, len(flags.graffiti))
	}

	return nil
}

func runGenCmd(out io.Writer, flags *GenCmdFlags, sedgeAction actions.SedgeActions, services []string) error {
	// Warn if ports are being exposed
	if flags.mapAllPorts {
		log.Warn(configs.MapAllPortsWarning)
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
	if flags.jwtPath == "" {
		flags.jwtPath, err = handleJWTSecret(generationPath)
		if err != nil {
			return err
		}
	} else {
		flags.jwtPath, err = loadJWTSecret(flags.jwtPath)
		if err != nil {
			return err
		}
	}

	// Overwrite feeRecipient and relayURLs for Lido Node
	if lidoNode {
		feeRecipient := contracts.FeeRecipient[network]
		flags.feeRecipient = feeRecipient.FeeRecipientAddress

		flags.relayURLs, _ = mevboostrelaylist.GetRelaysURI(network)
	}

	// Warning if no fee recipient is set
	if flags.feeRecipient == "" {
		log.Warn(configs.EmptyFeeRecipientError)
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
		Mev:                     !flags.noMev && utils.Contains(services, validator) && utils.Contains(services, consensus) && !flags.noValidator,
		MevImage:                flags.mevImage,
		LoggingDriver:           configs.GetLoggingDriver(logging),
		RelayURLs:               flags.relayURLs,
		MevBoostService:         utils.Contains(services, mevBoost),
		MevBoostEndpoint:        flags.mevBoostUrl,
		Services:                services,
		VLStartGracePeriod:      uint(vlStartGracePeriod.Seconds()),
		ExecutionApiUrl:         flags.executionApiUrl,
		ExecutionAuthUrl:        flags.executionAuthUrl,
		ConsensusApiUrl:         flags.consensusApiUrl,
		ECBootnodes:             flags.customEnodes,
		CCBootnodes:             flags.customEnrs,
		CustomChainSpecPath:     flags.CustomFlags.customChainSpec,
		CustomNetworkConfigPath: flags.CustomFlags.customNetworkConfig,
		CustomGenesisPath:       flags.CustomFlags.customGenesis,
		CustomDeployBlock:       flags.customDeployBlock,
		CustomDeployBlockPath:   flags.CustomFlags.customDeployBlock,
		MevBoostOnValidator:     flags.mevBoostOnVal,
		ContainerTag:            containerTag,
		LatestVersion:           flags.latestVersion,
	}
	_, err = sedgeAction.Generate(actions.GenerateOptions{
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

func handleJWTSecret(generationPath string) (string, error) {
	log.Info(configs.GeneratingJWTSecret)
	if !filepath.IsAbs(generationPath) {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, fmt.Errorf("generation path must be absolute"))
	}

	jwtSecret, err := crypto.GenerateJWTSecret()
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

	err = os.WriteFile(jwtPath, []byte(jwtSecret), 0o755)
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	log.Info(configs.JWTSecretGenerated)
	return jwtPath, nil
}

// TODO: Add unit tests
func loadJWTSecret(from string) (absFrom string, err error) {
	// Ensure from is absolute
	absFrom, err = filepath.Abs(from)
	if err != nil {
		return
	}
	// Check if file exists
	if f, err := os.Stat(absFrom); os.IsNotExist(err) || !f.Mode().IsRegular() {
		return "", fmt.Errorf("jwt secret file does not exist")
	}
	// Validate hex string
	jwtSecret, err := os.ReadFile(absFrom)
	if err != nil {
		return "", fmt.Errorf("could not read jwt secret file")
	}
	decodedJWT, err := hex.DecodeString(string(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("jwt secret is not a valid hex string")
	}
	if len(decodedJWT) != 32 {
		return "", fmt.Errorf("jwt secret must be 32 bytes long")
	}
	return
}

func validateLido(network string, flags *GenCmdFlags) error {
	if !flags.noMev {
		_, ok := mevboostrelaylist.DeployedContractAddresses[network]
		if !ok {
			options := mevboostrelaylist.GetLidoSupportedNetworksMevBoost()
			return fmt.Errorf(configs.InvalidNetworkForLidoMevBoost, options)
		}
	}
	_, ok := contracts.FeeRecipient[network]
	if !ok {
		options := contracts.GetLidoSupportedNetworks()
		return fmt.Errorf(configs.InvalidNetworkForLido, options)
	}

	return nil
}
