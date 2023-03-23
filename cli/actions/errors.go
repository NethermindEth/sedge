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
package actions

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedValidatorClient   = errors.New("unsupported validator client")
	ErrCreatingContextDir           = errors.New("error creating context dir")
	ErrValidatorImportCtBadExitCode = errors.New("validator import container exited with non-zero exit code")
	ErrUnknownLodestarPreset        = errors.New("unknown lodestar preset")
	ErrNetworkNotFound              = errors.New("network not found")
)

func newValidatorImportCtBadExitCodeError(ctId string, exitCode int64, logs string) error {
	return fmt.Errorf(`%w: validator-import service container with id %s ends with status code %d. Here are the logs for more details:
%s`, ErrValidatorImportCtBadExitCode, ctId, exitCode, logs)
}

func newUnknownLodestarPresetError(network string) error {
	return fmt.Errorf("%w for network %s", ErrUnknownLodestarPreset, network)
}
