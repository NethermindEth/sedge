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
package cli

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
	"time"

	bonds "github.com/NethermindEth/sedge/internal/lido/contracts/csaccounting"
	rewards "github.com/NethermindEth/sedge/internal/lido/contracts/csfeedistributor"
	"github.com/NethermindEth/sedge/internal/lido/contracts/csmodule"
	"github.com/NethermindEth/sedge/internal/ui"
	"github.com/cheggaaa/pb/v3"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type LidoData struct {
	NodeID   *big.Int
	NodeInfo csmodule.NodeOperator
	Keys     csmodule.Keys
	BondInfo bonds.BondInfo
	Rewards  *big.Int
}

var (
	rewardAddress    string
	networkName      string
	longDescriptions bool
)

const (
	nodeOpInfo  = `Node Operator Info`
	keysInfo    = `Keys`
	queueInfo   = `Queue`
	bondInfo    = `Bond`
	rewardsInfo = `Rewards`
)

func LidoStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lido-status [flags] [args]",
		Short: "Display status and information of Lido Node Operator",
		Long: `This command retrieves and displays the status and detailed information of Lido Node Operators. The information includes:
		- Node Operator ID.
		- Keys and queue information: available for deposit (in the queue), stuck, refunded, exited, deposited.
		- Bond and rewards information: total amount, amounts lower and higher than required, non-claimed rewards.
		- Alerts for penalties and exit requests.
		
		Valid args: reward address of Node Operator (rewards recipient)`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				if cobra.ExactArgs(1)(cmd, args) != nil {
					return errors.New("requires one argument")
				}
				rewardAddress = args[0]
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runListLidoStatusCmd(cmd, args); err != nil {
				return err
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&networkName, "network", "n", "holesky", "Target network. e.g. holesky, mainnet etc.")
	cmd.Flags().BoolVar(&longDescriptions, "l", false, "Show detailed descriptions for each value")
	cmd.Flags().SortFlags = false
	return cmd
}

func runListLidoStatusCmd(cmd *cobra.Command, args []string) error {
	log.Infof("Retrieving Lido Node Operator Info\n")
	nodeData, err := nodeData()
	if err != nil {
		return err
	}
	dataMap := buildLidoData(nodeData)

	// Extract headers and sort by weight
	headers := make([]string, 0, len(dataMap))
	for header := range dataMap {
		headers = append(headers, header)
	}

	// Sort headers by their weights
	sort.SliceStable(headers, func(i, j int) bool {
		return dataMap[headers[i]].weight < dataMap[headers[j]].weight
	})

	log.Infof("Listing Lido Node Operator Info\n")
	for _, header := range headers {
		ui.WriteLidoStatusTable(cmd.OutOrStdout(), dataMap[header].data, header)
	}
	return nil
}

// Get the data for the Node Operator
func nodeData() (*LidoData, error) {
	nodeData := &LidoData{}

	progressBar := pb.StartNew(5)
	defer progressBar.Finish()
	progressBar.SetCurrent(0)
	time.Sleep(1 * time.Second) // Simulate work
	progressBar.Increment()

	nodeID, err := csmodule.NodeID(networkName, rewardAddress)
	if err != nil {
		return nodeData, err
	}
	progressBar.Increment()

	nodeInfo, err := csmodule.NodeOperatorInfo(networkName, nodeID)
	if err != nil {
		return nodeData, err
	}
	progressBar.Increment()

	keys, err := csmodule.KeysStatus(networkName, nodeID)
	if err != nil {
		return nodeData, err
	}
	progressBar.Increment()

	bond, err := bonds.BondSummary(networkName, nodeID)
	if err != nil {
		return nodeData, err
	}

	reward, err := rewards.Rewards(networkName, nodeID)
	if err != nil {
		return nodeData, err
	}
	progressBar.Increment()

	nodeData.NodeID = nodeID
	nodeData.NodeInfo = nodeInfo
	nodeData.Keys = keys
	nodeData.BondInfo = bond
	nodeData.Rewards = reward

	return nodeData, nil
}

