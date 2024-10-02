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
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/NethermindEth/sedge/internal/utils"
)

// checkPrometheusServerUp checks that the prometheus server is up
func checkPrometheusServerUp(t *testing.T, port int) {
	t.Helper()
	response, err := utils.GetRequest(fmt.Sprintf("http://localhost:%d/metrics", port), 1*time.Second)
	require.NoError(t, err, "prometheus server should be up, it is not")
	require.Equal(t, http.StatusOK, response.StatusCode, "prometheus server should be up, but it is not")
}

// checkPrometheusServerDown checks that the prometheus server is down
func checkPrometheusServerDown(t *testing.T, port int) {
	t.Helper()
	response, err := utils.GetRequest(fmt.Sprintf("http://localhost:%d/metrics", port), 1*time.Second)
	require.Error(t, err, "prometheus server should be down, it is not")
	require.Nil(t, response, "prometheus server should be down, but it is not")
}

// checkMetrics checks that the metrics from the prometheus server are valid
// Should be called after checkPrometheusServerUp.
func checkMetrics(t *testing.T, port int) {
	t.Helper()
	// Get metrics from prometheus
	response, err := utils.GetRequest(fmt.Sprintf("http://localhost:%d/metrics", port), 1*time.Second)
	require.NoError(t, err, "request to prometheus server should succeed")
	require.Equal(t, http.StatusOK, response.StatusCode, "prometheus server should be up, but it is not")

	// Read response body into a string
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err, "should read response body")
	defer response.Body.Close()

	reader := strings.NewReader(string(body))
	metrics, err := parseMetrics(reader)
	require.NoError(t, err, "metrics should be parsed successfully")

	// Validate metrics
	validateMetrics(t, metrics, expectedMetrics)
}
