package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/NethermindEth/1Click/configs"
	log "github.com/sirupsen/logrus"
)

/*
InstallDependencies :
This function is responsible for installing the dependencies needed for 1Click setup

params :-
a. []string dependencies
List of dependencies needed for 1Click setup

returns :-
a. error
Error if any
*/
func InstallDependencies(dependencies []string) error {
	switch os := runtime.GOOS; os {
	case "linux":
		for _, dependency := range dependencies {
			err := install(dependency)
			if err != nil {
				return err
			}
		}
	case "windows":
		return fmt.Errorf("dependencies are not installed on your machine. Please install them and try again")
	default:
		log.Fatalf("Dependencies %s are not installed on your machine. Please install them and try again.", strings.Join(dependencies, ", "))
	}
	return nil
}

/*
install :
This function is responsible for installing the dependencies
for linux systems using pre-written bash scripts

params :-
a. string dependency
Dependency to be installed

returns :-
a. error
Error if any
*/
func install(dependency string) (err error) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}

	distro, err := GetOSInfo()
	if err != nil {
		return
	}

	scriptPath := fmt.Sprintf("%s/%s/linux/%s/%s.sh", pwd, configs.InstallScriptsPath, dependency, distro.Name)

	//TODO: Consider not showing the option to install the dependencies if the script is not available for the host machine distro
	if _, err := os.Stat(scriptPath); errors.Is(err, os.ErrNotExist) {
		// script does not exist
		return fmt.Errorf(configs.InstallNotSupported, dependency, distro.Name)
	}

	// Not sure which is the best way, the two below works
	cmd := exec.Command("bash", "-c", scriptPath)
	//cmd := exec.Command(fmt.Sprintf("%s/%s/linux/%s/%s.sh", pwd, configs.InstallScriptsPath, dependency, distro.Name))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
