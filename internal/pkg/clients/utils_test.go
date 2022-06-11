package clients

import (
	"fmt"
	"testing"
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
				"a": Client{
					"a",
					"A",
					true,
				},
				"b": Client{
					"b",
					"A",
					true,
				},
				"c": Client{
					"c",
					"A",
					true,
				},
			},
			false,
		},
		{
			ClientMap{
				"a": Client{
					"a",
					"A",
					true,
				},
				"b": Client{
					"b",
					"A",
					true,
				},
				"c": Client{
					"c",
					"A",
					false,
				},
			},
			false,
		},
		{
			ClientMap{
				"a": Client{
					"a",
					"A",
					false,
				},
				"b": Client{
					"b",
					"A",
					false,
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
			} else if !validateResultClient(res, input.clients) {
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
