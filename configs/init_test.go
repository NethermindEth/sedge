package configs_test

import (
	"testing"

	"github.com/NethermindEth/sedge/configs"
)

func TestENRs_Uniqueness(t *testing.T) {
	enrSet := make(map[string]struct{})
	for _, network := range configs.NetworksConfigs() {
		for _, enr := range network.DefaultCCBootnodes {
			if _, ok := enrSet[enr]; ok {
				t.Errorf("ENR '%s' is duplicated", enr)
			} else {
				enrSet[enr] = struct{}{}
			}
		}
	}
}
