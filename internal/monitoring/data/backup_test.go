package data

import (
	"archive/tar"
	"strconv"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBackupId(t *testing.T) {
	b := Backup{
		InstanceId: "mock-avs-default",
		Timestamp:  time.Unix(1696367916, 0),
		Version:    "v5.5.0",
		Commit:     "a3406616b848164358fdd24465b8eecda5f5ae34",
		Url:        "https://github.com/NethermindEth/mock-avs-pkg",
	}
	assert.Equal(t, b.Id(), "33de69fe9225b95c8fb909cb418e5102970c8d73")
}

func TestParseBackupName(t *testing.T) {
	tc := []struct {
		name       string
		backupName string
		instanceId string
		timestamp  time.Time
		err        error
	}{
		{
			name:       "valid backup name",
			backupName: "mock-avs-default-1696317683.tar",
			instanceId: "mock-avs-default",
			timestamp:  time.Unix(1696317683, 0),
			err:        nil,
		},
		{
			name:       "no .tar file",
			backupName: "mock-avs-default-1696317683",
			instanceId: "",
			timestamp:  time.Time{},
			err:        ErrInvalidBackupName,
		},
		{
			name:       "without dash separator between instance ID and timestamp",
			backupName: "mock-avs-default1696317683.tar",
			instanceId: "",
			timestamp:  time.Time{},
			err:        ErrInvalidBackupName,
		},
		{
			name:       "invalid timestamp",
			backupName: "mock-avs-default-1696317683a.tar",
			instanceId: "",
			timestamp:  time.Time{},
			err:        ErrInvalidBackupName,
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			instanceId, timestamp, err := ParseBackupName(tt.backupName)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.instanceId, instanceId)
				assert.Equal(t, tt.timestamp.Unix(), timestamp.Unix())
			}
		})
	}
}

func TestBackupFromTar(t *testing.T) {
	fs := afero.NewOsFs()

	// Create backup tar
	backupTar, err := afero.TempFile(fs, "", "backup-*.tar")
	require.NoError(t, err)
	defer backupTar.Close()

	// Create state.json
	stateFile, err := afero.TempFile(fs, "", "state-*.json")
	require.NoError(t, err)
	defer stateFile.Close()
	stateData := []byte(
		`{
			"name": "mock-avs",
			"url": "https://github.com/NethermindEth/mock-avs-pkg",
			"version": "v5.5.0",
			"spec_version": "v0.1.0",
			"commit": "a3406616b848164358fdd24465b8eecda5f5ae34",
			"profile": "option-returner",
			"tag": "default",
			"monitoring": {
			  "targets": [
				{
				  "service": "main-service",
				  "port": "8080",
				  "path": "/metrics"
				}
			  ]
			},
			"api": {
			  "service": "main-service",
			  "port": "8080"
			},
			"plugin": {
			  "image": "mock-avs-plugin:v0.1.0"
			}
		  }`,
	)
	_, err = stateFile.Write(stateData)
	require.NoError(t, err)

	// Create timestamp
	timestamp := time.Unix(1696367916, 0)
	timestampData := []byte(strconv.FormatInt(timestamp.Unix(), 10))
	timestampFile, err := afero.TempFile(fs, "", "timestamp-*")
	require.NoError(t, err)
	defer timestampFile.Close()
	timestampFile.Write(timestampData)

	// Build backup tar
	tarWriter := tar.NewWriter(backupTar)
	defer tarWriter.Close()
	// Add state.json
	stateFileInfo, err := stateFile.Stat()
	require.NoError(t, err)
	h, err := tar.FileInfoHeader(stateFileInfo, "")
	require.NoError(t, err)
	h.Name = "data/state.json"
	err = tarWriter.WriteHeader(h)
	require.NoError(t, err)
	_, err = tarWriter.Write(stateData)
	require.NoError(t, err)
	// Add timestamp
	timestampFileInfo, err := timestampFile.Stat()
	require.NoError(t, err)
	h, err = tar.FileInfoHeader(timestampFileInfo, "")
	require.NoError(t, err)
	h.Name = "timestamp"
	err = tarWriter.WriteHeader(h)
	require.NoError(t, err)
	_, err = tarWriter.Write(timestampData)
	require.NoError(t, err)
	// Close tar writer
	tarWriter.Close()

	// Check backup from tar
	b, err := BackupFromTar(fs, backupTar.Name())
	require.NoError(t, err)
	require.NotNil(t, b)
	assert.Equal(t,
		Backup{
			InstanceId: "mock-avs-default",
			Timestamp:  timestamp,
			Version:    "v5.5.0",
			Commit:     "a3406616b848164358fdd24465b8eecda5f5ae34",
			Url:        "https://github.com/NethermindEth/mock-avs-pkg",
		},
		*b)
}

func TestLoadBackupTarStateJson(t *testing.T) {
	fs := afero.NewOsFs()
	tarFile, err := afero.TempFile(fs, t.TempDir(), "backup-*.tar")
	require.NoError(t, err)
	defer tarFile.Close()
	tarWriter := tar.NewWriter(tarFile)
	tarAddStateJson(t, tarWriter, []byte(`
	{
		"name": "mock-avs",
		"url": "https://github.com/NethermindEth/mock-avs-pkg",
		"version": "v5.5.0",
		"spec_version": "v0.1.0",
		"commit": "a3406616b848164358fdd24465b8eecda5f5ae34",
		"profile": "option-returner",
		"tag": "second",
		"monitoring": {
		  "targets": [
			{
			  "service": "main-service",
			  "port": "8080",
			  "path": "/metrics"
			}
		  ]
		},
		"api": {
		  "service": "main-service",
		  "port": "8080"
		},
		"plugin": {
		  "image": "mock-avs-plugin:v0.1.0"
		}
	  }
	`))
	got, err := loadBackupTarStateJson(fs, tarFile.Name())
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, Instance{
		Name:        "mock-avs",
		Tag:         "second",
		URL:         "https://github.com/NethermindEth/mock-avs-pkg",
		Version:     "v5.5.0",
		SpecVersion: "v0.1.0",
		Commit:      "a3406616b848164358fdd24465b8eecda5f5ae34",
		Profile:     "option-returner",
		MonitoringTargets: MonitoringTargets{
			Targets: []MonitoringTarget{
				{
					Service: "main-service",
					Port:    "8080",
					Path:    "/metrics",
				},
			},
		},
		APITarget: &APITarget{
			Service: "main-service",
			Port:    "8080",
		},
		Plugin: &Plugin{
			Image: "mock-avs-plugin:v0.1.0",
		},
	}, *got)
}

func TestLoadBackupTarTimestamp(t *testing.T) {
	fs := afero.NewOsFs()
	tarFile, err := afero.TempFile(fs, t.TempDir(), "backup-*.tar")
	require.NoError(t, err)
	defer tarFile.Close()
	tarWriter := tar.NewWriter(tarFile)
	timestamp := time.Unix(1696367916, 0)
	tarAddTimestamp(t, tarWriter, timestamp)
	got, err := loadBackupTarTimestamp(fs, tarFile.Name())
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.True(t, timestamp.Equal(got))
}
