package csmodule

import (
	"io"
	"log"
	"math/big"
	"reflect"
	"testing"
)

func TestNodeOpIDs(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)
	tcs := []struct {
		name    string
		network string
	}{
		{
			"NodeOpIDs Holesky", "holesky",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			nodeOperatorIDs, err := NodeOpIDs(tc.network)
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

func TestNodeID(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)
	tcs := []struct {
		name            string
		network         string
		nodeID          *big.Int
		expectedAddress string
	}{
		{
			"NodeID Holesky", "holesky", big.NewInt(13), "0xC870Fd7316956C1582A2c8Fd2c42552cCEC70C88",
		},
		{
			"NodeID Holesky", "holesky", big.NewInt(4), "0xed1Fc097b5B9B007d40502e08aa0cddF477AaeaA",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			expectedNodeOp, err := NodeOperatorInfo(tc.network, tc.nodeID)
			if err != nil {
				t.Fatalf("failed to call NodeOpIDs: %v", err)
			}
			expectedNodeID, err := NodeID(tc.network, tc.expectedAddress)
			if err != nil {
				t.Fatalf("failed to call NodeID: %v", err)
			}
			if tc.nodeID.Cmp(expectedNodeID) != 0 {
				t.Errorf("Not same nodeID, expected %v, got: %v", expectedNodeID, tc.nodeID)
			}
			nodeOp, err := NodeOperatorInfo(tc.network, tc.nodeID)
			if err != nil {
				t.Fatalf("failed to call NodeOperatorInfo: %v", err)
			}
			if !reflect.DeepEqual(expectedNodeOp, nodeOp) {
				t.Errorf("Nodes do not match expected values\nGot: %+v\nExpected: %+v", nodeOp, expectedNodeOp)
			}
		})
	}
}
