package utils

import (
	"fmt"
	"testing"
)

func TestSkipLines(t *testing.T) {
	inputs := []struct {
		content string
		symbol  string
		result  string
	}{
		{},
		{"bbbb", "a", "bbbb"},
		{"aaaa", "a", ""},
		{"bbbb\naaaa", "a", "bbbb"},
		{"aaaa\nbbbb", "a", "bbbb"},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("SkipLines(%s, %s)", input.content, input.symbol)
		got := SkipLines(input.content, input.symbol)

		if got != input.result {
			t.Errorf("%s expected %s but got %s", descr, input.result, got)
		}
	}
}

func TestContains(t *testing.T) {
	inputs := [...]struct {
		list []string
		str  string
		want bool
	}{
		{},
		{[]string{"a", "A", "b", "B"}, "b", true},
		{[]string{"a", "A", "b", "B"}, "c", false},
		{[]string{}, "a", false},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("Contains(%s)", input.list)
		if got := Contains(input.list, input.str); got != input.want {
			t.Errorf("%s expected %t but got %t", descr, input.want, got)
		}
	}
}

func TestContainsOnly(t *testing.T) {
	inputs := [...]struct {
		list   []string
		target []string
		want   bool
	}{
		{[]string{}, []string{}, true},
		{[]string{}, []string{"a"}, true},
		{[]string{"b"}, []string{"a", "A", "b", "B"}, true},
		{[]string{"c"}, []string{"a", "A", "b", "B"}, false},
		{[]string{"a"}, []string{}, false},
		{[]string{"execution", "validator"}, []string{"execution", "consensus"}, false},
		{[]string{"execution", "consensus", "validator"}, []string{"execution", "consensus"}, false},
		{[]string{"execution", "validator"}, []string{"execution", "consensus", "validator"}, true},
		{[]string{"execution", "consensus", "validator"}, []string{"execution", "consensus", "validator"}, true},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("Contains(%s)", input.list)
		if got := ContainsOnly(input.list, input.target); got != input.want {
			t.Errorf("%s expected %t but got %t", descr, input.want, got)
		}
	}
}

func TestIsAddress(t *testing.T) {
	tcs := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"2131", false},
		{"dasd31gsd1231", false},
		{"0x2312313aaef2312312", false},
		{"0x5c00ABEf07604C59Ac72E859E5F93D5abZXCVF83", false},
		{"5c00ABEf07604C59Ac72E859E5F93D5ab8546F83", false},
		{"0x5c00ABEf07604C59Ac72E859E5F93D5ab8546F83", true},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprintf("IsAddress(%s)", tc.input), func(t *testing.T) {
			if got := IsAddress(tc.input); got != tc.want {
				t.Errorf("got != want. Expected %v, got %v", tc.want, tc.input)
			}
		})
	}
}
