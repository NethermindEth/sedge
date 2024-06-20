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

func TestGetRelays(t *testing.T) {
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
			relays, err := GetRelays(tc.network)
			if err != nil {
				t.Fatalf("Failed to call GetRelays %v", err)
			}

			if !reflect.DeepEqual(relays, tc.expectedRelays) {
				t.Errorf("Relays do not match expected values\nGot: %+v\nExpected: %+v", relays, tc.expectedRelays)
			}
		})
	}
}
