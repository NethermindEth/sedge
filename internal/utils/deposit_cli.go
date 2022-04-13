package utils

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/NethermindEth/1click/configs"
	"github.com/NethermindEth/1click/internal/pkg/commands"
	"github.com/NethermindEth/1click/templates"
	log "github.com/sirupsen/logrus"
)

/*
GenerateValidatorKey :
Generates a validator key using the eth2.0-deposit-cli tool.
The key can be generated using a new or existing mnemonic.
Key's path is set to $(pwd)/keystore.

params :-
a. existing bool
True if the key is to be generated using an existing mnemonic. False if the key is to be generated using a new mnemonic.
b. network string
Target network.

returns :-
a. error
Error if any
*/
func GenerateValidatorKey(existing bool, network, path string) (err error) {
	// Check if image already exists
	inspectCmd := commands.Runner.BuildDockerInspectCMD(commands.DockerInspectOptions{
		Name: configs.DepositCLIDockerImageName,
	})
	inspectCmd.GetOutput = true
	if out, err := commands.Runner.RunCMD(inspectCmd); err != nil {
		// Output is of type: []\n Error: <text>
		// TODO: Check if the error is not "Error: No such image: <image_name>" in Windows
		if strings.Contains(out, "No such object:") {
			// Image does not exist. Build it
			log.Infof(configs.ImageNotFound, configs.DepositCLIDockerImageName)
			if err := buildDepositCliImage(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf(configs.CommandError, inspectCmd.Cmd, out)
		}
	}

	data := DepositCLI{
		Network: network,
		Path:    path,
	}

	// Get the template file
	var rawTmp []byte
	if existing {
		rawTmp, err = templates.DepositCLI.ReadFile("deposit-cli/existing.tmpl")
	} else {
		rawTmp, err = templates.DepositCLI.ReadFile("deposit-cli/new.tmpl")
	}

	if err != nil {
		return
	}

	// Parse the template
	tmp, err := template.New("deposit-cli").Parse(string(rawTmp))
	if err != nil {
		return
	}

	// Get the command as a string
	var cmd bytes.Buffer
	err = tmp.Execute(&cmd, data)
	if err != nil {
		return
	}

	// Run the command
	log.Infof(configs.RunningCommand, cmd.String())
	_, err = commands.Runner.RunCMD(commands.Command{
		Cmd:       cmd.String(),
		GetOutput: false,
		RunInPty:  true,
	})
	if err != nil {
		return
	}
	log.Infof(configs.KeysFoundAt, path+"/keystore")

	return nil
}

func buildDepositCliImage() error {
	// Run docker build
	buildCMD := commands.Runner.BuildDockerBuildCMD(commands.DockerBuildOptions{
		Path: configs.DepositCLIDockerImageUrl,
		Tag:  configs.DepositCLIDockerImageName,
	})
	log.Infof(configs.RunningCommand, buildCMD.Cmd)
	_, err := commands.Runner.RunCMD(buildCMD)
	if err != nil {
		return err
	}

	return nil
}
