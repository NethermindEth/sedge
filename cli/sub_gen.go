package cli

import (
	"errors"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/cli/prompts"
	"github.com/spf13/cobra"
)

func FullNodeSubCmd(prompt prompts.Prompt, sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "full-node [flags]",
		Short: "Generate a full node config, with or without a validator",
		Args:  cobra.OnlyValidArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {

			return preValidationGenerateCmd(&flags)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, prompt, sedgeAction, []string{execution, consensus, validator})
		},
	}
	// Bind flags
	cmd.Flags().StringVarP(&flags.consensusName, "consensus", "c", "", "Consensus engine client, e.g. teku, lodestar, prysm, lighthouse, Nimbus. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVarP(&flags.executionName, "execution", "e", "", "Execution engine client, e.g. geth, nethermind, besu, erigon. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVarP(&flags.validatorName, "validator", "v", "", "Validator engine client, e.g. teku, lodestar, prysm, lighthouse, Nimbus. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")
	cmd.Flags().StringVar(&flags.checkpointSyncUrl, "checkpoint-sync-url", "", "Initial state endpoint (trusted synced consensus endpoint) for the consensus client to sync from a finalized checkpoint. Provide faster sync process for the consensus client and protect it from long-range attacks affored by Weak Subjetivity")
	cmd.Flags().StringVar(&flags.feeRecipient, "fee-recipient", "", "Suggested fee recipient. Is a 20-byte Ethereum address which the execution layer might choose to set as the coinbase and the recipient of other fees or rewards. There is no guarantee that an execution node will use the suggested fee recipient to collect fees, it may use any address it chooses. It is assumed that an honest execution node will use the suggested fee recipient, but users should note this trust assumption")
	cmd.Flags().BoolVar(&flags.noMev, "no-mev-boost", false, "Not use mev-boost if supported")
	cmd.Flags().StringVarP(&flags.mevImage, "mev-boost-image", "m", "", "Custom docker image to use for Mev Boost. Example: 'sedge generate full-node --mev-boost-image flashbots/mev-boost:latest-portable'")
	cmd.Flags().BoolVar(&flags.noValidator, "no-validator", false, "Exclude the validator from the full node setup. Designed for execution and consensus nodes setup without a validator node. Exclude also the validator from other flags. If set, mev-boost will not be used.")
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().StringVar(&flags.graffiti, "graffiti", "", "Graffiti to be used by the validator")
	cmd.Flags().BoolVar(&flags.mapAllPorts, "map-all", false, "Map all clients ports to host. Use with care. Useful to allow remote access to the clients")
	flags.fallbackEL = *cmd.Flags().StringSlice("fallback-execution-urls", []string{}, "Fallback/backup execution endpoints for the consensus client. Not supported by Teku. Example: 'sedge generate full-node -r --fallback-execution=https://mainnet.infura.io/v3/YOUR-PROJECT-ID,https://eth-mainnet.alchemyapi.io/v2/YOUR-PROJECT-ID'")
	flags.elExtraFlags = *cmd.Flags().StringArray("el-extra-flag", []string{}, "Additional flag to configure the execution client service in the generated docker-compose script. Example: 'sedge generate full-node --el-extra-flag \"<flag1>=value1\" --el-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	flags.clExtraFlags = *cmd.Flags().StringArray("cl-extra-flag", []string{}, "Additional flag to configure the consensus client service in the generated docker-compose script. Example: 'sedge generate full-node --cl-extra-flag \"<flag1>=value1\" --cl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	flags.vlExtraFlags = *cmd.Flags().StringArray("vl-extra-flag", []string{}, "Additional flag to configure the validator client service in the generated docker-compose script. Example: 'sedge generate full-node --vl-extra-flag \"<flag1>=value1\" --vl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	cmd.Flags().SortFlags = false
	return cmd
}

func ExecutionSubCmd(prompt prompts.Prompt, sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "execution [flags] [args]",
		Short: "Generate a execution node config",
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
			err := preValidationGenerateCmd(&flags)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, prompt, sedgeAction, []string{execution})
		},
	}

	// Bind flags
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().BoolVar(&flags.mapAllPorts, "map-all", false, "Map all clients ports to host. Use with care. Useful to allow remote access to the clients")
	flags.elExtraFlags = *cmd.Flags().StringArray("el-extra-flag", []string{}, "Additional flag to configure the execution client service in the generated docker-compose script. Example: 'sedge generate consensus--el-extra-flag \"<flag1>=value1\" --el-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	cmd.Flags().SortFlags = false
	return cmd
}

