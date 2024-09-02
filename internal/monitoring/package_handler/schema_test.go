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

import "testing"

func Test_validateYAMLSchema(t *testing.T) {
	type args struct {
		schemaFile   string
		documentFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid manifest",
			args: args{
				schemaFile:   "schema/manifest_schema.yml",
				documentFile: "testdata/manifests/valid-schema/manifest.yml",
			},
			wantErr: false,
		},
		{
			name: "invalid manifest",
			args: args{
				schemaFile:   "schema/manifest_schema.yml",
				documentFile: "testdata/manifests/invalid-schema/manifest.yml",
			},
			wantErr: true,
		},
		{
			name: "valid profile",
			args: args{
				schemaFile:   "schema/profile_schema.yml",
				documentFile: "testdata/profiles/valid-schema/profile.yml",
			},
			wantErr: false,
		},
		{
			name: "invalid profile",
			args: args{
				schemaFile:   "schema/profile_schema.yml",
				documentFile: "testdata/profiles/invalid-schema/profile.yml",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateYAMLSchema(tt.args.schemaFile, tt.args.documentFile); (err != nil) != tt.wantErr {
				t.Errorf("validateYAMLSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
