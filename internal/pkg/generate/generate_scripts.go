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
	"text/template"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/env"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/NethermindEth/sedge/templates"
	log "github.com/sirupsen/logrus"
)

/*
GenerateScripts :
This function is responsible for generating docker-compose files for execution, consensus and
validator clients.

params :-
a. gd GenerationData
Data object containing clients whose script are to be generated, path of generated scripts and special options for the clients configuration.

returns :-
a. string
Execution client json-rpc API port
b. string
Consensus client HTTP API port
a. error
Error if any
*/
func GenerateScripts(gd GenerationData) (elPort, clPort string, err error) {
	// Create scripts directory if not exists
	if _, err := os.Stat(gd.GenerationPath); os.IsNotExist(err) {
		err = os.MkdirAll(gd.GenerationPath, 0755)
		if err != nil {
			return "", "", err
		}
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
	ports, err := utils.AssingPorts("localhost", defaultsPorts)
	if err != nil {
		return "", "", fmt.Errorf(configs.PortOccupationError, err)
	}
	gd.Ports = ports
	// External endpoints will be configured here. Also Ports should be updated with external ports
	gd.ExecutionEndpoint = configs.OnPremiseExecutionURL
	gd.ConsensusEndpoint = configs.OnPremiseConsensusURL

	log.Info(configs.GeneratingDockerComposeScript)
	err = generateDockerComposeScripts(gd)
	if err != nil {
		return "", "", err
	}

	log.Info(configs.GeneratingEnvFile)
	err = generateEnvFile(gd)
	if err != nil {
		return "", "", err
	}

	return ports["ELApi"], ports["CLApi"], nil
}

/*
generateDockerComposeScripts :
This function is responsible for generating docker-compose scripts for execution, consensus and
validator clients.

params :-
a. executionClient string
Execution client whose script is to be generated
b. consensusClient string
Execution client whose script is to be generated
c. validatorClient string
Execution client whose script is to be generated
d. path string
Path of generated scripts

returns :-
a. error
Error if any
*/
func generateDockerComposeScripts(gd GenerationData) (err error) {
	rawBaseTmp, err := templates.Services.ReadFile(filepath.Join("services", "docker-compose_base.tmpl"))
	if err != nil {
		return
	}

	baseTmp, err := template.New("docker-compose").Parse(string(rawBaseTmp))
	if err != nil {
		return
	}

	clients := map[string]string{
		"execution": gd.ExecutionClient,
		"consensus": gd.ConsensusClient,
		"validator": gd.ValidatorClient,
	}
	for tmpKind, clientName := range clients {
		tmp, err := templates.Services.ReadFile(filepath.Join("services", configs.NetworksToServices[gd.Network], tmpKind, clientName+".tmpl"))
		if err != nil {
			return err
		}
		_, err = baseTmp.Parse(string(tmp))
		if err != nil {
			return err
		}
	}

	// Check for TTD in env base template
	TTD, err := env.CheckVariableBase(env.ReTTD, gd.Network)
	if err != nil {
		return err
	}

	// Check for prysm config
	ccPrysmCfg, err := env.CheckVariable(env.ReCONFIG, gd.Network, "consensus", gd.ConsensusClient)
	if err != nil {
		return err
	}
	vlPrysmCfg, err := env.CheckVariable(env.ReCONFIG, gd.Network, "validator", gd.ValidatorClient)
	if err != nil {
		return err
	}

	// Check for XEE_VERSION in teku
	xeeVersion, err := env.CheckVariable(env.ReXEEV, gd.Network, "consensus", gd.ConsensusClient)
	if err != nil {
		return err
	}

	// Check for Mev
	mev, err := env.CheckVariable(env.ReMEV, gd.Network, "validator", gd.ValidatorClient)
	if err != nil {
		return err
	}

	data := DockerComposeData{
		TTD:                 TTD,
		CcPrysmCfg:          ccPrysmCfg,
		VlPrysmCfg:          vlPrysmCfg,
		XeeVersion:          xeeVersion,
		Mev:                 mev && gd.Mev,
		MevPort:             gd.Ports["MevPort"],
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
		MapAllPorts:         gd.MapAllPorts,
	}

	// Print docker-compose file
	log.Infof(configs.PrintingFile, configs.DefaultDockerComposeScriptName)
	err = baseTmp.Execute(os.Stdout, data)
	if err != nil {
		return fmt.Errorf(configs.PrintingFileError, configs.DefaultDockerComposeScriptName, err)
	}
	fmt.Println()

	err = writeTemplateToFile(baseTmp, filepath.Join(gd.GenerationPath, configs.DefaultDockerComposeScriptName), data, false)
	if err != nil {
		return fmt.Errorf(configs.GeneratingScriptsError, gd.ExecutionClient, gd.ConsensusClient, gd.ValidatorClient, err)
	}
	log.Infof(configs.CreatedFile, filepath.Join(gd.GenerationPath, configs.DefaultDockerComposeScriptName))

	return nil
}

/*
generateEnvFile :
This function is responsible for generating the environment variable for the
generated docker-compose scripts for execution, consensus and
validator clients.

params :-
a. executionClient string
Execution client whose script was generated
b. consensusClient string
Execution client whose script was generated
c. validatorClient string
Execution client whose script was generated
d. path string
Path of generated scripts

returns :-
a. error
Error if any
*/
func generateEnvFile(gd GenerationData) (err error) {
	rawBaseTmp, err := templates.Envs.ReadFile(filepath.Join("envs", gd.Network, "env_base.tmpl"))
	if err != nil {
		return
	}

	baseTmp, err := template.New("env").Parse(string(rawBaseTmp))
	if err != nil {
		return
	}

	clients := map[string]string{
		"execution": gd.ExecutionClient,
		"consensus": gd.ConsensusClient,
		"validator": gd.ValidatorClient,
	}
	for tmpKind, clientName := range clients {
		tmp, err := templates.Envs.ReadFile(filepath.Join("envs", gd.Network, tmpKind, clientName+".tmpl"))
		if err != nil {
			return err
		}
		_, err = baseTmp.Parse(string(tmp))
		if err != nil {
			return err
		}
	}

	// TODO: Use OS wise delimiter for these data structs
	data := EnvData{
		ElImage:                   gd.ExecutionImage,
		ElDataDir:                 configs.ExecutionDefaultDataDir,
		CcImage:                   gd.ConsensusImage,
		CcDataDir:                 configs.ConsensusDefaultDataDir,
		VlImage:                   gd.ValidatorImage,
		VlDataDir:                 configs.ValidatorDefaultDataDir,
		ExecutionApiURL:           gd.ExecutionEndpoint + ":" + gd.Ports["ELApi"],
		ExecutionAuthURL:          gd.ExecutionEndpoint + ":" + gd.Ports["ELAuth"],
		ConsensusApiURL:           gd.ConsensusEndpoint + ":" + gd.Ports["CLApi"],
		ConsensusAdditionalApiURL: gd.ConsensusEndpoint + ":" + gd.Ports["CLAdditionalApi"],
		FeeRecipient:              gd.FeeRecipient,
		JWTSecretPath:             gd.JWTSecretPath,
		ExecutionEngineName:       gd.ExecutionClient,
		KeystoreDir:               configs.KeystoreDefaultDataDir,
	}

	// Fix prysm rpc url
	if gd.ValidatorClient == "prysm" {
		data.ConsensusAdditionalApiURL = fmt.Sprintf("%s:%s", "consensus", gd.Ports["CLAdditionalApi"])
	}

	// Print .env file
	log.Infof(configs.PrintingFile, ".env")
	err = baseTmp.Execute(os.Stdout, data)
	if err != nil {
		return fmt.Errorf(configs.PrintingFileError, ".env", err)
	}
	fmt.Println()

	err = writeTemplateToFile(baseTmp, filepath.Join(gd.GenerationPath, ".env"), data, false)
	if err != nil {
		return fmt.Errorf(configs.GeneratingScriptsError, gd.ExecutionClient, gd.ConsensusClient, gd.ValidatorClient, err)
	}
	log.Infof(configs.CreatedFile, filepath.Join(gd.GenerationPath, ".env"))

	return nil
}

/*
writeTemplateToFile :
Write template to `file`.

params :-
a. template *template.Template
Template to be written
b. file string
File's complete path
c. data interface{}
Data object to be applied to `template`
d. append bool
True to append the template to `file`. False to create it or overwrite it.

returns :-
a. err error
Error if any
*/
func writeTemplateToFile(template *template.Template, file string, data interface{}, append bool) (err error) {
	var f *os.File

	if append {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf(configs.CreatingFileError, file, err)
		}
	} else {
		f, err = os.Create(file)
		if err != nil {
			return fmt.Errorf(configs.CreatingFileError, file, err)
		}
	}

	// Just closing a file without checking any closing errors is a bad practice
	defer func() {
		cerr := f.Close()
		if err == nil && cerr != nil {
			log.Errorf(configs.ClosingFileError, file)
			err = cerr
		}
	}()

	err = template.Execute(f, data)
	if err != nil {
		return
	}

	return nil
}
