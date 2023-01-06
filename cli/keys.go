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
	"fmt"
	"os"
	"path/filepath"

	"github.com/NethermindEth/sedge/cli/prompts"
	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/keystores"
	eth2 "github.com/protolambda/zrnt/eth2/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type KeysCmdFlags struct {
	network               string
	path                  string
	eth1WithdrawalAddress string
	mnemonicPath          string
	passphrasePath        string
	existingVal           int64
	numberVal             int64
	randomPassphrase      bool
	install               bool
}

func KeysCmd(cmdRunner commands.CommandRunner, prompt prompts.Prompt) *cobra.Command {
	var (
		flags      KeysCmdFlags
		passphrase string
		mnemonic   string
	)
	// Cmd declaration
	cmd := &cobra.Command{
		Use:   "keys [flags]",
		Short: "Generate keystore folder",
		Long:  "Generate keystore folder using the eth2.0-deposit-cli tool",
		PreRun: func(cmd *cobra.Command, args []string) {
			// Validate network
			if err := configs.NetworkCheck(flags.network); err != nil {
				log.Fatal(err.Error())
			}
			// Ensure that path is absolute
			log.Debugf("Path to keystore file: %s", flags.path)
			absPath, err := filepath.Abs(flags.path)
			if err != nil {
				log.Fatalf(configs.InvalidVolumePathError, err)
			}
			flags.path = absPath
		},
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: allow usage of withdrawal address
			// Get keystore passphrase
			if !flags.randomPassphrase && flags.passphrasePath != "" {
				content, err := readFileContent(flags.passphrasePath)
				if err != nil {
					log.Fatalf(configs.PassphraseReadFileError, err)
				}
				if len(content) < 8 {
					log.Warn(configs.KeystorePasswordError)
				} else {
					passphrase = content
				}
			}
			if !flags.randomPassphrase && len(passphrase) < 8 {
				passphrase = prompt.Passphrase()
			}

			// Get or generate mnemonic
			if flags.mnemonicPath != "" {
				content, err := readFileContent(flags.mnemonicPath)
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
				if err := saveMnemonic(cmdRunner, mnemonic); err != nil {
					log.Fatal(err)
				}
			}

			// Get indexes
			if flags.mnemonicPath == "" {
				flags.existingVal = 0
			} else if flags.existingVal < 0 {
				flags.existingVal = prompt.ExistingVal()
			}

			if flags.numberVal <= 0 {
				flags.numberVal = prompt.NumberVal()
			}

			keystorePath := filepath.Join(flags.path, "keystore")

			data := keystores.ValidatorKeysGenData{
				Mnemonic:    mnemonic,
				Passphrase:  passphrase,
				OutputPath:  keystorePath,
				MinIndex:    uint64(flags.existingVal),
				MaxIndex:    uint64(flags.existingVal) + uint64(flags.numberVal),
				NetworkName: flags.network,
				ForkVersion: configs.NetworksConfigs()[flags.network].GenesisForkVersion,
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
	// Flag binds
	cmd.Flags().StringVarP(&flags.network, "network", "n", "mainnet", "Target network. e.g. mainnet, goerli, sepolia etc.")
	cmd.Flags().StringVarP(&flags.path, "path", "p", configs.DefaultDockerComposeScriptsPath, "Absolute path to keystore folder. e.g. /home/user/keystore")
	cmd.Flags().StringVar(&flags.eth1WithdrawalAddress, "eth1-withdrawal-address", "", "If this field is set and valid, the given Eth1 address will be used to create the withdrawal credentials. Otherwise, it will generate withdrawal credentials with the mnemonic-derived withdrawal public key in EIP-2334 format.")
	cmd.Flags().StringVar(&flags.mnemonicPath, "mnemonic-path", "", "Path to file with a existing mnemonic to use.")
	cmd.Flags().StringVar(&flags.passphrasePath, "passphrase-path", "", "Path to file with a keystores passphrase to use.")
	cmd.Flags().Int64Var(&flags.existingVal, "existing", -1, `Number of validators generated with the provided mnemonic. Will be ignored if "--mnemonic-path" its not set. This number will be used as the initial index for the generated keystores.`)
	cmd.Flags().Int64Var(&flags.numberVal, "num-validators", -1, "Number of validators to generate. This number will be used in addition to the existing flag as the end index for the generated keystores.")
	cmd.Flags().BoolVar(&flags.randomPassphrase, "random-passphrase", false, "Usa a randomly generated passphrase to encrypt keystores.")
	cmd.Flags().BoolVarP(&flags.install, "install", "i", false, "Install dependencies if not installed without asking")
	return cmd
}

func saveMnemonic(cmdRunner commands.CommandRunner, mnemonic string) error {
	file, err := os.CreateTemp(os.TempDir(), "sedge_mnemonic")
	if err != nil {
		return fmt.Errorf(configs.ShowMnemonicError, err)
	}
	defer os.Remove(file.Name())

	if _, err := file.WriteString(fmt.Sprintf(configs.MnemonicPresentation, mnemonic)); err != nil {
		return fmt.Errorf(configs.ShowMnemonicError, err)
	}

	if err := file.Sync(); err != nil {
		return fmt.Errorf(configs.ShowMnemonicError, err)
	}

	openTextEditorCmd := cmdRunner.BuildOpenTextEditor(commands.OpenTextEditorOptions{
		FilePath: file.Name(),
	})
	openTextEditorCmd.ForceNoSudo = true

	_, err = cmdRunner.RunCMD(openTextEditorCmd)
	if err != nil {
		return fmt.Errorf(configs.ShowMnemonicError, err)
	}

	return nil
}
