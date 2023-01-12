package prompts

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func (p *promptCli) Confirm(label string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
		Default:   "Y",
	}
	input, err := prompt.Run()
	var confirmation bool
	if strings.ToLower(input) == "y" {
		confirmation = true
	} else if strings.ToLower(input) == "n" {
		confirmation = false
	} else if input == "" {
		confirmation = true
	} else if err == nil {
		return false, fmt.Errorf("unexpected input %s", input)
	}
	if !confirmation && err != nil {
		err = nil
	}
	return confirmation, err
}
