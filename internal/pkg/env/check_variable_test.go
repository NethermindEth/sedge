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
package env

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/NethermindEth/sedge/internal/utils"
)

func TestCheckVariable(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name       string
		regex      *regexp.Regexp
		network    string
		clientType string
		client     string
		want       bool
		isErr      bool
	}{
		{
			"Test case 2, invalid network, error",
			ReMEV,
			"testnet",
			"consensus",
			"teku",
			false,
			true,
		},
		{
			"Test case 3, invalid clientType, error",
			ReMEV,
			"mainnet",
			"test",
			"test",
			false,
			true,
		},
		{
			"Test case 4, invalid client, error",
			ReMEV,
			"mainnet",
			"consensus",
			"test",
			false,
			true,
		},
		{
			"Test case 5, mainnet, no prysm config",
			ReMEV,
			"mainnet",
			"execution",
			"nethermind",
			false,
			false,
		},
		{
			"Test case 6, goerli, no prysm config in nethermind",
			ReMEV,
			"goerli",
			"execution",
			"nethermind",
			false,
			false,
		},
		{
			"Test case 7, chiado, prysm config, no consensus config",
			ReMEV,
			"chiado",
			"consensus",
			"prysm",
			false,
			true,
		},
		{
			"Test case 8, chiado, prysm config, no validator config",
			ReMEV,
			"chiado",
			"validator",
			"prysm",
			false,
			true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CheckVariable(tc.regex, tc.network, tc.clientType, tc.client)

			descr := fmt.Sprintf("CheckVariable(re, %s, %s, %s) with regex %v", tc.network, tc.clientType, tc.client, tc.regex)
			if err = utils.CheckErr(descr, tc.isErr, err); err != nil {
				t.Error(err)
			}

			if tc.want != got {
				t.Errorf("Expected %v, got %v. Function call: %s", tc.want, got, descr)
			}
		})
	}
}
