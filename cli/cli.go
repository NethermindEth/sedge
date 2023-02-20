package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	eth2 "github.com/protolambda/zrnt/eth2/configs"

	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/prompter"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/keystores"
	"github.com/NethermindEth/sedge/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	NetworkMainnet = "mainnet"
	NetworkGoerli  = "goerli"
	NetworkSepolia = "sepolia"
	NetworkGnosis  = "gnosis"
	NetworkChiado  = "chiado"
	NetworkCustom  = "custom"

	NodeTypeFullNode  = "full-node"
	NodeTypeExecution = "execution"
	NodeTypeConsensus = "consensus"
	NodeTypeValidator = "validator"

	Randomize = "randomize"

	SourceTypeExisting = "existing"
	SourceTypeCreate   = "create"
	SourceTypeSkip     = "skip"
	SourceTypeRandom   = "random"
)

type CliCmdOptions struct {
	genData                  generate.GenData
	generationPath           string
	nodeType                 string
	withValidator            bool
	importSlashingProtection bool
	slashingProtectionFrom   string
	customDeployBlock        string
	relayURL                 string
	jwtSourceType            string
	keystoreSourceType       string
	keystorePath             string
	keystoreMnemonicSource   string
	keystoreMnemonic         string
	keystorePassphraseSource string
	keystorePassphrasePath   string
	keystorePassphrase       string
	withdrawalAddress        string
	numberOfValidators       int64
	existingValidators       int64
	installDependencies      bool
}

func CliCmd(p prompter.Prompter, actions actions.SedgeActions) *cobra.Command {
	o := new(CliCmdOptions)
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "generate a node setup interactively",
		Long: `
This command will guide you through the process of setting up one of these node types:

- Full Node (execution + consensus + validator)
- Full Node without Validator (execution + consensus)
- Execution Node
- Consensus Node
- Validator Node

Follow the prompts to select the options you want for your node. At the end of the process, you will
be asked to run the generated setup or not. If you chose to run the setup, it will be executed for you
using docker compose command behind the scenes.
`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := selectNetwork(p, o); err != nil {
				log.Fatal(err)
			}
			if err := selectNodeType(p, o); err != nil {
				log.Fatal(err)
			}
			switch o.nodeType {
			case NodeTypeFullNode:
				o.genData.Services = append(o.genData.Services, "execution", "consensus")
				if err := setupFullNode(p, o, actions); err != nil {
					log.Fatal(err)
				}
			case NodeTypeExecution:
				o.genData.Services = append(o.genData.Services, "execution")
				if err := setupExecutionNode(p, o, actions); err != nil {
					log.Fatal(err)
				}
			case NodeTypeConsensus:
				o.genData.Services = append(o.genData.Services, "consensus")
				if err := setupConsensusNode(p, o, actions); err != nil {
					log.Fatal(err)
				}
			case NodeTypeValidator:
				o.genData.Services = append(o.genData.Services, "validator")
				if err := setupValidatorNode(p, o, actions); err != nil {
					log.Fatal(err)
				}
			}
		},
	}
	return cmd
}

func setupFullNode(p prompter.Prompter, o *CliCmdOptions, a actions.SedgeActions) error {
	if err := confirmWithValidator(p, o); err != nil {
		return err
	}
	if o.genData.Network == NetworkCustom {
		if err := runPromptActions(p, o,
			inputCustomNetworkConfig,
			inputCustomChainSpec,
			inputCustomGenesis,
			inputCustomTTD,
			inputCustomDeployBlock,
			inputExecutionBootNodes,
			inputConsensusBootNodes,
		); err != nil {
			return err
		}
	}
	if o.withValidator {
		o.genData.Services = append(o.genData.Services, "validator")
		if configs.SupportsMEVBoost(o.genData.Network) {
			if err := runPromptActions(p, o,
				inputMevImage,
				inputRelayURL,
			); err != nil {
				return err
			}
		}
		if err := runPromptActions(p, o,
			selectExecutionClient,
			selectConsensusClient,
			selectValidatorClient,
			inputValidatorGracePeriod,
			inputGraffiti,
			inputCheckpointSyncURL,
			inputFeeRecipient,
		); err != nil {
			return err
		}
	} else {
		if err := runPromptActions(p, o,
			selectExecutionClient,
			selectConsensusClient,
			inputCheckpointSyncURL,
			inputFeeRecipient,
		); err != nil {
			return err
		}
	}
	if err := runPromptActions(p, o,
		confirmExposeAllPorts,
		inputGenerationPath,
	); err != nil {
		return err
	}
	if err := setupJWT(p, o, false); err != nil {
		return err
	}
	// Set constant values
	o.genData.Mev = true
	// Call generate action
	if err := a.Generate(actions.GenerateOptions{
		GenerationData: &o.genData,
		GenerationPath: o.generationPath,
	}); err != nil {
		return err
	}
	_ = generate.GenData{
		// TODO add to simple consensus client setup
		ExecutionApiUrl:  "",
		ExecutionAuthUrl: "",

		// TODO: add to simple validator client setup
		ConsensusApiUrl:     "",
		MevBoostOnValidator: false,

		CustomDeployBlock:     "", // TODO add to prompts
		CustomDeployBlockPath: "", // TODO add to prompts
	}
	return postGenerate(p, o, a)
}

