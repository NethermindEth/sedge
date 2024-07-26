package csmodule

import (
	"io"
	"log"
	"math/big"
	"testing"
	"unicode/utf8"
)

func TestNodeOpIDs(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)
	tcs := []struct {
		name    string
		network string
	}{
		{
			"NodeOpIDs, Holesky", "holesky",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			nodeOperatorIDs, err := nodeOpIDs(tc.network)
			if err != nil {
				t.Errorf("failed to call  NodeOpIDs: %v", err)
			}

			nodeOperatorsCount, err := nodeOpsCount(tc.network)
			if err != nil {
				t.Errorf("failed to call  nodeOpsCount: %v", err)
			}

			if len(nodeOperatorIDs) != int(nodeOperatorsCount.Int64()) {
				t.Errorf("mismatch: nodeOperatorIDs size (%d) != nodeOperatorsCount (%d)", len(nodeOperatorIDs), nodeOperatorsCount.Int64())
			}
		})
	}
}

func TestNodeOperatorInfo(t *testing.T) {
	log.SetOutput(io.Discard)
	tcs := []struct {
		name            string
		network         string
		nodeID          *big.Int
		expectedAddress string
		wantErr         bool
	}{
		{
			"Valid NodeID, Holesky", "holesky", big.NewInt(13), "0xC870Fd7316956C1582A2c8Fd2c42552cCEC70C88", false,
		},
		{
			"Valid Address, Holesky", "holesky", big.NewInt(4), "0xed1Fc097b5B9B007d40502e08aa0cddF477AaeaA", false,
		},
		{
			"Invalid Address, Holesky", "holesky", big.NewInt(4), "0xC870Fd7316956C1582A2c8Fd2c46752cCEC70C99", true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			nodeOp, err := NodeOperatorInfo(tc.network, tc.nodeID)
			if err != nil {
				t.Fatalf("failed to call NodeOperatorInfo: %v", err)
			}
			if nodeOp.RewardAddress.String() != tc.expectedAddress && !tc.wantErr {
				t.Errorf("Not same Reward Address, expected %v, got: %v", tc.expectedAddress, nodeOp.RewardAddress.Hex())
			}
		})
	}
}

func TestNodeID(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)
	tcs := []struct {
		name           string
		network        string
		expectedNodeID *big.Int
		wantErr        bool
	}{
		{
			"Valid NodeID, Holesky #1", "holesky", big.NewInt(13), false,
		},
		{
			"Valid NodeID, Holesky #2", "holesky", big.NewInt(4), false,
		},
		{
			"Invalid NodeID, Holesky #1", "holesky", big.NewInt(-4), true,
		},
		{
			"Invalid NodeID, Holesky #2", "holesky", big.NewInt(20000), true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			nodeOp, err := NodeOperatorInfo(tc.network, tc.expectedNodeID)
			if err != nil {
				t.Fatalf("failed to call NodeOperatorInfo: %v", err)
			}
			nodeID, err := NodeID(tc.network, nodeOp.RewardAddress.Hex())
			if err != nil && !tc.wantErr {
				t.Fatalf("failed to call NodeID: %v", err)
			}
			if nodeID != nil && nodeID.Cmp(tc.expectedNodeID) != 0 {
				t.Errorf("not same nodeID, expected %v, got: %v", tc.expectedNodeID, nodeID)
			}
		})
	}
}

func FuzzTestNodeID(f *testing.F) {
	testcases := []struct {
		network string
		nodeID  *big.Int
	}{
		{"holesky", big.NewInt(13)},
		{"holesky", big.NewInt(-1)},
		{"holesky", big.NewInt(40000)},
	}

	for _, tc := range testcases {
		f.Add(tc.network, tc.nodeID.String())
	}

	f.Fuzz(func(t *testing.T, network string, nodeIDStr string) {
		// Convert nodeIDStr back to *big.Int
		nodeID, ok := new(big.Int).SetString(nodeIDStr, 10)
		if !ok {
			t.Skip("Skipping invalid big.Int string")
		}

		// Silence logger
		log.SetOutput(io.Discard)

		nodeOp, err := NodeOperatorInfo(network, nodeID)
		if err != nil {
			t.Logf("Expected failure in NodeOperatorInfo: %v", err)
			return
		}

		nodeIDReturned, err := NodeID(network, nodeOp.RewardAddress.Hex())
		if err != nil {
			t.Logf("Expected failure in NodeID: %v", err)
			return
		}

		if nodeIDReturned != nil && utf8.ValidString(network) && nodeIDReturned.Cmp(nodeID) != 0 {
			t.Errorf("not same nodeID, expected %v, got: %v", nodeID, nodeIDReturned)
		}
	})
}
