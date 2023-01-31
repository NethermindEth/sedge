package cli

import (
	"fmt"
	"path/filepath"
	"strings"

	eth2 "github.com/protolambda/zrnt/eth2/configs"

	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/generate"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/cli/prompts"
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

	ExecutionGeth       = "geth"
	ExecutionNethermind = "nethermind"
	ExecutionBesu       = "besu"
	ExecutionErigon     = "erigon"

	ConsensusTeku       = "teku"
	ConsensusLodestar   = "lodestar"
	ConsensusLighthouse = "lighthouse"
	ConsensusPrysm      = "prysm"

	ValidatorLighthouse = "lighthouse"
	ValidatorLodestar   = "lodestar"
	ValidatorPrysm      = "prysm"
	ValidatorTeku       = "teku"

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

func CliCmd(p prompts.Prompt, actions actions.SedgeActions) *cobra.Command {
	o := new(CliCmdOptions)
	cmd := &cobra.Command{
		Use: "cli",
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

func setupFullNode(p prompts.Prompt, o *CliCmdOptions, a actions.SedgeActions) error {
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
		FallbackELUrls: nil, // TODO add to consensus client setup

		ElExtraFlags: nil, // TODO add to prompts
		ClExtraFlags: nil, // TODO add to prompts
		VlExtraFlags: nil, // TODO add to prompts

		MevBoostOnValidator: false, // TODO add to prompts

		ExecutionApiUrl:  "", // TODO what is this?
		ExecutionAuthUrl: "", // TODO what is this?
		ConsensusApiUrl:  "", // TODO what is this?

		CustomDeployBlock:     "", // TODO ask to Carlos
		CustomDeployBlockPath: "", // TODO ask to carlos
	}
	return postGenerate(p, o, a)
}

func setupExecutionNode(p prompts.Prompt, o *CliCmdOptions, actions actions.SedgeActions) error {
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
	log.Error("unimplemented generation")
	return postGenerate(p, o, actions)
}

func setupConsensusNode(p prompts.Prompt, o *CliCmdOptions, actions actions.SedgeActions) error {
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
	log.Error("unimplemented generation")
	return postGenerate(p, o, actions)
}

func setupValidatorNode(p prompts.Prompt, o *CliCmdOptions, actions actions.SedgeActions) error {
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
	log.Error("unimplemented generation")
	return postGenerate(p, o, actions)
}

func setupJWT(p prompts.Prompt, o *CliCmdOptions, skip bool) error {
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

func postGenerate(p prompts.Prompt, o *CliCmdOptions, a actions.SedgeActions) error {
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
	return nil
}

func generateKeystore(p prompts.Prompt, o *CliCmdOptions, a actions.SedgeActions) error {
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
		case SourceTypeRandom:
			// TODO generate random passphrase
			log.Error("unimplemented random passphrase generation")
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

type promptAction func(prompts.Prompt, *CliCmdOptions) error

func runPromptActions(p prompts.Prompt, o *CliCmdOptions, actions ...promptAction) error {
	for _, action := range actions {
		if err := action(p, o); err != nil {
			return err
		}
	}
	return nil
}

func selectNetwork(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.Network, err = p.Select("Select network", NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom)
	return
}

func selectNodeType(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.nodeType, err = p.Select("Select node type", NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator)
	return
}

func selectExecutionClient(p prompts.Prompt, o *CliCmdOptions) (err error) {
	// TODO: support randomize
	selectedExecutionClient, err := p.Select("Select execution client", ExecutionNethermind, ExecutionGeth, ExecutionBesu, ExecutionErigon)
	if err != nil {
		return err
	}
	o.genData.ExecutionClient = &clients.Client{
		Name: selectedExecutionClient,
		Type: "execution",
	}
	return nil
}

func selectConsensusClient(p prompts.Prompt, o *CliCmdOptions) (err error) {
	// TODO: support randomize
	selectedConsensusClient, err := p.Select("Select consensus client", ConsensusLighthouse, ConsensusLodestar, ConsensusPrysm, ConsensusTeku)
	if err != nil {
		return err
	}
	o.genData.ConsensusClient = &clients.Client{
		Name: selectedConsensusClient,
		Type: "consensus",
	}
	return nil
}

func selectValidatorClient(p prompts.Prompt, o *CliCmdOptions) (err error) {
	// TODO: support randomize
	selectedValidatorClient, err := p.Select("Select validator client", ValidatorLighthouse, ValidatorLodestar, ValidatorPrysm, ValidatorTeku)
	if err != nil {
		return err
	}
	o.genData.ValidatorClient = &clients.Client{
		Name: selectedValidatorClient,
		Type: "validator",
	}
	return nil
}

func selectJWTSource(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.jwtSourceType, err = p.Select("Select JWT source", SourceTypeCreate, SourceTypeExisting)
	return
}

func selectJWTSourceOrSkip(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.jwtSourceType, err = p.Select("Select JWT source", SourceTypeCreate, SourceTypeExisting, SourceTypeSkip)
	return
}

func selectKeystoreSource(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.keystoreSourceType, err = p.Select("Select keystore source", SourceTypeCreate, SourceTypeExisting, SourceTypeSkip)
	return
}

func selectKeystoreMnemonicSource(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.keystoreMnemonicSource, err = p.Select("Select mnemonic source", SourceTypeCreate, SourceTypeExisting)
	return
}

func selectKeystorePassphraseSource(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.keystorePassphraseSource, err = p.Select("Select passphrase source", SourceTypeRandom, SourceTypeExisting, SourceTypeCreate)
	return
}

func confirmWithValidator(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.withValidator, err = p.Confirm("Do you want to set up a validator")
	return
}

func confirmExposeAllPorts(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.MapAllPorts, err = p.Confirm("Do you want to expose all ports?")
	return
}

func confirmImportSlashingProtection(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.importSlashingProtection, err = p.Confirm("Do you want to import slashing protection data?")
	return
}

func confirmInstallDependencies(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.installDependencies, err = p.Confirm("Install dependencies?")
	return
}

func inputCustomNetworkConfig(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.CustomNetworkConfigPath, err = p.InputFilePath("Custom Network Config", true)
	return
}

func inputCustomChainSpec(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.CustomChainSpecPath, err = p.InputFilePath("Custom ChainSpec", true)
	return
}

func inputCustomGenesis(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.CustomGenesisPath, err = p.InputFilePath("Custom Genesis", true)
	return
}

func inputCustomTTD(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.CustomTTD, err = p.Input("Custom TTD", false)
	return
}

func inputCustomDeployBlock(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.customDeployBlock, err = p.Input("Custom deploy block", false)
	return
}

func inputExecutionBootNodes(p prompts.Prompt, o *CliCmdOptions) (err error) {
	bootNodesInput, err := p.Input("Execution boot nodes", false)
	if err != nil {
		return err
	}
	bootNodesList := strings.Split(bootNodesInput, ",")
	o.genData.ECBootnodes = &bootNodesList
	return nil
}

func inputConsensusBootNodes(p prompts.Prompt, o *CliCmdOptions) (err error) {
	bootNodesInput, err := p.Input("Consensus boot nodes", false)
	if err != nil {
		return err
	}
	bootNodesList := strings.Split(bootNodesInput, ",")
	o.genData.CCBootnodes = &bootNodesList
	return nil
}

func inputMevImage(p prompts.Prompt, o *CliCmdOptions) (err error) {
	// Default value is set in the template
	o.genData.MevImage, err = p.Input("Mev-Boost image", false)
	return
}

func inputMevBoostEndpoint(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.MevBoostEndpoint, err = p.Input("Mev-Boost endpoint", false)
	return
}

func inputRelayURL(p prompts.Prompt, o *CliCmdOptions) (err error) {
	// TODO add default relay URL value, it is not present in the generate command
	o.genData.RelayURL, err = p.Input("Relay URL", false)
	return
}

func inputGraffiti(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.Graffiti, err = p.Input("Graffiti", false)
	return
}

func inputCheckpointSyncURL(p prompts.Prompt, o *CliCmdOptions) (err error) {
	// Default value is set in the template
	o.genData.CheckpointSyncUrl, err = p.Input("Checkpoint sync URL", false)
	return
}

func inputFeeRecipient(p prompts.Prompt, o *CliCmdOptions) (err error) {
	// TODO: Is necessary to add a default value?
	o.genData.FeeRecipient, err = p.Input("Fee recipient", true)
	return
}

func inputValidatorGracePeriod(p prompts.Prompt, o *CliCmdOptions) (err error) {
	epochs, err := p.InputNumber("Validator grace period (epochs)")
	if err != nil {
		return err
	}
	o.genData.VLStartGracePeriod = uint(epochs * int64(configs.NetworkEpochTime(o.genData.Network)))
	return nil
}

func inputGenerationPath(p prompts.Prompt, o *CliCmdOptions) (err error) {
	// TODO: add default
	o.generationPath, err = p.Input("Generation path", true)
	return
}

func inputJWTPath(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.genData.JWTSecretPath, err = p.Input("JWT path", true)
	return
}

func inputKeystoreMnemonicPath(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.keystoreMnemonic, err = p.Input("Mnemonic path", true)
	return
}

func inputKeystorePassphrasePath(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.keystorePassphrasePath, err = p.Input("Passphrase path", true)
	return
}

func inputKeystorePassphrase(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.keystorePassphrase, err = p.InputHide("Passphrase")
	return
}

func inputWithdrawalAddress(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.withdrawalAddress, err = p.Input("Withdrawal address", false)
	return
}

func inputNumberOfValidators(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.numberOfValidators, err = p.InputNumber("Number of validators")
	return
}

func inputNumberOfExistingValidators(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.existingValidators, err = p.InputNumber("Existing validators")
	return
}

func inputKeystorePath(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.keystorePath, err = p.Input("Keystore path", true)
	return
}

func inputImportSlashingProtectionFrom(p prompts.Prompt, o *CliCmdOptions) (err error) {
	o.slashingProtectionFrom, err = p.InputFilePath("Interchange slashing protection file", true)
	return
}
