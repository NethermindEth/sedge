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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	eth2 "github.com/protolambda/zrnt/eth2/configs"

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/mevboostrelaylist"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/ui"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/keystores"
	"github.com/NethermindEth/sedge/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	NetworkMainnet = "mainnet"
	NetworkSepolia = "sepolia"
	NetworkGnosis  = "gnosis"
	NetworkChiado  = "chiado"
	NetworkHolesky = "holesky"
	NetworkCustom  = "custom"

	NodeTypeFullNode  = "full-node"
	NodeTypeExecution = "execution"
	NodeTypeConsensus = "consensus"
	NodeTypeValidator = "validator"

	EthereumNode = "ethereum-node"
	LidoNode     = "lido-node"

	Randomize = "randomize"

	SourceTypeExisting = "existing"
	SourceTypeCreate   = "create"
	SourceTypeSkip     = "skip"
	SourceTypeRandom   = "random"
)

var ErrCancelled = errors.New("cancelled by the user")

type CliCmdOptions struct {
	genData                  generate.GenData
	generationPath           string
	nodeType                 string
	withValidator            bool
	withMevBoost             bool
	importSlashingProtection bool
	slashingProtectionFrom   string
	jwtSourceType            string
	keystoreSourceType       string
	keystorePath             string
	keystoreMnemonicSource   string
	keystoreMnemonic         string
	keystoreMnemonicPath     string
	keystorePassphraseSource string
	keystorePassphrasePath   string
	keystorePassphrase       string
	withdrawalAddress        string
	nodeSetup                string
	numberOfValidators       int64
	existingValidators       int64
	installDependencies      bool
}

