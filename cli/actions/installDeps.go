package actions

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
)

type InstallDependenciesOptions struct {
	Dependencies []string
	Install      bool
}

func (s *sedgeActions) InstallDependencies(options InstallDependenciesOptions) error {
	log.Infof(configs.CheckingDependencies, strings.Join(options.Dependencies, ", "))

	// Check if dependencies are installed. Keep checking dependencies until they are all installed
	for pending := utils.CheckDependencies(options.Dependencies); len(pending) > 0; pending = utils.CheckDependencies(options.Dependencies) {
		log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
		if options.Install {
			// Install dependencies directly
			if err := installDependencies(s.commandRunner, pending); err != nil {
				return err
			}
		} else {
			// Let the user decide to see the instructions for installing dependencies and exit or let the tool install them and continue
			if err := installOrShowInstructions(s.commandRunner, pending); err != nil {
				return err
			}
		}
	}
	log.Info(configs.DependenciesOK)
	return nil
}

func installDependencies(cmdRunner commands.CommandRunner, pending []string) error {
	if runtime.GOOS != "windows" { // Windows doesn't support docker installation through scripts
		if err := utils.HandleInstructions(cmdRunner, pending, utils.InstallDependency); err != nil {
			return fmt.Errorf(configs.InstallingDependenciesError, err)
		}
	}
	return nil
}

func installOrShowInstructions(cmdRunner commands.CommandRunner, pending []string) (err error) {
	// notest
	optInstall, optExit := "Install dependencies", "Exit. You will manage this dependencies on your own"
	prompt := promptui.Select{
		Label: "Select how to proceed with the pending dependencies",
		Items: []string{optInstall, optExit},
	}

	if err = utils.HandleInstructions(cmdRunner, pending, utils.ShowInstructions); err != nil {
		return fmt.Errorf(configs.ShowingInstructionsError, err)
	}
	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf(configs.PromptFailedError, err)
	}

	switch result {
	case optInstall:
		return installDependencies(cmdRunner, pending)
	default:
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	return nil
}
