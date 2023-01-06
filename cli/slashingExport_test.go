package cli_test

import (
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/cli/actions/mock"
	"github.com/NethermindEth/sedge/configs"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSlashingExport_ValidatorIsRequired(t *testing.T) {
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
			args:        []string{"--network", "sepolia", "--out", "slashing_data.json"},
			expectedErr: cli.ErrInvalidNumberOfArguments,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {})

		slashingImportCmd := cli.SlashingExportCmd(nil)
		slashingImportCmd.SetArgs(tt.args)
		err := slashingImportCmd.Execute()
		assert.ErrorIs(t, err, tt.expectedErr)
	}
}

func TestSlashingExport_Params(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		actionOptions actions.SlashingExportOptions
	}{
		{
			name: "validator argument",
			args: []string{"lighthouse"},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "lighthouse",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultDockerComposeScriptsPath,
				Out:             filepath.Join(configs.DefaultDockerComposeScriptsPath, "slashing_export.json"),
			},
		},
		{
			name: "network flag",
			args: []string{"lighthouse", "--network", "sepolia"},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "lighthouse",
				Network:         "sepolia",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultDockerComposeScriptsPath,
				Out:             filepath.Join(configs.DefaultDockerComposeScriptsPath, "slashing_export.json"),
			},
		},
		{
			name: "network shorthand flag",
			args: []string{"lighthouse", "-n", "sepolia"},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "lighthouse",
				Network:         "sepolia",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultDockerComposeScriptsPath,
				Out:             filepath.Join(configs.DefaultDockerComposeScriptsPath, "slashing_export.json"),
			},
		},
		{
			name: "stop-validator flag",
			args: []string{"prysm", "--stop-validator"},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "prysm",
				Network:         "mainnet",
				StopValidator:   true,
				StartValidator:  false,
				GenerationPath:  configs.DefaultDockerComposeScriptsPath,
				Out:             filepath.Join(configs.DefaultDockerComposeScriptsPath, "slashing_export.json"),
			},
		},
		{
			name: "start-validator flag",
			args: []string{"teku", "--start-validator"},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "teku",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  true,
				GenerationPath:  configs.DefaultDockerComposeScriptsPath,
				Out:             filepath.Join(configs.DefaultDockerComposeScriptsPath, "slashing_export.json"),
			},
		},
		{
			name: "path flag",
			args: []string{"teku", "--path", filepath.Join("custom", "dir")},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "teku",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  filepath.Join("custom", "dir"),
				Out:             filepath.Join("custom", "dir", "slashing_export.json"),
			},
		},
		{
			name: "path shorthand flag",
			args: []string{"teku", "-p", filepath.Join("custom", "dir")},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "teku",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  filepath.Join("custom", "dir"),
				Out:             filepath.Join("custom", "dir", "slashing_export.json"),
			},
		},
		{
			name: "out flag",
			args: []string{"lodestar", "--out", filepath.Join("custom", "out", "file.json")},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "lodestar",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultDockerComposeScriptsPath,
				Out:             filepath.Join("custom", "out", "file.json"),
			},
		},
		{
			name: "out shorthand flag",
			args: []string{"lodestar", "-o", filepath.Join("custom", "out", "file.json")},
			actionOptions: actions.SlashingExportOptions{
				ValidatorClient: "lodestar",
				Network:         "mainnet",
				StopValidator:   false,
				StartValidator:  false,
				GenerationPath:  configs.DefaultDockerComposeScriptsPath,
				Out:             filepath.Join("custom", "out", "file.json"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			actions := mock_actions.NewMockSedgeActions(ctrl)
			actions.EXPECT().ExportSlashingInterchangeData(tt.actionOptions).Times(1)

			slashingImportCmd := cli.SlashingExportCmd(actions)
			slashingImportCmd.SetArgs(tt.args)
			err := slashingImportCmd.Execute()

			assert.Nil(t, err)
		})
	}
}
