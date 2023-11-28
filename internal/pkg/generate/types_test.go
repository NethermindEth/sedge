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
				Services: []string{"execution", "consensus", "validator", "mev-boost"},
			},
			want: true,
		},
		{
			name: "with consensus client",
			data: DockerComposeData{
				Services: []string{"consensus"},
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
			name: "without consensus client",
			data: DockerComposeData{
				Services: []string{"execution", "mev-boost"},
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
			assert.Equal(t, out, tC.want, "services: %v", tC.data.Services)
		})
	}
}

func TestDockerComposeData_WithStarknetClient(t *testing.T) {
	tests := []struct {
		name string
		data DockerComposeData
		want bool
		}{
		{
			name: "with consensus and execution client",
			data: DockerComposeData{
				Services: []string{"execution", "consensus", "starknet"},
			},
			want: true,
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
			out := tC.data.WithFullFlagStarknet()
			assert.Equal(t, out, tC.want, "services: %v", tC.data.Services)
		})
	}
}

func TestDockerComposeData_WithValidatorClient(t *testing.T) {
	tests := []struct {
		name string
		data DockerComposeData
		want bool
	}{
		{
			name: "with validator client",
			data: DockerComposeData{
				Services: []string{"execution", "consensus", "validator", "mev-boost"},
			},
			want: true,
		},
		{
			name: "with validator client",
			data: DockerComposeData{
				Services: []string{"validator"},
			},
			want: true,
		},
		{
			name: "without validator client",
			data: DockerComposeData{
				Services: []string{"execution", "consensus", "mev-boost"},
			},
			want: false,
		},
		{
			name: "without validator client",
			data: DockerComposeData{
				Services: []string{"execution"},
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
			out := tC.data.WithValidatorClient()
			assert.Equal(t, out, tC.want, "services: %v", tC.data.Services)
		})
	}
}

func TestDockerComposeData_WithMevBoostClient(t *testing.T) {
	tests := []struct {
		name string
		data EnvData
		want bool
	}{
		{
			name: "with mev-boost client",
			data: EnvData{
				Services: []string{"execution", "consensus", "validator", "mev-boost"},
			},
			want: true,
		},
		{
			name: "with mev-boost client",
			data: EnvData{
				Services: []string{"mev-boost"},
			},
			want: true,
		},
		{
			name: "without mev-boost client",
			data: EnvData{
				Services: []string{"execution", "consensus"},
			},
			want: false,
		},
		{
			name: "without mev-boost client",
			data: EnvData{
				Services: []string{"execution"},
			},
			want: false,
		},
		{
			name: "with nil services",
			data: EnvData{
				Services: nil,
			},
			want: false,
		},
	}
	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			out := tC.data.WithMevBoostClient()
			assert.Equal(t, out, tC.want, "services: %v", tC.data.Services)
		})
	}
}
