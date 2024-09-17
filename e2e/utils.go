package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/docker/docker/client"
)

func runSedge(t *testing.T, binaryPath string, args ...string) error {
	dataDir := filepath.Join(filepath.Dir(binaryPath), "sedge-data")
	return runCommand(t, binaryPath, append([]string{"--path", dataDir}, args...)...)
}

func runCommand(t *testing.T, path string, args ...string) error {
	_, err := runCommandOutput(t, path, args...)
	return err
}

func runCommandOutput(t *testing.T, path string, args ...string) ([]byte, error) {
	t.Helper()
	t.Logf("Binary path: %s", path)
	t.Logf("Running command: sedge %s", strings.Join(args, " "))
	out, err := exec.Command(path, args...).CombinedOutput()
	t.Logf("===== OUTPUT =====\n%s\n==================", out)
	return out, err
}

func dataDirPath() (string, error) {
	userDataHome := os.Getenv("XDG_DATA_HOME")
	if userDataHome == "" {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		userDataHome = filepath.Join(userHome, ".local", "share")
	}
	return filepath.Join(userDataHome, ".sedge"), nil
}

type Target struct {
	Labels Labels `json:"labels"`
	Health string `json:"health"`
}

type Labels map[string]string

type Data struct {
	ActiveTargets []Target `json:"activeTargets"`
}

type PrometheusTargetsResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

func prometheusTargets(t *testing.T) (*PrometheusTargetsResponse, error) {
	response, err := utils.GetRequest("http://localhost:9090/api/v1/targets", time.Second)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("prometheus targets status code should be 200")
	}
	var r PrometheusTargetsResponse
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func logAndPipeError(t *testing.T, prefix string, err error) error {
	t.Helper()
	if err != nil {
		t.Log(prefix, err)
	}
	return err
}

func getContainerIDByName(containerName string) (string, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return "", err
	}
	defer cli.Close()

	dockerServiceManager := services.NewDockerServiceManager(cli)

	containerID, err := dockerServiceManager.ContainerID(containerName)
	if err != nil {
		return "", err
	}

	return containerID, nil
}
