Here goes the code for the lido-exporter

You can run it using the command
```
go run cmd\lido-exporter\main.go LExporter
```


# From the readme of the PoC:

The steps it goes through are:

1. Connect to RPC
2. Deduct the block range to monitor past events
3. Get the ABI of the contract:
    - If the ABI is in the JSON file "abi's.json," read the ABI from there
    - Otherwise, connect to the Etherscan API and fetch the ABI
    - If the contract is a proxy, call the `proxy__getImplementation` method to obtain the implementation address, and then retrieve the correct ABI
4. Using the `FilterQuery` data structure and the `FilterLogs` and `SubscribeFilterLogs` from the `ethereum` library, obtain the information from the required events
5. Process the information and expose some metrics

## What should you modify
- Introduce a valid `apiKey` for Etherscan as a constant named `apiKey`
- Modify the following constant variables depending on which event and node operator you wish to monitor: `rpcAddress`, `etherscanAPIURL`, `stringAddress`, `eventName`, `nodeOperatorId`. The three examples provided in this PoC are hardcoded in the constant declaration block at the start of the `main.go` file. They are in three different blocks separated by a blank line, for the events `Transfer` from the Tether smart contract, `ValidatorExitRequest`, and `ELRewardsStealingPenaltyReported`, respectively.
