package ui

import (
	"fmt"
	"io"
	"io/ioutil"
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

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf(configs.PrintFileError, err)
	}

	fmt.Fprintln(w, string(content))
	return nil
}
