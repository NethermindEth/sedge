package factory

import (
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/clients"
)

type fakeArbFlags struct {
	executionApiUrl string
}

func (f *fakeArbFlags) GetExecutionName() string            { return "" }
func (f *fakeArbFlags) GetConsensusName() string            { return "" }
func (f *fakeArbFlags) GetValidatorName() string            { return "" }
func (f *fakeArbFlags) GetOptimismName() string             { return "" }
func (f *fakeArbFlags) GetTaikoName() string                { return "" }
func (f *fakeArbFlags) GetSurgeName() string                { return "" }
func (f *fakeArbFlags) GetArbitrumName() string             { return "" }
func (f *fakeArbFlags) GetL2ExecutionName() string          { return "" }
func (f *fakeArbFlags) GetDistributedValidatorName() string { return "" }
func (f *fakeArbFlags) GetExecutionApiUrl() string          { return f.executionApiUrl }
func (f *fakeArbFlags) GetNetwork() string                  { return "" }
func (f *fakeArbFlags) IsNoValidator() bool                 { return false }

func TestNewArbitrumNodeInitializer(t *testing.T) {
	init := NewArbitrumNodeInitializer()
	if init == nil {
		t.Fatal("nil initializer")
	}
	if init.serviceType != "arbitrum" {
		t.Errorf("serviceType: want arbitrum, got %q", init.serviceType)
	}
	if init.l2ExecutionType != "arbexecution" {
		t.Errorf("l2ExecutionType: want arbexecution, got %q", init.l2ExecutionType)
	}
	if init.config.clientType != "arbitrum" {
		t.Errorf("config.clientType: want arbitrum, got %q", init.config.clientType)
	}
	if init.config.forceName != "nitro" {
		t.Errorf("config.forceName: want nitro, got %q", init.config.forceName)
	}
}

func TestArbitrumUpdateResult_ExternalL1(t *testing.T) {
	// Without external-L1: Execution/Consensus stay intact.
	exec := &clients.Client{Name: "geth"}
	cons := &clients.Client{Name: "lighthouse"}
	rollup := &clients.Client{Name: "nitro"}
	execL2 := &clients.Client{Name: "nethermind-arbitrum"}

	init := NewArbitrumNodeInitializer()
	init.execClient = execL2
	init.flags = &fakeArbFlags{} // executionApiUrl == ""

	result := &clients.Clients{Execution: exec, Consensus: cons}
	init.UpdateResult(result, rollup)
	if result.Execution != exec || result.Consensus != cons {
		t.Errorf("non-external mode: should not have cleared L1 clients")
	}
	if result.Arbitrum != rollup || result.L2Execution != execL2 {
		t.Errorf("non-external mode: rollup/EL not set")
	}

	// With external-L1: Execution/Consensus get cleared.
	init.flags = &fakeArbFlags{executionApiUrl: "http://l1:8545"}
	result = &clients.Clients{Execution: exec, Consensus: cons}
	init.UpdateResult(result, rollup)
	if result.Execution != nil || result.Consensus != nil {
		t.Errorf("external mode: should have cleared L1 clients, got Execution=%v Consensus=%v", result.Execution, result.Consensus)
	}
}
