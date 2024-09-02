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
package profile

import (
	"errors"
	"fmt"
	"math"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/NethermindEth/sedge/internal/monitoring/utils"
)

var pathRe = regexp.MustCompile(`^(/|./|../|[^/ ]([^/ ]*/)*[^/ ]*$)`)

// Profile represents a profile file of a package
type Profile struct {
	Name                          string                         `yaml:"-"`
	HardwareRequirementsOverrides *HardwareRequirementsOverrides `yaml:"hardware_requirements_overrides,omitempty"`
	PluginOverrides               PluginOverrides                `yaml:"plugin_overrides"`
	Options                       []Option                       `yaml:"options"`
	Monitoring                    Monitoring                     `yaml:"monitoring"`
	API                           *APITarget                     `yaml:"api,omitempty"`
}

// Validate validates the profile file
func (p *Profile) Validate() error {
	var missingFields []string
	if len(p.Options) == 0 {
		missingFields = append(missingFields, "options")
	}

	invalidOptionsErr := errors.New("invalid options")
	invalidOptions := false
	for i, option := range p.Options {
		if err := option.validate(i); err != nil {
			invalidOptions = true
			invalidOptionsErr = fmt.Errorf("%w: %w", invalidOptionsErr, err)
		}
	}

	invalidMonitoringErr := p.Monitoring.validate()

	if len(missingFields) > 0 || invalidOptions || invalidMonitoringErr != nil {
		var err error = InvalidProfileError{
			message:       "Invalid profile",
			missingFields: missingFields,
		}
		if invalidOptions {
			err = fmt.Errorf("%w: %w", err, invalidOptionsErr)
		}
		if invalidMonitoringErr != nil {
			err = fmt.Errorf("%w: %w", err, invalidMonitoringErr)
		}
		return err
	}

	return nil
}

// HardwareRequirementsOverrides represents the hardware requirements overrides field of a profile
type HardwareRequirementsOverrides struct {
	MinCPUCores                 int  `yaml:"min_cpu_cores"`
	MinRAM                      int  `yaml:"min_ram"`
	MinFreeSpace                int  `yaml:"min_free_space"`
	StopIfRequirementsAreNotMet bool `yaml:"stop_if_requirements_are_not_met"`
}

// TODO: add validation for hardware requirements overrides

// PluginOverrides represents the plugin overrides field of a profile
type PluginOverrides struct {
	Image string `yaml:"image"`
}

// TODO: add validation for plugin overrides

// Option represents an option within the options field of a profile
type Option struct {
	Name        string    `yaml:"name"`
	Target      string    `yaml:"target"`
	Type        string    `yaml:"type"`
	Default     string    `yaml:"default"`
	Help        string    `yaml:"help"`
	Hidden      bool      `yaml:"hidden"`
	ValidateDef *Validate `yaml:"validate,omitempty"`
}

