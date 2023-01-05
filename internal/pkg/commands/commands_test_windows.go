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
	"fmt"
	"testing"
	"text/template"
)

func TestRunCmd(t *testing.T) {
	inputs := []struct {
		cmd       string
		getOutput bool
		output    string
		isErr     bool
	}{
		{
			cmd:       "echo 'hello world'",
			getOutput: true,
			output:    "hello world\n",
			isErr:     false,
		},
		{
			cmd:       "wr0n6",
			getOutput: true,
			isErr:     true,
		},
	}

	runner := NewCMDRunner(CMDRunnerOptions{
		RunAsAdmin: false,
	})

	for _, input := range inputs {
		descr := fmt.Sprintf("RunCmd(%s,%t)", input.cmd, input.getOutput)

		got, err := runner.RunCMD(Command{
			Cmd:       input.cmd,
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

func TestRunBashScript(t *testing.T) {
	inputs := []struct {
		cmd       string
		getOutput bool
		output    string
		isErr     bool
	}{
		{
			cmd:       "echo 'hello world'",
			getOutput: true,
			output:    "hello world\n",
			isErr:     false,
		},
		{
			cmd:       "wr0n6",
			getOutput: true,
			isErr:     true,
		},
	}

	runner := NewCMDRunner(CMDRunnerOptions{
		RunAsAdmin: false,
	})

	for _, input := range inputs {
		descr := fmt.Sprintf("RunBashCmd(%s,%t)", input.cmd, input.getOutput)

		tmp, err := template.New("script").Parse(string(input.cmd))
		if err != nil {
			t.Fatalf("Unexpected error at case %q: %v", input.cmd, err)
		}

		got, err := runner.RunScript(ScriptFile{
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
