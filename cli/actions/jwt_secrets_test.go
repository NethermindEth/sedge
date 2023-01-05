package actions_test

import (
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestCreateJwtSecrets(t *testing.T) {
	tests := []struct {
		name    string
		options actions.CreateJwtSecretOptions
		err     bool
	}{
		{
			name: "Wrong Jwt Path",
			options: actions.CreateJwtSecretOptions{
				GenerationPath: t.TempDir(),
				JwtPath:        filepath.Join("a", "b", "c"),
				Network:        "mainnet",
			},
		},
		{
			name: "Missing network",
			options: actions.CreateJwtSecretOptions{
				GenerationPath: t.TempDir(),
				JwtPath:        "",
			},
			err: true,
		},
		{
			name: "Generation of jwt",
			options: actions.CreateJwtSecretOptions{
				GenerationPath: t.TempDir(),
				JwtPath:        "",
				Network:        "mainnet",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sedgeActions := actions.NewSedgeActions(nil, nil, nil)
			jwtPath, err := sedgeActions.CreateJwtSecrets(tc.options)
			if err != nil {
				if !tc.err {
					t.Errorf("Got: %v, want error: %v", err, tc.err)
				}
				return
			}
			if tc.err || tc.options.JwtPath != "" {
				return
			}
			assert.FileExists(t, jwtPath)
		})
	}
}
