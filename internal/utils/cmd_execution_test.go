package utils

import (
	"fmt"
	"testing"
)

func TestRunCmd(t *testing.T) {
	//TODO: improve test and fix pty commands tests
	inputs := []struct {
		cmd       string
		getOutput bool
		tty       bool
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
			tty:       true,
			output:    "hello world\n",
			isErr:     true,
		},
		{
			cmd:   "wr0n6",
			isErr: true,
		},
		{
			cmd:   "wr0n6",
			tty:   true,
			isErr: true,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("RunCmd(%s,%t,%t)", input.cmd, input.getOutput, input.tty)

		got, err := runCmd(input.cmd, input.getOutput, input.tty)
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

func TestRunBashCmd(t *testing.T) {
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

		got, err := RunBashCmd(input.cmd, input.getOutput)
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
