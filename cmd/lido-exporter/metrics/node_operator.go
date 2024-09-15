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
package metrics

import (
	"context"
	"math/big"
	"time"

	"github.com/NethermindEth/sedge/internal/lido/contracts/csmodule"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
)

var (
	nodeOperatorIDGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_node_operator_id",
		Help: "Unique identifier for the node operator",
	}, []string{"node_operator_id", "network"})

	nodeOperatorManagerAddressGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_node_operator_manager_address",
		Help: "Address used to perform routine management operations regarding the CSM Node Operator",
	}, []string{"node_operator_id", "network"})

	nodeOperatorRewardAddressGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_node_operator_reward_address",
		Help: "Address that is the ultimate recipient of the rewards",
	}, []string{"node_operator_id", "network"})
)

func collectNodeOperatorInfo(ctx context.Context, network string, nodeOperatorID *big.Int, scrapeTime time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(scrapeTime):
			nodeOperator, err := csmodule.NodeOperatorInfo(network, nodeOperatorID)
			if err != nil {
				log.Errorf("Failed to get node operator info: %v", err)
				return
			}

			nodeOperatorIDGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(nodeOperatorID.Int64()))
			nodeOperatorManagerAddressGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(nodeOperator.ManagerAddress.Big().Int64()))
			nodeOperatorRewardAddressGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(nodeOperator.RewardAddress.Big().Int64()))

			log.Infof("Processed node operator data")
		}
	}
}
