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
