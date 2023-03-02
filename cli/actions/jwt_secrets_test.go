package actions_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/stretchr/testify/assert"
)

func TestCreateJwtSecrets(t *testing.T) {
	tempDir := t.TempDir()
	tests := []struct {
		name    string
		options actions.CreateJWTSecretOptions
		err     bool
	}{
		{
			name: "Wrong Jwt Path",
			options: actions.CreateJWTSecretOptions{
				GenerationPath: tempDir,
				JWTPath:        tempDir,
				Network:        "mainnet",
			},
		},
		{
			name: "Missing network",
			options: actions.CreateJWTSecretOptions{
				GenerationPath: tempDir,
				JWTPath:        "",
			},
			err: true,
		},
		{
			name: "Generation of jwt",
			options: actions.CreateJWTSecretOptions{
				GenerationPath: tempDir,
				JWTPath:        "",
				Network:        "mainnet",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sedgeActions := actions.NewSedgeActions(actions.SedgeActionsOptions{})
			if tc.options.JWTPath != "" {
				file, err := os.Create(filepath.Join(tc.options.JWTPath, "jwtSecret"))
				if err != nil {
					t.Error(err)
				}
				tc.options.JWTPath = filepath.Join(tc.options.JWTPath, "jwtSecret")
				defer os.Remove(file.Name())
				defer file.Close()
			}
			jwtPath, err := sedgeActions.CreateJWTSecrets(tc.options)
			if tc.err {
				assert.NotNil(t, err)
				return
			} else {
				assert.Nil(t, err)
			}
			if jwtPath == "" {
				return
			}
			assert.FileExists(t, jwtPath)
		})
	}
}
