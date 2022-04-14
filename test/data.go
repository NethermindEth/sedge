package test

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

func getAllFiles(entries []fs.FileInfo, srcPath, dstPath string) ([]copyOperation, error) {
	ops := []copyOperation{}
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "." && entry.Name() != ".." {
			newSrc := filepath.Join(srcPath, entry.Name())
			newDst := filepath.Join(dstPath, entry.Name())
			err := os.Mkdir(newDst, os.ModePerm)
			if err != nil {
				return nil, err
			}
			newEntries, err := ioutil.ReadDir(newSrc)
			if err != nil {
				return nil, err
			}
			newOps, err := getAllFiles(newEntries, newSrc, newDst)
			if err != nil {
				return nil, err
			}
			ops = append(ops, newOps...)
		} else {
			filePath := filepath.Join(srcPath, entry.Name())
			fileDstPath := filepath.Join(dstPath, entry.Name())
			file, err := os.Open(filePath)
			if err != nil {
				return nil, err
			}
			ops = append(ops, copyOperation{
				src: file,
				dst: fileDstPath,
			})
		}
	}
	return ops, nil
}

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
