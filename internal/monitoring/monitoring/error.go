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
package monitoring

import "errors"

var (
	ErrInitializingMonitoringMngr    = errors.New("error initializing monitoring manager")
	ErrCheckingMonitoringStack       = errors.New("error checking monitoring stack status")
	ErrRunningMonitoringStack        = errors.New("error running monitoring stack")
	ErrInstallingMonitoringMngr      = errors.New("error installing monitoring manager")
	ErrConfiguringMonitoringServices = errors.New("error configuring monitoring services")
	ErrNonexistingTarget             = errors.New("target to remove does not exist")
)
