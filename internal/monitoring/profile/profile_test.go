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
package profile

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/NethermindEth/sedge/internal/monitoring/package_handler/testdata"
)

func TestOptionValidate(t *testing.T) {
	afs := afero.NewMemMapFs()
	testDir, err := afero.TempDir(afs, "", "test")
	require.NoError(t, err)
	testdata.SetupDir(t, "options", testDir, afs)
	message := "Option #1 is invalid"

	tests := []struct {
		name     string
		filePath string
		want     error
	}{
		{
			name:     "Full OK Option",
			filePath: "full-ok/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Missing Fields Option",
			filePath: "missing-fields/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"options.help"},
			},
		},
		{
			name:     "Missing and Invalid Fields Option",
			filePath: "missing-invalid-fields/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"options.name", "options.target", "options.help"},
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Full Missing Fields Option",
			filePath: "full-missing/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"options.name", "options.target", "options.type", "options.help"},
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Invalid Type in Option",
			filePath: "invalid-type/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"options.target", "options.help"},
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type int",
			filePath: "check-invalid-int/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type int with min-max value",
			filePath: "check-invalid-int-validate/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check valid type int",
			filePath: "check-valid-int/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type int with validate",
			filePath: "check-valid-int-validate/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type int without max value",
			filePath: "check-valid-int-without-max/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type int without min value",
			filePath: "check-valid-int-without-min/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check invalid type port",
			filePath: "check-invalid-port/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type port with negative value",
			filePath: "check-negative-port/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type port with huge value",
			filePath: "check-huge-port/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type port with zero value",
			filePath: "check-zero-port/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type port with decimal value",
			filePath: "check-decimal-port/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type bool",
			filePath: "check-invalid-bool/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type float",
			filePath: "check-invalid-float/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type float with min-max value",
			filePath: "check-invalid-float-validate/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check valid type float",
			filePath: "check-valid-float/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type float with validate",
			filePath: "check-valid-float-validate/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type float without max value",
			filePath: "check-valid-float-without-max/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type float without min value",
			filePath: "check-valid-float-without-min/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check invalid type uri",
			filePath: "check-invalid-uri/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type uri with scheme",
			filePath: "check-invalid-uri-scheme/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check valid type uri",
			filePath: "check-valid-uri/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type uri with invalid scheme",
			filePath: "check-valid-uri-invalid-scheme/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check type select",
			filePath: "check-type-select/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"options.validate"},
			},
		},
		{
			name:     "Check type select with validate",
			filePath: "check-select-validate/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check type str",
			filePath: "check-type-str/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type path_dir",
			filePath: "check-valid-path_dir/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type str with validate",
			filePath: "check-valid-str-validate/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check invalid type str with validate",
			filePath: "check-invalid-str-validate/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check valid type path_file",
			filePath: "check-valid-path-file/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check valid type path_file with validate",
			filePath: "check-valid-path-file-validate/pkg/option.yml",
			want:     nil,
		},
		{
			name:     "Check invalid type path_file",
			filePath: "check-invalid-path-file/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
		{
			name:     "Check invalid type path_file with validate",
			filePath: "check-invalid-path-file-validate/pkg/option.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"options.default"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			data, err := afero.ReadFile(afs, filepath.Join(testDir, "options", tt.filePath))
			if err != nil {
				t.Fatalf("failed reading data from yaml file: %s", err)
			}

			option := Option{}
			if err := yaml.Unmarshal(data, &option); err != nil {
				t.Fatalf("failed unmarshalling yaml: %s", err)
			}

			got := option.validate(0)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMonitoringTargetValidate(t *testing.T) {
	afs := afero.NewMemMapFs()
	testDir, err := afero.TempDir(afs, "", "test")
	require.NoError(t, err)
	testdata.SetupDir(t, "monitoring-targets", testDir, afs)
	message := "Monitoring target #1 is invalid"

	tests := []struct {
		name     string
		filePath string
		want     error
	}{
		{
			name:     "Full OK Monitoring Target",
			filePath: "ok/pkg/target.yml",
		},
		{
			name:     "Invalid Path Monitoring Target",
			filePath: "invalid-path/pkg/target.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"monitoring.targets.path"},
			},
		},
		{
			name:     "Invalid Port Monitoring Target",
			filePath: "invalid-port/pkg/target.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"monitoring.targets.port"},
			},
		},
		{
			name:     "Invalid Service Monitoring Target",
			filePath: "invalid-service/pkg/target.yml",
			want: InvalidProfileError{
				message:       message,
				invalidFields: []string{"monitoring.targets.service"},
			},
		},
		{
			name:     "Invalid Targets Monitoring Target",
			filePath: "invalid-targets/pkg/target.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"monitoring.targets.service", "monitoring.targets.port", "monitoring.targets.path"},
			},
		},
		{
			name:     "Missing Path Monitoring Target",
			filePath: "missing-path/pkg/target.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"monitoring.targets.path"},
			},
		},
		{
			name:     "Missing Port Monitoring Target",
			filePath: "missing-port/pkg/target.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"monitoring.targets.port"},
			},
		},
		{
			name:     "Missing Service Monitoring Target",
			filePath: "missing-service/pkg/target.yml",
			want: InvalidProfileError{
				message:       message,
				missingFields: []string{"monitoring.targets.service"},
			},
		},
		{
			name:     "Missing Targets Monitoring Target",
			filePath: "missing-targets/pkg/target.yml",
			want:     errors.New("invalid monitoring: there must be at least one monitoring target in the profile file"),
		},
		{
			name:     "Multiple Targets Monitoring Target",
			filePath: "multiple-targets/pkg/target.yml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			data, err := afero.ReadFile(afs, filepath.Join(testDir, "monitoring-targets", tt.filePath))
			if err != nil {
				t.Fatalf("failed reading data from yaml file: %s", err)
			}

			monitoring := Monitoring{}
			if err := yaml.Unmarshal(data, &monitoring); err != nil {
				t.Fatalf("failed unmarshalling yaml: %s", err)
			}

			got := monitoring.validate()
			if tt.want == nil {
				assert.NoError(t, got)
			} else {
				assert.Error(t, got)
				assert.ErrorContains(t, got, tt.want.Error())
			}
		})
	}
}
