package utils_test

import (
	"testing"

	"github.com/NethermindEth/1Click/internal/utils"

	"github.com/NethermindEth/1Click/configs"
)

func TestGetClients(t *testing.T) {

	testClient := configs.GetClients("")
	if testClient == nil {
		t.Fail()
	}
}

func TestCheckDependencies(t *testing.T) {
	testDep := utils.CheckDependencies(configs.GetDependencies())

	if testDep == nil {
		t.Fail()
	}
}

func TestGetOSInfo(t *testing.T) {

	distro, err := utils.GetOSInfo()

	print("%s %s", distro.Name, err)
	if err == nil {
		t.Fail()
	}
}

/*
func TestHandleInstructions(t *testing.T) {
	handler := error
	testGetScriptPath := utils.HandleInstructions(configs.GetDependencies(), handler)
	if testGetScriptPath == nil {
		t.Fail()
	}
}
*/
