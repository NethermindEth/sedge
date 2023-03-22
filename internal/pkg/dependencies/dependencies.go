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
package dependencies

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/NethermindEth/sedge/templates"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -package=sedge_mocks -destination=../../../mocks/dependenciesManager.go github.com/NethermindEth/sedge/internal/pkg/dependencies DependenciesManager
type DependenciesManager interface {
	Supported(dependencies []string) (supported []string, unsupported []string, err error)
	Check(dependencies []string) (installed []string, pending []string)
	Install(dependency string) error
	ShowInstructions(dependency string) error
	DockerEngineIsOn() error
	DockerComposeIsInstalled() error
}

type dependenciesManager struct {
	cmdRunner commands.CommandRunner
}

func NewDependenciesManager(cmdRunner commands.CommandRunner) DependenciesManager {
	return &dependenciesManager{
		cmdRunner: cmdRunner,
	}
}

// Supported returns the supported and unsupported dependencies of the given list
func (d *dependenciesManager) Supported(dependencies []string) (supported []string, unsupported []string, err error) {
	for _, dependency := range dependencies {
		if _, err := getScriptPath(dependency); err != nil {
			if errors.Is(err, ErrUnsupportedInstallForOS) || errors.Is(err, ErrUnsupportedDependency) {
				unsupported = append(unsupported, dependency)
			} else {
				return nil, nil, err
			}
		} else {
			supported = append(supported, dependency)
		}
	}
	return
}

// Check checks if the dependencies are installed on the host machine
func (d *dependenciesManager) Check(dependencies []string) (installed []string, pending []string) {
	for _, dependency := range dependencies {
		_, err := exec.LookPath(dependency)
		if err != nil {
			log.Error(fmt.Errorf("%w: %s", ErrDependencyNotInstalled, dependency))
			pending = append(pending, dependency)
		} else {
			installed = append(installed, dependency)
		}
	}
	return
}

// Install installs the dependency on the host machine
func (d *dependenciesManager) Install(dependency string) error {
	scriptPath, err := getScriptPath(dependency)
	if err != nil {
		return err
	}
	rawScript, err := templates.Setup.ReadFile(scriptPath)
	if err != nil {
		return err
	}
	tmp, err := template.New("script").Parse(string(rawScript))
	if err != nil {
		return err
	}
	script := commands.ScriptFile{
		Tmp:       tmp,
		GetOutput: false,
		Data:      struct{}{},
	}
	_, err = d.cmdRunner.RunScript(script)
	return err
}

// ShowInstructions shows the instructions to install the dependency on the host machine
func (d *dependenciesManager) ShowInstructions(dependency string) error {
	scriptPath, err := getScriptPath(dependency)
	if err != nil {
		return err
	}

	content, err := templates.Setup.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf(configs.ReadingInstructionError, scriptPath)
	}

	trimmedContent := utils.SkipLines(string(content), "#")
	trimmedContent = strings.Trim(trimmedContent, "\n")
	log.Infof(configs.InstructionsFor, dependency)
	fmt.Printf("\n%s\n\n", trimmedContent)

	return nil
}

// DockerEngineIsOn checks if docker engine is on by calling `docker ps -a`
func (d *dependenciesManager) DockerEngineIsOn() error {
	dockerPsCMD := d.cmdRunner.BuildDockerPSCMD(commands.DockerPSOptions{All: true})
	log.Debugf(configs.RunningCommand, dockerPsCMD.Cmd)
	dockerPsCMD.GetOutput = true
	if _, _, err := d.cmdRunner.RunCMD(dockerPsCMD); err != nil {
		return ErrDockerEngineIsNotRunning
	}
	return nil
}

// DockerComposeIsInstalled checks if docker-compose is installed calling `docker-compose version`
func (d *dependenciesManager) DockerComposeIsInstalled() error {
	dockerComposeVersionCMD := d.cmdRunner.BuildDockerComposeVersionCMD()
	log.Debugf(configs.RunningCommand, dockerComposeVersionCMD.Cmd)
	dockerComposeVersionCMD.GetOutput = true
	if _, _, err := d.cmdRunner.RunCMD(dockerComposeVersionCMD); err != nil {
		return fmt.Errorf("%w: %s", ErrDependencyNotInstalled, "docker-compose")
	}
	return nil
}

func getScriptPath(dependency string) (string, error) {
	if runtime.GOOS != "linux" {
		return "", fmt.Errorf("%w: %s", ErrUnsupportedInstallForOS, runtime.GOOS)
	}
	distro, err := getOSInfo()
	if err != nil {
		return "", fmt.Errorf(configs.DistroInfoError, err)
	}

	scriptPath := fmt.Sprintf("setup/%s/%s/%s_%s.sh", runtime.GOOS, dependency, distro.Name, distro.Version)
	_, err = templates.Setup.Open(scriptPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("%w: %s", ErrUnsupportedDependency, dependency)
		} else {
			return "", err
		}
	}
	return scriptPath, nil
}
