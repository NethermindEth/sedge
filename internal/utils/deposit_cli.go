package utils

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/templates"
	log "github.com/sirupsen/logrus"
)

/*
GenerateValidatorKey :
Generates a validator key using the eth2.0-deposit-cli tool.
The key can be generated using a new or existing mnemonic.
Key's path is set to $(pwd)/keystore.

params :-
a. arg ValidatorKeyData
Data for keystore generation

returns :-
a. error
Error if any
*/
func GenerateValidatorKey(arg ValidatorKeyData) (err error) {
	// Check if image already exists
	inspectCmd := commands.Runner.BuildDockerInspectCMD(commands.DockerInspectOptions{
		Name: configs.DepositCLIDockerImageName,
	})
	inspectCmd.GetOutput = true
	if out, err := commands.Runner.RunCMD(inspectCmd); err != nil {
		// Output is of type: []\n Error: <text>
		// TODO: Check if the error is not "Error: No such image: <image_name>" exists in Windows
		if strings.Contains(out, "No such object:") {
			// TODO: Allow user to choose between building the image from the staking-deposit-cli repo or use netherminderth/staking-deposit-cli image
			// Image does not exist. Build it
			// log.Infof(configs.ImageNotFoundBuilding, configs.DepositCLIDockerImageName)
			// if err := buildDepositCliImage(); err != nil {
			// 	return err
			// }
			// Image does not exist. Pull it
			log.Infof(configs.ImageNotFoundPulling, configs.DepositCLIDockerImageName)
			if err := pullDepositCliImage(); err != nil {
				return err
			}
		} else {
			return fmt.Errorf(configs.CommandError, inspectCmd.Cmd, out)
		}
	}

	// Get the template file
	var rawTmp []byte
	if arg.Existing {
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

	lenPass := make([]struct{}, len(arg.Password))
	for i := 0; i < len(arg.Password); i++ {
		lenPass[i] = struct{}{}
	}

	data := DepositCLI{
		Network:               arg.Network,
		Path:                  arg.Path,
		LenPass:               lenPass,
		Image:                 configs.DepositCLIDockerImageName,
		Eth1WithdrawalAddress: arg.Eth1WithdrawalAddress,
	}

	// Get the command as a string with password hidden to print it
	var cmd bytes.Buffer
	err = tmp.Execute(&cmd, data)
	if err != nil {
		return
	}
	log.Infof(configs.RunningCommand, cmd.String())

	// Get the command as a string with password to execute it
	data.Password = arg.Password
	cmd.Reset()
	err = tmp.Execute(&cmd, data)
	if err != nil {
		return
	}

	// Run the command
	_, err = commands.Runner.RunCMD(commands.Command{
		Cmd:       cmd.String(),
		GetOutput: false,
		RunInPty:  true,
	})
	if err != nil {
		return
	}
	log.Info("deposit-cli tool exited")

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

func pullDepositCliImage() error {
	// Run docker pull
	pullCMD := commands.Runner.BuildDockerPullCMD(commands.DockerBuildOptions{
		Tag: configs.DepositCLIDockerImageName,
	})
	log.Infof(configs.RunningCommand, pullCMD.Cmd)
	_, err := commands.Runner.RunCMD(pullCMD)
	if err != nil {
		return err
	}

	return nil
}
