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
	"github.com/NethermindEth/sedge/internal/lido/contracts/csmodule"
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
	elRewardsStealingPenaltyReportedCh := make(chan *csmodule.CsmoduleELRewardsStealingPenaltyReported)
	initialSlashingSubmittedCh := make(chan *csmodule.CsmoduleInitialSlashingSubmitted)
	withdrawalSubmittedCh := make(chan *csmodule.CsmoduleWithdrawalSubmitted)

	csmoduleFilterer, err := csmodule.NewCsmoduleFilterer(common.HexToAddress(contracts.DeployedAddresses(contracts.CSModule)[network]), client)
	if err != nil {
		log.Errorf("Failed to create CSModule filterer: %v", err)
		return
	}

	// Subscribe to penalty events
	_, err = csmoduleFilterer.WatchELRewardsStealingPenaltyReported(filterOpts, elRewardsStealingPenaltyReportedCh, []*big.Int{nodeOperatorID})
	if err != nil {
		log.Errorf("Failed to watch ELRewardsStealingPenaltyReported events: %v", err)
	}

	_, err = csmoduleFilterer.WatchInitialSlashingSubmitted(filterOpts, initialSlashingSubmittedCh, []*big.Int{nodeOperatorID})
	if err != nil {
		log.Errorf("Failed to watch InitialSlashingSubmitted events: %v", err)
	}

	_, err = csmoduleFilterer.WatchWithdrawalSubmitted(filterOpts, withdrawalSubmittedCh, []*big.Int{nodeOperatorID})
	if err != nil {
		log.Errorf("Failed to watch WithdrawalSubmitted events: %v", err)
	}

	for {
		select {
		case event := <-elRewardsStealingPenaltyReportedCh:
			penaltiesTotal.WithLabelValues(nodeOperatorID.String(), "el_rewards_stealing", event.Raw.TxHash.Hex()).Inc()
			log.Infof("Processed EL rewards stealing penalty event")
		case event := <-initialSlashingSubmittedCh:
			penaltiesTotal.WithLabelValues(nodeOperatorID.String(), "initial_slashing", event.Raw.TxHash.Hex()).Inc()
			log.Infof("Processed initial slashing penalty event")
		case event := <-withdrawalSubmittedCh:
			// Amount is in Wei, but we want to count only less than 32 ETH
			if event.Amount.Cmp(new(big.Int).Mul(big.NewInt(32), big.NewInt(1e18))) < 0 {
				penaltiesTotal.WithLabelValues(nodeOperatorID.String(), "withdrawal", event.Raw.TxHash.Hex()).Inc()
			}
			log.Infof("Processed withdrawal penalty event")
		case <-ctx.Done():
			return
		}
	}
}
