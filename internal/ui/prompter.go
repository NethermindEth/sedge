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
package ui

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// notest

//go:generate mockgen -package=sedge_mocks -destination=../../mocks/prompter.go github.com/NethermindEth/sedge/internal/ui Prompter
type Prompter interface {
	Select(message string, defaultValue string, options []string) (int, error)
	Confirm(string, bool) (bool, error)
	Input(prompt, defaultValue string, required bool) (result string, err error)
	InputInt64(prompt string, defaultValue int64) (result int64, err error)
	InputFilePath(prompt, defaultValue string, required bool, fileExtensions ...string) (result string, err error)
	InputURL(prompt, defaultValue string, required bool) (result string, err error)
	InputSecret(prompt string) (result string, err error)
	EthAddress(prompt string, defaultValue string, required bool) (result string, err error)
	InputList(prompt string, defaultValue []string, validator func([]string) error) (result []string, err error)
}

func NewPrompter() Prompter {
	return &prompter{}
}

type prompter struct{}

func (p *prompter) Select(message, defaultValue string, options []string) (result int, err error) {
	q := &survey.Select{
		Message:  message,
		Options:  options,
		PageSize: 10,
	}
	if defaultValue != "" {
		q.Default = defaultValue
	}
	err = survey.AskOne(q, &result)
	return
}

func (p *prompter) Confirm(question string, defaultValue bool) (answer bool, err error) {
	err = survey.AskOne(&survey.Confirm{
		Message: question,
		Default: defaultValue,
	}, &answer)
	return
}

func (p *prompter) Input(prompt, defaultValue string, required bool) (result string, err error) {
	options := make([]survey.AskOpt, 0)
	if required {
		options = append(options, survey.WithValidator(survey.Required))
	}
	q := &survey.Input{
		Message: prompt,
	}
	if defaultValue != "" {
		q.Default = defaultValue
	}
	err = survey.AskOne(q, &result, options...)
	return
}

func (p *prompter) InputFilePath(prompt, defaultValue string, required bool, fileExtensions ...string) (result string, err error) {
	options := []survey.AskOpt{
		survey.WithValidator(FilePathValidator),
	}
	if required {
		options = append(options, survey.WithValidator(survey.Required))
	}
	if len(fileExtensions) > 0 {
		options = append(options, survey.WithValidator(fileExtensionValidator(fileExtensions)))
	}
	q := &survey.Input{
		Message: prompt,
	}
	if defaultValue != "" {
		q.Default = defaultValue
	}
	err = survey.AskOne(q, &result, options...)
	return
}

func (p *prompter) InputURL(prompt, defaultValue string, required bool) (result string, err error) {
	options := []survey.AskOpt{
		survey.WithValidator(URLValidator),
	}
	if required {
		options = append(options, survey.WithValidator(survey.Required))
	}
	q := &survey.Input{
		Message: prompt,
	}
	if defaultValue != "" {
		q.Default = defaultValue
	}
	err = survey.AskOne(q, &result, options...)
	return
}

func (p *prompter) InputSecret(prompt string) (result string, err error) {
	options := []survey.AskOpt{
		survey.WithValidator(survey.MinLength(8)),
	}
	q := &survey.Password{
		Message: prompt,
	}
	err = survey.AskOne(q, &result, options...)
	return
}

func (p *prompter) InputInt64(prompt string, defaultValue int64) (result int64, err error) {
	options := []survey.AskOpt{
		survey.WithValidator(Int64Validator),
	}
	q := &survey.Input{
		Message: prompt,
		Default: fmt.Sprintf("%d", defaultValue),
	}
	err = survey.AskOne(q, &result, options...)
	return
}

func (p *prompter) EthAddress(prompt string, defaultValue string, required bool) (result string, err error) {
	options := []survey.AskOpt{
		survey.WithValidator(EthAddressValidator),
	}
	if required {
		options = append(options, survey.WithValidator(survey.Required))
	}
	q := &survey.Input{
		Message: prompt,
		Default: defaultValue,
	}
	err = survey.AskOne(q, &result, options...)
	return
}

func (p *prompter) InputList(prompt string, defaultValue []string, validator func([]string) error) (result []string, err error) {
	var (
		text            string
		defaultValueStr string
		options         []survey.AskOpt
	)
	if len(defaultValue) > 0 {
		prompt += " Default values are listed below."
		defaultValueStr = fmt.Sprintf("\n%s\n", strings.Join(defaultValue, "\n"))
	}
	q := &survey.Multiline{
		Message: prompt,
		Default: defaultValueStr,
	}
	if validator != nil {
		f := func(val interface{}) error {
			if input, ok := val.(string); ok {
				var toValidate []string
				for _, item := range strings.Split(input, "\n") {
					item = strings.TrimSpace(item)
					if item != "" {
						toValidate = append(toValidate, item)
					}
				}
				if err := validator(toValidate); err != nil {
					return err
				}
			}
			return nil
		}
		options = append(options, survey.WithValidator(f))
	}
	err = survey.AskOne(q, &text, options...)
	for _, item := range strings.Split(text, "\n") {
		item = strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return
}
