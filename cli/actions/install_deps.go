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

type DependenciesHandlers interface {
	InstallDependencies(cmdRunner commands.CommandRunner, dependencies []string) error
	InstallOrShowInstructions(cmdRunner commands.CommandRunner, dependencies []string) error
}

type dependenciesHandlers struct {
}

func (dh *dependenciesHandlers) InstallDependencies(cmdRunner commands.CommandRunner, pending []string) error {
	// notest
	if runtime.GOOS == "linux" { // Windows and MacOS don't support docker installation through scripts
		if err := utils.HandleInstructions(cmdRunner, pending, utils.InstallDependency); err != nil {
			return fmt.Errorf(configs.InstallingDependenciesError, err)
		}
	}
	return nil
}

func (dh *dependenciesHandlers) InstallOrShowInstructions(cmdRunner commands.CommandRunner, pending []string) (err error) {
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
		return dh.InstallDependencies(cmdRunner, pending)
	default:
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	return nil
}

func NewDependenciesHandlers() DependenciesHandlers {
	return &dependenciesHandlers{}
}

type ManageDependenciesOptions struct {
	Dependencies []string
	Install      bool
}

func (s *sedgeActions) ManageDependencies(options ManageDependenciesOptions) error {
	log.Infof(configs.CheckingDependencies, strings.Join(options.Dependencies, ", "))

	// Check if dependencies are installed. Keep checking dependencies until they are all installed
	for pending := utils.CheckDependencies(options.Dependencies); len(pending) > 0; pending = utils.CheckDependencies(options.Dependencies) {
		log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
		if options.Install {
			// Install dependencies directly
			if err := s.depsHandlers.InstallDependencies(s.commandRunner, pending); err != nil {
				return err
			}
		} else {
			// Let the user decide to see the instructions for installing dependencies and exit or let the tool install them and continue
			if err := s.depsHandlers.InstallOrShowInstructions(s.commandRunner, pending); err != nil {
				return err
			}
		}
	}
	log.Info(configs.DependenciesOK)
	return nil
}
