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
