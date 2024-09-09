package lido_exporter

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus"
)

func MonitorEvents(client *ethclient.Client, contractAddress string) {
	currentBlock, fromBlock, toBlock := getBlockRange(client)

	fmt.Printf("Deducting block range...\n")
	fmt.Printf("From block: %d\n", fromBlock)
	fmt.Printf("To block: %d\n", toBlock)
	fmt.Printf("Current block: %d\n\n", currentBlock)

	fmt.Printf("Getting contract ABI...\n")
	contractABI, _ := getContractABI(client, contractAddress)

	// Fetch and process past logs
	pastLogs := fetchPastLogs(client, fromBlock, toBlock, contractAddress, getTopics(contractABI, eventName, nodeOperatorId))
	processPastLogs(pastLogs, contractABI, client)

	// Subscribe to new logs
	subscribeToNewLogs(client, currentBlock, contractAddress, contractABI, getTopics(contractABI, eventName, nodeOperatorId))
}

func getBlockRange(client *ethclient.Client) (uint64, uint64, uint64) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get the latest block: %v", err)
	}

	currentBlock := header.Number.Uint64()
	blockTime := 13 * time.Second
	blocksPerDay := uint64(24 * time.Hour / blockTime)
	fromBlock := currentBlock - (blocksPerDay * numberOfDaysToCheck)
	toBlock := currentBlock

	fromBlock = 1945061 // for testing ExitRequests
	fromBlock = 1978983 //for testing Penalties
	toBlock = fromBlock + 2

	return currentBlock, fromBlock, toBlock
}

func getContractABI(client *ethclient.Client, contractAddressAsString string) (abi.ABI, string) {

	contractABI := readABIFromJson(contractAddressAsString)

	if contractABI == "" {
		var err error
		contractABI, err = fetchABI(contractAddressAsString)
		if err != nil {
			log.Fatalf("Failed to fetch ABI: %v", err)
		}
		WriteABIToJson(contractABI, contractAddressAsString)
	}

	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	contractAddress := common.HexToAddress(contractAddressAsString)
	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: parsedABI.Methods["proxy__getImplementation"].ID,
	}, nil)

	if err == nil {
		contractAddressAsString = common.BytesToAddress(result[:]).Hex()
		contractABI = readABIFromJson(contractAddressAsString)

		if contractABI == "" {
			var err error
			contractABI, err = fetchABI(contractAddressAsString)
			if err != nil {
				log.Fatalf("Failed to fetch ABI: %v", err)
			}
			WriteABIToJson(contractABI, contractAddressAsString)
		}

		parsedABI, err = abi.JSON(strings.NewReader(contractABI))
		if err != nil {
			log.Fatalf("Failed to parse ABI: %v", err)
		}
	}

	return parsedABI, contractAddressAsString
}

func fetchPastLogs(client *ethclient.Client, fromBlock, toBlock uint64, contractAddress string, topics [][]common.Hash) []types.Log {
	fmt.Printf("Fetching past logs...\n")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics:    topics,
	}

	pastLogs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	return pastLogs
}

func processPastLogs(pastLogs []types.Log, contractABI abi.ABI, client *ethclient.Client) {
	fmt.Printf("Processing past logs (%d logs found):\n", len(pastLogs))
	for _, vLog := range pastLogs {
		processLog(vLog, contractABI, client)
		eventCounter.With(prometheus.Labels{"event_type": "past_events"}).Inc()
	}
}

func subscribeToNewLogs(client *ethclient.Client, currentBlock uint64, contractAddress string, contractABI abi.ABI, topics [][]common.Hash) {
	fmt.Printf("Subscribing to new logs...\n")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(currentBlock)),
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics:    topics,
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	fmt.Println("Listening for new logs...")
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Error: %v", err)
		case vLog := <-logs:
			processLog(vLog, contractABI, client)
			eventCounter.With(prometheus.Labels{"event_type": "incoming_events"}).Inc()
		}
	}
}

func intToHash(value int64) common.Hash {
	// Convert the integer to a big.Int
	bigIntValue := new(big.Int).SetInt64(value)

	// Convert big.Int to a byte slice
	byteSlice := bigIntValue.Bytes()

	// Create a common.Hash and fill it with the byte slice
	var hash common.Hash

	// Copy the byte slice into the hash, ensuring it is right-padded with zeros if necessary
	copy(hash[32-len(byteSlice):], byteSlice)

	return hash
}

func getTopics(contractABI abi.ABI, eventName string, operatorId int64) [][]common.Hash {
	eventSignatureHash := getEventSignatureHash(contractABI, eventName)
	topics := [][]common.Hash{{eventSignatureHash}}

	var elementsToAdd = [][]common.Hash{}
	if eventName == "ValidatorExitRequest" {
		elementsToAdd = [][]common.Hash{{}, {intToHash(operatorId)}}
	}

	if eventName == "ELRewardsStealingPenaltyReported " {
		elementsToAdd = [][]common.Hash{{intToHash(operatorId)}}
	}

	topics = append(topics, elementsToAdd...)
	return topics
}
