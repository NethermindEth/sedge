package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type ChainConfig struct {
	Name      string `yaml:"name"`
	Sequencer string `yaml:"sequencer"`
}

type SuperChainConfig struct {
	Networks []ChainConfig `yaml:"networks"`
}

// GetSequencerURL returns the sequencer URL for a given network and chain name.
// network should be either "mainnet" or "sepolia"
// chainName is the name of the chain (e.g., "base", "op", etc.)
func GetSequencerURL(network, chainName string) (string, error) {
	configContent := GetSuperchainConfig(network)
	if configContent == "" {
		return "", fmt.Errorf("no config found for network %s", network)
	}
	data := []byte(configContent)

	var config SuperChainConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return "", fmt.Errorf("failed to parse config file: %w", err)
	}

	for _, chain := range config.Networks {
		if chain.Name == chainName {
			if chain.Sequencer == "" {
				return "", fmt.Errorf("no sequencer URL configured for chain %s in network %s", chainName, network)
			}
			return chain.Sequencer, nil
		}
	}

	return "", fmt.Errorf("chain %s not found in network %s", chainName, network)
}
