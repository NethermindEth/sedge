package utils

import (
	"fmt"
	"testing"
)

func validatePending(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for _, dep1 := range got {
		found := false
		for _, dep2 := range want {
			if dep1 == dep2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestCheckDependencies(t *testing.T) {
	inputs := [...]struct {
		dependencies []string
		pending      []string
	}{
		{},
		{
			[]string{"wr0n9"},
			[]string{"wr0n9"},
		},
		{
			[]string{""},
			[]string{""},
		},
		{
			[]string{"curl"},
			[]string{},
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("CheckDependencies(%s)", input.dependencies)
		got := CheckDependencies(input.dependencies)
		if !validatePending(got, input.pending) {
			t.Errorf("%s expected %s but got %s", descr, input.pending, got)
		}

	}
}

func TestPreCheck(t *testing.T) {
	//TODO: fix problems with sudo
	inputs := [...]struct {
		path  string
		isErr bool
	}{
		// {
		// 	t.TempDir(),
		// 	false,
		// },
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("PreCheck(%s)", input.path)
		err := PreCheck(input.path)
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}

type checkContainersTC struct {
	path     string
	runner   commands.CommandRunner
	isErr    bool
	psRunned int
}

func buildCheckContainersTestCase(t *testing.T, caseName string, isErr bool) *checkContainersTC {
	dcPath := t.TempDir()
	test.PrepareTestCaseDir(filepath.Join("testdata", "checks_tests", caseName, "docker-compose-scripts"), dcPath)

	tc := checkContainersTC{}
	tc.path = dcPath
	tc.runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			if strings.Contains(c.Cmd, "docker-compose") && strings.Contains(c.Cmd, "ps") {
				tc.psRunned += 1
				_, err := os.Lstat(filepath.Join(dcPath, configs.DefaultDockerComposeScriptName))
				return "", err
			}
			return "", nil
		},
		SRunBash: func(bs commands.BashScript) (string, error) {
			return "", nil
		},
	}
	tc.isErr = isErr
	return &tc
}

func TestCheckContainers(t *testing.T) {
	tcs := [...]*checkContainersTC{
		buildCheckContainersTestCase(
			t,
			"case_1",
			false,
		),
		buildCheckContainersTestCase(
			t,
			"case_2",
			true,
		),
	}

	for _, tc := range tcs {
		descr := fmt.Sprintf("CheckContainers(%s)", tc.path)

		commands.InitRunner(func() commands.CommandRunner {
			return tc.runner
		})

		_, err := CheckContainers(tc.path)
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if tc.psRunned < 1 {
				t.Errorf("%s didn't run docker-compose ps", descr)
			}
		}
	}
}
