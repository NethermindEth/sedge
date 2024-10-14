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

	bonds "github.com/NethermindEth/sedge/internal/lido/contracts/csaccounting"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	log "github.com/sirupsen/logrus"
)

var (
	bondCurrentGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_bond_current",
		Help: "The current amount of bonded ETH",
	}, []string{"node_operator_id", "network"})

	bondRequiredGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_bond_required",
		Help: "The required amount of ETH to maintain",
	}, []string{"node_operator_id", "network"})

	bondExcessGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_bond_excess",
		Help: "The amount of excess bond over the required amount",
	}, []string{"node_operator_id", "network"})

	bondMissedGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_bond_missed",
		Help: "The amount of bond that is missing",
	}, []string{"node_operator_id", "network"})
)

func collectBondInfo(ctx context.Context, network string, nodeOperatorID *big.Int, scrapeTime time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(scrapeTime):
			bondInfo, err := bonds.BondSummary(network, nodeOperatorID)
			if err != nil {
				log.Errorf("Failed to get bond summary: %v", err)
				return
			}

			bondCurrentGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(bondInfo.Current.Int64()))
			bondRequiredGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(bondInfo.Required.Int64()))
			bondExcessGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(bondInfo.Excess.Int64()))
			bondMissedGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(bondInfo.Missed.Int64()))
			log.Infof("Processed bond data")
		}
	}
}