// Structure the data to be displayed
func buildLidoData(node *LidoData) map[string]struct {
	data   []string
	weight int
} {
	var nodeOpDetailed, keysDetailed, queueDetailed, bondDetailed, rewardsDetailed string
	rewardAddressLink := fmt.Sprintf(`https://etherscan.io/address/%s`, node.NodeInfo.RewardAddress)

	detailedDescriptions := map[string]string{
		nodeOpInfo: `
## Description 
- Node Operator ID: Unique identifier for the node operator.
- Reward Address: Address tha is the ultimate recipient of the rewards
- Manager Address: Address used to perform routine management operations regarding the CSM Node Operator.`,

		keysInfo: `
## Description 
- Stuck Keys Count: Number of keys stuck in the system. A validator is considered to be "stuck" if it has not been exited timely following an exit signal from the protocol.
- Refunded Keys Count: Number of keys that were refunded.
- Exited Keys Count: Number of keys that have exited.
- Deposited Keys Count: Number of keys currently deposited.
- Depositable Keys Count: Number of keys eligible for deposits.`,

		queueInfo: `
## Description 
- Keys in the deposit queue: Number of the depositable keys that are in the deposit queue.`,

		bondInfo: `
## Description 
- Bond : a security collateral that Node Operators must submit before uploading validator keys into CSM. It covers possible losses caused by inappropriate actions on the Node Operator's side.
- Current Bond: The current amount of bonded ETH.
- Required Bond: The required amount of ETH to maintain.
- Excess Bond: The amount of excess bond over the required amount.
- Missed Bond: The amount of bond that is missing.`,

		rewardsInfo: `
## Description 
- Non-claimed Rewards: The amount of rewards available for claiming.`,
	}

	if longDescriptions {
		nodeOpDetailed = detailedDescriptions[nodeOpInfo]
		keysDetailed = detailedDescriptions[keysInfo]
		queueDetailed = detailedDescriptions[queueInfo]
		bondDetailed = detailedDescriptions[bondInfo]
		rewardsDetailed = detailedDescriptions[rewardsInfo]
	}

	data := map[string]struct {
		data   []string
		weight int
	}{
		nodeOpInfo: {
			data: []string{
				fmt.Sprintf(`- **Node Operator ID:** %s`, node.NodeID.String()),
				fmt.Sprintf(`- **Reward Address:** %s`, node.NodeInfo.RewardAddress.String()),
				fmt.Sprintf(`- **Manager Address:** %s`, node.NodeInfo.ManagerAddress.String()),
				fmt.Sprintf(`- [Reward Address Link on etherscan](%s)`, rewardAddressLink),
				nodeOpDetailed,
			},
			weight: 1,
		},
		keysInfo: {
			data: []string{
				fmt.Sprintf(`- **Stuck Keys Count:** %s`, node.Keys.StuckValidatorsCount.String()),
				fmt.Sprintf(`- **Refunded Keys Count:** %s`, node.Keys.RefundedValidatorsCount.String()),
				fmt.Sprintf(`- **Exited Keys Count:** %s`, node.Keys.ExitedValidators.String()),
				fmt.Sprintf(`- **Deposited Keys Count:** %s`, node.Keys.DepositedValidators.String()),
				fmt.Sprintf(`- **Depositable Keys Count:** %s`, node.Keys.DepositableValidatorsCount.String()),
				keysDetailed,
			},
			weight: 2,
		},
		queueInfo: {
			data: []string{
				fmt.Sprintf(`- **Keys in the deposit queue:** %d`, node.NodeInfo.EnqueuedCount),
				queueDetailed,
			},
			weight: 3,
		},
		bondInfo: {
			data: []string{
				fmt.Sprintf(`- **Current Bond:** %s`, weiToEth(node.BondInfo.Current).String()),
				fmt.Sprintf(`- **Required Bond:** %s`, weiToEth(node.BondInfo.Required).String()),
				fmt.Sprintf(`- **Excess Bond:** %s`, weiToEth(node.BondInfo.Excess).String()),
				fmt.Sprintf(`- **Missed Bond:** %s`, weiToEth(node.BondInfo.Missed).String()),
				bondDetailed,
			},
			weight: 4,
		},
		rewardsInfo: {
			data: []string{
				fmt.Sprintf(`- **Non-claimed Rewards:** %s`, weiToEth(node.Rewards).String()),
				rewardsDetailed,
			},
			weight: 5,
		},
	}

	return data
}

// Convert Wei to Ether
func weiToEth(wei *big.Int) decimal.Decimal {
	weiToEther := decimal.NewFromBigInt(big.NewInt(1e18), 0)
	weiDecimal := decimal.NewFromBigInt(wei, 0)
	return weiDecimal.Div(weiToEther)
}
