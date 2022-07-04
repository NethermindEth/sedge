# 1click
[![Go Report Card](https://goreportcard.com/badge/github.com/NethermindEth/1click)](https://goreportcard.com/report/github.com/NethermindEth/1click)

A one click setup tool for PoS network/chain validators. Currently, `1click` is designed primarily for solo stakers and testnet devs of the Beacon Chain and the Merge (Ethereum). `1click` generates docker-compose scripts for the entire on-premise validator setup based on the chosen client.

## 💥 How this all started?

As people who actively work on The Merge, we know how hard it is to set up an Ethereum validator:
- You need to prepare procure at least three (compatible) nodes: an execution node (geth, nethermind, etc), a consensus node, and a validator node (lighthouse, prysm, etc)
- You then need to execute them, connect them, monitor them, and secure the validator keys (which includes staking 32 ETH).
- There may be several valid combinations of clients to choose for your setup, so you need to go through each of the client's docs, evaluate it, get instructions for it and test it. You also need to feel comfortable executing commands in the cli, know docker, and understand basics of networking. On top of this, there are many different settings you must read up on and consider for your client node.
- In the case of working with the Ethereum Mainnet, you are working with real money that can potentially be lost in the event of having downtime or being slashed. To avoid losing real value, you must be aware of and follow best practices on the validator setup, and correctly monitor your nodes.
- Have you heard of MEV? Flashbots is working on an MEV-Boost component which will take your validator to another level of awesomeness. You most likely want to always be running the latest version, but you also most likely don’t have the time to understand the MEV-Boost architecture in and out, or how to successfully implement it into your environment.
  
> Enter 1click

We want `1click` to take care of all of the above for you. With just a few clicks or steps, `1click` can create an entire ethereum staking architecture that supports client diversity and Ethereum's latest features, while being completely free and open source. We want `1click` to save you from making costly mistakes in this complex setup; along with hours or days of research, reading and testing. We want you to be able to stake easily with or without blockchain knowledge by giving you the tools to help this amazing community (and earn some good money of course 😉).

We want to share our knowledge in this topic and create something that allows everyone to easily and safely set up lots of diverse validators. 

We don't want to stop at Ethereum. We also want to help stakers of other PoS networks/chains, so if your favourite chain is not here, you are more than welcome to contribute!

## ⚡️ Quick start

### Installation (Only UNIX systems)

#### Using Go

If you have at least `go1.18.2` installed then this command will install the `1click` executable along with the library and its dependencies in your system:

```
go install github.com/NethermindEth/1click/cmd/1click@latest
```

The executable will be in your `$GOBIN` (`$GOPATH/bin`) 

#### Manual

Generate the executable manually (need Go installed):

```
git clone https://github.com/NethermindEth/1click.git
cd 1click
go build -o 1click cmd/1click/main.go
```

or if you have `make` installed:

```
git clone https://github.com/NethermindEth/1click.git
cd 1click
make compile
```

The executable will be in the `1click/build` folder

---
In case you want the binary in your main PATH (or you don't have `$GOBIN` in your PATH), please copy the executable to `/usr/local/bin`:

```
# Using go
sudo $GOPATH/bin/1click /usr/local/bin/
# Manually
sudo cp 1click/build/1click /usr/local/bin/
```

### Dependencies
`1click` dependencies are `docker` with `docker compose` plugin, but if you don't have those installed, `1click` will show instructions to install them, or install them for you.

### Quick run
With `1click cli` you can go through the entire workflow setup:
1. Check dependencies
2. Generate jwtsecret (not for mainnet and prater)
3. Generate a `docker-compose` script with randomized clients selection and `.env`
4. Execute the `docker-compose` script (only execution and consensus nodes will be executed by default)
  
## 🔥 What can you do with 1click today?

- Select an execution, consensus and validator node (manually or automatically) and generate a `docker-compose` script with production-tested configurations to run the setup you want.
- Generate the keystore folder using the [staking-deposit-cli](https://github.com/ethereum/staking-deposit-cli) tool with `1click keys`
- Don't remember `docker-compose` commands or flags for your setup? Check docker logs of the running services with `1click logs`, and shut them down with `1click down`

> The setup is currently designed to start all three nodes required to run a validator (execution, consensus and validator node). Soon `1click` will let you directly connect to a public or remote node. The execution and consensus nodes will be executed first, and the validator node will be executed automatically after those nodes are synced, giving you time to prepare the keystore file and make the deposit for your staked ether.

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
The following roadmap covers the main features and ideas we want to implement but doesn't cover everything we are planning for this tool. Stay in touch if you are interested, a lot of improvements are coming in the next two months.

### Version 0.1 (coming June-July 2022)
- [x] Generate `docker-compose` scripts and `.env` files for selected clients with a cli tool
- [x] Generate keystore folder with the cli
- [x] Test coverage (unit tests)
- [x] Integrate Kiln network
- [ ] Integrate MEV-Boost as an option

### Version 0.X
- [ ] Enable use of public execution and consensus nodes
- [ ] Include monitoring tool for alerting, tracking validator balance, and tracking sync progress and status of nodes
- [ ] TUI for guided and more interactive setup (better UX)
- [ ] Integrate Ropsten network
- [ ] Integrate Sepolia network
- [ ] Integrate Prater network
- [ ] Off-premise setup support
- [ ] Cross platform support and documentation
- [ ] Documentation with examples
- [ ] More tests!!!
- [ ] Integrate other PoS networks

### Version 1.0
Full Ethereum 2 support with MEV-Boost

## 💪 Want to contribute?
Please check our Contributing Guidelines, Code of Conduct and our issues. In case you want to report or suggest something (any help is welcome) please file an issue first so the main team is aware and it can be discussed.

If you know of any good tricks for validator setup that other people could make good use of as well, please consider adding it to `1click`. Your efforts will be greatly appreciated by the community.

## ⚠️ License

`1click` is a Nethermind free and open-source software licensed under the [Apache 2.0 License](https://github.com/NethermindEth/1click/blob/main/LICENSE).
