/*
Copyright 2023 Nethermind

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
package actions

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/google/go-github/v54/github"
	log "github.com/sirupsen/logrus"
)

type GetCustomConfigsOptions struct {
	GenerationPath     string
	CustomConfigSource string // a file or github url to folder
}

type CustomConfigsResults struct {
	NetworkID          string
	DeployBlockValue   string
	Path               string
	CustomConfigs      map[string]string
	ExecutionBootnodes []string
	ConsensusBootnodes []string
}

func (results CustomConfigsResults) ToCustomNetworkData() generate.CustomNetworkData {
	return generate.CustomNetworkData{
		// General
		Path:      results.Path,
		NetworkID: results.NetworkID,
		// Execution
		NethermindChainspec: results.CustomConfigs[configs.NethermindChainspecFileName],
		GethGenesis:         results.CustomConfigs[configs.GethGenesisFileName],
		BesuGenesis:         results.CustomConfigs[configs.BesuGenesisFileName],
		// Consensus
		ConsensusConfig:          results.CustomConfigs[configs.ConsensusConfigFileName],
		GenesisState:             results.CustomConfigs[configs.GenesisStateFileName],
		DeployBlock:              results.CustomConfigs[configs.DeployBlockFileName],
		DeployBlockValue:         results.DeployBlockValue,
		DepositContract:          results.CustomConfigs[configs.DepositContractFileName],
		DepositContractBlock:     results.CustomConfigs[configs.DepositContractBlockFileName],
		DepositContractBlockHash: results.CustomConfigs[configs.DepositContractBlockHashFileName],
		TrustedSetupTxt:          results.CustomConfigs[configs.TrustedSetupTxtFileName],
		TrustedSetupJson:         results.CustomConfigs[configs.TrustedSetupJsonFileName],
	}
}

type customConfigItem struct {
	fileName           string
	allowedSourceNames []string
}

func (s *sedgeActions) GetCustomConfigs(options GetCustomConfigsOptions) (CustomConfigsResults, error) {
	log.Info("Getting custom network configs")
	results := CustomConfigsResults{}
	results.CustomConfigs = map[string]string{}

	sourcePath, err := s.getCustomConfigsSourcePath(options.CustomConfigSource)
	if err != nil {
		return results, err
	}

	customConfigTargets := []customConfigItem{
		// Execution clients custom configs
		{
			fileName: configs.NethermindChainspecFileName,
			allowedSourceNames: []string{
				"nethermind_chainspec.json",
				"chainspec.json",
			},
		},
		{
			fileName: configs.GethGenesisFileName,
			allowedSourceNames: []string{
				"geth_genesis.json",
				"genesis.json",
			},
		},
		{
			fileName: configs.BesuGenesisFileName,
			allowedSourceNames: []string{
				"besu_genesis.json",
				"besu.json",
			},
		},
		// Consensus clients custom configs
		{
			fileName: configs.ConsensusConfigFileName,
			allowedSourceNames: []string{
				"config.yaml",
				"config.yml",
			},
		},
		{
			fileName: configs.GenesisStateFileName,
			allowedSourceNames: []string{
				"genesis.ssz",
			},
		},
		{
			fileName: configs.DeployBlockFileName,
			allowedSourceNames: []string{
				"deploy_block.txt",
			},
		},
		{
			fileName: configs.DepositContractFileName,
			allowedSourceNames: []string{
				"deposit_contract.txt",
			},
		},
		{
			fileName: configs.DepositContractBlockFileName,
			allowedSourceNames: []string{
				"deposit_contract_block.txt",
			},
		},
		{
			fileName: configs.DepositContractBlockHashFileName,
			allowedSourceNames: []string{
				"deposit_contract_block_hash.txt",
			},
		},
		{
			fileName: configs.TrustedSetupTxtFileName,
			allowedSourceNames: []string{
				"trusted_setup.txt",
			},
		},
		{
			fileName: configs.TrustedSetupJsonFileName,
			allowedSourceNames: []string{
				"trusted_setup.json",
			},
		},
	}

	results.Path = path.Join(options.GenerationPath, configs.CustomNetworkConfigsFolder)
	err = s.copyCustomConfigs(
		sourcePath,
		results.Path,
		&results,
		customConfigTargets,
		[]string{
			"bootnode.txt",
			"bootstrap_nodes_execution.txt",
		},
		[]string{
			"bootstrap_nodes.txt",
		},
	)
	if err != nil {
		return results, err
	}

	// Get network id
	gethPath, gethOk := results.CustomConfigs[configs.GethGenesisFileName]
	besuPath, besuOk := results.CustomConfigs[configs.BesuGenesisFileName]
	if gethOk && gethPath != "" {
		results.NetworkID, err = getNetworkIDFrom(gethPath)
		if err != nil {
			return results, err
		}
	} else if besuOk && besuPath != "" {
		results.NetworkID, err = getNetworkIDFrom(besuPath)
		if err != nil {
			return results, err
		}
	} else {
		log.Warn("Network ID not found in custom configs. This is required for some execution clients.")
	}

	// Get deploy block
	deployBlockFile, ok := results.CustomConfigs[configs.DeployBlockFileName]
	if ok && deployBlockFile != "" {
		data, err := os.ReadFile(deployBlockFile)
		if err != nil {
			return results, err
		}
		results.DeployBlockValue = string(data)
	} else {
		log.Warn("Deploy Block not found in custom configs. This is required for some consensus clients.")
	}

	// Clean temp directory if downloaded
	if sourcePath != options.CustomConfigSource {
		err = os.RemoveAll(sourcePath)
		if err != nil {
			return results, err
		}
	}

	return results, nil
}

// Get or download custom configs files from source. returns a folder path
func (s sedgeActions) getCustomConfigsSourcePath(
	source string,
) (string, error) {
	newSourcePath := ""
	err := utils.HandleUrlOrPath(
		source,
		func(rawUrl string) error {
			// Download github folder link
			// TODO: Non github urls not supported yet
			uri, err := url.ParseRequestURI(rawUrl)
			if err != nil {
				return err
			}
			if uri.Scheme != "https" {
				return fmt.Errorf("invalid url scheme")
			}
			if uri.Hostname() != "github.com" {
				return fmt.Errorf("invalid url hostname")
			}

			pathParts := strings.Split(uri.Path, "/")
			if len(pathParts) < 5 {
				fmt.Println("invalid url path")
			}
			owner := pathParts[1]
			repo := pathParts[2]
			ref := pathParts[4]
			path := strings.Join(pathParts[5:], "/")
			client := github.NewClient(nil)
			newSourcePath = os.TempDir()

			err = utils.DownloadGithubObject(
				client,
				newSourcePath,
				owner,
				repo,
				path,
				ref,
			)
			return err
		},
		func(path string) error {
			// Get local folder path
			newSourcePath = path
			return nil
		},
	)
	return newSourcePath, err
}

// Copy custom configs files from source
func (s sedgeActions) copyCustomConfigs(
	sourcePath string,
	destPath string,
	results *CustomConfigsResults,
	copyTargets []customConfigItem,
	executionBootnodesAllowedNames []string,
	consensusBootnodesAllowedNames []string,
) error {
	contents, err := os.ReadDir(sourcePath)
	if err != nil {
		return err
	}

	for _, item := range contents {
		if item.IsDir() {
			continue
		}
		// Execution bootnodes
		for _, allowedName := range executionBootnodesAllowedNames {
			if item.Name() == allowedName {
				bootnodes, err := s.extractBootnodesFromFile(path.Join(sourcePath, item.Name()))
				if err != nil {
					return fmt.Errorf("error extracting bootnodes from file %s: %v", item.Name(), err)
				}
				results.ExecutionBootnodes = append(results.ExecutionBootnodes, bootnodes...)
				break
			}
		}
		// Consensus bootnodes
		for _, allowedName := range consensusBootnodesAllowedNames {
			if item.Name() == allowedName {
				bootnodes, err := s.extractBootnodesFromFile(path.Join(sourcePath, item.Name()))
				if err != nil {
					return fmt.Errorf("error extracting bootnodes from file %s: %v", item.Name(), err)
				}
				results.ConsensusBootnodes = append(results.ConsensusBootnodes, bootnodes...)
				break
			}
		}
		// Other custom config file
	targetsLoop:
		for _, target := range copyTargets {
			_, ok := results.CustomConfigs[target.fileName]
			if ok {
				continue
			}
			for _, allowedName := range target.allowedSourceNames {
				if item.Name() == allowedName {
					srcPath := path.Join(sourcePath, item.Name())
					dstPath := path.Join(destPath, target.fileName)
					if err = utils.CopyFile(srcPath, dstPath); err != nil {
						return fmt.Errorf(configs.ErrorCopyingFile, srcPath, err)
					}
					results.CustomConfigs[target.fileName] = dstPath
					break targetsLoop
				}
			}
		}
	}

	return nil
}

// Extract bootnodes from file
func (s sedgeActions) extractBootnodesFromFile(
	filePath string,
) ([]string, error) {
	bootnodes := []string{}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return bootnodes, err
	}

	rawBootnodes := string(bytes)
	for _, line := range strings.Split(rawBootnodes, "\n") {
		trimmedLine := strings.TrimSpace(line)
		for _, part := range strings.Split(trimmedLine, " ") {
			if part != "" && (strings.HasPrefix(part, "enode://") || strings.HasPrefix(part, "enr:")) {
				bootnodes = append(bootnodes, part)
			}
		}
	}

	return bootnodes, nil
}

func getNetworkIDFrom(
	gethGenesisFile string,
) (string, error) {
	data, err := os.ReadFile(gethGenesisFile)
	if err != nil {
		return "", err
	}

	var genesis struct {
		Config struct {
			ChainID int `json:"chainId"`
		} `json:"config"`
	}

	if err = json.Unmarshal(data, &genesis); err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", genesis.Config.ChainID), nil
}