func setupExecutionNode(p prompter.Prompter, o *CliCmdOptions, a actions.SedgeActions) error {
	if err := selectExecutionClient(p, o); err != nil {
		return err
	}
	if o.genData.Network == NetworkCustom {
		if err := runPromptActions(p, o,
			inputCustomChainSpec,
			inputCustomTTD,
			inputExecutionBootNodes,
		); err != nil {
			return err
		}
	}
	if err := runPromptActions(p, o,
		confirmExposeAllPorts,
		inputGenerationPath,
	); err != nil {
		return err
	}
	if err := setupJWT(p, o, true); err != nil {
		return err
	}
	if err := a.Generate(actions.GenerateOptions{
		GenerationData: &o.genData,
		GenerationPath: o.generationPath,
	}); err != nil {
		return err
	}
	return postGenerate(p, o, a)
}

func setupConsensusNode(p prompter.Prompter, o *CliCmdOptions, a actions.SedgeActions) error {
	if err := selectConsensusClient(p, o); err != nil {
		return err
	}
	if o.genData.Network == NetworkCustom {
		if err := runPromptActions(p, o,
			inputCustomNetworkConfig,
			inputCustomGenesis,
			inputCustomDeployBlock,
			inputConsensusBootNodes,
		); err != nil {
			return err
		}
	} else {
		if err := inputCheckpointSyncURL(p, o); err != nil {
			return err
		}
		if configs.SupportsMEVBoost(o.genData.Network) {
			if err := inputMevBoostEndpoint(p, o); err != nil {
				return err
			}
		}
	}
	if err := runPromptActions(p, o,
		inputFeeRecipient,
		confirmExposeAllPorts,
		inputGenerationPath,
	); err != nil {
		return err
	}
	if err := setupJWT(p, o, true); err != nil {
		return err
	}
	if err := a.Generate(actions.GenerateOptions{
		GenerationData: &o.genData,
		GenerationPath: o.generationPath,
	}); err != nil {
		return err
	}
	return postGenerate(p, o, a)
}

func setupValidatorNode(p prompter.Prompter, o *CliCmdOptions, a actions.SedgeActions) error {
	if err := selectValidatorClient(p, o); err != nil {
		return err
	}
	if o.genData.Network == NetworkCustom {
		if err := runPromptActions(p, o,
			inputCustomNetworkConfig,
			inputCustomGenesis,
			inputCustomDeployBlock,
			inputConsensusBootNodes,
		); err != nil {
			return err
		}
	}
	if err := runPromptActions(p, o,
		inputGraffiti,
		inputValidatorGracePeriod,
		inputFeeRecipient,
		inputGenerationPath,
	); err != nil {
		return err
	}
	if err := a.Generate(actions.GenerateOptions{
		GenerationData: &o.genData,
		GenerationPath: o.generationPath,
	}); err != nil {
		return err
	}
	return postGenerate(p, o, a)
}

func setupJWT(p prompter.Prompter, o *CliCmdOptions, skip bool) error {
	if skip {
		if err := selectJWTSourceOrSkip(p, o); err != nil {
			return err
		}
	} else {
		if err := selectJWTSource(p, o); err != nil {
			return err
		}
	}
	switch o.jwtSourceType {
	case SourceTypeCreate:
		jwtPath, err := generateJWTSecret("")
		o.genData.JWTSecretPath = jwtPath
		if err != nil {
			return err
		}
	case SourceTypeExisting:
		if err := inputJWTPath(p, o); err != nil {
			return err
		}
	case SourceTypeSkip:
		break
	default:
		return fmt.Errorf("unknown JWT source type %s", o.jwtSourceType)
	}
	return nil
}

