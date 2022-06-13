package utils

import (
	"sort"
	"strings"
	"testing"
)

var networks = []string{"mainnet", "kiln"}

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
