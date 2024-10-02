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

	"github.com/stretchr/testify/assert"
)

func TestPathMatcher(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		input    interface{}
		want     bool
	}{
		{
			name:     "matching paths with different separators",
			expected: "/path/to/file",
			input:    "\\path\\to\\file",
			want:     true,
		},
		{
			name:     "matching paths with drive letter",
			expected: "/path/to/file",
			input:    "C:\\path\\to\\file",
			want:     true,
		},
		{
			name:     "non-matching paths",
			expected: "/path/to/file",
			input:    "/path/to/other",
			want:     false,
		},
		{
			name:     "input is not a string",
			expected: "/path/to/file",
			input:    123,
			want:     false,
		},
		{
			name:     "both paths with different drive letters",
			expected: "C:/path/to/file",
			input:    "D:\\path\\to\\file",
			want:     true,
		},
		{
			name:     "expected path with drive letter",
			expected: "C:/path/to/file",
			input:    "/path/to/file",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matcher := PathMatcher{Expected: tt.expected}
			assert.Equal(t, tt.want, matcher.Matches(tt.input))
		})
	}
}

func TestPathMatcherString(t *testing.T) {
	matcher := PathMatcher{Expected: "/path/to/file"}
	expected := "is equal to /path/to/file (ignoring drive letter and separators)"
	assert.Equal(t, expected, matcher.String())
}

func TestRemoveDriveLetter(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "path with drive letter",
			input: "C:\\path\\to\\file",
			want:  "\\path\\to\\file",
		},
		{
			name:  "path without drive letter",
			input: "/path/to/file",
			want:  "/path/to/file",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "single character",
			input: "a",
			want:  "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeDriveLetter(tt.input))
		})
	}
}
