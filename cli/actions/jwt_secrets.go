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
		return "", ErrNetworkNotFound
	}
	// Generate JWT secret if necessary
	var err error
	jwtPath := options.JWTPath
	if jwtPath == "" && !configs.NetworksConfigs()[options.Network].NoJWT {
		return handleJWTSecret(options.GenerationPath)
	}
	// Handle relative paths by joining with generation path
	if !filepath.IsAbs(jwtPath) {
		jwtPath = filepath.Join(options.GenerationPath, jwtPath)
	}
	// Create parent directories if they don't exist
	if err = os.MkdirAll(filepath.Dir(jwtPath), 0o755); err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
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

	// Convert the absolute path to a relative path
	relPath, err := filepath.Rel(generationPath, jwtPath)
	if err != nil {
		return "", fmt.Errorf(configs.GenerateJWTSecretError, err)
	}
	return relPath, nil
}
