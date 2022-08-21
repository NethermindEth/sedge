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
package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	log "github.com/sirupsen/logrus"
)

/*
CheckDependencies :
This function is responsible for checking if on-premise setup dependencies are installed on host machine

params :-
a. dependencies []string
List of dependencies to be checked

returns :-
a. []string
List of dependencies that are not installed
*/
func CheckDependencies(dependencies []string) (pending []string) {
	for _, dependency := range dependencies {
		_, err := exec.LookPath(dependency)
		if err != nil {
			log.Errorf(configs.DependencyNotInstalledError, dependency)
			pending = append(pending, dependency)
		}
	}
	return
}

/*
PreCheck :
Check if docker-compose can be used to interact with the generated docker-compose script

params :-
a. generationPath string
Path to the generated docker-compose script

returns :-
a. error
Error if any
*/
func PreCheck(generationPath string) error {
	// Check docker is installed
	pending := CheckDependencies([]string{"docker"})
	for _, dependency := range pending {
		log.Errorf(configs.DependencyNotInstalledError, dependency)
	}
	if len(pending) > 0 {
		return fmt.Errorf(configs.DependenciesMissingError)
	}

	// Check docker engine is on
	dockerPsCMD := commands.Runner.BuildDockerPSCMD(commands.DockerPSOptions{All: true})
	log.Debugf(configs.RunningCommand, dockerPsCMD.Cmd)
	dockerPsCMD.GetOutput = true
	_, err := commands.Runner.RunCMD(dockerPsCMD)
	if err != nil {
		return fmt.Errorf(configs.DockerEngineOffError, err)
	}

	// Check if docker-compose script was generated
	file := generationPath + "/" + configs.DefaultDockerComposeScriptName
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Errorf(configs.OpeningFileError, file, err)
		return fmt.Errorf(configs.DockerComposeScriptNotFoundError, generationPath, configs.DefaultDockerComposeScriptsPath)
	}

	// Check that compose plugin is installed with docker running 'docker compose ps'
	dockerComposePsCMD := commands.Runner.BuildDockerComposePSCMD(commands.DockerComposePsOptions{Path: file})
	log.Debugf(configs.RunningCommand, dockerComposePsCMD.Cmd)
	dockerComposePsCMD.GetOutput = true
	_, err = commands.Runner.RunCMD(dockerComposePsCMD)
	if err != nil {
		return fmt.Errorf(configs.DockerComposeOffError, err)
	}

	return nil
}

/*
CheckContainers :
Check if containers of generated docker-compose script are running

params :-
a. generationPath string
Path to the generated docker-compose script

returns :-
a. string
Output of 'docker ps --services --filter status=running'
b. error
Error if any
*/
func CheckContainers(generationPath string) (string, error) {
	// Check if docker-compose script is running
	psCMD := commands.Runner.BuildDockerComposePSCMD(commands.DockerComposePsOptions{
		Path:          filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services:      true,
		FilterRunning: true,
	})
	log.Debugf(configs.RunningCommand, psCMD.Cmd)
	psCMD.GetOutput = true
	rawServices, err := commands.Runner.RunCMD(psCMD)
	if err != nil || rawServices == "\n" {
		if rawServices == "\n" && err == nil {
			err = fmt.Errorf(configs.DockerComposePsReturnedEmptyError)
		}
		return "", fmt.Errorf(configs.ScriptIsNotRunningError, err)
	}

	return rawServices, nil
}
