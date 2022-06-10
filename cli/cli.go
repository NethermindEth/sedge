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
	"os"
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
	"github.com/spf13/pflag"
)

var (
	executionName  string
	consensusName  string
	validatorName  string
	generationPath string
	randomize      bool
	install        bool
	run            bool
	y              bool
	services       *[]string
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

First it will check if dependencies like docker and docker-compose are installed on your machine
and provide instructions for installing them if they are not installed.

Second, it will generate docker-compose scripts to run the full setup according to your selection.

Finally, it will run the generated docker-compose script. Only execution and consensus clients will be executed by default.

Running the command without flags (except global flag'--config') is equivalent to '1click cli -r' `,
	Args: cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := preRunCliCmd(cmd, args); err != nil {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if errs := runCliCmd(cmd, args); len(errs) > 0 {
			for _, err := range errs {
				log.Error(err)
			}
			os.Exit(1)
		}
	},
}

func preRunCliCmd(cmd *cobra.Command, args []string) error {
	// Count flags being set
	count := 0
	// HACKME: LocalFlags() doesn't work, so we count manually and check for parent flag config
	cmd.Flags().Visit(func(f *pflag.Flag) {
		if f.Name != "config" {
			count++
		}
	})

	if count == 0 {
		// No flag behavior
		randomize = true
	}

	// Quick run
	if y {
		randomize, install, run = true, true, true
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
	} else if !utils.ContainsOnly(*services, []string{execution, consensus, validator}) {
		return fmt.Errorf(configs.RunClientsError, strings.Join(*services, ","), strings.Join([]string{execution, consensus, validator}, ","))
	}
	return nil
}

func runCliCmd(cmd *cobra.Command, args []string) []error {
	// Get all clients: supported + configured
	clientsMap, errors := clients.GetClients([]string{execution, consensus, validator})
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

	// Generate docker-compose scripts
	if err = generate.GenerateScripts(combinedClients.Execution.Name, combinedClients.Consensus.Name, combinedClients.Validator.Name, generationPath); err != nil {
		return []error{err}
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
		// Track sync of execution and consensus clients
		// TODO: Parameterize wait arg of trackSync
		if err = trackSync(monitor, time.Minute); err != nil {
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

	// Local flags
	cliCmd.Flags().StringVarP(&executionName, "execution", "e", "", "Execution engine client, e.g. Geth, Nethermind, Besu, Erigon")

	cliCmd.Flags().StringVarP(&consensusName, "consensus", "c", "", "Consensus engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")

	cliCmd.Flags().StringVarP(&validatorName, "validator", "v", "", "Validator engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")

	cliCmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")

	cliCmd.Flags().BoolVarP(&randomize, "randomize", "r", false, "Randomize combination of clients")

	cliCmd.Flags().BoolVarP(&install, "install", "i", false, "Install dependencies if not installed without asking")

	cliCmd.Flags().BoolVar(&run, "run", false, "Run the generated docker-compose scripts without asking")

	cliCmd.Flags().BoolVarP(&y, "yes", "y", false, "Shortcut for '1click cli -r -i --run'. Run without prompts")

	services = cliCmd.Flags().StringSlice("run-clients", []string{execution, consensus}, "Run only the specified clients. Possible values: execution, consensus, validator, all. The 'all' option must be used alone. Example: '1click cli -r --run-clients=consensus,validator'")

}
