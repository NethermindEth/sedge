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

	"github.com/stretchr/testify/assert"
)

func TestRandomChoice(t *testing.T) {
	inputs := [...]struct {
		clients ClientMap
		isErr   bool
	}{
		{
			ClientMap{},
			true,
		},
		{
			ClientMap{
				"a": &Client{
					Name:      "a",
					Type:      "A",
					Supported: true,
				},
				"b": &Client{
					Name:      "b",
					Type:      "A",
					Supported: true,
				},
				"c": &Client{
					Name:      "c",
					Type:      "A",
					Supported: true,
				},
			},
			false,
		},
		{
			ClientMap{
				"a": &Client{
					Name:      "a",
					Type:      "A",
					Supported: true,
				},
				"b": &Client{
					Name:      "b",
					Type:      "A",
					Supported: true,
				},
				"c": &Client{
					Name:      "c",
					Type:      "A",
					Supported: false,
				},
			},
			false,
		},
		{
			ClientMap{
				"a": &Client{
					Name:      "a",
					Type:      "A",
					Supported: false,
				},
				"b": &Client{
					Name:      "b",
					Type:      "A",
					Supported: false,
				},
			},
			true,
		},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("RandomChoice(%v)", input.clients)

		res, err := RandomChoice(input.clients)
		if input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr {
			if err != nil {
				t.Errorf("%s failed: %v", descr, err)
			} else if !validateResultClient(*res, input.clients) {
				t.Errorf("%s got invalid result: %v", descr, res)
			}

			if !res.Supported {
				t.Errorf("Got non supported client. RandomChoice(%+v) should only return a supported client", input.clients)
			}
		}

	}
}

func validateResultClient(client Client, clients ClientMap) bool {
	for _, other := range clients {
		if client.Name == other.Name && client.Supported == other.Supported && client.Type == other.Type {
			return true
		}
	}
	return false
}

// HACKME: Change this test to use fuzzy tests
func TestRandomClientName(t *testing.T) {
	tests := []struct {
		name    string
		clients []string
		err     error
	}{
		{
			name:    "empty",
			clients: []string{},
			err:     EmptyClientsListError,
		},
		{
			name:    "one",
			clients: []string{"a"},
			err:     nil,
		},
		{
			name:    "many",
			clients: []string{"a", "b", "c"},
			err:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomClientName(tt.clients)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Contains(t, tt.clients, got)
			}
		})
	}
}
