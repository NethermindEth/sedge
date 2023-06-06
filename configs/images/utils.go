package clientsimages

import (
	"fmt"
	"regexp"
)

type ClientsImagesGuesser struct{}

// Nethermind
func (cig ClientsImagesGuesser) NethermindRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^nethermind/nethermind:\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) IsNethermind(image string) bool {
	re := cig.NethermindRegexp()
	return re.MatchString(image)
}

// Erigon
func (cig ClientsImagesGuesser) ErigonRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^thorax/erigon:v\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) IsErigon(image string) bool {
	re := cig.ErigonRegexp()
	return re.MatchString(image)
}

// Geth
func (cig ClientsImagesGuesser) GethRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^ethereum/client-go:v\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) IsGeth(image string) bool {
	re := cig.GethRegexp()
	return re.MatchString(image)
}

// Besu
func (cig ClientsImagesGuesser) BesuRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^hyperledger/besu:\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) IsBesu(image string) bool {
	re := cig.BesuRegexp()
	return re.MatchString(image)
}

// Lighthouse
func (cig ClientsImagesGuesser) LighthouseConsensusRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^sigp/lighthouse:v\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) LighthouseValidatorRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^sigp/lighthouse:v\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) IsLighthouseConsensus(image string) bool {
	re := cig.LighthouseConsensusRegexp()
	return re.MatchString(image)
}

func (cig ClientsImagesGuesser) IsLighthouseValidator(image string) bool {
	re := cig.LighthouseValidatorRegexp()
	return re.MatchString(image)
}

// Teku
func (cig ClientsImagesGuesser) TekuConsensusRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^consensys/teku:\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) TekuValidatorRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^consensys/teku:\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) IsTekuConsensus(image string) bool {
	re := cig.TekuConsensusRegexp()
	return re.MatchString(image)
}

func (cig ClientsImagesGuesser) IsTekuValidator(image string) bool {
	re := cig.TekuValidatorRegexp()
	return re.MatchString(image)
}

// Lodestar
func (cig ClientsImagesGuesser) LodestarConsensusRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^chainsafe/lodestar:v\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) LodestarValidatorRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^chainsafe/lodestar:v\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) IsLodestarConsensus(image string) bool {
	re := cig.LodestarConsensusRegexp()
	return re.MatchString(image)
}

func (cig ClientsImagesGuesser) IsLodestarValidator(image string) bool {
	re := cig.LodestarValidatorRegexp()
	return re.MatchString(image)
}

// Prysm
func (cig ClientsImagesGuesser) PrysmConsensusRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^gcr.io/prysmaticlabs/prysm/beacon-chain:v\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) PrysmValidatorRegexp() *regexp.Regexp {
	return regexp.MustCompile(`^gcr.io/prysmaticlabs/prysm/validator:v\d+\.\d+\.\d+$`)
}

func (cig ClientsImagesGuesser) IsPrysmConsensus(image string) bool {
	re := cig.PrysmConsensusRegexp()
	return re.MatchString(image)
}

func (cig ClientsImagesGuesser) IsPrysmValidator(image string) bool {
	re := cig.PrysmValidatorRegexp()
	return re.MatchString(image)
}

// // Execution
func (cig ClientsImagesGuesser) IsExecution(image string) bool {
	return cig.IsNethermind(image) || cig.IsErigon(image) || cig.IsGeth(image) || cig.IsBesu(image)
}

// // Consensus
func (cig ClientsImagesGuesser) IsConsensus(image string) bool {
	return cig.IsLighthouseConsensus(image) ||
		cig.IsTekuConsensus(image) ||
		cig.IsLodestarConsensus(image) ||
		cig.IsPrysmConsensus(image)
}

// // Validator
func (cig ClientsImagesGuesser) IsValidator(image string) bool {
	return cig.IsLighthouseValidator(image) ||
		cig.IsTekuValidator(image) ||
		cig.IsLodestarValidator(image) ||
		cig.IsPrysmValidator(image)
}

func (cig ClientsImagesGuesser) GuessClientImageUpdate(image string, updates ClientsImages) (string, error) {
	for _, candidate := range []struct {
		re     *regexp.Regexp
		update string
	}{
		{cig.NethermindRegexp(), updates.Execution().Nethermind().String()},
		{cig.ErigonRegexp(), updates.Execution().Erigon().String()},
		{cig.GethRegexp(), updates.Execution().Geth().String()},
		{cig.BesuRegexp(), updates.Execution().Besu().String()},
		{cig.LighthouseConsensusRegexp(), updates.Consensus().Lighthouse().String()},
		{cig.TekuConsensusRegexp(), updates.Consensus().Teku().String()},
		{cig.LodestarConsensusRegexp(), updates.Consensus().Lodestar().String()},
		{cig.PrysmConsensusRegexp(), updates.Consensus().Prysm().String()},
		{cig.LighthouseValidatorRegexp(), updates.Validator().Lighthouse().String()},
		{cig.TekuValidatorRegexp(), updates.Validator().Teku().String()},
		{cig.LodestarValidatorRegexp(), updates.Validator().Lodestar().String()},
		{cig.PrysmValidatorRegexp(), updates.Validator().Prysm().String()},
	} {
		if candidate.re.MatchString(image) {
			return candidate.update, nil
		}
	}

	return "", fmt.Errorf("could not find update for image %s", image)
}
