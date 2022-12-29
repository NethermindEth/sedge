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
	"path/filepath"
	"text/template"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/env"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/NethermindEth/sedge/templates"
	"io"
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
	if gd.ValidatorClient == nil || gd.ValidatorClient.Omited {
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
	if gd.ExecutionClient == nil || gd.ExecutionClient.Omited {
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
	if gd.ConsensusClient == nil || gd.ConsensusClient.Omited {
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
			cls[i] = &clients.Client{Omited: true}
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
	defaultsPorts := map[string]string{
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
		if client.Omited {
			name = empty
		}
		tmp, err := templates.Services.ReadFile(filepath.Join("services",
			configs.NetworksConfigs[gd.Network].NetworkService,
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
	if !cls[validator].Omited {
		validatorBlockerTemplate = "validator-blocker"
	} else {
		validatorBlockerTemplate = "empty-validator-blocker"
	}

	// Parse validator-blocker template
	tmp, err := templates.Services.ReadFile(filepath.Join("services", validatorBlockerTemplate+".tmpl"))
	if err != nil {
		return err
	}
	if _, err = baseTmp.Parse(string(tmp)); err != nil {
		return err
	}

	// Check for TTD in env base template
	TTD, err := env.CheckVariableBase(env.ReTTD, gd.Network)
	if err != nil {
		return err
	}

	// Check for splitted network flags
	splittedNetwork, err := env.CheckVariableBase(env.ReSPLITTED, gd.Network)
	if err != nil {
		return err
	}

	// Check vars related to Consensus service
	var ccRemoteCfg, ccRemoteGen, ccRemoteDpl, xeeVersion, clCheckpointSyncUrl bool
	var bootnodes []string

	if !cls[execution].Omited {
		gd.ExecutionClient.Endpoint = configs.OnPremiseExecutionURL
	}

	if !cls[consensus].Omited {
		gd.ConsensusClient.Endpoint = configs.OnPremiseConsensusURL
		// Check for custom network config
		ccRemoteCfg, err = env.CheckVariable(env.ReCONFIG, gd.Network, "consensus", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}
		ccRemoteGen, err = env.CheckVariable(env.ReGENESIS, gd.Network, "consensus", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}
		ccRemoteDpl, err = env.CheckVariable(env.ReDEPLOY, gd.Network, "consensus", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}
		// Check for XEE_VERSION in teku
		xeeVersion, err = env.CheckVariable(env.ReXEEV, gd.Network, "consensus", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}

		clCheckpointSyncUrl, err = env.CheckVariable(env.ReCHECKPOINT, gd.Network, "consensus", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}

		// Check for Bootstrap nodes
		bootnodes, err = env.GetBootnodes(gd.Network, gd.ConsensusClient.Name)
		if err != nil {
			return err
		}
	}

	// Check vars related to Validator service
	var vlRemoteCfg, vlRemoteGen, vlRemoteDpl, mevSupported bool

	if !cls[validator].Omited {
		vlRemoteCfg, err = env.CheckVariable(env.ReCONFIG, gd.Network, "validator", gd.ValidatorClient.Name)
		if err != nil {
			return err
		}
		vlRemoteGen, err = env.CheckVariable(env.ReGENESIS, gd.Network, "validator", gd.ValidatorClient.Name)
		if err != nil {
			return err
		}
		vlRemoteDpl, err = env.CheckVariable(env.ReDEPLOY, gd.Network, "validator", gd.ValidatorClient.Name)
		if err != nil {
			return err
		}

		// Check for Mev
		mevSupported, err = env.CheckVariable(env.ReMEV, gd.Network, "validator", gd.ValidatorClient.Name)
		if err != nil {
			return err
		}
	}

	// If consensus is running with other services, and not set the MevBoostEndpoint, set it to the default
	if !cls[consensus].Omited && (!cls[execution].Omited || !cls[validator].Omited) && gd.MevBoostEndpoint == "" && gd.Mev {
		mevSupported, err = env.CheckVariable(env.ReMEV, gd.Network, "validator", gd.ConsensusClient.Name)
		if err != nil {
			return err
		}
		if mevSupported {
			gd.MevBoostEndpoint = configs.DefaultMevBoostEndpoint + ":" + gd.Ports["MevPort"]
		}
	}

	data := DockerComposeData{
		Services:            gd.Services,
		TTD:                 TTD,
		CcCustomCfg:         ccRemoteCfg || ccRemoteGen || ccRemoteDpl,
		CcRemoteCfg:         ccRemoteCfg,
		CcRemoteGen:         ccRemoteGen,
		CcRemoteDpl:         ccRemoteDpl,
		VlCustomCfg:         vlRemoteCfg || vlRemoteGen || vlRemoteDpl,
		VlRemoteCfg:         vlRemoteCfg,
		VlRemoteGen:         vlRemoteGen,
		VlRemoteDpl:         vlRemoteDpl,
		XeeVersion:          xeeVersion,
		Mev:                 gd.MevBoostService || (mevSupported && gd.Mev),
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
		Bootnodes:           bootnodes,
		MapAllPorts:         gd.MapAllPorts,
		SplittedNetwork:     splittedNetwork,
		ClCheckpointSyncUrl: clCheckpointSyncUrl,
		LoggingDriver:       gd.LoggingDriver,
		VLStartGracePeriod:  gd.VLStartGracePeriod,
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
		if client.Omited {
			tmp, err = templates.Services.ReadFile(filepath.Join(
				"services",
				configs.NetworksConfigs[gd.Network].NetworkService,
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
		executionApiUrl = cls[execution].Endpoint + ":" + gd.Ports["ELApi"]
	}
	executionAuthUrl := gd.ExecutionAuthUrl
	if executionAuthUrl == "" {
		executionAuthUrl = cls[execution].Endpoint + ":" + gd.Ports["ELAuth"]
	}
	consensusApiUrl := gd.ConsensusApiUrl
	var consensusAdditionalApiUrl string
	if consensusApiUrl == "" {
		consensusAdditionalApiUrl = cls[consensus].Endpoint + ":" + gd.Ports["CLAdditionalApi"]
		consensusApiUrl = cls[consensus].Endpoint + ":" + gd.Ports["CLApi"]
	} else {
		if cls[consensus].Name == "prysm" {
			consensusAdditionalApiUrl = fmt.Sprintf("%s:%s", "consensus", gd.Ports["CLAdditionalApi"])
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

	// Save to writer
	err = baseTmp.Execute(at, data)
	if err != nil {
		return err
	}

	return nil
}
