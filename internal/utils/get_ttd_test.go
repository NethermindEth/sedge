package utils

import (
	"fmt"
	"testing"
)

func TestTTD(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name       string
		network    string
		clientType string
		client     string
		want       bool
		isErr      bool
	}{
		{
			"Test case 1, mainnet, no TTD",
			"mainnet",
			"execution",
			"nethermind",
			false,
			false,
		},
		{
			"Test case 2, invalid network, error",
			"testnet",
			"consensus",
			"teku",
			false,
			true,
		},
		{
			"Test case 3, invalid clientType, error",
			"mainnet",
			"test",
			"test",
			false,
			true,
		},
		{
			"Test case 4, invalid client, error",
			"mainnet",
			"consensus",
			"test",
			false,
			true,
		},
		{
			"Test case 5, kiln, TTD in nethermind",
			"kiln",
			"execution",
			"nethermind",
			true,
			false,
		},
		{
			"Test case 6, kiln, TTD in lighthouse",
			"kiln",
			"consensus",
			"lighthouse",
			true,
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := TTD(tc.network, tc.clientType, tc.client)

			descr := fmt.Sprintf("TTD(%s, %s, %s)", tc.network, tc.clientType, tc.client)
			if err = CheckErr(descr, tc.isErr, err); err != nil {
				t.Error(err)
			}

			if tc.want != got {
				t.Errorf("Expected %v, got %v. Function call: %s", tc.want, got, descr)
			}
		})
	}
}
