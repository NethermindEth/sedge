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
package actions

import (
	"os"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	log "github.com/sirupsen/logrus"
)

type GenerateOptions struct {
	GenerationData generate.GenData
	GenerationPath string
}

func (s *sedgeActions) Generate(options GenerateOptions) (generate.GenData, error) {
	// Create scripts directory if not exists
	if _, err := os.Stat(options.GenerationPath); os.IsNotExist(err) {
		err = os.MkdirAll(options.GenerationPath, 0o755)
		if err != nil {
			return options.GenerationData, err
		}
	}

	// Setup custom configs files if needed
	customConfigsPaths, err := generate.CustomNetworkConfigs(options.GenerationPath, options.GenerationData.Network, generate.CustomConfigsSources{
		ChainSpecSrc:     options.GenerationData.CustomChainSpecPath,
		NetworkConfigSrc: options.GenerationData.CustomNetworkConfigPath,
		GenesisSrc:       options.GenerationData.CustomGenesisPath,
		DeployBlockSrc:   options.GenerationData.CustomDeployBlockPath,
	})
	if err != nil {
		return options.GenerationData, err
	}
	options.GenerationData.CustomChainSpecPath = customConfigsPaths.ChainSpecPath
	options.GenerationData.CustomNetworkConfigPath = customConfigsPaths.NetworkConfigPath
	options.GenerationData.CustomGenesisPath = customConfigsPaths.GenesisPath
	options.GenerationData.CustomDeployBlockPath = customConfigsPaths.DeployBlockPath

	log.Info(configs.GeneratingDockerComposeScript)
	// open output file
	out, err := os.Create(filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName))
	if err != nil {
		return options.GenerationData, err
	}
	defer out.Close()
	err = generate.ComposeFile(&options.GenerationData, out)
	if err != nil {
		return options.GenerationData, err
	}
	log.Info(configs.GeneratedDockerComposeScript)

	log.Info(configs.GeneratingEnvFile)
	// open output file
	outEnv, err := os.Create(filepath.Join(options.GenerationPath, configs.DefaultEnvFileName))
	if err != nil {
		return options.GenerationData, err
	}
	defer outEnv.Close()
	err = generate.EnvFile(&options.GenerationData, outEnv)
	if err != nil {
		return options.GenerationData, err
	}
	log.Info(configs.GeneratedEnvFile)

	log.Info(configs.CleaningGeneratedFiles)
	err = generate.CleanGenerated(options.GenerationPath)
	if err != nil {
		return options.GenerationData, err
	}
	log.Info(configs.CleanedGeneratedFiles)

	// create datadir folders
	datadirs := []struct {
		path     string
		createIf bool
	}{
		{
			path:     filepath.Join(options.GenerationPath, configs.ExecutionDir),
			createIf: options.GenerationData.ExecutionClient != nil,
		},
		{
			path:     filepath.Join(options.GenerationPath, configs.ConsensusDir),
			createIf: options.GenerationData.ConsensusClient != nil,
		},
		{
			path:     filepath.Join(options.GenerationPath, configs.ValidatorDir),
			createIf: options.GenerationData.ValidatorClient != nil,
		},
	}
	for _, datadir := range datadirs {
		if datadir.createIf {
			_, err := os.Stat(datadir.path)
			if os.IsNotExist(err) {
				err = os.MkdirAll(datadir.path, 0o755)
				if err != nil {
					return options.GenerationData, err
				}
			} else {
				return options.GenerationData, err
			}
		}
	}

	return options.GenerationData, nil
}
