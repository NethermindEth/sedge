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
package csfeedistributor

import (
	"io"
	"log"
	"math/big"
	"testing"

	bond "github.com/NethermindEth/sedge/internal/lido/contracts/csaccounting"
)

func TestRewards(t *testing.T) {
	log.SetOutput(io.Discard)
	tcs := []struct {
		name    string
		network string
		nodeID  *big.Int
		wantErr bool
	}{
		{
			name:    "Rewards for nodeID 1, Mainnet",
			network: "mainnet",
			nodeID:  big.NewInt(1),
			wantErr: false,
		},
		{
			name:    "Invalid nodeID, Mainnet",
			network: "mainnet",
			nodeID:  big.NewInt(-15),
			wantErr: true,
		},
		{
			name:    "Rewards for nodeID 1, Mainnet",
			network: "mainnet",
			nodeID:  big.NewInt(1),
			wantErr: false,
		},
		{
			name:    "Invalid nodeID, Hoodi",
			network: "hoodi",
			nodeID:  big.NewInt(-20),
			wantErr: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			treeCID, err := treeCID(tc.network)
			if err != nil {
				t.Fatalf("failed to call treeCID: %v", err)
			}

			rewards, err := Rewards(tc.network, tc.nodeID)
			if err != nil && !tc.wantErr {
				t.Fatalf("failed to call Rewards: %v", err)
			}

			fees, err := cumulativeFeeShares(treeCID, tc.nodeID)
			if err != nil && !tc.wantErr {
				t.Fatalf("failed to call cumulativeFeeShares: %v", err)
			}

			bond, err := bond.BondSummary(tc.network, tc.nodeID)
			if err != nil && !tc.wantErr {
				t.Fatalf("failed to call BondSummary: %v", err)
			}

			if rewards == nil && tc.wantErr {
				t.Skipf("Expected nil value for rewards")
			} else if rewards == nil && !tc.wantErr {
				t.Fatalf("invalid rewards value: expected a value, got nil")
			}
			expectedRewards := new(big.Int).Add(bond.Excess, fees)
			if rewards.Cmp(expectedRewards) != 0 {
				t.Errorf("invalid rewards amount, expected %v, got: %v", expectedRewards, rewards)
			}
		})
	}
}

func TestConvertTreeValuesToBigInt(t *testing.T) {
	tcs := []struct {
		input    interface{}
		expected *big.Int
	}{
		{12345.0, big.NewInt(12345)},
		{67890.0, big.NewInt(67890)},
		{0.0, big.NewInt(0)},
	}

	for _, tc := range tcs {
		result, err := convertTreeValuesToBigInt(tc.input)
		if err != nil {
			t.Errorf("convertTreeValuesToBigInt(%v) returned error: %v", tc.input, err)
		}

		if result.Cmp(tc.expected) != 0 {
			t.Errorf("convertTreeValuesToBigInt(%v) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}

func TestGetSharesFromReports(t *testing.T) {
	// Create test data with actual IPFS structure
	testOperators := map[string]Operator{
		"0": {
			DistributedRewards: big.NewInt(51439079431843106),
			PerformanceCoefficients: PerformanceCoefficients{
				AttestationsWeight: 54,
				BlocksWeight:       4,
				SyncWeight:         2,
			},
			Validators: map[string]Validator{
				"1208835": {
					AttestationDuty: Duty{
						Assigned: 1575,
						Included: 1565,
					},
					DistributedRewards: big.NewInt(1435293025283508),
					Performance:        0.9936507936507937,
					ProposalDuty: Duty{
						Assigned: 0,
						Included: 0,
					},
					RewardsShare: 1.0,
					Slashed:      false,
					Strikes:      0,
					SyncDuty: Duty{
						Assigned: 0,
						Included: 0,
					},
					Threshold: 0.8776934372451525,
				},
			},
		},
		"1": {
			DistributedRewards: big.NewInt(0),
			PerformanceCoefficients: PerformanceCoefficients{
				AttestationsWeight: 54,
				BlocksWeight:       8,
				SyncWeight:         2,
			},
			Validators: map[string]Validator{
				"1232516": {
					AttestationDuty: Duty{
						Assigned: 115,
						Included: 0,
					},
					DistributedRewards: big.NewInt(0),
					Performance:        0.0,
					ProposalDuty: Duty{
						Assigned: 0,
						Included: 0,
					},
					RewardsShare: 0.5834,
					Slashed:      false,
					Strikes:      1,
					SyncDuty: Duty{
						Assigned: 0,
						Included: 0,
					},
					Threshold: 0.8976934372451525,
				},
			},
		},
	}

	newTreeData := &NewTreeData{
		Blockstamp: Blockstamp{
			BlockHash:      "0xa9f6b5644560aadf834f04398ef0a4dcbf0567354e0e3f63f236d66b2a78ca42",
			BlockNumber:    1431724,
			BlockTimestamp: 1760693388,
			RefEpoch:       48124,
			RefSlot:        1539999,
			SlotNumber:     1539999,
			StateRoot:      "0x3713312e3a66235f0365d068d607b6048156c4e6cf6de276127710566c8635d2",
		},
		Distributable:      big.NewInt(907586356879081627),
		DistributedRewards: big.NewInt(610421466330066286),
		Frame:              []int64{46550, 48124},
		Operators:          testOperators,
		RebateToProtocol:   big.NewInt(297164890549015341),
	}

	tcs := []struct {
		name     string
		nodeID   *big.Int
		expected *big.Int
		wantErr  bool
	}{
		{
			name:     "Valid nodeID 0",
			nodeID:   big.NewInt(0),
			expected: big.NewInt(51439079431843106),
			wantErr:  false,
		},
		{
			name:     "Valid nodeID 1 with zero rewards",
			nodeID:   big.NewInt(1),
			expected: big.NewInt(0),
			wantErr:  false,
		},
		{
			name:     "Invalid nodeID",
			nodeID:   big.NewInt(999),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getSharesFromReports(newTreeData, tc.nodeID)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.Cmp(tc.expected) != 0 {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
