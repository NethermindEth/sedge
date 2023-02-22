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
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/NethermindEth/sedge/internal/utils"
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

func installOrShowInstructions(cmdRunner commands.CommandRunner, pending []string) error {
	// notest
	optInstall, optExit := "Install dependencies", "Exit. You will manage this dependencies on your own"
	prompt := ui.NewPrompter()
	options := []string{optInstall, optExit}
	index, err := prompt.Select("Select how to proceed with the pending dependencies", "", options)
	if err != nil {
		return err
	}
	if err = utils.HandleInstructions(cmdRunner, pending, utils.ShowInstructions); err != nil {
		return fmt.Errorf(configs.ShowingInstructionsError, err)
	}

	switch options[index] {
	case optInstall:
		return installDependencies(cmdRunner, pending)
	default:
		log.Info(configs.Exiting)
		os.Exit(0)
	}

	return nil
}
