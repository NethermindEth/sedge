package e2e

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/keystores"
	"github.com/stretchr/testify/assert"
)

type depositDataKey struct {
	Account               string `json:"account"`
	Amount                int    `json:"amount"`
	DepositCliVersion     string `json:"deposit_cli_version"`
	DepositDataRoot       string `json:"deposit_data_root"`
	DepositMessageRoot    string `json:"deposit_message_root"`
	ForkVersion           string `json:"fork_version"`
	NetworkName           string `json:"network_name"`
	Pubkey                string `json:"pubkey"`
	Signature             string `json:"signature"`
	Version               int    `json:"version"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
}

func TestKeys_Eth1_Withdrawal_Keys_Mainnet(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		nil,
		// Act
		func(t *testing.T, binaryPath, dataDirPath string) {
			runErr = runSedge(t, binaryPath, "keys",
				"--eth1-withdrawal-address", "0xb794f5ea0ba39494ce839613fffba74279579268",
				"--network", "mainnet",
				"--num-validators", "10",
				"--random-passphrase",
				"--path", dataDirPath)
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.NoError(t, runErr, "keys command should not fail")
			// Check if the deposit_data.json was created
			depositDataFilePath := filepath.Join(dataDirPath, "keystore", keystores.DepositDataFileName)
			assert.FileExists(t, depositDataFilePath, "deposit_data.json should be created")

			// Check if the deposit_data.json is valid
			var keys []depositDataKey
			jsonData, err := os.ReadFile(depositDataFilePath)
			assert.NoError(t, err, "error reading deposit_data.json")
			err = json.Unmarshal([]byte(jsonData), &keys)
			assert.NoError(t, err, "error unmarshalling json")

			pattern := `^010000000000000000000000[a-fA-F0-9]{40}$`
			regex := regexp.MustCompile(pattern)
			for _, key := range keys {
				assert.Regexp(t, regex, key.WithdrawalCredentials, "withdrawal_credentials should match the pattern")
				assert.Equal(t, key.NetworkName, "mainnet", "network_name should be mainnet")
			}
		},
	)
	// Run test case
	e2eTest.run()
}
