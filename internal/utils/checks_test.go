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

func TestCheckContainers(t *testing.T) {
	//TODO: fix problems with sudo and containers
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
		descr := fmt.Sprintf("CheckContainers(%s)", input.path)
		_, err := CheckContainers(input.path)
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
