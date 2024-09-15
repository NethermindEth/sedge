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

func RunSedge(t *testing.T, binaryPath string, args ...string) error {
	dataDir := filepath.Join(filepath.Dir(binaryPath), "sedge-data")
	return RunCommand(t, binaryPath, "sedge", append([]string{"--path", dataDir}, args...)...)
}

func RunCommand(t *testing.T, path string, binaryName string, args ...string) error {
	_, _, err := runCommandOutput(t, path, binaryName, args...)
	return err
}

func RunCommandCMD(t *testing.T, path string, binaryName string, args ...string) *exec.Cmd {
	t.Helper()
	t.Logf("Running command: %s %s", binaryName, strings.Join(args, " "))
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start command: %s %s", binaryName, strings.Join(args, " "))
	}
	return cmd
}

func runCommandOutput(t *testing.T, path string, binaryName string, args ...string) ([]byte, *exec.Cmd, error) {
	t.Helper()
	t.Logf("Binary path: %s", path)
	t.Logf("Running command: %s %s", binaryName, strings.Join(args, " "))
	cmd := exec.Command(path, args...)
	out, err := cmd.CombinedOutput()
	t.Logf("===== OUTPUT =====\n%s\n==================", out)
	return out, cmd, err
}

func LogAndPipeError(t *testing.T, prefix string, err error) error {
	t.Helper()
	if err != nil {
		t.Log(prefix, err)
	}
	return err
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
