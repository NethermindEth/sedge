package configs_test

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetDependencies(t *testing.T) {
	testingService := viper.GetStringSlice("dependencies")
	if len(testingService) > 0 {
		t.SkipNow()
	}
}
