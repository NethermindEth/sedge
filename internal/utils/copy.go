package utils

import (
	"fmt"
	"io"
	"os"
	"path"
)

// CopyFile:
// Copies the file from src to dst, if the directory of the dst doesn't
// exist then is created before the copy.
//
// params :-
// a. src string
// Source path
// b. dst string
// Destination path
func CopyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	if err := os.MkdirAll(path.Dir(dst), os.ModePerm); err != nil {
		return err
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}
