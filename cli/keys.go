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
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/keystores"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/manifoldco/promptui"
	eth2 "github.com/protolambda/zrnt/eth2/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	path                  string
	mnemonicPath          string
	passphrasePath        string
	eth1WithdrawalAddress string
	existingVal           int64
	numberVal             int64
	randomPassphrase      bool
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
		// Validate network
		network = strings.ToLower(network)
		_, ok := configs.NetworksConfigs[network]
		if !ok {
			log.Fatalf(configs.UnknownNetworkError, network)
		}

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
		// Handle mainnet case
		if network == "mainnet" {
			runKeysWithStakingDeposit(cmd, args)
			return
		}

		// TODO: allow usage of withdrawal address
		// Get keystore passphrase
		var passphrase string
		if !randomPassphrase && passphrasePath != "" {
			content, err := readFileContent(passphrasePath)
			if err != nil {
				log.Fatalf(configs.PassphraseReadFileError, err)
			}
			if len(content) < 8 {
				log.Warn(configs.KeystorePasswordError)
			} else {
				passphrase = content
			}
		}
		if !randomPassphrase && passphrase == "" || len(passphrase) < 8 {
			passphrase = passphrasePrompt()
		}

		// Get or generate mnemonic
		var mnemonic string
		if mnemonicPath != "" {
			content, err := readFileContent(mnemonicPath)
			if err != nil {
				log.Fatalf(configs.MnemonicReadFileError, err)
			}
			mnemonic = content
		}
		if mnemonic == "" {
			log.Warn(configs.GeneratingMnemonic)
			candidate, err := keystores.CreateMnemonic()
			if err != nil {
				log.Fatal(err)
			}
			mnemonic = candidate
			// TODO: improve prompts for the generated mnemonic. This should confirm user have copied the mnemonic by asking to input it again.
			// TODO: clean screen after the generated mnemonic is printed.
			fmt.Fprintf(cmd.OutOrStdout(), "Mnemonic:\n\n%s\n\n", mnemonic)
			prompt := promptui.Prompt{
				Label: configs.StoreMnemonic,
			}
			prompt.Run()
		}

		// Get indexes
		if mnemonicPath == "" {
			existingVal = 0
		} else if existingVal < 0 {
			existingVal = existingValPrompt()
		}

		if numberVal <= 0 {
			numberVal = numberValPrompt()
		}

		keystorePath := filepath.Join(path, "keystore")

		data := keystores.ValidatorKeysGenData{
			Mnemonic:    mnemonic,
			Passphrase:  passphrase,
			OutputPath:  keystorePath,
			MinIndex:    uint64(existingVal),
			MaxIndex:    uint64(existingVal) + uint64(numberVal),
			NetworkName: network,
			ForkVersion: configs.NetworksConfigs[network].GenesisForkVersion,
			// Constants
			UseUniquePassphrase: true,
			Insecure:            false,
			AmountGwei:          uint64(eth2.Mainnet.MAX_EFFECTIVE_BALANCE),
			AsJsonList:          true,
		}

		log.Info(configs.GeneratingKeystores)
		if err := keystores.CreateKeystores(data); err != nil {
			log.Fatal(err)
		}
		log.Info(configs.KeystoresGenerated)
		log.Info(configs.GeneratingDepositData)
		if err := keystores.CreateDepositData(data); err != nil {
			log.Fatal(err)
		}
		log.Info(configs.DepositDataGenerated)
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)

	// Local flags
	keysCmd.Flags().StringVarP(&network, "network", "n", "mainnet", "Target network. e.g. mainnet, goerli, sepolia etc.")

	keysCmd.Flags().StringVarP(&path, "path", "p", configs.DefaultSedgeDataPath, "Absolute path to keystore folder. e.g. /home/user/keystore")

	keysCmd.Flags().StringVar(&eth1WithdrawalAddress, "eth1-withdrawal-address", "", "If this field is set and valid, the given Eth1 address will be used to create the withdrawal credentials. Otherwise, it will generate withdrawal credentials with the mnemonic-derived withdrawal public key in EIP-2334 format.")

	keysCmd.Flags().StringVar(&mnemonicPath, "mnemonic-path", "", "Path to file with a existing mnemonic to use.")

	keysCmd.Flags().StringVar(&passphrasePath, "passphrase-path", "", "Path to file with a keystores passphrase to use.")

	keysCmd.Flags().Int64Var(&existingVal, "existing", -1, `Number of validators generated with the provided mnemonic. Will be ignored if "--mnemonic-path" its not set. This number will be used as the initial index for the generated keystores.`)

	keysCmd.Flags().Int64Var(&numberVal, "num-validators", -1, "Number of validators to generate. This number will be used in addition to the existing flag as the end index for the generated keystores.")

	keysCmd.Flags().BoolVar(&randomPassphrase, "random-passphrase", false, "Usa a randomly generated passphrase to encrypt keystores.")
}

func passphrasePrompt() string {
	// notest
	validate := func(input string) error {
		if len(input) < 8 {
			return errors.New(configs.KeystorePasswordError)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter the passphrase you will use for the validator keystore",
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
		Label:    "Please re-enter the passphrase. Press Ctrl+C to retry",
		Validate: validate,
		Mask:     '*',
	}

	_, err = prompt.Run()

	if err != nil {
		log.Errorf(configs.PromptFailedError, err)
		return passphrasePrompt()
	}

	return result
}

func existingValPrompt() int64 {
	// notest
	validate := func(input string) error {
		if value, err := strconv.ParseInt(input, 10, 64); err != nil || value < 0 {
			if value < 0 {
				err = fmt.Errorf("value must be positive")
			}
			return fmt.Errorf(configs.InvalidNumberOfValidatorsError, err)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter the number of previous validators keystores generated with this mnemonic. This number will be used as the initial index for the generated keystores",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf(configs.PromptFailedError, err)
	}

	validate = func(input string) error {
		if input != result {
			return errors.New(configs.NumberOfValidatorsRetryError)
		}
		return nil
	}

	prompt = promptui.Prompt{
		Label:    "Please confirm the number of previous validators keystores generated. Press Ctrl+C to retry",
		Validate: validate,
	}

	_, err = prompt.Run()

	if err != nil {
		log.Errorf(configs.PromptFailedError, err)
		return existingValPrompt()
	}

	index, _ := strconv.ParseInt(result, 10, 64)
	return index
}

func numberValPrompt() int64 {
	// notest
	validate := func(input string) error {
		if value, err := strconv.ParseInt(input, 10, 64); err != nil || value <= 0 {
			if value <= 0 {
				err = fmt.Errorf("value must be greater than 0")
			}
			return fmt.Errorf(configs.InvalidNumberOfValidatorsError, err)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter the number of validators keystores to generate with this mnemonic. This number will be used in addition to the existing validators as the end index for the generated keystores",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf(configs.PromptFailedError, err)
	}

	index, _ := strconv.ParseInt(result, 10, 64)
	return index
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

func readFileContent(path string) (string, error) {
	raw, err := os.ReadFile(path)
	content := strings.TrimSpace(strings.TrimSuffix(string(raw), "\n"))
	return content, err
}
