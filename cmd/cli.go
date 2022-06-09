/*
Copyright © 2022 Nethermind hello.nethermind.io

*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/pkg/clients"
	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/internal/pkg/generate"
	"github.com/NethermindEth/1click/internal/ui"
	"github.com/NethermindEth/1click/internal/utils"
	posmoni "github.com/NethermindEth/posmoni/pkg/eth2"
	posmonidb "github.com/NethermindEth/posmoni/pkg/eth2/db"
	posmoninet "github.com/NethermindEth/posmoni/pkg/eth2/networking"
	"github.com/manifoldco/promptui"
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
		if err = runAndShowContainers(); err != nil {
			return []error{err}
		}
	} else {
		// Let the user decide to see the instructions for executing the scripts and exit or let the tool execute them
		if err = runScriptOrExit(); err != nil {
			return []error{err}
		}
	}

	// Initialize Eth2 Monitoring tool
	moniCfg := posmoni.ConfigOpts{
		Checkers: []posmoni.CfgChecker{
			{Key: posmoni.Execution, ErrMsg: posmoni.NoExecutionFoundError, Data: []string{configs.OnPremiseExecutionURL}},
			{Key: posmoni.Consensus, ErrMsg: posmoni.NoConsensusFoundError, Data: []string{configs.OnPremiseConsensusURL}},
		},
	}
	m, err := posmoni.NewEth2Monitor(
		posmonidb.EmptyRepository{},
		&posmoninet.BeaconClient{RetryDuration: time.Second},
		&posmoninet.ExecutionClient{RetryDuration: time.Second},
		posmoninet.SubscribeOpts{},
		moniCfg,
	)
	if err != nil {
		return []error{fmt.Errorf(configs.MonitoringToolInitError, err)}
	}

	// Track sync of execution and consensus clients
	// TODO: Parameterize wait arg of trackSync
	if err = trackSync(m, time.Minute); err != nil {
		return []error{err}
	}

	// TODO: Prompt for waiting for keystore and validator registration to run the validator

	// Run docker-compose script validator
	upCMD := commands.Runner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services: []string{"validator"},
	})
	log.Infof(configs.RunningCommand, upCMD.Cmd)
	if _, err := commands.Runner.RunCMD(upCMD); err != nil {
		return []error{fmt.Errorf(configs.CommandError, upCMD.Cmd, err)}
	}

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

func installOrShowInstructions(pending []string) (err error) {
	optInstall, optExit := "Install dependencies", "Exit. You will manage this dependencies on your own"
	prompt := promptui.Select{
		Label: "Select how to proceed with the pending dependencies",
		Items: []string{optInstall, optExit},
	}

	if err = utils.HandleInstructions(pending, utils.ShowInstructions); err != nil {
		return fmt.Errorf(configs.ShowingInstructionsError, err)
	}
	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf(configs.PromptFailedError, err)
	}

	switch result {
	case optInstall:
		return installDependencies(pending)
	default:
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	return nil
}

func installDependencies(pending []string) error {
	if err := utils.HandleInstructions(pending, utils.InstallDependency); err != nil {
		return fmt.Errorf(configs.InstallingDependenciesError, err)
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

func validateClients(allClients clients.OrderedClients, w io.Writer) (clients.Clients, error) {
	var combinedClients clients.Clients
	var err error

	if randomize {
		// Select a random execution client, consensus client and validator client
		combinedClients, err = randomizeClients(allClients)
		if err != nil {
			return combinedClients, err
		}

		log.Infof("Listing randomized clients\n\n")
		ui.WriteRandomizedClientsTable(w, ui.RandomizedClientsTable{
			Clients: []string{
				combinedClients.Execution.Name,
				combinedClients.Consensus.Name,
				combinedClients.Validator.Name,
			},
			ClientTypes: []string{
				combinedClients.Execution.Type,
				combinedClients.Consensus.Type,
				combinedClients.Validator.Type,
			},
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

		exec, ok := allClients[execution][executionName]
		if !ok {
			exec.Name = executionName
		}
		cons, ok := allClients[consensus][consensusName]
		if !ok {
			cons.Name = consensusName
		}
		val, ok := allClients[validator][validatorName]
		if !ok {
			val.Name = validatorName
		}

		combinedClients = clients.Clients{
			Execution: exec,
			Consensus: cons,
			Validator: val,
		}
	}

	if err = clients.ValidateClient(combinedClients.Execution, execution); err != nil {
		return combinedClients, err
	}
	if err = clients.ValidateClient(combinedClients.Consensus, consensus); err != nil {
		return combinedClients, err
	}
	if err = clients.ValidateClient(combinedClients.Validator, validator); err != nil {
		return combinedClients, err
	}

	return combinedClients, nil
}

func runScriptOrExit() (err error) {
	optRun, optExit := fmt.Sprintf("Run the script with the selected services %s", strings.Join(*services, ",")), "Exit"
	prompt := promptui.Select{
		Label: "Select how to proceed with the generated docker-compose script",
		Items: []string{optRun, optExit},
	}

	log.Infof(configs.InstructionsFor, "running docker-compose script")
	upCMD := commands.Runner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     generationPath,
		Services: *services,
	})
	fmt.Printf("\n%s\n\n", upCMD.Cmd)

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed %s", err)
	}

	switch result {
	case optRun:
		if err = runAndShowContainers(); err != nil {
			return err
		}
	default:
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	return nil
}

func runAndShowContainers() error {
	// TODO: (refac) Put this check to checks.go and call it from there
	// Check if docker engine is on
	log.Info(configs.CheckingDockerEngine)
	psCMD := commands.Runner.BuildDockerPSCMD(commands.DockerPSOptions{
		All: true,
	})
	psCMD.GetOutput = true
	log.Infof(configs.RunningCommand, psCMD.Cmd)
	if _, err := commands.Runner.RunCMD(psCMD); err != nil {
		return fmt.Errorf(configs.DockerEngineOffError, err)
	}

	// Run docker-compose script
	upCMD := commands.Runner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services: *services,
	})
	log.Infof(configs.RunningCommand, upCMD.Cmd)
	if _, err := commands.Runner.RunCMD(upCMD); err != nil {
		return fmt.Errorf(configs.CommandError, upCMD.Cmd, err)
	}

	// Run docker-compose ps --filter status=running to show script running containers
	dcpsCMD := commands.Runner.BuildDockerComposePSCMD(commands.DockerComposePsOptions{
		Path:     filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services: false,
	})
	log.Infof(configs.RunningCommand, dcpsCMD.Cmd)
	if _, err := commands.Runner.RunCMD(dcpsCMD); err != nil {
		return fmt.Errorf(configs.CommandError, dcpsCMD.Cmd, err)
	}

	return nil
}

// Interface for Posmoni Eth2 monitor
type monitor interface {
	TrackSync(done <-chan struct{}, beaconEndpoints, executionEndpoints []string, wait time.Duration) <-chan posmoni.EndpointSyncStatus
}

func trackSync(m monitor, wait time.Duration) error {
	done := make(chan struct{})
	statuses := m.TrackSync(done, []string{configs.OnPremiseExecutionURL}, []string{configs.OnPremiseConsensusURL}, time.Minute)

	var esynced, csynced bool
	for s := range statuses {
		if s.Error != nil {
			return fmt.Errorf(configs.TrackSyncError, s.Endpoint, s.Error)
		}
		esynced = esynced || (s.Synced && s.Endpoint == configs.OnPremiseExecutionURL)
		csynced = csynced || (s.Synced && s.Endpoint == configs.OnPremiseConsensusURL)
		if esynced && csynced {
			// Stop tracking
			close(done)
			log.Info(configs.NodesSynced)
		}
	}

	return nil
}
