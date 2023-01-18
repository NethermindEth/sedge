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
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/NethermindEth/sedge/internal/pkg/env"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/NethermindEth/sedge/templates"
)

const (
	execution       = "execution"
	consensus       = "consensus"
	validator       = "validator"
	validatorImport = "validator-import"
	mevBoost        = "mev-boost"
	configConsensus = "config_consensus"
	empty           = "empty"
)

// validateClients validates each client in GenData
func validateClients(gd *GenData) error {
	c := clients.ClientInfo{Network: gd.Network}
	if err := validateConsensus(gd, &c); err != nil {
		return err
	}
	if err := validateExecution(gd, &c); err != nil {
		return err
	}
	if err := validateValidator(gd, &c); err != nil {
		return err
	}
	return nil
}

// validateValidator validates the validator client in GenData
func validateValidator(gd *GenData, c *clients.ClientInfo) error {
	if gd.ValidatorClient == nil {
		return nil
	}
	validatorClients, err := c.SupportedClients(validator)
	if err != nil {
		return ErrUnableToGetClientsInfo
	}
	if !utils.Contains(validatorClients, gd.ValidatorClient.Name) {
		return ErrValidatorClientNotValid
	}
	return nil
}

// validateExecution validates the execution client in GenData
func validateExecution(gd *GenData, c *clients.ClientInfo) error {
	if gd.ExecutionClient == nil {
		return nil
	}
	executionClients, err := c.SupportedClients(execution)
	if err != nil {
		return ErrUnableToGetClientsInfo
	}
	if !utils.Contains(executionClients, gd.ExecutionClient.Name) {
		return ErrExecutionClientNotValid
	}
	return nil
}

// validateConsensus validates the consensus client in GenData
func validateConsensus(gd *GenData, c *clients.ClientInfo) error {
	if gd.ConsensusClient == nil {
		return nil
	}

	consensusClients, err := c.SupportedClients(consensus)
	if err != nil {
		return ErrUnableToGetClientsInfo
	}
	if !utils.Contains(consensusClients, gd.ConsensusClient.Name) {
		return ErrConsensusClientNotValid
	}
	return nil
}

// mapClients convert genData clients to clients.Clients
func mapClients(gd *GenData) map[string]*clients.Client {
	cls := map[string]*clients.Client{
		execution: gd.ExecutionClient,
		consensus: gd.ConsensusClient,
		validator: gd.ValidatorClient,
	}

	return cls
}

