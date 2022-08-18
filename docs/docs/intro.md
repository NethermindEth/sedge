---
sidebar_position: 1
---

# Introduction

**Sedge** is a one click setup tool for PoS network/chain validators and nodes written entirely in the [Go programming language](https://golang.org/). Currently, **Sedge** is designed primarily for solo stakers and testnet devs of the Beacon Chain and the Merge (Ethereum), so these guides are focused only on Ethereum. **Sedge** takes care of the entire on-premise validator setup based on the chosen client; using generated docker-compose scripts based on the desired configuration.

:::caution Disclaimer

While Sedge assists in installing the validator, it is not designed to register or maintain it. Users are solely responsible for ensuring that they monitor and maintain the validator as required, so that they do not incur penalties and/or financial losses. This includes promptly updating the tool to ensure the latest stable releases of clients are used.

:::

:::caution Disclaimer

Users acknowledge that no warranty is being made of a successful installation. Sedge is a tool and ultimately depends on you to use it correctly and following all the best practice guidance, as found in the project's README and this documentation.

:::