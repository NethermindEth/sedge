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
	"strings"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/utils"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var logLevel string

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sedge",
		Short: "A brief description of your application",
		Long:  `A tool to allow deploying validators with ease.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initLogging()
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
	cmd.PersistentFlags().StringVar(&logLevel, "logLevel", "info", "Set Log Level, e.gg panic, fatal, error, warn, warning, info, debug. trace")
	return cmd
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
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{configs.Component},
		TimestampFormat: "2006-01-02 15:04:05 --",
	})

	level, err := log.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		log.WithField(configs.Component, "Logger Init").Error(err)
		return
	}
	log.SetLevel(level)
	log.WithField(configs.Component, "Logger Init").Infof("Log level: %+v", logLevel)
}