// ComposeFile generates a docker-compose file with the provided GenData
func ComposeFile(gd *GenData, at io.Writer) error {
	// Check empty data
	if gd == nil {
		return ErrEmptyData
	}
	err := validateClients(gd)
	if err != nil {
		return err
	}
	// Check for port occupation
	defaultsPorts := map[string]uint16{
		"ELDiscovery":     configs.DefaultDiscoveryPortEL,
		"ELMetrics":       configs.DefaultMetricsPortEL,
		"ELApi":           configs.DefaultApiPortEL,
		"ELAuth":          configs.DefaultAuthPortEL,
		"ELWS":            configs.DefaultWSPortEL,
		"CLDiscovery":     configs.DefaultDiscoveryPortCL,
		"CLMetrics":       configs.DefaultMetricsPortCL,
		"CLApi":           configs.DefaultApiPortCL,
		"CLAdditionalApi": configs.DefaultAdditionalApiPortCL,
		"VLMetrics":       configs.DefaultMetricsPortVL,
		"MevPort":         configs.DefaultMevPort,
	}
	ports, err := utils.AssignPorts("localhost", defaultsPorts)
	if err != nil {
		// notest
		return fmt.Errorf(configs.PortOccupationError, err)
	}
	gd.Ports = ports

	rawBaseTmp, err := templates.Services.ReadFile(strings.Join([]string{"services", "docker-compose_base.tmpl"}, "/"))
	if err != nil {
		return err
	}

	baseTmp, err := template.New("docker-compose").Parse(string(rawBaseTmp))
	if err != nil {
		return err
	}

	cls := mapClients(gd)

	for tmpKind, client := range cls {
		var name string
		if client == nil {
			name = empty
		} else {
			name = client.Name
		}
		tmp, err := templates.Services.ReadFile(strings.Join([]string{
			"services",
			configs.NetworksConfigs()[gd.Network].NetworkService,
			tmpKind,
			name + ".tmpl",
		}, "/"))
		if err != nil {
			return err
		}
		_, err = baseTmp.Parse(string(tmp))
		if err != nil {
			return err
		}
	}
	validatorBlockerTemplate := ""
	if cls[validator] != nil {
		validatorBlockerTemplate = "validator-blocker"
	} else {
		validatorBlockerTemplate = "empty-validator-blocker"
	}

	// Parse validator-blocker template
	tmp, err := templates.Services.ReadFile(strings.Join([]string{"services", validatorBlockerTemplate + ".tmpl"}, "/"))
	if err != nil {
		return err
	}
	if _, err = baseTmp.Parse(string(tmp)); err != nil {
		return err
	}

	// Check for splitted network flags
	splittedNetwork, err := env.CheckVariableBase(env.ReSPLITTED, gd.Network)
	if err != nil {
		return err
	}

	// Check vars related to Consensus service
	var xeeVersion, clCheckpointSyncUrl bool

	if cls[execution] != nil {
		gd.ExecutionClient.Endpoint = configs.OnPremiseExecutionURL
	}

	if cls[consensus] != nil {
		gd.ConsensusClient.Endpoint = configs.OnPremiseConsensusURL
		// Check for XEE_VERSION in teku
		xeeVersion, err = env.CheckVariable(env.ReXEEV, gd.Network, "consensus", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}

		clCheckpointSyncUrl, err = env.CheckVariable(env.ReCHECKPOINT, gd.Network, "consensus", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}
	}

	ttd := gd.CustomTTD
	if len(ttd) == 0 {
		ttd = configs.NetworksConfigs()[gd.Network].DefaultTTD
	}

	// Check for CC Bootnode nodes
	var ccBootnodes []string
	if gd.CCBootnodes != nil {
		ccBootnodes = *gd.CCBootnodes
	}
	if len(ccBootnodes) == 0 {
		ccBootnodes = configs.NetworksConfigs()[gd.Network].DefaultCCBootnodes
	}

	// Check for Bootnode nodes
	var ecBootnodes []string
	if gd.ECBootnodes != nil {
		ecBootnodes = *gd.ECBootnodes
	}
	if len(ecBootnodes) == 0 {
		ecBootnodes = configs.NetworksConfigs()[gd.Network].DefaultECBootnodes
	}
	var mevSupported bool
	if cls[validator] != nil {
		mevSupported, err = env.CheckVariable(env.ReMEV, gd.Network, "validator", gd.ValidatorClient.Name)
		if err != nil {
			return err
		}
	}
	// If consensus is running with other services, and not set the MevBoostEndpoint, set it to the default
	if cls[consensus] != nil && cls[validator] != nil && gd.MevBoostEndpoint == "" && gd.Mev {
		mevSupported, err = env.CheckVariable(env.ReMEV, gd.Network, "validator", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}
		if mevSupported {
			gd.MevBoostEndpoint = fmt.Sprintf("%s:%v", configs.DefaultMevBoostEndpoint, gd.Ports["MevPort"])
		}
	}

	data := DockerComposeData{
		Services:            gd.Services,
		TTD:                 ttd,
		XeeVersion:          xeeVersion,
		Mev:                 gd.MevBoostService || (mevSupported && gd.Mev) || gd.MevBoostOnValidator,
		MevPort:             gd.Ports["MevPort"],
		MevBoostEndpoint:    gd.MevBoostEndpoint,
		MevImage:            gd.MevImage,
		CheckpointSyncUrl:   gd.CheckpointSyncUrl,
		FeeRecipient:        gd.FeeRecipient,
		ElDiscoveryPort:     gd.Ports["ELDiscovery"],
		ElMetricsPort:       gd.Ports["ELMetrics"],
		ElApiPort:           gd.Ports["ELApi"],
		ElAuthPort:          gd.Ports["ELAuth"],
		ElWsPort:            gd.Ports["ELWS"],
		ClDiscoveryPort:     gd.Ports["CLDiscovery"],
		ClMetricsPort:       gd.Ports["CLMetrics"],
		ClApiPort:           gd.Ports["CLApi"],
		ClAdditionalApiPort: gd.Ports["CLAdditionalApi"],
		VlMetricsPort:       gd.Ports["VLMetrics"],
		FallbackELUrls:      arrayOrEmpty(gd.FallbackELUrls),
		ElExtraFlags:        arrayOrEmpty(gd.ElExtraFlags),
		ClExtraFlags:        arrayOrEmpty(gd.ClExtraFlags),
		VlExtraFlags:        arrayOrEmpty(gd.VlExtraFlags),
		ECBootnodesList:     ecBootnodes,
		CCBootnodesList:     ccBootnodes,
		ECBootnodes:         strings.Join(ecBootnodes, ","),
		CCBootnodes:         strings.Join(ccBootnodes, ","),
		MapAllPorts:         gd.MapAllPorts,
		SplittedNetwork:     splittedNetwork,
		ClCheckpointSyncUrl: clCheckpointSyncUrl,
		LoggingDriver:       gd.LoggingDriver,
		VLStartGracePeriod:  gd.VLStartGracePeriod,
		CustomNetwork:       gd.Network == configs.CustomNetwork.Name, // Used custom templates
		CustomConsensusConfigs: gd.CustomNetworkConfigPath != "" ||
			gd.CustomGenesisPath != "" ||
			gd.CustomDeployBlockPath != "", // Have custom configs paths
		CustomChainSpecPath:     gd.CustomChainSpecPath,     // Path to chainspec.json
		CustomNetworkConfigPath: gd.CustomNetworkConfigPath, // Path to config.yaml
		CustomGenesisPath:       gd.CustomGenesisPath,       // Path to genesis.ssz
		CustomDeployBlockPath:   gd.CustomDeployBlockPath,   // Path to deploy_block.txt
		UID:                     os.Geteuid(),
		GID:                     os.Getegid(),
	}

	// Save to writer
	err = baseTmp.Execute(at, data)
	if err != nil {
		return err
	}

	return nil
}

