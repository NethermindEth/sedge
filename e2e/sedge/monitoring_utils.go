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
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/NethermindEth/sedge/internal/pkg/services"
	"github.com/NethermindEth/sedge/internal/utils"
	"github.com/docker/docker/client"
)

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

func monitoringDotEnv(t *testing.T) map[string]string {
	t.Helper()
	dataDir, err := dataDirPath()
	if err != nil {
		t.Fatal(err)
	}
	envPath := filepath.Join(dataDir, "monitoring", ".env")
	raw, err := os.ReadFile(envPath)
	if err != nil {
		// Not installed yet, or cleaned already.
		return map[string]string{}
	}
	env := map[string]string{}
	for _, line := range strings.Split(string(raw), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		env[parts[0]] = parts[1]
	}
	return env
}

func prometheusBaseURL(t *testing.T) string {
	t.Helper()
	env := monitoringDotEnv(t)
	port := env["PROM_PORT"]
	if port == "" {
		port = "9090"
	}
	return fmt.Sprintf("http://localhost:%s", port)
}

func grafanaBaseURL(t *testing.T) string {
	t.Helper()
	env := monitoringDotEnv(t)
	port := env["GRAFANA_PORT"]
	if port == "" {
		port = "3000"
	}
	return fmt.Sprintf("http://localhost:%s", port)
}

func prometheusTargets(t *testing.T) (*PrometheusTargetsResponse, error) {
	response, err := utils.GetRequest(prometheusBaseURL(t)+"/api/v1/targets", time.Second)
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
