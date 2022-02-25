package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"text/template"

	"github.com/NethermindEth/1Click/configs"
)

/*
RunCmd :
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
func RunCmd(cmd string, getOutput bool) (out string, err error) {
	tmp, err := template.New("script").Parse(cmd)
	if err != nil {
		return "", err
	}

	script := Script{
		Tmp:       tmp,
		GetOutput: getOutput,
		Data:      struct{}{},
	}

	if out, err = executeScript(script); err != nil {
		return "", fmt.Errorf(configs.RunningCMDError, cmd, err)
	}

	return out, nil
}

/*
executeScript :
Execute the script in the given template.

params :-
a. script Script
Script object to be executed

returns :-
a. string
The output of the script.
b. error
Error if any
*/
func executeScript(script Script) (out string, err error) {
	var scriptBuffer, combinedOut bytes.Buffer
	if err = script.Tmp.Execute(&scriptBuffer, script.Data); err != nil {
		return
	}

	cmd := exec.Command("bash")

	// Prepare pipes for stdin, stdout and stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	wait := sync.WaitGroup{}

	// Prepare channel to receive errors from goroutines
	errChans := make([]<-chan error, 0)
	// cmd executes any instrucctions coming from stdin
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
		return
	}

	// Check for errors from goroutines
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

	if script.GetOutput {
		out = combinedOut.String()
	}

	return out, nil
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
