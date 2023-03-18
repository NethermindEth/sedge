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
package utils

import (
	"fmt"
	"path/filepath"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	log "github.com/sirupsen/logrus"
)

/*
CheckContainers :
Check if containers of generated docker-compose script are running

params :-
a. generationPath string
Path to the generated docker-compose script

returns :-
a. string
Output of 'docker ps --services --filter status=running'
b. error
Error if any
*/
func CheckContainers(cmdRunner commands.CommandRunner, generationPath string) (string, error) {
	// Check if docker-compose script is running
	psCMD := cmdRunner.BuildDockerComposePSCMD(commands.DockerComposePsOptions{
		Path:          filepath.Join(generationPath, configs.DefaultDockerComposeScriptName),
		Services:      true,
		FilterRunning: true,
	})
	log.Debugf(configs.RunningCommand, psCMD.Cmd)
	psCMD.GetOutput = true
	rawServices, _, err := cmdRunner.RunCMD(psCMD)
	if err != nil || rawServices == "\n" {
		if rawServices == "\n" && err == nil {
			err = fmt.Errorf(configs.DockerComposePsReturnedEmptyError)
		}
		return "", fmt.Errorf(configs.ScriptIsNotRunningError, err)
	}

	return rawServices, nil
}
