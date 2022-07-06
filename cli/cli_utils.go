package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/pkg/clients"
	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/internal/utils"
	"github.com/NethermindEth/1click/templates"
	posmoni "github.com/NethermindEth/posmoni/pkg/eth2"
	"github.com/manifoldco/promptui"
)

// Interface for Posmoni Eth2 monitor
type MonitoringTool interface {
	TrackSync(done <-chan struct{}, beaconEndpoints, executionEndpoints []string, wait time.Duration) <-chan posmoni.EndpointSyncStatus
}

var monitor MonitoringTool

func initMonitor(builder func() MonitoringTool) {
	monitor = builder()
}

func installOrShowInstructions(pending []string) (err error) {
	// notest
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
	var executionClient, consensusClient clients.Client
	var combinedClients clients.Clients

	executionClient, err := clients.RandomChoice(allClients[execution])
	if err != nil {
		return combinedClients, err
	}
	consensusClient, err = clients.RandomChoice(allClients[consensus])
	if err != nil {
		return combinedClients, err
	}

	combinedClients = clients.Clients{
		Execution: executionClient,
		Consensus: consensusClient,
		Validator: consensusClient}
	return combinedClients, nil
}

func validateClients(allClients clients.OrderedClients, w io.Writer) (clients.Clients, error) {
	var combinedClients clients.Clients
	var err error

	// Select a random execution client, consensus client and validator client
	randomizedClients, err := randomizeClients(allClients)
	if err != nil {
		return combinedClients, err
	}

	// Randomize missing clients, and choose same pair of client for consensus and validator if at least one of them is missing
	if executionName == "" {
		log.Warnf(configs.ExecutionClientNotSpecifiedWarn, randomizedClients.Execution.Name)
		executionName = randomizedClients.Execution.Name
	}
	if consensusName == "" && validatorName == "" {
		log.Warnf(configs.CLNotSpecifiedWarn, randomizedClients.Consensus.Name)
		consensusName = randomizedClients.Consensus.Name
		validatorName = randomizedClients.Validator.Name
	} else if consensusName == "" {
		log.Warn(configs.ConsensusClientNotSpecifiedWarn)
		consensusName = validatorName
	} else if validatorName == "" {
		log.Warn(configs.ValidatorClientNotSpecifiedWarn)
		validatorName = consensusName
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
	// notest
	log.Infof(configs.InstructionsFor, "running docker-compose script")
	upCMD := commands.Runner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services: *services,
	})
	fmt.Printf("\n%s\n\n", upCMD.Cmd)

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Run the script with the selected services %s", strings.Join(*services, ", ")),
		IsConfirm: true,
		Default:   "Y",
	}
	_, err = prompt.Run()
	if err != nil {
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	if err = runAndShowContainers(*services); err != nil {
		return err
	}

	return nil
}

func runAndShowContainers(services []string) error {
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
		Services: services,
	})
	log.Infof(configs.RunningCommand, upCMD.Cmd)
	if _, err := commands.Runner.RunCMD(upCMD); err != nil {
		return fmt.Errorf(configs.CommandError, upCMD.Cmd, err)
	}

	// Run docker compose ps --filter status=running to show script running containers
	dcpsCMD := commands.Runner.BuildDockerComposePSCMD(commands.DockerComposePsOptions{
		Path:          filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		FilterRunning: true,
	})
	log.Infof(configs.RunningCommand, dcpsCMD.Cmd)
	if _, err := commands.Runner.RunCMD(dcpsCMD); err != nil {
		return fmt.Errorf(configs.CommandError, dcpsCMD.Cmd, err)
	}

	return nil
}

type container struct {
	NetworkSettings networkSettings
}
type networkSettings struct {
	Networks map[string]networks
}
type networks struct {
	IPAddress string
}

func parseNetwork(js string) (string, error) {
	var c []container
	if err := json.NewDecoder(bytes.NewReader([]byte(js))).Decode(&c); err != nil {
		return "", err
	}
	if len(c) == 0 {
		return "", errors.New(configs.NoOutputDockerInspectError)
	}
	if ip := c[0].NetworkSettings.Networks["1click_network"].IPAddress; ip != "" {
		return ip, nil
	}
	return "", errors.New(configs.IPNotFoundError)
}

