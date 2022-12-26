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
	"github.com/NethermindEth/sedge/internal/pkg/generation"
	"os"
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/utils"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sedge",
		Short: "A brief description of your application",
		Long:  `A tool to allow deploying validators with ease.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initConfig()
			checkVersion()
		},
		// TODO: Start the TUI engine in this callback. Default behavior
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
	// Disable completion default cmd
	cmd.CompletionOptions.DisableDefaultCmd = true
	// Persistent flags
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sedge.yaml)")
	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".sedge" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(configs.ConfigFileName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, "Config file not found on the path provided nor in the home directory")

		// Generate config file
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		fmt.Printf("Generating config file in the %s directory\n", home)

		err = generation.GenerateConfig(home)
		cobra.CheckErr(err)

		viper.ReadInConfig()
		cobra.CheckErr(err)
	}

	initLogging()
}

func checkVersion() {
	// Check version
	ok, err := utils.IsLatestVersion()
	if err != nil {
		log.Warnf("%s %e", configs.UnableToCheckVersion, err)
	} else if !ok {
		log.Warnf("%s %s", configs.NeedVersionUpdate, utils.CurrentVersion())
	} else {
		log.Infof("%s %s", configs.VersionUpdated, utils.CurrentVersion())
	}
}

/*
initLogging :
This function is responsible for :-
initializing the logging configurations
params :-
none
returns :-
none
*/
func initLogging() {
	var config configs.LogConfig

	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{configs.Component},
		TimestampFormat: "2006-01-02 15:04:05 --",
	})

	err := viper.UnmarshalKey("logs", &config)
	if err != nil {
		log.WithField(configs.Component, "Logger Init").Errorf("Unable to decode into struct, %v", err)
		return
	}
	log.WithField(configs.Component, "Logger Init").Infof("Logging configuration: %+v", config)

	level, err := log.ParseLevel(strings.ToLower(config.Level))
	if err != nil {
		log.WithField(configs.Component, "Logger Init").Error(err)
		return
	}
	log.SetLevel(level)
}
