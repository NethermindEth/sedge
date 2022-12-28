package actions_test

import (
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
	"github.com/stretchr/testify/assert"
)

func TestRunContainers(t *testing.T) {
	tests := []struct {
		name            string
		options         actions.RunContainersOptions
		expectedCommand string
	}{
		{
			name: "with services",
			options: actions.RunContainersOptions{
				GenerationPath: "/a/b/c/d",
				Services:       []string{"validator", "consensus", "execution"},
			},
			expectedCommand: "docker compose -f /a/b/c/d/docker-compose.yml up -d validator consensus execution",
		},
		{
			name: "without services",
			options: actions.RunContainersOptions{
				GenerationPath: "/a/b/c/d",
				Services:       []string{},
			},
			expectedCommand: "docker compose -f /a/b/c/d/docker-compose.yml up -d",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			commandRunner := &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					assert.Equal(t, tc.expectedCommand, c.Cmd)
					return "", nil
				},
			}
			sedgeActions := actions.NewSedgeActions(nil, nil, commandRunner)
			sedgeActions.RunContainers(tc.options)
		})
	}
}
