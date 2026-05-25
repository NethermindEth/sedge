package factory

import "testing"

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
