---
sidebar_position: 1
---

# Introduction

**Sedge** is a one-click setup tool for PoS network/chain validators and nodes written entirely in the [Go programming language](https://golang.org/). **Sedge** takes care of the entire on-premise full node setup based on the chosen client, using generated docker-compose scripts based on the desired configuration. The following list shows some of the use cases for Sedge:

- **Staking**. You can use Sedge to set up and deploy an Ethereum and Gnosis full node, including a validator node and nodes with mev-boost. Sedge's on-premise setup design favors solo stakers but can also be integrated with other staking solutions.
- **Automated testing**. If you need to set up nodes very often to test your application or node, you can integrate Sedge to an automated solution, and make use of Sedge as a one-click tool with which you can quickly set up a full node in just a single step. This can be the case for protocol and dapps developers.
- **API and Checkpoint Sync endpoints**. Sedge allows you to set up a full node without a validator so that you can use this node either to expose the execution JSON-RPC API or the Beacon Chain HTTP API, or to expose a [Checkpoint Sync](concepts/checkpoint-sync.md) endpoint.

> Whether or not you choose to set up a node with a validator, and whether or not you are exposing the API, you are still contributing to the network, which is great!

:::caution Disclaimer

While Sedge assists in installing the validator, it is not designed to register or maintain it. Users are solely responsible for ensuring that they monitor and maintain the validator as required, so that they do not incur penalties and/or financial losses. This includes promptly updating the tool to ensure the latest stable releases of clients are used.

:::

:::caution Disclaimer

Users acknowledge that no warranty is being made of a successful installation. Sedge is a tool and ultimately depends on you to use it correctly and following all the best practice guidance, as found in the project's README and this documentation.

:::

## Why you should try Sedge

Sedge is a tool that focuses on ease of use and as such lowers the entry barrier for non technical users. With Sedge, you can run a node or a validator on Ethereum with little to no technical background and set it all up in less than 5 minutes! Don't know which client to choose? No problem! Let Sedge choose it for you.

If you have enough technical knowledge, you can also take advantage of Sedge and its features. You can generate the setup files, modify them according to your needs, and integrate Sedge with other solutions.

Sedge codebase design allow us to create many features in a way Sedge will be able to satisfy most of your needs: setup a single node or a full node, cross platform support, alerting mechanisms, support for other networks and much more. [Stay tuned](https://twitter.com/nethermindeth) to get the most of these features as soon as we ship them.

