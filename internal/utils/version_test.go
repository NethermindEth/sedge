package utils

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestVersion(t *testing.T) {
	tests := []struct {
		name     string
		response string
		status   int
		want     bool
		tag      string
		err      error
	}{
		{
			name:     "No Version",
			response: `[]`,
			want:     false,
			tag:      "",
			status:   200,
			err:      ErrorNoTag,
		},
		{
			name:     "Error Checking Version",
			response: `[]`,
			want:     false,
			tag:      "",
			status:   500,
			err:      ErrorCheckingVersion,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.github.com/repos/NethermindEth/sedge/tags",
				httpmock.NewStringResponder(tt.status, tt.response))
			got, err := IsLatestVersion()
			if err != nil {
				if err != tt.err {
					t.Errorf("LatestVersionOnGithub() error = %v, wantErr %v", err, tt.err)
				}
				return
			}

			if got != tt.want {
				t.Errorf("latestVersionOnGithub() = %v, want %v", got, tt.want)
			}

			if tt.want && tt.tag != CurrentVersion() {
				t.Errorf("latestVersionOnGithub() = %v, want %v", CurrentVersion(), tt.tag)
			}
		})
	}
}
