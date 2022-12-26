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
package generate

import "github.com/NethermindEth/sedge/internal/pkg/clients"

// EnvData : Struct Data object to be applied to the docker-compose script environment (.env) template
type EnvData struct {
	ElImage                   string
	ElDataDir                 string
	CcImage                   string
	CcDataDir                 string
	VlImage                   string
	VlDataDir                 string
	ExecutionApiURL           string
	ExecutionAuthURL          string
	ConsensusApiURL           string
	ConsensusAdditionalApiURL string
	FeeRecipient              string
	JWTSecretPath             string
	ExecutionEngineName       string
	ConsensusClientName       string
	KeystoreDir               string
	Graffiti                  string
	ECBootnodes               []string
	CCBootnodes               []string
	CustomTTD                 string
	CustomDeployBlock         string
}

// GenerationData : Struct Data object for script's generation
type GenerationData struct {
	Services                []string
	ExecutionClient         clients.Client
	ConsensusClient         clients.Client
	ValidatorClient         clients.Client
	GenerationPath          string
	Network                 string
	CheckpointSyncUrl       string
	FeeRecipient            string
	JWTSecretPath           string
	FallbackELUrls          []string
	ElExtraFlags            []string
	ClExtraFlags            []string
	VlExtraFlags            []string
	MapAllPorts             bool
	Mev                     bool
	MevImage                string
	Ports                   map[string]string
	Graffiti                string
	LoggingDriver           string
	ECBootnodes             []string
	CCBootnodes             []string
	CustomTTD               string
	CustomChainSpecPath     string
	CustomNetworkConfigPath string
	CustomGenesisPath       string
	CustomDeployBlock       string
	CustomDeployBlockPath   string
	VLStartGracePeriod      uint
}

// DockerComposeData : Struct Data object to be applied to docker-compose script
type DockerComposeData struct {
	Services                []string
	TTD                     bool
	XeeVersion              bool
	Mev                     bool
	MevPort                 string
	MevImage                string
	CheckpointSyncUrl       string
	FeeRecipient            string
	ElDiscoveryPort         string
	ElMetricsPort           string
	ElApiPort               string
	ElAuthPort              string
	ElWsPort                string
	ClDiscoveryPort         string
	ClMetricsPort           string
	ClApiPort               string
	ClAdditionalApiPort     string
	VlMetricsPort           string
	FallbackELUrls          []string
	ElExtraFlags            []string
	ClExtraFlags            []string
	VlExtraFlags            []string
	ECBootnodes             []string
	CCBootnodes             []string
	MapAllPorts             bool
	SplittedNetwork         bool
	ClCheckpointSyncUrl     bool
	LoggingDriver           string
	CustomCfgDownload       bool
	RemoteCfg               bool
	RemoteGen               bool
	RemoteDpl               bool
	VlCustomCfgDownload     bool
	VlRemoteCfg             bool
	VlRemoteGen             bool
	VlRemoteDpl             bool
	CustomConsensusConfigs  bool
	CustomNetwork           bool
	CustomChainSpecPath     string
	CustomNetworkConfigPath string
	CustomGenesisPath       string
	CustomDeployBlockPath   string
	VLStartGracePeriod      uint
}

// WithConsensusClient returns true if the consensus client is explicitly required
// by the user, with the --run-clients flag.
func (d DockerComposeData) WithConsensusClient() bool {
	for _, service := range d.Services {
		if service == "consensus" {
			return true
		}
	}
	return false
}

// GenerationResults: Struct for storing results of the generation process
type GenerationResults struct {
	DockerComposePath string
	EnvFilePath       string
	ELPort            string
	CLPort            string
}
