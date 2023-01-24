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
package cli_test

import (
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSlashingImport_ValidatorIsRequired(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectedErr error
	}{
		{
			name:        "without flags",
			args:        []string{},
			expectedErr: cli.ErrInvalidNumberOfArguments,
		},
		{
			name:        "with flags",
			args:        []string{"--network", "sepolia", "--from", "slashing0import.json"},
			expectedErr: cli.ErrInvalidNumberOfArguments,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {})

		slashingImportCmd := cli.SlashingImportCmd(nil)
		slashingImportCmd.SetArgs(tt.args)
		err := slashingImportCmd.Execute()
		assert.ErrorIs(t, err, tt.expectedErr)
	}
}

func TestSlashingImport_Params(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		actionOptions actions.SlashingImportOptions
	}{
		{
			name: "validator argument",
			args: []string{"lighthouse"},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "lighthouse",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "slashing-export.json"),
			},
		},
		{
			name: "network flag",
			args: []string{"lighthouse", "--network", "sepolia"},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "lighthouse",
				Network:         "sepolia",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "slashing-export.json"),
			},
		},
		{
			name: "network shorthand flag",
			args: []string{"lighthouse", "-n", "sepolia"},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "lighthouse",
				Network:         "sepolia",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "slashing-export.json"),
			},
		},
		{
			name: "stop-validator flag",
			args: []string{"prysm", "--stop-validator"},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "prysm",
				Network:         "mainnet",
				StopValidator:   true,
				StartValidator:  false,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "slashing-export.json"),
			},
		},
		{
			name: "start-validator flag",
			args: []string{"teku", "--start-validator"},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "teku",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  true,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "slashing-export.json"),
			},
		},
		{
			name: "path flag",
			args: []string{"teku", "--path", filepath.Join("custom", "dir")},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "teku",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  filepath.Join("custom", "dir"),
				From:            filepath.Join("custom", "dir", "slashing-export.json"),
			},
		},
		{
			name: "path shorthand flag",
			args: []string{"teku", "-p", filepath.Join("custom", "dir")},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "teku",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  filepath.Join("custom", "dir"),
				From:            filepath.Join("custom", "dir", "slashing-export.json"),
			},
		},
		{
			name: "from flag",
			args: []string{"lodestar", "--from", filepath.Join("custom", "from", "file.json")},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "lodestar",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join("custom", "from", "file.json"),
			},
		},
		{
			name: "from shorthand flag",
			args: []string{"lodestar", "-f", filepath.Join("custom", "from", "file.json")},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "lodestar",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join("custom", "from", "file.json"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			actions := sedge_mocks.NewMockSedgeActions(ctrl)
			actions.EXPECT().ImportSlashingInterchangeData(tt.actionOptions).Times(1)

			slashingImportCmd := cli.SlashingImportCmd(actions)
			slashingImportCmd.SetArgs(tt.args)
			err := slashingImportCmd.Execute()

			assert.Nil(t, err)
		})
	}
}
