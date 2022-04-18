# 1click
[![Go Report Card](https://goreportcard.com/badge/github.com/NethermindEth/1click)](https://goreportcard.com/report/github.com/NethermindEth/1click)

A one click setup tool for PoS networks/chains validators. Currently, 1click is designed mainly for solo stakers and testnets devs of Beacon Chain and The Merge (Ethereum). 1click generates docker-compose scripts for the whole validator setup given clients selection in an on-premise way.

## ⚡️ Quick start

### Installation (Only UNIX systems)

#### Using Go

If you have at least `go1.18.0` installed then this command will install the 1click executable along with the library and its dependencies:

```
go install github.com/NethermindEth/1click/cmd/1click@latest
```

The executable will be in `$GOBIN` (`$GOPATH/bin`) 

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
In case you want the binary in PATH (in case you don't have `$GOBIN` either in PATH), copy it to `/usr/local/bin`:

```
# Using go
sudo $GOPATH/bin/1click /usr/local/bin/
# Manual
sudo cp 1click/build/1click /usr/local/bin/
```

#### Download the binary

> This is temporary until the first release

Download directly the binary and put it in `/usr/local/bin`:

```
sudo curl -LJ -o /usr/local/bin/1click https://github.com/NethermindEth/1click/raw/feature/main/build/1click
sudo chmod +x /usr/local/bin/1click
```

### Dependencies
1click dependencies are `docker` and `docker-compose`, but it you don't have those installed, 1click will show instrucctions to install them, or install them for you.

### Quick run
With `1click cli` you can go through the entire workflow setup:
1. Check dependencies
2. Generate a `docker-compose` script with randomized clients selection and `.env`
3. Execute the `docker-compose` script (only execution and consensus nodes will be executed by default)
  
## 🔥 What can you do right now with this tool?

- Select an execution, consensus and validator node (manually or automatically) and generate a `docker-compose` script with production tested configurations to run the setup as you want.
- Generate the keystore folder using the [staking-deposit-cli](https://github.com/ethereum/staking-deposit-cli) tool with `1click keys`
- Don't remember `docker-compose` commands or flags? No problem, you can check the logs of the running services of the generated `docker-compose` script with `1click logs` and shutdown the services with `1click down`

> The setup is currently designed to start all the three nodes required to start a validator (execution, consensus and validator node). This will change soon and 1click will let you connect to a public or remote node, or to automatically start the validator node when the execution and consensus nodes in the setup are synced. Although you can do all of this after generating the docker-compose script 😉

## Supported networks and clients

### Mainnet

| Execution  | Consensus  | Validator  |
| ---------- | ---------- | ---------- |
| Geth       | Lighthouse | Lighthouse |
| Nethermind | Lodestar   | Lodestar   |
|            | Prysm      | Prysm      |
|            | Teku       | Teku       |

## ✅ Roadmap
The following roadmap covers the main features and ideas we want to implement but doesn't cover everything we are planning for this tool. Stay touched if you are interested, a lot of improvements are to come in the next two months.

### Version 0.1 (comming soon in May-June 2022)
- [x] Generate `docker-compose` scripts and `.env` files for selected clients with a cli tool
- [x] Generate keystore folder with the cli
- [ ] Test coverage (unit and integration tests)
- [ ] Monitoring tool for alerting, tracking validator balance, and tracking sync progress and status of nodes
- [ ] Integrate MEV-Boost as recommended setting  
- [ ] Use public execution and consensus nodes

### Version 0.X
- [ ] TUI for guided and more interactive setup (better UX)
- [ ] Integrate Kiln network
- [ ] Integrate Prater network
- [ ] Off-premise setup support
- [ ] Improve documentation
- [ ] Cross platform support
- [ ] More tests!!!

### Version 1.0
Full Ethereum 2 support with MEV-Boost

## 💪 Want to contribute?
Please check our Contributing Guidelines, Code of Conduct and our issues. In case you want to report or suggest something (any help welcomed) please file an issue first and the main team will reach you and discuss about it.

## ⚠️ License

`1click` is a Nethermind free and open-source software1 licensed under the [Apache 2.0 License](https://github.com/NethermindEth/1click/blob/main/LICENSE).
