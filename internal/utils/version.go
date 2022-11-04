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
	"context"
	"errors"
	"text/template"

	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/google/go-github/v47/github"
)

const (
	// owner is the owner of the repository
	owner = "NethermindEth"
	// repo is the name of the repository
	repo = "sedge"
)

var (
	ErrorNoTag           = errors.New("no tag found on Github")
	ErrorCheckingVersion = errors.New("error while checking for new Version, please check your internet connection")
)

// Version represents the current version of Sedge
var Version string

func init() {
	if Version == "" {
		version, err := versionFromGit()
		if err != nil {
			Version = "v0.0.0"
		} else {
			Version = version
		}
	}
}

/*
versionFromGit :
Extract the version using git on the terminal.

params :-
None

returns :-
a. string
Version extracted from git calling a command on the terminal
a. err error
Error if any
*/
func versionFromGit() (string, error) {
	const cmd = `
#!/bin/bash

git tag | sort | tail -n 1
`
	r := commands.NewCMDRunner(commands.CMDRunnerOptions{
		RunAsAdmin: false,
	})
	tt := template.Must(template.New("").Parse(cmd))
	out, err := r.RunBash(commands.BashScript{
		Tmp:       tt,
		GetOutput: true,
		Data:      struct{}{},
	})
	if err != nil {
		return "", err
	}
	outVersion := out
	if len(outVersion) > 0 {
		return outVersion[:len(outVersion)-1], nil
	}
	return outVersion, errors.New("no version found")
}

/*
latestVersionOnGithub :
Fetch the latest version on GitHub

params :-
None

returns :-
a. string
Version of Sedge fetched from GitHub
a. err error
Error if any
*/
func latestVersionOnGithub() (string, error) {
	client := github.NewClient(nil)
	tags, _, err := client.Repositories.ListTags(context.Background(), owner, repo, nil)
	if err != nil {
		return "", ErrorCheckingVersion
	}
	if len(tags) > 0 {
		latestTag := tags[0]
		return latestTag.GetName(), nil
	}
	return "", ErrorNoTag
}

/*
CurrentVersion :
Export the current version of Sedge

params :-
None

returns :-
a. string
Return the current version of Sedge
*/
func CurrentVersion() string {
	return Version
}

/*
IsLatestVersion :
Check if we are running the latest version that was released on GitHub

params :-
None

returns :-
a. bool
Returns true if we are on the latest version, but false if we have a different version that latest on Github.
a. err error
Error if any
*/
func IsLatestVersion() (bool, error) {
	versionOnGithub, err := latestVersionOnGithub()
	if err != nil {
		return false, err
	}
	if versionOnGithub == "" {
		return false, ErrorNoTag
	}
	if Version == versionOnGithub {
		return true, nil
	}
	return false, nil
}
