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
	"time"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/cli/factory"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/crypto"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	sedgeOpts "github.com/NethermindEth/sedge/internal/pkg/options"
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
	execution, consensus, validator, distributedValidator, mevBoost, optimism, opExecution, taiko, tExecution = "execution", "consensus", "validator", "distributedValidator", "mev-boost", "optimism", "opexecution", "taiko", "texecution"
	jwtPathName                                                                                               = "jwtsecret"
)

type CustomFlags struct {
	customChainSpec     string
	customNetworkConfig string
	customGenesis       string
	customDeployBlock   string
}

type L2Execution struct {
	l2ExecutionName string
	ell2ExtraFlags  []string
}

type OptimismFlags struct {
	optimismName string
	opExtraFlags []string
	chain        string
}

type TaikoFlags struct {
	taikoName       string
	taikoExtraFlags []string
}

// GenCmdFlags is a struct that holds the flags of the generate command
type GenCmdFlags struct {
	CustomFlags
	L2Execution
	OptimismFlags
	TaikoFlags
	executionName            string
	consensusName            string
	validatorName            string
	distributed              bool
	distributedValidatorName string
	checkpointSyncUrl        string
	feeRecipient             string
	noMev                    bool
	mevImage                 string
	mevBoostOnVal            bool
	noValidator              bool
	jwtPath                  string
	graffiti                 string
	mapAllPorts              bool
	fallbackEL               []string
	elExtraFlags             []string
	clExtraFlags             []string
	vlExtraFlags             []string
	dvExtraFlags             []string
	relayURLs                []string
	mevBoostUrl              string
	executionApiUrl          string
	executionAuthUrl         string
	consensusApiUrl          string
	waitEpoch                int
	customEnodes             []string
	customEnrs               []string
	latestVersion            bool
	network                  string
}

// Implement ClientFlags interface
func (flags *GenCmdFlags) GetExecutionName() string {
	return flags.executionName
}

func (flags *GenCmdFlags) GetConsensusName() string {
	return flags.consensusName
}

func (flags *GenCmdFlags) GetValidatorName() string {
	return flags.validatorName
}

func (flags *GenCmdFlags) GetOptimismName() string {
	return flags.optimismName
}

func (flags *GenCmdFlags) GetTaikoName() string {
	return flags.taikoName
}

func (flags *GenCmdFlags) GetL2ExecutionName() string {
	return flags.l2ExecutionName
}

func (flags *GenCmdFlags) GetDistributedValidatorName() string {
	return flags.distributedValidatorName
}

func (flags *GenCmdFlags) GetExecutionApiUrl() string {
	return flags.executionApiUrl
}

func (flags *GenCmdFlags) GetNetwork() string {
	return flags.network
}

func (flags *GenCmdFlags) IsNoValidator() bool {
	return flags.noValidator
}

