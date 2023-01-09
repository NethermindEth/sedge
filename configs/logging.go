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
package configs

import "fmt"

var loggingDrivers = map[string]string{
	"none": "",
	"json": "json-file",
}

/*
ValidateLoggingFlag:
Validates the provider loggingFlag. A loggingFlag is valid
if is exactly equal to supported flags, for example: none, json.

params :-
a. loggingFlag string
Flag to validate.

returns :-
a. error
If the provided flag is valid then this error will be nil when returns, in
another case, the proper error is returned.
*/
func ValidateLoggingFlag(loggingFlag string) error {
	if _, ok := loggingDrivers[loggingFlag]; !ok {
		return fmt.Errorf(InvalidLoggingFlag, loggingFlag)
	}
	return nil
}

/*
GetLoggingDriver:
Returns the logging driver name assocaited with the provided logging flag. Panics
if the loggingFlag is not valid.

params :-
a. loggingFlag string
Flag of which want to know the associated driver.

returns :-
a. string
The associated logging driver.
*/
func GetLoggingDriver(loggingFlag string) string {
	return loggingDrivers[loggingFlag]
}

/*
ValidLoggingFlags:
Provides the list of supported logging flags.

returns :-
a. []string
The list of supported logging flags.
*/
func ValidLoggingFlags() []string {
	flags := make([]string, 0, len(loggingDrivers))
	for flag := range loggingDrivers {
		flags = append(flags, flag)
	}
	return flags
}