func CliCmd(p ui.Prompter, actions actions.SedgeActions, depsMgr dependencies.DependenciesManager) *cobra.Command {
	o := new(CliCmdOptions)
	cmd := &cobra.Command{
		Use:   "cli",
		Short: "Generate a node setup interactively",
		Long: `This command will guide you through the process of setting up one of these node types:

- Full Node (execution + consensus + validator)
- Full Node without Validator (execution + consensus)
- Execution Node
- Consensus Node
- Validator Node
- Lido CSM Node

Follow the prompts to select the options you want for your node. At the end of the process, you will
be asked to run the generated setup or not. If you chose to run the setup, it will be executed for you
using docker compose command behind the scenes.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runPromptActions(p, o,
				selectNodeSetup,
				selectNetwork,
				selectNodeType,
				inputGenerationPath,
				inputContainerTag,
			); err != nil {
				return err
			}
			switch o.nodeType {
			case NodeTypeFullNode:
				return setupFullNode(p, o, actions, depsMgr)
			case NodeTypeExecution:
				return setupExecutionNode(p, o, actions, depsMgr)
			case NodeTypeConsensus:
				return setupConsensusNode(p, o, actions, depsMgr)
			case NodeTypeValidator:
				return setupValidatorNode(p, o, actions, depsMgr)
			}
			return nil
		},
	}
	return cmd
}

func setupFullNode(p ui.Prompter, o *CliCmdOptions, a actions.SedgeActions, depsManager dependencies.DependenciesManager) (err error) {
	o.genData.Services = []string{"execution", "consensus"}
	if err := confirmWithValidator(p, o); err != nil {
		return err
	}
	if o.genData.Network == NetworkCustom {
		if err := runPromptActions(p, o,
			inputCustomNetworkConfig,
			inputCustomChainSpec,
			inputCustomGenesis,
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
			if err := confirmEnableMEVBoost(p, o); err != nil {
				return err
			}
			if o.withMevBoost {
				o.genData.Mev = o.withMevBoost
				o.genData.Services = append(o.genData.Services, "mev-boost")
			}
			if o.withMevBoost {
				if err := runPromptActions(p, o,
					inputMevImage,
					inputRelayURL,
				); err != nil {
					return err
				}
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
			inputFeeRecipientNoValidator,
		); err != nil {
			return err
		}
	}
	if err := confirmExposeAllPorts(p, o); err != nil {
		return err
	}
	if err := setupJWT(p, o, false); err != nil {
		return err
	}
	// Call generate action
	o.genData, err = a.Generate(actions.GenerateOptions{
		GenerationData: o.genData,
		GenerationPath: o.generationPath,
	})
	if err != nil {
		return err
	}
	return postGenerate(p, o, a, depsManager)
}

func setupExecutionNode(p ui.Prompter, o *CliCmdOptions, a actions.SedgeActions, depsManager dependencies.DependenciesManager) (err error) {
	o.genData.Services = []string{"execution"}
	if err := selectExecutionClient(p, o); err != nil {
		return err
	}
	if o.genData.Network == NetworkCustom {
		if err := runPromptActions(p, o,
			inputCustomChainSpec,
			inputExecutionBootNodes,
		); err != nil {
			return err
		}
	}
	if err := confirmExposeAllPorts(p, o); err != nil {
		return err
	}
	if err := setupJWT(p, o, true); err != nil {
		return err
	}
	o.genData, err = a.Generate(actions.GenerateOptions{
		GenerationData: o.genData,
		GenerationPath: o.generationPath,
	})
	if err != nil {
		return err
	}
	return postGenerate(p, o, a, depsManager)
}

func setupConsensusNode(p ui.Prompter, o *CliCmdOptions, a actions.SedgeActions, depsManager dependencies.DependenciesManager) (err error) {
	o.genData.Services = []string{"consensus"}
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
		inputExecutionAPIUrl,
		inputExecutionAuthUrl,
		inputFeeRecipientNoValidator,
		confirmExposeAllPorts,
	); err != nil {
		return err
	}
	if err := setupJWT(p, o, true); err != nil {
		return err
	}
	o.genData, err = a.Generate(actions.GenerateOptions{
		GenerationData: o.genData,
		GenerationPath: o.generationPath,
	})
	if err != nil {
		return err
	}
	return postGenerate(p, o, a, depsManager)
}

func setupValidatorNode(p ui.Prompter, o *CliCmdOptions, a actions.SedgeActions, depsManager dependencies.DependenciesManager) (err error) {
	o.genData.Services = []string{"validator"}
	if err := selectValidatorClient(p, o); err != nil {
		return err
	}
	if o.genData.Network == NetworkCustom {
		if err := runPromptActions(p, o,
			inputCustomNetworkConfig,
			inputCustomGenesis,
			inputCustomDeployBlock,
		); err != nil {
			return err
		}
	}
	if err := runPromptActions(p, o,
		inputConsensusAPIUrl,
		inputGraffiti,
		inputValidatorGracePeriod,
		inputFeeRecipient,
	); err != nil {
		return err
	}
	if configs.SupportsMEVBoost(o.genData.Network) {
		if err := confirmEnableMEVBoost(p, o); err != nil {
			return err
		}
		o.genData.MevBoostOnValidator = o.withMevBoost
	}
	o.genData, err = a.Generate(actions.GenerateOptions{
		GenerationData: o.genData,
		GenerationPath: o.generationPath,
	})
	if err != nil {
		return err
	}
	return postGenerate(p, o, a, depsManager)
}

func setupJWT(p ui.Prompter, o *CliCmdOptions, skip bool) error {
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
		jwtPath, err := handleJWTSecret(o.generationPath)
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

func postGenerate(p ui.Prompter, o *CliCmdOptions, a actions.SedgeActions, depsMgr dependencies.DependenciesManager) error {
	if o.withValidator || o.nodeType == NodeTypeValidator {
		if err := generateKeystore(p, o, a, depsMgr); err != nil {
			return err
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
		if err := checkCLIDependencies(p, o, a, depsMgr); err != nil {
			return err
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
		if o.withValidator {
			log.Info(configs.HappyStakingRun)
		} else {
			log.Infof(configs.HappySedgingRun, o.generationPath)
		}
	} else {
		log.Infof(configs.HappySedgingNoRun, o.generationPath)
	}
	return nil
}

func generateKeystore(p ui.Prompter, o *CliCmdOptions, a actions.SedgeActions, depsMgr dependencies.DependenciesManager) error {
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
			if mnemonic, err := readFileContent(o.keystoreMnemonicPath); err != nil {
				return err
			} else {
				o.keystoreMnemonic = mnemonic
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
		o.keystorePath = filepath.Join(o.generationPath, "keystore")
		// Check if file exists
		if f, err := os.Stat(o.keystorePath); err == nil {
			if f.IsDir() {
				overwrite, err := p.Confirm(fmt.Sprintf("%s already exists. Do you want to overwrite it?", o.keystorePath), false)
				if err != nil {
					return err
				}
				if overwrite {
					if err := os.RemoveAll(o.keystorePath); err != nil {
						return err
					}
				} else {
					return fmt.Errorf("%s already exists", o.keystorePath)
				}
			} else {
				return fmt.Errorf("%s is not a directory", o.keystorePath)
			}
		}

		// TODO: Create an Action for keystore generation
		log.Info("Generating keystores...")
		data := keystores.ValidatorKeysGenData{
			Mnemonic:    o.keystoreMnemonic,
			Passphrase:  o.keystorePassphrase,
			OutputPath:  o.keystorePath,
			MinIndex:    uint64(o.existingValidators),
			MaxIndex:    uint64(o.existingValidators) + uint64(o.numberOfValidators),
			NetworkName: o.genData.Network,
			ForkVersion: configs.NetworksConfigs()[o.genData.Network].GenesisForkVersion,
			// Constants
			UseUniquePassphrase: true,
			Insecure:            false,
			AmountGwei:          uint64(eth2.Mainnet.MAX_EFFECTIVE_BALANCE),
			AsJsonList:          true,
		}
		if err := keystores.CreateKeystores(data); err != nil {
			return err
		}
		log.Info(configs.KeystoresGenerated)
		log.Info(configs.GeneratingDepositData)
		if err := keystores.CreateDepositData(data); err != nil {
			log.Fatal(err)
		}
		log.Info(configs.DepositDataGenerated)
	case SourceTypeExisting:
		if err := inputKeystorePath(p, o); err != nil {
			return err
		}
		validationErrors := keystores.ValidateKeystoreDir(o.keystorePath)
		if len(validationErrors) > 0 {
			log.Warnf("Keystore folder %s has %d validation errors. Check the following:", o.keystorePath, len(validationErrors))
			for index, e := range validationErrors {
				log.Warnf("%d. %s", index+1, e.Error())
			}
			cont, err := p.Confirm("Do you want to continue regardless the keystore folder validation warnings?", false)
			if err != nil {
				return err
			} else if !cont {
				return ErrCancelled
			}
		} else {
			log.Infof("Keystore folder %s is valid", o.keystorePath)
		}
	}
	if err := checkCLIDependencies(p, o, a, depsMgr); err != nil {
		return err
	}
	log.Info("Importing validator keys into the validator client...")
	err := a.SetupContainers(actions.SetupContainersOptions{
		GenerationPath: o.generationPath,
		Services:       []string{validator},
	})
	if err != nil {
		return err
	}
	err = a.ImportValidatorKeys(actions.ImportValidatorKeysOptions{
		ValidatorClient: o.genData.ValidatorClient.Name,
		Network:         o.genData.Network,
		GenerationPath:  o.generationPath,
		From:            o.keystorePath,
		ContainerTag:    o.genData.ContainerTag,
		CustomConfig: actions.ImportValidatorKeysCustomOptions{
			NetworkConfigPath: o.genData.CustomNetworkConfigPath,
			GenesisPath:       o.genData.CustomGenesisPath,
			DeployBlockPath:   o.genData.CustomDeployBlockPath,
		},
	})
	if err != nil {
		return err
	}
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
			ContainerTag:    o.genData.ContainerTag,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func checkCLIDependencies(p ui.Prompter, o *CliCmdOptions, a actions.SedgeActions, depsMgr dependencies.DependenciesManager) error {
	_, pendingDependencies := depsMgr.Check([]string{dependencies.Docker})
	if len(pendingDependencies) > 0 {
		supported, unsupported, err := depsMgr.Supported(pendingDependencies)
		if err != nil {
			if errors.Is(err, dependencies.ErrUnsupportedInstallForOS) {
				log.Warnf(err.Error())
				return nil
			} else {
				return err
			}
		}
		if len(unsupported) > 0 {
			log.Warnf("unsupported install dependencies %s", strings.Join(unsupported, " "))
			return nil
		}
		// FIXME: There is an issue with the cli command and sudo permissions. Sedge deps install don't have this issue. This should be investigated and solved before uncommenting the commented code below.
		// if err := confirmInstallDependencies(p, o); err != nil {
		// 	return err
		// }
		o.installDependencies = false
		if !o.installDependencies {
			for _, s := range supported {
				if err := depsMgr.ShowInstructions(s); err != nil {
					return err
				}
			}
			return fmt.Errorf("%w: %s. To install dependencies if supported run: 'sedge deps install'", ErrMissingDependencies, strings.Join(pendingDependencies, ", "))
		}
		// for _, s := range supported {
		// 	if err := depsMgr.Install(s); err != nil {
		// 		return err
		// 	}
		// }
	}
	if err := depsMgr.DockerEngineIsOn(); err != nil {
		return err
	}
	return depsMgr.DockerComposeIsInstalled()
}

type promptAction func(ui.Prompter, *CliCmdOptions) error

func runPromptActions(p ui.Prompter, o *CliCmdOptions, actions ...promptAction) error {
	for _, action := range actions {
		if err := action(p, o); err != nil {
			return err
		}
	}
	return nil
}

func selectNodeSetup(p ui.Prompter, o *CliCmdOptions) (err error) {
	options := []string{EthereumNode, LidoNode}
	index, err := p.Select("Select node setup", "", options)
	if err != nil {
		return err
	}
	o.nodeSetup = options[index]
	return
}

func selectNetwork(p ui.Prompter, o *CliCmdOptions) error {
	var options []string
	if o.nodeSetup == LidoNode {
		options = []string{NetworkMainnet, NetworkHolesky, NetworkSepolia}
	} else {
		options = []string{NetworkMainnet, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}
	}
	index, err := p.Select("Select network", "", options)
	if err != nil {
		return err
	}
	o.genData.Network = options[index]
	return nil
}

func selectNodeType(p ui.Prompter, o *CliCmdOptions) error {
	options := []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}
	index, err := p.Select("Select node type", "", options)
	if err != nil {
		return err
	}
	o.nodeType = options[index]
	return nil
}

func selectExecutionClient(p ui.Prompter, o *CliCmdOptions) (err error) {
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
		log.Info("Random execution client selected: ", selectedExecutionClient)
	}
	o.genData.ExecutionClient = &clients.Client{
		Name: selectedExecutionClient,
		Type: "execution",
	}
	o.genData.ExecutionClient.SetImageOrDefault("")
	return nil
}

func selectConsensusClient(p ui.Prompter, o *CliCmdOptions) (err error) {
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
		log.Info("Random consensus client selected: ", selectedConsensusClient)
	}
	o.genData.ConsensusClient = &clients.Client{
		Name: selectedConsensusClient,
		Type: "consensus",
	}
	o.genData.ConsensusClient.SetImageOrDefault("")
	return nil
}

func selectValidatorClient(p ui.Prompter, o *CliCmdOptions) (err error) {
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
		log.Info("Random validator client selected: ", selectedValidatorClient)
	}
	o.genData.ValidatorClient = &clients.Client{
		Name: selectedValidatorClient,
		Type: "validator",
	}
	o.genData.ValidatorClient.SetImageOrDefault("")
	return nil
}

func selectJWTSource(p ui.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeCreate, SourceTypeExisting}
	index, err := p.Select("Select JWT source", "", options)
	if err != nil {
		return err
	}
	o.jwtSourceType = options[index]
	return nil
}

func selectJWTSourceOrSkip(p ui.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}
	index, err := p.Select("Select JWT source", "", options)
	if err != nil {
		return err
	}
	o.jwtSourceType = options[index]
	return nil
}

func selectKeystoreSource(p ui.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}
	index, err := p.Select("Select keystore source", "", options)
	if err != nil {
		return err
	}
	o.keystoreSourceType = options[index]
	return nil
}

func selectKeystoreMnemonicSource(p ui.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeCreate, SourceTypeExisting}
	index, err := p.Select("Select mnemonic source", "", options)
	if err != nil {
		return err
	}
	o.keystoreMnemonicSource = options[index]
	return nil
}

func selectKeystorePassphraseSource(p ui.Prompter, o *CliCmdOptions) error {
	options := []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}
	index, err := p.Select("Select passphrase source", "", options)
	if err != nil {
		return err
	}
	o.keystorePassphraseSource = options[index]
	return nil
}

func confirmWithValidator(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.withValidator, err = p.Confirm("Do you want to set up a validator?", true)
	return
}

func confirmExposeAllPorts(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.MapAllPorts, err = p.Confirm("Do you want to expose all ports?", false)
	return
}

func confirmImportSlashingProtection(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.importSlashingProtection, err = p.Confirm("Do you want to import slashing protection data?", false)
	return
}

func confirmInstallDependencies(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.installDependencies, err = p.Confirm("Install dependencies?", false)
	return
}

func confirmEnableMEVBoost(p ui.Prompter, o *CliCmdOptions) (err error) {
	if o.nodeSetup == LidoNode {
		_, ok := mevboostrelaylist.DeployedContractAddresses[o.genData.Network]
		if !ok {
			return
		}
	}
	o.withMevBoost, err = p.Confirm("Enable MEV Boost?", true)
	return
}

func inputCustomNetworkConfig(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CustomNetworkConfigPath, err = p.InputFilePath("Custom network config file path", "", true, ".yml", ".yaml")
	if err != nil {
		return err
	}
	return absPathInPlace(&o.genData.CustomNetworkConfigPath)
}

func inputCustomChainSpec(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CustomChainSpecPath, err = p.InputFilePath("File path or url to use as custom network chainSpec for execution client", "", true, ".json")
	if err != nil {
		return err
	}
	return absPathInPlace(&o.genData.CustomChainSpecPath)
}

func inputCustomGenesis(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CustomGenesisPath, err = p.InputFilePath("File path or URL to use as custom network genesis for consensus client", "", true, ".ssz")
	if err != nil {
		return err
	}
	return absPathInPlace(&o.genData.CustomGenesisPath)
}

func inputCustomDeployBlock(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CustomDeployBlock, err = p.Input("Custom deploy block", "0", false, ui.DigitsStringValidator)
	return
}

func inputExecutionBootNodes(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.ECBootnodes, err = p.InputList("Execution boot nodes", []string{}, utils.ENodesValidator)
	return
}

func inputConsensusBootNodes(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CCBootnodes, err = p.InputList("Consensus boot nodes", []string{}, utils.ENRValidator)
	return
}

func inputMevImage(p ui.Prompter, o *CliCmdOptions) (err error) {
	// Default value is set in the template
	o.genData.MevImage, err = p.Input("Mev-Boost image", "flashbots/mev-boost:latest", false, nil)
	return
}

func inputMevBoostEndpoint(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.MevBoostEndpoint, err = p.InputURL("Mev-Boost endpoint", "", false)
	return
}

func inputRelayURL(p ui.Prompter, o *CliCmdOptions) (err error) {
	if o.nodeSetup == LidoNode {
		relayURLs, err := mevboostrelaylist.GetRelaysURI(o.genData.Network)
		if err != nil {
			return err
		}
		o.genData.RelayURLs = relayURLs
		return nil
	}
	var defaultValue []string = configs.NetworksConfigs()[o.genData.Network].RelayURLs
	relayURLs, err := p.InputList("Insert relay URLs if you don't want to use the default values listed below", defaultValue, func(list []string) error {
		badUri, ok := utils.UriValidator(list)
		if !ok {
			return fmt.Errorf(configs.InvalidUrlFlagError, "relay", badUri)
		}
		return nil
	})
	o.genData.RelayURLs = relayURLs
	return
}

func inputGraffiti(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.Graffiti, err = p.Input("Graffiti to be used by the validator (press enter to skip it)", "", false, ui.GraffitiValidator)
	return
}

func inputCheckpointSyncURL(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.CheckpointSyncUrl, err = p.InputURL("Checkpoint sync URL", configs.NetworksConfigs()[o.genData.Network].CheckpointSyncURL, false)
	return
}

func inputFeeRecipient(p ui.Prompter, o *CliCmdOptions) (err error) {
	if o.nodeSetup == LidoNode {
		feeRecipient := contracts.FeeRecipient[o.genData.Network]
		o.genData.FeeRecipient = feeRecipient.FeeRecipientAddress
		return
	}
	o.genData.FeeRecipient, err = p.EthAddress("Please enter the Fee Recipient address", "", true)
	return
}

func inputFeeRecipientNoValidator(p ui.Prompter, o *CliCmdOptions) (err error) {
	if o.nodeSetup == LidoNode {
		feeRecipient := contracts.FeeRecipient[o.genData.Network]
		o.genData.FeeRecipient = feeRecipient.FeeRecipientAddress
		return
	}
	o.genData.FeeRecipient, err = p.EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false)
	return
}

func inputValidatorGracePeriod(p ui.Prompter, o *CliCmdOptions) (err error) {
	epochs, err := p.InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", 1)
	if err != nil {
		return err
	}
	o.genData.VLStartGracePeriod = uint((configs.NetworkEpochTime(o.genData.Network) * time.Duration(epochs)).Seconds())
	return nil
}

func inputGenerationPath(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.generationPath, err = p.Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil)
	if err != nil {
		return err
	}
	if o.generationPath == "" {
		o.generationPath = configs.DefaultAbsSedgeDataPath
	}
	return absPathInPlace(&o.generationPath)
}

func inputJWTPath(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.JWTSecretPath, err = p.InputFilePath("JWT path", "", true)
	if err != nil {
		return err
	}
	return absPathInPlace(&o.genData.JWTSecretPath)
}

func inputKeystoreMnemonicPath(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.keystoreMnemonicPath, err = p.InputFilePath("Mnemonic path", "", true)
	if err != nil {
		return err
	}
	return absPathInPlace(&o.keystoreMnemonicPath)
}

func inputKeystorePassphrasePath(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.keystorePassphrasePath, err = p.InputFilePath("Passphrase path", "", true)
	if err != nil {
		return err
	}
	return absPathInPlace(&o.keystorePassphrasePath)
}

func inputKeystorePassphrase(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.keystorePassphrase, err = p.InputSecret("Enter keystore passphrase (min 8 characters):")
	return
}

func inputWithdrawalAddress(p ui.Prompter, o *CliCmdOptions) (err error) {
	if o.nodeSetup == LidoNode {
		_, ok := contracts.WithdrawalAddress[network]
		if ok {
			o.withdrawalAddress = contracts.WithdrawalAddress[network].WithdrawalAddress
		}
		return
	}
	o.withdrawalAddress, err = p.Input("Withdrawal address", "", false, func(s string) error { return ui.EthAddressValidator(s, true) })
	return
}

func inputNumberOfValidators(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.numberOfValidators, err = p.InputInt64("Number of validators", 1)
	return
}

func inputNumberOfExistingValidators(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.existingValidators, err = p.InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", 0)
	return
}

func inputKeystorePath(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.keystorePath, err = p.InputDirPath("Keystore path", filepath.Join(o.generationPath, "keystore"), true)
	if err != nil {
		return err
	}
	return absPathInPlace(&o.keystorePath)
}

func inputImportSlashingProtectionFrom(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.slashingProtectionFrom, err = p.InputFilePath("Interchange slashing protection file path", "", true, ".json")
	if err != nil {
		return err
	}
	return absPathInPlace(&o.slashingProtectionFrom)
}

func inputExecutionAPIUrl(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.ExecutionApiUrl, err = p.InputURL("Execution API URL", "", true)
	return
}

func inputExecutionAuthUrl(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.ExecutionAuthUrl, err = p.InputURL("Execution Auth API URL", "", true)
	return
}

func inputConsensusAPIUrl(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.ConsensusApiUrl, err = p.InputURL("Consensus API URL", "", true)
	return
}

func inputContainerTag(p ui.Prompter, o *CliCmdOptions) (err error) {
	o.genData.ContainerTag, err = p.Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil)
	return
}

func absPathInPlace(path *string) error {
	absPath, err := filepath.Abs(*path)
	if err != nil {
		return err
	}
	*path = absPath
	return nil
}
