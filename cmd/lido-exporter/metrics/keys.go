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
	keysExitedValidatorsCountGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_keys_exited_validators_count",
		Help: "Number of keys exited",
	}, []string{"node_operator_id", "network"})

	keysDepositedValidatorsCountGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_keys_deposited_validators_count",
		Help: "Number of keys that already received deposits including withdrawn keys",
	}, []string{"node_operator_id", "network"})

	keysDepositableValidatorsCountGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_keys_depositable_validators_count",
		Help: "Number of keys eligible for deposits",
	}, []string{"node_operator_id", "network"})

	addedKeysCountGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_added_keys_count",
		Help: "Number of keys added",
	}, []string{"node_operator_id", "network"})

	withdrawnKeysCountGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_withdrawn_keys_count",
		Help: "Number of keys withdrawn",
	}, []string{"node_operator_id", "network"})

	vettedKeysCountGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_vetted_keys_count",
		Help: "Number of keys vetted",
	}, []string{"node_operator_id", "network"})

	enqueuedKeysCountGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "csm_enqueued_keys_count",
		Help: "Number of keys in the deposit queue",
	}, []string{"node_operator_id", "network"})
)

func collectKeysInfo(ctx context.Context, network string, nodeOperatorID *big.Int, scrapeTime time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(scrapeTime):
			keysStatus, err := csmodule.KeysStatus(network, nodeOperatorID)
			if err != nil {
				log.Errorf("Failed to get keys status: %v", err)
				return
			}

			keysExitedValidatorsCountGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(keysStatus.ExitedValidators.Int64()))
			keysDepositedValidatorsCountGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(keysStatus.DepositedValidators.Int64()))
			keysDepositableValidatorsCountGauge.WithLabelValues(nodeOperatorID.String(), network).Set(float64(keysStatus.DepositableValidatorsCount.Int64()))

			log.Info("Processed keys data")
		}
	}
}
