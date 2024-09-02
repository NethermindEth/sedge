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

import "strings"

type InvalidProfileError struct {
	message       string
	invalidFields []string
	missingFields []string
}

func (e InvalidProfileError) Error() string {
	// Nil error
	if e.message == "" {
		return ""
	}

	if len(e.invalidFields) == 0 && len(e.missingFields) == 0 {
		return e.message
	}

	msg := e.message + " -> "
	if len(e.invalidFields) > 0 {
		msg += "invalid fields: " + strings.Join(e.invalidFields, ", ")
	}
	if len(e.missingFields) > 0 {
		msg += "missing fields: " + strings.Join(e.missingFields, ", ")
	}
	return msg
}
