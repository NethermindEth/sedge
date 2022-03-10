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
package cmd

import (
	"os"
	"regexp"
	"strings"

	"github.com/NethermindEth/1Click/configs"
	"github.com/NethermindEth/1Click/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	network          string
	path             string
	existingMnemonic bool
)

var (
	//Windows and Unix path
	rePath = regexp.MustCompile(`^[a-zA-Z]+:(\\[a-zA-Z0-9_.-]+)+|^~{0,1}\/[a-zA-Z0-9~]+(\/[a-zA-Z0-9_.-]+)+`)
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys [flags]",
	Short: "Generate keystore folder",
	Long: `Generate keystore folder using the eth2.0-deposit-cli tool.
	
New mnemonic will be generated if -e/--existing flag is not provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Validate network when several networks are supported

		// Validate path. It must be an absolute and correct path
		log.Debugf("Path to keystore file: %s", path)
		if !rePath.MatchString(path) {
			log.Fatalf(configs.InvalidVolumePathError, path)
		}

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

		log.Info(configs.GeneratingKeystore)
		if err := utils.GenerateValidatorKey(existingMnemonic, network, path); err != nil {
			log.Fatalf(configs.GeneratingKeystoreError, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)

	// Get PWD
	pwd, err := os.Getwd()
	if err != nil {
		log.WithField(configs.Component, "Root Init").Fatal(err)
	}
	log.Debug(pwd)

	// Local flags
	keysCmd.Flags().StringVarP(&network, "network", "n", "mainnet", "Target network. e.g. mainnet, prater, etc.")

	keysCmd.Flags().StringVarP(&path, "path", "p", pwd, "Absolute path to keystore folder. e.g. /home/user/keystore")

	keysCmd.Flags().BoolVarP(&existingMnemonic, "existing", "e", false, "Use existing mnemonic")
}
