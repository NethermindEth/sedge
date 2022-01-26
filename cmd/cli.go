/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/pkg/clients"
	"github.com/NethermindEth/1Click/internal/ui"
	"github.com/NethermindEth/1Click/internal/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	executionName  string
	consensusName  string
	validatorName  string
	generationPath string
	randomize      bool
)

const (
	execution, consensus, validator = "execution", "consensus", "validator"
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
		// Get all clients: supported + configured
		clientsMap, errors := clients.GetClients([]string{execution, consensus, validator})
		if len(errors) > 0 {
			for _, err := range errors {
				log.Error(err)
			}
			os.Exit(1)
		}

		executionClient, consensusClient, validatorClient, err := validateClients(clientsMap)
		if err != nil {
			log.Fatal(err)
		}

		// Check if dependencies are installed
		dependencies := configs.GetDependencies()
		log.Infof(configs.CheckingDependencies, strings.Join(dependencies, ", "))
		pending := utils.CheckDependencies(dependencies)

		if len(pending) > 0 {
			log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
			// Let the user decide to see the instructions for installing dependencies and exit or let the tool install them and continue
			err := installOrShowInstructions(pending)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Info(configs.DependenciesOK)

		// Generate docker-compose scripts
		err = utils.GenerateScripts(executionClient.Name, consensusClient.Name, validatorClient.Name, generationPath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// Local flags
	cliCmd.Flags().StringVar(&executionName, "execution", "", "Execution engine client, e.g. Geth, Nethermind, Besu, Erigon")

	cliCmd.Flags().StringVar(&consensusName, "consensus", "", "Consensus engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")

	cliCmd.Flags().StringVar(&validatorName, "validator", "", "Validator engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")

	cliCmd.Flags().StringVar(&generationPath, "path", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")

	cliCmd.Flags().BoolVarP(&randomize, "randomize", "r", false, "Randomize combination of clients")
}

func installOrShowInstructions(pending []string) (err error) {
	optShow, optInstall, optExit := "Show instructions for installing dependencies", "Install dependencies", "Exit. You will manage this dependencies on your own"
	prompt := promptui.Select{
		Label: "Select how to proceed with the pending dependencies",
		Items: []string{optShow, optInstall, optExit},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed %s", err)
	}

	switch result {
	case optShow:
		err = utils.HandleInstructions(pending, utils.ShowInstructions)
		if err != nil {
			return fmt.Errorf(configs.ShowingInstructionsError, err)
		}
	case optInstall:
		err = utils.HandleInstructions(pending, utils.InstallDependency)
		if err != nil {
			return fmt.Errorf(configs.InstallingDependenciesError, err)
		}
	default:
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	return nil
}

func randomizeClients(clientsMap map[string][]clients.Client) (clients.Client, clients.Client, clients.Client, error) {
	var executionClient, consensusClient, validatorClient clients.Client

	executionClient, err := clients.RandomChoice(clientsMap[execution])
	if err != nil {
		return executionClient, consensusClient, validatorClient, err
	}
	consensusClient, err = clients.RandomChoice(clientsMap[consensus])
	if err != nil {
		return executionClient, consensusClient, validatorClient, err
	}
	validatorClient, err = clients.RandomChoice(clientsMap[validator])
	if err != nil {
		return executionClient, consensusClient, validatorClient, err
	}

	return executionClient, consensusClient, validatorClient, nil
}

func validateClients(clientsMap map[string][]clients.Client) (clients.Client, clients.Client, clients.Client, error) {
	var executionClient, consensusClient, validatorClient clients.Client
	var err error

	if randomize {
		// Select a random execution client and a random consensus client
		executionClient, consensusClient, validatorClient, err = randomizeClients(clientsMap)
		if err != nil {
			return executionClient, consensusClient, validatorClient, err
		}

		log.Infof("Listing randomized clients\n\n")
		ui.WriteRandomizedClientsTable([][]string{{"Execution client", executionClient.Name}, {"Consensus client", consensusClient.Name}, {"Validator client", validatorClient.Name}})
	} else {
		if executionName == "" {
			return executionClient, consensusClient, validatorClient, fmt.Errorf(configs.ClientNotSpecifiedError, execution)
		}
		if consensusName == "" {
			return executionClient, consensusClient, validatorClient, fmt.Errorf(configs.ClientNotSpecifiedError, consensus)
		}
		if validatorName == "" {
			return executionClient, consensusClient, validatorClient, fmt.Errorf(configs.ClientNotSpecifiedError, validator)
		}

		executionClient, consensusClient, validatorClient = clients.Select(clientsMap[execution], executionName), clients.Select(clientsMap[consensus], consensusName), clients.Select(clientsMap[validator], validatorName)
	}

	err = clients.ValidateClient(executionClient)
	if err != nil {
		return executionClient, consensusClient, validatorClient, err
	}
	err = clients.ValidateClient(consensusClient)
	if err != nil {
		return executionClient, consensusClient, validatorClient, err
	}
	err = clients.ValidateClient(validatorClient)
	if err != nil {
		return executionClient, consensusClient, validatorClient, err
	}

	return executionClient, consensusClient, validatorClient, nil
}
