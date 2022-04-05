package generate

import (
	"fmt"
	"testing"
)

func validateGeneratedConfig(t *testing.T, path string) {
	//TODO: validate generated config
}

func TestGenerateConfig(t *testing.T) {
	inputs := [...]struct {
		path  string
		isErr bool
	}{
		{t.TempDir(), false},
		{"", true},
	}

	for _, input := range inputs {
		descr := fmt.Sprintf("GenerateScripts(%s,)", input.path)

		if err := GenerateConfig(input.path); input.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !input.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
		validateGeneratedConfig(t, input.path)
	}
}