func postGenerate(p prompter.Prompter, o *CliCmdOptions, a actions.SedgeActions) error {
	if o.withValidator || o.nodeType == NodeTypeValidator {
		if err := generateKeystore(p, o, a); err != nil {
			return err
		}
	}
	pendingDependencies := utils.CheckDependencies([]string{"docker"})
	if len(pendingDependencies) > 0 {
		supported, unsupported := utils.DependenciesSupported(pendingDependencies)
		if len(unsupported) > 0 {
			log.Warnf("unsupported install dependencies %s", strings.Join(unsupported, " "))
			return nil
		}
		if err := confirmInstallDependencies(p, o); err != nil {
			return err
		}
		if !o.installDependencies {
			for _, s := range supported {
				if err := utils.ShowInstructions(a.GetCommandRunner(), s); err != nil {
					return err
				}
			}
			return nil
		}
		for _, s := range supported {
			if err := utils.InstallDependency(a.GetCommandRunner(), s); err != nil {
				return err
			}
		}
	}
	var services []string
	switch o.nodeType {
	case NodeTypeFullNode:
		services = []string{"execution", "consensus"}
		if o.withValidator {
			services = append(services, "validator")
		}
	case NodeTypeExecution:
		services = []string{"execution"}
	case NodeTypeConsensus:
		services = []string{"consensus"}
	case NodeTypeValidator:
		services = []string{"validator"}
	}
	run, err := p.Confirm("Run services now?", false)
	if err != nil {
		return err
	}
	if run {
		if err := a.SetupContainers(actions.SetupContainersOptions{
			GenerationPath: o.generationPath,
			Services:       services,
		}); err != nil {
			return err
		}
		if err := a.RunContainers(actions.RunContainersOptions{
			GenerationPath: o.generationPath,
			Services:       services,
		}); err != nil {
			return err
		}
		// TODO: Final tips
		log.Error("show final tips")
	}
	return nil
}

