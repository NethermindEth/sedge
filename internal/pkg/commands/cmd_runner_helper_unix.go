//go:build linux || darwin
// +build linux darwin

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
package commands

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"sync"
)

/*
runCmd :
Executes a command and returns the output.

params :-
a. cmd string
The command to be executed.
b. bool getOutput
True if the output is to be returned.

returns :-
a. string
The output of the command.
b. error
Error if any
*/
func runCmd(cmd string, getOutput bool) (out string, exitCode int, err error) {
	r := strings.ReplaceAll(cmd, "\n", "")
	spl := strings.Split(r, " ")
	c, args := spl[0], spl[1:]

	exc := exec.Command(c, args...)

	var combinedOut bytes.Buffer
	if getOutput {
		// If the cmd is to get the output, then use an unified buffer to combine stdout and stderr
		exc.Stdout = &combinedOut
		exc.Stderr = &combinedOut
	} else {
		// Pipe output to stdout and stderr
		exc.Stdout = os.Stdout
		exc.Stderr = os.Stderr
	}

	// Start and wait for the command to finish
	if err = exc.Start(); err != nil {
		return out, exitCode, err
	}
	// Return this error at the end as we need to check if the output from stderr is to be returned
	err = exc.Wait()
	exitCode = exc.ProcessState.ExitCode()

	if getOutput {
		out = combinedOut.String()
	}

	return out, exitCode, err
}

/*
executeBashScript :
Execute the bash script in the given template.

params :-
a. script Script
Script object to be executed

returns :-
a. string
The output of the script.
b. error
Error if any
*/
func executeBashScript(script ScriptFile, runWithSudo bool) (out string, err error) {
	var scriptBuffer, combinedOut bytes.Buffer
	if err = script.Tmp.Execute(&scriptBuffer, script.Data); err != nil {
		return out, err
	}

	var cmd *exec.Cmd
	if runWithSudo {
		cmd = exec.Command("sudo", "bash")
	} else {
		cmd = exec.Command("bash")
	}

	// Prepare pipes for stdin, stdout and stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return out, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return out, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return out, err
	}

	wait := sync.WaitGroup{}

	// Prepare channel to receive errors from goroutines
	errChans := make([]<-chan error, 0)
	// cmd executes any instructions coming from stdin
	errChans = append(errChans, goCopy(&wait, stdin, &scriptBuffer, true))

	if script.GetOutput {
		// If the script is to get the output, then use an unified buffer to combine stdout and stderr
		cmd.Stdout = &combinedOut
		cmd.Stderr = &combinedOut
	} else {
		// If the script is not to get the output, then pipe the output to stdout and stderr
		errChans = append(errChans, goCopy(&wait, os.Stdout, stdout, false))
		errChans = append(errChans, goCopy(&wait, os.Stderr, stderr, false))
	}

	if err = cmd.Start(); err != nil {
		return out, err
	}

	// Check for errors from goroutines
	for _, errChan := range errChans {
		err = <-errChan
		if err != nil {
			return out, err
		}
	}

	wait.Wait()

	if err = cmd.Wait(); err != nil {
		return out, err
	}

	if script.GetOutput {
		out = combinedOut.String()
	}

	return out, nil
}
