package prompts

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func (p *promptCli) Select(label string, options ...string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: options,
		Templates: &promptui.SelectTemplates{
			Selected: fmt.Sprintf("%s: {{ . | cyan }}", label),
		},
		Searcher: func(input string, index int) bool {
			return strings.HasPrefix(options[index], input)
		},
	}
	_, selected, err := prompt.Run()
	return selected, err
}
