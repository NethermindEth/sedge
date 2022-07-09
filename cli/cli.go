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

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/pkg/clients"
	"github.com/NethermindEth/1click/internal/pkg/generate"
	"github.com/NethermindEth/1click/internal/utils"
	posmoni "github.com/NethermindEth/posmoni/pkg/eth2"
	posmonidb "github.com/NethermindEth/posmoni/pkg/eth2/db"
	posmoninet "github.com/NethermindEth/posmoni/pkg/eth2/networking"
	"github.com/spf13/cobra"
)

var (
	executionName     string
	executionImage    string
	consensusName     string
	consensusImage    string
	validatorName     string
	validatorImage    string
	generationPath    string
	checkpointSyncUrl string
	network           string
	feeRecipient      string
	jwtPath           string
	install           bool
	run               bool
	y                 bool
	services          *[]string
	fallbackEL        *[]string
	elExtraFlags      *[]string
	clExtraFlags      *[]string
	vlExtraFlags      *[]string
	mapAllPorts       bool
	noMev             bool
)

const (
	execution, consensus, validator = "execution", "consensus", "validator"
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli [flags]",
	Short: "Quick start 1click",
	Long: `Run the setup tool on-premise in a quick way. Provide only the command line
options and the tool will do all the work.

First it will check if dependencies such as docker are installed on your machine
and provide instructions for installing them if they are not installed.

Second, it will generate docker-compose scripts to run the full setup according to your selection.

Finally, it will run the generated docker-compose script. Only execution and consensus clients will be executed by default.`,
	Args: cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		// notest
		if err := preRunCliCmd(cmd, args); err != nil {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// notest
		if errs := runCliCmd(cmd, args); len(errs) > 0 {
			for _, err := range errs {
				log.Error(err)
			}
			os.Exit(1)
		}
	},
}

func preRunCliCmd(cmd *cobra.Command, args []string) error {
	// Quick run
	if y {
		install, run = true, true
	}

	// Validate run-clients flag
	if utils.Contains(*services, "all") {
		if len(*services) == 1 {
			// all used correctly
			services = &[]string{execution, consensus, validator}
		} else {
			// Ambiguous value
			return fmt.Errorf(configs.RunClientsFlagAmbiguousError, *services)
		}
	} else if utils.Contains(*services, "none") {
		if len(*services) == 1 {
			// all used correctly
			services = &[]string{}
		} else {
			// Ambiguous value
			return fmt.Errorf(configs.RunClientsFlagAmbiguousError, *services)
		}
	} else if !utils.ContainsOnly(*services, []string{execution, consensus, validator}) {
		return fmt.Errorf(configs.RunClientsError, strings.Join(*services, ","), strings.Join([]string{execution, consensus, validator}, ","))
	}

	// Validate network
	networks, err := utils.SupportedNetworks()
	if err != nil {
		return fmt.Errorf(configs.NetworkValidationFailedError, err)
	}
	if !utils.Contains(networks, network) {
		return fmt.Errorf(configs.UnknownNetworkError, network)
	}

	// Validate fee recipient
	if feeRecipient != "" && !utils.IsAddress(feeRecipient) {
		return errors.New(configs.InvalidFeeRecipientError)
	}

	// Prepare custom images
	if executionName != "" {
		executionParts := strings.Split(executionName, ":")
		executionName = executionParts[0]
		executionImage = strings.Join(executionParts[1:], ":")
	}
	if consensusName != "" {
		consensusParts := strings.Split(consensusName, ":")
		consensusName = consensusParts[0]
		consensusImage = strings.Join(consensusParts[1:], ":")
	}
	if validatorName != "" {
		validatorParts := strings.Split(validatorName, ":")
		validatorName = validatorParts[0]
		validatorImage = strings.Join(validatorParts[1:], ":")
	}

	return nil
}

