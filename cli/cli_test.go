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
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
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

	ETHClients := map[string][]string{
		"execution": clients.AllClients["execution"],
		"consensus": clients.AllClients["consensus"],
		"validator": clients.AllClients["validator"],
		"starknet":  clients.AllClients["starknet"],
	}
	ETHClients["execution"] = append(ETHClients["execution"], "randomize")
	ETHClients["consensus"] = append(ETHClients["consensus"], "randomize")
	ETHClients["validator"] = append(ETHClients["validator"], "randomize")
	ETHClients["starknet"] = append(ETHClients["starknet"], "randomize")

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
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeStarknet}).Return(0, nil),
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
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().Input("Withdrawal address", "", false, gomock.AssignableToTypeOf(func(s string) error { return ui.EthAddressValidator(s, true) })).Return("0x00000007abca72jmd83jd8u3jd9kdn32j38abc", nil),
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
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeStarknet}).Return(0, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", true).Return(false, nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(1, nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "full node without validator holesky",
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
						Name:  "lodestar",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Lodestar.String(),
					},
					Network:           "holesky",
					CheckpointSyncUrl: "https://checkpoint-sync.holesky.ethpandaops.io/",
					FeeRecipient:      "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
					MapAllPorts:       true,
					ContainerTag:      "tag",
					JWTSecretPath:     filepath.Join(generationPath, "jwtsecret"),
				}
				gomock.InOrder(
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(5, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(0, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", true).Return(false, nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(3, nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("https://checkpoint-sync.holesky.ethpandaops.io/", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
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
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeStarknet}).Return(1, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(2, nil),
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
			name: "execution node holesky",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"execution"},
					ExecutionClient: &clients.Client{
						Name:  "nethermind",
						Type:  "execution",
						Image: configs.ClientImages.Execution.Nethermind.String(),
					},
					Network:      "holesky",
					MapAllPorts:  true,
					ContainerTag: "tag",
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(5, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(1, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(2, nil),
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
					Network:           NetworkGoerli,
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
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(1, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(2, nil),
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
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "consensus node holesky",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"consensus"},
					ConsensusClient: &clients.Client{
						Name:  "lodestar",
						Type:  "consensus",
						Image: configs.ClientImages.Consensus.Lodestar.String(),
					},
					Network:           NetworkHolesky,
					CheckpointSyncUrl: "https://checkpoint-sync.holesky.ethpandaops.io/",
					FeeRecipient:      "0x2d07a21ebadde0c13e8b91022a7e5732eb6bf5d5",
					MapAllPorts:       false,
					ExecutionApiUrl:   "http://execution:5051",
					ExecutionAuthUrl:  "http://execution:5051",
					ContainerTag:      "tag",
					JWTSecretPath:     filepath.Join(generationPath, "jwtsecret"),
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(5, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(2, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(3, nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("https://checkpoint-sync.holesky.ethpandaops.io/", nil),
					prompter.EXPECT().InputURL("Execution API URL", "", true).Return("http://execution:5051", nil),
					prompter.EXPECT().InputURL("Execution Auth API URL", "", true).Return("http://execution:5051", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("0x2d07a21ebadde0c13e8b91022a7e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(false, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
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
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeStarknet}).Return(3, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select validator client", "", ETHClients["validator"]).Return(1, nil),
					prompter.EXPECT().InputURL("Consensus API URL", "", true).Return("http://localhost:5051", nil),
					prompter.EXPECT().Input("Graffiti to be used by the validator (press enter to skip it)", "", false, gomock.AssignableToTypeOf(ui.GraffitiValidator)).Return("test graffiti", nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address", "", true).Return("0x2d07a31ebadce0a13e8a91022a5e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Enable MEV Boost?", true).Return(true, nil),
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
			name: "starknet node mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"starknet"},
					StarknetClient: &clients.Client{
						Name:  "juno",
						Type:  "starknet",
						Image: configs.ClientImages.Starknet.Juno.String(),
					},
					ExecutionApiUrl: "ws://execution:8545",
					Network:         "mainnet",
					MapAllPorts:     true,
					ContainerTag:    "tag",
					Full:            false,
					RemoteDbURL:     "http://example.com",
					JWTSecretPath:   filepath.Join(generationPath, "jwtsecret"),
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeStarknet}).Return(4, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select starknet client", "", ETHClients["starknet"]).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to set up a full starknet node (with execution and consensus)?", true).Return(false, nil),
					prompter.EXPECT().InputURL("Input Eth 1 Endpoint", "", true).Return("ws://execution:8545", nil),
					prompter.EXPECT().InputURL("Input GRPC URL of a remote Juno node (press Enter to skip)", "", false).Return("http://example.com", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "starknet node sepolia",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"starknet"},
					StarknetClient: &clients.Client{
						Name:  "juno",
						Type:  "starknet",
						Image: configs.ClientImages.Starknet.Juno.String(),
					},
					ExecutionApiUrl: "ws://execution:8545",
					Network:         "sepolia",
					MapAllPorts:     true,
					ContainerTag:    "tag",
					Full:            false,
					RemoteDbURL:     "http://example.com",
					JWTSecretPath:   filepath.Join(generationPath, "jwtsecret"),
				}

				gomock.InOrder(
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(2, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeStarknet}).Return(4, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select starknet client", "", ETHClients["starknet"]).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to set up a full starknet node (with execution and consensus)?", true).Return(false, nil),
					prompter.EXPECT().InputURL("Input Eth 1 Endpoint", "", true).Return("ws://execution:8545", nil),
					prompter.EXPECT().InputURL("Input GRPC URL of a remote Juno node (press Enter to skip)", "", false).Return("http://example.com", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				)
			},
		},
		{
			name: "full starknet node with consensus and execution",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter, depsMgr *sedge_mocks.MockDependenciesManager) {
				generationPath := t.TempDir()
				genData := generate.GenData{
					Services: []string{"starknet", "execution", "consensus"},
					StarknetClient: &clients.Client{
						Name:  "juno",
						Type:  "starknet",
						Image: configs.ClientImages.Starknet.Juno.String(),
					},
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
					Network:            "mainnet",
					CheckpointSyncUrl:  "http://checkpoint.sync",
					FeeRecipient:       "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
					MapAllPorts:        true,
					SLStartGracePeriod: 840,
					ContainerTag:       "tag",
					JWTSecretPath:      filepath.Join(generationPath, "jwtsecret"),
				}
				gomock.InOrder(
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkHolesky}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator, NodeTypeStarknet}).Return(4, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false, nil).Return(generationPath, nil),
					prompter.EXPECT().Input("Container tag, sedge will add to each container and the network, a suffix with the tag", "", false, nil).Return("tag", nil),
					prompter.EXPECT().Select("Select starknet client", "", ETHClients["starknet"]).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to set up a full starknet node (with execution and consensus)?", true).Return(true, nil),
					prompter.EXPECT().Select("Select execution client", "", ETHClients["execution"]).Return(0, nil),
					prompter.EXPECT().Select("Select consensus client", "", ETHClients["consensus"]).Return(1, nil),
					prompter.EXPECT().InputInt64("Starknet grace period. This is the number of epochs the starknet client will wait before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().InputURL("Checkpoint sync URL", configs.NetworksConfigs()[genData.Network].CheckpointSyncURL, false).Return("http://checkpoint.sync", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address (press enter to skip it)", "", false).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: genData,
					})).Return(genData, nil),
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
			defer ctrl.Finish()

			tt.setup(t, sedgeActions, prompter, depsMgr)

			c := CliCmd(prompter, sedgeActions, depsMgr)
			c.Execute()
		})
	}
}
