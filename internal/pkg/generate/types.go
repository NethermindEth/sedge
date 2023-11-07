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

import (
	"github.com/NethermindEth/sedge/internal/pkg/clients"
)

// EnvData : Struct Data object to be applied to the docker-compose script environment (.env) template
type EnvData struct {
	Services                  []string
	Mev                       bool
	ElImage                   string
	L2Image                   string //for juno
	ElDataDir                 string
	L2DataDir                 string //for juno
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
	RelayURLs                 string
	CheckpointSyncUrl         string
}

// GenData : Struct Data object for script's generation
type GenData struct {
	Services                []string
	ExecutionClient         *clients.Client
	ConsensusClient         *clients.Client
	ValidatorClient         *clients.Client
	StarknetClient          *clients.Client //starknet
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
	RelayURLs               []string
	MevImage                string
	MevBoostService         bool
	MevBoostEndpoint        string
	MevBoostOnValidator     bool
	Ports                   map[string]uint16
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
	SLStartGracePeriod      uint
	ExecutionApiUrl         string
	ExecutionAuthUrl        string
	ConsensusApiUrl         string
	ContainerTag            string

	// juno flags
	DbPath              string
	HttpPort            string
	WsPort              string
	MetricsPort         string
	GrpcPort            string
	PendingPollInterval string
}

// DockerComposeData : Struct Data object to be applied to docker-compose script
type DockerComposeData struct {
	Services            []string
	Network             string
	TTD                 string
	XeeVersion          bool
	Mev                 bool
	MevBoostOnValidator bool
	MevPort             uint16
	MevImage            string
	MevBoostEndpoint    string
	CheckpointSyncUrl   string
	FeeRecipient        string
	ElDiscoveryPort     uint16
	ElMetricsPort       uint16
	ElApiPort           uint16
	ElAuthPort          uint16

	L2ApiPort               uint16
	L2WsPort                uint16
	L2MetricsPort           uint16
	L2GrpcPort              uint16
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
	SLStartGracePeriod      uint
	UID                     int // Needed for teku
	GID                     int // Needed for teku
	ContainerTag            string

	// juno flags
	DbPath              string
	HttpPort            string
	WsPort              string
	MetricsPort         string
	GrpcPort            string
	PendingPollInterval string
}

// WithConsensusClient returns true if the consensus client is set
func (d DockerComposeData) WithConsensusClient() bool {
	for _, service := range d.Services {
		if service == consensus {
			return true
		}
	}
	return false
}

// WithValidatorClient returns true if the validator client is set
func (d DockerComposeData) WithValidatorClient() bool {
	for _, service := range d.Services {
		if service == validator {
			return true
		}
	}
	return false
}

// WithMevBoostClient returns true if the Mev-Boost client is set
func (d EnvData) WithMevBoostClient() bool {
	for _, service := range d.Services {
		if service == mevBoost {
			return true
		}
	}
	return false
}

// WithFullFlagStarknet returns true if Juno Node is run with the full flag
func (d DockerComposeData) WithFullFlagStarknet() bool {
	for _, service := range d.Services {
		if service == consensus || service == execution {
			return true
		}
	}
	return false
}

type ComposeData struct {
	Version  string    `yaml:"version,omitempty"`
	Services *Services `yaml:"services"`
	Networks *Networks `yaml:"networks,omitempty"`
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
	Logging         *Logging `yaml:"logging,omitempty"`
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
	Logging         *Logging `yaml:"logging,omitempty"`
}
type ValidatorBlocker struct {
	Image         string `yaml:"image"`
	ContainerName string `yaml:"container_name"`
	Command       string `yaml:"command"`
}
type StarknetBlocker struct {
	Image         string `yaml:"image"`
	ContainerName string `yaml:"container_name"`
	Command       string `yaml:"command"`
}
type ValidatorImportDependsOn struct {
	Condition string `yaml:"condition"`
}
type StarknetImportDependsOn struct {
	Condition string `yaml:"condition"`
}
type DependsOn struct {
	ValidatorImport *ValidatorImportDependsOn `yaml:"validator-import"`
	StarknetImport  *StarknetImportDependsOn  `yaml:"starknet-import"`
}

type Validator struct {
	ContainerName string     `yaml:"container_name"`
	Image         string     `yaml:"image"`
	DependsOn     *DependsOn `yaml:"depends_on"`
	Networks      []string   `yaml:"networks"`
	Ports         []string   `yaml:"ports"`
	Volumes       []string   `yaml:"volumes"`
	Command       []string   `yaml:"command"`
	Logging       *Logging   `yaml:"logging,omitempty"`
}
type ConfigConsensus struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Volumes       []string `yaml:"volumes"`
	Command       []string `yaml:"command"`
	Logging       *Logging `yaml:"logging,omitempty"`
}
type Services struct {
	Execution        *Execution        `yaml:"execution,omitempty"`
	Mevboost         *Mevboost         `yaml:"mev-boost,omitempty"`
	Consensus        *Consensus        `yaml:"consensus,omitempty"`
	ValidatorBlocker *ValidatorBlocker `yaml:"validator-blocker,omitempty"`
	StarknetBlocker  *StarknetBlocker  `yaml:"starknet-blocker,omitempty"`
	Validator        *Validator        `yaml:"validator,omitempty"`
	ConfigConsensus  *ConfigConsensus  `yaml:"config_consensus,omitempty"`
	Starknet         *Starknet         `yaml:"config_starknet,omitempty"` //starknet services
}
type Sedge struct {
	Name string `yaml:"name"`
}
type Networks struct {
	Sedge *Sedge `yaml:"sedge"`
}

// starknet
type Starknet struct {
	StopGracePeriod string     `yaml:"stop_grace_period"`
	ContainerName   string     `yaml:"container_name"`
	Restart         string     `yaml:"restart"`
	Image           string     `yaml:"image"`
	DependsOn       *DependsOn `yaml:"depends_on"`
	Networks        []string   `yaml:"networks"`
	Volumes         []string   `yaml:"volumes"`
	Ports           []string   `yaml:"ports"`
	Expose          []int      `yaml:"expose"`
	Command         []string   `yaml:"command"`
	Logging         *Logging   `yaml:"logging,omitempty"`
}
