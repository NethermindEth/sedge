# Sedge
[![Go Report Card](https://goreportcard.com/badge/github.com/NethermindEth/sedge)](https://goreportcard.com/report/github.com/NethermindEth/sedge)

A one click setup tool for PoS network/chain validators. Currently, Sedge is designed primarily for solo stakers and testnet devs of the Beacon Chain and the Merge (Ethereum). Sedge generates docker-compose scripts for the entire on-premise validator setup based on the chosen client.

## 💥 How this all started?

As people who actively work on The Merge, we know how hard it is to set up an Ethereum validator:
- You need to procure at least three (compatible) nodes: an execution node (geth, nethermind, etc), a consensus node, and a validator node (lighthouse, prysm, etc)
- You then need to execute them, connect them, monitor them, and secure the validator keys (which includes staking 32 ETH).
- There may be several valid combinations of clients to choose for your setup, so you need to go through each of the client's docs, evaluate it, get instructions for it and test it. You also need to feel comfortable executing commands in the cli, know docker, and understand basics of networking. On top of this, there are many different settings you must read up on and consider for your client node.
- In the case of working with the Ethereum Mainnet, you are working with real money that can potentially be lost in the event of having downtime or being slashed. To avoid losing real value, you must be aware of and follow best practices on the validator setup, and correctly monitor your nodes.
- Have you heard of MEV? Flashbots is working on an MEV-Boost component which will take your validator to another level of awesomeness. You most likely want to always be running the latest version, but you also most likely don’t have the time to understand the MEV-Boost architecture in and out, or how to successfully implement it into your environment.
  
> Enter sedge

We want Sedge to take care of all of the above for you. With just a few clicks or steps, Sedge can create an entire ethereum staking architecture that supports client diversity and Ethereum's latest features, while being completely free and open source. We want Sedge to save you from making costly mistakes in this complex setup; along with hours or days of research, reading and testing. We want you to be able to stake easily with or without blockchain knowledge by giving you the tools to help this amazing community (and earn some good money of course 😉).

We want to share our knowledge in this topic and create something that allows everyone to easily and safely set up lots of diverse validators. 

We don't want to stop at Ethereum. We also want to help stakers of other PoS networks/chains, so if your favourite chain is not here, you are more than welcome to contribute!

## ⚡️ Quick start
### Dependencies
Sedge dependencies are `docker` with `docker compose` plugin, but if you don't have those installed, Sedge will show instructions to install them, or install them for you.

### Quick run
With `sedge cli` you can go through the entire workflow setup:
1. Check dependencies
2. Generate jwtsecret (not for mainnet and prater)
3. Generate a `docker-compose` script with randomized clients selection and `.env`
4. Execute the `docker-compose` script (only execution and consensus nodes will be executed by default)
5. Validator client will be executed automatically after execution and consensus nodes are synced.
  
Between steps 4 and 5 you can generate the validator(s) keystore folder using `sedge keys`. 

The entire process is interactive, although you can use the `-y` flag to run Sedge without prompts.

Check all the options and flags with `sedge cli --help`. More instructions or guides about sedge's features will come soon!

### Configuration file
When you run Sedge for the first time, it generates a `.sedge.yml` on your HOME directory. This file should look like this:

```yaml
dependencies:
  - docker

executionClients:
  - geth
  - nethermind

consensusClients:
  - lighthouse
  - lodestar
  - prysm
  - teku
 
validatorClients:
  - lighthouse
  - lodestar
  - prysm
  - teku

logs:
  logLevel: info
```

If you want to know what Sedge does in every step, then just read the logs. Sedge is logging every step. Some of the applied steps or commands are shown on the debug logs only. To see these logs you need to replace `info` for `debug` on the `logLevel` field.

You can modify the clients there to customize Sedge's random selection of clients. If you remove a client there, it won't be randomly selected. Be careful not to add an unsupported client, if this client is choosed you will face an error. Sedge knows very well which clients it supports.

You don't need to modify the `dependencies` field at all. If you alter it, Sedge may not work as expected.

## 🔥 What can you do with sedge today?

- Select an execution, consensus and validator node (manually or automatically) and generate a `docker-compose` script with production-tested configurations to run the setup you want.
- Generate the keystore folder using the [staking-deposit-cli](https://github.com/ethereum/staking-deposit-cli) tool with `sedge keys`
- Don't remember `docker-compose` commands or flags for your setup? Check docker logs of the running services with `sedge logs`, and shut them down with `sedge down`

> The setup is currently designed to start all three nodes required to run a validator (execution, consensus and validator node). Soon Sedge will let you directly connect to a public or remote node. The execution and consensus nodes will be executed first, and the validator node will be executed automatically after those nodes are synced, giving you time to prepare the keystore file and make the deposit for your staked ether.

If you are familiar with `docker`, `docker compose`, and the validator setup, then you can use Sedge to generate a base docker-compose script with the recommended settings, stop Sedge instead of letting it execute the script, and then edit the script as much as you want. Is a lot more easier than doing everything from scratch!

> Although Sedge supports several clients, is still on beta. Some settings may not work because -at least on the testnets- the clients are constantly evolving. Please let us know any issues you encounter!

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

### Ropsten

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Geth       | Lighthouse | Lighthouse |
| Nethermind | Lodestar   | Lodestar   |
|            | Prysm      | Prysm      |
|            | Teku       | Teku       |


### CL clients with Mev-Boost

| Client     | Mev-Boost | Networks   |
| ---------- | --------- | ---------- |
| Lighthouse | no*       | Ropsten    |
| Lodestar   | no        | -          |
| Prysm      | no        | -          |
| Teku       | yes       | Ropsten    |

> Settings for Lighthouse with mev-boost are quite ready, we are waiting for an official and stable Lighthouse docker image with mev-boost support
## Supported Linux flavours for dependency installation

| OS             | Versions                |
| -------------- | ----------------------- |
| Ubuntu         | 22.04,21.10,21.04,20.04 |
| Debian         | 11,10,9,8               |
| Fedora         | 35,34                   |
| CentOS         | 8                       |
| Arch           | -                       |
| Amazon Linux 2 | -                       |
| Alpine         | 3.15,3.14,3.14.3        |

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
- [ ] Integrate Goerli/Prater network
- [ ] Documentation with examples

### Version 0.3
- [ ] Enable use of public execution and consensus nodes
- [ ] Integrate Sepolia network
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

If you know of any good tricks for validator setup that other people could make good use of as well, please consider adding it to Sedge. Your efforts will be greatly appreciated by the community.

## ⚠️ License

Sedge is a Nethermind free and open-source software licensed under the [Apache 2.0 License](https://github.com/NethermindEth/sedge/blob/main/LICENSE).
