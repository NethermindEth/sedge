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
package services

import "errors"

var (
	ErrStartingContainer = errors.New("error starting container")
	ErrContainerNotFound = errors.New("container not found")
	ErrStoppingContainer = errors.New("error stopping container")
	ErrNetworksNotFound  = errors.New("networks not found")
)
