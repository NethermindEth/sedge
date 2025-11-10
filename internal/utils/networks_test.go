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
package utils

import (
	"sort"
	"strings"
	"testing"
)

var networks = []string{"mainnet", "custom", "sepolia", "chiado", "gnosis", "hoodi"}

func TestSupportedNetworks(t *testing.T) {
	names, err := SupportedNetworks()
	if err != nil {
		t.Errorf("GetSupportedNetworks() failed, gave error: %v", err)
	}

	sort.Strings(networks)
	sort.Strings(names)

	zip, err := ZipString(networks, names)
	if err != nil {
		t.Errorf("Unexpected error after zipping list %v and %v. Error:%v", networks, names, err)
	}

	for _, l := range zip {
		if !strings.EqualFold(l[0], l[1]) {
			t.Errorf("Incorrect supported network. Want %v, got %v", l[0], l[1])
		}
	}
}
