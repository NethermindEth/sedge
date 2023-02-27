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
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/test"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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
	subCommand subCmd
	args       GenCmdFlags
	globalArgs globalFlags
	err        error
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
	tErr error,
) *generateCmdTestCase {
	tc := generateCmdTestCase{}
	configPath := t.TempDir()

	dcPath := filepath.Join(configPath, "docker-compose-scripts")
	err := os.Mkdir(dcPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	tc.name = name
	tc.args = args
	tc.globalArgs = globalArgs
	tc.globalArgs.generationPath = t.TempDir()
	tc.subCommand = subCommand
	tc.err = tErr
	return &tc
}

func TestGenerateCmd(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)
	configs.InitNetworksConfigs()
	tcs := []generateCmdTestCase{
		*buildGenerateTestCase(
			t,
			"Execution, bad number of arguments", "case_1",
			GenCmdFlags{},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "execution",
				args: []string{"nethermind", "besu"},
			},
			errors.New("requires one argument"),
		),
		*buildGenerateTestCase(
			t,
			"Consensus, bad number of arguments", "case_1",
			GenCmdFlags{
				executionAuthUrl: "http://localhost:8545",
				executionApiUrl:  "http://localhost:8545",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "consensus",
				args: []string{"teku", "lodestar"},
			},
			errors.New("requires one argument"),
		),
		*buildGenerateTestCase(
			t,
			"Validator, bad number of arguments", "case_1",
			GenCmdFlags{
				consensusApiUrl: "http://localhost:4000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "",
				logging:        "",
			},
			subCmd{
				name: "validator",
				args: []string{"teku", "lodestar"},
			},
			errors.New("requires one argument"),
		),
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
				logging: "",
			},
			subCmd{
				name: "full-node",
				args: []string{},
			},
			nil,
		),
		*buildGenerateTestCase(
			t,
			"full-node Random clients, no feeRecipient", "case_1",
			GenCmdFlags{},
			globalFlags{
				install: false,
				logging: "",
			},
			subCmd{
				name: "full-node",
				args: []string{},
			},
			nil,
		),
		*buildGenerateTestCase(
			t,
			"full-node Random clients, custom Ckpt sync endpoint", "case_1",
			GenCmdFlags{
				checkpointSyncUrl: "http://localhost:8080",
			},
			globalFlags{
				install: false,
				logging: "",
			},
			subCmd{
				name: "full-node",
				args: []string{},
			},
			nil,
		),
		*buildGenerateTestCase(
			t,
			"consensus Random client, custom Ckpt sync endpoint", "case_1",
			GenCmdFlags{
				executionAuthUrl:  "http://localhost:8545",
				executionApiUrl:   "http://localhost:8545",
				checkpointSyncUrl: "http://localhost:8080",
			},
			globalFlags{
				install: false,
				logging: "",
			},
			subCmd{
				name: "consensus",
				args: []string{},
			},
			nil,
		),
		*buildGenerateTestCase(
			t,
			"consensus Random client, bad custom Ckpt sync endpoint", "case_1",
			GenCmdFlags{
				executionAuthUrl:  "http://localhost:8545",
				executionApiUrl:   "http://localhost:8545",
				checkpointSyncUrl: "./8080",
			},
			globalFlags{
				install: false,
				logging: "",
			},
			subCmd{
				name: "consensus",
				args: []string{},
			},
			fmt.Errorf(configs.InvalidCkptSyncURL, "./8080"),
		),
		*buildGenerateTestCase(
			t,
			"consensus Random client, invalid api url", "case_1",
			GenCmdFlags{
				executionAuthUrl: "http://localhost:8545",
				executionApiUrl:  "localhost/8545",
			},
			globalFlags{
				install: false,
				logging: "",
			},
			subCmd{
				name: "consensus",
				args: []string{},
			},
			fmt.Errorf("invalid execution api url %s", "localhost/8545"),
		),
		*buildGenerateTestCase(
			t,
			"consensus Random client, invalid auth url", "case_1",
			GenCmdFlags{
				executionAuthUrl: "localhost/8545",
				executionApiUrl:  "http://localhost:8545",
			},
			globalFlags{
				install: false,
				logging: "",
			},
			subCmd{
				name: "consensus",
				args: []string{},
			},
			fmt.Errorf("invalid execution auth url %s", "localhost/8545"),
		),
		*buildGenerateTestCase(
			t,
			"Wrong sub command", "case_1",
			GenCmdFlags{},
			globalFlags{
				install: false,
				network: "",
				logging: "",
			},
			subCmd{
				name: "full",
				args: []string{},
			},
			nil,
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
			nil,
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
			nil,
		),
		*buildGenerateTestCase(
			t,
			"Bad network input", "case_1",
			GenCmdFlags{
				feeRecipient: "0x0000000000000000000000000000000000000000",
			},
			globalFlags{
				install:        false,
				generationPath: "",
				network:        "wrong",
				logging:        "",
			},
			subCmd{
				name: "consensus",
				args: []string{"lighthouse"},
			},
			errors.New("unknown network \"wrong\". Please provide correct network name. Use 'networks' command to see the list of supported networks"),
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
			nil,
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
			errors.New("required flag(s) \"execution-auth-url\" not set"),
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
			errors.New("required flag(s) \"execution-api-url\" not set"),
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
			errors.New("invalid consensus client"),
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
			nil,
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
			errors.New("required flag(s) \"consensus-url\" not set"),
		),
		*buildGenerateTestCase(
			t,
			"Validator good client", "case_1",
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
			nil,
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
			nil,
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
			errors.New("unknown command \"wrong\" for \"sedge generate mevboost\""),
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
			nil,
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
			errors.New("invalid execution client"),
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
			errors.New("custom flags used without --network custom")),
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
			nil),
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
			nil),
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
			errors.New("custom flags used without --network custom")),
		*buildGenerateTestCase(
			t,
			"Full-node, custom TTD", "case_1",
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
			nil),
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
			nil,
		),
		*buildGenerateTestCase(
			t,
			"Full-node, waitEpoch set", "case_1",
			GenCmdFlags{
				feeRecipient: "0x0000000000000000000000000000000000000000",
				waitEpoch:    5,
			},
			globalFlags{
				network: "chiado",
			},
			subCmd{
				name: "full-node",
			},
			nil),
		*buildGenerateTestCase(
			t,
			"Validator, waitEpoch set", "case_1",
			GenCmdFlags{
				feeRecipient:    "0x0000000000000000000000000000000000000000",
				waitEpoch:       50,
				consensusApiUrl: "http://localhost:4000",
			},
			globalFlags{
				network: "goerli",
			},
			subCmd{
				name: "validator",
				args: []string{"lodestar"},
			},
			nil),
		*buildGenerateTestCase(
			t,
			"Validator, invalid consensus api url", "case_1",
			GenCmdFlags{
				feeRecipient:    "0x0000000000000000000000000000000000000000",
				consensusApiUrl: "localhost/4000",
			},
			globalFlags{
				network: "goerli",
			},
			subCmd{
				name: "validator",
				args: []string{"lodestar"},
			},
			errors.New("invalid consensus api url localhost/4000")),
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			descr := fmt.Sprintf("sedge generate %s %s %s", strings.Join(tc.subCommand.argsList(), " "), tc.args.toString(), strings.Join(tc.globalArgs.argsList(), " "))
			sedgeActions := actions.NewSedgeActions(actions.SedgeActionsOptions{})

			rootCmd := RootCmd()
			rootCmd.AddCommand(GenerateCmd(sedgeActions))
			argsL := append([]string{"generate"}, tc.subCommand.argsList()...)
			argsL = append(argsL, tc.args.argsList()...)
			argsL = append(argsL, tc.globalArgs.argsList()...)
			rootCmd.SetArgs(argsL)
			rootCmd.SetOutput(io.Discard)

			err := rootCmd.Execute()

			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error(), descr)
			} else {
				assert.NoError(t, err, descr)
			}
		})
	}
}

