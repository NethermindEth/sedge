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
package package_handler

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidFilePath            = errors.New("invalid file path")
	ErrInvalidDirPath             = errors.New("invalid directory path")
	ErrInvalidChecksum            = errors.New("invalid checksum")
	ErrNoVersionsFound            = errors.New("no versions found")
	ErrInvalidVersion             = errors.New("invalid version")
	ErrVersionNotFound            = errors.New("version not found")
	ErrProfileNotFound            = errors.New("profile not found")
	ErrNoPlugin                   = errors.New("no plugin found")
	ErrProfileComposeFileNotFound = errors.New("profile compose file not found")
	ErrBuildContextNotAllowed     = errors.New("build context not allowed")
)

// PackageFileNotFoundError is returned when a package file is not found.
type PackageFileNotFoundError struct {
	fileRelativePath string
	packagePath      string
}

func (e PackageFileNotFoundError) Error() string {
	return "package file not found: " + e.fileRelativePath + " in package " + e.packagePath
}

// PackageDirNotFoundError is returned when a package directory is not found.
type PackageDirNotFoundError struct {
	dirRelativePath string
	packagePath     string
}

func (e PackageDirNotFoundError) Error() string {
	return "package directory not found: " + e.dirRelativePath + " in package " + e.packagePath
}

// InvalidConfError is returned when a manifest or profile file is invalid or incomplete.
type InvalidConfError struct {
	message       string
	invalidFields []string
	missingFields []string
}

func (e InvalidConfError) Error() string {
	// Nil error
	if e.message == "" {
		return ""
	}

	if len(e.invalidFields) == 0 && len(e.missingFields) == 0 {
		return e.message
	}

	msg := e.message + " -> "
	if len(e.invalidFields) > 0 {
		msg += "invalid fields: " + strings.Join(e.invalidFields, ", ")
	}
	if len(e.missingFields) > 0 {
		msg += "missing fields: " + strings.Join(e.missingFields, ", ")
	}
	return msg
}

// ReadingProfileError is returned when a profile cannot be read.
type ReadingProfileError struct {
	profileName string
}

func (e ReadingProfileError) Error() string {
	return "error reading profile " + e.profileName
}

// ParsingProfileError is returned when a profile cannot be parsed.
type ParsingProfileError struct {
	profileName string
}

func (e ParsingProfileError) Error() string {
	return "error parsing profile " + e.profileName
}

// ReadingManifestError is returned when a manifest cannot be read.
type ReadingManifestError struct {
	pkgPath string
}

func (e ReadingManifestError) Error() string {
	return "error reading manifest in package " + e.pkgPath
}

// ParsingManifestError is returned when a manifest cannot be parsed.
type ParsingManifestError struct {
	pkgPath string
}

func (e ParsingManifestError) Error() string {
	return "error parsing manifest in package " + e.pkgPath
}

// ReadingDotEnvError is returned when a .env file cannot be read.
type ReadingDotEnvError struct {
	pkgPath     string
	profileName string
}

func (e ReadingDotEnvError) Error() string {
	return "error reading .env file of profile " + e.profileName + " in package " + e.pkgPath
}

// RepositoryNotFoundOrPrivateError is returned when the specified repository URL
// cannot be found or accessed due to its private status. This error typically occurs
// when no credentials are provided and the repository is either private or does not exist.
type RepositoryNotFoundOrPrivateError struct {
	URL string
}

func (e RepositoryNotFoundOrPrivateError) Error() string {
	return fmt.Sprintf("repository %s not found or private", e.URL)
}

// RepositoryNotFoundError is returned when the given repository URL is not found.
type RepositoryNotFoundError struct {
	URL string
}

func (e RepositoryNotFoundError) Error() string {
	return fmt.Sprintf("repository %s not found", e.URL)
}
