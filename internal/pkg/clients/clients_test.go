/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package clients

import (
	"fmt"
	"testing"

	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestSupportedClients(t *testing.T) {
	inputs := [...]struct {
		clientType ClientType
		network    string
		want       []string
		isErr      bool
	}{
		{ExecutionClientType, "gnosis", []string{"nethermind"}, false},
		{ConsensusClientType, "gnosis", utils.Filter(AllClients[ConsensusClientType], func(c string) bool { return c != "prysm" }), false},
		{ExecutionClientType, "mainnet", AllClients[ExecutionClientType], false},
		{ConsensusClientType, "mainnet", AllClients[ConsensusClientType], false},
		{ValidatorClientType, "mainnet", AllClients[ValidatorClientType], false},
		{"random", "mainnet", []string{}, true},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("SupportedClients(%s)", input.clientType)

		c := ClientInfo{Network: input.network}
		if res, err := c.SupportedClients(input.clientType); input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else {
				assert.EqualValues(t, input.want, res, descr)
			}
		}
	}
}

type clientsTestCase struct {
	configClientsTypes map[ClientType][]string
	query              []ClientType
	network            string
	isErr              bool
}

func validateClients(resultClients OrderedClients, tc clientsTestCase) bool {
	// Check if all query clients types are in the result types
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

func TestClients(t *testing.T) {
	inputs := [...]clientsTestCase{
		{
			map[ClientType][]string{
				ConsensusClientType: {"lighthouse", "prysm", "teku", "lodestar"},
				ValidatorClientType: {"lighthouse", "prysm", "teku", "lodestar"},
				ExecutionClientType: {"nethermind", "geth", "besu", "erigon"},
			},
			[]ClientType{ConsensusClientType},
			"mainnet",
			false,
		},
		{
			map[ClientType][]string{
				ConsensusClientType: {"lighthouse"},
				ExecutionClientType: {"nethermind"},
				ValidatorClientType: {"lighthouse"},
			},
			[]ClientType{"other"},
			"mainnet",
			true,
		},
		{
			map[ClientType][]string{
				ValidatorClientType: {"lighthouse", "prysm", "teku", "lodestar"},
				ExecutionClientType: {"nethermind", "geth", "besu", "erigon"},
			},
			[]ClientType{ExecutionClientType, ValidatorClientType},
			"mainnet",
			false,
		},
		{
			map[ClientType][]string{
				ValidatorClientType: {"lighthouse", "prysm", "teku", "lodestar"},
				ConsensusClientType: {"lighthouse", "prysm", "teku", "lodestar"},
				ExecutionClientType: {"nethermind", "geth", "besu", "erigon"},
			},
			[]ClientType{ConsensusClientType, "other"},
			"mainnet",
			true,
		},
		{
			map[ClientType][]string{
				ValidatorClientType: {"lighthouse", "teku", "lodestar"},
				ConsensusClientType: {"lighthouse", "teku", "lodestar"},
				ExecutionClientType: {"nethermind"},
			},
			[]ClientType{ConsensusClientType, ExecutionClientType, ValidatorClientType},
			"gnosis",
			false,
		},
	}

	for i, input := range inputs {
		t.Run(fmt.Sprintf("Network %s, testcase: %d", input.network, i), func(t *testing.T) {
			descr := fmt.Sprintf("Clients(%s)", input.query)

			c := ClientInfo{Network: input.network}
			if res, err := c.Clients(input.query); input.isErr && err == nil {
				t.Errorf("%s expected to fail", descr)
			} else if !input.isErr {
				if err != nil {
					t.Errorf("%s failed: %v", descr, err)
				} else if !validateClients(res, input) {
					t.Errorf("%s got invalid result: %v", descr, res)
				}
			}
		})
	}
}

func TestValidateClient(t *testing.T) {
	inputs := [...]struct {
		client     Client
		clientType ClientType
		isErr      bool
	}{
		{
			client:     Client{},
			clientType: "",
			isErr:      true,
		},
		{
			client: Client{
				Name:      "nethermind",
				Type:      ExecutionClientType,
				Supported: true,
			},
			clientType: ExecutionClientType,
			isErr:      false,
		},
		{
			client: Client{
				Name:      "nethermind",
				Type:      ExecutionClientType,
				Supported: false,
			},
			clientType: ExecutionClientType,
			isErr:      true,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("ValidateClient(%v, %s)", input.client, input.clientType)

		if err := ValidateClient(&input.client, input.clientType); input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
