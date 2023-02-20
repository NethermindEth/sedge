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
package cli

import (
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	mock_actions "github.com/NethermindEth/sedge/cli/actions/mock"
	"github.com/NethermindEth/sedge/configs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestImportKeys_NumberOfArguments(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "no flags",
			args: []string{},
		},
		{
			name: "with flags",
			args: []string{"--network", "goerli"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := ImportKeysCmd(nil)
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			assert.ErrorIs(t, err, ErrInvalidNumberOfArguments)
		})
	}
}

func TestImportKeys_ArgsAndFlags(t *testing.T) {
	tests := []struct {
		name            string
		args            []string
		expectedOptions actions.ImportValidatorKeysOptions
	}{
		{
			name: "no flags",
			args: []string{"lighthouse"},
			expectedOptions: actions.ImportValidatorKeysOptions{
				ValidatorClient: "lighthouse",
				Network:         "mainnet",
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "keystore"),
			},
		},
		{
			name: "with flags",
			args: []string{
				"prysm",
				"--network", "goerli",
				"--from", "/tmp/keystore",
				"--path", "/tmp/sedge",
				"--start-validator",
			},
			expectedOptions: actions.ImportValidatorKeysOptions{
				ValidatorClient: "prysm",
				Network:         "goerli",
				StartValidator:  true,
				GenerationPath:  "/tmp/sedge",
				From:            "/tmp/keystore",
			},
		},
		{
			name: "with shorthand flags",
			args: []string{
				"teku",
				"-n", "goerli",
				"--from", "/tmp/keystore",
				"-p", "/tmp/sedge",
				"--stop-validator",
			},
			expectedOptions: actions.ImportValidatorKeysOptions{
				ValidatorClient: "teku",
				Network:         "goerli",
				From:            "/tmp/keystore",
				GenerationPath:  "/tmp/sedge",
				StopValidator:   true,
			},
		},
		{
			name: "with custom configs",
			args: []string{
				"lighthouse",
				"--custom-config", "/tmp/config",
				"--custom-genesis", "/tmp/genesis",
				"--custom-deploy-block", "custom-deploy-block",
			},
			expectedOptions: actions.ImportValidatorKeysOptions{
				ValidatorClient: "lighthouse",
				Network:         "mainnet",
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "keystore"),
				CustomConfig: actions.ImportValidatorKeysCustomOptions{
					NetworkConfigPath: "/tmp/config",
					GenesisPath:       "/tmp/genesis",
					DeployBlockPath:   "custom-deploy-block",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actions := mock_actions.NewMockSedgeActions(gomock.NewController(t))

			actions.EXPECT().ImportValidatorKeys(tt.expectedOptions).Times(1)

			cmd := ImportKeysCmd(actions)
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			assert.NoError(t, err)
		})
	}
}
