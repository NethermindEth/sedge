package prompts

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
)

func (p *promptCli) Input(label string, required bool) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if required && len(s) == 0 {
				return errors.New("required input")
			}
			return nil
		},
	}
	return prompt.Run()
}

func (p *promptCli) InputHide(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: true,
	}
	return prompt.Run()
}

func (p *promptCli) InputNumber(label string) (int64, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			_, err := strconv.ParseInt(s, 10, 64)
			return err
		},
	}
	input, err := prompt.Run()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(input, 10, 64)
}

func (p *promptCli) InputFilePath(label string, required bool) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if required && len(s) == 0 {
				return errors.New("required input")
			}
			fileInfo, err := os.Stat(s)
			if err != nil {
				return err
			}
			if fileInfo.IsDir() {
				return fmt.Errorf("is a directory not a file")
			}
			return nil
		},
	}
	return prompt.Run()
}
