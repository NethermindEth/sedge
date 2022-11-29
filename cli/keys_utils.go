package cli

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/NethermindEth/sedge/cli/prompts"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runKeysWithStakingDeposit(cmd *cobra.Command, args []string, flags *KeysCmdFlags, prompt prompts.Prompt) error {
	existingMnemonic := flags.mnemonicPath != ""
	// Check if dependencies are installed. Keep checking dependencies until they are all installed
	for pending := utils.CheckDependencies([]string{"docker"}); len(pending) > 0; {
		log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
		if flags.install {
			// Install dependencies directly
			if err := installDependencies(pending); err != nil {
				log.Fatal(err)
			}
		} else {
			// Let the user decide to see the instructions for installing dependencies and exit or let the tool install them and continue
			if err := installOrShowInstructions(pending); err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Info(configs.DependenciesOK)

	// Prompt for eth1WithdrawalAddress
	if flags.eth1WithdrawalAddress == "" {
		eth1Address, err := prompt.Eth1Withdrawal()
		if err != nil {
			return err
		}
		flags.eth1WithdrawalAddress = eth1Address
	}

	// Get keystore password
	password := passwordPrompt()

	// Create keystore folder
	log.Info(configs.GeneratingKeystoresLegacy)
	if err := os.MkdirAll(filepath.Join(flags.path, "keystore"), 0o766); err != nil {
		log.Fatal(err)
	}

	keystorePath := filepath.Join(flags.path, "keystore", "validator_keys")
	data := utils.ValidatorKeyData{
		Existing:              existingMnemonic,
		Network:               flags.network,
		Path:                  keystorePath,
		Password:              password,
		Eth1WithdrawalAddress: flags.eth1WithdrawalAddress,
	}
	if err := utils.GenerateValidatorKey(data); err != nil {
		log.Fatalf(configs.GeneratingKeystoreError, err)
	}

	// Check if keystore generation went ok
	if !emptyKeystore(flags.path) {
		log.Infof(configs.KeysFoundAt, keystorePath)
		if err := createKeystorePassword(password, flags.path); err != nil {
			log.Fatalf(configs.CreatingKeystorePasswordError, err)
		}

		log.Warn(configs.ReviewKeystorePath)
	}
	return nil
}

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
