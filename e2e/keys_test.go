package e2e

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/NethermindEth/sedge/internal/lido/contracts"
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
		func(t *testing.T, binaryPath string) error {
			mnemonicPathFile := filepath.Join(filepath.Dir(binaryPath), "mnemonic.txt")
			file, err := os.Create(mnemonicPathFile)
			if err != nil {
				return err
			}
			defer file.Close()
			mnemonicText := "science ill robust clump oxygen intact barely horror athlete eyebrow cave target hero input entry citizen wink affair entire alert sick flight gossip refuse"
			_, err = file.WriteString(mnemonicText)
			return err
		},
		// Act
		func(t *testing.T, binaryPath, dataDirPath string) {
			mnemonicPathFile := filepath.Join(filepath.Dir(binaryPath), "mnemonic.txt")
			runErr = runSedge(t, binaryPath, "keys",
				"--eth1-withdrawal-address", "0xb794f5ea0ba39494ce839613fffba74279579268",
				"--network", "mainnet",
				"--num-validators", "10",
				"--mnemonic-path", mnemonicPathFile,
				"--existing", "0",
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

func TestKeys_Lido_Mainnet(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		func(t *testing.T, binaryPath string) error {
			mnemonicPathFile := filepath.Join(filepath.Dir(binaryPath), "mnemonic.txt")
			file, err := os.Create(mnemonicPathFile)
			if err != nil {
				return err
			}
			defer file.Close()
			mnemonicText := "science ill robust clump oxygen intact barely horror athlete eyebrow cave target hero input entry citizen wink affair entire alert sick flight gossip refuse"
			_, err = file.WriteString(mnemonicText)
			return err
		},
		// Act
		func(t *testing.T, binaryPath, dataDirPath string) {
			mnemonicPathFile := filepath.Join(filepath.Dir(binaryPath), "mnemonic.txt")
			runErr = runSedge(t, binaryPath, "keys",
				"--lido",
				"--network", "mainnet",
				"--num-validators", "10",
				"--mnemonic-path", mnemonicPathFile,
				"--existing", "0",
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
				expectedWithdrawalAddress := "010000000000000000000000" + (contracts.WithdrawalAddress["mainnet"].WithdrawalAddress[2:])
				assert.Equal(t, expectedWithdrawalAddress, key.WithdrawalCredentials, "WithdrawalAddress value should match expected value")
			}
		},
	)
	// Run test case
	e2eTest.run()
}

func TestKeys_Lido_Eth1Withdrawal_HoleskyInvalid(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		func(t *testing.T, binaryPath string) error {
			mnemonicPathFile := filepath.Join(filepath.Dir(binaryPath), "mnemonic.txt")
			file, err := os.Create(mnemonicPathFile)
			if err != nil {
				return err
			}
			defer file.Close()
			mnemonicText := "science ill robust clump oxygen intact barely horror athlete eyebrow cave target hero input entry citizen wink affair entire alert sick flight gossip refuse"
			_, err = file.WriteString(mnemonicText)
			return err
		},
		// Act
		func(t *testing.T, binaryPath, dataDirPath string) {
			mnemonicPathFile := filepath.Join(filepath.Dir(binaryPath), "mnemonic.txt")
			runErr = runSedge(t, binaryPath, "keys",
				"--lido",
				"--eth1-withdrawal-address", "0xb794f5ea0ba39494ce839613fffba74279579268",
				"--network", "holesky",
				"--num-validators", "10",
				"--mnemonic-path", mnemonicPathFile,
				"--existing", "0",
				"--random-passphrase",
				"--path", dataDirPath)
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "keys command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}

func TestKeys_Lido_GnosisUnsupported(t *testing.T) {
	// Test context
	var (
		runErr error
	)
	// Build test case
	e2eTest := newE2ETestCase(
		t,
		// Arrange
		func(t *testing.T, binaryPath string) error {
			mnemonicPathFile := filepath.Join(filepath.Dir(binaryPath), "mnemonic.txt")
			file, err := os.Create(mnemonicPathFile)
			if err != nil {
				return err
			}
			defer file.Close()
			mnemonicText := "science ill robust clump oxygen intact barely horror athlete eyebrow cave target hero input entry citizen wink affair entire alert sick flight gossip refuse"
			_, err = file.WriteString(mnemonicText)
			return err
		},
		// Act
		func(t *testing.T, binaryPath, dataDirPath string) {
			mnemonicPathFile := filepath.Join(filepath.Dir(binaryPath), "mnemonic.txt")
			runErr = runSedge(t, binaryPath, "keys",
				"--lido",
				"--eth1-withdrawal-address", "0xb794f5ea0ba39494ce839613fffba74279579268",
				"--network", "gnosis",
				"--num-validators", "10",
				"--mnemonic-path", mnemonicPathFile,
				"--existing", "0",
				"--random-passphrase",
				"--path", dataDirPath)
		},
		// Assert
		func(t *testing.T, dataDirPath string) {
			assert.Error(t, runErr, "keys command should fail")
		},
	)
	// Run test case
	e2eTest.run()
}
