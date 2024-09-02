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
package data

import "errors"

var (
	ErrInstanceAlreadyExists       = errors.New("instance already exists")
	ErrInstanceNotFound            = errors.New("instance not found")
	ErrInvalidInstance             = errors.New("invalid instance")
	ErrInvalidInstanceDir          = errors.New("invalid instance directory")
	ErrTempDirDoesNotExist         = errors.New("temp directory does not exist")
	ErrTempIsNotDir                = errors.New("temp is not a directory")
	ErrMonitoringStackNotFound     = errors.New("monitoring stack not found")
	ErrInitializingMonitoringStack = errors.New("failed monitoring stack initialization")
	ErrReadingFile                 = errors.New("failed reading file")
	ErrWritingFile                 = errors.New("failed writing file")
	ErrStackNotInitialized         = errors.New("stack not initialized")
	ErrBackupAlreadyExists         = errors.New("backup already exists")
	ErrCreatingBackup              = errors.New("failed creating backup")
	ErrInvalidBackupName           = errors.New("invalid backup name")
	ErrBackupNotFound              = errors.New("backup not found")
)