// arrayOrEmpty returns an empty array if the input is nil, otherwise returns the input
func arrayOrEmpty(array *[]string) []string {
	if array == nil {
		return []string{}
	}
	return *array
}

// EnvFile generates a .env file with the provided GenData
func EnvFile(gd *GenData, at io.Writer) error {
	rawBaseTmp, err := templates.Envs.ReadFile(strings.Join([]string{"envs", gd.Network, "env_base.tmpl"}, "/"))
	if err != nil {
		return ErrTemplateNotFound
	}

	baseTmp, err := template.New("env").Parse(string(rawBaseTmp))
	if err != nil {
		return err
	}

	cls := mapClients(gd)

	for tmpKind, client := range cls {
		var tmp []byte
		if client == nil {
			tmp, err = templates.Services.ReadFile(strings.Join([]string{
				"services",
				configs.NetworksConfigs()[gd.Network].NetworkService,
				tmpKind,
				"empty.tmpl",
			}, "/"))
			if err != nil {
				return err
			}
		} else {
			tmp, err = templates.Envs.ReadFile(strings.Join([]string{"envs", gd.Network, tmpKind, client.Name + ".tmpl"}, "/"))
			if err != nil {
				return err
			}
		}
		_, err = baseTmp.Parse(string(tmp))
		if err != nil {
			return err
		}
	}
	executionApiUrl := gd.ExecutionApiUrl
	executionAuthUrl := gd.ExecutionAuthUrl
	if cls[execution] != nil {
		if executionApiUrl == "" {
			executionApiUrl = fmt.Sprintf("%s:%v", cls[execution].Endpoint, gd.Ports["ELApi"])
		}
		if executionAuthUrl == "" {
			executionAuthUrl = fmt.Sprintf("%s:%v", cls[execution].Endpoint, gd.Ports["ELAuth"])
		}
	}
	consensusApiUrl := gd.ConsensusApiUrl
	var consensusAdditionalApiUrl string
	if consensusApiUrl == "" {
		consensusAdditionalApiUrl = fmt.Sprintf("%s:%v", endpointOrEmpty(cls[consensus]), gd.Ports["CLAdditionalApi"])
		consensusApiUrl = fmt.Sprintf("%s:%v", endpointOrEmpty(cls[consensus]), gd.Ports["CLApi"])
	} else {
		if cls[consensus] != nil && cls[consensus].Name == "prysm" {
			consensusAdditionalApiUrl = fmt.Sprintf("%s:%v", "consensus", gd.Ports["CLAdditionalApi"])
		} else {
			consensusAdditionalApiUrl = consensusApiUrl
		}
	}

	var mevSupported bool
	if cls[validator] != nil {
		mevSupported, err = env.CheckVariable(env.ReMEV, gd.Network, "validator", gd.ValidatorClient.Name)
		if err != nil {
			return err
		}
	}
	// If consensus is running with the validator, and the MevBoostEndpoint is not set, set it to the default value
	if cls[consensus] != nil && cls[validator] != nil && gd.MevBoostEndpoint == "" && gd.Mev {
		mevSupported, err = env.CheckVariable(env.ReMEV, gd.Network, "validator", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}
		if mevSupported {
			gd.MevBoostEndpoint = fmt.Sprintf("%s:%v", configs.DefaultMevBoostEndpoint, gd.Ports["MevPort"])
		}
	}
	// TODO: Use OS wise delimiter for these data structs
	data := EnvData{
		Mev:                       gd.MevBoostService || (mevSupported && gd.Mev) || gd.MevBoostOnValidator,
		ElImage:                   imageOrEmpty(cls[execution]),
		ElDataDir:                 "./" + configs.ExecutionDir,
		CcImage:                   imageOrEmpty(cls[consensus]),
		CcDataDir:                 "./" + configs.ConsensusDir,
		VlImage:                   imageOrEmpty(cls[validator]),
		VlDataDir:                 "./" + configs.ValidatorDir,
		ExecutionApiURL:           executionApiUrl,
		ExecutionAuthURL:          executionAuthUrl,
		ConsensusApiURL:           consensusApiUrl,
		ConsensusAdditionalApiURL: consensusAdditionalApiUrl,
		FeeRecipient:              gd.FeeRecipient,
		JWTSecretPath:             gd.JWTSecretPath,
		ExecutionEngineName:       nameOrEmpty(cls[execution]),
		ConsensusClientName:       nameOrEmpty(cls[consensus]),
		KeystoreDir:               "./" + configs.KeystoreDir,
		Graffiti:                  gd.Graffiti,
		RelayURL:                  gd.RelayURL,
	}

	// Fix prysm rpc url
	if cls[validator] != nil && cls[validator].Name == "prysm" {
		data.ConsensusAdditionalApiURL = fmt.Sprintf("%s:%d", "consensus", gd.Ports["CLAdditionalApi"])
	}

	// Save to writer
	err = baseTmp.Execute(at, data)
	if err != nil {
		return err
	}

	return nil
}

// endpointOrEmpty returns the endpoint of the client if it is not nil, otherwise returns an empty string
func endpointOrEmpty(cls *clients.Client) string {
	if cls != nil {
		return cls.Endpoint
	}
	return ""
}

// nameOrEmpty returns the name of the client if it is not nil, otherwise returns an empty string
func nameOrEmpty(cls *clients.Client) string {
	if cls != nil {
		return cls.Name
	}
	return ""
}

// imageOrEmpty returns the image of the client if it is not nil, otherwise returns an empty string
func imageOrEmpty(cls *clients.Client) string {
	if cls != nil {
		return cls.Image
	}
	return ""
}
