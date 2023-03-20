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
package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	"github.com/NethermindEth/sedge/internal/utils"
)

var ErrMissingDependencies = errors.New("missing dependencies")

func readFileContent(path string) (string, error) {
	raw, err := os.ReadFile(path)
	content := strings.TrimSpace(strings.TrimSuffix(string(raw), "\n"))
	return content, err
}

func checkDependencies(depsManager dependencies.DependenciesManager, dockerCompose bool, dependencies ...string) error {
	_, pending := depsManager.Check(dependencies)
	if len(pending) > 0 {
		return fmt.Errorf("%w: %s", ErrMissingDependencies, strings.Join(pending, ", "))
	}
	if utils.Contains(dependencies, "docker") {
		if err := depsManager.DockerEngineIsOn(); err != nil {
			return err
		}
	}
	if dockerCompose {
		if err := depsManager.DockerComposeIsInstalled(); err != nil {
			return err
		}
	}
	return nil
}
