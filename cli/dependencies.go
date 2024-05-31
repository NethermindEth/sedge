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

	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deps = []string{
	dependencies.Docker,
}

func DependenciesCommand(depsMgr dependencies.DependenciesManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deps",
		Short: "Manage dependencies",
		Long:  "Checks and install dependencies needed to run Sedge.",
	}
	cmd.AddCommand(dependenciesCheckCommand(depsMgr))
	cmd.AddCommand(dependenciesInstallCommand(depsMgr))
	return cmd
}

func dependenciesCheckCommand(depsMgr dependencies.DependenciesManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check dependencies",
		Long: `Checks if the following dependencies are installed on the host machine:
	- docker
	- docker compose

Also checks if the docker engine is running`,
		RunE: func(cmd *cobra.Command, args []string) error {
			checks := []depCheck{
				checkDockerIsInstalled,
				dockerDaemonIsRunning,
				dockerComposeIsInstalled,
			}
			checksOk := true
			for _, check := range checks {
				checkOk, err := check(depsMgr)
				if err != nil {
					return err
				}
				checksOk = checksOk && checkOk
			}
			if !checksOk {
				return fmt.Errorf("%w. To install dependencies if supported run: 'sedge deps install'", ErrMissingDependencies)
			} else {
				log.Info("All dependencies are installed and running")
			}
			return nil
		},
	}
	return cmd
}

func dependenciesInstallCommand(depsMgr dependencies.DependenciesManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install dependencies",
		Long: `Checks if docker is installed in the host machine. If not, it will try to install it.
Installation is only supported on Linux.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			installed, pending := depsMgr.Check(deps)
			for _, dep := range installed {
				log.Infof("%s is installed", dep)
			}
			ok := true
			for _, dep := range pending {
				log.Infof("Trying to install %s", dep)
				if err := depsMgr.Install(dep); err != nil {
					log.Error(err.Error())
					ok = false
				}
				log.Infof("%s is installed", dep)
			}
			if !ok {
				return ErrMissingDependencies
			}
			log.Info("All dependencies are installed.")
			return nil
		},
	}
	return cmd
}

type depCheck func(dependencies.DependenciesManager) (bool, error)

func checkDockerIsInstalled(depsMgr dependencies.DependenciesManager) (bool, error) {
	// Check dependencies
	installed, pending := depsMgr.Check(deps)
	for _, dep := range installed {
		log.Infof("%s is installed", dep)
	}
	for _, dep := range pending {
		log.Warnf("%s is not installed", dep)
	}
	return len(pending) == 0, nil
}

func dockerDaemonIsRunning(depsMgr dependencies.DependenciesManager) (bool, error) {
	// Check docker daemon is running
	if err := depsMgr.DockerEngineIsOn(); err != nil {
		if errors.Is(err, dependencies.ErrDockerEngineIsNotRunning) {
			log.Warn("Docker engine is not running. Please start it and try again.")
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func dockerComposeIsInstalled(depsMgr dependencies.DependenciesManager) (bool, error) {
	// Check docker compose is installed
	if err := depsMgr.DockerComposeIsInstalled(); err != nil {
		if errors.Is(err, dependencies.ErrDependencyNotInstalled) {
			log.Warn("Docker compose is not installed.")
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