func getContainerIP(service string) (ip string, err error) {
	// Run docker compose ps --quiet <service> to show service's ID
	dcpsCMD := commands.Runner.BuildDockerComposePSCMD(commands.DockerComposePsOptions{
		Path:        filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Quiet:       true,
		ServiceName: service,
	})
	log.Infof(configs.RunningCommand, dcpsCMD.Cmd)
	dcpsCMD.GetOutput = true
	id, err := commands.Runner.RunCMD(dcpsCMD)
	if err != nil {
		return ip, fmt.Errorf(configs.CommandError, dcpsCMD.Cmd, err)
	}

	// Run docker inspect <id> to get IP address
	inspectCmd := commands.Runner.BuildDockerInspectCMD(commands.DockerInspectOptions{
		Name: id,
	})
	log.Infof(configs.RunningCommand, inspectCmd.Cmd)
	inspectCmd.GetOutput = true
	data, err := commands.Runner.RunCMD(inspectCmd)
	if err != nil {
		return
	}

	ip, err = parseNetwork(data)
	return
}

func trackSync(m MonitoringTool, elPort, clPort string, wait time.Duration) error {
	done := make(chan struct{})
	defer close(done)

	log.Info(configs.GettingContainersIP)
	executionIP, errE := getContainerIP(execution)
	if errE != nil {
		log.Errorf(configs.GetContainerIPError, execution, errE)
	}
	consensusIP, errC := getContainerIP(consensus)
	if errC != nil {
		log.Errorf(configs.GetContainerIPError, consensus, errC)
		if errE != nil {
			// Both IP were not detected, both containers probably failed
			return errors.New(configs.UnableToTrackSyncError)
		}
	}

	consensusUrl := fmt.Sprintf("http://%s:%s", consensusIP, clPort)
	executionUrl := fmt.Sprintf("http://%s:%s", executionIP, elPort)

	statuses := m.TrackSync(done, []string{consensusUrl}, []string{executionUrl}, wait)

	var esynced, csynced bool
	for s := range statuses {
		if s.Error != nil {
			return fmt.Errorf(configs.TrackSyncError, s.Endpoint, s.Error)
		}
		esynced = esynced || (s.Synced && s.Endpoint == executionUrl)
		csynced = csynced || (s.Synced && s.Endpoint == consensusUrl)
		if esynced && csynced {
			// Stop tracking
			done <- struct{}{}
			log.Info(configs.NodesSynced)
			break // statuses channel might still have data before closing done channel
		}
	}

	return nil
}

func RunValidatorOrExit() error {
	// notest
	log.Infof(configs.InstructionsFor, "running validator service of docker-compose script")
	upCMD := commands.Runner.BuildDockerComposeUpCMD(commands.DockerComposeUpOptions{
		Path:     filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services: []string{validator},
	})
	fmt.Printf("\n%s\n\n", upCMD.Cmd)

	prompt := promptui.Prompt{
		Label:     "Run validator service",
		IsConfirm: true,
		Default:   "Y",
	}
	_, err := prompt.Run()
	if err != nil {
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	if err = runAndShowContainers([]string{validator}); err != nil {
		return err
	}

	return nil
}

func handleJWTSecret() error {
	log.Info(configs.GeneratingJWTSecret)

	rawScript, err := templates.Scripts.ReadFile(filepath.Join("scripts", "jwt_secret.sh"))
	if err != nil {
		return fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	tmp, err := template.New("script").Parse(string(rawScript))
	if err != nil {
		return fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	script := commands.BashScript{
		Tmp:       tmp,
		GetOutput: false,
		Data: map[string]string{
			"Path": generationPath,
		},
	}

	if _, err = commands.Runner.RunBash(script); err != nil {
		return fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	jwtPath, err = filepath.Abs(filepath.Join(generationPath, "jwtsecret"))
	if err != nil {
		return fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	log.Info(configs.JWTSecretGenerated)
	return nil
}

func feeRecipientPrompt() error {
	// notest
	validate := func(input string) error {
		if input != "" && !utils.IsAddress(input) {
			return errors.New(configs.InvalidFeeRecipientError)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter the Fee Recipient address. You can leave it blank and press enter (not recommended)",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		return fmt.Errorf(configs.PromptFailedError, err)
	}

	feeRecipient = result
	return nil
}

func preRunTeku() error {
	log.Info(configs.PreparingTekuDatadir)
	for _, s := range *services {
		if s == "all" || s == consensus {
			// Prepare consensus datadir
			path := filepath.Join(generationPath, configs.ConsensusDefaultDataDir)
			if err := os.MkdirAll(path, 0777); err != nil {
				return fmt.Errorf(configs.TekuDatadirError, consensus, err)
			}
		}
		if s == "all" || s == validator {
			// Prepare validator datadir
			path := filepath.Join(generationPath, configs.ValidatorDefaultDataDir)
			if err := os.MkdirAll(path, 0777); err != nil {
				return fmt.Errorf(configs.TekuDatadirError, validator, err)
			}
		}
	}
	return nil
}
