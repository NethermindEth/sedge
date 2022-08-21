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
package env

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/NethermindEth/sedge/internal/utils"
)

func TestGetBootnodes(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name    string
		network string
		client  string
		want    []string
		isErr   bool
	}{
		{
			name:    "Test case 1, no bootnodes",
			network: "mainnet",
			client:  "lighthouse",
			want:    nil,
			isErr:   false,
		},
		{
			name:    "Test case 2, bootnodes, ropsten, prysm",
			network: "ropsten",
			client:  "prysm",
			want:    []string{"enr:-Iq4QMCTfIMXnow27baRUb35Q8iiFHSIDBJh6hQM5Axohhf4b6Kr_cOCu0htQ5WvVqKvFgY28893DHAg8gnBAXsAVqmGAX53x8JggmlkgnY0gmlwhLKAlv6Jc2VjcDI1NmsxoQK6S-Cii_KmfFdUJL2TANL3ksaKUnNXvTCv1tLwXs0QgIN1ZHCCIyk", "enr:-KG4QMJSJ7DHk6v2p-W8zQ3Xv7FfssZ_1E3p2eY6kN13staMObUonAurqyWhODoeY6edXtV8e9eL9RnhgZ9va2SMDRQMhGV0aDKQS-iVMYAAAHD0AQAAAAAAAIJpZIJ2NIJpcIQDhAAhiXNlY3AyNTZrMaEDXBVUZhhmdy1MYor1eGdRJ4vHYghFKDgjyHgt6sJ-IlCDdGNwgiMog3VkcIIjKA"},
			isErr:   false,
		},
		{
			name:    "Test case 3, error, invalid network",
			network: "test",
			client:  "prysm",
			want:    nil,
			isErr:   true,
		},
		{
			name:    "Test case 4, error, invalid client",
			network: "mainnet",
			client:  "test",
			want:    nil,
			isErr:   true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetBootnodes(tc.network, tc.client)

			descr := fmt.Sprintf("GetBootnodes(%s, %s)", tc.network, tc.client)
			if err = utils.CheckErr(descr, tc.isErr, err); err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Expected %v, got %v. Function call: %s", tc.want, got, descr)
			}
		})
	}
}
