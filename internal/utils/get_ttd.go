package utils

import (
	"path/filepath"
	"regexp"

	"github.com/NethermindEth/1click/templates"
)

var reTTD = regexp.MustCompile(`(.*)_TTD=(.*)`)

func TTD(network, clientType, client string) (bool, error) {
	content, err := templates.Envs.ReadFile(filepath.Join("envs", network, clientType, client+".tmpl"))
	if err != nil {
		return false, err
	}

	if m := reTTD.FindStringSubmatch(string(content)); m != nil {
		return true, nil
	}

	return false, nil
}
