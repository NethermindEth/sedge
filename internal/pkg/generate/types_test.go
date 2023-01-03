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
package generate

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestDockerComposeData_WithConsensusClient(t *testing.T) {
	tests := []struct {
		name string
		data DockerComposeData
		want bool
	}{
		{
			name: "with consensus client",
			data: DockerComposeData{
				Services: []string{"execution", "consensus", "validator"},
			},
			want: true,
		},
		{
			name: "without consensus client",
			data: DockerComposeData{
				Services: []string{"execution", "validator"},
			},
			want: false,
		},
		{
			name: "with nil services",
			data: DockerComposeData{
				Services: nil,
			},
			want: false,
		},
	}
	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			out := tC.data.WithConsensusClient()
			assert.Equal(t, out, tC.want)
		})
	}
}
