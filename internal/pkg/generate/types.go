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
	ExecutionIsRemote         bool
	ConsensusApiURL           string
	ConsensusAdditionalApiURL string
	ConsensusIsRemote         bool
	FeeRecipient              string
	JWTSecretPath             string
	ExecutionEngineName       string
	KeystoreDir               string
}

// GenerationData : Struct Data object for script's generation
type GenerationData struct {
	ExecutionClient   string
	ExecutionImage    string
	ExecutionEndpoint string
	ExecutionIsRemote bool
	ConsensusClient   string
	ConsensusImage    string
	ConsensusEndpoint string
	ConsensusIsRemote bool
	ValidatorClient   string
	ValidatorImage    string
	GenerationPath    string
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
	Ports             map[string]string
}

// DockerComposeData : Struct Data object to be applied to docker-compose script
type DockerComposeData struct {
	TTD                 bool
	CcPrysmCfg          bool
	VlPrysmCfg          bool
	XeeVersion          bool
	Mev                 bool
	MevPort             string
	CheckpointSyncUrl   string
	FeeRecipient        string
	ElDiscoveryPort     string
	ElMetricsPort       string
	ElApiPort           string
	ElAuthPort          string
	ElWsPort            string
	ExecutionIsRemote   bool
	ClDiscoveryPort     string
	ClMetricsPort       string
	ClApiPort           string
	ClAdditionalApiPort string
	ConsensusIsRemote   bool
	VlMetricsPort       string
	FallbackELUrls      []string
	ElExtraFlags        []string
	ClExtraFlags        []string
	VlExtraFlags        []string
	MapAllPorts         bool
	SplittedNetwork     bool
}
