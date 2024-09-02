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
package package_handler

import (
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/NethermindEth/sedge/internal/monitoring/package_handler/testdata"
)

func TestManifest_ValidateFromYML(t *testing.T) {
	afs := afero.NewMemMapFs()
	testDir, err := afero.TempDir(afs, "", "test")
	require.NoError(t, err)
	testdata.SetupDir(t, "manifests", testDir, afs)

	tests := []struct {
		name      string
		filePath  string
		wantError string
	}{
		{
			name:      "Full OK Manifest",
			filePath:  "full-ok/pkg/manifest.yml",
			wantError: "",
		},
		{
			name:      "Invalid Fields Manifest",
			filePath:  "invalid-fields/pkg/manifest.yml",
			wantError: "Invalid manifest file: Invalid hardware requirements -> invalid fields: hardware_requirements.min_cpu_cores -> (negative value), hardware_requirements.min_ram -> (negative value), hardware_requirements.min_free_space -> (negative value): Invalid plugin -> invalid fields: plugin.image -> (invalid docker image: invalid reference format)",
		},
		{
			name:      "Minimal Manifest",
			filePath:  "minimal/pkg/manifest.yml",
			wantError: "",
		},
		{
			name:      "Missing Fields Manifest",
			filePath:  "missing-fields/pkg/manifest.yml",
			wantError: "Invalid manifest file -> missing fields: version, name, upgrade, profiles",
		},
		{
			name:      "Missing Fields Manifest in profile",
			filePath:  "missing-fields-profile/pkg/manifest.yml",
			wantError: "Invalid manifest file -> missing fields: version, name, upgrade: invalid profiles: profile 0: profile 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			data, err := afero.ReadFile(afs, filepath.Join(testDir, "manifests", tt.filePath))
			if err != nil {
				t.Fatalf("failed reading data from yaml file: %s", err)
			}

			manifest := Manifest{}
			if err := yaml.Unmarshal(data, &manifest); err != nil {
				t.Fatalf("failed unmarshalling yaml: %s", err)
			}

			err = manifest.validate()
			if tt.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantError)
			}
		})
	}
}

// TODO: Add more test cases
func TestManifest_Validate(t *testing.T) {
	tests := []struct {
		name     string
		manifest *Manifest
		wantErr  bool
	}{
		{
			name: "valid manifest",
			manifest: &Manifest{
				Version: "1.0.0",
				Name:    "test-package",
				Upgrade: "manual",
				HardwareRequirements: hardwareRequirements{
					MinCPUCores:                 2,
					MinRAM:                      4096,
					MinFreeSpace:                1024,
					StopIfRequirementsAreNotMet: true,
				},
				Plugin: &Plugin{
					Image: "test-image:latest",
				},
				Profiles: []string{"test-profile"},
			},
			wantErr: false,
		},
		{
			name: "invalid manifest",
			manifest: &Manifest{
				Version: "1.0.0",
				Name:    "",
				Upgrade: "manual",
				HardwareRequirements: hardwareRequirements{
					MinCPUCores:                 -1,
					MinRAM:                      4096,
					MinFreeSpace:                1024,
					StopIfRequirementsAreNotMet: true,
				},
				Plugin: &Plugin{
					Image: "",
				},
				Profiles: []string{"test-profile"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.manifest.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Manifest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlugin_Validate(t *testing.T) {
	tests := []struct {
		name    string
		plugin  *Plugin
		wantErr bool
	}{
		{
			name:    "valid image",
			plugin:  &Plugin{Image: "nginx:latest"},
			wantErr: false,
		},
		{
			name:    "valid image with tag",
			plugin:  &Plugin{Image: "localhost:5000/my-image"},
			wantErr: false,
		},
		{
			name:    "valid image with complex tag",
			plugin:  &Plugin{Image: "my-registry.com:8080/my-image:1.0"},
			wantErr: false,
		},
		{
			name:    "valid image with sub-registry",
			plugin:  &Plugin{Image: "my-registry.com/sub-registry/my-image:abc123"},
			wantErr: false,
		},
		{
			name:    "invalid image with double colon",
			plugin:  &Plugin{Image: "nginx::latest"},
			wantErr: true,
		},
		{
			name:    "invalid image with double colon and port",
			plugin:  &Plugin{Image: "localhost::5000/my-image"},
			wantErr: true,
		},
		{
			name:    "invalid image with double colon and port",
			plugin:  &Plugin{Image: "my-registry.com::8080/my-image:1.0"},
			wantErr: true,
		},
		{
			name:    "invalid image with double at symbol",
			plugin:  &Plugin{Image: "my-registry.com/my-image@@sha256:abc123"},
			wantErr: true,
		},
		{
			name:    "invalid image with uppercase letters",
			plugin:  &Plugin{Image: "UPPERCASEINVALID"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.plugin.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Plugin.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
