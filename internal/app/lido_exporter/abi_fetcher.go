package lido_exporter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

func fetchABI(contractAddress string) (string, error) {
	url := fmt.Sprintf("https://%s?module=contract&action=getabi&address=%s&apikey=%s", etherscanAPIURL, contractAddress, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if result["status"] == "1" {
		return result["result"].(string), nil
	} else {
		return "", fmt.Errorf("error fetching ABI: %s", result["message"])
	}
}

func readABIFromJson(address string) string {

	data, err := readJSON(JsonFilename)
	if err != nil {
		return ""
	} else {
		return data[address]
	}

}

func WriteABIToJson(abi string, address string) {

	if readABIFromJson(address) == abi {
		return
	}

	data, err := readJSON(JsonFilename)

	if err != nil {
		data = map[string]string{}
	}

	data[address] = abi

	err = writeJSON(JsonFilename, data)
}

func getEventSignatureHash(parsedABI abi.ABI, eventName string) common.Hash {

	// Get the event by name
	event, ok := parsedABI.Events[eventName]
	if !ok {
		fmt.Printf("Event %s not found in ABI\n", eventName)
		return common.Hash{}
	}

	// Construct the event signature string
	eventSignature := fmt.Sprintf("%s(", event.Name)
	for i, input := range event.Inputs {
		if i > 0 {
			eventSignature += ","
		}
		eventSignature += input.Type.String()
	}
	eventSignature += ")"

	fmt.Printf("event signature: %s\n", eventSignature)

	// Hash the event signature using keccak256 (SHA3-256)
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(eventSignature))
	var eventHash common.Hash
	hash.Sum(eventHash[:0])

	return eventHash
}
