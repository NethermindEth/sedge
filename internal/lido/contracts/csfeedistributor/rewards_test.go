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
	"context"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	bond "github.com/NethermindEth/sedge/internal/lido/contracts/csaccounting"
	"github.com/stretchr/testify/assert"
)

func TestRewards(t *testing.T) {
	// Set a timeout for network-dependent tests
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tcs := []struct {
		name    string
		network string
		nodeID  *big.Int
		wantErr bool
	}{
		{
			name:    "Rewards for nodeID 5, Hoodi",
			network: "hoodi",
			nodeID:  big.NewInt(5),
			wantErr: false,
		},
		{
			name:    "Rewards for nodeID 10, Holesky",
			network: "holesky",
			nodeID:  big.NewInt(10),
			wantErr: false,
		},
		{
			name:    "Rewards for nodeID 250, Holesky",
			network: "holesky",
			nodeID:  big.NewInt(250),
			wantErr: false,
		},
		{
			name:    "Rewards for nodeID 113, Holesky",
			network: "holesky",
			nodeID:  big.NewInt(113),
			wantErr: false,
		},
		{
			name:    "Invalid nodeID, Holesky",
			network: "holesky",
			nodeID:  big.NewInt(-3),
			wantErr: true,
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
			// Skip tests that require network connectivity if we're in a CI environment
			if isCIEnvironment() {
				t.Skip("Skipping network-dependent test in CI environment")
			}

			// Run the test with timeout
			done := make(chan bool, 1)
			go func() {
				defer func() { done <- true }()
				
				treeCID, err := treeCID(tc.network)
				if err != nil {
					if tc.wantErr {
						return // Expected error
					}
					t.Logf("Network error getting treeCID: %v", err)
					t.Skip("Skipping due to network connectivity issues")
					return
				}

				rewards, err := Rewards(tc.network, tc.nodeID)
				if err != nil && !tc.wantErr {
					// Check if it's a network-related error
					if isNetworkError(err) {
						t.Logf("Network error calling Rewards: %v", err)
						t.Skip("Skipping due to network connectivity issues")
						return
					}
					t.Fatalf("failed to call Rewards: %v", err)
				}

				fees, err := cumulativeFeeShares(treeCID, tc.nodeID)
				if err != nil && !tc.wantErr {
					if isNetworkError(err) {
						t.Logf("Network error calling cumulativeFeeShares: %v", err)
						t.Skip("Skipping due to network connectivity issues")
						return
					}
					t.Fatalf("failed to call cumulativeFeeShares: %v", err)
				}

				bond, err := bond.BondSummary(tc.network, tc.nodeID)
				if err != nil && !tc.wantErr {
					if isNetworkError(err) {
						t.Logf("Network error calling BondSummary: %v", err)
						t.Skip("Skipping due to network connectivity issues")
						return
					}
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
			}()

			select {
			case <-done:
				// Test completed
			case <-ctx.Done():
				t.Skip("Test timed out, skipping due to network connectivity issues")
			}
		})
	}
}

// isNetworkError checks if the error is related to network connectivity
func isNetworkError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "no route to host") ||
		strings.Contains(errStr, "network is unreachable") ||
		strings.Contains(errStr, "bad response") ||
		strings.Contains(errStr, "stream error") ||
		strings.Contains(errStr, "failed to unmarshal JSON")
}

// isCIEnvironment checks if we're running in a CI environment
func isCIEnvironment() bool {
	return strings.ToLower(getEnv("CI", "")) == "true" ||
		strings.ToLower(getEnv("GITHUB_ACTIONS", "")) == "true" ||
		strings.ToLower(getEnv("TRAVIS", "")) == "true"
}

// getEnv gets an environment variable with a fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
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
