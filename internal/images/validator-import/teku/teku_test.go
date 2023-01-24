package teku_test

import (
	"embed"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/internal/images/validator-import/teku"
	"github.com/stretchr/testify/assert"
)

//go:embed context
var tekuContext embed.FS

func TestInitContext(t *testing.T) {
	entries, err := tekuContext.ReadDir("context")
	assert.NoError(t, err)
	for _, entry := range entries {
		t.Run(entry.Name(), func(t *testing.T) {
			contextDir, err := teku.InitContext()
			assert.NoError(t, err)
			data, err := ioutil.ReadFile(filepath.Join(contextDir, entry.Name()))
			assert.NoError(t, err)
			expectedData, err := tekuContext.ReadFile("context/" + entry.Name())
			assert.NoError(t, err)
			assert.Equal(t, expectedData, data)
		})
	}
}
