package cli_test

import (
	"bytes"
	"errors"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/NethermindEth/sedge/cli"
	"github.com/NethermindEth/sedge/cli/actions"
	"github.com/NethermindEth/sedge/internal/pkg/commands"
	"github.com/NethermindEth/sedge/internal/pkg/dependencies"
	sedge_mocks "github.com/NethermindEth/sedge/mocks"
	"github.com/NethermindEth/sedge/test"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	executionContainerName = "execution-client"
	consensusContainerName = "consensus-client"
	validatorContainerName = "validator-client"
	starknetContainerName  = "starknet-client"
	mevBoostContainerName  = "mev-boost"

	executionContainerIp    = "192.168.1.1"
	executionContainerImage = "nethermind/nethermind"
	consensusContainerIp    = "192.168.1.2"
	consensusContainerImage = "sigp/lighthouse"
	validatorContainerIp    = "192.168.1.3"
	validatorContainerImage = "consensys/teku"
	mevBoostContainerIp     = "192.168.1.4"
	mevBoostContainerImage  = "flashbots/mev-boost"
	starknetContainerIp     = "192.168.1.5"
	starknetContainerImage  = "nethermind/juno"
)

func buildInspectResults(
	containerName,
	containerImage,
	containerIp string,
) types.ContainerJSON {
	return types.ContainerJSON{
		ContainerJSONBase: &types.ContainerJSONBase{
			Name: containerName,
		},
		Config: &container.Config{
			Image: containerImage,
		},
		NetworkSettings: &types.NetworkSettings{
			Networks: map[string]*network.EndpointSettings{
				"sedge-network": { // FIXME: fix in case of network data renaming
					IPAddress: containerIp,
				},
			},
		},
	}
}

func TestShow(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	tcs := []struct {
		name  string
		setup func(*testing.T, *sedge_mocks.MockDependenciesManager, *sedge_mocks.MockSedgeActions) string
		err   string
	}{
		{
			name: "Valid flags and valid docker-compose",
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				generationPath = t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", "valid"), generationPath)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}

				containersData := actions.ContainersData{
					Containers: []actions.ContainerData{
						{
							Name:  executionContainerName,
							Image: executionContainerImage,
							Ip:    executionContainerIp,
						},
						{
							Name:  consensusContainerName,
							Image: consensusContainerImage,
							Ip:    consensusContainerIp,
						},
						{
							Name:  validatorContainerName,
							Image: validatorContainerImage,
							Ip:    validatorContainerIp,
						},
						{
							Name:  mevBoostContainerName,
							Image: mevBoostContainerImage,
							Ip:    mevBoostContainerIp,
						},
						{
							Name:  starknetContainerName,
							Image: starknetContainerImage,
							Ip:    starknetContainerIp,
						},
					},
				}
				options := actions.GetContainersDataOptions{
					DockerComposePath: filepath.Join(generationPath, "docker-compose.yml"),
				}

				gomock.InOrder(
					d.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil).Times(1),
					d.EXPECT().DockerEngineIsOn().Return(nil).Times(1),
					d.EXPECT().DockerComposeIsInstalled().Return(nil).Times(1),
					a.EXPECT().ValidateDockerComposeFile(filepath.Join(generationPath, "docker-compose.yml")).Return(nil).Times(1),
					a.EXPECT().GetContainersData(options).Return(containersData, nil).Times(1),
				)
				return
			},
		},
		{
			name: "Valid flags, bad docker-compose",
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				generationPath = t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", "bad_services"), generationPath)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				gomock.InOrder(
					d.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil).Times(1),
					d.EXPECT().DockerEngineIsOn().Return(nil).Times(1),
					d.EXPECT().DockerComposeIsInstalled().Return(nil).Times(1),
					a.EXPECT().ValidateDockerComposeFile(filepath.Join(generationPath, "docker-compose.yml")).Return(errors.New("error")).Times(1),
				)
				return
			},
			err: "error",
		},
		{
			name: "Flag pointing to non-existing docker-compose file",
			setup: func(t *testing.T, d *sedge_mocks.MockDependenciesManager, a *sedge_mocks.MockSedgeActions) (generationPath string) {
				generationPath = t.TempDir()
				err := test.PrepareTestCaseDir(filepath.Join("testdata", "run_tests", "no_compose"), generationPath)
				if err != nil {
					t.Fatalf("Error setting up test case: %v", err)
				}
				gomock.InOrder(
					d.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil).Times(1),
					d.EXPECT().DockerEngineIsOn().Return(nil).Times(1),
					d.EXPECT().DockerComposeIsInstalled().Return(nil).Times(1),
					a.EXPECT().ValidateDockerComposeFile(filepath.Join(generationPath, "docker-compose.yml")).Return(errors.New("error")).Times(1),
				)
				return
			},
			err: "error",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mActions := sedge_mocks.NewMockSedgeActions(ctrl)
			depsMgr := sedge_mocks.NewMockDependenciesManager(ctrl)
			generationPath := tc.setup(t, depsMgr, mActions)

			args := []string{"--path", generationPath}
			t.Logf("Running test with args: %v", args)

			var buff bytes.Buffer

			runner := &test.SimpleCMDRunner{
				SRunCMD: func(c commands.Command) (string, int, error) {
					if strings.Contains(c.Cmd, "docker compose") && strings.Contains(c.Cmd, "ps") && strings.Contains(c.Cmd, "--filter status=running") {
						return "", 0, nil
					}
					return "", 1, errors.New("runner error")
				},
				SRunBash: func(bs commands.ScriptFile) (string, error) {
					return "", nil
				},
			}

			showCmd := cli.ShowCmd(runner, mActions, depsMgr)
			showCmd.SetArgs(args)
			showCmd.SetOutput(&buff)
			err := showCmd.Execute()

			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
				assert.Contains(t, buff.String(), "help for show")
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
