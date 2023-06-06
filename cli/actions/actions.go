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
	clientsimages "github.com/NethermindEth/sedge/configs/images"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/docker/docker/client"
)

//go:generate mockgen -package=sedge_mocks -destination=../../mocks/sedgeActions.go github.com/NethermindEth/sedge/cli/actions SedgeActions
type SedgeActions interface {
	ClientsImages() clientsimages.ClientsImages
	GetCommandRunner() commands.CommandRunner
	ImportSlashingInterchangeData(SlashingImportOptions) error
	ExportSlashingInterchangeData(SlashingExportOptions) error
	SetupContainers(SetupContainersOptions) error
	RunContainers(RunContainersOptions) error
	Generate(GenerateOptions) (generate.GenData, error)
	CreateJWTSecrets(CreateJWTSecretOptions) (string, error)
	ImportValidatorKeys(ImportValidatorKeysOptions) error
	ValidateDockerComposeFile(path string, services ...string) error
	GetContainersData(GetContainersDataOptions) (ContainersData, error)
	UpdateClients(UpdateClientsOptions) error
}

type sedgeActions struct {
	dockerClient   client.APIClient
	serviceManager services.ServiceManager
	commandRunner  commands.CommandRunner
	clientsImages  clientsimages.ClientsImages
}

type SedgeActionsOptions struct {
	DockerClient   client.APIClient
	ServiceManager services.ServiceManager
	CommandRunner  commands.CommandRunner
	ClientsImages  clientsimages.ClientsImages
}

func NewSedgeActions(options SedgeActionsOptions) SedgeActions {
	return &sedgeActions{
		dockerClient:   options.DockerClient,
		serviceManager: options.ServiceManager,
		commandRunner:  options.CommandRunner,
		clientsImages:  options.ClientsImages,
	}
}

func (s sedgeActions) GetCommandRunner() commands.CommandRunner {
	return s.commandRunner
}

func (s sedgeActions) ClientsImages() clientsimages.ClientsImages {
	return s.clientsImages
}
