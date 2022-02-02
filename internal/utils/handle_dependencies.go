package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
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

	if err = executeScript(tmp); err != nil {
		return
	}

	return nil
}

/*
executeScript :
Execute the script in the given template and print its output to stdout.

params :-
a. tmp *template.Template
Script template

returns :-
a. error
Error if any
*/
func executeScript(tmp *template.Template) (err error) {
	var script bytes.Buffer
	if err = tmp.Execute(&script, struct{}{}); err != nil {
		return
	}

	cmd := exec.Command("bash")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	wait := sync.WaitGroup{}

	errChans := make([]<-chan error, 0)
	errChans = append(errChans, goCopy(&wait, stdin, &script, true))
	errChans = append(errChans, goCopy(&wait, os.Stdout, stdout, false))
	errChans = append(errChans, goCopy(&wait, os.Stderr, stderr, false))

	if err = cmd.Start(); err != nil {
		return
	}

	for _, errChan := range errChans {
		err = <-errChan
		if err != nil {
			return
		}
	}

	wait.Wait()

	if err = cmd.Wait(); err != nil {
		return
	}

	return nil
}

/*
goCopy :
Copy the content from reader(src) to writer(dst).

params :-
a. wait *sync.WaitGroup
Wait group to wait for copying to finish
b. dst io.Writer
Destination to write to
c. src io.Reader
Source to read from
d. isStdin bool
True if the destination is stdin, false otherwise

returns :-
a. chan error
Channel to where error will be sent
*/
func goCopy(wait *sync.WaitGroup, dst io.WriteCloser, src io.Reader, isStdin bool) <-chan error {
	errChan := make(chan error)
	wait.Add(1)
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			errChan <- err
			return
		}
		if isStdin {
			if err := dst.Close(); err != nil {
				errChan <- err
				return
			}
		}
		close(errChan)
		wait.Done()
	}()
	return errChan
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
