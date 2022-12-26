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
	RelayURL                  string
}

// GenData : Struct Data object for script's generation
type GenData struct {
	ExecutionClient   *clients.Client
	ConsensusClient   *clients.Client
	ValidatorClient   *clients.Client
	MevBoostService   bool
	Network           string
	CheckpointSyncUrl string
	FeeRecipient      string
	JWTSecretPath     string
	FallbackELUrls    []string
	ElExtraFlags      []string
	ClExtraFlags      []string
	VlExtraFlags      []string
	MapAllPorts       bool
	Mev               bool
	MevImage          string
	Ports             map[string]string
	Graffiti          string
	LoggingDriver     string
	RelayURL          string
	MevBoostEndpoint  string
}

// DockerComposeData : Struct Data object to be applied to docker-compose script
type DockerComposeData struct {
	TTD                 bool
	CcCustomCfg         bool
	CcRemoteCfg         bool
	CcRemoteGen         bool
	CcRemoteDpl         bool
	VlCustomCfg         bool
	VlRemoteCfg         bool
	VlRemoteGen         bool
	VlRemoteDpl         bool
	XeeVersion          bool
	Mev                 bool
	MevPort             string
	MevImage            string
	CheckpointSyncUrl   string
	FeeRecipient        string
	ElDiscoveryPort     string
	ElMetricsPort       string
	ElApiPort           string
	ElAuthPort          string
	ElWsPort            string
	ClDiscoveryPort     string
	ClMetricsPort       string
	ClApiPort           string
	ClAdditionalApiPort string
	VlMetricsPort       string
	FallbackELUrls      []string
	ElExtraFlags        []string
	ClExtraFlags        []string
	VlExtraFlags        []string
	Bootnodes           []string
	MapAllPorts         bool
	SplittedNetwork     bool
	ClCheckpointSyncUrl bool
	LoggingDriver       string
	MevBoostEndpoint    string
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
type ValidatorImport struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Networks      []string `yaml:"networks"`
	Volumes       []string `yaml:"volumes"`
	Command       string   `yaml:"command"`
	Logging       *Logging `yaml:"logging"`
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
	Execution       *Execution       `yaml:"execution"`
	Mevboost        *Mevboost        `yaml:"mevboost"`
	Consensus       *Consensus       `yaml:"consensus"`
	ValidatorImport *ValidatorImport `yaml:"validator-import"`
	Validator       *Validator       `yaml:"validator"`
	ConfigConsensus *ConfigConsensus `yaml:"config_consensus"`
}
type Sedge struct {
	Name string `yaml:"name"`
}
type Networks struct {
	Sedge *Sedge `yaml:"sedge"`
}
