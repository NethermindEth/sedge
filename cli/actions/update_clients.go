package actions

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	clientsimages "github.com/NethermindEth/sedge/configs/images"
)

type UpdateClientsOptions struct {
	GenerationPath  string
	UpdateExecution bool
	ExecutionImage  string
	UpdateConsensus bool
	ConsensusImage  string
	UpdateValidator bool
	ValidatorImage  string
}

var (
	executionImageEnvVarRegexp = regexp.MustCompile(`^ *(?P<VAR>EC_IMAGE_VERSION) *= *(?P<VAL>.+) *$`)
	executionImageEnvVar       = "EC_IMAGE_VERSION"
	consensusImageEnvVarRegexp = regexp.MustCompile(`^ *(?P<VAR>CC_IMAGE_VERSION) *= *(?P<VAL>.+) *$`)
	consensusImageEnvVar       = "CC_IMAGE_VERSION"
	validatorImageEnvVarRegexp = regexp.MustCompile(`^ *(?P<VAR>VL_IMAGE_VERSION) *= *(?P<VAL>.+) *$`)
	validatorImageEnvVar       = "VL_IMAGE_VERSION"
)

func (s *sedgeActions) getFinalImage(actualImage, newImage string) (string, error) {
	if newImage == "" {
		var guesser clientsimages.ClientsImagesGuesser
		return guesser.GuessClientImageUpdate(actualImage, s.clientsImages)
	}
	return newImage, nil
}

func (s *sedgeActions) UpdateClients(options UpdateClientsOptions) error {
	envFilePath := path.Join(options.GenerationPath, configs.DefaultEnvFileName)
	envFileRaw, err := os.ReadFile(envFilePath)
	if err != nil {
		return err
	}
	envFile := string(envFileRaw)
	lines := strings.Split(envFile, "\n")
	for idx, line := range lines {
		if options.UpdateExecution {
			result := executionImageEnvVarRegexp.FindStringSubmatch(line)
			if len(result) >= 3 {
				finalImage, err := s.getFinalImage(
					result[2],
					options.ExecutionImage,
				)
				if err != nil {
					return err
				}
				lines[idx] = fmt.Sprintf("%s=%s", executionImageEnvVar, finalImage)
			}
		}
		if options.UpdateConsensus {
			result := consensusImageEnvVarRegexp.FindStringSubmatch(line)
			if len(result) >= 3 {
				finalImage, err := s.getFinalImage(
					result[2],
					options.ConsensusImage,
				)
				if err != nil {
					return err
				}
				lines[idx] = fmt.Sprintf("%s=%s", consensusImageEnvVar, finalImage)
			}
		}
		if options.UpdateValidator {
			result := validatorImageEnvVarRegexp.FindStringSubmatch(line)
			if len(result) >= 3 {
				finalImage, err := s.getFinalImage(
					result[2],
					options.ValidatorImage,
				)
				if err != nil {
					return err
				}
				lines[idx] = fmt.Sprintf("%s=%s", validatorImageEnvVar, finalImage)
			}
		}
	}
	return nil
}
