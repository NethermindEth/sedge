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
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSlashingImport_ValidatorIsRequired(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

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
		slashingImportCmd.SetOutput(io.Discard)
		err := slashingImportCmd.Execute()
		assert.ErrorIs(t, err, tt.expectedErr)
	}
}

func TestSlashingImport_Params(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	customDir := t.TempDir()
	from := t.TempDir()
	if _, err := os.Create(filepath.Join(from, "slashing-data.json")); err != nil {
		t.Fatal(err)
	}

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
			args: []string{"teku", "--path", customDir},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "teku",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  customDir,
				From:            filepath.Join(customDir, "slashing-export.json"),
			},
		},
		{
			name: "path shorthand flag",
			args: []string{"teku", "-p", customDir},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "teku",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  customDir,
				From:            filepath.Join(customDir, "slashing-export.json"),
			},
		},
		{
			name: "from flag",
			args: []string{"lodestar", "--from", filepath.Join(from, "slashing-data.json")},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "lodestar",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(from, "slashing-data.json"),
			},
		},
		{
			name: "from shorthand flag",
			args: []string{"lodestar", "-f", filepath.Join(from, "slashing-data.json")},
			actionOptions: actions.SlashingImportOptions{
				ValidatorClient: "lodestar",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(from, "slashing-data.json"),
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
			slashingImportCmd.SetOutput(io.Discard)
			err := slashingImportCmd.Execute()

			assert.Nil(t, err)
		})
	}
}

func TestSlashingImport_Errors(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	tests := []struct {
		name string
		args []string
		run  bool
		err  error
	}{
		{
			name: "invalid network",
			args: []string{"lighthouse", "--network", "invalid_network"},
			err:  errors.New("invalid network: invalid_network"),
		},
		{
			name: "action error",
			args: []string{"lighthouse"},
			run:  true,
			err:  errors.New("action error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			actions := sedge_mocks.NewMockSedgeActions(ctrl)
			if tt.run {
				actions.EXPECT().ImportSlashingInterchangeData(gomock.Any()).Return(errors.New("action error")).Times(1)
			}

			slashingImportCmd := cli.SlashingImportCmd(actions)
			slashingImportCmd.SetArgs(tt.args)
			slashingImportCmd.SetOutput(io.Discard)
			err := slashingImportCmd.Execute()

			assert.EqualError(t, err, tt.err.Error())
		})
	}
}
