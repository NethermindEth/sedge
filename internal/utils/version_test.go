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
		{
			name:     "Correct version",
			response: "[\n  {\n    \"name\": \"v0.3.0\",\n    \"zipball_url\": \"https://api.github.com/repos/NethermindEth/sedge/zipball/refs/tags/v0.3.0\",\n    \"tarball_url\": \"https://api.github.com/repos/NethermindEth/sedge/tarball/refs/tags/v0.3.0\",\n    \"commit\": {\n      \"sha\": \"8eedbb2f400eeb2decc4bcf3198884ede3d83bfb\",\n      \"url\": \"https://api.github.com/repos/NethermindEth/sedge/commits/8eedbb2f400eeb2decc4bcf3198884ede3d83bfb\"\n    },\n    \"node_id\": \"REF_kwDOGoTRGLByZWZzL3RhZ3MvdjAuMy4w\"\n  },\n  {\n    \"name\": \"v0.2.0\",\n    \"zipball_url\": \"https://api.github.com/repos/NethermindEth/sedge/zipball/refs/tags/v0.2.0\",\n    \"tarball_url\": \"https://api.github.com/repos/NethermindEth/sedge/tarball/refs/tags/v0.2.0\",\n    \"commit\": {\n      \"sha\": \"10445f607fc0d44abb42aead5fd0f6fd051a0c95\",\n      \"url\": \"https://api.github.com/repos/NethermindEth/sedge/commits/10445f607fc0d44abb42aead5fd0f6fd051a0c95\"\n    },\n    \"node_id\": \"REF_kwDOGoTRGLByZWZzL3RhZ3MvdjAuMi4w\"\n  },\n  {\n    \"name\": \"0.1.2\",\n    \"zipball_url\": \"https://api.github.com/repos/NethermindEth/sedge/zipball/refs/tags/0.1.2\",\n    \"tarball_url\": \"https://api.github.com/repos/NethermindEth/sedge/tarball/refs/tags/0.1.2\",\n    \"commit\": {\n      \"sha\": \"90d0cba09c5c2c2143c9d0e32b02f7eacf99b0f7\",\n      \"url\": \"https://api.github.com/repos/NethermindEth/sedge/commits/90d0cba09c5c2c2143c9d0e32b02f7eacf99b0f7\"\n    },\n    \"node_id\": \"REF_kwDOGoTRGK9yZWZzL3RhZ3MvMC4xLjI\"\n  },\n  {\n    \"name\": \"0.1.1\",\n    \"zipball_url\": \"https://api.github.com/repos/NethermindEth/sedge/zipball/refs/tags/0.1.1\",\n    \"tarball_url\": \"https://api.github.com/repos/NethermindEth/sedge/tarball/refs/tags/0.1.1\",\n    \"commit\": {\n      \"sha\": \"6b551d9f270ec45332a30a76d738db282017a11f\",\n      \"url\": \"https://api.github.com/repos/NethermindEth/sedge/commits/6b551d9f270ec45332a30a76d738db282017a11f\"\n    },\n    \"node_id\": \"REF_kwDOGoTRGK9yZWZzL3RhZ3MvMC4xLjE\"\n  },\n  {\n    \"name\": \"0.1.0\",\n    \"zipball_url\": \"https://api.github.com/repos/NethermindEth/sedge/zipball/refs/tags/0.1.0\",\n    \"tarball_url\": \"https://api.github.com/repos/NethermindEth/sedge/tarball/refs/tags/0.1.0\",\n    \"commit\": {\n      \"sha\": \"e15fdc5f5d6e7e11961e7c6647a27330d3fbbdab\",\n      \"url\": \"https://api.github.com/repos/NethermindEth/sedge/commits/e15fdc5f5d6e7e11961e7c6647a27330d3fbbdab\"\n    },\n    \"node_id\": \"REF_kwDOGoTRGK9yZWZzL3RhZ3MvMC4xLjA\"\n  }\n]",
			want:     true,
			tag:      "v0.3.0",
			status:   200,
			err:      nil,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "https://api.github.com/repos/NethermindEth/sedge/tags",
				httpmock.NewStringResponder(tt.status, tt.response))
			Version = tt.tag
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
