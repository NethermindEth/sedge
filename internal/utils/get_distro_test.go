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
	"runtime"
	"testing"
)

func TestGetDistro(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Test running in a non-linux environment. GOOS: %s", runtime.GOOS)
	}
	_, err := getOSInfo()
	if err != nil {
		t.Errorf("getOsInfo() failed: %v", err)
	}

	// TODO: validate distro info
}

func TestGetDistroName(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Test running in a non-linux environment. GOOS: %s", runtime.GOOS)
	}
	_, err := GetDistroName()
	if err != nil {
		t.Errorf("GetDistroName() failed: %v", err)
	}

	// TODO: validate distro name
}
