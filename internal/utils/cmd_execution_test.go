package utils

import (
	"testing"
	"text/template"

	"github.com/NethermindEth/1Click/configs"
	"github.com/stretchr/testify/assert"
)

func TestRunCmd(t *testing.T) {
	testingService, v := RunCmd("", false)

	assert.Nil(t, v)
	assert.Equal(t, testingService, testingService)
}

func TestExecuteScript(t *testing.T) {
	cmd := configs.DockerPsCMD
	tmp, err := template.New("script").Parse(cmd)
	if err != nil {
		t.FailNow()
	}

	script := Script{
		Tmp:       tmp,
		GetOutput: true,
		Data:      struct{}{},
	}
	testingString, err := executeScript(script)
	assert.Nil(t, err, nil)
	assert.NotEmpty(t, testingString)
}
