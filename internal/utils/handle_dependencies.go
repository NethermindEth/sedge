package utils

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/templates"
	log "github.com/sirupsen/logrus"
)

/*
HandleInstructions :
This function is responsible for handling the dependencies needed for 1Click setup
If install support for a dependency exists, then `handler` will process it, for example,
installing it or showing instructions for it.

params :-
a. []string dependencies
List of dependencies needed for 1Click setup
b. func(string) error handler
Handler for each dependency

returns :-
a. error
Error if any
*/
func HandleInstructions(dependencies []string, handler func(string) error) error {
	pending := make([]string, 0)

	for _, dependency := range dependencies {
		if dependencySupported(dependency) {
			err := handler(dependency)
			if err != nil {
				log.Error(err)
				pending = append(pending, dependency)
			}
		} else {
			log.Errorf(configs.InstallNotSupported, dependency, runtime.GOOS)
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
func ShowInstructions(dependency string) error {
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
func InstallDependency(dependency string) (err error) {
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

	script := Script{
		Tmp:    tmp,
		Output: false,
		Data:   struct{}{},
	}

	if _, err = executeScript(script); err != nil {
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
	if runtime.GOOS != "linux" {
		return "", DistroInfo{}, fmt.Errorf(configs.OSNotSupported, runtime.GOOS)
	}

	distro, err = GetOSInfo()
	if err != nil {
		return
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
