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

import "path/filepath"

// PathMatcher is a custom matcher for filepath comparison that implements the gomock.Matcher interface
type PathMatcher struct {
	Expected string
}

func (m PathMatcher) Matches(x interface{}) bool {
	path, ok := x.(string)
	if !ok {
		return false
	}
	// Remove drive letter if present
	path = removeDriveLetter(path)
	expected := removeDriveLetter(m.Expected)
	return filepath.Clean(path) == filepath.Clean(expected)
}

func (m PathMatcher) String() string {
	return "is equal to " + m.Expected + " (ignoring drive letter and separators)"
}

// removeDriveLetter removes the drive letter from a Windows path
func removeDriveLetter(path string) string {
	if len(path) >= 2 && path[1] == ':' {
		return path[2:]
	}
	return path
}
