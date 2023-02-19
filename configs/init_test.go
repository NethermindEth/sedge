package configs_test

import (
	"testing"

	"github.com/NethermindEth/sedge/configs"
)

func TestBootNodes_Uniqueness(t *testing.T) {
	bootNodesSet := make(map[string]struct{})
	for _, network := range configs.NetworksConfigs() {
		for _, bootNode := range network.DefaultCCBootnodes {
			if _, ok := bootNodesSet[bootNode]; ok {
				t.Errorf("Boot node %s is duplicated", bootNode)
			} else {
				bootNodesSet[bootNode] = struct{}{}
			}
		}
	}
}
