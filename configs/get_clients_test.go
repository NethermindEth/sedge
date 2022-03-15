package configs_test

import (
	"testing"

	"github.com/NethermindEth/1Click/configs"
)

func TestGetClientsExecution(t *testing.T) {
	executionClients := configs.GetClients("executionClients")

	if executionClients != nil {
		t.Skip()
	}
}

func TestGetClientsConsensus(t *testing.T) {
	consensusClients := configs.GetClients("consensusClients")
	if consensusClients != nil {
		t.Skip()
	}
}

func TestGetClientsEmpty(t *testing.T) {
	testingService := configs.GetClients("")
	if testingService == nil {
		t.Skip()
	}
}
