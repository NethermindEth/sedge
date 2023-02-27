package keystore_test_data

import (
	"embed"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/internal/pkg/keystores"
)

//go:embed valid
var testdata embed.FS

func SetupTestDataDir(t *testing.T) (string, error) {
	t.Helper()
	dir := t.TempDir()
	depositDataContent, err := testdata.ReadFile(path.Join("valid", keystores.DepositDataFileName))
	if err != nil {
		return "", err
	}
	if f, err := os.Create(filepath.Join(dir, keystores.DepositDataFileName)); err != nil {
		return "", err
	} else {
		if _, err := f.Write(depositDataContent); err != nil {
			return "", err
		}
		f.Close()
	}
	keystorePasswordContent, err := testdata.ReadFile(path.Join("valid", keystores.KeystorePasswordFileName))
	if err != nil {
		return "", err
	}
	if f, err := os.Create(filepath.Join(dir, keystores.KeystorePasswordFileName)); err != nil {
		return "", err
	} else {
		if _, err := f.Write(keystorePasswordContent); err != nil {
			return "", err
		}
		f.Close()
	}
	validatorKeysDir, err := testdata.ReadDir(path.Join("valid", keystores.ValidatorKeysDirName))
	if err != nil {
		return "", err
	}
	os.MkdirAll(filepath.Join(dir, keystores.ValidatorKeysDirName), 0o755)
	for _, source := range validatorKeysDir {
		if dest, err := os.Create(filepath.Join(dir, keystores.ValidatorKeysDirName, source.Name())); err != nil {
			return "", err
		} else {
			sourceContent, err := testdata.ReadFile(path.Join("valid", keystores.ValidatorKeysDirName, source.Name()))
			if err != nil {
				return "", err
			}
			if _, err := dest.Write(sourceContent); err != nil {
				return "", err
			}
			dest.Close()
		}
	}
	return dir, nil
}
