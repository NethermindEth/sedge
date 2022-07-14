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
package configs

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

type getConfigClientsTestCase struct {
	queryClientType    string
	configClientsTypes map[string][]string
	resultClients      []string
	isErr              bool
}

func prepareGetConfigClientsTestCase(testCase getConfigClientsTestCase) {
	for key, value := range testCase.configClientsTypes {
		viper.Set(key, value)
	}
}

func cleanGetConfigClientsTestCase(_ getConfigClientsTestCase) {
	viper.Reset()
}

func cleanAll() {
	viper.Reset()
}

func verifyGetConfigClientsResult(t *testing.T, resultClients []string, input getConfigClientsTestCase, descr string) {
	for _, client := range resultClients {
		contained := false
		for _, expected := range input.resultClients {
			if client == expected {
				contained = true
				break
			}
		}
		if !contained {
			t.Logf("using config: %s", input.configClientsTypes)
			t.Errorf("%s expected %s but got: %s", descr, input.resultClients, resultClients)
		}
	}
}

func TestGetConfigClients(t *testing.T) {
	inputs := [...]getConfigClientsTestCase{
		{"execution", map[string][]string{"executionClients": {"a", "b"}, "consensusClients": {"c"}, "validatorClients": {"d", "e"}}, []string{"a", "b"}, false},
		{"all", map[string][]string{"executionClients": {"a"}, "consensusClients": {"b"}, "validatorClients": {"c"}}, []string{}, true},
		{"consensusClients", map[string][]string{"executionClients": {"a"}, "consensusClients": {"b"}, "validatorClients": {"c"}}, []string{}, true},
	}
	t.Cleanup(cleanAll)

	for _, input := range inputs {
		prepareGetConfigClientsTestCase(input)
		descr := fmt.Sprintf("GetConfigClients(%s)", input.queryClientType)

		if res, err := ConfigClients(input.queryClientType); input.isErr && err == nil {
			t.Logf("using config: %s", input.configClientsTypes)
			t.Errorf("%s expected to fail.", descr)
		} else if !input.isErr {
			if err != nil {
				t.Logf("using config: %s", input.configClientsTypes)
				t.Errorf("%s failed: %v", descr, err)
			} else {
				verifyGetConfigClientsResult(t, res, input, descr)
			}
		}

		cleanGetConfigClientsTestCase(input)
	}
}
