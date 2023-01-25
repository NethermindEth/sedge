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
	Mev                       bool
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
	RelayURL                  string
}

// GenData : Struct Data object for script's generation
type GenData struct {
	Services                []string
	ExecutionClient         *clients.Client
	ConsensusClient         *clients.Client
	ValidatorClient         *clients.Client
	Network                 string
	CheckpointSyncUrl       string
	FeeRecipient            string
	JWTSecretPath           string
	FallbackELUrls          *[]string
	ElExtraFlags            *[]string
	ClExtraFlags            *[]string
	VlExtraFlags            *[]string
	MapAllPorts             bool
	Mev                     bool
	RelayURL                string
	MevImage                string
	MevBoostService         bool
	MevBoostEndpoint        string
	MevBoostOnValidator     bool
	Ports                   map[string]uint16
	Graffiti                string
	LoggingDriver           string
	ECBootnodes             *[]string
	CCBootnodes             *[]string
	CustomTTD               string
	CustomChainSpecPath     string
	CustomNetworkConfigPath string
	CustomGenesisPath       string
	CustomDeployBlock       string
	CustomDeployBlockPath   string
	VLStartGracePeriod      uint
	ExecutionApiUrl         string
	ExecutionAuthUrl        string
	ConsensusApiUrl         string
}

// DockerComposeData : Struct Data object to be applied to docker-compose script
type DockerComposeData struct {
	Services                []string
	TTD                     string
	XeeVersion              bool
	Mev                     bool
	MevPort                 uint16
	MevImage                string
	MevBoostEndpoint        string
	CheckpointSyncUrl       string
	FeeRecipient            string
	ElDiscoveryPort         uint16
	ElMetricsPort           uint16
	ElApiPort               uint16
	ElAuthPort              uint16
	ElWsPort                uint16
	ClDiscoveryPort         uint16
	ClMetricsPort           uint16
	ClApiPort               uint16
	ClAdditionalApiPort     uint16
	VlMetricsPort           uint16
	FallbackELUrls          []string
	ElExtraFlags            []string
	ClExtraFlags            []string
	VlExtraFlags            []string
	ECBootnodes             string
	ECBootnodesList         []string
	CCBootnodes             string
	CCBootnodesList         []string
	MapAllPorts             bool
	SplittedNetwork         bool
	ClCheckpointSyncUrl     bool
	LoggingDriver           string
	CustomConsensusConfigs  bool
	CustomNetwork           bool
	CustomChainSpecPath     string
	CustomNetworkConfigPath string
	CustomGenesisPath       string
	CustomDeployBlock       bool
	CustomDeployBlockPath   string // Needed for lighthouse
	VLStartGracePeriod      uint
	UID                     int // Needed for teku
	GID                     int // Needed for teku
}

// WithConsensusClient returns true if the consensus client is explicitly required
// by the user, with the --run-clients flag.
func (d DockerComposeData) WithConsensusClient() bool {
	for _, service := range d.Services {
		if service == consensus {
			return true
		}
	}
	return false
}

type ComposeData struct {
	Version  string    `yaml:"version"`
	Services *Services `yaml:"services"`
	Networks *Networks `yaml:"networks"`
}
type Options struct {
	MaxSize string `yaml:"max-size"`
	MaxFile string `yaml:"max-file"`
}
type Logging struct {
	Driver  string   `yaml:"driver"`
	Options *Options `yaml:"options"`
}
type Execution struct {
	StopGracePeriod string   `yaml:"stop_grace_period"`
	ContainerName   string   `yaml:"container_name"`
	Restart         string   `yaml:"restart"`
	Image           string   `yaml:"image"`
	Networks        []string `yaml:"networks"`
	Volumes         []string `yaml:"volumes"`
	Ports           []string `yaml:"ports"`
	Expose          []int    `yaml:"expose"`
	Command         []string `yaml:"command"`
	Logging         *Logging `yaml:"logging"`
}
type Mevboost struct {
	Image         string   `yaml:"image"`
	Networks      []string `yaml:"networks"`
	ContainerName string   `yaml:"container_name"`
	Restart       string   `yaml:"restart"`
	Entrypoint    []string `yaml:"entrypoint"`
}
type Consensus struct {
	StopGracePeriod string   `yaml:"stop_grace_period"`
	ContainerName   string   `yaml:"container_name"`
	Restart         string   `yaml:"restart"`
	Image           string   `yaml:"image"`
	Networks        []string `yaml:"networks"`
	Volumes         []string `yaml:"volumes"`
	Ports           []string `yaml:"ports"`
	Expose          []int    `yaml:"expose"`
	Command         []string `yaml:"command"`
	Logging         *Logging `yaml:"logging"`
}
type ValidatorBlocker struct {
	Image   string `yaml:"image"`
	Command string `yaml:"command"`
}
type ValidatorImportDependsOn struct {
	Condition string `yaml:"condition"`
}
type DependsOn struct {
	ValidatorImport *ValidatorImportDependsOn `yaml:"validator-import"`
}
type Validator struct {
	ContainerName string     `yaml:"container_name"`
	Image         string     `yaml:"image"`
	DependsOn     *DependsOn `yaml:"depends_on"`
	Networks      []string   `yaml:"networks"`
	Ports         []string   `yaml:"ports"`
	Volumes       []string   `yaml:"volumes"`
	Command       []string   `yaml:"command"`
	Logging       *Logging   `yaml:"logging"`
}
type ConfigConsensus struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Volumes       []string `yaml:"volumes"`
	Command       []string `yaml:"command"`
	Logging       *Logging `yaml:"logging"`
}
type Services struct {
	Execution        *Execution        `yaml:"execution"`
	Mevboost         *Mevboost         `yaml:"mevboost"`
	Consensus        *Consensus        `yaml:"consensus"`
	ValidatorBlocker *ValidatorBlocker `yaml:"validator-blocker"`
	Validator        *Validator        `yaml:"validator"`
	ConfigConsensus  *ConfigConsensus  `yaml:"config_consensus"`
}
type Sedge struct {
	Name string `yaml:"name"`
}
type Networks struct {
	Sedge *Sedge `yaml:"sedge"`
}
