package commands

import (
	"fmt"
	"testing"
	"text/template"
)

func TestRunCmd(t *testing.T) {
	//TODO: improve test and fix pty commands tests
	inputs := []struct {
		cmd       string
		getOutput bool
		runInPty  bool
		output    string
		isErr     bool
	}{
		{
			cmd:       "echo hello world",
			getOutput: true,
			output:    "hello world\n",
			isErr:     false,
		},
		{
			cmd:       "echo hello world",
			getOutput: true,
			runInPty:  true,
			output:    "hello world\n",
			isErr:     true,
		},
		{
			cmd:   "wr0n6",
			isErr: true,
		},
	}

	InitRunner(func() CommandRunner {
		return NewCMDRunner(CMDRunnerOptions{
			RunAsAdmin: false,
		})
	})

	for _, input := range inputs {
		descr := fmt.Sprintf("RunCmd(%s,%t,%t)", input.cmd, input.getOutput, input.runInPty)

		got, err := Runner.RunCMD(Command{
			Cmd:       input.cmd,
			GetOutput: input.getOutput,
			RunInPty:  input.runInPty,
		})
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if input.getOutput && input.output != got {
				t.Errorf("%s expected %s but got %s", descr, input.output, got)
			}
		}
	}
}

func TestRunBashScript(t *testing.T) {
	inputs := []struct {
		cmd       string
		getOutput bool
		output    string
		isErr     bool
	}{
		{
			cmd:       "echo hello world",
			getOutput: true,
			output:    "hello world\n",
			isErr:     false,
		},
		{
			cmd:   "wr0n6",
			isErr: true,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("RunBashCmd(%s,%t)", input.cmd, input.getOutput)

		tmp, err := template.New("script").Parse(string(input.cmd))
		if err != nil {
			t.Fatalf("Unexpected error at case %q: %v", input.cmd, err)
		}

		got, err := Runner.RunBash(BashScript{
			Tmp:       tmp,
			GetOutput: input.getOutput,
		})
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if input.getOutput && input.output != got {
				t.Errorf("%s expected %s but got %s", descr, input.output, got)
			}
		}
	}
}

//TODO: add test cases for building and executing docker commands
