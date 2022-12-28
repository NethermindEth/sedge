package actions

import (
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/generate"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type GenerateComposeOptions struct {
	GenerationData *generate.GenData
	GenerationPath string
}

func (s *sedgeActions) GenerateCompose(options GenerateComposeOptions) error {
	// Create scripts directory if not exists
	if _, err := os.Stat(options.GenerationPath); os.IsNotExist(err) {
		err = os.MkdirAll(options.GenerationPath, 0o755)
		if err != nil {
			return err
		}
	}

	log.Info(configs.GeneratingDockerComposeScript)
	// open output file
	out, err := os.Create(filepath.Join(options.GenerationPath, configs.DefaultDockerComposeScriptName))
	if err != nil {
		return generate.CanNotCreateComposeFileError
	}
	err = generate.ComposeFile(options.GenerationData, out)
	if err != nil {
		return err
	}

	log.Info(configs.GeneratingEnvFile)
	// open output file
	out, err = os.Create(filepath.Join(options.GenerationPath, configs.DefaultEnvFileName))
	err = generate.EnvFile(options.GenerationData, out)
	if err != nil {
		return err
	}

	err = generate.CleanGenerated(options.GenerationPath)
	if err != nil {
		return err
	}

	return nil
}
