package csmodule

import (
	"math/big"
	"testing"
)

func TestKeysStatus(t *testing.T) {
	keys, err := KeysStatus("holesky", big.NewInt(-1))
	if err != nil {
		t.Fatalf("failed to call KeysStatus: %v", err)
	}
	nodeInfo, err := NodeOperatorInfo("holesky", big.NewInt(-1))
	if err != nil {
		t.Fatalf("failed to call NodeOperatorInfo: %v", err)
	}
	expectedDeposited := big.NewInt(int64(nodeInfo.TotalDepositedKeys))
	deposited := keys.DepositedValidators
	if deposited.Cmp(expectedDeposited) != 0 {
		t.Errorf("Not same nodeID, expected %v, got: %v", expectedDeposited, deposited)
	}
}

func TestSigningKeys(t *testing.T) {
	signingKeys, err := SigningKeys("holesky", big.NewInt(1))
	if err != nil {
		t.Fatalf("failed to call SigningKeys: %v", err)
	}
	keysCount, err := nodeOpNonWithdrawnKeys("holesky", big.NewInt(1))
	if err != nil {
		t.Fatalf("failed to call nodeOpNonWithdrawnKeys: %v", err)
	}
	if len(signingKeys) != int(keysCount.Int64()) {
		t.Errorf("mismatch: keys size (%d) != keysCount (%d)", len(signingKeys), keysCount.Int64())
	}
}
