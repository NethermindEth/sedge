package cli

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
)

func passwordPrompt() string {
	// notest
	validate := func(input string) error {
		if len(input) < 8 {
			return errors.New(configs.KeystorePasswordError)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter the password you will use for the validator keystore",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	if err != nil {
		log.Errorf(configs.PromptFailedError, err)
		return ""
	}

	validate = func(input string) error {
		if input != result {
			return errors.New(configs.KeystorePasswordRetryError)
		}
		return nil
	}

	prompt = promptui.Prompt{
		Label:    "Please re-enter the password. Press Ctrl+C to retry",
		Validate: validate,
		Mask:     '*',
	}

	_, err = prompt.Run()

	if err != nil {
		log.Errorf(configs.PromptFailedError, err)
		return passwordPrompt()
	}

	return result
}

func createKeystorePassword(password string, path string) error {
	log.Debug(configs.CreatingKeystorePassword)

	// Create file keystore_password.txt
	filename := filepath.Join(path, "keystore", "keystore_password.txt")
	_, err := commands.Runner.RunCMD(commands.Runner.BuildCreateFileCMD(commands.CreateFileOptions{
		FileName: filename,
	}))
	if err != nil {
		return err
	}

	// Write password to file
	_, err = commands.Runner.RunCMD(commands.Runner.BuildEchoToFileCMD(commands.EchoToFileOptions{
		FileName: filename,
		Content:  password,
	}))
	if err != nil {
		return err
	}

	log.Info(configs.KeystorePasswordCreated)
	return nil
}

// Check if keystore folder is not empty
func emptyKeystore(path string) bool {
	f, err := os.Open(filepath.Join(path, "keystore"))
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	// Either not empty or error, suits both cases
	return err == io.EOF
}
