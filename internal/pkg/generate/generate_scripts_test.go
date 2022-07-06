package generate

import (
	"fmt"
	"testing"

	"github.com/NethermindEth/1click/internal/pkg/clients"
	"github.com/NethermindEth/1click/internal/utils"
)

const (
	wrongDep = "wrong_dep"
	network  = "mainnet"
)

type generateTestCase struct {
	execution, consensus, validator, path string
	isErr                                 bool
}

func generateTestCases(t *testing.T) (tests []generateTestCase) {
	tests = []generateTestCase{}
	c := clients.ClientInfo{Network: network}

	executionClients, err := c.SupportedClients("execution")
	if err != nil {
		t.Errorf("SupportedClients(\"execution\") failed: %v", err)
	}
	consensusClients, err := c.SupportedClients("consensus")
	if err != nil {
		t.Errorf("SupportedClients(\"consensus\") failed: %v", err)
	}
	validatorClients, err := c.SupportedClients("validator")
	if err != nil {
		t.Errorf("SupportedClients(\"validator\") failed: %v", err)
	}

	// TODO: Add CheckpointSyncUrl, FallbackELUrls and FeeRecipient to test data

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
	t.Parallel()
	inputs := generateTestCases(t)

	for i, input := range inputs {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			gd := GenerationData{
				ExecutionClient: input.execution,
				ConsensusClient: input.consensus,
				ValidatorClient: input.validator,
				GenerationPath:  input.path,
				Network:         network,
			}
			descr := fmt.Sprintf("GenerateScripts(%+v)", gd)

			if _, _, err := GenerateScripts(gd); input.isErr && err == nil {
				t.Errorf("%s expected to fail", descr)
			} else if !input.isErr && err != nil {
				t.Errorf("%s failed: %v", descr, err)
			}
			validateGeneratedFiles(t, input)
		})
	}
}
