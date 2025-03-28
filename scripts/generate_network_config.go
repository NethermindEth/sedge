package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Network struct {
	Name      string `yaml:"name"`
	Sequencer string `yaml:"sequencer"`
}

type NetworkConfig struct {
	Networks []Network `yaml:"networks"`
}

type OutputConfig struct {
	Network            string `json:"network"`
	CL                 string `json:"cl"`
	CLImage            string `json:"cl_image"`
	CheckpointSyncURL  string `json:"checkpoint-sync-url"`
	Timeout            int    `json:"timeout"`
	Agent              string `json:"agent"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: generate_network_config [mainnet|sepolia]")
		os.Exit(1)
	}

	networkType := os.Args[1]
	if networkType != "mainnet" && networkType != "sepolia" {
		fmt.Println("Network type must be either 'mainnet' or 'sepolia'")
		os.Exit(1)
	}

	// Read the YAML file
	configFile := fmt.Sprintf("configs/superchain/%s.yaml", networkType)
	yamlData, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	var networkConfig NetworkConfig
	err = yaml.Unmarshal(yamlData, &networkConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)
		os.Exit(1)
	}

	var outputConfigs []OutputConfig
	for _, network := range networkConfig.Networks {
		// Skip empty networks
		if network.Name == "superchain" {
			continue
		}

		config := OutputConfig{
			Network:            network.Name + "-" + networkType,
			CL:                "",
			CLImage:           "",
			CheckpointSyncURL: "",
			Timeout:           600,
			Agent:            "g6-standard-6",
		}
		outputConfigs = append(outputConfigs, config)
	}

	// Write the JSON output directly in scripts folder
	outputFile := fmt.Sprintf("scripts/network_config_%s.json", networkType)
	jsonData, err := json.MarshalIndent(outputConfigs, "", "  ")
	if err != nil {
		fmt.Printf("Error generating JSON: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %s\n", outputFile)
}
