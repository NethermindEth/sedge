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
