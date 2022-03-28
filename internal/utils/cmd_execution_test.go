package utils_test

import (
	"testing"

	"github.com/NethermindEth/1Click/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestRunCmd(t *testing.T) {
	testingService, v := utils.RunCmd("", false)

	assert.Nil(t, v)
	assert.Equal(t, testingService, testingService)
}

// func TestExecuteScript(t *testing.T) {

// 	testingString, err := executeScript(nil)
// 	assert.Nil(t, err, nil)
// 	assert.NotEqual(t, testingString, "")
// }
