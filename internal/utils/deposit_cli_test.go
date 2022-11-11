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
package utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
)

type generateValidatorKeyTestCase struct {
	name     string
	runner   commands.CommandRunner
	existing bool
	network  string
	path     string
	isErr    bool
}

func TestGenerateValidatorKey(t *testing.T) {
	tcs := []generateValidatorKeyTestCase{
		{
			name: "Test case 1",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					return "", nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			},
			existing: false,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    false,
		},
		{
			name: "Test case 2",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "", fmt.Errorf("unexpected error")
					}
					return "", nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			},
			existing: true,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    true,
		},
		{
			name: "Test case 3",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					return "", nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			},
			existing: true,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    false,
		},
		{
			name: "Test case 4",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					return "", nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			},
			existing: false,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    false,
		},
		{
			name: "Test case 5",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					if strings.Contains(c.Cmd, "docker pull") {
						return "", fmt.Errorf("error")
					}
					return "", nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			},
			existing: true,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    true,
		},
		{
			name: "Test case 6",
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					if strings.Contains(c.Cmd, "docker pull") {
						return "", fmt.Errorf("error")
					}
					return "", nil
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			},
			existing: false,
			network:  "mainnet",
			path:     t.TempDir(),
			isErr:    true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			commands.InitRunner(func() commands.CommandRunner {
				return tc.runner
			})
			descr := fmt.Sprintf("GenerateValidatorKey(%t, %s, %s, %s)", tc.existing, tc.network, tc.path, "password")

			data := ValidatorKeyData{
				Existing:              tc.existing,
				Network:               tc.network,
				Path:                  tc.path,
				Password:              "password",
				Eth1WithdrawalAddress: "0x5c00ABEf07604C59Ac72E859E5F93D5ab8546F83",
			}
			err := GenerateValidatorKey(data)
			if tc.isErr && err == nil {
				t.Errorf("%s expected to fail.", descr)
			} else if !tc.isErr && err != nil {
				t.Errorf("%s failed: %v", descr, err)
			}
		})
	}
}
