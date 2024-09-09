package lido_exporter

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func processLog(vLog types.Log, contractABI abi.ABI, client *ethclient.Client) {
	eventSignature := vLog.Topics[0].Hex()
	event, err := contractABI.EventByID(common.HexToHash(eventSignature))

	if err != nil {
		log.Printf("Unknown event signature: %s", eventSignature)
		return
	}

	fmt.Printf("Event: %s\n", event.Name)

	unpackedData := make(map[string]interface{})
	err = contractABI.UnpackIntoMap(unpackedData, event.Name, vLog.Data)
	if err != nil {
		log.Fatalf("Failed to unpack non-indexed log data: %v", err)
	}

	for i, input := range event.Inputs {
		if input.Indexed {
			topicIndex := i + 1
			if topicIndex < len(vLog.Topics) {
				topicValue := vLog.Topics[topicIndex].Hex()
				switch input.Type.String() {
				case "address":
					unpackedData[input.Name] = common.HexToAddress(topicValue)
				case "uint256":
					unpackedData[input.Name] = new(big.Int).SetBytes(common.FromHex(topicValue))
				default:
					log.Printf("Unsupported indexed parameter type: %s", input.Type.String())
				}
			}
		}
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		log.Fatalf("Failed to get block details: %v", err)
	}

	blockTime := time.Unix(int64(block.Time()), 0)
	fmt.Printf("Block timestamp: %s\n", blockTime.Format(time.RFC3339))
	fmt.Printf("Block number: %d\n", vLog.BlockNumber)

	for name, value := range unpackedData {
		fmt.Printf("%s: %v\n", name, value)
	}
	fmt.Printf("\n")
}
