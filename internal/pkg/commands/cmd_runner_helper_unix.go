//go:build linux || darwin
// +build linux darwin

package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/NethermindEth/sedge/configs"
	"github.com/creack/pty"
	log "github.com/sirupsen/logrus"
	"golang.org/x/term"
)

/*
runCmd :
Executes a command and returns the output.

params :-
a. cmd string
The command to be executed.
b. bool getOutput
True if the output is to be returned.
c. bool runInPty
True if the command is to be run in a pty, false otherwise.

returns :-
a. string
The output of the command.
b. error
Error if any
*/
func runCmd(cmd string, getOutput, runInPTY bool) (out string, err error) {
	r := strings.ReplaceAll(cmd, "\n", "")
	spl := strings.Split(r, " ")
	c, args := spl[0], spl[1:]

	exc := exec.Command(c, args...)

	if runInPTY {
		return runInPty(exc, getOutput)
	}

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
		return
	}
	// Return this error at the end as we need to check if the output from stderr is to be returned
	err = exc.Wait()

	if getOutput {
		out = combinedOut.String()
	}

	return
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
func executeBashScript(script BashScript) (out string, err error) {
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

/*
runInPty :
Executes a command in a pty and returns the output.

params :-
a. cmd *exec.Cmd
The command to be executed.
b. bool getOutput
True if the output is to be returned.

returns :-
a. string
The output of the command.
b. error
Error if any
*/
func runInPty(cmd *exec.Cmd, getOutput bool) (out string, err error) {
	// Start the command with a pty.
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return "", err
	}
	// Make sure to close the pty at the end.
	defer func() {
		cErr := ptmx.Close()
		if err == nil && cErr != nil {
			log.Error(cErr)
			err = cErr
		}
	}()

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	errCh := make(chan error)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for sig := range ch {
			log.Debug(sig)
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Error(err)
				errCh <- fmt.Errorf(configs.ResizingPtyError, err)
			}
		}
		close(errCh)
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	for {
		select {
		case err = <-errCh:
			// Check resizing errors
			if err != nil {
				return
			}
		default:
			// Normal workflow
			// Set stdin in raw mode.
			oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
			if err != nil {
				return "", err
			}
			defer func() {
				rErr := term.Restore(int(os.Stdin.Fd()), oldState)
				if err == nil && rErr != nil {
					err = rErr
				}
			}()

			// Copy stdin to the pty (where are not using stdin at the moment)
			// NOTE: The goroutine will keep reading until the next keystroke before returning.
			go func() {
				_, err = io.Copy(ptmx, os.Stdin)
				log.Error(err)
			}()

			// Handle output
			var output bytes.Buffer
			if getOutput {
				// Copy the pty to out
				_, _ = io.Copy(&output, ptmx)
				out = output.String()
			} else {
				// Copy the pty to stdout
				_, _ = io.Copy(os.Stdout, ptmx)
			}

			return out, nil
		}
	}
}
