package clients

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

// TODO: Add testcases for other networks

func validateSupportedClients(t *testing.T, clientType string, supportedClients []string) {
	//TODO: validate supported clients
}

func TestGetSupportedClients(t *testing.T) {
	inputs := [...]struct {
		clientType string
		network    string
		isErr      bool
	}{
		{"execution", "mainnet", false},
		{"consensus", "mainnet", false},
		{"validator", "mainnet", false},
		{"random", "mainnet", true},
		{"", "mainnet", true},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("GetSupportedClients(%s)", input.clientType)

		if res, err := GetSupportedClients(input.clientType, input.network); input.isErr && err == nil {
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

type getClientsTestCase struct {
	configClientsTypes map[string][]string
	query              []string
	network            string
	isErr              bool
}

func prepareGetConfigClientsTestCase(testCase getClientsTestCase) {
	for key, value := range testCase.configClientsTypes {
		viper.Set(key+"Clients", value)
	}
}

func cleanGetConfigClientsTestCase(_ getClientsTestCase) {
	viper.Reset()
}

func cleanAll() {
	viper.Reset()
}

func validateClients(resultClients OrderedClients, tc getClientsTestCase) bool {
	//Check if all query clients types are in the result types
Loop1:
	for _, queryType := range tc.query {
		for resultType := range resultClients {
			if queryType == resultType {
				continue Loop1
			}
		}
		return false
	}

	for resultType, resultTypeClients := range resultClients {
	Loop2:
		for _, resultClient := range resultTypeClients {
			for _, configClientName := range tc.configClientsTypes[resultType] {
				if resultClient.Name == configClientName {
					continue Loop2
				}
			}
			return false
		}
	}
	return true
}

func TestGetClients(t *testing.T) {
	inputs := [...]getClientsTestCase{
		{},
		{
			map[string][]string{
				"consensus": {"lighthouse"},
				"execution": {"nethermind"},
				"validator": {"lighthouse"},
			},
			[]string{"consensus"},
			"mainnet",
			false,
		},
		{
			map[string][]string{
				"consensus": {"lighthouse"},
				"execution": {"nethermind"},
				"validator": {"lighthouse"},
			},
			[]string{"other"},
			"mainnet",
			true,
		},
		{
			map[string][]string{
				"consensus": {"lighthouse"},
				"execution": {"nethermind"},
				"validator": {"lighthouse"},
			},
			[]string{"execution", "validator"},
			"mainnet",
			false,
		},
		{
			map[string][]string{
				"consensus": {"lighthouse"},
				"execution": {"nethermind"},
				"validator": {"lighthouse"},
			},
			[]string{"consensus", "other"},
			"mainnet",
			true,
		},
	}

	t.Cleanup(cleanAll)

	for _, input := range inputs {
		descr := fmt.Sprintf("GetClients(%s)", input.query)

		prepareGetConfigClientsTestCase(input)

		if res, err := GetClients(input.query, input.network); input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if !validateClients(res, input) {
				t.Errorf("%s got invalid result: %v", descr, res)
			}
		}

		cleanGetConfigClientsTestCase(input)
	}
}

func TestValidateClient(t *testing.T) {
	inputs := [...]struct {
		client     Client
		clientType string
		isErr      bool
	}{
		{
			client:     Client{},
			clientType: "",
			isErr:      true,
		},
		{
			client: Client{
				"nethermind",
				"execution",
				true,
			},
			clientType: "execution",
			isErr:      false,
		},
		{
			client: Client{
				"nethermind",
				"execution",
				false,
			},
			clientType: "execution",
			isErr:      true,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("ValidateClient(%v, %s)", input.client, input.clientType)

		if err := ValidateClient(input.client, input.clientType); input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
