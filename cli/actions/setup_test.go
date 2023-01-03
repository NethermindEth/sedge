/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package actions_test

import (
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
	"github.com/stretchr/testify/assert"
)

func TestSetupContainers(t *testing.T) {
	tests := []struct {
		name              string
		options           actions.SetupContainersOptions
		expectedBuildCmd  string
		expectedPullCmd   string
		expectedCreateCmd string
	}{
		{
			name: "with services",
			options: actions.SetupContainersOptions{
				GenerationPath: "/a/b/c/d",
				Services:       []string{"validator", "consensus", "execution"},
			},
			expectedBuildCmd:  "docker compose -f /a/b/c/d/docker-compose.yml build validator consensus execution",
			expectedPullCmd:   "docker compose -f /a/b/c/d/docker-compose.yml pull validator consensus execution",
			expectedCreateCmd: "docker compose -f /a/b/c/d/docker-compose.yml create validator consensus execution",
		},
		{
			name: "without services",
			options: actions.SetupContainersOptions{
				GenerationPath: "/a/b/c/d",
				Services:       []string{},
			},
			expectedBuildCmd:  "docker compose -f /a/b/c/d/docker-compose.yml build",
			expectedPullCmd:   "docker compose -f /a/b/c/d/docker-compose.yml pull",
			expectedCreateCmd: "docker compose -f /a/b/c/d/docker-compose.yml create",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			commandRunner := &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, error) {
					assert.Contains(t, []string{tc.expectedBuildCmd, tc.expectedPullCmd, tc.expectedCreateCmd}, c.Cmd)
					return "", nil
				},
			}
			sedgeActions := actions.NewSedgeActions(nil, nil, commandRunner)
			sedgeActions.SetupContainers(tc.options)
		})
	}
}
