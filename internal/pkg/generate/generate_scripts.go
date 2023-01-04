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
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/NethermindEth/sedge/internal/pkg/env"

	"io"

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

func validateValidator(gd *GenData, c *clients.ClientInfo) error {
	if gd.ValidatorClient == nil || gd.ValidatorClient.Omitted {
		return nil
	}
	validatorClients, err := c.SupportedClients(validator)
	if err != nil {
		return UnableToGetClientsInfoError
	}
	if !utils.Contains(validatorClients, gd.ValidatorClient.Name) {
		return ValidatorClientNotValidError
	}
	return nil
}

func validateExecution(gd *GenData, c *clients.ClientInfo) error {
	if gd.ExecutionClient == nil || gd.ExecutionClient.Omitted {
		return nil
	}
	executionClients, err := c.SupportedClients(execution)
	if err != nil {
		return UnableToGetClientsInfoError
	}
	if !utils.Contains(executionClients, gd.ExecutionClient.Name) {
		return ExecutionClientNotValidError
	}
	return nil
}

func validateConsensus(gd *GenData, c *clients.ClientInfo) error {
	if gd.ConsensusClient == nil || gd.ConsensusClient.Omitted {
		return nil
	}

	consensusClients, err := c.SupportedClients(consensus)
	if err != nil {
		return UnableToGetClientsInfoError
	}
	if !utils.Contains(consensusClients, gd.ConsensusClient.Name) {
		return ConsensusClientNotValidError
	}
	return nil
}

// getClients convert genData clients to clients.Clients
func getClients(gd *GenData) map[string]*clients.Client {
	cls := map[string]*clients.Client{
		execution: gd.ExecutionClient,
		consensus: gd.ConsensusClient,
		validator: gd.ValidatorClient,
	}

	for i := range cls {
		if cls[i] == nil {
			cls[i] = &clients.Client{Omitted: true}
		}
	}

	return cls
}

