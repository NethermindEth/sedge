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

		combinedClients, err := validateClients(clientsMap)
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
		err = utils.GenerateScripts(combinedClients.Execution.Name, combinedClients.Consensus.Name, combinedClients.Validator.Name, generationPath)
		if err != nil {
			log.Fatal(err)
		}

		// Run docker-compose scripts
		err = utils.RunDockerCompose(generationPath + "/docker-compose.yml")
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// Local flags
	cliCmd.Flags().StringVarP(&executionName, "execution", "e", "", "Execution engine client, e.g. Geth, Nethermind, Besu, Erigon")

	cliCmd.Flags().StringVarP(&consensusName, "consensus", "c", "", "Consensus engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")

	cliCmd.Flags().StringVarP(&validatorName, "validator", "v", "", "Validator engine client, e.g. Teku, Lodestar, Prysm, Lighthouse, Nimbus")

	cliCmd.Flags().StringVarP(&generationPath, "path", "p", configs.DefaultDockerComposeScriptsPath, "docker-compose scripts generation path")

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
		err = installOrShowInstructions(pending)
		return
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

func randomizeClients(allClients clients.OrderedClients) (clients.Clients, error) {
	var executionClient, consensusClient, validatorClient clients.Client
	var combinedClients clients.Clients

	executionClient, err := clients.RandomChoice(allClients[execution])
	if err != nil {
		return combinedClients, err
	}
	consensusClient, err = clients.RandomChoice(allClients[consensus])
	if err != nil {
		return combinedClients, err
	}
	validatorClient, err = clients.RandomChoice(allClients[validator])
	if err != nil {
		return combinedClients, err
	}

	combinedClients = clients.Clients{
		Execution: executionClient,
		Consensus: consensusClient,
		Validator: validatorClient}
	return combinedClients, nil
}

func validateClients(allClients clients.OrderedClients) (clients.Clients, error) {
	var combinedClients clients.Clients
	var err error

	if randomize {
		// Select a random execution client, consensus client and validator client
		combinedClients, err = randomizeClients(allClients)
		if err != nil {
			return combinedClients, err
		}

		log.Infof("Listing randomized clients\n\n")
		ui.WriteRandomizedClientsTable([][]string{
			{"Execution", combinedClients.Execution.Name},
			{"Consensus", combinedClients.Consensus.Name},
			{"Validator", combinedClients.Validator.Name},
		})
	} else {
		notProvidedClients := make([]string, 0)
		if executionName == "" {
			notProvidedClients = append(notProvidedClients, execution+" client")
		}
		if consensusName == "" {
			notProvidedClients = append(notProvidedClients, consensus+" client")
		}
		if validatorName == "" {
			notProvidedClients = append(notProvidedClients, validator+" client")
		}

		if len(notProvidedClients) > 0 {
			var msg string
			if len(notProvidedClients) == 1 {
				msg = notProvidedClients[0]
			} else {
				msg = strings.Join(notProvidedClients[:len(notProvidedClients)-1], ", ")
				msg = msg + " and " + notProvidedClients[len(notProvidedClients)-1]
			}
			return combinedClients, fmt.Errorf(configs.ClientNotSpecifiedError, msg)
		}

		combinedClients = clients.Clients{
			Execution: allClients[execution][executionName],
			Consensus: allClients[consensus][consensusName],
			Validator: allClients[validator][validatorName],
		}

		err = clients.ValidateClient(combinedClients.Execution, execution)
		if err != nil {
			return combinedClients, err
		}
		err = clients.ValidateClient(combinedClients.Consensus, consensus)
		if err != nil {
			return combinedClients, err
		}
		err = clients.ValidateClient(combinedClients.Validator, validator)
		if err != nil {
			return combinedClients, err
		}
	}

	return combinedClients, nil
}
