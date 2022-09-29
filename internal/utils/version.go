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
	"github.com/google/go-github/v47/github"
	"os/exec"
)

const (
	// owner is the owner of the repository
	owner = "NethermindEth"
	// repo is the name of the repository
	repo = "sedge"
	// MsgNeedVersionUpdate is the message to display when a new Version is available
	MsgUnableToCheckVersion = "Unable to check for new Version. Please check manually at " +
		"https://github.com/NethermindEth/sedge/releases, with error:"
	MsgNeedVersionUpdate = "A new Version of sedge is available. Please update to the latest Version. See " +
		"https://github.com/NethermindEth/sedge/releases for more information. Latest detected tag:"
	MsgVersionUpdated = "You are running the latest version of sedge. Version: "
)

var ErrorNoTag = errors.New("no tag found on Github")
var ErrorCheckingVersion = errors.New("error while checking for new Version, please check your internet connection")

var Version string

func init() {
	if Version == "" {
		Version = versionFromGit()
	}
}

func versionFromGit() string {
	cmd := "git tag | sort | tail -n 1"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return ""
	}
	outVersion := string(out)
	return outVersion[:len(outVersion)-1]
}

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

func CurrentVersion() string {
	return Version
}

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