func generateKeystore(p prompter.Prompter, o *CliCmdOptions, a actions.SedgeActions) error {
	if err := selectKeystoreSource(p, o); err != nil {
		return err
	}
	switch o.keystoreSourceType {
	case SourceTypeSkip:
		return nil
	case SourceTypeCreate:
		// Get the mnemonic
		if err := selectKeystoreMnemonicSource(p, o); err != nil {
			return err
		}
		switch o.keystoreMnemonicSource {
		case SourceTypeCreate:
			candidate, err := keystores.CreateMnemonic()
			if err != nil {
				return err
			}
			if err := saveMnemonic(a.GetCommandRunner(), candidate); err != nil {
				return err
			}
			o.keystoreMnemonic = candidate
		case SourceTypeExisting:
			if err := inputKeystoreMnemonicPath(p, o); err != nil {
				return err
			}
		}
		// Get the passphrase
		if err := selectKeystorePassphraseSource(p, o); err != nil {
			return err
		}
		switch o.keystorePassphraseSource {
		case SourceTypeExisting:
			if err := inputKeystorePassphrasePath(p, o); err != nil {
				return err
			}
			if passphrase, err := readFileContent(o.keystorePassphrasePath); err != nil {
				return err
			} else {
				o.keystorePassphrase = passphrase
			}
		case SourceTypeCreate:
			if err := inputKeystorePassphrase(p, o); err != nil {
				return err
			}
		}
		if err := runPromptActions(p, o,
			inputWithdrawalAddress,
			inputNumberOfValidators,
			inputNumberOfExistingValidators,
		); err != nil {
			return err
		}
		if err := keystores.CreateKeystores(keystores.ValidatorKeysGenData{
			Mnemonic:    o.keystoreMnemonic,
			Passphrase:  o.keystorePassphrase,
			OutputPath:  filepath.Join(o.generationPath, "keystores"),
			MinIndex:    uint64(o.existingValidators),
			MaxIndex:    uint64(o.existingValidators) + uint64(o.numberOfValidators),
			NetworkName: o.genData.Network,
			ForkVersion: configs.NetworksConfigs()[o.genData.Network].GenesisForkVersion,
			// Constants
			UseUniquePassphrase: true,
			Insecure:            false,
			AmountGwei:          uint64(eth2.Mainnet.MAX_EFFECTIVE_BALANCE),
			AsJsonList:          true,
		}); err != nil {
			return err
		}

	case SourceTypeExisting:
		if err := inputKeystorePath(p, o); err != nil {
			return err
		}
	}
	// TODO call validator import
	log.Error("missing call to import validator keys")
	if err := confirmImportSlashingProtection(p, o); err != nil {
		return err
	}
	if o.importSlashingProtection {
		if err := inputImportSlashingProtectionFrom(p, o); err != nil {
			return err
		}
		err := a.ImportSlashingInterchangeData(actions.SlashingImportOptions{
			ValidatorClient: o.genData.ValidatorClient.Name,
			Network:         o.genData.Network,
			GenerationPath:  o.generationPath,
			From:            o.slashingProtectionFrom,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type promptAction func(prompter.Prompter, *CliCmdOptions) error

func runPromptActions(p prompter.Prompter, o *CliCmdOptions, actions ...promptAction) error {
	for _, action := range actions {
		if err := action(p, o); err != nil {
			return err
		}
	}
	return nil
}

func selectNetwork(p prompter.Prompter, o *CliCmdOptions) error {
	options := []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}
	index, err := p.Select("Select network", "", options)
	if err != nil {
		return err
	}
	o.genData.Network = options[index]
	return nil
}

func selectNodeType(p prompter.Prompter, o *CliCmdOptions) error {
	options := []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}
	index, err := p.Select("Select node type", "", options)
	if err != nil {
		return err
	}
	o.nodeType = options[index]
	return nil
}

func selectExecutionClient(p prompter.Prompter, o *CliCmdOptions) (err error) {
	c := clients.ClientInfo{Network: o.genData.Network}
	supportedClients, err := c.SupportedClients(execution)
	if err != nil {
		return err
	}
	options := append(supportedClients, Randomize)
	index, err := p.Select("Select execution client", "", options)
	if err != nil {
		return err
	}
	selectedExecutionClient := options[index]
	// In case random is selected, select a random client
	if selectedExecutionClient == Randomize {
		randomName, err := clients.RandomClientName(supportedClients)
		if err != nil {
			return err
		}
		selectedExecutionClient = randomName
	}
	o.genData.ExecutionClient = &clients.Client{
		Name: selectedExecutionClient,
		Type: "execution",
	}
	return nil
}

func selectConsensusClient(p prompter.Prompter, o *CliCmdOptions) (err error) {
	c := clients.ClientInfo{Network: o.genData.Network}
	supportedClients, err := c.SupportedClients(consensus)
	if err != nil {
		return err
	}
	options := append(supportedClients, Randomize)
	index, err := p.Select("Select consensus client", "", options)
	if err != nil {
		return err
	}
	selectedConsensusClient := options[index]
	// In case random is selected, select a random client
	if selectedConsensusClient == Randomize {
		randomName, err := clients.RandomClientName(supportedClients)
		if err != nil {
			return err
		}
		selectedConsensusClient = randomName
	}
	o.genData.ConsensusClient = &clients.Client{
		Name: selectedConsensusClient,
		Type: "consensus",
	}
	return nil
}

func selectValidatorClient(p prompter.Prompter, o *CliCmdOptions) (err error) {
	c := clients.ClientInfo{Network: o.genData.Network}
	supportedClients, err := c.SupportedClients(validator)
	if err != nil {
		return err
	}
	options := append(supportedClients, Randomize)
	index, err := p.Select("Select validator client", "", options)
	if err != nil {
		return err
	}
	selectedValidatorClient := options[index]
	// In case random is selected, select a random client
	if selectedValidatorClient == Randomize {
		randomName, err := clients.RandomClientName(supportedClients)
		if err != nil {
			return err
		}
		selectedValidatorClient = randomName
	}
	o.genData.ValidatorClient = &clients.Client{
		Name: selectedValidatorClient,
		Type: "validator",
	}
	return nil
}

func selectJWTSource(p prompter.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeCreate, SourceTypeExisting}
	index, err := p.Select("Select JWT source", "", options)
	if err != nil {
		return err
	}
	o.jwtSourceType = options[index]
	return nil
}

func selectJWTSourceOrSkip(p prompter.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}
	index, err := p.Select("Select JWT source", "", options)
	if err != nil {
		return err
	}
	o.jwtSourceType = options[index]
	return nil
}

func selectKeystoreSource(p prompter.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}
	index, err := p.Select("Select keystore source", "", options)
	if err != nil {
		return err
	}
	o.keystoreSourceType = options[index]
	return nil
}

func selectKeystoreMnemonicSource(p prompter.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeCreate, SourceTypeExisting}
	index, err := p.Select("Select mnemonic source", "", options)
	if err != nil {
		return err
	}
	o.keystoreMnemonicSource = options[index]
	return nil
}

func selectKeystorePassphraseSource(p prompter.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}
	index, err := p.Select("Select passphrase source", "", options)
	if err != nil {
		return err
	}
	o.keystorePassphraseSource = options[index]
	return nil
}

func confirmWithValidator(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.withValidator, err = p.Confirm("Do you want to set up a validator?", false)
	return
}

func confirmExposeAllPorts(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.MapAllPorts, err = p.Confirm("Do you want to expose all ports?", false)
	return
}