func ComposeFile(gd *GenData, at io.Writer) error {
	// Check empty data
	if gd == nil {
		return EmptyDataError
	}
	if gd.ExecutionClient == nil && gd.ConsensusClient == nil && gd.ValidatorClient == nil {
		return EmptyDataError
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

	rawBaseTmp, err := templates.Services.ReadFile(filepath.Join("services", "docker-compose_base.tmpl"))
	if err != nil {
		return err
	}

	baseTmp, err := template.New("docker-compose").Parse(string(rawBaseTmp))
	if err != nil {
		return err
	}

	cls := getClients(gd)

	for tmpKind, client := range cls {
		name := client.Name
		if client.Omitted {
			name = empty
		}
		tmp, err := templates.Services.ReadFile(filepath.Join("services",
			configs.NetworksConfigs()[gd.Network].NetworkService,
			tmpKind,
			name+".tmpl"))
		if err != nil {
			return err
		}
		_, err = baseTmp.Parse(string(tmp))
		if err != nil {
			return err
		}
	}
	validatorBlockerTemplate := ""
	if !cls[validator].Omitted {
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

	if !cls[execution].Omitted {
		gd.ExecutionClient.Endpoint = configs.OnPremiseExecutionURL
	}

	if !cls[consensus].Omitted {
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
	ccBootnodes := gd.CCBootnodes
	if ccBootnodes == nil || len(ccBootnodes) == 0 {
		ccBootnodes = configs.NetworksConfigs()[gd.Network].DefaultCCBootnodes
	}

	// Check for Bootnode nodes
	ecBootnodes := gd.ECBootnodes
	if ecBootnodes == nil || len(ecBootnodes) == 0 {
		ecBootnodes = configs.NetworksConfigs()[gd.Network].DefaultECBootnodes
	}
	var mevSupported bool
	if !cls[validator].Omitted {
		mevSupported, err = env.CheckVariable(env.ReMEV, gd.Network, "validator", gd.ValidatorClient.Name)
		if err != nil {
			return err
		}
	}
	// If consensus is running with other services, and not set the MevBoostEndpoint, set it to the default
	if !cls[consensus].Omitted && (!cls[execution].Omitted || !cls[validator].Omitted) && gd.MevBoostEndpoint == "" && gd.Mev {
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
		FallbackELUrls:      gd.FallbackELUrls,
		ElExtraFlags:        gd.ElExtraFlags,
		ClExtraFlags:        gd.ClExtraFlags,
		VlExtraFlags:        gd.VlExtraFlags,
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

func EnvFile(gd *GenData, at io.Writer) error {
	if gd.ExecutionClient == nil && gd.ConsensusClient == nil && gd.ValidatorClient == nil {
		return EmptyDataError
	}
	rawBaseTmp, err := templates.Envs.ReadFile(filepath.Join("envs", gd.Network, "env_base.tmpl"))
	if err != nil {
		return TemplateNotFoundError
	}

	baseTmp, err := template.New("env").Parse(string(rawBaseTmp))
	if err != nil {
		return err
	}

	cls := getClients(gd)

	for tmpKind, client := range cls {
		var tmp []byte
		if client.Omitted {
			tmp, err = templates.Services.ReadFile(filepath.Join(
				"services",
				configs.NetworksConfigs()[gd.Network].NetworkService,
				tmpKind,
				"empty.tmpl"))
			if err != nil {
				return err
			}
		} else {
			tmp, err = templates.Envs.ReadFile(filepath.Join("envs", gd.Network, tmpKind, client.Name+".tmpl"))
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
	if executionApiUrl == "" {
		executionApiUrl = fmt.Sprintf("%s:%v", cls[execution].Endpoint, gd.Ports["ELApi"])
	}
	executionAuthUrl := gd.ExecutionAuthUrl
	if executionAuthUrl == "" {
		executionAuthUrl = fmt.Sprintf("%s:%v", cls[execution].Endpoint, gd.Ports["ELAuth"])
	}
	consensusApiUrl := gd.ConsensusApiUrl
	var consensusAdditionalApiUrl string
	if consensusApiUrl == "" {
		consensusAdditionalApiUrl = fmt.Sprintf("%s:%v", cls[consensus].Endpoint, gd.Ports["CLAdditionalApi"])
		consensusApiUrl = fmt.Sprintf("%s:%v", cls[consensus].Endpoint, gd.Ports["CLApi"])
	} else {
		if cls[consensus].Name == "prysm" {
			consensusAdditionalApiUrl = fmt.Sprintf("%s:%v", "consensus", gd.Ports["CLAdditionalApi"])
		} else {
			consensusAdditionalApiUrl = consensusApiUrl
		}
	}

	// TODO: Use OS wise delimiter for these data structs
	data := EnvData{
		ElImage:                   cls[execution].Image,
		ElDataDir:                 "./" + configs.ExecutionDir,
		CcImage:                   cls[consensus].Image,
		CcDataDir:                 "./" + configs.ConsensusDir,
		VlImage:                   cls[validator].Image,
		VlDataDir:                 "./" + configs.ValidatorDir,
		ExecutionApiURL:           executionApiUrl,
		ExecutionAuthURL:          executionAuthUrl,
		ConsensusApiURL:           consensusApiUrl,
		ConsensusAdditionalApiURL: consensusAdditionalApiUrl,
		FeeRecipient:              gd.FeeRecipient,
		JWTSecretPath:             gd.JWTSecretPath,
		ExecutionEngineName:       cls[execution].Name,
		ConsensusClientName:       cls[consensus].Name,
		KeystoreDir:               "./" + configs.KeystoreDir,
		Graffiti:                  gd.Graffiti,
		RelayURL:                  gd.RelayURL,
	}

	// Fix prysm rpc url
	if cls[validator].Name == "prysm" {
		data.ConsensusAdditionalApiURL = fmt.Sprintf("%s:%d", "consensus", gd.Ports["CLAdditionalApi"])
	}

	// Save to writer
	err = baseTmp.Execute(at, data)
	if err != nil {
		return err
	}

	return nil
}
