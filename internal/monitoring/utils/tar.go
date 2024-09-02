package utils

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func CompressToTarGz(srcDir string, tarFile io.Writer) error {
	gw := gzip.NewWriter(tarFile)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// walk through every file in the folder
	err := filepath.Walk(srcDir, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		header.Name, err = filepath.Rel(srcDir, file)
		if err != nil {
			return err
		}

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func DecompressTarGz(tarFile io.Reader, destDir string) error {
	log.Debugf("Decompressing tar file to %s", destDir)
	gr, err := gzip.NewReader(tarFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)

	for {
		header, err := tr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}
		target := filepath.Join(destDir, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			targetInfo, err := os.Stat(target)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					err = os.MkdirAll(target, 0o755)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			} else if !targetInfo.IsDir() {
				return fmt.Errorf("cannot decompress tar file: %s is not a directory", target)
			}
		case tar.TypeReg:
			targetDir := filepath.Dir(target)
			err = os.MkdirAll(targetDir, 0o755)
			if err != nil {
				return err
			}
			targetF, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			defer func() {
				closeErr := targetF.Close()
				if err == nil {
					err = closeErr
				}
			}()
			_, err = io.Copy(targetF, tr)
			if err != nil {
				return err
			}
		}
	}
}
