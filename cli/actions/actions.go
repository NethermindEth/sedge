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
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/docker/docker/client"
)

type SedgeActions interface {
	ImportSlashingInterchangeData(SlashingImportOptions) error
	ExportSlashingInterchangeData(SlashingExportOptions) error
	SetupContainers(SetupContainersOptions) error
	RunContainers(RunContainersOptions) error
	InstallDependencies(InstallDependenciesOptions) error
	Generate(GenerateOptions) error
	CreateJWTSecrets(CreateJWTSecretOptions) (string, error)
	ImportValidatorKeys(ImportValidatorKeysOptions) error
}

type sedgeActions struct {
	dockerClient   client.APIClient
	serviceManager services.ServiceManager
	commandRunner  commands.CommandRunner
}

func NewSedgeActions(dockerClient client.APIClient, serviceManager services.ServiceManager, commandRunner commands.CommandRunner) SedgeActions {
	return &sedgeActions{
		dockerClient:   dockerClient,
		serviceManager: serviceManager,
		commandRunner:  commandRunner,
	}
}
