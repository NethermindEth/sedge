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
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/lido/contracts/mevboostrelaylist"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	sedgeOpts "github.com/NethermindEth/sedge/internal/pkg/options"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/NethermindEth/sedge/internal/utils"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/NethermindEth/sedge/test"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
)

func TestCli(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	mainnetMevboostRelayListUris, _ := mevboostrelaylist.RelaysURI("mainnet")
	hoodiMevboostRelayListUris, _ := mevboostrelaylist.RelaysURI("hoodi")

	ETHClients := map[string][]string{
		"execution": clients.AllClients["execution"],
		"consensus": clients.AllClients["consensus"],
		"validator": clients.AllClients["validator"],
	}
	ETHClients["execution"] = append(ETHClients["execution"], "randomize")
	ETHClients["consensus"] = append(ETHClients["consensus"], "randomize")
	ETHClients["validator"] = append(ETHClients["validator"], "randomize")

	GnosisClients := map[string][]string{
		"execution": {"nethermind"},
		"consensus": utils.Filter(clients.AllClients["consensus"], func(c string) bool { return c != "prysm" }),
		"validator": utils.Filter(clients.AllClients["validator"], func(c string) bool { return c != "prysm" }),
	}
	GnosisClients["execution"] = append(GnosisClients["execution"], "randomize")
	GnosisClients["consensus"] = append(GnosisClients["consensus"], "randomize")
	GnosisClients["validator"] = append(GnosisClients["validator"], "randomize")

	tests := []struct {
		name  string
		setup func(*testing.T, *sedge_mocks.MockSedgeActions, *sedge_mocks.MockPrompter, *sedge_mocks.MockDependenciesManager)
	}{
		{
			name: "full node with validator mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"execution", "consensus", "validator", "mev-boost"},
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					ConsensusClient: &clients.Client{
						Name:  "prysm",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Prysm.String(),
					},
					ValidatorClient: &clients.Client{
						Name:  "prysm",
						Type:  "validator",
						Image: configs.ClientImages.Validator.Prysm.String(),
					},
					Network:            "mainnet",
					CheckpointSyncUrl:  "http://checkpoint.sync",
					FeeRecipient:       "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
					MapAllPorts:        true,
					Graffiti:           "test graffiti",
					VLStartGracePeriod: 840,
					Mev:                true,
					MevImage:           "flashbots/mev-boost:latest",
					RelayURLs:          configs.NetworksConfigs()[NetworkMainnet].RelayURLs,
					ContainerTag:       "tag",
					JWTSecretPath:      filepath.Join(generationPath, "jwtsecret"),
				}
				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})
				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(0, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", true).Return(true, nil),
					prompter.EXPECT().Confirm("Enable MEV Boost?", true).Return(true, nil),
					prompter.EXPECT().Input("Mev-Boost image", "flashbots/mev-boost:latest", false, nil).Return("flashbots/mev-boost:latest", nil),
					prompter.EXPECT().InputList("Insert relay URLs if you don't want to use the default values listed below", configs.NetworksConfigs()[NetworkMainnet].RelayURLs, gomock.AssignableToTypeOf(func([]string) error { return nil })).Return(configs.NetworksConfigs()[NetworkMainnet].RelayURLs, nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(1, nil),
					prompter.EXPECT().Select("Select validator client", "", ETHClients["validator"]).Return(1, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Graffiti to be used by the validator (press enter to skip it)", "", false, gomock.AssignableToTypeOf(ui.GraffitiValidator)).Return("test graffiti", nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address", "", true).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().Input("Withdrawal address", "", false, gomock.AssignableToTypeOf(func(s string) error { return ui.EthAddressValidator(s, true) })).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					depsMgr.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil),
					depsMgr.EXPECT().DockerEngineIsOn().Return(nil),
					depsMgr.EXPECT().DockerComposeIsInstalled().Return(nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         NetworkMainnet,
						GenerationPath:  generationPath,
						From:            filepath.Join(generationPath, "keystore"),
						ContainerTag:    "tag",
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "full node without validator mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"execution", "consensus"},
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					ConsensusClient: &clients.Client{
						Name:  "prysm",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Prysm.String(),
					},
					Network:           "mainnet",
					CheckpointSyncUrl: "http://checkpoint.sync",
					FeeRecipient:      "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
					MapAllPorts:       true,
					ContainerTag:      "tag",
					JWTSecretPath:     filepath.Join(generationPath, "jwtsecret"),
				}
				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(0, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", true).Return(false, nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(1, nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "execution node",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"execution"},
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					Network:      "mainnet",
					MapAllPorts:  true,
					ContainerTag: "tag",
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(1, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(2, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(true, nil),
					depsMgr.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil),
					depsMgr.EXPECT().DockerEngineIsOn().Return(nil),
					depsMgr.EXPECT().DockerComposeIsInstalled().Return(nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution"},
					}),
					sedgeActions.EXPECT().RunContainers(actions.RunContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution"},
					}),
				)
			},
		},
		{
			name: "consensus node",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"consensus"},
					ConsensusClient: &clients.Client{
						Name:  "lodestar",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Lodestar.String(),
					},
					Network:           NetworkMainnet,
					CheckpointSyncUrl: "http://checkpoint.sync",
					FeeRecipient:      "0x2d07a21ebadde0c13e8b91022a7e5732eb6bf5d5",
					MapAllPorts:       true,
					ExecutionApiUrl:   "http://execution:5051",
					ExecutionAuthUrl:  "http://execution:5051",
					MevBoostEndpoint:  "http://mev-boost:3030",
					ContainerTag:      "tag",
					JWTSecretPath:     filepath.Join(generationPath, "jwtsecret"),
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(2, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(3, nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().InputURL("Mev-Boost endpoint", "", false).Return("http://mev-boost:3030", nil),
					prompter.EXPECT().InputURL("Execution API URL", "", true).Return("http://execution:5051", nil),
					prompter.EXPECT().InputURL("Execution Auth API URL", "", true).Return("http://execution:5051", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("0x2d07a21ebadde0c13e8b91022a7e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "validator mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()

				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})
				genData := generate.GenData{
					Services: []string{"validator"},
					ValidatorClient: &clients.Client{
						Name:  "prysm",
						Type:  "validator",
						Image: configs.ClientImages.Validator.Prysm.String(),
					},
					Network:             "mainnet",
					FeeRecipient:        "0x2d07a31ebadce0a13e8a91022a5e5732eb6bf5d5",
					Graffiti:            "test graffiti",
					VLStartGracePeriod:  840,
					MevBoostOnValidator: true,
					ConsensusApiUrl:     "http://localhost:5051",
					ContainerTag:        "tag",
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(3, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select validator client", "", ETHClients["validator"]).Return(1, nil),
					prompter.EXPECT().InputURL("Consensus API URL", "", true).Return("http://localhost:5051", nil),
					prompter.EXPECT().Input("Graffiti to be used by the validator (press enter to skip it)", "", false, gomock.AssignableToTypeOf(ui.GraffitiValidator)).Return("test graffiti", nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address", "", true).Return("0x2d07a31ebadce0a13e8a91022a5e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Enable MEV Boost?", true).Return(true, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().Input("Withdrawal address", "", false, gomock.AssignableToTypeOf(func(s string) error { return ui.EthAddressValidator(s, true) })).Return("0x2d07a21ebadde0c13e6b91022a7e5732eb6bf5d5", nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					depsMgr.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil),
					depsMgr.EXPECT().DockerEngineIsOn().Return(nil),
					depsMgr.EXPECT().DockerComposeIsInstalled().Return(nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         NetworkMainnet,
						From:            filepath.Join(generationPath, "keystore"),
						GenerationPath:  generationPath,
						ContainerTag:    "tag",
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "consensus node",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"consensus"},
					ConsensusClient: &clients.Client{
						Name:  "lodestar",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Lodestar.String(),
					},
					Network:           NetworkMainnet,
					CheckpointSyncUrl: "http://checkpoint.sync",
					FeeRecipient:      "0x2d07a21ebadde0c13e8b91022a7e5732eb6bf5d5",
					MapAllPorts:       true,
					ExecutionApiUrl:   "http://execution:5051",
					ExecutionAuthUrl:  "http://execution:5051",
					MevBoostEndpoint:  "http://mev-boost:3030",
					ContainerTag:      "tag",
					JWTSecretPath:     filepath.Join(generationPath, "jwtsecret"),
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(2, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(3, nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().InputURL("Mev-Boost endpoint", "", false).Return("http://mev-boost:3030", nil),
					prompter.EXPECT().InputURL("Execution API URL", "", true).Return("http://execution:5051", nil),
					prompter.EXPECT().InputURL("Execution Auth API URL", "", true).Return("http://execution:5051", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("0x2d07a21ebadde0c13e8b91022a7e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "full node with Lido, mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"execution", "consensus", "validator", "mev-boost"},
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					ConsensusClient: &clients.Client{
						Name:  "prysm",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Prysm.String(),
					},
					ValidatorClient: &clients.Client{
						Name:  "prysm",
						Type:  "validator",
						Image: configs.ClientImages.Validator.Prysm.String(),
					},
					Network:            "mainnet",
					CheckpointSyncUrl:  "http://checkpoint.sync",
					FeeRecipient:       "0x388C818CA8B9251b393131C08a736A67ccB19297",
					MapAllPorts:        true,
					Graffiti:           "test graffiti",
					VLStartGracePeriod: 840,
					Mev:                true,
					MevImage:           "flashbots/mev-boost:latest",
					RelayURLs:          mainnetMevboostRelayListUris,
					ContainerTag:       "tag",
					JWTSecretPath:      filepath.Join(generationPath, "jwtsecret"),
				}
				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})
				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(1, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(0, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", true).Return(true, nil),
					prompter.EXPECT().Input("Mev-Boost image", "flashbots/mev-boost:latest", false, nil).Return("flashbots/mev-boost:latest", nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(1, nil),
					prompter.EXPECT().Select("Select validator client", "", ETHClients["validator"]).Return(1, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Graffiti to be used by the validator (press enter to skip it)", "", false, gomock.AssignableToTypeOf(ui.GraffitiValidator)).Return("test graffiti", nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					depsMgr.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil),
					depsMgr.EXPECT().DockerEngineIsOn().Return(nil),
					depsMgr.EXPECT().DockerComposeIsInstalled().Return(nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         NetworkMainnet,
						GenerationPath:  generationPath,
						From:            filepath.Join(generationPath, "keystore"),
						ContainerTag:    "tag",
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "full node with validator, hoodi",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"execution", "consensus", "validator", "mev-boost"},
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					ConsensusClient: &clients.Client{
						Name:  "prysm",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Prysm.String(),
					},
					ValidatorClient: &clients.Client{
						Name:  "prysm",
						Type:  "validator",
						Image: configs.ClientImages.Validator.Prysm.String(),
					},
					Network:            "hoodi",
					CheckpointSyncUrl:  "http://checkpoint.sync",
					FeeRecipient:       "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
					MapAllPorts:        true,
					Graffiti:           "test graffiti",
					VLStartGracePeriod: 840,
					Mev:                true,
					MevImage:           "flashbots/mev-boost:1.9rc3",
					RelayURLs:          configs.NetworksConfigs()[NetworkHoodi].RelayURLs,
					ContainerTag:       "tag",
					JWTSecretPath:      filepath.Join(generationPath, "jwtsecret"),
				}
				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})
				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(1, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(0, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", true).Return(true, nil),
					prompter.EXPECT().Confirm("Enable MEV Boost?", true).Return(true, nil),
					prompter.EXPECT().Input("Mev-Boost image", "flashbots/mev-boost:latest", false, nil).Return("flashbots/mev-boost:1.9rc3", nil),
					prompter.EXPECT().InputList("Insert relay URLs if you don't want to use the default values listed below", configs.NetworksConfigs()[NetworkHoodi].RelayURLs, gomock.AssignableToTypeOf(func([]string) error { return nil })).Return(configs.NetworksConfigs()[NetworkHoodi].RelayURLs, nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(1, nil),
					prompter.EXPECT().Select("Select validator client", "", ETHClients["validator"]).Return(1, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Graffiti to be used by the validator (press enter to skip it)", "", false, gomock.AssignableToTypeOf(ui.GraffitiValidator)).Return("test graffiti", nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address", "", true).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().Input("Withdrawal address", "", false, gomock.AssignableToTypeOf(func(s string) error { return ui.EthAddressValidator(s, true) })).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					depsMgr.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil),
					depsMgr.EXPECT().DockerEngineIsOn().Return(nil),
					depsMgr.EXPECT().DockerComposeIsInstalled().Return(nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         NetworkHoodi,
						GenerationPath:  generationPath,
						From:            filepath.Join(generationPath, "keystore"),
						ContainerTag:    "tag",
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "full node with Lido, hoodi",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"execution", "consensus", "validator", "mev-boost"},
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					ConsensusClient: &clients.Client{
						Name:  "prysm",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Prysm.String(),
					},
					ValidatorClient: &clients.Client{
						Name:  "prysm",
						Type:  "validator",
						Image: configs.ClientImages.Validator.Prysm.String(),
					},
					Network:            "hoodi",
					CheckpointSyncUrl:  "http://checkpoint.sync",
					FeeRecipient:       "0x9b108015fe433F173696Af3Aa0CF7CDb3E104258",
					MapAllPorts:        true,
					Graffiti:           "test graffiti",
					VLStartGracePeriod: 840,
					Mev:                true,
					MevImage:           "flashbots/mev-boost:1.9rc3",
					RelayURLs:          hoodiMevboostRelayListUris,
					ContainerTag:       "tag",
					JWTSecretPath:      filepath.Join(generationPath, "jwtsecret"),
				}
				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})
				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(1, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia}).Return(1, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(0, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", true).Return(true, nil),
					prompter.EXPECT().Input("Mev-Boost image", "flashbots/mev-boost:latest", false, nil).Return("flashbots/mev-boost:1.9rc3", nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(1, nil),
					prompter.EXPECT().Select("Select validator client", "", ETHClients["validator"]).Return(1, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Graffiti to be used by the validator (press enter to skip it)", "", false, gomock.AssignableToTypeOf(ui.GraffitiValidator)).Return("test graffiti", nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					depsMgr.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil),
					depsMgr.EXPECT().DockerEngineIsOn().Return(nil),
					depsMgr.EXPECT().DockerComposeIsInstalled().Return(nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         NetworkHoodi,
						GenerationPath:  generationPath,
						From:            filepath.Join(generationPath, "keystore"),
						ContainerTag:    "tag",
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "full node with nimbus validator mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"execution", "consensus", "validator", "mev-boost"},
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					ConsensusClient: &clients.Client{
						Name:  "nimbus",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Nimbus.String(),
					},
					ValidatorClient: &clients.Client{
						Name:  "nimbus",
						Type:  "validator",
						Image: configs.ClientImages.Validator.Nimbus.String(),
					},
					Network:            "mainnet",
					CheckpointSyncUrl:  "http://checkpoint.sync",
					FeeRecipient:       "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
					MapAllPorts:        true,
					Graffiti:           "test graffiti",
					VLStartGracePeriod: 840,
					Mev:                true,
					MevImage:           "flashbots/mev-boost:latest",
					RelayURLs:          configs.NetworksConfigs()[NetworkMainnet].RelayURLs,
					ContainerTag:       "tag",
					JWTSecretPath:      filepath.Join(generationPath, "jwtsecret"),
				}
				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})
				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(0, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", true).Return(true, nil),
					prompter.EXPECT().Confirm("Enable MEV Boost?", true).Return(true, nil),
					prompter.EXPECT().Input("Mev-Boost image", "flashbots/mev-boost:latest", false, nil).Return("flashbots/mev-boost:latest", nil),
					prompter.EXPECT().InputList("Insert relay URLs if you don't want to use the default values listed below", configs.NetworksConfigs()[NetworkMainnet].RelayURLs, gomock.AssignableToTypeOf(func([]string) error { return nil })).Return(configs.NetworksConfigs()[NetworkMainnet].RelayURLs, nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(4, nil),
					prompter.EXPECT().Select("Select validator client", "", ETHClients["validator"]).Return(4, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Graffiti to be used by the validator (press enter to skip it)", "", false, gomock.AssignableToTypeOf(ui.GraffitiValidator)).Return("test graffiti", nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address", "", true).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().Input("Withdrawal address", "", false, gomock.AssignableToTypeOf(func(s string) error { return ui.EthAddressValidator(s, true) })).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					depsMgr.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil),
					depsMgr.EXPECT().DockerEngineIsOn().Return(nil),
					depsMgr.EXPECT().DockerComposeIsInstalled().Return(nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator", "consensus"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "nimbus",
						Network:         NetworkMainnet,
						GenerationPath:  generationPath,
						From:            filepath.Join(generationPath, "keystore"),
						ContainerTag:    "tag",
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "aztec sequencer sepolia",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				keystorePath, err := filepath.Abs(filepath.Join("testdata", "cli_tests", "aztec_keystore", "valid_single", "keystore.json"))
				if err != nil {
					t.Fatal(err)
				}

				genData := generate.GenData{
					Services:      []string{"execution", "consensus", "aztec"},
					AztecNodeType: aztecNodeTypeSequencer,
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					ConsensusClient: &clients.Client{
						Name:  "lighthouse",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Lighthouse.String(),
					},
					AztecClient: &clients.Client{
						Name:  "aztec",
						Type:  "aztec",
						Image: configs.ClientImages.Aztec.Aztec.String(),
					},
					Network:                    "sepolia",
					CheckpointSyncUrl:          "http://checkpoint.sync",
					MapAllPorts:                false,
					ContainerTag:               "",
					JWTSecretPath:              filepath.Join(generationPath, "jwtsecret"),
					AztecSequencerKeystorePath: keystorePath,
					AztecP2pIp:                 "192.168.1.100",
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(2, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(4, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("", nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(0, nil),
					prompter.EXPECT().Select("Select aztec client", "", []string{"aztec", Randomize}).Return(0, nil),
					prompter.EXPECT().Select("Select aztec node type", "", []string{aztecNodeTypeFullNode, aztecNodeTypeSequencer}).Return(1, nil),
					prompter.EXPECT().Input("Aztec node P2P IP address", "", true, gomock.AssignableToTypeOf(func(string) error { return nil })).Return("192.168.1.100", nil),
					prompter.EXPECT().InputFilePath("Aztec sequencer keystore.json path", "", true, ".json").Return(keystorePath, nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(false, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Any()).DoAndReturn(func(opts actions.GenerateOptions) (generate.GenData, error) {
						if opts.GenerationPath != generationPath {
							t.Fatalf("unexpected GenerationPath. got %q want %q", opts.GenerationPath, generationPath)
						}

						got := opts.GenerationData
						if got.Network != genData.Network {
							t.Fatalf("unexpected Network. got %q want %q", got.Network, genData.Network)
						}
						if strings.Join(got.Services, ",") != strings.Join(genData.Services, ",") {
							t.Fatalf("unexpected Services. got %v want %v", got.Services, genData.Services)
						}
						if got.ExecutionClient == nil || got.ExecutionClient.Name != genData.ExecutionClient.Name {
							t.Fatalf("unexpected ExecutionClient. got %+v want name %q", got.ExecutionClient, genData.ExecutionClient.Name)
						}
						if got.ConsensusClient == nil || got.ConsensusClient.Name != genData.ConsensusClient.Name {
							t.Fatalf("unexpected ConsensusClient. got %+v want name %q", got.ConsensusClient, genData.ConsensusClient.Name)
						}
						if got.AztecClient == nil || got.AztecClient.Name != genData.AztecClient.Name {
							t.Fatalf("unexpected AztecClient. got %+v want name %q", got.AztecClient, genData.AztecClient.Name)
						}
						if got.AztecNodeType != genData.AztecNodeType {
							t.Fatalf("unexpected AztecNodeType. got %q want %q", got.AztecNodeType, genData.AztecNodeType)
						}
						if got.AztecSequencerKeystorePath != genData.AztecSequencerKeystorePath {
							t.Fatalf("unexpected AztecSequencerKeystorePath. got %q want %q", got.AztecSequencerKeystorePath, genData.AztecSequencerKeystorePath)
						}
						if got.AztecP2pIp != genData.AztecP2pIp {
							t.Fatalf("unexpected AztecP2pIp. got %q want %q", got.AztecP2pIp, genData.AztecP2pIp)
						}
						return got, nil
					}),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "aztec full node sepolia",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()

				genData := generate.GenData{
					Services:      []string{"execution", "consensus", "aztec"},
					AztecNodeType: aztecNodeTypeFullNode,
					AztecP2pIp:    "192.168.1.100",
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					ConsensusClient: &clients.Client{
						Name:  "lighthouse",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Lighthouse.String(),
					},
					AztecClient: &clients.Client{
						Name:  "aztec",
						Type:  "aztec",
						Image: configs.ClientImages.Aztec.Aztec.String(),
					},
					Network:           "sepolia",
					CheckpointSyncUrl: "http://checkpoint.sync",
					MapAllPorts:       false,
					ContainerTag:      "",
					JWTSecretPath:     filepath.Join(generationPath, "jwtsecret"),
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select node setup", "", []string{sedgeOpts.EthereumNode, sedgeOpts.LidoNode}).Return(0, nil),
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkHoodi, NetworkSepolia, NetworkGnosis, NetworkChiado}).Return(2, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeAztec}).Return(4, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("", nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(0, nil),
					prompter.EXPECT().Select("Select aztec client", "", []string{"aztec", Randomize}).Return(0, nil),
					prompter.EXPECT().Select("Select aztec node type", "", []string{aztecNodeTypeFullNode, aztecNodeTypeSequencer}).Return(0, nil),
					prompter.EXPECT().Input("Aztec node P2P IP address", "", true, gomock.AssignableToTypeOf(func(string) error { return nil })).Return("192.168.1.100", nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(false, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to enable the monitoring stack?", false).Return(false, nil),
					sedgeActions.EXPECT().Generate(gomock.Any()).DoAndReturn(func(opts actions.GenerateOptions) (generate.GenData, error) {
						if opts.GenerationPath != generationPath {
							t.Fatalf("unexpected GenerationPath. got %q want %q", opts.GenerationPath, generationPath)
						}

						got := opts.GenerationData
						if got.Network != genData.Network {
							t.Fatalf("unexpected Network. got %q want %q", got.Network, genData.Network)
						}
						if strings.Join(got.Services, ",") != strings.Join(genData.Services, ",") {
							t.Fatalf("unexpected Services. got %v want %v", got.Services, genData.Services)
						}
						if got.ExecutionClient == nil || got.ExecutionClient.Name != genData.ExecutionClient.Name {
							t.Fatalf("unexpected ExecutionClient. got %+v want name %q", got.ExecutionClient, genData.ExecutionClient.Name)
						}
						if got.ConsensusClient == nil || got.ConsensusClient.Name != genData.ConsensusClient.Name {
							t.Fatalf("unexpected ConsensusClient. got %+v want name %q", got.ConsensusClient, genData.ConsensusClient.Name)
						}
						if got.AztecClient == nil || got.AztecClient.Name != genData.AztecClient.Name {
							t.Fatalf("unexpected AztecClient. got %+v want name %q", got.AztecClient, genData.AztecClient.Name)
						}
						if got.AztecNodeType != genData.AztecNodeType {
							t.Fatalf("unexpected AztecNodeType. got %q want %q", got.AztecNodeType, genData.AztecNodeType)
						}
						if got.AztecSequencerKeystorePath != "" {
							t.Fatalf("unexpected AztecSequencerKeystorePath. got %q want empty", got.AztecSequencerKeystorePath)
						}
						if got.AztecP2pIp != genData.AztecP2pIp {
							t.Fatalf("unexpected AztecP2pIp. got %q want %q", got.AztecP2pIp, genData.AztecP2pIp)
						}
						return got, nil
					}),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			sedgeActions := sedge_mocks.NewMockSedgeActions(ctrl)
			prompter := sedge_mocks.NewMockPrompter(ctrl)
			depsMgr := sedge_mocks.NewMockDependenciesManager(ctrl)
			monitoringMgr := sedge_mocks.NewMockMonitoringManager(ctrl)
			defer ctrl.Finish()

			tt.setup(t, sedgeActions, prompter, depsMgr)

			c := CliCmd(prompter, sedgeActions, depsMgr, monitoringMgr)
			c.Execute()
		})
	}
}
