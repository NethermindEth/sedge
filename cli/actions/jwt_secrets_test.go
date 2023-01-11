package actions_test

import (
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/configs"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestCreateJwtSecrets(t *testing.T) {
	configs.InitNetworksConfigs()
	tests := []struct {
		name    string
		options actions.CreateJWTSecretOptions
		err     bool
	}{
		{
			name: "Wrong Jwt Path",
			options: actions.CreateJWTSecretOptions{
				GenerationPath: t.TempDir(),
				JWTPath:        filepath.Join("a", "b", "c"),
				Network:        "mainnet",
			},
		},
		{
			name: "Missing network",
			options: actions.CreateJWTSecretOptions{
				GenerationPath: t.TempDir(),
				JWTPath:        "",
			},
			err: true,
		},
		{
			name: "Generation of jwt",
			options: actions.CreateJWTSecretOptions{
				GenerationPath: t.TempDir(),
				JWTPath:        "",
				Network:        "mainnet",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sedgeActions := actions.NewSedgeActions(nil, nil, nil)
			jwtPath, err := sedgeActions.CreateJwtSecrets(tc.options)
			if err != nil {
				assert.True(t, tc.err)
				return
			}
			assert.True(t, !tc.err)
			if tc.options.JWTPath != "" {
				return
			}
			assert.FileExists(t, jwtPath)
		})
	}
}
