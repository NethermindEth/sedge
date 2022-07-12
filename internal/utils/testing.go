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
package utils

import (
	"fmt"
)

/*
CheckErr :
Helper function for tests. Check if a given error was expected.

params :-
a. desc string
Description of the error in case of failure
b. isError bool
Whether the error was expected or not
c. err error
Error to check

returns :-
a. error
Error if any
*/
func CheckErr(descr string, isErr bool, err error) error {
	l := err == nil && isErr
	r := err != nil && !isErr

	if l || r {
		return fmt.Errorf("%s failed, unexpected error value: %v", descr, err)
	}
	return nil
}
