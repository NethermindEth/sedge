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
	"errors"
	"testing"
)

func TestCheckErr(t *testing.T) {
	testErr := errors.New("test error")

	tcs := []struct {
		name    string
		isErr   bool
		err     error
		wantErr bool
	}{
		{
			name:    "Test case 1, got error and don't expect error",
			isErr:   false,
			err:     testErr,
			wantErr: true,
		},
		{
			name:    "Test case 2, got error and expect error",
			isErr:   true,
			err:     testErr,
			wantErr: false,
		},
		{
			name:    "Test case 3, got no error and don't expect error",
			isErr:   false,
			err:     nil,
			wantErr: false,
		},
		{
			name:    "Test case 4, got no error and expect error",
			isErr:   true,
			err:     nil,
			wantErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckErr("", tc.isErr, tc.err)
			if tc.wantErr && err == nil {
				t.Errorf("CheckErr(\"\", %v, %v) failed: expected error, got nil", tc.isErr, tc.err)
			}

			if !tc.wantErr && err != nil {
				t.Errorf("CheckErr(\"\", %v, %v) failed: expected nil, got error", tc.isErr, tc.err)
			}
		})
	}
}
