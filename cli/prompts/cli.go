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
package prompts

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/NethermindEth/sedge/configs"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
)

type promptCli struct{}

func NewPromptCli() Prompt {
	return &promptCli{}
}

func (p *promptCli) Passphrase() string {
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
		return p.Passphrase()
	}

	return result
}

func (p *promptCli) ExistingVal() int64 {
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
		return p.ExistingVal()
	}

	index, _ := strconv.ParseInt(result, 10, 64)
	return index
}

func (p *promptCli) NumberVal() int64 {
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

func (p *promptCli) Eth1Withdrawal() (string, error) {
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
		return "", fmt.Errorf(configs.PromptFailedError, err)
	}
	return result, nil
}

func (p *promptCli) FeeRecipient() (string, error) {
	// notest
	validate := func(input string) error {
		if input != "" && !utils.IsAddress(input) {
			return errors.New(configs.InvalidFeeRecipientError)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter the Fee Recipient address. You can leave it blank and press enter (not recommended)",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf(configs.PromptFailedError, err)
	}
	return result, nil
}
