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
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	path                  string
	eth1WithdrawalAddress string
	existingMnemonic      bool
)

// Windows and Unix path
var rePath = regexp.MustCompile(`^[a-zA-Z]+:(\\[a-zA-Z0-9_.-]+)+|^~{0,1}\/[a-zA-Z0-9~]+(\/[a-zA-Z0-9_.-]+)+`)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys [flags]",
	Short: "Generate keystore folder",
	Long: `Generate keystore folder using the eth2.0-deposit-cli tool.
	
New mnemonic will be generated if -e/--existing flag is not provided.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		var err error
		// Ensure that path is absolute
		log.Debugf("Path to keystore file: %s", path)
		if !filepath.IsAbs(path) {
			path, err = filepath.Abs(path)
			if err != nil {
				log.Fatalf(configs.InvalidVolumePathError, err)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Validate network when several networks are supported

		// Check if dependencies are installed. Keep checking dependencies until they are all installed
		for pending := utils.CheckDependencies([]string{"docker"}); len(pending) > 0; {
			log.Infof(configs.DependenciesPending, strings.Join(pending, ", "))
			if install {
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
		if eth1WithdrawalAddress == "" {
			eth1WithdrawalPrompt()
		}

		// Get keystore password
		password := passwordPrompt()

		// Create keystore folder
		log.Info(configs.GeneratingKeystore)
		if err := os.MkdirAll(filepath.Join(path, "keystore"), 0o766); err != nil {
			log.Fatal(err)
		}

		keystorePath := filepath.Join(path, "keystore", "validator_keys")
		data := utils.ValidatorKeyData{
			Existing:              existingMnemonic,
			Network:               network,
			Path:                  keystorePath,
			Password:              password,
			Eth1WithdrawalAddress: eth1WithdrawalAddress,
		}
		if err := utils.GenerateValidatorKey(data); err != nil {
			log.Fatalf(configs.GeneratingKeystoreError, err)
		}

		// Check if keystore generation went ok
		if !emptyKeystore() {
			log.Infof(configs.KeysFoundAt, keystorePath)
			if err := createKeystorePassword(password); err != nil {
				log.Fatalf(configs.CreatingKeystorePasswordError, err)
			}

			log.Warn(configs.ReviewKeystorePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)

	// Local flags
	keysCmd.Flags().StringVarP(&network, "network", "n", "mainnet", "Target network. e.g. mainnet, goerli, sepolia etc.")

	keysCmd.Flags().StringVarP(&path, "path", "p", configs.DefaultDockerComposeScriptsPath, "Absolute path to keystore folder. e.g. /home/user/keystore")

	keysCmd.Flags().StringVar(&eth1WithdrawalAddress, "eth1-withdrawal-address", "", "If this field is set and valid, the given Eth1 address will be used to create the withdrawal credentials. Otherwise, it will generate withdrawal credentials with the mnemonic-derived withdrawal public key in EIP-2334 format.")

	keysCmd.Flags().BoolVarP(&existingMnemonic, "existing", "e", false, "Use existing mnemonic")
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

func createKeystorePassword(password string) error {
	log.Debug(configs.CreatingKeystorePassword)

	// Create file keystore_password.txt
	file, err := os.Create(filepath.Join(path, "keystore", "keystore_password.txt"))
	if err != nil {
		return err
	}
	defer file.Close()

	// Write password to file
	_, err = file.WriteString(password)
	if err != nil {
		return err
	}

	log.Info(configs.KeystorePasswordCreated)
	return nil
}

// Check if keystore folder is not empty
func emptyKeystore() bool {
	f, err := os.Open(filepath.Join(path, "keystore"))
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	// Either not empty or error, suits both cases
	return err == io.EOF
}

func eth1WithdrawalPrompt() error {
	// notest
	validate := func(input string) error {
		if input != "" && !utils.IsAddress(input) {
			return errors.New("invalid ETH1 address")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter a Eth1 address to be used to create the withdrawal credentials. You can leave it blank and press enter",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf(configs.PromptFailedError, err)
	}

	eth1WithdrawalAddress = result
	return nil
}
