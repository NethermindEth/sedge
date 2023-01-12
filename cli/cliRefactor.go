package cli

import (
	"fmt"
	"strings"

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

type CliRefactorOptions struct {
	generationPath           string
	network                  string
	nodeType                 string
	withValidator            bool
	exposeAllPorts           bool
	importSlashingProtection bool
	slashingProtectionFrom   string
	customNetworkConfig      string
	customChainSpec          string
	customGenesis            string
	customTTD                string
	customDeployBlock        string
	executionBootNodes       string
	consensusBootNodes       string
	mevBoostImage            string
	mevBoostEndpoint         string
	relayURL                 string
	executionClient          string
	consensusClient          string
	validatorClient          string
	graffiti                 string
	ckptSyncURL              string
	feeRecipient             string
	jwtSourceType            string
	jwtPath                  string
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

func CliRefactorCmd(p prompts.Prompt, actions actions.SedgeActions) *cobra.Command {
	o := new(CliRefactorOptions)
	cmd := &cobra.Command{
		Use: "cli-new",
		Run: func(cmd *cobra.Command, args []string) {
			if err := selectNetwork(p, o); err != nil {
				log.Fatal(err)
			}
			if err := selectNodeType(p, o); err != nil {
				log.Fatal(err)
			}
			switch o.nodeType {
			case NodeTypeFullNode:
				if err := setupFullNode(p, o, actions); err != nil {
					log.Fatal(err)
				}
			case NodeTypeExecution:
				if err := setupExecutionNode(p, o, actions); err != nil {
					log.Fatal(err)
				}
			case NodeTypeConsensus:
				if err := setupConsensusNode(p, o, actions); err != nil {
					log.Fatal(err)
				}
			case NodeTypeValidator:
				if err := setupValidatorNode(p, o, actions); err != nil {
					log.Fatal(err)
				}
			}
		},
	}
	return cmd
}

func setupFullNode(p prompts.Prompt, o *CliRefactorOptions, actions actions.SedgeActions) error {
	if err := confirmWithValidator(p, o); err != nil {
		return err
	}
	if o.network == NetworkCustom {
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
		if configs.SupportsMEVBoost(o.network) {
			if err := runPromptActions(p, o,
				inputMevBoostImage,
				inputRelayURL,
			); err != nil {
				return err
			}
		}
		if err := runPromptActions(p, o,
			selectExecutionClient,
			selectConsensusClient,
			selectValidatorClient,
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
	// TODO call generate action
	log.Error("unimplemented generation")
	return postGenerate(p, o, actions)
}

func setupExecutionNode(p prompts.Prompt, o *CliRefactorOptions, actions actions.SedgeActions) error {
	if err := selectExecutionClient(p, o); err != nil {
		return err
	}
	if o.network == NetworkCustom {
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

func setupConsensusNode(p prompts.Prompt, o *CliRefactorOptions, actions actions.SedgeActions) error {
	if err := selectConsensusClient(p, o); err != nil {
		return err
	}
	if o.network == NetworkCustom {
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
		if configs.SupportsMEVBoost(o.network) {
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

func setupValidatorNode(p prompts.Prompt, o *CliRefactorOptions, actions actions.SedgeActions) error {
	if err := selectValidatorClient(p, o); err != nil {
		return err
	}
	if o.network == NetworkCustom {
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
		inputFeeRecipient,
		inputGenerationPath,
	); err != nil {
		return err
	}
	log.Error("unimplemented generation")
	return postGenerate(p, o, actions)
}

func setupJWT(p prompts.Prompt, o *CliRefactorOptions, skip bool) error {
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
		log.Error("unimplemented JWT generation")
		// TODO: generate JWT
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

func postGenerate(p prompts.Prompt, o *CliRefactorOptions, a actions.SedgeActions) error {
	if o.withValidator || o.nodeType == NodeTypeValidator {
		if err := generateKeystore(p, o, a); err != nil {
			return err
		}
	}
	pendingDependencies := utils.CheckDependencies([]string{"docker", "invalid-dep"})
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
	a.SetupContainers(actions.SetupContainersOptions{
		GenerationPath: o.generationPath,
		Services:       services,
	})
	a.RunContainers(actions.RunContainersOptions{
		GenerationPath: o.generationPath,
		Services:       services,
	})
	// TODO: Final tips
	log.Error("show final tips")
	return nil
}

func generateKeystore(p prompts.Prompt, o *CliRefactorOptions, a actions.SedgeActions) error {
	if err := selectKeystoreSource(p, o); err != nil {
		return err
	}
	switch o.keystoreSourceType {
	case SourceTypeSkip:
		return nil
	case SourceTypeCreate:
		// TODO generate keystore
		log.Error("unimplemented keystore generation")
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
			// TODO open file and set passphrase
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
			ValidatorClient: o.validatorClient,
			Network:         o.network,
			GenerationPath:  o.generationPath,
			From:            o.slashingProtectionFrom,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type promptAction func(prompts.Prompt, *CliRefactorOptions) error

func runPromptActions(p prompts.Prompt, o *CliRefactorOptions, actions ...promptAction) error {
	for _, action := range actions {
		if err := action(p, o); err != nil {
			return err
		}
	}
	return nil
}

func selectNetwork(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.network, err = p.Select("Select network", NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom)
	return
}

func selectNodeType(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.nodeType, err = p.Select("Select node type", NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator)
	return
}

func selectExecutionClient(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	// TODO: support randomize
	o.executionClient, err = p.Select("Select execution client", ExecutionNethermind, ExecutionGeth, ExecutionBesu, ExecutionErigon)
	return
}

func selectConsensusClient(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	// TODO: support randomize
	o.consensusClient, err = p.Select("Select consensus client", ConsensusLighthouse, ConsensusLodestar, ConsensusPrysm, ConsensusTeku)
	return
}

func selectValidatorClient(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	// TODO: support randomize
	o.validatorClient, err = p.Select("Select validator client", ValidatorLighthouse, ValidatorLodestar, ValidatorPrysm, ValidatorTeku)
	return
}

func selectJWTSource(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.jwtSourceType, err = p.Select("Select JWT source", SourceTypeCreate, SourceTypeExisting)
	return
}

func selectJWTSourceOrSkip(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.jwtSourceType, err = p.Select("Select JWT source", SourceTypeCreate, SourceTypeExisting, SourceTypeSkip)
	return
}

func selectKeystoreSource(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.keystoreSourceType, err = p.Select("Select keystore source", SourceTypeCreate, SourceTypeExisting, SourceTypeSkip)
	return
}

func selectKeystoreMnemonicSource(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.keystoreMnemonicSource, err = p.Select("Select mnemonic source", SourceTypeCreate, SourceTypeExisting)
	return
}

func selectKeystorePassphraseSource(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.keystorePassphraseSource, err = p.Select("Select passphrase source", SourceTypeRandom, SourceTypeExisting, SourceTypeCreate)
	return
}

func confirmWithValidator(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.withValidator, err = p.Confirm("Do you want to set up a validator")
	return
}

func confirmExposeAllPorts(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.exposeAllPorts, err = p.Confirm("Do you want to expose all ports?")
	return
}

func confirmImportSlashingProtection(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.importSlashingProtection, err = p.Confirm("Do you want to import slashing protection data?")
	return
}

func confirmInstallDependencies(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.installDependencies, err = p.Confirm("Install dependencies?")
	return
}

func inputCustomNetworkConfig(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.customNetworkConfig, err = p.InputFilePath("Custom Network Config", true)
	return
}

func inputCustomChainSpec(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.customChainSpec, err = p.InputFilePath("Custom ChainSpec", true)
	return
}

func inputCustomGenesis(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.customGenesis, err = p.InputFilePath("Custom Genesis", true)
	return
}

func inputCustomTTD(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.customTTD, err = p.Input("Custom TTD", false)
	return
}

func inputCustomDeployBlock(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.customDeployBlock, err = p.Input("Custom deploy block", false)
	return
}

func inputExecutionBootNodes(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.executionBootNodes, err = p.Input("Execution boot nodes", false)
	return
}

func inputConsensusBootNodes(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.consensusBootNodes, err = p.Input("Consensus boot nodes", false)
	return
}

func inputMevBoostImage(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	// TODO add default value
	o.mevBoostImage, err = p.Input("Mev-Boost image", false)
	return
}

func inputMevBoostEndpoint(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.mevBoostEndpoint, err = p.Input("Mev-Boost endpoint", false)
	return
}

func inputRelayURL(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	// TODO add default value
	o.relayURL, err = p.Input("Relay URL", false)
	return
}

func inputGraffiti(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.graffiti, err = p.Input("Graffiti", false)
	return
}

func inputCheckpointSyncURL(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	// TODO: add default value
	o.ckptSyncURL, err = p.Input("Checkpoint sync URL", false)
	return
}

func inputFeeRecipient(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	// TODO: Is necessary to add a default value?
	o.feeRecipient, err = p.Input("Fee recipient", true)
	return
}

func inputGenerationPath(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	// TODO: add default
	o.generationPath, err = p.Input("Generation path", true)
	return
}

func inputJWTPath(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.jwtPath, err = p.Input("JWT path", true)
	return
}

func inputKeystoreMnemonicPath(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.keystoreMnemonic, err = p.Input("Mnemonic path", true)
	return
}

func inputKeystorePassphrasePath(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.keystorePassphrasePath, err = p.Input("Passphrase path", true)
	return
}

func inputKeystorePassphrase(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.keystorePassphrase, err = p.InputHide("Passphrase")
	return
}

func inputWithdrawalAddress(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.withdrawalAddress, err = p.Input("Withdrawal address", false)
	return
}

func inputNumberOfValidators(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.numberOfValidators, err = p.InputNumber("Number of validators")
	return
}

func inputNumberOfExistingValidators(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.existingValidators, err = p.InputNumber("Existing validators")
	return
}

func inputKeystorePath(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.keystorePath, err = p.Input("Keystore path", true)
	return
}

func inputImportSlashingProtectionFrom(p prompts.Prompt, o *CliRefactorOptions) (err error) {
	o.slashingProtectionFrom, err = p.InputFilePath("Interchange slashing protection file", true)
	return
}
