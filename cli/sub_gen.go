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

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/spf13/cobra"
)

var CustomFlagsUsedWithoutCustomNetwork = errors.New("custom flags used without --network custom")

func validateCustomNetwork(flags *CustomFlags, net string) error {
	if net != "custom" {
		if len(flags.customTTD) != 0 || len(flags.customChainSpec) != 0 || len(flags.customNetworkConfig) != 0 ||
			len(flags.customGenesis) != 0 || len(flags.customDeployBlock) != 0 {
			// TODO add error on expected place
			return CustomFlagsUsedWithoutCustomNetwork
		}
	}
	return nil
}

func FullNodeSubCmd(sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "full-node [flags]",
		Short: "Generate a full node config, with or without a validator",
		Long: `Generate a docker-compose and an environment file with a full node configuration.

It will not generate a validator configuration if the --no-validator flag is set to true.

On mainnet, sepolia and goerli, mev-boost will be activated by default unless you run it with --no-mev-boost flag.

If you don't provide a execution, consensus or validator client, it will be chosen randomly. If one of consensus or validator is provided, but the other one is omitted, then the same pair of clients will be used for both consensus and validator.

Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client, for example 'sedge generate full-node --execution nethermind:docker.image'. If you want to use the default docker image, just use the client name`,
		Args: cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validateCustomNetwork(&flags.CustomFlags, network); err != nil {
				return err
			}
			return preValidationGenerateCmd(network, logging)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, sedgeAction, []string{execution, consensus, validator})
		},
	}
	// Bind flags
	cmd.Flags().StringVarP(&flags.consensusName, "consensus", "c", "", "Consensus engine client, e.g. teku, lodestar, prysm, lighthouse, Nimbus. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVarP(&flags.executionName, "execution", "e", "", "Execution engine client, e.g. geth, nethermind, besu, erigon. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVarP(&flags.validatorName, "validator", "v", "", "Validator engine client, e.g. teku, lodestar, prysm, lighthouse, Nimbus. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVar(&flags.checkpointSyncUrl, "checkpoint-sync-url", "", "Initial state endpoint (trusted synced consensus endpoint) for the consensus client to sync from a finalized checkpoint. Provide faster sync process for the consensus client and protect it from long-range attacks affored by Weak Subjetivity. Each network has a default checkpoint sync url.")
	cmd.Flags().StringVar(&flags.feeRecipient, "fee-recipient", "", "Suggested fee recipient. Is a 20-byte Ethereum address which the execution layer might choose to set as the coinbase and the recipient of other fees or rewards. There is no guarantee that an execution node will use the suggested fee recipient to collect fees, it may use any address it chooses. It is assumed that an honest execution node will use the suggested fee recipient, but users should note this trust assumption")
	cmd.Flags().BoolVar(&flags.noMev, "no-mev-boost", false, "Not use mev-boost if supported")
	cmd.Flags().StringVarP(&flags.mevImage, "mev-boost-image", "m", "", "Custom docker image to use for Mev Boost. Example: 'sedge generate full-node --mev-boost-image flashbots/mev-boost:latest-portable'")
	cmd.Flags().BoolVar(&flags.noValidator, "no-validator", false, "Exclude the validator from the full node setup. Designed for execution and consensus nodes setup without a validator node. Exclude also the validator from other flags. If set, mev-boost will not be used.")
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().StringVar(&flags.graffiti, "graffiti", "", "Graffiti to be used by the validator")
	cmd.Flags().BoolVar(&flags.mapAllPorts, "map-all", false, "Map all clients ports to host. Use with care. Useful to allow remote access to the clients")
	flags.fallbackEL = cmd.Flags().StringSlice("fallback-execution-urls", []string{}, "Fallback/backup execution endpoints for the consensus client. Not supported by Teku. Example: 'sedge generate full-node -r --fallback-execution=https://mainnet.infura.io/v3/YOUR-PROJECT-ID,https://eth-mainnet.alchemyapi.io/v2/YOUR-PROJECT-ID'")
	flags.elExtraFlags = cmd.Flags().StringArray("el-extra-flag", []string{}, "Additional flag to configure the execution client service in the generated docker-compose script. Example: 'sedge generate full-node --el-extra-flag \"<flag1>=value1\" --el-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	flags.clExtraFlags = cmd.Flags().StringArray("cl-extra-flag", []string{}, "Additional flag to configure the consensus client service in the generated docker-compose script. Example: 'sedge generate full-node --cl-extra-flag \"<flag1>=value1\" --cl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	flags.vlExtraFlags = cmd.Flags().StringArray("vl-extra-flag", []string{}, "Additional flag to configure the validator client service in the generated docker-compose script. Example: 'sedge generate full-node --vl-extra-flag \"<flag1>=value1\" --vl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	cmd.Flags().StringVar(&flags.customTTD, "custom-ttd", "", "Custom Terminal Total Difficulty to use for the execution client")
	cmd.Flags().StringVar(&flags.customChainSpec, "custom-chainSpec", "", "File path or url to use as custom network chainSpec for execution client.")
	cmd.Flags().StringVar(&flags.customNetworkConfig, "custom-config", "", "File path or url to use as custom network config file for consensus client.")
	cmd.Flags().StringVar(&flags.customGenesis, "custom-genesis", "", "File path or url to use as custom network genesis for consensus client.")
	cmd.Flags().StringVar(&flags.customDeployBlock, "custom-deploy-block", "", "Custom network deploy block to use for consensus client.")
	cmd.Flags().IntVar(&flags.waitEpoch, "wait-epoch", 0, "Number of epochs to wait before starting and restarting of the validator client.")
	flags.customEnodes = cmd.Flags().StringSlice("execution-bootnodes", []string{}, "List of comma separated enodes to use as custom network peers for execution client.")
	flags.customEnrs = cmd.Flags().StringSlice("consensus-bootnodes", []string{}, "List of comma separated enrs to use as custom network peers for consensus client.")
	cmd.Flags().SortFlags = false
	return cmd
}

func ExecutionSubCmd(sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "execution [flags] [args]",
		Short: "Generate a execution node config",
		Long: "Generate a docker-compose and an environment file with a execution node configuration.\n" +
			"Valid args: name of execution clients according to network\n\n" +
			"Should be one of: nethermind, geth, besu, erigon. If you don't provide one, it will chosen randomly.\n" +
			"Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client, for example 'sedge generate execution nethermind:docker.image'. If you want to use the default docker image, just use the client name",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				if cobra.ExactArgs(1)(cmd, args) != nil {
					return errors.New("requires one argument")
				}
				flags.executionName = args[0]
			}
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validateCustomNetwork(&flags.CustomFlags, network); err != nil {
				return err
			}
			err := preValidationGenerateCmd(network, logging)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, sedgeAction, []string{execution})
		},
	}

	// Bind flags
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().BoolVar(&flags.mapAllPorts, "map-all", false, "Map all clients ports to host. Use with care. Useful to allow remote access to the clients")
	cmd.Flags().StringVar(&flags.customTTD, "custom-ttd", "", "Custom Terminal Total Difficulty to use for the execution client")
	cmd.Flags().StringVar(&flags.customChainSpec, "custom-chainSpec", "", "File path or url to use as custom network chainSpec for execution client.")
	flags.customEnodes = cmd.Flags().StringSlice("execution-bootnodes", []string{}, "List of comma separated enodes to use as custom network peers for execution client.")
	flags.elExtraFlags = cmd.Flags().StringArray("el-extra-flag", []string{}, "Additional flag to configure the execution client service in the generated docker-compose script. Example: 'sedge generate consensus--el-extra-flag \"<flag1>=value1\" --el-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	cmd.Flags().SortFlags = false
	return cmd
}

