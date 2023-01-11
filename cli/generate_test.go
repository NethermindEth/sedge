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
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/test"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type subCmd struct {
	name string
	args []string
}

func (cmd subCmd) argsList() []string {
	s := make([]string, 0)
	s = append(s, cmd.name)
	s = append(s, cmd.args...)
	return s
}

type globalFlags struct {
	install        bool
	generationPath string
	network        string
	logging        string
}

type generateCmdTestCase struct {
	name       string
	configPath string
	subCommand subCmd
	args       GenCmdFlags
	globalArgs globalFlags
	isErr      bool
}

func (flags *globalFlags) argsList() []string {
	s := make([]string, 0)
	if flags.install {
		s = append(s, "--install")
	}
	if flags.generationPath != "" {
		s = append(s, "--path", flags.generationPath)
	}
	if flags.network != "" {
		s = append(s, "--network", flags.network)
	}
	if flags.logging != "" {
		s = append(s, "--logging", flags.logging)
	}
	return s
}

func (flags *GenCmdFlags) argsList() []string {
	s := make([]string, 0)
	if flags.executionName != "" {
		s = append(s, "-e", flags.executionName)
	}
	if flags.consensusName != "" {
		s = append(s, "-c", flags.consensusName)
	}
	if flags.validatorName != "" {
		s = append(s, "-v", flags.validatorName)
	}
	if flags.checkpointSyncUrl != "" {
		s = append(s, "--checkpoint-sync-url", flags.checkpointSyncUrl)
	}
	if flags.executionAuthUrl != "" {
		s = append(s, "--execution-auth-url", flags.executionAuthUrl)
	}
	if flags.executionApiUrl != "" {
		s = append(s, "--execution-api-url", flags.executionApiUrl)
	}

	if flags.noMev {
		s = append(s, "--no-mev-boost")
	}
	if flags.noValidator {
		s = append(s, "--no-validator")
	}
	if flags.mevImage != "" {
		s = append(s, "--mev-boost-image", flags.mevImage)
	}
	if flags.jwtPath != "" {
		s = append(s, "--jwt-secret-path", flags.jwtPath)
	}
	if flags.graffiti != "" {
		s = append(s, "--graffiti", flags.graffiti)
	}
	if flags.consensusApiUrl != "" {
		s = append(s, "--consensus-url", flags.consensusApiUrl)
	}
	if flags.feeRecipient != "" {
		s = append(s, "--fee-recipient", flags.feeRecipient)
	}
	if flags.mapAllPorts {
		s = append(s, "--map-all")
	}
	if flags.customTTD != "" {
		s = append(s, "--custom-ttd", flags.customTTD)
	}
	if flags.customChainSpec != "" {
		s = append(s, "--custom-chainSpec", flags.customChainSpec)
	}
	if flags.customNetworkConfig != "" {
		s = append(s, "--custom-config", flags.customNetworkConfig)
	}
	if flags.customGenesis != "" {
		s = append(s, "--custom-genesis", flags.customGenesis)
	}
	if flags.customDeployBlock != "" {
		s = append(s, "--custom-deploy-block", flags.customDeployBlock)
	}
	if flags.customEnodes != nil && len(*flags.customEnodes) != 0 {
		s = append(s, "--execution-bootnodes", strings.Join(*flags.customEnodes, ","))
	}
	if flags.customEnrs != nil && len(*flags.customEnrs) != 0 {
		s = append(s, "--consensus-bootnodes", strings.Join(*flags.customEnrs, ","))
	}
	return s
}

func (flags *GenCmdFlags) toString() string {
	return strings.Join(flags.argsList(), " ")
}

