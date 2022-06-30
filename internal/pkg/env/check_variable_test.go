package env

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/NethermindEth/1click/internal/utils"
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
			"Test case 1, mainnet, no TTD",
			ReCONFIG,
			"mainnet",
			"execution",
			"nethermind",
			false,
			false,
		},
		{
			"Test case 2, invalid network, error",
			ReCONFIG,
			"testnet",
			"consensus",
			"teku",
			false,
			true,
		},
		{
			"Test case 3, invalid clientType, error",
			ReCONFIG,
			"mainnet",
			"test",
			"test",
			false,
			true,
		},
		{
			"Test case 4, invalid client, error",
			ReCONFIG,
			"mainnet",
			"consensus",
			"test",
			false,
			true,
		},
		{
			"Test case 5, mainnet, no prysm config",
			ReCONFIG,
			"mainnet",
			"execution",
			"nethermind",
			false,
			false,
		},
		{
			"Test case 6, kiln, no prysm config in nethermind",
			ReCONFIG,
			"kiln",
			"execution",
			"nethermind",
			false,
			false,
		},
		{
			"Test case 7, kiln, prysm config, consensus",
			ReCONFIG,
			"kiln",
			"consensus",
			"prysm",
			true,
			false,
		},
		{
			"Test case 8, kiln, prysm config, validator",
			ReCONFIG,
			"kiln",
			"validator",
			"prysm",
			true,
			false,
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

func TestCheckVariableBase(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name    string
		regex   *regexp.Regexp
		network string
		want    bool
		isErr   bool
	}{
		{
			"Test case 1, mainnet, no TTD",
			ReTTD,
			"mainnet",
			false,
			false,
		},
		{
			"Test case 2, invalid network, error",
			ReTTD,
			"testnet",
			false,
			true,
		},
		{
			"Test case 3, kiln, TTD",
			ReTTD,
			"kiln",
			true,
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CheckVariableBase(tc.regex, tc.network)

			descr := fmt.Sprintf("CheckVariableBase(re, %s) with regex %v", tc.network, tc.regex)
			if err = utils.CheckErr(descr, tc.isErr, err); err != nil {
				t.Error(err)
			}

			if tc.want != got {
				t.Errorf("Expected %v, got %v. Function call: %s", tc.want, got, descr)
			}
		})
	}
}
