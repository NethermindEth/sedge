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
	"io"
	"log"
	"os"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

type RelayData struct {
	Mainnet []Relay `yaml:"mainnet"`
	Holesky []Relay `yaml:"holesky"`
}

/*
loadRelays :
This function is responsible for:
loading relays for each network from .yaml file
params :-
filename (string): .yaml file that has the relays
returns :-
a. map[string][]Relay
Map of network name and its relays
b. error
Error if any
*/
func loadRelays(filename string) (map[string][]Relay, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var relayData RelayData
	err = yaml.Unmarshal(data, &relayData)
	if err != nil {
		return nil, err
	}

	return map[string][]Relay{
		"mainnet": relayData.Mainnet,
		"holesky": relayData.Holesky,
	}, nil
}

func TestRelays(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	// Load relays from the YAML file
	expectedRelaysMap, err := loadRelays("relays.yaml")
	if err != nil {
		t.Fatalf("Failed to load relays: %v", err)
	}

	tcs := []struct {
		name           string
		network        string
		expectedRelays []Relay
	}{
		{
			"GetRelays Mainnet", "mainnet", expectedRelaysMap["mainnet"],
		},
		{
			"GetRelays Holesky", "holesky", expectedRelaysMap["holesky"],
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			relays, err := Relays(tc.network)
			if err != nil {
				t.Fatalf("Failed to call GetRelays: %v", err)
			}

			if !reflect.DeepEqual(relays, tc.expectedRelays) {
				t.Errorf("Relays do not match expected values\nGot: %+v\nExpected: %+v", relays, tc.expectedRelays)
			}
		})
	}
}
