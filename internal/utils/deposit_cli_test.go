package utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/test"
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
				SRunBash: func(bs commands.BashScript) (string, error) {
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
				SRunBash: func(bs commands.BashScript) (string, error) {
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
				SRunBash: func(bs commands.BashScript) (string, error) {
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
				SRunBash: func(bs commands.BashScript) (string, error) {
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
				SRunBash: func(bs commands.BashScript) (string, error) {
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
				SRunBash: func(bs commands.BashScript) (string, error) {
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

			err := GenerateValidatorKey(tc.existing, tc.network, tc.path, "password")
			if tc.isErr && err == nil {
				t.Errorf("%s expected to fail.", descr)
			} else if !tc.isErr && err != nil {
				t.Errorf("%s failed: %v", descr, err)
			}
		})
	}
}
