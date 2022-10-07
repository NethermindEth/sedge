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
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/keystores"
	"github.com/NethermindEth/sedge/test"
	log "github.com/sirupsen/logrus"
)

type keysCmdTestCase struct {
	name           string
	configPath     string
	network        string
	keystorePath   string
	passphrasePath string
	mnemnonicPath  string
	existingVal    int64
	numVal         int64
	fdOut          *bytes.Buffer
	isErr          bool
}

func resetKeysCmd() {
	cfgFile = ""
	path = ""
	network = ""
	mnemonicPath = ""
	passphrasePath = ""
	eth1WithdrawalAddress = ""
	existingVal = 0
	numberVal = 0
}

func buildKeysTestCase(t *testing.T, caseName, caseDataPath, caseNetwork string, existing, num int64, isErr bool) *keysCmdTestCase {
	tc := keysCmdTestCase{}
	configPath := t.TempDir()

	err := test.PrepareTestCaseDir(filepath.Join("testdata", "keys_tests", caseDataPath, "config"), configPath)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	// Create Mnemonic file
	mnemonicPath := filepath.Join(configPath, "mnemonic.txt")
	file, err := os.Create(mnemonicPath)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}
	defer file.Close()
	testMnemonic, err := keystores.CreateMnemonic()
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}
	file.WriteString(testMnemonic)

	keystorePath := filepath.Join(configPath, "keystore")
	err = os.MkdirAll(keystorePath, os.ModePerm)
	if err != nil {
		t.Fatalf("Can't build test case: %v", err)
	}

	tc.name = caseName
	tc.configPath = filepath.Join(configPath, "config.yaml")
	tc.network = network
	tc.keystorePath = keystorePath
	tc.mnemnonicPath = mnemonicPath
	tc.passphrasePath = filepath.Join(configPath, "pass.txt")
	tc.existingVal = existing
	tc.numVal = num
	tc.fdOut = new(bytes.Buffer)
	tc.isErr = isErr
	return &tc
}

func TestKeysCmd(t *testing.T) {
	// TODO: allow to test error programs
	tcs := []keysCmdTestCase{
		*buildKeysTestCase(t, "Mainnet", "case_1", "mainnet", 0, 1, false),
		*buildKeysTestCase(t, "Unknow network", "case_1", "none", 0, 1, true),
		*buildKeysTestCase(t, "Bigger number", "case_1", "sepolia", 0, 100, false),
		*buildKeysTestCase(t, "Existing validators", "case_1", "sepolia", 100, 10, false),
	}

	t.Cleanup(resetKeysCmd)

	for _, tc := range tcs {
		resetKeysCmd()
		t.Run(
			tc.name,
			func(t *testing.T) {
				rootCmd.SetArgs([]string{"keys",
					"--config", tc.configPath,
					"--network", tc.network,
					"--path", tc.keystorePath,
					"--mnemonic-path", tc.mnemnonicPath,
					"--passphrase-path", tc.passphrasePath,
					"--existing", fmt.Sprint(tc.existingVal),
					"--num-validators", fmt.Sprint(tc.numVal),
				})
				rootCmd.SetOut(tc.fdOut)
				log.SetOutput(tc.fdOut)

				descr := fmt.Sprintf("sedge keys --network %s --existing %d --num-validators %d", tc.network, tc.existingVal, tc.numVal)
				err := rootCmd.Execute()
				if tc.isErr && err == nil {
					t.Errorf("%s expected to fail", descr)
				} else if !tc.isErr && err != nil {
					t.Errorf("%s failed: %v", descr, err)
				}
			},
		)
	}
}
