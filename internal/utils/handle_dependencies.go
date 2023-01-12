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
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/templates"
	log "github.com/sirupsen/logrus"
)

/*
HandleInstructions :
This function is responsible for handling the dependencies needed for sedge setup
If install support for a dependency exists, then `handler` will process it, for example,
installing it or showing instructions for it.

params :-
a. []string dependencies
List of dependencies needed for sedge setup
b. func(string) error handler
Handler for each dependency

returns :-
a. error
Error if any
*/
func HandleInstructions(cmdRunner commands.CommandRunner, dependencies []string, handler func(commands.CommandRunner, string) error) (err error) {
	pending := make([]string, 0)

	for _, dependency := range dependencies {
		if dependencySupported(dependency) {
			err = handler(cmdRunner, dependency)
			if err != nil {
				log.Error(err)
				pending = append(pending, dependency)
			}
		} else {
			var OS string
			if runtime.GOOS == "linux" {
				if OS, err = GetDistroName(); err != nil {
					return
				}
			} else {
				// TODO: Get OS version for other OS like Windows and Darwin
				OS = runtime.GOOS
			}
			log.Errorf(configs.InstallNotSupported, dependency, OS)
			pending = append(pending, dependency)
		}
	}

	if len(pending) > 0 {
		return fmt.Errorf(configs.DependenciesPending, strings.Join(pending, ", "))
	}
	return nil
}

/*
ShowInstructions :
This function is responsible for showing instructions for installing
each dependency

params :-
a. string dependency
Dependency whose instruction will be showed

returns :-
a. error
Error if any
*/
func ShowInstructions(_ commands.CommandRunner, dependency string) error {
	scriptPath, _, err := getScriptPath(dependency)
	if err != nil {
		return fmt.Errorf(configs.ScriptPathError, err)
	}

	content, err := templates.Setup.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf(configs.ReadingInstructionError, scriptPath)
	}

	trimmedContent := SkipLines(string(content), "#")
	trimmedContent = strings.Trim(trimmedContent, "\n")
	log.Infof(configs.InstructionsFor, dependency)
	fmt.Printf("\n%s\n\n", trimmedContent)

	return nil
}

/*
InstallDependency :
This function is responsible for installing the dependencies using pre-written bash scripts

params :-
a. string dependency
Dependency to be installed

returns :-
a. error
Error if any
*/
func InstallDependency(cmdRunner commands.CommandRunner, dependency string) (err error) {
	scriptPath, _, err := getScriptPath(dependency)
	if err != nil {
		return
	}

	rawScript, err := templates.Setup.ReadFile(scriptPath)
	if err != nil {
		return
	}

	tmp, err := template.New("script").Parse(string(rawScript))
	if err != nil {
		return
	}

	script := commands.ScriptFile{
		Tmp:       tmp,
		GetOutput: false,
		Data:      struct{}{},
	}

	if _, err = cmdRunner.RunScript(script); err != nil {
		return
	}

	return nil
}

/*
getScriptPath :
Give the path to dependency embedded install script.
Currently support only for linux.

params :-
a. string dependency
Dependency to be installed

returns :-
a. string
Path to the install script
b. DistroInfo
Linux distribution information
c. error
Error if any
*/
func getScriptPath(dependency string) (path string, distro DistroInfo, err error) {
	// DEV: Only for linux
	if runtime.GOOS != "linux" {
		return "", DistroInfo{}, fmt.Errorf(configs.OSNotSupported, runtime.GOOS)
	}

	distro, err = getOSInfo()
	if err != nil {
		return path, distro, fmt.Errorf(configs.DistroInfoError, err)
	}

	path = fmt.Sprintf("setup/%s/%s/%s_%s.sh", runtime.GOOS, dependency, distro.Name, distro.Version)
	return path, distro, nil
}

/*
dependencySupported :
Check if exist support for dependency installation.

params :-
a. string dependency
Dependency to be installed

returns :-
a. bool
True if support exists, false otherwise
*/
func dependencySupported(dependency string) bool {
	if runtime.GOOS != "linux" {
		return false
	}

	scriptPath, _, err := getScriptPath(dependency)
	if err != nil {
		return false
	}

	if _, err := templates.Setup.Open(scriptPath); errors.Is(err, os.ErrNotExist) {
		// script does not exist
		return false
	}

	return true
}

func DependenciesSupported(dependencies []string) (supported []string, unsupported []string) {
	for _, dependency := range dependencies {
		if dependencySupported(dependency) {
			supported = append(supported, dependency)
		} else {
			unsupported = append(unsupported, dependency)
		}
	}
	return
}
