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
	"github.com/NethermindEth/sedge/internal/lido/contracts/vebo"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	log "github.com/sirupsen/logrus"
)

var exitRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "csm_onchain_exit_requests_total",
	Help: "Number of validator exit requests",
}, []string{"node_operator_id", "network", "tx_hash"})

// collectExitRequests listens for ValidatorExitRequest events on the Vebo contract and updates the exit requests counter.
// `filterOpts` is used to specify the start block for the event logs.
func collectExitRequests(ctx context.Context, network string, nodeOperatorID *big.Int, client *ethclient.Client, filterOpts *bind.WatchOpts) {
	validatorExitRequestCh := make(chan *vebo.VeboValidatorExitRequest)

	// Get staking module ID
	stakingModuleID, err := contracts.StakingModuleID(network)
	if err != nil {
		log.Errorf("Failed to get staking module ID: %v", err)
		return
	}

	veboFilterer, err := vebo.NewVeboFilterer(common.HexToAddress(contracts.DeployedAddresses(contracts.Vebo)[network]), client)

	// Should ValidatorIndex be added?
	// Subscribe to ValidatorExitRequest events
	_, err = veboFilterer.WatchValidatorExitRequest(filterOpts, validatorExitRequestCh, []*big.Int{stakingModuleID}, []*big.Int{nodeOperatorID}, nil)
	if err != nil {
		log.Printf("Failed to watch ValidatorExitRequest events: %v", err)
	}

	for {
		select {
		case event := <-validatorExitRequestCh:
			exitRequestsTotal.WithLabelValues(nodeOperatorID.String(), network, event.Raw.TxHash.Hex()).Inc()
		case <-ctx.Done():
			return
		}
	}
}
