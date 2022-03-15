package configs_test

import (
	"testing"

	"github.com/NethermindEth/1Click/configs"
)

func TestLogConfig(t *testing.T) {
	var testLogService configs.LogConfig

	if len(testLogService.Level) != 0 {
		t.Skip()
	}
}
