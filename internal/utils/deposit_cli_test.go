package utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/test"
)

type generateValidatorKeyTestCase struct {
	runner   commands.CommandRunner
	existing bool
	network  string
	path     string
	isErr    bool
}

func TestGenerateValidatorKey(t *testing.T) {
	tcs := []generateValidatorKeyTestCase{
		{
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
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					if strings.Contains(c.Cmd, "docker build") {
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
			runner: &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					if strings.Contains(c.Cmd, "inspect") {
						return "No such object: image", fmt.Errorf("error")
					}
					if strings.Contains(c.Cmd, "docker build") {
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
		commands.InitRunner(func() commands.CommandRunner {
			return tc.runner
		})
		descr := fmt.Sprintf("GenerateValidatorKey(%t, %s, %s)", tc.existing, tc.network, tc.path)

		err := GenerateValidatorKey(tc.existing, tc.network, tc.path)
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail.", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
