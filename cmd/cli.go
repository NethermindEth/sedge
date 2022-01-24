/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/ui"
	"github.com/NethermindEth/1Click/internal/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	executionClient string
	consensusClient string
	validatorClient string
	generationPath  string
	randomize       bool
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Quick start 1Click",
	Long: `Run the setup tool on-premise in a quick way. Provide only the command line
options and the tool will do all the work.

First it will check if dependencies like docker and docker-compose are installed on your machine
and provide instructions for installing them if they are not installed.

Second, it will generate docker-compose scripts to run the full setup according to your selection.

Finally, it will run the generated docker-compose script`,
	Run: func(cmd *cobra.Command, args []string) {
		executionClients, consensusClients, validatorClients := configs.GetClients("executionClients"), configs.GetClients("consensusClients"), configs.GetClients("validatorClients")
		log.Debugf("Execution clients: %v", executionClients)
		log.Debugf("Consensus clients: %v", consensusClients)
		log.Debugf("Validator clients: %v", validatorClients)
		if executionClients == nil || consensusClients == nil || validatorClients == nil {
			log.Fatal(configs.NoClientsFound)
		}

		if randomize {
			// Select a random execution client and a random consensus client
			executionClient, consensusClient, validatorClient = utils.RandomChoice(executionClients), utils.RandomChoice(consensusClients), utils.RandomChoice(validatorClients)
			log.Infof("Listing randomized clients\n\n")
			ui.WriteRandomizedClientsTable([][]string{{"Execution client", executionClient}, {"Consensus client", consensusClient}, {"Validator client", validatorClient}})
		} else if executionClient == "" || consensusClient == "" || validatorClient == "" {
			log.Error(configs.ProvideClients)
			os.Exit(1)
		} else {
			// Validate clients
			if !utils.Contains(executionClients, executionClient) {
				log.Errorf(configs.IncorrectClient, executionClient)
				os.Exit(1)
			} else if !utils.Contains(consensusClients, consensusClient) {
				log.Errorf(configs.IncorrectClient, consensusClient)
				os.Exit(1)
			} else if !utils.Contains(validatorClients, validatorClient) {
				log.Errorf(configs.IncorrectClient, validatorClient)
				os.Exit(1)
			}
		}

		// Check if dependencies are installed
		dependencies := configs.GetDependencies()
		log.Infof(configs.CheckingDependencies, strings.Join(dependencies, ", "))
		pending := utils.CheckDependencies(dependencies)

		if len(pending) > 0 {
			log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
			// Let the user decide to see the instructions for installing dependencies and exit or let the tool install them and continue
			installOrShowInstructions(pending)
		}
		log.Info(configs.DependenciesOK)

		// Generate docker-compose scripts
		utils.GenerateScripts(executionClient, consensusClient, validatorClient, generationPath)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// Local flags
	cliCmd.Flags().StringVar(&executionClient, "execution", "", "Execution engine client, e.g. Geth, Nethermind, Besu, Erigon")

	cliCmd.Flags().StringVar(&consensusClient, "consensus", "", "Consensus engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")

	cliCmd.Flags().StringVar(&validatorClient, "validator", "", "Validator engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")

	cliCmd.Flags().StringVar(&generationPath, "path", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")

	cliCmd.Flags().BoolVarP(&randomize, "randomize", "r", false, "Randomize combination of clients")
}

func installOrShowInstructions(pending []string) {
	optShow, optInstall, optExit := "Show instructions for installing dependencies", "Install dependencies", "Exit. You will manage this dependencies on your own"
	prompt := promptui.Select{
		Label: "Select how to proceed with the pending dependencies",
		Items: []string{optShow, optInstall, optExit},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("prompt failed %v", err)
	}

	switch result {
	case optShow:
		err = utils.HandleInstructions(pending, utils.ShowInstructions)
		if err != nil {
			log.Fatalf(configs.ShowingInstructionsError, err)
		}
	case optInstall:
		err = utils.HandleInstructions(pending, utils.InstallDependency)
		if err != nil {
			log.Fatalf(configs.InstallingDependenciesError, err)
		}
	default:
		log.Info(configs.Exiting)
		os.Exit(0)
	}
}
