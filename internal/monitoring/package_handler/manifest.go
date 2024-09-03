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
package package_handler

import (
	"errors"
	"fmt"

	"github.com/distribution/reference"
)

// Manifest represents the manifest file of a package
type Manifest struct {
	Version              string               `yaml:"version"`
	Name                 string               `yaml:"name"`
	Upgrade              string               `yaml:"upgrade"`
	HardwareRequirements hardwareRequirements `yaml:"hardware_requirements"`
	Plugin               *Plugin              `yaml:"plugin"`
	Profiles             []string             `yaml:"profiles"`
}

func (m *Manifest) validate() error {
	var missingFields []string
	if m.Version == "" {
		missingFields = append(missingFields, "version")
	}
	if m.Name == "" {
		missingFields = append(missingFields, "name")
	}
	if m.Upgrade == "" {
		missingFields = append(missingFields, "upgrade")
	}
	if len(m.Profiles) == 0 {
		missingFields = append(missingFields, "profiles")
	}

	hardReqErr := m.HardwareRequirements.validate()

	var pluginErr error
	if m.Plugin != nil {
		pluginErr = m.Plugin.validate()
	}

	profileErr := errors.New("invalid profiles")
	invalidProfiles := false
	for i, profile := range m.Profiles {
		if profile == "" {
			invalidProfiles = true
			profileErr = fmt.Errorf("%w: profile %d", profileErr, i)
		}
	}

	if hardReqErr != nil || pluginErr != nil || invalidProfiles || len(missingFields) > 0 {
		var err error = InvalidConfError{
			message:       "Invalid manifest file",
			missingFields: missingFields,
		}
		if hardReqErr != nil {
			err = fmt.Errorf("%w: %w", err, hardReqErr)
		}
		if pluginErr != nil {
			err = fmt.Errorf("%w: %w", err, pluginErr)
		}
		if invalidProfiles {
			err = fmt.Errorf("%w: %w", err, profileErr)
		}
		return err
	}

	return nil
}

type hardwareRequirements struct {
	MinCPUCores                 int  `yaml:"min_cpu_cores"`
	MinRAM                      int  `yaml:"min_ram"`
	MinFreeSpace                int  `yaml:"min_free_space"`
	StopIfRequirementsAreNotMet bool `yaml:"stop_if_requirements_are_not_met"`
}

func (h *hardwareRequirements) validate() error {
	var invalidFields []string
	if h.MinCPUCores < 0 {
		invalidFields = append(invalidFields, "hardware_requirements.min_cpu_cores -> (negative value)")
	}
	if h.MinRAM < 0 {
		invalidFields = append(invalidFields, "hardware_requirements.min_ram -> (negative value)")
	}
	if h.MinFreeSpace < 0 {
		invalidFields = append(invalidFields, "hardware_requirements.min_free_space -> (negative value)")
	}
	if len(invalidFields) > 0 {
		return InvalidConfError{
			message:       "Invalid hardware requirements",
			invalidFields: invalidFields,
		}
	}
	return nil
}

type Plugin struct {
	Image string `yaml:"image"`
}

func (p *Plugin) validate() error {
	var invalidFields []string
	// Validate plugin image field is a valid docker image
	if p.Image != "" {
		// Parse the image name
		if _, err := reference.ParseNormalizedNamed(p.Image); err != nil {
			invalidFields = append(invalidFields, fmt.Sprintf("plugin.image -> (invalid docker image: %v)", err))
		}
	} else {
		invalidFields = append(invalidFields, "plugin.image -> (empty)")
	}
	if len(invalidFields) > 0 {
		return InvalidConfError{
			message:       "Invalid plugin",
			invalidFields: invalidFields,
		}
	}
	return nil
}
