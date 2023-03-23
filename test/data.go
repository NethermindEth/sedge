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
package test

// notest

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func copy(srcFile fs.File, dest string) error {
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}

type copyOperation struct {
	src fs.File
	dst string
}

// Create copyOperations for every file on any subdirectory in "srcPath" to "dstPath".
func getAllFiles(entries []fs.FileInfo, srcPath, dstPath string) ([]copyOperation, error) {
	ops := []copyOperation{}
	for _, entry := range entries { // loop through files infos
		if entry.IsDir() && entry.Name() != "." && entry.Name() != ".." { // entry is a directory
			// build new source directory path
			newSrc := filepath.Join(srcPath, entry.Name())
			// build new destiny directory path
			newDst := filepath.Join(dstPath, entry.Name())
			// create destiny directory path if it doesn't exist
			err := os.Mkdir(newDst, os.ModePerm)
			if err != nil {
				return nil, err
			}
			// get new entries from new source directory
			newEntries, err := ioutil.ReadDir(newSrc)
			if err != nil {
				return nil, err
			}
			// recursively create copy operations from new source to new destiny using new entries
			newOps, err := getAllFiles(newEntries, newSrc, newDst)
			if err != nil {
				return nil, err
			}
			// add copy operations to result operations
			ops = append(ops, newOps...)
		} else { // entry is a file
			// create source file path
			filePath := filepath.Join(srcPath, entry.Name())
			// create destiny file path
			fileDstPath := filepath.Join(dstPath, entry.Name())
			// open file from file path
			file, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}
			// add copy operation to result operations
			ops = append(ops, copyOperation{
				src: file,
				dst: fileDstPath,
			})
		}
	}
	return ops, nil
}

// Copy all sub directory and files from "srcPath" to "dstPath"
// maintaining the same structure.
func PrepareTestCaseDir(srcPath, dstPath string) error {
	srcDirs, err := ioutil.ReadDir(srcPath)
	if err != nil {
		return err
	}

	copyOps, err := getAllFiles(srcDirs, srcPath, dstPath)
	if err != nil {
		return err
	}

	for _, copyOp := range copyOps {
		err := copy(copyOp.src, copyOp.dst)
		defer copyOp.src.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
