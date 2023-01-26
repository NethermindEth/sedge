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
	"testing"

	log "github.com/sirupsen/logrus"
)

type networksTestCase struct {
	name     string
	logsOut  *bytes.Buffer
	tableOut *bytes.Buffer
	isErr    bool
}

func buildNetworksTestCase(t *testing.T, name string, isErr bool) listClientsTestCase {
	tc := listClientsTestCase{}

	tc.name = name
	tc.logsOut = new(bytes.Buffer)
	tc.tableOut = new(bytes.Buffer)
	tc.isErr = isErr
	return tc
}

func TestNetworksCmd(t *testing.T) {
	tcs := [...]listClientsTestCase{
		buildListClientTestCase(t, "Ok", false),
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			rootCmd := RootCmd()
			rootCmd.SetOut(tc.tableOut)
			rootCmd.AddCommand(NetworksCmd())
			rootCmd.SetArgs([]string{"networks"})
			initLogging()
			log.SetOutput(tc.logsOut)

			err := rootCmd.Execute()
			if tc.isErr && err == nil {
				t.Error("sedge networks expected to fail")
			} else if !tc.isErr && err != nil {
				t.Errorf("sedge networks failed: %v", err)
			}
		})
	}
}
