package utils

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/NethermindEth/1Click/configs"
	"github.com/stretchr/testify/assert"
)

func TestGetScriptPath(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.FailNow()
	}

	distro, err := GetOSInfo()
	assert.Nil(t, err)
	path := fmt.Sprintf("setup/%s/%s/%s_%s.sh", runtime.GOOS, "docker", distro.Name, distro.Version)

	assert.Equal(t, distro.Name, "pop")
	fmt.Println(path)
	assert.NotEmpty(t, path)
}

func TestDependencySupported(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.FailNow()
	}
	distro, err := GetOSInfo()
	scriptPathDocker, _, err := getScriptPath("docker")
	assert.Nil(t, err, nil)
	assert.Equal(t, scriptPathDocker, "setup/linux/docker/"+distro.Name+"_"+distro.Version+".sh")

	scriptPathDockerCompose, _, err := getScriptPath("docker-compose")
	assert.Nil(t, err)
	fmt.Println(scriptPathDockerCompose)
	assert.EqualValues(t, scriptPathDockerCompose, "setup/linux/docker-compose/"+distro.Name+"_"+distro.Version+".sh")
}

func TestInstallDependency(t *testing.T) {
	testingService := InstallDependency("docker")
	assert.Nil(t, testingService)

}

func TestShowInstructions(t *testing.T) {
	testingService := ShowInstructions("docker")
	fmt.Println(testingService)
	assert.Nil(t, testingService)
}

func TestHandleInstructions(t *testing.T) {
	dependencies := configs.GetDependencies()
	// assert.NotEmpty(t, dependencies)
	pending := CheckDependencies(dependencies)
	// assert.NotEmpty(t, pending)
	testingService := HandleInstructions(pending, ShowInstructions)

	assert.Nil(t, testingService)

}
