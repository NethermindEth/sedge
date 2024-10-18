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
	"io"
	"strings"
	"testing"

	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"gotest.tools/v3/assert"
)

type ExpectedMetric struct {
	Name   string
	Type   string
	Help   string
	Labels []string
}

var expectedMetrics = map[string]ExpectedMetric{
	"csm_onchain_penalties_total": {
		Name:   "csm_onchain_penalties_total",
		Type:   "counter",
		Help:   "Number of penalties",
		Labels: []string{"node_operator_id", "penalty_type", "tx_hash"},
	},
	"csm_onchain_exit_requests_total": {
		Name:   "csm_onchain_exit_requests_total",
		Type:   "counter",
		Help:   "Number of validator exit requests",
		Labels: []string{"node_operator_id", "network", "tx_hash"},
	},
	"csm_node_operator_id": {
		Name:   "csm_node_operator_id",
		Type:   "gauge",
		Help:   "Unique identifier for the node operator",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_node_operator_manager_address": {
		Name:   "csm_node_operator_manager_address",
		Type:   "gauge",
		Help:   "Address used to perform routine management operations regarding the CSM Node Operator",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_node_operator_reward_address": {
		Name:   "csm_node_operator_reward_address",
		Type:   "gauge",
		Help:   "Address that is the ultimate recipient of the rewards",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_keys_stuck_validators_count": {
		Name:   "csm_keys_stuck_validators_count",
		Type:   "gauge",
		Help:   "Number of keys stuck in the system. A validator is considered to be stuck if it has not been exited timely following an exit signal from the protocol",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_keys_refunded_validators_count": {
		Name:   "csm_keys_refunded_validators_count",
		Type:   "gauge",
		Help:   "Number of keys refunded",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_keys_exited_validators_count": {
		Name:   "csm_keys_exited_validators_count",
		Type:   "gauge",
		Help:   "Number of keys exited",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_keys_deposited_validators_count": {
		Name:   "csm_keys_deposited_validators_count",
		Type:   "gauge",
		Help:   "Number of keys that already received deposits including withdrawn keys",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_keys_depositable_validators_count": {
		Name:   "csm_keys_depositable_validators_count",
		Type:   "gauge",
		Help:   "Number of keys eligible for deposits",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_added_keys_count": {
		Name:   "csm_added_keys_count",
		Type:   "gauge",
		Help:   "Number of keys added",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_withdrawn_keys_count": {
		Name:   "csm_withdrawn_keys_count",
		Type:   "gauge",
		Help:   "Number of keys withdrawn",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_vetted_keys_count": {
		Name:   "csm_vetted_keys_count",
		Type:   "gauge",
		Help:   "Number of keys vetted",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_enqueued_keys_count": {
		Name:   "csm_enqueued_keys_count",
		Type:   "gauge",
		Help:   "Number of keys in the deposit queue",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_bond_current": {
		Name:   "csm_bond_current",
		Type:   "gauge",
		Help:   "The current amount of bonded ETH",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_bond_required": {
		Name:   "csm_bond_required",
		Type:   "gauge",
		Help:   "The required amount of ETH to maintain",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_bond_excess": {
		Name:   "csm_bond_excess",
		Type:   "gauge",
		Help:   "The amount of excess bond over the required amount",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_bond_missed": {
		Name:   "csm_bond_missed",
		Type:   "gauge",
		Help:   "The amount of bond that is missing",
		Labels: []string{"node_operator_id", "network"},
	},
	"csm_non_claimed_rewards": {
		Name:   "csm_non_claimed_rewards",
		Type:   "gauge",
		Help:   "The amount of rewards available for claiming",
		Labels: []string{"node_operator_id", "network"},
	},
}

func parseMetrics(reader io.Reader) (map[string]*dto.MetricFamily, error) {
	parser := expfmt.TextParser{}
	return parser.TextToMetricFamilies(reader)
}

func validateMetrics(t *testing.T, parsedMetrics map[string]*dto.MetricFamily, expectedMetrics map[string]ExpectedMetric) {
	t.Helper()
	t.Logf("Validating metrics")
	foundMetrics := make(map[string]bool)
	for name, expected := range expectedMetrics {
		metricFamily, found := parsedMetrics[name]
		if !found {
			t.Logf("Metric %s not found in the response\n", name)
			continue
		}
		foundMetrics[name] = true

		// Check the metric type
		assert.Equal(t, metricFamily.GetType().String(), strings.ToUpper(expected.Type), "Metric %s has incorrect type. Expected: %s, Got: %s\n", name, strings.ToUpper(expected.Type), metricFamily.GetType())

		// Check the help string
		assert.Equal(t, metricFamily.GetHelp(), expected.Help, "Metric %s has incorrect help string. Expected: %s, Got: %s\n", name, expected.Help, metricFamily.GetHelp())

		// Check labels
		for _, metric := range metricFamily.Metric {
			labelNames := make(map[string]bool)
			for _, label := range metric.Label {
				labelNames[label.GetName()] = true
			}
			for _, expectedLabel := range expected.Labels {
				assert.Equal(t, labelNames[expectedLabel], true, "Metric %s is missing expected label: %s\n", name, expectedLabel)
			}
		}
	}
}
