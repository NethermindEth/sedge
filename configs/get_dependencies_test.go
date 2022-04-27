package configs_test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetDependencies(t *testing.T) {
	testingService := viper.GetStringSlice("dependencies")
	assert.NotEmpty(t, testingService)
}
