package utils

import (
	"fmt"
	"testing"
)

type zipTestCase struct {
	lists [][]string
	isErr bool
	out   [][]string
}

func evalZipTestCase(t *testing.T, tc zipTestCase) {
	res, err := ZipString(tc.lists...)
	descr := fmt.Sprintf("ZipString(%s)", tc.lists)

	if !tc.isErr && err != nil {
		t.Errorf("%s failed: %v", descr, err)
	} else if tc.isErr && err == nil {
		t.Errorf("%s expected to fail", descr)
	}

Loop:
	for x, list := range res {
		for y, got := range list {
			if x > len(tc.out) || y > len(tc.out[x]) || got != tc.out[x][y] {
				t.Errorf("%s expected %s but got %s", descr, tc.out, res)
				break Loop
			}
		}
	}

}

func TestZipStrings(t *testing.T) {
	inputs := [...]zipTestCase{
		{},
		{[][]string{{"a", "a", "a"}, {"b", "b", "b"}}, false, [][]string{{"a", "b"}, {"a", "b"}, {"a", "b"}}},
		{[][]string{{"a", "a"}, {"b", "b"}, {"c", "c"}}, false, [][]string{{"a", "b", "c"}, {"a", "b", "c"}}},
		{lists: [][]string{{}, {"a"}}, isErr: true},
	}

	for _, tc := range inputs {
		evalZipTestCase(t, tc)
	}
}