func ConsensusSubCmd(prompt prompts.Prompt, sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "consensus [flags] [args]",
		Short: "Generate a consensus node config",
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
			return preValidationGenerateCmd(&flags)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, prompt, sedgeAction, []string{consensus})
		},
	}
	// Bind flags
	cmd.Flags().StringVar(&flags.checkpointSyncUrl, "checkpoint-sync-url", "", "Initial state endpoint (trusted synced consensus endpoint) for the consensus client to sync from a finalized checkpoint. Provide faster sync process for the consensus client and protect it from long-range attacks affored by Weak Subjetivity")
	cmd.Flags().StringVar(&flags.feeRecipient, "fee-recipient", "", "Suggested fee recipient. Is a 20-byte Ethereum address which the execution layer might choose to set as the coinbase and the recipient of other fees or rewards. There is no guarantee that an execution node will use the suggested fee recipient to collect fees, it may use any address it chooses. It is assumed that an honest execution node will use the suggested fee recipient, but users should note this trust assumption")
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().StringVar(&flags.mevBoostUrl, "mev-boost-url", "", "Mev Boost endpoint")
	cmd.Flags().BoolVar(&flags.mapAllPorts, "map-all", false, "Map all clients ports to host. Use with care. Useful to allow remote access to the clients")
	flags.fallbackEL = *cmd.Flags().StringSlice("fallback-execution-urls", []string{}, "Fallback/backup execution endpoints for the consensus client. Not supported by Teku. Example: 'sedge cli -r --fallback-execution=https://mainnet.infura.io/v3/YOUR-PROJECT-ID,https://eth-mainnet.alchemyapi.io/v2/YOUR-PROJECT-ID'")
	flags.clExtraFlags = *cmd.Flags().StringArray("cl-extra-flag", []string{}, "Additional flag to configure the consensus client service in the generated docker-compose script. Example: 'sedge generate consensus --cl-extra-flag \"<flag1>=value1\" --cl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	cmd.Flags().StringVar(&flags.executionApiUrl, "execution-api-url", "", "Execution API endpoint for the consensus client. Example: 'sedge generate consensus -r --execution-api-url=https://mainnet.infura.io/v3/YOUR-PROJECT-ID'")
	cmd.Flags().StringVar(&flags.executionAuthUrl, "execution-auth-url", "", "Execution AUTH endpoint for the consensus client. Example: 'sedge generate consensus -r --execution-auth-url=https://mainnet .infura.io/v3/YOUR-PROJECT-ID'")
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

func ValidatorSubCmd(prompt prompts.Prompt, sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "validator [flags] [args]",
		Short: "Generate a validator node config",
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
			return preValidationGenerateCmd(&flags)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, prompt, sedgeAction, []string{validator})
		},
	}
	// Bind flags
	cmd.Flags().StringVar(&flags.feeRecipient, "fee-recipient", "", "Suggested fee recipient. Is a 20-byte Ethereum address which the execution layer might choose to set as the coinbase and the recipient of other fees or rewards. There is no guarantee that an execution node will use the suggested fee recipient to collect fees, it may use any address it chooses. It is assumed that an honest execution node will use the suggested fee recipient, but users should note this trust assumption")
	cmd.Flags().StringVar(&flags.jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")
	cmd.Flags().StringVar(&flags.graffiti, "graffiti", "", "Graffiti to be used by the validator")
	flags.vlExtraFlags = *cmd.Flags().StringArray("vl-extra-flag", []string{}, "Additional flag to configure the validator client service in the generated docker-compose script. Example: 'sedge generate validator --vl-extra-flag \"<flag1>=value1\" --vl-extra-flag \"<flag2>=\\\"value2\\\"\"'")
	cmd.Flags().StringVar(&flags.consensusApiUrl, "consensus-url", "", "Consensus endpoint for the validator client to connect to. Example: 'sedge generate validator --consensus-url http://localhost:8545'")
	err := cmd.MarkFlagRequired("consensus-url")
	if err != nil {
		return nil
	}
	cmd.Flags().SortFlags = false
	return cmd
}

func MevBoostSubCmd(prompt prompts.Prompt, sedgeAction actions.SedgeActions) *cobra.Command {
	var flags GenCmdFlags

	cmd := &cobra.Command{
		Use:   "mevboost [flags]",
		Short: "Generate a mev-boost node config",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preValidationGenerateCmd(&flags)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenCmd(cmd.OutOrStdout(), &flags, prompt, sedgeAction, []string{mevBoost})
		},
	}
	// Bind flags
	cmd.Flags().StringVar(&flags.relayURL, "relay-url", "", "Relay URL used to connect to mev relay.")
	cmd.Flags().StringVarP(&flags.mevImage, "mev-boost-image", "m", "", "Custom docker image to use for Mev Boost. Example: 'sedge generate mevboost --mev-boost-image flashbots/mev-boost:latest-portable'")
	cmd.Flags().SortFlags = false
	return cmd
}
