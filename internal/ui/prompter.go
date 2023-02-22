package ui

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/NethermindEth/sedge/internal/utils"
)

//go:generate mockgen -source=prompter.go -destination=prompter_mock.go -package=ui
type Prompter interface {
	Select(message string, defaultValue string, options []string) (int, error)
	Confirm(string, bool) (bool, error)
	Input(prompt, defaultValue string, required bool) (result string, err error)
	InputInt64(prompt string, defaultValue int64) (result int64, err error)
	InputFilePath(prompt, defaultValue string, required bool) (result string, err error)
	InputSecret(prompt string) (result string, err error)
	EthAddress(prompt string, defaultValue string) (result string, err error)
}

func NewPrompter() Prompter {
	return &prompter{}
}

type prompter struct{}

func (p *prompter) Select(message, defaultValue string, options []string) (result int, err error) {
	q := &survey.Select{
		Message:  message,
		Options:  options,
		PageSize: 5,
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

func (p *prompter) InputFilePath(prompt, defaultValue string, required bool) (result string, err error) {
	options := []survey.AskOpt{
		survey.WithValidator(func(ans interface{}) error {
			if str, ok := ans.(string); ok {
				fileInfo, err := os.Stat(str)
				if err != nil {
					return err
				}
				if fileInfo.IsDir() {
					return errors.New("is a directory not a file")
				}
			}
			return nil
		}),
	}
	if required {
		options = append(options, survey.WithValidator(survey.Required))
	}
	options = append(options)
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
		survey.WithValidator(func(ans interface{}) error {
			if str, ok := ans.(string); !ok {
				_, err := strconv.ParseInt(str, 10, 64)
				return err
			}
			return nil
		}),
	}
	q := &survey.Input{
		Message: prompt,
		Default: fmt.Sprintf("%d", defaultValue),
	}
	err = survey.AskOne(q, &result, options...)
	return
}

func (p *prompter) EthAddress(prompt string, defaultValue string) (result string, err error) {
	options := []survey.AskOpt{
		survey.WithValidator(func(ans interface{}) error {
			if str, ok := ans.(string); ok && !utils.IsAddress(str) {
				if len(str) > 0 && !utils.IsAddress(str) {
					return errors.New("not a valid ethereum address")
				}
			}
			return nil
		}),
	}
	q := &survey.Input{
		Message: prompt,
		Default: defaultValue,
	}
	err = survey.AskOne(q, &result, options...)
	return
}
