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

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/keystores"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/NethermindEth/sedge/internal/utils"
	eth2 "github.com/protolambda/zrnt/eth2/configs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type KeysCmdFlags struct {
	network              string
	path                 string
	ethWithdrawalAddress string
	mnemonicPath         string
	passphrasePath       string
	existingVal          int64
	numberVal            int64
	randomPassphrase     bool
	install              bool
	lidoNode             bool
}

func KeysCmd(cmdRunner commands.CommandRunner, p ui.Prompter) *cobra.Command {
	var (
		flags        KeysCmdFlags
		passphrase   string
		mnemonic     string
		keystorePath string
	)
	// Cmd declaration
	cmd := &cobra.Command{
		Use:   "keys [flags]",
		Short: "Generate keystore folder",
		Long:  "Generate keystore folder using the eth2.0-deposit-cli tool",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Build keystores path
			keystorePath = filepath.Join(flags.path, "keystore")
			keystoreAbsPath, err := filepath.Abs(keystorePath)
			if err != nil {
				return err
			}
			keystorePath = keystoreAbsPath
			// Check if file exists
			if f, err := os.Stat(keystorePath); err == nil {
				if f.IsDir() {
					overwrite, err := p.Confirm(fmt.Sprintf("%s already exists. Do you want to overwrite it?", keystorePath), false)
					if err != nil {
						return err
					}
					if overwrite {
						if err := os.RemoveAll(keystorePath); err != nil {
							return err
						}
					} else {
						return fmt.Errorf("%s already exists", keystorePath)
					}
				} else {
					return fmt.Errorf("%s is not a directory", keystorePath)
				}
			}
			// Validate network
			if err := configs.NetworkCheck(flags.network); err != nil {
				log.Fatal(err.Error())
			}
			// Validate fee recipient
			if flags.ethWithdrawalAddress != "" && !utils.IsAddress(flags.ethWithdrawalAddress) {
				log.Fatal(configs.ErrInvalidWithdrawalAddr)
			}
			// Ensure that path is absolute
			log.Debugf("Path to keystore folder: %s", flags.path)
			absPath, err := filepath.Abs(flags.path)
			if err != nil {
				log.Fatalf(configs.InvalidVolumePathError, err)
			}
			flags.path = absPath
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Incompatible --lido and --eth1-withdrawal-address together
			if flags.lidoNode && flags.ethWithdrawalAddress != "" {
				log.Fatalf(configs.IncompatibleLidoAndEth1Withdrawal)
			}
			// validate network for Lido
			if flags.lidoNode {
				supported := contracts.NetworkSupportedByLidoWithdrawal(flags.network)
				if !supported {
					log.Fatalf(configs.InvalidNetworkForLidoKeys, contracts.LidoWithdrawalSupportedNetworks())
				}
			}
			// Warn about withdrawal address
			if flags.ethWithdrawalAddress != "" {
				log.Warn(configs.WithdrawalAddressDefinedWarning)
			}
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
				input, err := p.InputSecret("Enter keystore passphrase (min 8 characters):")
				if err != nil {
					log.Fatal(err)
				}
				passphrase = input
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
				existingVal, err := p.InputInt64("Enter the number of existing validators (0 if none):", 0)
				if err != nil {
					log.Fatal(err)
				}
				flags.existingVal = existingVal
			}

			if flags.numberVal <= 0 {
				numberVal, err := p.InputInt64("Enter the number of validators to generate:", 1)
				if err != nil {
					log.Fatal(err)
				}
				flags.numberVal = numberVal
			}

			keystorePath := filepath.Join(flags.path, "keystore")

			var withdrawalAddress string
			if flags.lidoNode {
				withdrawalAddress, _ = contracts.WithdrawalAddress(flags.network)
				withdrawalAddress = withdrawalAddress[2:]
			} else if flags.ethWithdrawalAddress != "" {
				withdrawalAddress = flags.ethWithdrawalAddress[2:]
			}

			data := keystores.ValidatorKeysGenData{
				Mnemonic:          mnemonic,
				Passphrase:        passphrase,
				OutputPath:        keystorePath,
				MinIndex:          uint64(flags.existingVal),
				MaxIndex:          uint64(flags.existingVal) + uint64(flags.numberVal),
				NetworkName:       flags.network,
				ForkVersion:       configs.NetworksConfigs()[flags.network].GenesisForkVersion,
				WithdrawalAddress: withdrawalAddress,
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
			log.Infof(configs.KeystorePath, keystorePath)
		},
	}
	// Flag binds
	cmd.PersistentFlags().BoolVar(&flags.lidoNode, "lido", false, "Enable Lido CSM compatible keys. Similar to using --eth-withdrawal-address with the Lido Withdrawal Vault address.")
	cmd.Flags().StringVarP(&flags.network, "network", "n", "mainnet", "Target network. e.g. mainnet,sepolia, holesky, gnosis, chiado etc.")
	cmd.Flags().StringVarP(&flags.path, "path", "p", configs.DefaultAbsSedgeDataPath, "Absolute path to keystore folder. e.g. /home/user/keystore")
	cmd.Flags().StringVar(&flags.ethWithdrawalAddress, "eth-withdrawal-address", "", "If this field is set and valid, the given Eth address will be used to create the withdrawal credentials. Otherwise, it will generate withdrawal credentials with the mnemonic-derived withdrawal public key in EIP-2334 format.")
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

	_, _, err = cmdRunner.RunCMD(openTextEditorCmd)
	if err != nil {
		return fmt.Errorf(configs.ShowMnemonicError, err)
	}

	return nil
}
