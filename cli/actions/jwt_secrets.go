package actions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/crypto"
	log "github.com/sirupsen/logrus"
)

type CreateJWTSecretOptions struct {
	JWTPath        string
	Network        string
	GenerationPath string
}

func (s *sedgeActions) CreateJWTSecrets(options CreateJWTSecretOptions) (string, error) {
	if options.Network == "" {
		return "", ErrorNetworkNotFound
	}
	// Generate JWT secret if necessary
	var err error
	jwtPath := options.JWTPath
	if jwtPath == "" && configs.NetworksConfigs()[options.Network].RequireJWT {
		return handleJWTSecret(options.GenerationPath)
	} else if filepath.IsAbs(jwtPath) { // Ensure jwtPath is absolute
		if jwtPath, err = filepath.Abs(jwtPath); err != nil {
			return jwtPath, err
		}
	}
	return jwtPath, nil
}

func handleJWTSecret(generationPath string) (string, error) {
	log.Info(configs.GeneratingJWTSecret)

	jwtSecret, err := crypto.GenerateJWTSecret()
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	jwtPath, err := filepath.Abs(filepath.Join(generationPath, "jwtsecret"))
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	if err = os.MkdirAll(filepath.Dir(jwtPath), 0o755); err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	err = os.WriteFile(jwtPath, []byte(jwtSecret), 0o755)
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}

	log.Info(configs.JWTSecretGenerated)
	return jwtPath, nil
}
