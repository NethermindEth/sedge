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
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/test"
	log "github.com/sirupsen/logrus"
)

type keysCmdTestCase struct {
	configPath   string
	runner       commands.CommandRunner
	mnemonic     bool
	network      string
	keystorePath string
	fdOut        *bytes.Buffer
	isErr        bool
}

func resetKeysCmd() {
	cfgFile = ""
	path = ""
	network = ""
	existingMnemonic = false
}

func buildKeysTestCase(t *testing.T, caseName, caseNetwork string, mnemonic, isErr bool) *keysCmdTestCase {
	tc := keysCmdTestCase{}
	configPath := t.TempDir()

	err := test.PrepareTestCaseDir(filepath.Join("testdata", "keys_tests", caseName, "config"), configPath)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	keystorePath := filepath.Join(configPath, "keystore")
	err = os.MkdirAll(keystorePath, os.ModePerm)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	// TODO: allow runner edition
	tc.runner = &test.SimpleCMDRunner{
		SRunCMD: func(c commands.Command) (string, error) {
			return "", nil
		},
		SRunBash: func(bs commands.BashScript) (string, error) {
			return "", nil
		},
	}

	tc.configPath = filepath.Join(configPath, "config.yaml")
	tc.network = network
	tc.keystorePath = keystorePath
	tc.mnemonic = mnemonic
	tc.fdOut = new(bytes.Buffer)
	tc.isErr = isErr
	return &tc
}

func TestKeysCmd(t *testing.T) {
	//TODO: allow to test error programs
	tcs := []keysCmdTestCase{
		*buildKeysTestCase(t, "case_1", "mainnet", false, false),
		*buildKeysTestCase(t, "case_1", "mainnet", true, false),
	}

	t.Cleanup(resetKeysCmd)

	for _, tc := range tcs {
		resetKeysCmd()
		rootCmd.SetArgs([]string{"keys", "--config", tc.configPath, "--path", tc.keystorePath})
		rootCmd.SetOut(tc.fdOut)
		log.SetOutput(tc.fdOut)

		commands.InitRunner(func() commands.CommandRunner {
			return tc.runner
		})

		descr := "sedge keys"
		err := rootCmd.Execute()
		if tc.isErr && err == nil {
			t.Errorf("%s expected to fail", descr)
		} else if !tc.isErr && err != nil {
			t.Errorf("%s failed: %v", descr, err)
		}
	}
}
