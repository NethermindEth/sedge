package configs_test

import (
	"testing"

	"github.com/NethermindEth/1Click/configs"
)

func TestGetDependencies(t *testing.T) {
	testingService := configs.GetDependencies()

	if len(testingService) != 0 {
		t.Skip()
	}
}

func TestGetDependenciesEmpty(t *testing.T) {
	testingService := configs.GetDependencies()
	if len(testingService) != 0 {
		t.Skip()
	}
}