func buildGenerateTestCase(
	t *testing.T,
	name,
	caseTestDataDir string,
	args GenCmdFlags,
	globalArgs globalFlags,
	subCommand subCmd,
	isErr bool,
) *generateCmdTestCase {
	tc := generateCmdTestCase{}
	configPath := t.TempDir()

	err := test.PrepareTestCaseDir(filepath.Join("testdata", "cli_tests", caseTestDataDir, "config"), configPath)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}
	dcPath := filepath.Join(configPath, "docker-compose-scripts")
	err = os.Mkdir(dcPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	tc.name = name
	tc.args = args
	tc.globalArgs = globalArgs
	tc.globalArgs.generationPath = t.TempDir()
	tc.subCommand = subCommand
	tc.configPath = filepath.Join(configPath, "config.yaml")
	tc.isErr = isErr
	return &tc
}

func TestGenerateCmd(t *testing.T) {
	configs.InitNetworksConfigs()
	tcs := []generateCmdTestCase{
		*buildGenerateTestCase(
			t,
			"full-node Fixed clients", "case_1",
			GenCmdFlags{
				executionName: "nethermind",
				consensusName: "lighthouse",
				validatorName: "lighthouse",
				feeRecipient:  "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install: false,
				network: "mainnet",
				logging: "",
			},
			subCmd{
				name: "full-node",
				args: []string{},
			},
			false,
		),
		*buildGenerateTestCase(
			t,
			"Wrong sub command", "case_1",
			GenCmdFlags{
				executionName:     "nethermind",
				validatorName:     "lighthouse",
				feeRecipient:      "0x0000000000000000000000000000000000000000",
				checkpointSyncUrl: "http://localhost:8545",
			},
			globalFlags{
				install: false,
				network: "",
				logging: "",
			},
			subCmd{
				name: "full",
				args: []string{},
			},
			true,
		),
		*buildGenerateTestCase(
			t,
			"Missing validator client", "case_1",
			GenCmdFlags{
				executionName: "nethermind",
				consensusName: "lighthouse",
				feeRecipient:  "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "mainnet",
				logging:        "",
			},
			subCmd{
				name: "full-node",
				args: []string{},
			},
			false,
		),
		*buildGenerateTestCase(
			t,
			"Good network input", "case_1",
			GenCmdFlags{
				feeRecipient: "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "mainnet",
				logging:        "",
			},
			subCmd{
				name: "full-node",
				args: []string{},
			},
			false,
		),
		*buildGenerateTestCase(
			t,
			"Bad network input", "case_1",
			GenCmdFlags{
				executionName: "nethermind",
				consensusName: "lighthouse",
				feeRecipient:  "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "wrong",
				logging:        "",
			},
			subCmd{
				name: "consensus",
				args: []string{},
			},
			true,
		),
		*buildGenerateTestCase(
			t,
			"Consensus fixed client", "case_1",
			GenCmdFlags{
				executionAuthUrl: "http://localhost:8545",
				executionApiUrl:  "http://localhost:8545",
				feeRecipient:     "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "consensus",
				args: []string{"teku"},
			},
			false,
		),
		*buildGenerateTestCase(
			t,
			"Consensus Missing execution-auth-url", "case_1",
			GenCmdFlags{
				executionApiUrl: "http://localhost:8545",
				feeRecipient:    "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "consensus",
				args: []string{"teku"},
			},
			true,
		),
		*buildGenerateTestCase(
			t,
			"Consensus Missing execution-api-url", "case_1",
			GenCmdFlags{
				executionAuthUrl: "http://localhost:8545",
				feeRecipient:     "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "consensus",
				args: []string{"teku"},
			},
			true,
		),
		*buildGenerateTestCase(
			t,
			"Consensus wrong client", "case_1",
			GenCmdFlags{
				executionAuthUrl: "http://localhost:8545",
				executionApiUrl:  "http://localhost:8545",
				feeRecipient:     "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "consensus",
				args: []string{"wrong"},
			},
			true,
		),
		*buildGenerateTestCase(
			t,
			"Validator", "case_1",
			GenCmdFlags{
				consensusApiUrl: "http://localhost:4000",
				feeRecipient:    "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "validator",
				args: []string{},
			},
			false,
		),
		*buildGenerateTestCase(
			t,
			"Validator missing consensus-api", "case_1",
			GenCmdFlags{
				feeRecipient: "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "validator",
				args: []string{},
			},
			true,
		),
		*buildGenerateTestCase(
			t,
			"validator good client", "case_1",
			GenCmdFlags{
				consensusApiUrl: "http://localhost:4000",
				feeRecipient:    "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "validator",
				args: []string{"teku"},
			},
			false,
		),
		*buildGenerateTestCase(
			t,
			"MevBoost", "case_1",
			GenCmdFlags{},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "mevboost",
				args: []string{},
			},
			false,
		),
		*buildGenerateTestCase(
			t,
			"MevBoost wrong argument", "case_1",
			GenCmdFlags{},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "mevboost",
				args: []string{"wrong"},
			},
			true,
		),
		*buildGenerateTestCase(
			t,
			"Execution ", "case_1",
			GenCmdFlags{
				mapAllPorts: true,
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "sepolia",
				logging:        "",
			},
			subCmd{
				name: "execution",
				args: []string{"nethermind"},
			},
			false,
		),
		*buildGenerateTestCase(
			t,
			"Execution wrong client on gnosis", "case_1",
			GenCmdFlags{},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "gnosis",
				logging:        "",
			},
			subCmd{
				name: "execution",
				args: []string{"geth"},
			},
			true,
		),
		*buildGenerateTestCase(
			t,
			"Mainnet Network, custom ttd, should fail", "case_1",
			GenCmdFlags{
				CustomFlags: CustomFlags{
					customTTD: "some",
				},
			},
			globalFlags{
				network: "mainnet",
			},
			subCmd{
				name: "full-node",
			},
			true),
		*buildGenerateTestCase(
			t,
			"Custom Network and custom ttd, should work", "case_1",
			GenCmdFlags{
				feeRecipient: "0x0000000000000000000000000000000000000000",
				CustomFlags: CustomFlags{
					customTTD: "some",
				},
			},
			globalFlags{
				network: "custom",
			},
			subCmd{
				name: "full-node",
			},
			false),
		*buildGenerateTestCase(
			t,
			"Custom Network and custom ttd, execution node, should work", "case_1",
			GenCmdFlags{
				CustomFlags: CustomFlags{
					customTTD: "some",
				},
			},
			globalFlags{
				network: "custom",
			},
			subCmd{
				name: "execution",
			},
			false),
		*buildGenerateTestCase(
			t,
			"Mainnet Network custom ChainSpec, execution node, shouldn't work", "case_1",
			GenCmdFlags{
				CustomFlags: CustomFlags{
					customTTD: "some",
				},
			},
			globalFlags{
				network: "mainnet",
			},
			subCmd{
				name: "execution",
			},
			true),
		*buildGenerateTestCase(
			t,
			"Validator", "case_1",
			GenCmdFlags{
				feeRecipient: "0x0000000000000000000000000000000000000000",
				CustomFlags: CustomFlags{
					customTTD: "some",
				},
			},
			globalFlags{
				network: "custom",
			},
			subCmd{
				name: "full-node",
			},
			false),
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			descr := fmt.Sprintf("sedge generate %s %s %s", tc.subCommand.argsList(), tc.args.toString(), tc.globalArgs.argsList())

			sedgeActions := actions.NewSedgeActions(nil, nil, nil)

			rootCmd := RootCmd()
			rootCmd.AddCommand(GenerateCmd(sedgeActions))
			argsL := append([]string{"generate"}, tc.subCommand.argsList()...)
			argsL = append(argsL, tc.args.argsList()...)
			argsL = append(argsL, tc.globalArgs.argsList()...)
			rootCmd.SetArgs(argsL)

			if err := rootCmd.Execute(); !tc.isErr && err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if tc.isErr && err == nil {
				t.Errorf("%s expected to fail", descr)
			}
		})
	}
}
