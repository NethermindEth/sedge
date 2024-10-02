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

// WaitExitError container waiting error, if any
type WaitExitError struct {
	Message string `json:"Message,omitempty"`
}

// WaitResponse ContainerWaitResponse
// OK response to ContainerWait operation
type WaitResponse struct {
	// error
	Error *WaitExitError `json:"Error,omitempty"`

	// Exit code of the container
	// Required: true
	StatusCode int64 `json:"StatusCode"`
}
