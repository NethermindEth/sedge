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
package ui

import (
	"fmt"
	"io"
	"os"

	"github.com/NethermindEth/sedge/configs"
)

/*
PrintFileContent
This functions outputs the contents of a file to a writer

params :-
a. w io.Writer
Writer to be use for writing the file content
b. path string
Path to the file to write its content

returns :-
a. error
Error if any
*/
func PrintFileContent(w io.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf(configs.PrintFileError, err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf(configs.PrintFileError, err)
	}

	fmt.Fprintln(w, string(content))
	return nil
}