func valClients(allClients clients.OrderedClients, flags *GenCmdFlags, services []string) (*clients.Clients, error) {
	nodeFactory := factory.NewNodeFactory()
	return nodeFactory.InitializeClients(allClients, flags, services)
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
- Optimism Full Node
- Execution Node
- Consensus Node
- Validator Node
- Mev-Boost Instance
- Lido CSM node
- Poa Node
`,
		Args: cobra.NoArgs,
	}

	cmd.AddCommand(FullNodeSubCmd(sedgeAction))
	cmd.AddCommand(ExecutionSubCmd(sedgeAction))
	cmd.AddCommand(ConsensusSubCmd(sedgeAction))
	cmd.AddCommand(ValidatorSubCmd(sedgeAction))
	cmd.AddCommand(MevBoostSubCmd(sedgeAction))
	cmd.AddCommand(OpFullNodeSubCmd(sedgeAction))
	cmd.AddCommand(TaikoFullNodeSubCmd(sedgeAction))

	cmd.PersistentFlags().BoolVar(&lidoNode, "lido", false, "generate Lido CSM node")
	cmd.PersistentFlags().StringVarP(&generationPath, "path", "p", configs.DefaultAbsSedgeDataPath, "generation path for sedge data. Default is sedge-data")
	cmd.PersistentFlags().StringVarP(&network, "network", "n", "mainnet", "Target network. e.g. mainnet,sepolia, holesky, gnosis, chiado, volta, energyweb, joc-mainnet, joc-testnet, linea-mainnet, linea-sepolia etc.")
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
		flags.jwtPath, err = handleJWTSecret(generationPath, jwtPathName)
		if err != nil {
			return err
		}
	} else {
		flags.jwtPath, err = loadJWTSecret(flags.jwtPath)
		if err != nil {
			return err
		}
	}
	var jwtSecretL2 string
	var sequencerURL string
	// If optimism or taiko is included in the services, generate the jwt secret for it
	if utils.Contains(services, optimism) || utils.Contains(services, taiko) {
		jwtSecretL2, err = handleJWTSecret(generationPath, jwtPathName+"-l2")
		if err != nil {
			return err
		}
	}
	if utils.Contains(services, optimism) {
		sequencerURL, err = configs.GetSequencerURL(network, flags.chain)
		if err != nil {
			return err
		}
	}

	// Overwrite feeRecipient and relayURLs for Lido Node
	if lidoNode {
		opts := sedgeOpts.CreateSedgeOptions(sedgeOpts.LidoNode)
		flags.feeRecipient = opts.FeeRecipient(network)
		flags.relayURLs, _ = opts.RelayURLs(network)
	}

	// Warning if no fee recipient is set
	if flags.feeRecipient == "" {
		log.Warn(configs.EmptyFeeRecipientError)
	}

	vlStartGracePeriod := configs.NetworkEpochTime(network) * time.Duration(flags.waitEpoch)

	var consensusApiUrl string
	var executionApiUrl string
	var executionAuthUrl string
	if combinedClients.Consensus == nil {
		consensusApiUrl = flags.consensusApiUrl
	}

	if combinedClients.Execution == nil {
		executionApiUrl = flags.executionApiUrl
		executionAuthUrl = flags.executionAuthUrl
	}

	if network == configs.NetworkHoodi {
		flags.mevImage = "flashbots/mev-boost:1.9rc3"
	}

	// Generate docker-compose scripts
	gd := generate.GenData{
		ExecutionClient:            combinedClients.Execution,
		ConsensusClient:            combinedClients.Consensus,
		ValidatorClient:            combinedClients.Validator,
		Distributed:                flags.distributed,
		DistributedValidatorClient: combinedClients.DistributedValidator,
		DvExtraFlags:               flags.dvExtraFlags,
		L2ExecutionClient:          combinedClients.L2Execution,
		OptimismClient:             combinedClients.Optimism,
		TaikoClient:                combinedClients.Taiko,
		Network:                    network,
		CheckpointSyncUrl:          flags.checkpointSyncUrl,
		FeeRecipient:               flags.feeRecipient,
		JWTSecretPath:              flags.jwtPath,
		Graffiti:                   flags.graffiti,
		FallbackELUrls:             flags.fallbackEL,
		ElExtraFlags:               flags.elExtraFlags,
		ClExtraFlags:               flags.clExtraFlags,
		VlExtraFlags:               flags.vlExtraFlags,
		ElL2ExtraFlags:             flags.ell2ExtraFlags,
		OpExtraFlags:               flags.opExtraFlags,
		TaikoExtraFlags:            flags.taikoExtraFlags,
		Chain:                      flags.chain,
		Sequencer:                  sequencerURL,
		MapAllPorts:                flags.mapAllPorts,
		Mev:                        !flags.noMev && utils.Contains(services, validator) && utils.Contains(services, consensus) && !flags.noValidator,
		MevImage:                   flags.mevImage,
		LoggingDriver:              configs.GetLoggingDriver(logging),
		RelayURLs:                  flags.relayURLs,
		MevBoostService:            utils.Contains(services, mevBoost),
		MevBoostEndpoint:           flags.mevBoostUrl,
		Services:                   services,
		VLStartGracePeriod:         uint(vlStartGracePeriod.Seconds()),
		ExecutionApiUrl:            executionApiUrl,
		ExecutionAuthUrl:           executionAuthUrl,
		ConsensusApiUrl:            consensusApiUrl,
		ECBootnodes:                flags.customEnodes,
		CCBootnodes:                flags.customEnrs,
		CustomChainSpecPath:        flags.CustomFlags.customChainSpec,
		CustomNetworkConfigPath:    flags.CustomFlags.customNetworkConfig,
		CustomGenesisPath:          flags.CustomFlags.customGenesis,
		CustomDeployBlock:          flags.customDeployBlock,
		CustomDeployBlockPath:      flags.CustomFlags.customDeployBlock,
		MevBoostOnValidator:        flags.mevBoostOnVal,
		ContainerTag:               containerTag,
		LatestVersion:              flags.latestVersion,
		JWTSecretL2:                jwtSecretL2,
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

func handleJWTSecret(generationPath, name string) (string, error) {
	log.Info(configs.GeneratingJWTSecret)
	if !filepath.IsAbs(generationPath) {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, fmt.Errorf("generation path must be absolute"))
	}

	jwtSecret, err := crypto.GenerateJWTSecret()
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	jwtPath, err := filepath.Abs(filepath.Join(generationPath, name))
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

func nodeType() string {
	var nodeType string
	if lidoNode {
		nodeType = sedgeOpts.LidoNode
	} else {
		nodeType = sedgeOpts.EthereumNode
	}
	return nodeType
}
