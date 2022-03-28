package utils

import (
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
	assert.Equal(t, distro.Name, "pop")
}

func TestDependencySupported(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.FailNow()
	}

	scriptPathDocker, _, err := getScriptPath("docker")
	assert.Nil(t, err, nil)
	assert.NotEmpty(t, scriptPathDocker)

	scriptPathDockerCompose, _, err := getScriptPath("docker-compose")
	assert.Nil(t, err)
	assert.NotEmpty(t, scriptPathDockerCompose)
}

// func TestInstallDependency(t *testing.T) {
// 	scriptPath, _, err := getScriptPath("")
// 	assert.Nil(t, err)
// 	rawScript, err := templates.Setup.ReadFile(scriptPath)
// 	assert.Nil(t, err)
// 	_, err = template.New("script").Parse(string(rawScript))
// 	assert.Nil(t, err)

// 	testingService := InstallDependency("")
// 	assert.Nil(t, testingService)
// }

// func TestShowInstructions(t *testing.T) {
// 	testingService := ShowInstructions("docker-compose")
// 	fmt.Println(testingService)
// 	assert.Nil(t, testingService)
// }

func TestHandleInstructions(t *testing.T) {
	dependencies := configs.GetDependencies()
	// assert.NotEmpty(t, dependencies)
	pending := CheckDependencies(dependencies)
	// assert.NotEmpty(t, pending)
	testingService := HandleInstructions(pending, ShowInstructions)

	assert.Nil(t, testingService)

}
