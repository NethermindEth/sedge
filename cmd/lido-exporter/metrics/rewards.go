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

	rewards "github.com/NethermindEth/sedge/internal/lido/contracts/csfeedistributor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
)

var nonClaimedRewardsGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "csm_non_claimed_rewards",
	Help: "The amount of rewards available for claiming",
}, []string{"node_operator_id", "network"})

func collectRewardsInfo(ctx context.Context, network string, nodeOperatorID *big.Int, scrapeTime time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(scrapeTime):
			rewards, err := rewards.Rewards(network, nodeOperatorID)
			if err != nil {
				log.Errorf("Failed to get rewards: %v", err)
				return
			}

			nonClaimedRewardsGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(rewards.Int64()))
		}
	}
}
