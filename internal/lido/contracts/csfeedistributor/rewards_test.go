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
	"reflect"
	"testing"

	bond "github.com/NethermindEth/sedge/internal/lido/contracts/csaccounting"
)

func TestRewards(t *testing.T) {
	log.SetOutput(io.Discard)
	tcs := []struct {
		name         string
		network      string
		nodeID       *big.Int
		proofStrings []string
		wantErr      bool
	}{
		{
			name:    "Rewards for nodeID 1, Holesky",
			network: "holesky",
			nodeID:  big.NewInt(1),
			proofStrings: []string{
				"0x6d20abd8b1d6bd278c57e818ff4984c023509ff6b5ad82d14b04725b19bcc70c",
				"0xed6bfbd3f1755e450ea146c18bff02dd4ca5e6a2c8a354ba3fa50256294175d6",
				"0x778ee311998c39a6d013c940e5a4f882adbe565d7656ed01a69b158416f94275",
				"0x75f2f37a15fc968b78f6fa2f62e2e2c5f77040564a86e437910a22576afbdb5f",
				"0x2f524e0fdbcf86b1cea60e6939811f5b8d262b33f1bffa1f2956d1228e43ea1c",
				"0x4bba91a73d5d1974dbb75b61622306ba3742af7dc2a15d3d600473a7a06df426",
				"0xc5cb054e6f48e9f610c115c212cd2c0125009bf7e63b6c472f2b2262c0868895",
			},
			wantErr: false,
		},
		{
			name:    "Rewards for nodeID 113, Holesky",
			network: "holesky",
			nodeID:  big.NewInt(113),
			proofStrings: []string{
				"0xbc078e2dffa88e24d5750eb8d5f4cc1fe4d24c582d4341fbde368ca793b3df5a",
				"0x162abf1ad8a2615e339cb77a75555b4dcea66d708bf77e4c4e5a5afe1dbfa48c",
				"0x0290ed7853afaceb04d7a7e568aa5460ea00503b713f3672aabae6d9ca9f5f0b",
				"0xda8b6b6e6c9367cc58ccf7e2888d899921447781c64605eba898c78367e537a5",
				"0x1bf82b211eb3ef2cb5703f985d9a17428dda63dafb993fbdaeec4b4f4db28632",
				"0x68b6269fbb92d54881088f6f4c83ae6dbff9d6005602f5338a07805a7764e321",
				"0x88d938487dd1dbe31cd9cb3acdb96d786f9b4002d52a55650a48605cc1d90d54",
			},
			wantErr: false,
		},
		{
			name:    "Invalid proof for nodeID 11, Holesky",
			network: "holesky",
			nodeID:  big.NewInt(11),
			proofStrings: []string{
				"0x16e3715ab4eaf054d8ee629bb5cb2c916a8fbb4a2fbcaca78374d8c776617999",
				"0x44fe6c1164f8f937a319fd6e120da4ac0a1e4c060af442df230be65b14210a61",
				"0x9402cbac957071d74d44f34897c53992fce48e875f45e782c0427ae72d7b5a3e",
				"0x43abe2721fdd9122dd0ae484cfe8aadf04bf18b88ef24fc787a11066ae0e7959",
			},
			wantErr: true,
		},
		{
			name:    "Invalid nodeID, Holesky",
			network: "holesky",
			nodeID:  big.NewInt(-3),
			proofStrings: []string{
				"0x6d20abd8b1d6bd278c57e818ff4984c023509ff6b5ad82d14b04725b19bcc70c",
				"0xed6bfbd3f1755e450ea146c18bff02dd4ca5e6a2c8a354ba3fa50256294175d6",
				"0x778ee311998c39a6d013c940e5a4f882adbe565d7656ed01a69b158416f94275",
				"0x75f2f37a15fc968b78f6fa2f62e2e2c5f77040564a86e437910a22576afbdb5f",
				"0x2f524e0fdbcf86b1cea60e6939811f5b8d262b33f1bffa1f2956d1228e43ea1c",
				"0x4bba91a73d5d1974dbb75b61622306ba3742af7dc2a15d3d600473a7a06df426",
				"0xc5cb054e6f48e9f610c115c212cd2c0125009bf7e63b6c472f2b2262c0868895",
			},
			wantErr: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			treeCID, err := treeCID(tc.network)
			if err != nil {
				t.Fatalf("failed to call treeCID: %v", err)
			}

			rewards, err := Rewards(tc.network, tc.nodeID, tc.proofStrings)
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

func TestConvertProofToByte(t *testing.T) {
	proofStrings := []string{
		"0x6d20abd8b1d6bd278c57e818ff4984c023509ff6b5ad82d14b04725b19bcc70c",
		"0xed6bfbd3f1755e450ea146c18bff02dd4ca5e6a2c8a354ba3fa50256294175d6",
		"0x778ee311998c39a6d013c940e5a4f882adbe565d7656ed01a69b158416f94275",
		"0x75f2f37a15fc968b78f6fa2f62e2e2c5f77040564a86e437910a22576afbdb5f",
		"0x2f524e0fdbcf86b1cea60e6939811f5b8d262b33f1bffa1f2956d1228e43ea1c",
		"0x4bba91a73d5d1974dbb75b61622306ba3742af7dc2a15d3d600473a7a06df426",
		"0xc5cb054e6f48e9f610c115c212cd2c0125009bf7e63b6c472f2b2262c0868895",
	}

	expectedProofChunks := [][32]byte{
		{109, 32, 171, 216, 177, 214, 189, 39, 140, 87, 232, 24, 255, 73, 132, 192, 35, 80, 159, 246, 181, 173, 130, 209, 75, 4, 114, 91, 25, 188, 199, 12},
		{237, 107, 251, 211, 241, 117, 94, 69, 14, 161, 70, 193, 139, 255, 2, 221, 76, 165, 230, 162, 200, 163, 84, 186, 63, 165, 2, 86, 41, 65, 117, 214},
		{119, 142, 227, 17, 153, 140, 57, 166, 208, 19, 201, 64, 229, 164, 248, 130, 173, 190, 86, 93, 118, 86, 237, 1, 166, 155, 21, 132, 22, 249, 66, 117},
		{117, 242, 243, 122, 21, 252, 150, 139, 120, 246, 250, 47, 98, 226, 226, 197, 247, 112, 64, 86, 74, 134, 228, 55, 145, 10, 34, 87, 106, 251, 219, 95},
		{47, 82, 78, 15, 219, 207, 134, 177, 206, 166, 14, 105, 57, 129, 31, 91, 141, 38, 43, 51, 241, 191, 250, 31, 41, 86, 209, 34, 142, 67, 234, 28},
		{75, 186, 145, 167, 61, 93, 25, 116, 219, 183, 91, 97, 98, 35, 6, 186, 55, 66, 175, 125, 194, 161, 93, 61, 96, 4, 115, 167, 160, 109, 244, 38},
		{197, 203, 5, 78, 111, 72, 233, 246, 16, 193, 21, 194, 18, 205, 44, 1, 37, 0, 155, 247, 230, 59, 108, 71, 47, 43, 34, 98, 192, 134, 136, 149},
	}

	actualProofChunks, err := convertProofToByte(proofStrings)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actualProofChunks, expectedProofChunks) {
		t.Errorf("unexpected proof chunks\nexpected: %v\nactual:   %v", expectedProofChunks, actualProofChunks)
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
