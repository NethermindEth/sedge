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
	log "github.com/sirupsen/logrus"
)

/*
RunCmd :
Executes a command and prints its output to stdout.

params :-
a. cmd string
The command to be executed.
b. args []string
The arguments to be passed to the command.

returns :-
a. string
The output of the command.
b. error
Error if any
*/
func RunCmd(cmd string, args ...string) error {
	log.Info(configs.RunningCommand)
	fullCmd := fmt.Sprintf(cmd, args)
	tmp, err := template.New("script").Parse(fullCmd)
	if err != nil {
		return err
	}

	return executeScript(tmp)
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
