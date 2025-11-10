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
package mevboostrelaylist

import (
	"fmt"
	"sort"

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Relay : Struct represent data of MEV-Boost Relay
type Relay struct {
	Uri         string `yaml:"Uri"`
	Operator    string `yaml:"Operator"`
	IsMandatory bool   `yaml:"IsMandatory"`
	Description string `yaml:"Description"`
}

var contractName = contracts.MEVBoostRelayAllowedList

/*
Relays :
This function is responsible for :-
retrieving a list of relays from the MEV-Boost Allowed List contract for a given network.
params :-
network (string): The name of the network (e.g., "mainnet", "hoodi").
returns :-
a. []Relay
List of relays
b. error
Error if any
*/
func Relays(network string) ([]Relay, error) {
	var relays []Relay
	contract, client, err := mevBoostRelayListContract(network)
	if err != nil {
		return relays, fmt.Errorf("failed to call mevBoostRelayListContract: %w", err)
	}
	defer client.Close()

	result, err := contract.GetRelays(nil)
	if err != nil {
		return relays, fmt.Errorf("failed to call GetRelays contract method: %w", err)
	}

	for _, r := range result {
		relay := Relay{
			Uri:         r.Uri,
			Operator:    r.Operator,
			IsMandatory: r.IsMandatory,
			Description: r.Description,
		}
		relays = append(relays, relay)
	}

	return relays, nil
}

/*
RelaysURI :
This function is responsible for :-
retrieving a list of relays URI from the MEV-Boost Allowed List contract for a given network.
params :-
network (string): The name of the network (e.g., "mainnet", "hoodi").
returns :-
a. []string
List of relays URI
b. error
Error if any
*/
func RelaysURI(network string) ([]string, error) {
	relays, err := Relays(network)
	if err != nil {
		return nil, err
	}
	relayURIs := []string{}
	for _, relay := range relays {
		relayURIs = append(relayURIs, relay.Uri)
	}
	return relayURIs, err
}

func LidoSupportedNetworksMevBoost() []string {
	networks := []string{}
	for network := range contracts.DeployedAddresses(contractName) {
		networks = append(networks, network)
	}
	sort.Strings(networks)
	return networks
}

func NetworkSupportedByLidoMevBoost(network string) ([]string, bool) {
	supportedNetworks := LidoSupportedNetworksMevBoost()
	var supported bool
	for _, supportedNetwork := range supportedNetworks {
		if network == supportedNetwork {
			supported = true
		}
	}
	if !supported {
		return supportedNetworks, supported
	}
	return nil, supported
}

func mevBoostRelayListContract(network string) (*Mevboostrelaylist, *ethclient.Client, error) {
	client, err := contracts.ConnectClient(network, false)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to client: %w", err)
	}

	contractAddress, err := contracts.ContractAddressByNetwork(contractName, network)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get deployed contract address: %w", err)
	}

	address := common.HexToAddress(contractAddress)
	contract, err := NewMevboostrelaylist(address, client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Mevboostrelaylist contract instance: %w", err)
	}
	return contract, client, nil
}
