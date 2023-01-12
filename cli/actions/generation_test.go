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
package actions_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/clients"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	mock_client "github.com/NethermindEth/sedge/test/mock_docker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func newAction(t *testing.T, ctrl *gomock.Controller) actions.SedgeActions {
	t.Helper()
	dockerClient := mock_client.NewMockAPIClient(ctrl)
	serviceManager := services.NewServiceManager(dockerClient)
	return actions.NewSedgeActions(dockerClient, serviceManager, nil)
}

// Test that the generated compose file with dump data is generated correctly
func TestGenerateDockerCompose(t *testing.T) {
	configs.InitNetworksConfigs()
	samplePath := t.TempDir()
	sampleData := &generate.GenData{
		ExecutionClient: &clients.Client{Name: "nethermind"},
		Network:         "mainnet",
	}
	sedgeAction := newAction(t, nil)

	err := sedgeAction.Generate(actions.GenerateOptions{
		GenerationData: sampleData,
		GenerationPath: samplePath,
	})
	if err != nil {
		t.Error("GenerateDockerComposeAndEnvFile() failed", err)
		return
	}

	// Check that docker-compose file exists
	assert.FileExists(t, filepath.Join(samplePath, configs.DefaultDockerComposeScriptName))
	// Check that .env file doesn't exist
	assert.FileExists(t, filepath.Join(samplePath, configs.DefaultEnvFileName))

	// Validate that Execution Client info matches the sample data
	// load the docker-compose file
	composeFile, err := os.ReadFile(filepath.Join(samplePath, configs.DefaultDockerComposeScriptName))
	assert.Nilf(t, err, "unable to read docker-compose.yml")

	var composeData generate.ComposeData
	err = yaml.Unmarshal(composeFile, &composeData)
	assert.Nilf(t, err, "unable to parse docker-compose.yml")

	// Check that the execution client is nethermind
	assert.Equal(t, composeData.Services.Execution.ContainerName, "execution-client")

	// Check other services are nil
	assert.Nil(t, composeData.Services.Consensus)

}

func TestFolderCreationOnCompose(t *testing.T) {
	configs.InitNetworksConfigs()
	samplePath := t.TempDir() + "test"
	sampleData := &generate.GenData{
		ExecutionClient: &clients.Client{Name: "nethermind"},
		Network:         "mainnet",
	}
	sedgeAction := newAction(t, nil)

	err := sedgeAction.Generate(actions.GenerateOptions{
		GenerationData: sampleData,
		GenerationPath: samplePath,
	})
	if !assert.Nilf(t, err, "GenerateDockerComposeAndEnvFile() failed") {
		return
	}

	// Check that the folder was created
	assert.DirExists(t, samplePath)
	// Check that docker-compose file exists
	assert.FileExists(t, filepath.Join(samplePath, configs.DefaultDockerComposeScriptName))
	// Check that .env file doesn't exist
	assert.FileExists(t, filepath.Join(samplePath, configs.DefaultEnvFileName))
	// Remove the folder
	err = os.RemoveAll(samplePath)
	if !assert.Nilf(t, err, "unable to remove sample folder") {
		return
	}
	assert.NoDirExists(t, samplePath)
}