func runCliCmd(cmd *cobra.Command, args []string) []error {
	// Warnings
	// Warn if custom images are used
	if executionImage != "" || consensusImage != "" || validatorImage != "" {
		log.Warn(configs.CustomImagesWarning)
	}
	// Warn if exposed ports are used
	if mapAllPorts {
		log.Warn(configs.MapAllPortsWarning)
	}

	// Get all clients: supported + configured
	c := clients.ClientInfo{Network: network}
	clientsMap, errors := c.Clients([]string{execution, consensus, validator})
	if len(errors) > 0 {
		return errors
	}

	// Handle selection and validation of clients
	combinedClients, err := validateClients(clientsMap, cmd.OutOrStdout())
	if err != nil {
		return []error{err}
	}

	dependencies := configs.GetDependencies()
	log.Infof(configs.CheckingDependencies, strings.Join(dependencies, ", "))

	// Check if dependencies are installed. Keep checking dependencies until they are all installed
	for pending := utils.CheckDependencies(dependencies); len(pending) > 0; pending = utils.CheckDependencies(dependencies) {
		log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
		if install {
			// Install dependencies directly
			if err := installDependencies(pending); err != nil {
				return []error{err}
			}
		} else {
			// Let the user decide to see the instructions for installing dependencies and exit or let the tool install them and continue
			if err := installOrShowInstructions(pending); err != nil {
				return []error{err}
			}
		}
	}
	log.Info(configs.DependenciesOK)

	// Generate JWT secret if necessary
	if jwtPath == "" && configs.JWTNetworks[network] {
		if err = handleJWTSecret(); err != nil {
			return []error{err}
		}
	} else if filepath.IsAbs(jwtPath) { //Ensure jwtPath is absolute
		if jwtPath, err = filepath.Abs(jwtPath); err != nil {
			return []error{err}
		}
	}

	// Get fee recipient
	if !y && feeRecipient == "" {
		if err = feeRecipientPrompt(); err != nil {
			return []error{err}
		}
	}

	// Generate docker-compose scripts
	gd := generate.GenerationData{
		ExecutionClient:   combinedClients.Execution.Name,
		ExecutionImage:    executionImage,
		ConsensusClient:   combinedClients.Consensus.Name,
		ConsensusImage:    consensusImage,
		ValidatorClient:   combinedClients.Validator.Name,
		ValidatorImage:    validatorImage,
		GenerationPath:    generationPath,
		Network:           network,
		CheckpointSyncUrl: checkpointSyncUrl,
		FeeRecipient:      feeRecipient,
		JWTSecretPath:     jwtPath,
		FallbackELUrls:    *fallbackEL,
		ElExtraFlags:      *elExtraFlags,
		ClExtraFlags:      *clExtraFlags,
		VlExtraFlags:      *vlExtraFlags,
		MapAllPorts:       mapAllPorts,
		Mev:               !noMev,
	}
	elPort, clPort, err := generate.GenerateScripts(gd)
	if err != nil {
		return []error{err}
	}

	// If --run-clients=none was set then exit and don't run anything
	if len(*services) == 0 {
		log.Info(configs.HappyStaking2)
		return nil
	}

	// If teku is chosen, then prepare datadir with 777 permissions
	if combinedClients.Consensus.Name == "teku" {
		preRunTeku()
	}

	if run {
		if err = runAndShowContainers(*services); err != nil {
			return []error{err}
		}
	} else {
		// Let the user decide to see the instructions for executing the scripts and exit or let the tool execute them
		if err = runScriptOrExit(); err != nil {
			return []error{err}
		}
	}

	log.Info(configs.ValidatorTips)

	// Run validator after execution and consensus clients are synced, unless the user intencionally wants to run the validator service  in the previous step
	if !utils.Contains(*services, validator) {
		// Wait for clients to start
		//log.Info(configs.WaitingForNodesToStart)
		//time.Sleep(waitingTime)
		// Track sync of execution and consensus clients
		// TODO: Parameterize wait arg of trackSync
		if err = trackSync(monitor, elPort, clPort, time.Minute*5); err != nil {
			return []error{err}
		}

		// TODO: Prompt for waiting for keystore and validator registration to run the validator
		if run {
			if err = runAndShowContainers([]string{validator}); err != nil {
				return []error{err}
			}
		} else {
			// Let the user decide to see the instructions for executing the validator and exit or let the tool execute it
			if err = RunValidatorOrExit(); err != nil {
				return []error{err}
			}
		}
	}
	log.Info(configs.HappyStaking)

	return nil
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().SortFlags = false

	// Local flags
	cliCmd.Flags().StringVarP(&executionName, "execution", "e", "", "Execution engine client, e.g. geth, nethermind, besu, erigon. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")

	cliCmd.Flags().StringVarP(&consensusName, "consensus", "c", "", "Consensus engine client, e.g. teku, lodestar, prysm, lighthouse, Nimbus. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")

	cliCmd.Flags().StringVarP(&validatorName, "validator", "v", "", "Validator engine client, e.g. teku, lodestar, prysm, lighthouse, Nimbus. Additionally, you can use this syntax '<CLIENT>:<DOCKER_IMAGE>' to override the docker image used for the client. If you want to use the default docker image, just use the client name")

	cliCmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")

	cliCmd.Flags().StringVar(&checkpointSyncUrl, "checkpoint-sync-url", "", "Initial state endpoint (trusted synced consensus endpoint) for the consensus client to sync from a finalized checkpoint. Provide faster sync process for the consensus client and protect it from long-range attacks affored by Weak Subjetivity")

	cliCmd.Flags().StringVarP(&network, "network", "n", "mainnet", "Target network. e.g. mainnet, prater, kiln, etc.")

	cliCmd.Flags().StringVar(&feeRecipient, "fee-recipient", "", "Suggested fee recipient. Is 20-byte Ethereum address which the execution layer might choose to set as the coinbase and the recipient of other fees or rewards. There is no guarantee that an execution node will use the suggested fee recipient to collect fees, it may use any address it chooses. It is assumed that an honest execution node will use the suggested fee recipient, but users should note this trust assumption")

	cliCmd.Flags().BoolVar(&noMev, "no-mev-boost", false, "Not use mev-boost if supported")

	cliCmd.Flags().StringVar(&jwtPath, "jwt-secret-path", "", "Path to the JWT secret file")

	cliCmd.Flags().BoolVarP(&install, "install", "i", false, "Install dependencies if not installed without asking")

	cliCmd.Flags().BoolVarP(&run, "run", "r", false, "Run the generated docker-compose scripts without asking")

	cliCmd.Flags().BoolVarP(&y, "yes", "y", false, "Shortcut for '1click cli -r -i --run'. Run without prompts")

	cliCmd.Flags().BoolVar(&mapAllPorts, "map-all", false, "Map all clients ports to host. Use with care. Useful to allow remote access to the clients")

	services = cliCmd.Flags().StringSlice("run-clients", []string{execution, consensus}, "Run only the specified clients. Possible values: execution, consensus, validator, all, none. The 'all' and 'none' option must be used alone. Example: '1click cli -r --run-clients=consensus,validator'")

	fallbackEL = cliCmd.Flags().StringSlice("fallback-execution-urls", []string{}, "Fallback/backup execution endpoints for the consensus client. Not supported by Teku. Example: '1click cli -r --fallback-execution=https://mainnet.infura.io/v3/YOUR-PROJECT-ID,https://eth-mainnet.alchemyapi.io/v2/YOUR-PROJECT-ID'")

	elExtraFlags = cliCmd.Flags().StringArray("el-extra-flag", []string{}, "Additional flag to configure the execution client service in the generated docker-compose script. Example: '1click cli --el-extra-flag \"<flag1>=value1\" --el-extra-flag \"<flag2>=\\\"value2\\\"\"'")

	clExtraFlags = cliCmd.Flags().StringArray("cl-extra-flag", []string{}, "Additional flag to configure the consensus client service in the generated docker-compose script. Example: '1click cli --cl-extra-flag \"<flag1>=value1\" --cl-extra-flag \"<flag2>=\\\"value2\\\"\"'")

	vlExtraFlags = cliCmd.Flags().StringArray("vl-extra-flag", []string{}, "Additional flag to configure the validator client service in the generated docker-compose script. Example: '1click cli --vl-extra-flag \"<flag1>=value1\" --vl-extra-flag \"<flag2>=\\\"value2\\\"\"'")

	// Initialize monitoring tool
	initMonitor(func() MonitoringTool {
		// Initialize Eth2 Monitoring tool
		moniCfg := posmoni.ConfigOpts{
			Checkers: []posmoni.CfgChecker{
				{Key: posmoni.Execution, ErrMsg: posmoni.NoExecutionFoundError, Data: []string{configs.OnPremiseExecutionURL}},
				{Key: posmoni.Consensus, ErrMsg: posmoni.NoConsensusFoundError, Data: []string{configs.OnPremiseConsensusURL}},
			},
		}
		m, err := posmoni.NewEth2Monitor(
			posmonidb.EmptyRepository{},
			&posmoninet.BeaconClient{RetryDuration: time.Minute * 10},
			&posmoninet.ExecutionClient{RetryDuration: time.Minute * 10},
			posmoninet.SubscribeOpts{},
			moniCfg,
		)
		if err != nil {
			log.Fatalf(configs.MonitoringToolInitError, err)
		}

		return m
	})

	waitingTime = time.Minute
}
