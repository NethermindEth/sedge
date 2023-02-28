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
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/pkg/keystores/testdata"
	"github.com/NethermindEth/sedge/internal/utils"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/NethermindEth/sedge/test"
	"github.com/golang/mock/gomock"
)

func TestCli_FullNode(t *testing.T) {
	configs.InitNetworksConfigs()

	tests := []struct {
		name  string
		setup func(*testing.T, *sedge_mocks.MockSedgeActions, *sedge_mocks.MockPrompter)
	}{
		{
			name: "full node with validator mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()

				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", false).Return(true, nil),
					prompter.EXPECT().Input("Mev-Boost image", "flashbots/mev-boost:latest", false).Return("flashbots/mev-boost:latest", nil),
					prompter.EXPECT().InputList("Relay URLs", configs.MainnetRelayURLs(), nil).Return(configs.MainnetRelayURLs(), nil),
					prompter.EXPECT().Select("Select execution client", "", []string{"besu", "erigon", "geth", "nethermind", "randomize"}).Return(3, nil),
					prompter.EXPECT().Select("Select consensus client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().Select("Select validator client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Graffiti", "", false).Return("test graffiti", nil),
					prompter.EXPECT().Input("Checkpoint sync URL", "", false).Return("http://localhost:5052", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address.", "", true).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
							Services: []string{"execution", "consensus", "validator"},
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
							CheckpointSyncUrl:  "http://localhost:5052",
							FeeRecipient:       "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
							MapAllPorts:        true,
							Graffiti:           "test graffiti",
							VLStartGracePeriod: 840000000000,
							Mev:                true,
							MevImage:           "flashbots/mev-boost:latest",
							RelayURL:           strings.Join(configs.MainnetRelayURLs(), ","),
						},
					})).Return(nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().Input("Withdrawal address", "", false).Return("0x00000007abca72jmd83jd8u3jd9kdn32j38abc", nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         NetworkMainnet,
						GenerationPath:  generationPath,
						From:            filepath.Join(generationPath, "keystores"),
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
		{
			name: "full node without validator mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", false).Return(false, nil),
					prompter.EXPECT().Select("Select execution client", "", []string{"besu", "erigon", "geth", "nethermind", "randomize"}).Return(3, nil),
					prompter.EXPECT().Select("Select consensus client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().Input("Checkpoint sync URL", "", false).Return("http://localhost:5052", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address.", "", true).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
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
							CheckpointSyncUrl: "http://localhost:5052",
							FeeRecipient:      "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
							MapAllPorts:       true,
							Mev:               true,
						},
					})).Return(nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
		{
			name: "full node with validator in custom network",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()
				testData := t.TempDir()
				keystoreDir, err := keystore_test_data.SetupTestDataDir(t)
				if err != nil {
					t.Error(err)
				}

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(5, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(0, nil),
					prompter.EXPECT().Confirm("Do you want to set up a validator?", false).Return(true, nil),
					prompter.EXPECT().InputFilePath("Custom network config file path", "", true).Return("testdata/networkConfig.json", nil),
					prompter.EXPECT().InputFilePath("Custom ChainSpec", "", true).Return("testdata/chainSpec.json", nil),
					prompter.EXPECT().InputFilePath("Custom Genesis", "", true).Return("testdata/genesis.json", nil),
					prompter.EXPECT().Input("Custom TTD (Terminal Total Difficulty)", "", false).Return("58750000000000", nil),
					prompter.EXPECT().Input("Custom deploy block", "", false).Return("2355021", nil),
					prompter.EXPECT().InputList("Execution boot nodes", gomock.Len(0), gomock.AssignableToTypeOf(utils.BootNodesValidator)).Return([]string{"enode://ecnode1", "enode://ecnode2"}, nil),
					prompter.EXPECT().InputList("Consensus boot nodes", gomock.Len(0), gomock.AssignableToTypeOf(utils.BootNodesValidator)).Return([]string{"enode://ccnode1", "enode://ccnode2"}, nil),
					prompter.EXPECT().Select("Select execution client", "", []string{"besu", "nethermind", "randomize"}).Return(1, nil),
					prompter.EXPECT().Select("Select consensus client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().Select("Select validator client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Graffiti", "", false).Return("test graffiti", nil),
					prompter.EXPECT().Input("Checkpoint sync URL", "", false).Return("http://localhost:5052", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address.", "", true).Return("0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(1, nil),
					prompter.EXPECT().InputFilePath("JWT path", "", true).Return(filepath.Join(testData, "jwt"), nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
							Services: []string{"execution", "consensus", "validator"},
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
							Network:                 NetworkCustom,
							CheckpointSyncUrl:       "http://localhost:5052",
							FeeRecipient:            "0x2d07a21ebadde0c13e6b91022a7e5722eb6bf5d5",
							MapAllPorts:             true,
							Graffiti:                "test graffiti",
							VLStartGracePeriod:      840000000000,
							Mev:                     true,
							CustomNetworkConfigPath: "testdata/networkConfig.json",
							CustomChainSpecPath:     "testdata/chainSpec.json",
							CustomGenesisPath:       "testdata/genesis.json",
							CustomTTD:               "58750000000000",
							CustomDeployBlock:       "2355021",
							ECBootnodes:             &[]string{"enode://ecnode1", "enode://ecnode2"},
							CCBootnodes:             &[]string{"enode://ccnode1", "enode://ccnode2"},
							JWTSecretPath:           filepath.Join(testData, "jwt"),
						},
					})).Return(nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(1, nil),
					prompter.EXPECT().Input("Keystore path", "", true).Return(keystoreDir, nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         "custom",
						From:            keystoreDir,
						GenerationPath:  generationPath,
						CustomConfig: actions.ImportValidatorKeysCustomOptions{
							NetworkConfigPath: "testdata/networkConfig.json",
							GenesisPath:       "testdata/genesis.json",
						},
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
		{
			name: "execution node",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(1, nil),
					prompter.EXPECT().Select("Select execution client", "", []string{"besu", "erigon", "geth", "nethermind", "randomize"}).Return(3, nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(2, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
							Services: []string{"execution"},
							ExecutionClient: &clients.Client{
								Name:  "nethermind",
								Type:  "execution",
								Image: configs.ClientImages.Execution.Nethermind.String(),
							},
							Network:     "mainnet",
							MapAllPorts: true,
						},
					})).Return(nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(true, nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution"},
					}),
					sedgeActions.EXPECT().RunContainers(actions.RunContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution"},
					}),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
		{
			name: "execution node, custom network",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(5, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(1, nil),
					prompter.EXPECT().Select("Select execution client", "", []string{"besu", "nethermind", "randomize"}).Return(1, nil),
					prompter.EXPECT().InputFilePath("Custom ChainSpec", "", true).Return("testdata/chainSpec.json", nil),
					prompter.EXPECT().Input("Custom TTD (Terminal Total Difficulty)", "", false).Return("58750000000000", nil),
					prompter.EXPECT().InputList("Execution boot nodes", gomock.Len(0), gomock.AssignableToTypeOf(utils.BootNodesValidator)).Return([]string{"enode://ecnode1", "enode://ecnode2"}, nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
							Services: []string{"execution"},
							ExecutionClient: &clients.Client{
								Name:  "nethermind",
								Type:  "execution",
								Image: configs.ClientImages.Execution.Nethermind.String(),
							},
							Network:             NetworkCustom,
							MapAllPorts:         true,
							CustomChainSpecPath: "testdata/chainSpec.json",
							CustomTTD:           "58750000000000",
							ECBootnodes:         &[]string{"enode://ecnode1", "enode://ecnode2"},
						},
					})).Return(nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(true, nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution"},
					}),
					sedgeActions.EXPECT().RunContainers(actions.RunContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"execution"},
					}),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
		{
			name: "consensus node",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(1, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(2, nil),
					prompter.EXPECT().Select("Select consensus client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().Input("Checkpoint sync URL", "", false).Return("http://localhost:5052", nil),
					prompter.EXPECT().Input("Mev-Boost endpoint", "", false).Return("http://localhost:3030", nil),
					prompter.EXPECT().Input("Execution API URL", "", false).Return("http://localhost:5051", nil),
					prompter.EXPECT().Input("Execution Auth API URL", "", false).Return("http://localhost:5051", nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address.", "", true).Return("0x2d07a21ebadde0c13e8b91022a7e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
							Services: []string{"consensus"},
							ConsensusClient: &clients.Client{
								Name:  "prysm",
								Type:  "consensus",
								Image: configs.ClientImages.Consensus.Prysm.String(),
							},
							Network:           NetworkGoerli,
							CheckpointSyncUrl: "http://localhost:5052",
							FeeRecipient:      "0x2d07a21ebadde0c13e8b91022a7e5732eb6bf5d5",
							MapAllPorts:       true,
							ExecutionApiUrl:   "http://localhost:5051",
							ExecutionAuthUrl:  "http://localhost:5051",
							MevBoostEndpoint:  "http://localhost:3030",
						},
					})).Return(nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
		{
			name: "consensus node, custom network",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(5, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(2, nil),
					prompter.EXPECT().Select("Select consensus client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().InputFilePath("Custom network config file path", "", true).Return("testdata/networkConfig.json", nil),
					prompter.EXPECT().InputFilePath("Custom Genesis", "", true).Return("testdata/genesis.json", nil),
					prompter.EXPECT().Input("Custom deploy block", "", false).Return("2355021", nil),
					prompter.EXPECT().InputList("Consensus boot nodes", gomock.Len(0), gomock.AssignableToTypeOf(utils.BootNodesValidator)).Return([]string{"enode://ccnode1", "enode://ccnode2"}, nil),
					prompter.EXPECT().Input("Execution API URL", "", false).Return("http://localhost:5051", nil),
					prompter.EXPECT().Input("Execution Auth API URL", "", false).Return("http://localhost:5051", nil),
					prompter.EXPECT().Select("Select JWT source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address.", "", true).Return("0x2d07a21ebadce0c13e8a91022a7e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Do you want to expose all ports?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
							Services: []string{"consensus"},
							ConsensusClient: &clients.Client{
								Name:  "prysm",
								Type:  "consensus",
								Image: configs.ClientImages.Consensus.Prysm.String(),
							},
							Network:                 NetworkCustom,
							FeeRecipient:            "0x2d07a21ebadce0c13e8a91022a7e5732eb6bf5d5",
							MapAllPorts:             true,
							ExecutionApiUrl:         "http://localhost:5051",
							ExecutionAuthUrl:        "http://localhost:5051",
							CustomNetworkConfigPath: "testdata/networkConfig.json",
							CustomGenesisPath:       "testdata/genesis.json",
							CustomDeployBlock:       "2355021",
							CCBootnodes:             &[]string{"enode://ccnode1", "enode://ccnode2"},
						},
					})).Return(nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
		{
			name: "validator mainnet",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()

				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(0, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(3, nil),
					prompter.EXPECT().Select("Select validator client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Consensus API URL", "", false).Return("http://localhost:5051", nil),
					prompter.EXPECT().Input("Graffiti", "", false).Return("test graffiti", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address.", "", true).Return("0x2d07a31ebadce0a13e8a91022a5e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Enable MEV Boost?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
							Services: []string{"validator"},
							ValidatorClient: &clients.Client{
								Name:  "prysm",
								Type:  "validator",
								Image: configs.ClientImages.Validator.Prysm.String(),
							},
							Network:             "mainnet",
							FeeRecipient:        "0x2d07a31ebadce0a13e8a91022a5e5732eb6bf5d5",
							Graffiti:            "test graffiti",
							VLStartGracePeriod:  840000000000,
							MevBoostOnValidator: true,
							ConsensusApiUrl:     "http://localhost:5051",
						},
					})).Return(nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().Input("Withdrawal address", "", false).Return("0x2d07a21ebadde0c13e6b91022a7e5732eb6bf5d5", nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         NetworkMainnet,
						From:            filepath.Join(generationPath, "keystores"),
						GenerationPath:  generationPath,
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
		{
			name: "validator custom",
			setup: func(t *testing.T, sedgeActions *sedge_mocks.MockSedgeActions, prompter *sedge_mocks.MockPrompter) {
				generationPath := t.TempDir()

				sedgeActions.EXPECT().GetCommandRunner().Return(&test.SimpleCMDRunner{})

				callsSequence := []*gomock.Call{
					prompter.EXPECT().Select("Select network", "", []string{NetworkMainnet, NetworkGoerli, NetworkSepolia, NetworkGnosis, NetworkChiado, NetworkCustom}).Return(5, nil),
					prompter.EXPECT().Select("Select node type", "", []string{NodeTypeFullNode, NodeTypeExecution, NodeTypeConsensus, NodeTypeValidator}).Return(3, nil),
					prompter.EXPECT().Select("Select validator client", "", []string{"lighthouse", "lodestar", "prysm", "teku", "randomize"}).Return(2, nil),
					prompter.EXPECT().InputFilePath("Custom network config file path", "", true).Return("testdata/networkConfig.json", nil),
					prompter.EXPECT().InputFilePath("Custom Genesis", "", true).Return("testdata/genesis.json", nil),
					prompter.EXPECT().Input("Custom deploy block", "", false).Return("2355021", nil),
					prompter.EXPECT().InputList("Consensus boot nodes", gomock.Len(0), gomock.AssignableToTypeOf(utils.BootNodesValidator)).Return([]string{"enode://ccnode1", "enode://ccnode2"}, nil),
					prompter.EXPECT().InputInt64("Validator grace period. This is the number of epochs the validator will wait for security reasons before starting", int64(1)).Return(int64(2), nil),
					prompter.EXPECT().Input("Consensus API URL", "", false).Return("http://localhost:5051", nil),
					prompter.EXPECT().Input("Graffiti", "", false).Return("test graffiti", nil),
					prompter.EXPECT().EthAddress("Please enter the Fee Recipient address.", "", true).Return("0x2d07a31ebadce0a13e8a91022a5e5732eb6bf5d5", nil),
					prompter.EXPECT().Confirm("Enable MEV Boost?", false).Return(true, nil),
					prompter.EXPECT().Input("Generation path", configs.DefaultAbsSedgeDataPath, false).Return(generationPath, nil),
					sedgeActions.EXPECT().Generate(gomock.Eq(actions.GenerateOptions{
						GenerationPath: generationPath,
						GenerationData: generate.GenData{
							Services: []string{"validator"},
							ValidatorClient: &clients.Client{
								Name:  "prysm",
								Type:  "validator",
								Image: configs.ClientImages.Validator.Prysm.String(),
							},
							Network:                 NetworkCustom,
							FeeRecipient:            "0x2d07a31ebadce0a13e8a91022a5e5732eb6bf5d5",
							Graffiti:                "test graffiti",
							VLStartGracePeriod:      840000000000,
							MevBoostOnValidator:     true,
							ConsensusApiUrl:         "http://localhost:5051",
							CustomNetworkConfigPath: "testdata/networkConfig.json",
							CustomGenesisPath:       "testdata/genesis.json",
							CustomDeployBlock:       "2355021",
							CCBootnodes:             &[]string{"enode://ccnode1", "enode://ccnode2"},
						},
					})).Return(nil),
					prompter.EXPECT().Select("Select keystore source", "", []string{SourceTypeCreate, SourceTypeExisting, SourceTypeSkip}).Return(0, nil),
					prompter.EXPECT().Select("Select mnemonic source", "", []string{SourceTypeCreate, SourceTypeExisting}).Return(0, nil),
					prompter.EXPECT().Select("Select passphrase source", "", []string{SourceTypeRandom, SourceTypeExisting, SourceTypeCreate}).Return(0, nil),
					prompter.EXPECT().Input("Withdrawal address", "", false).Return("0x2d07a21ebadde0c13e6b91022a7e5732eb6bf5d5", nil),
					prompter.EXPECT().InputInt64("Number of validators", int64(1)).Return(int64(1), nil),
					prompter.EXPECT().InputInt64("Existing validators. This number will be used as the initial index for the generated keystores.", int64(0)).Return(int64(0), nil),
					sedgeActions.EXPECT().SetupContainers(actions.SetupContainersOptions{
						GenerationPath: generationPath,
						Services:       []string{"validator"},
					}),
					sedgeActions.EXPECT().ImportValidatorKeys(actions.ImportValidatorKeysOptions{
						ValidatorClient: "prysm",
						Network:         NetworkCustom,
						From:            filepath.Join(generationPath, "keystores"),
						GenerationPath:  generationPath,
						CustomConfig: actions.ImportValidatorKeysCustomOptions{
							NetworkConfigPath: "testdata/networkConfig.json",
							GenesisPath:       "testdata/genesis.json",
						},
					}).Return(nil),
					prompter.EXPECT().Confirm("Do you want to import slashing protection data?", false).Return(false, nil),
					prompter.EXPECT().Confirm("Run services now?", false).Return(false, nil),
				}

				var prevCall *gomock.Call
				for _, call := range callsSequence {
					call.Times(1)
					if prevCall != nil {
						call.After(prevCall)
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			sedgeActions := sedge_mocks.NewMockSedgeActions(ctrl)
			prompter := sedge_mocks.NewMockPrompter(ctrl)
			defer ctrl.Finish()

			tt.setup(t, sedgeActions, prompter)

			c := CliCmd(prompter, sedgeActions)
			c.Execute()
		})
	}
}
