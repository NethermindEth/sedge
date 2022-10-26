---
sidebar_position: 1
---

# Introduction

**Sedge** is a one click setup tool for PoS network/chain validators and nodes written entirely in the [Go programming language](https://golang.org/). **Sedge** takes care of the entire on-premise full node setup based on the chosen client; using generated docker-compose scripts based on the desired configuration. The following list shows some of the use cases for Sedge:

- **Staking**. You can use use Sedge to set up and deploy an Ethereum and Gnosis full node, including a validator node and with mev-boost. The on-premise setup design Sedge uses favors solo stakers but it can also be integrated in other staking solutions.
- **Automated testing**. If you need to set up nodes very often to test your application or node, you can integrate Sedge to an automated solution, taking advantage that Sedge can set up a full node very quickly in just one step. This can be the case for protocol and dapps developers.
- **API and Checkpoint Sync endpoints**. Sedge allows you to set up a full node without a validator so that you can use this node either to expose the execution JSON-RPC API or the Beacon Chain HTTP API, or to expose a [Checkpoint Sync](concepts/checkpoint-sync.md) endpoint.

> If you set up a node with or without a validator, exposing or not the API, you would be contributing to the network anyway, which is amazing.

:::caution Disclaimer

While Sedge assists in installing the validator, it is not designed to register or maintain it. Users are solely responsible for ensuring that they monitor and maintain the validator as required, so that they do not incur penalties and/or financial losses. This includes promptly updating the tool to ensure the latest stable releases of clients are used.

:::

:::caution Disclaimer

Users acknowledge that no warranty is being made of a successful installation. Sedge is a tool and ultimately depends on you to use it correctly and following all the best practice guidance, as found in the project's README and this documentation.

:::

## Why would you use Sedge

Sedge focus on ease of use and little to none background to run a node or validator. You don't need to be a highly technical person to run a node using Sedge. This creates a lower barrier for people to get involved in the ecosystem. However, if you are have enough technical knowledge you can take advantage of sedge capabilities, as you can generate the setup files and modify them according to your needs and also integrate Sedge on other solutions. Today you can setup a full node or a validator in Ethereum in just one step using Sedge and in less than 5 minutes! Don't know which client to choose? No problem! Let Sedge choose it for you. 

Sedge codebase design allow us to create many features in a way Sedge will be able to satisfy most of your needs: automated updates, cross platform support, alerting mechanisms, off-premise setup, setup a single node or a full node, improved UI/UX, other networks and much more. Stay tuned to get the most of these features as soon as we ship them.

