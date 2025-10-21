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

	"github.com/NethermindEth/sedge/internal/lido/contracts"
	"github.com/NethermindEth/sedge/internal/lido/contracts/csexitpenalties"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
)

var penaltiesTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "csm_onchain_penalties_total",
	Help: "Number of penalties",
}, []string{"node_operator_id", "penalty_type", "tx_hash"})

// collectPenalties listens for Penalty events on the CSModule contract and updates the penalties counter.
// `filterOpts` is used to specify the start block for the event logs.
func collectPenalties(ctx context.Context, network string, nodeOperatorID *big.Int, client *ethclient.Client, filterOpts *bind.WatchOpts) {
	strikesPenaltyProcessedCh := make(chan *csexitpenalties.CsexitpenaltiesStrikesPenaltyProcessed)
	validatorExitDelayProcessedCh := make(chan *csexitpenalties.CsexitpenaltiesValidatorExitDelayProcessed)
	triggeredExitFeeRecordedCh := make(chan *csexitpenalties.CsexitpenaltiesTriggeredExitFeeRecorded)

	csexitpenaltiesFilterer, err := csexitpenalties.NewCsexitpenaltiesFilterer(common.HexToAddress(contracts.DeployedAddresses(contracts.CSExitPenalties)[network]), client)
	if err != nil {
		log.Errorf("Failed to create CSExitPenalties filterer: %v", err)
		return
	}

	// Subscribe to penalty events
	_, err = csexitpenaltiesFilterer.WatchStrikesPenaltyProcessed(filterOpts, strikesPenaltyProcessedCh, []*big.Int{nodeOperatorID})
	if err != nil {
		log.Errorf("Failed to watch StrikesPenaltyProcessed events: %v", err)
	}

	_, err = csexitpenaltiesFilterer.WatchValidatorExitDelayProcessed(filterOpts, validatorExitDelayProcessedCh, []*big.Int{nodeOperatorID})
	if err != nil {
		log.Errorf("Failed to watch ValidatorExitDelayProcessed events: %v", err)
	}

	_, err = csexitpenaltiesFilterer.WatchTriggeredExitFeeRecorded(filterOpts, triggeredExitFeeRecordedCh, []*big.Int{nodeOperatorID}, []*big.Int{big.NewInt(0), big.NewInt(1)})
	if err != nil {
		log.Errorf("Failed to watch TriggeredExitFeeRecorded events: %v", err)
	}

	for {
		select {
		case event := <-strikesPenaltyProcessedCh:
			penaltiesTotal.WithLabelValues(nodeOperatorID.String(), "strikes_penalty", event.Raw.TxHash.Hex()).Inc()
			log.Infof("Processed strikes penalty event")
		case event := <-validatorExitDelayProcessedCh:
			penaltiesTotal.WithLabelValues(nodeOperatorID.String(), "validator_exit_delay_penalty", event.Raw.TxHash.Hex()).Inc()
			log.Infof("Processed validator exit delay processed event")
		case event := <-triggeredExitFeeRecordedCh:
			penaltiesTotal.WithLabelValues(nodeOperatorID.String(), "triggered_exit_fee", event.Raw.TxHash.Hex()).Inc()
			log.Infof("Processed triggered exit fee recorded event")
		case <-ctx.Done():
			return
		}
	}
}
