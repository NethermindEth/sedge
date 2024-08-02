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
package csaccounting

import (
	"io"
	"log"
	"math/big"
	"testing"
)

func TestBondSummary(t *testing.T) {
	log.SetOutput(io.Discard)
	tcs := []struct {
		name      string
		network   string
		nodeID    *big.Int
		invalidId bool
	}{
		{
			"BondSummary with valid ID, Holesky #1", "holesky", big.NewInt(2), false,
		},
		{
			"BondSummary with valid ID, Holesky #2", "holesky", big.NewInt(14), false,
		},
		{
			"BondSummary with invalid ID, Holesky ", "holesky", big.NewInt(-6), true,
		},
	}
	var expectedExcess, expectedMissed *big.Int
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			bond, err := BondSummary(tc.network, tc.nodeID)
			if err != nil {
				t.Fatalf("failed to call BondsSummary: %v", err)
			}
			if tc.invalidId && bond.Current.Sign() > 0 {
				t.Errorf("invalid current bond amount, expected %v, got: %v", big.NewInt(0), bond.Current)
			}
			if tc.invalidId && bond.Required.Sign() > 0 {
				t.Errorf("invalid required bond amount, expected %v, got: %v", big.NewInt(0), bond.Required)
			}

			// if current amount is greater than required
			if bond.Current.Cmp(bond.Required) == 1 {
				expectedExcess = new(big.Int).Sub(bond.Current, bond.Required)
				expectedMissed = big.NewInt(0)
			} else if bond.Current.Cmp(bond.Required) == -1 {
				expectedExcess = big.NewInt(0)
				expectedMissed = new(big.Int).Sub(bond.Required, bond.Current)
			} else {
				expectedExcess = big.NewInt(0)
				expectedMissed = big.NewInt(0)
			}

			if bond.Excess.Cmp(expectedExcess) != 0 {
				t.Errorf("not same excess bond amount, expected %v, got: %v", expectedExcess, bond.Excess)
			}
			if bond.Missed.Cmp(expectedMissed) != 0 {
				t.Errorf("not same missed bond amount, expected %v, got: %v", expectedMissed, bond.Missed)
			}
		})
	}
}
