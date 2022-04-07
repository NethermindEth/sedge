package utils

import (
	"fmt"
	"regexp"
	"runtime"
	"testing"
)

func TestGetScriptPath(t *testing.T) {
	inputs := []struct {
		dependency string
		isErr      bool
	}{
		{
			"",
			false,
		},
		{
			"docker",
			false,
		},
		{
			"wR0n9",
			false,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("getScriptPath(%s)", input.dependency)

		path, _, err := getScriptPath(input.dependency)
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail.", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else {
				//TODO: improve results tests
				match, err := regexp.MatchString(`.*`+input.dependency+`/.*/`+`\.sh`, path)
				if err != nil || match {
					t.Errorf("returned path %s is invalid", path)
				}

				//TODO: validate distro
			}
		}

	}
}

func TestDependencySupported(t *testing.T) {
	if runtime.GOOS != "linux" { //TODO: update when other os's are supported
		t.Skipf("this os is not supported")
	}

	inputs := []struct {
		dependency string
		result     bool
	}{
		{
			"",
			false,
		},
		{
			"docker",
			true,
		},
		{
			"wR0n9",
			false,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("dependencySupported(%s)", input.dependency)
		got := dependencySupported(input.dependency)
		if got != input.result {
			t.Errorf("%s expected %t but got %t", descr, input.result, got)
		}
	}
}

func TestHandleInstructions(t *testing.T) {
	inputs := []struct {
		dependencies  []string
		handlerResult error
		isErr         bool
	}{
		{
			[]string{},
			nil,
			false,
		},
		{
			[]string{"docker", "docker-compose"},
			nil,
			false,
		},
		{
			[]string{"docker", "docker-compose"},
			fmt.Errorf("unexpected error"),
			true,
		},
		{
			[]string{"docker", "wR0n9"},
			nil,
			true,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("HandlerInstruction(%s,...)", input.dependencies)
		err := HandleInstructions(input.dependencies, func(dependency string) error {
			contained := false
			for _, expected := range input.dependencies {
				if dependency == expected {
					contained = true
				}
			}
			if !contained {
				t.Errorf("%s handler called on unnexpected dependency %s", descr, dependency)
			}

			t.Logf("%s handler returning error %v", descr, input.handlerResult)
			return input.handlerResult
		})
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}

func TestShowInstructions(t *testing.T) {
	inputs := []struct {
		dependency string
		isErr      bool
	}{
		{
			"",
			true,
		},
		{
			"docker",
			false,
		},
		{
			"wR0n9",
			true,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("ShowInstructions(%s)", input.dependency)

		if err := ShowInstructions(input.dependency); input.isErr && err == nil {
			t.Errorf("%s expected to fail.", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}

func TestInstallDependency(t *testing.T) {
	//TODO: test dependency installing
}