func ConsensusSubCmd(sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "consensus [flags] --execution-api-url <URL> --execution-auth-url <URL> [args]",
		Short: "Generate a consensus node config",
		Long: "Generate a docker-compose and an environment file with a consensus node configuration\n" +
			"Valid args: name of execution clients according to network\n\n" +
			"Should be one of: lighthouse, teku, prysm, lodestar. If you don't provide one, it will chosen randomly.\n" +
			"Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client, for example 'sedge generate consensus prysm:docker.image'. If you want to use the default docker image, just use the client name" +
			"\n\n" +
			"Required flags:\n" +
			"- '--execution-api-url'\n" +
			"- '--execution-auth-url'",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				if cobra.ExactArgs(1)(cmd, args) != nil {
					return errors.New("requires one argument")
				}
				flags.consensusName = args[0]
			}
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validateCustomNetwork(&flags.CustomFlags, network); err != nil {
				return err
			}
			return preValidationGenerateCmd(network, logging)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, sedgeAction, []string{consensus})
		},
	}
	// Bind flags
	cmd.Flags().StringVar(&flags.executionApiUrl, "execution-api-url", "", "Execution API endpoint for the consensus client. Example: 'sedge generate consensus -r --execution-api-url=https://api.url.endpoint'")
	cmd.Flags().StringVar(&flags.executionAuthUrl, "execution-auth-url", "", "Execution AUTH endpoint for the consensus client. Example: 'sedge generate consensus -r --execution-auth-url=https://auth.url.endpoint'")
	cmd.Flags().StringVar(&flags.checkpointSyncUrl, "checkpoint-sync-url", "", "Initial state endpoint (trusted synced consensus endpoint) for the consensus client to sync from a finalized checkpoint. Provide faster sync process for the consensus client and protect it from long-range attacks affored by Weak Subjetivity. Each network has a default checkpoint sync url.")
	cmd.Flags().StringVar(&flags.feeRecipient, "fee-recipient", "", "Suggested fee recipient. Is a 20-byte Ethereum address which the execution layer might choose to set as the coinbase and the recipient of other fees or rewards. There is no guarantee that an execution node will use the suggested fee recipient to collect fees, it may use any address it chooses. It is assumed that an honest execution node will use the suggested fee recipient, but users should note this trust assumption")
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().StringVar(&flags.mevBoostUrl, "mev-boost-url", "", "If you are running a mev boost node, and you want to connect to it, you need to set mev-boost-url, if not set, node will not load any mev boost related config.")
	cmd.Flags().BoolVar(&flags.mapAllPorts, "map-all", false, "Map all clients ports to host. Use with care. Useful to allow remote access to the clients")
	flags.fallbackEL = cmd.Flags().StringSlice("fallback-execution-urls", []string{}, "Fallback/backup execution endpoints for the consensus client. Not supported by Teku. Example: 'sedge generate consensus --fallback-execution=https://mainnet.infura.io/v3/YOUR-PROJECT-ID,https://eth-mainnet.alchemyapi.io/v2/YOUR-PROJECT-ID'")
	flags.clExtraFlags = cmd.Flags().StringArray("cl-extra-flag", []string{}, "Additional flag to configure the consensus client service in the generated docker-compose script. Example: 'sedge generate consensus --cl-extra-flag \"<flag1>=value1\" --cl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	cmd.Flags().StringVar(&flags.customNetworkConfig, "custom-config", "", "File path or url to use as custom network config file for consensus client.")
	cmd.Flags().StringVar(&flags.customGenesis, "custom-genesis", "", "File path or url to use as custom network genesis for consensus client.")
	cmd.Flags().StringVar(&flags.customDeployBlock, "custom-deploy-block", "", "Custom network deploy block to use for consensus client.")
	flags.customEnrs = cmd.Flags().StringSlice("consensus-bootnodes", []string{}, "List of comma separated enrs to use as custom network peers for consensus client.")
	err := cmd.MarkFlagRequired("execution-api-url")
	if err != nil {
		return nil
	}
	err = cmd.MarkFlagRequired("execution-auth-url")
	if err != nil {
		return nil
	}
	cmd.Flags().SortFlags = false
	return cmd
}