func TestGeneratePathCases(t *testing.T) {
	configs.InitNetworksConfigs()
	// Silence logger
	log.SetOutput(io.Discard)

	// Custom Generation path
	path := t.TempDir()
	descr := fmt.Sprintf("Generation path error, sedge generate execution --path %s", path)
	sedgeActions := actions.NewSedgeActions(actions.SedgeActionsOptions{})

	rootCmd := RootCmd()
	rootCmd.AddCommand(GenerateCmd(sedgeActions))
	argsL := []string{"generate", "execution", "--path", path}
	rootCmd.SetArgs(argsL)
	rootCmd.SetOutput(io.Discard)

	err := rootCmd.Execute()

	assert.NoError(t, err, descr)

	// Init generation path
	path = t.TempDir()
	if err = os.Remove(path); err != nil {
		t.Fatal(err)
	}
	descr = fmt.Sprintf("Init generation path, sedge generate execution --path %s", path)
	sedgeActions = actions.NewSedgeActions(actions.SedgeActionsOptions{})

	rootCmd = RootCmd()
	rootCmd.AddCommand(GenerateCmd(sedgeActions))
	argsL = []string{"generate", "execution", "--path", path}
	rootCmd.SetArgs(argsL)
	rootCmd.SetOutput(io.Discard)

	err = rootCmd.Execute()

	assert.NoError(t, err, descr)

	// Custom jwt secret path, good
	path = t.TempDir()
	descr = fmt.Sprintf("Custom jwt secret path, good, sedge generate execution --jwt-secret-path %s", path)
	err = test.PrepareTestCaseDir(filepath.Join("testdata", "cli_tests", "jwtsecret"), path)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}
	jwtPath := filepath.Join(path, "jwtsecret")

	sedgeActions = actions.NewSedgeActions(actions.SedgeActionsOptions{})

	rootCmd = RootCmd()
	rootCmd.AddCommand(GenerateCmd(sedgeActions))
	argsL = []string{"generate", "execution", "--path", path, "--jwt-secret-path", jwtPath}
	rootCmd.SetArgs(argsL)
	rootCmd.SetOutput(io.Discard)

	err = rootCmd.Execute()

	assert.NoError(t, err, descr)

	// Custom jwt secret path, error
	path = t.TempDir()
	descr = fmt.Sprintf("Custom jwt secret path, error, sedge generate execution --jwt-secret-path %s", path)
	sedgeActions = actions.NewSedgeActions(actions.SedgeActionsOptions{})

	rootCmd = RootCmd()
	rootCmd.AddCommand(GenerateCmd(sedgeActions))
	argsL = []string{"generate", "execution", "--jwt-secret-path", path}
	rootCmd.SetArgs(argsL)
	rootCmd.SetOutput(io.Discard)

	err = rootCmd.Execute()

	assert.Error(t, err, descr)
}
