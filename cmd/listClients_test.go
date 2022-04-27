package cmd

import (
	"fmt"
	"testing"

	"github.com/NethermindEth/1Click/internal/pkg/clients"
	"github.com/stretchr/testify/assert"
)

func TestBuildData(t *testing.T) {
	supportClients := clients.GetSupportedClients

	testingService, err := buildData(supportClients)
	assert.NotNil(t, testingService)
	assert.Nil(t, err)
	fmt.Println(testingService)

	// testingService, err = buildData(configs.GetConfigClients(supportClients["execution"]))
	// assert.Nil(t, err)
	// assert.NotNil(t, testingService)
	// fmt.Println(testingService, err)

}