func confirmImportSlashingProtection(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.importSlashingProtection, err = p.Confirm("Do you want to import slashing protection data?", false)
	return
}

func confirmInstallDependencies(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.installDependencies, err = p.Confirm("Install dependencies?", false)
	return
}

func inputCustomNetworkConfig(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CustomNetworkConfigPath, err = p.InputFilePath("Custom network config file path", "", true)
	return
}

func inputCustomChainSpec(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CustomChainSpecPath, err = p.InputFilePath("Custom ChainSpec", "", true)
	return
}

func inputCustomGenesis(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CustomGenesisPath, err = p.InputFilePath("Custom Genesis", "", true)
	return
}

func inputCustomTTD(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CustomTTD, err = p.Input("Custom TTD", "", false)
	return
}

func inputCustomDeployBlock(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.customDeployBlock, err = p.Input("Custom deploy block", "", false)
	return
}

func inputExecutionBootNodes(p prompter.Prompter, o *CliCmdOptions) (err error) {
	bootNodesInput, err := p.Input("Execution boot nodes", "", false)
	if err != nil {
		return err
	}
	bootNodesList := strings.Split(bootNodesInput, ",")
	o.genData.ECBootnodes = &bootNodesList
	return nil
}

func inputConsensusBootNodes(p prompter.Prompter, o *CliCmdOptions) (err error) {
	bootNodesInput, err := p.Input("Consensus boot nodes", "", false)
	if err != nil {
		return err
	}
	bootNodesList := strings.Split(bootNodesInput, ",")
	o.genData.CCBootnodes = &bootNodesList
	return nil
}

func inputMevImage(p prompter.Prompter, o *CliCmdOptions) (err error) {
	// Default value is set in the template
	o.genData.MevImage, err = p.Input("Mev-Boost image", "flashbots/mev-boost:latest", false)
	return
}

func inputMevBoostEndpoint(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.MevBoostEndpoint, err = p.Input("Mev-Boost endpoint", "", false)
	return
}

func inputRelayURL(p prompter.Prompter, o *CliCmdOptions) (err error) {
	// TODO add default relay URL value and remove it from the template
	// TODO support multiple relay URLs
	o.genData.RelayURL, err = p.Input("Relay URL", "", false)
	return
}

func inputGraffiti(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.Graffiti, err = p.Input("Graffiti", "", false)
	return
}

func inputCheckpointSyncURL(p prompter.Prompter, o *CliCmdOptions) (err error) {
	// Default value is set in the template
	o.genData.CheckpointSyncUrl, err = p.Input("Checkpoint sync URL", "", false)
	return
}

func inputFeeRecipient(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.FeeRecipient, err = p.EthAddress("Please enter the Fee Recipient address. You can leave it blank and press enter (not recommended)", "")
	return
}

func inputValidatorGracePeriod(p prompter.Prompter, o *CliCmdOptions) (err error) {
	epochs, err := p.InputInt64("Validator grace period (epochs)", 1)
	if err != nil {
		return err
	}
	o.genData.VLStartGracePeriod = uint(epochs * int64(configs.NetworkEpochTime(o.genData.Network)))
	return nil
}

func inputGenerationPath(p prompter.Prompter, o *CliCmdOptions) (err error) {
	// TODO: show default value in the prompt
	o.generationPath, err = p.Input("Generation path", configs.DefaultAbsSedgeDataPath, false)
	if o.generationPath == "" {
		o.generationPath = configs.DefaultAbsSedgeDataPath
	}
	return
}

func inputJWTPath(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.genData.JWTSecretPath, err = p.InputFilePath("JWT path", "", true)
	return
}

func inputKeystoreMnemonicPath(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.keystoreMnemonic, err = p.Input("Mnemonic path", "", true)
	return
}

func inputKeystorePassphrasePath(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.keystorePassphrasePath, err = p.Input("Passphrase path", "", true)
	return
}

func inputKeystorePassphrase(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.keystorePassphrase, err = p.InputSecret("Passphrase")
	return
}

func inputWithdrawalAddress(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.withdrawalAddress, err = p.Input("Withdrawal address", "", false)
	return
}

func inputNumberOfValidators(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.numberOfValidators, err = p.InputInt64("Number of validators", 1)
	return
}

func inputNumberOfExistingValidators(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.existingValidators, err = p.InputInt64("Existing validators", 0)
	return
}

func inputKeystorePath(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.keystorePath, err = p.Input("Keystore path", "", true)
	return
}

func inputImportSlashingProtectionFrom(p prompter.Prompter, o *CliCmdOptions) (err error) {
	o.slashingProtectionFrom, err = p.InputFilePath("Interchange slashing protection file", "", true)
	return
}
