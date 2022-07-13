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
package cli

import (
	"bytes"
	"regexp"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestVersionCmdExecute(t *testing.T) {
	fdOut := new(bytes.Buffer)

	versionRegexp := regexp.MustCompile(`sedge [0-9][0-9|\.]+[0-9]`)

	rootCmd.SetArgs([]string{"version"})
	rootCmd.SetOut(fdOut)
	log.SetOutput(fdOut)

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("versionCmd.Execute() failed: %v", err)
	}

	if !versionRegexp.Match(fdOut.Bytes()) {
		t.Errorf("versionCmd.Execute() has an invalid output: %s", fdOut.String())
	}
}
