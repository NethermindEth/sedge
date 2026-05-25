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
package configs

import "testing"

func TestArbitrumClientImagesPresent(t *testing.T) {
	cases := []struct {
		role, client string
		got          Image
		wantName     string
		wantVersion  string
	}{
		{
			role:        "arbexecution",
			client:      "nethermind-arbitrum",
			got:         ClientImages.ArbExecution.NethermindArbitrum,
			wantName:    "nethermind/nethermind-arbitrum",
			wantVersion: "0.1.0-alpha",
		},
		{
			role:        "arbitrum",
			client:      "nitro",
			got:         ClientImages.Arbitrum.Nitro,
			wantName:    "offchainlabs/nitro-node",
			wantVersion: "v3.10.0-rc.2-746bda2",
		},
	}
	for _, tc := range cases {
		if tc.got.Name == "" || tc.got.Version == "" {
			t.Fatalf("%s/%s missing from client_images.yaml (got %q:%q)", tc.role, tc.client, tc.got.Name, tc.got.Version)
		}
		if tc.got.Name != tc.wantName || tc.got.Version != tc.wantVersion {
			t.Errorf("%s/%s: got %s:%s, want %s:%s", tc.role, tc.client,
				tc.got.Name, tc.got.Version, tc.wantName, tc.wantVersion)
		}
	}
}
