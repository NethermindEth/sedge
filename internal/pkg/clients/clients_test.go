package clients

import (
	"fmt"
	"testing"
)

func validateSupportedClients(t *testing.T, clientType string, supportedClients []string) {
	//TODO: validate supported clients
}

func TestGetSupportedClients(t *testing.T) {
	inputs := [...]struct {
		clientType string
		isErr      bool
	}{
		{"execution", false},
		{"consensus", false},
		{"validator", false},
		{"random", true},
		{"", true},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("GetSupportedClients(%s)", input.clientType)

		if res, err := GetSupportedClients(input.clientType); input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else {
				validateSupportedClients(t, input.clientType, res)
			}
		}
	}
}
