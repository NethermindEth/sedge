package actions_test

import (
	"embed"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/test"
	mock_client "github.com/NethermindEth/sedge/test/mock_docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"
)

func TestImportKeys_ValidatorNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := validatorNotFoundHelper(t, ctrl)

	err := s.ImportValidatorKeys(actions.ImportValidatorKeysOptions{})
	assert.ErrorIs(t, err, services.ErrContainerNotFound)
}

func TestImportKeys_CheckValidatorFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wantError := errors.New("error")
	s := checkValidatorFailureHelper(t, ctrl, wantError)

	err := s.ImportValidatorKeys(actions.ImportValidatorKeysOptions{})
	assert.ErrorIs(t, err, wantError)
}

func TestImportKeys_ValidatorStopFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	s := validatorStopFailureHelper(t, ctrl)

	err := s.ImportValidatorKeys(actions.ImportValidatorKeysOptions{})
	assert.ErrorIs(t, err, services.ErrStoppingContainer)
}

func TestImportKeys_ValidatorRunning(t *testing.T) {
	clients := []string{"prysm", "lodestar"}
	for _, validatorClient := range clients {
		t.Run(validatorClient, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dockerClient := importKeysGoldenPath(t, ctrl, false)
			serviceManager := services.NewServiceManager(dockerClient)
			cmdRunner := test.SimpleCMDRunner{}
			s := actions.NewSedgeActions(dockerClient, serviceManager, &cmdRunner)

			from, err := setupKeystoreDir(t)
			if err != nil {
				t.Fatal(err)
			}

			generationPath := t.TempDir()

			s.ImportValidatorKeys(actions.ImportValidatorKeysOptions{
				ValidatorClient: validatorClient,
				Network:         "sepolia",
				From:            from,
				GenerationPath:  generationPath,
			})
		})
	}
}

func TestImportKeysCustom_ValidatorRunning(t *testing.T) {
	clients := []string{"lighthouse", "teku"}
	for _, validatorClient := range clients {
		t.Run(validatorClient, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dockerClient := importKeysGoldenPath(t, ctrl, true)
			serviceManager := services.NewServiceManager(dockerClient)
			cmdRunner := test.SimpleCMDRunner{}
			s := actions.NewSedgeActions(dockerClient, serviceManager, &cmdRunner)

			from, err := setupKeystoreDir(t)
			if err != nil {
				t.Fatal(err)
			}

			generationPath := t.TempDir()

			s.ImportValidatorKeys(actions.ImportValidatorKeysOptions{
				ValidatorClient: validatorClient,
				Network:         "sepolia",
				From:            from,
				GenerationPath:  generationPath,
			})
		})
	}
}

func TestImportKeys_UnsupportedClient(t *testing.T) {
	clients := []string{"", "unsupported", "kfjkdshjkr24"}
	for _, validatorClient := range clients {
		t.Run(validatorClient, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := unsupportedClientsHelper(t, ctrl)

			from, err := setupKeystoreDir(t)
			if err != nil {
				t.Fatal(err)
			}

			generationPath := t.TempDir()

			err = s.ImportValidatorKeys(actions.ImportValidatorKeysOptions{
				ValidatorClient: validatorClient,
				Network:         "sepolia",
				From:            from,
				GenerationPath:  generationPath,
			})
			assert.ErrorIs(t, err, actions.ErrUnsupportedValidatorClient)
		})
	}
}

//go:embed testdata/keystore
var keystoreTestData embed.FS

func setupKeystoreDir(t *testing.T) (string, error) {
	t.Helper()
	tempKeystore := t.TempDir()

	baseTestDir := "testdata/keystore"
	dirs := []string{""}
	for len(dirs) > 0 {
		currentDir := dirs[0]
		dirEntries, err := keystoreTestData.ReadDir(path.Join(baseTestDir, currentDir))
		if err != nil {
			return "", err
		}
		for _, entry := range dirEntries {
			if entry.IsDir() {
				dirs = append(dirs, filepath.Join(currentDir, entry.Name()))
			} else {
				entryData, err := keystoreTestData.ReadFile(path.Join(baseTestDir, currentDir, entry.Name()))
				if err != nil {
					return "", err
				}
				if err := os.MkdirAll(filepath.Join(tempKeystore, currentDir), 0o755); err != nil {
					return "", err
				}
				if err := ioutil.WriteFile(filepath.Join(tempKeystore, currentDir, entry.Name()), entryData, 0o755); err != nil {
					return "", err
				}
			}
		}
		dirs = dirs[1:]
	}
	return tempKeystore, nil
}

// importKeysGoldenPath returns a mocked docker client interface with all the
// required responses for a correct validator import keys container execution.
func importKeysGoldenPath(t *testing.T, ctrl *gomock.Controller, withCustomImage bool) client.APIClient {
	t.Helper()
	dockerClient := mock_client.NewMockAPIClient(ctrl)

	validatorCtId := "validatorctid"
	validatorImportCtName := "validator-import-client"
	validatorImportCtId := "validator-import-ct-id"

	// Mock ContainerList
	dockerClient.EXPECT().
		ContainerList(gomock.Any(), types.ContainerListOptions{
			All:     true,
			Filters: filters.NewArgs(filters.Arg("name", services.ServiceCtValidator)),
		}).
		Return([]types.Container{
			{ID: validatorCtId},
		}, nil)
	// Mock ContainerInspect
	inspectCall := dockerClient.EXPECT().
		ContainerInspect(gomock.Any(), services.ServiceCtValidator).
		Return(types.ContainerJSON{
			ContainerJSONBase: &types.ContainerJSONBase{
				ID: validatorCtId,
				State: &types.ContainerState{
					Running: true,
				},
			},
		}, nil)
	if withCustomImage {
		inspectCall.Times(2)
	} else {
		inspectCall.Times(3)
	}
	// Mock ContainerStop
	dockerClient.EXPECT().
		ContainerStop(gomock.Any(), validatorCtId, gomock.Any()).
		Return(nil)
	// Mock ContainerCreate
	dockerClient.EXPECT().
		ContainerCreate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), validatorImportCtName).
		Return(container.ContainerCreateCreatedBody{ID: validatorImportCtId}, nil).
		Times(1)
	// Mock ContainerStart
	dockerClient.EXPECT().
		ContainerStart(gomock.Any(), validatorImportCtId, gomock.Any()).
		Return(nil).
		Times(1)
	dockerClient.EXPECT().
		ContainerStart(gomock.Any(), services.ServiceCtValidator, gomock.Any()).
		Return(nil).
		Times(1)
	// Mock ContainerWait
	exitCh := make(chan container.ContainerWaitOKBody, 1)
	exitCh <- container.ContainerWaitOKBody{
		StatusCode: 0,
	}
	dockerClient.EXPECT().
		ContainerWait(gomock.Any(), validatorImportCtName, container.WaitConditionNextExit).
		Return(exitCh, make(chan error)).
		Times(1)
	// Mock ContainerRemove
	dockerClient.EXPECT().
		ContainerRemove(gomock.Any(), validatorImportCtId, types.ContainerRemoveOptions{}).
		Return(nil).
		Times(1)

	return dockerClient
}