// Validate validates the option
func (o *Option) validate(idx int) error {
	var missingFields, invalidFields []string
	if o.Name == "" {
		missingFields = append(missingFields, "options.name")
	}
	if o.Target == "" {
		missingFields = append(missingFields, "options.target")
	}
	if o.Type == "" {
		missingFields = append(missingFields, "options.type")
	}
	if o.Help == "" {
		missingFields = append(missingFields, "options.help")
	}

	var invalidDefault bool
	if o.Default != "" {
		switch o.Type {
		case "str":
			if o.ValidateDef != nil {
				invalidDefault = !regexp.MustCompile(o.ValidateDef.Re2Regex).MatchString(o.Default)
			}
		case "int":
			val, err := strconv.Atoi(o.Default)
			invalidDefault = err != nil
			if o.ValidateDef != nil {
				if o.ValidateDef.MinValue != nil && val < int(*o.ValidateDef.MinValue) {
					invalidDefault = true
				}
				if o.ValidateDef.MaxValue != nil && val > int(*o.ValidateDef.MaxValue) {
					invalidDefault = true
				}
			}
		case "port":
			port, err := strconv.Atoi(o.Default)
			invalidDefault = err != nil || port <= 0 || port > math.MaxUint16
		case "float":
			val, err := strconv.ParseFloat(o.Default, 64)
			invalidDefault = err != nil
			if o.ValidateDef != nil {
				if o.ValidateDef.MinValue != nil && val < *o.ValidateDef.MinValue {
					invalidDefault = true
				}
				if o.ValidateDef.MaxValue != nil && val > *o.ValidateDef.MaxValue {
					invalidDefault = true
				}
			}
		case "bool":
			_, err := strconv.ParseBool(o.Default)
			invalidDefault = err != nil
		case "path_dir":
			invalidDefault = !pathRe.MatchString(o.Default)
		case "path_file":
			invalidDefault = !pathRe.MatchString(o.Default)
			if o.ValidateDef != nil {
				invalidDefault = filepath.Ext(o.Default) != o.ValidateDef.Format
			}
		case "uri":
			gotUrl, err := url.Parse(o.Default)
			if err != nil {
				invalidDefault = true
			} else if o.ValidateDef != nil {
				invalidDefault = !utils.Contains(o.ValidateDef.UriScheme, gotUrl.Scheme)
			}
		case "select":
			if o.ValidateDef == nil {
				missingFields = append(missingFields, "options.validate")
			} else {
				invalidDefault = !utils.Contains(o.ValidateDef.Options, o.Default)
			}
		default:
			invalidDefault = true
		}
	}
	if invalidDefault {
		invalidFields = append(invalidFields, "options.default")
	}

	if len(missingFields) > 0 || len(invalidFields) > 0 {
		return InvalidProfileError{
			message:       "Option #" + strconv.Itoa(idx+1) + " is invalid",
			missingFields: missingFields,
			invalidFields: invalidFields,
		}
	}

	return nil
}

// Validate represents the validate field of an option
type Validate struct {
	Re2Regex  string   `yaml:"re2_regex"`
	Format    string   `yaml:"format"`
	UriScheme []string `yaml:"uri_scheme"`
	MinValue  *float64 `yaml:"min_value,omitempty"`
	MaxValue  *float64 `yaml:"max_value,omitempty"`
	Options   []string `yaml:"options"`
}

// Monitoring represents the monitoring field of a profile
type Monitoring struct {
	Targets []MonitoringTarget `yaml:"targets"`
}

func (m *Monitoring) validate() error {
	err := errors.New("invalid monitoring")

	if len(m.Targets) == 0 {
		return fmt.Errorf("%w: %s", err, "there must be at least one monitoring target in the profile file")
	}

	ok := true
	for i, target := range m.Targets {
		if valErr := target.validate(i); valErr != nil {
			ok = false
			err = fmt.Errorf("%w: %w", err, valErr)
		}
	}
	if !ok {
		return err
	}
	return nil
}

// MonitoringTarget represents a monitoring target within the targets field of a monitoring
type MonitoringTarget struct {
	Service string `yaml:"service"`
	Port    *int   `yaml:"port"`
	Path    string `yaml:"path"`
}

func (m *MonitoringTarget) validate(idx int) error {
	var missingFields, invalidFields []string

	if m.Service == "" {
		missingFields = append(missingFields, "monitoring.targets.service")
	} else if len(strings.Split(m.Service, " ")) != 1 {
		invalidFields = append(invalidFields, "monitoring.targets.service")
	}

	if m.Port == nil {
		missingFields = append(missingFields, "monitoring.targets.port")
	} else if *m.Port <= 0 || *m.Port > math.MaxUint16 {
		invalidFields = append(invalidFields, "monitoring.targets.port")
	}

	if m.Path == "" {
		missingFields = append(missingFields, "monitoring.targets.path")
	} else {
		tmpUri := "http://localhost:8080" + m.Path
		if _, err := url.Parse(tmpUri); err != nil {
			invalidFields = append(invalidFields, "monitoring.targets.path")
		}
	}

	if len(missingFields) > 0 || len(invalidFields) > 0 {
		return InvalidProfileError{
			message:       "Monitoring target #" + strconv.Itoa(idx+1) + " is invalid",
			missingFields: missingFields,
			invalidFields: invalidFields,
		}
	}

	return nil
}

// API represents the api field of a profile
type APITarget struct {
	Service string `yaml:"service"`
	Port    int    `yaml:"port"`
}
