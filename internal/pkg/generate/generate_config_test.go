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
package generate

import (
	"fmt"
	"testing"
)

func validateGeneratedConfig(t *testing.T, path string) {
	//TODO: validate generated config
}

func TestGenerateConfig(t *testing.T) {
	inputs := [...]struct {
		path  string
		isErr bool
	}{
		{t.TempDir(), false},
		{"", true},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("GenerateScripts(%s,)", input.path)

		if err := GenerateConfig(input.path); input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
		validateGeneratedConfig(t, input.path)
	}
}
