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
	"testing"
)

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
