package generate

import (
	"fmt"
	"testing"

	"github.com/NethermindEth/1click/internal/pkg/clients"
	"github.com/NethermindEth/1click/internal/utils"
)

const wrongDep string = "wrong_dep"

type generateTestCase struct {
	execution, consensus, validator, path string
	isErr                                 bool
}

func generateTestCases(t *testing.T) (tests []generateTestCase) {
	tests = []generateTestCase{}

	executionClients, err := clients.GetSupportedClients("execution")
	if err != nil {
		t.Errorf("GetSupportedClients(\"execution\") failed: %v", err)
	}
	consensusClients, err := clients.GetSupportedClients("consensus")
	if err != nil {
		t.Errorf("GetSupportedClients(\"consensus\") failed: %v", err)
	}
	validatorClients, err := clients.GetSupportedClients("validator")
	if err != nil {
		t.Errorf("GetSupportedClients(\"validator\") failed: %v", err)
	}

	tests = append(tests, generateTestCase{isErr: true})

	for _, execution := range executionClients {
		path := t.TempDir()
		tests = append(tests, generateTestCase{execution, wrongDep, wrongDep, path, true})
		for _, consensus := range consensusClients {
			if utils.Contains(validatorClients, consensus) {
				path := t.TempDir()
				tests = append(tests, generateTestCase{
					execution, consensus, consensus, path, false,
				}, generateTestCase{execution, consensus, consensus, "", true})
			}
		}
	}

	return
}

func validateGeneratedFiles(t *testing.T, testCase generateTestCase) {
	//TODO: validate generated files
}

func TestGenerateScripts(t *testing.T) {
	inputs := generateTestCases(t)

	for _, input := range inputs {
		descr := fmt.Sprintf("GenerateScripts(%s,%s,%s,%s)", input.execution, input.consensus, input.validator, input.path)

		if err := GenerateScripts(input.execution, input.consensus, input.validator, input.path); input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
		validateGeneratedFiles(t, input)
	}

}
