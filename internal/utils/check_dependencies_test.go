package utils_test

import (
	"os/exec"
	"testing"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCheckDependencies(t *testing.T) {
	dependencies := configs.GetDependencies()
	var pending []string
	var err error
	for _, dependency := range dependencies {
		_, err := exec.LookPath(dependency)
		if err != nil {
			log.Errorf(configs.DependencyNotInstalledError, dependency)
			pending = append(pending, dependency)
		}
	}

	testCheckDependencies := utils.CheckDependencies(dependencies)

	assert.Nil(t, err)
	assert.Equal(t, testCheckDependencies, pending)
}
