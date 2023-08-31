package utils

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v54/github"
)

func DownloadGithubFile(
	client *github.Client,
	dest,
	owner,
	repository,
	path,
	ref string,
) (err error) {
	stream, _, err := client.Repositories.DownloadContents(
		context.Background(),
		owner,
		repository,
		path,
		&github.RepositoryContentGetOptions{
			Ref: ref,
		},
	)
	if err != nil {
		return
	}

	file, err := os.Create(dest)
	if err != nil {
		return
	}
	defer func() {
		closeErr := file.Close()
		if err == nil {
			err = closeErr
		}
	}()

	_, err = io.Copy(file, stream)
	if err != nil {
		return
	}

	return nil
}

func DownloadGithubObject(
	client *github.Client,
	dest,
	owner,
	repository,
	path,
	ref string,
) error {
	filedata, dirdata, _, err := client.Repositories.GetContents(
		context.Background(),
		owner,
		repository,
		path,
		&github.RepositoryContentGetOptions{
			Ref: ref,
		},
	)
	if err != nil {
		return err
	}

	if filedata != nil {
		err = DownloadGithubFile(
			client,
			strings.Join([]string{dest, filedata.GetName()}, "/"),
			owner,
			repository,
			path,
			ref,
		)
		if err != nil {
			return err
		}
	}

	if dirdata != nil {
		err = os.MkdirAll(
			dest,
			0o755,
		)
		if err != nil {
			return err
		}

		for _, content := range dirdata {
			if content != nil {
				newDest := dest
				if content.GetType() != "file" {
					filepath.Join(newDest, content.GetName())
				}
				err = DownloadGithubObject(
					client,
					dest,
					owner,
					repository,
					content.GetPath(),
					ref,
				)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
