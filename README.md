# Sedge
[![Go Report Card](https://goreportcard.com/badge/github.com/NethermindEth/sedge)](https://goreportcard.com/report/github.com/NethermindEth/sedge)

A one click setup tool for PoS network/chain validators. Currently, `sedge` is designed primarily for solo stakers and testnet devs of the Beacon Chain and the Merge (Ethereum). `sedge` generates docker-compose scripts for the entire on-premise validator setup based on the chosen client.

## 💥 How this all started?

As people who actively work on The Merge, we know how hard it is to set up an Ethereum validator:
- You need to procure at least three (compatible) nodes: an execution node (geth, nethermind, etc), a consensus node, and a validator node (lighthouse, prysm, etc)
- You then need to execute them, connect them, monitor them, and secure the validator keys (which includes staking 32 ETH).
- There may be several valid combinations of clients to choose for your setup, so you need to go through each of the client's docs, evaluate it, get instructions for it and test it. You also need to feel comfortable executing commands in the cli, know docker, and understand basics of networking. On top of this, there are many different settings you must read up on and consider for your client node.
- In the case of working with the Ethereum Mainnet, you are working with real money that can potentially be lost in the event of having downtime or being slashed. To avoid losing real value, you must be aware of and follow best practices on the validator setup, and correctly monitor your nodes.
- Have you heard of MEV? Flashbots is working on an MEV-Boost component which will take your validator to another level of awesomeness. You most likely want to always be running the latest version, but you also most likely don’t have the time to understand the MEV-Boost architecture in and out, or how to successfully implement it into your environment.
  
> Enter sedge

We want `sedge` to take care of all of the above for you. With just a few clicks or steps, `sedge` can create an entire ethereum staking architecture that supports client diversity and Ethereum's latest features, while being completely free and open source. We want `sedge` to save you from making costly mistakes in this complex setup; along with hours or days of research, reading and testing. We want you to be able to stake easily with or without blockchain knowledge by giving you the tools to help this amazing community (and earn some good money of course 😉).

We want to share our knowledge in this topic and create something that allows everyone to easily and safely set up lots of diverse validators. 

We don't want to stop at Ethereum. We also want to help stakers of other PoS networks/chains, so if your favourite chain is not here, you are more than welcome to contribute!

## ⚡️ Quick start

### Installation (Only UNIX systems)

#### Using Go

If you have at least `go1.18.2` installed then this command will install the `sedge` executable along with the library and its dependencies in your system:

```
go install github.com/NethermindEth/sedge/cmd/sedge@latest
```

The executable will be in your `$GOBIN` (`$GOPATH/bin`) 

#### Manual

Generate the executable manually (need Go installed):

```
git clone https://github.com/NethermindEth/sedge.git
cd sedge
go build -o sedge cmd/sedge/main.go
```

or if you have `make` installed:

```
git clone https://github.com/NethermindEth/sedge.git
cd sedge
make compile
```

The executable will be in the `sedge/build` folder

---
In case you want the binary in your main PATH (or you don't have `$GOBIN` in your PATH), please copy the executable to `/usr/local/bin`:

```
# Using go install
sudo cp $GOPATH/bin/sedge /usr/local/bin/
# Manually
sudo cp sedge/build/sedge /usr/local/bin/
```

### Dependencies
`sedge` dependencies are `docker` with `docker compose` plugin, but if you don't have those installed, `sedge` will show instructions to install them, or install them for you.

### Quick run
With `sedge cli` you can go through the entire workflow setup:
1. Check dependencies
2. Generate jwtsecret (not for mainnet and prater)
3. Generate a `docker-compose` script with randomized clients selection and `.env`
4. Execute the `docker-compose` script (only execution and consensus nodes will be executed by default)
5. Validator client will be executed automatically after execution and consensus nodes are synced.
  
Between steps 4 and 5 you can generate the validator(s) keystore folder using `sedge keys`. 

Check all the options and flags with `sedge cli --help`. More instructions or guides about sedge's features will come soon!
## 🔥 What can you do with sedge today?

- Select an execution, consensus and validator node (manually or automatically) and generate a `docker-compose` script with production-tested configurations to run the setup you want.
- Generate the keystore folder using the [staking-deposit-cli](https://github.com/ethereum/staking-deposit-cli) tool with `sedge keys`
- Don't remember `docker-compose` commands or flags for your setup? Check docker logs of the running services with `sedge logs`, and shut them down with `sedge down`

> The setup is currently designed to start all three nodes required to run a validator (execution, consensus and validator node). Soon `sedge` will let you directly connect to a public or remote node. The execution and consensus nodes will be executed first, and the validator node will be executed automatically after those nodes are synced, giving you time to prepare the keystore file and make the deposit for your staked ether.

## Supported networks and clients

### Mainnet

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Geth       | Lighthouse | Lighthouse |
| Nethermind | Lodestar   | Lodestar   |
|            | Prysm      | Prysm      |
|            | Teku       | Teku       |

### Kiln

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Geth       | Lighthouse | Lighthouse |
| Nethermind | Lodestar   | Lodestar   |
|            | Prysm      | Prysm      |
|            | Teku       | Teku       |

## ✅ Roadmap
The following roadmap covers the main features and ideas we want to implement but doesn't cover everything we are planning for this tool. Stay in touch if you are interested, a lot of improvements are coming in the next two months. Please note that this Roadmap is continually changing until version 1.0.

### Version 0.1 (Actual)
- [x] Generate `docker-compose` scripts and `.env` files for selected clients with a cli tool
- [x] Generate keystore folder with the cli
- [x] Test coverage (unit tests)
- [x] Integrate Kiln network
- [x] Integrate MEV-Boost as an option
- [x] Integrate Ropsten network

### Version 0.2
- [ ] Set up and run only one node (execution/consensus/validator)
- [ ] Integrate Sepolia network
- [ ] Documentation with examples

### Version 0.3
- [ ] Enable use of public execution and consensus nodes
- [ ] Include monitoring tool for alerting, tracking validator balance, and tracking sync progress and status of nodes

### Version 0.4
- [ ] Integrate Gnosis network

### Version 0.X
- [ ] TUI for guided and more interactive setup (better UX)
- [ ] Integrate Prater network
- [ ] Off-premise setup support
- [ ] Cross platform support and documentation
- [ ] More tests!!!
- [ ] Integrate other PoS networks

### Version 1.0
Full Ethereum 2 support with MEV-Boost

## 💪 Want to contribute?
Please check our Contributing Guidelines, Code of Conduct and our issues. In case you want to report or suggest something (any help is welcome) please file an issue first so the main team is aware and it can be discussed.

If you know of any good tricks for validator setup that other people could make good use of as well, please consider adding it to `sedge`. Your efforts will be greatly appreciated by the community.

## ⚠️ License

`sedge` is a Nethermind free and open-source software licensed under the [Apache 2.0 License](https://github.com/NethermindEth/sedge/blob/main/LICENSE).