func ValidatorSubCmd(sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "validator [flags] --consensus-url <URL> [args]",
		Short: "Generate a validator node config",
		Long: "Generate a docker-compose and an environment file with a validator node configuration\n" +
			"Valid args: name of execution clients according to network\n\n" +
			"Should be one of: lighthouse, teku, prysm, lodestar. If you don't provide one, it will chosen randomly.\n" +
			"Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client, for example 'sedge generate validator prysm:docker.image'. If you want to use the default docker image, just use the client name" +
			"\n\n" +
			"Required flags:\n" +
			"- `--consensus-url`",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				if cobra.ExactArgs(1)(cmd, args) != nil {
					return errors.New("requires one argument")
				}
				flags.validatorName = args[0]
			}
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validateCustomNetwork(&flags.CustomFlags, network); err != nil {
				return err
			}
			return preValidationGenerateCmd(network, logging)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, sedgeAction, []string{validator})
		},
	}
	// Bind flags
	cmd.Flags().StringVar(&flags.consensusApiUrl, "consensus-url", "", "Consensus endpoint for the validator client to connect to. Example: 'sedge generate validator --consensus-url http://localhost:4000'")
	cmd.Flags().StringVar(&flags.feeRecipient, "fee-recipient", "", "Suggested fee recipient. Is a 20-byte Ethereum address which the execution layer might choose to set as the coinbase and the recipient of other fees or rewards. There is no guarantee that an execution node will use the suggested fee recipient to collect fees, it may use any address it chooses. It is assumed that an honest execution node will use the suggested fee recipient, but users should note this trust assumption")
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().StringVar(&flags.graffiti, "graffiti", "", "Graffiti to be used by the validator")
	cmd.Flags().BoolVar(&flags.mevBoostOnVal, "mev-boost", false, "Use mev-boost while turning on validator node")
	cmd.Flags().StringVar(&flags.customNetworkConfig, "custom-config", "", "File path or url to use as custom network config file for consensus client.")
	cmd.Flags().StringVar(&flags.customGenesis, "custom-genesis", "", "File path or url to use as custom network genesis for consensus client.")
	cmd.Flags().StringVar(&flags.customDeployBlock, "custom-deploy-block", "", "Custom network deploy block to use for consensus client.")
	cmd.Flags().IntVar(&flags.waitEpoch, "wait-epoch", 0, "Number of epochs to wait before starting and restarting of the validator client.")
	flags.vlExtraFlags = cmd.Flags().StringArray("vl-extra-flag", []string{}, "Additional flag to configure the validator client service in the generated docker-compose script. Example: 'sedge generate validator --vl-extra-flag \"<flag1>=value1\" --vl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	err := cmd.MarkFlagRequired("consensus-url")
	if err != nil {
		return nil
	}
	cmd.Flags().SortFlags = false
	return cmd
}

func MevBoostSubCmd(sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "mevboost [flags]",
		Short: "Generate a mev-boost node config",
		Long:  "Generate a docker-compose and an environment file with a mev-boost node configuration",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preValidationGenerateCmd(network, logging)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, sedgeAction, []string{mevBoost})
		},
	}
	// Bind flags
	cmd.Flags().StringVar(&flags.relayURL, "relay-url", "", "Relay URL used to connect to mev relay.")
	cmd.Flags().StringVarP(&flags.mevImage, "mev-boost-image", "m", "", "Custom docker image to use for Mev Boost. Example: 'sedge generate mevboost --mev-boost-image flashbots/mev-boost:latest-portable'")
	cmd.Flags().SortFlags = false
	return cmd
}
