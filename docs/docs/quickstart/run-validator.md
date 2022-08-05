---
sidebar_position: 5
id: run-validator
---

# Run a validator on Mainnet with one click

This guide shows you how to setup and run a validator on Mainnet using random execution, consensus, and validator clients; configuring Sedge to run without any prompts or confirmations, and to install dependencies automatically if missing.

First make sure you have Sedge installed and in your PATH following the [installation guide](install-guide.mdx).

:::tip

If you don't have Sedge in your PATH, just open your Terminal on the folder which Sedge's executable / binary is and run `./sedge` instead of only `sedge`.

:::

Run the following command from your Terminal start the setup of a validator on Mainnet (default network):

```
sedge cli -y
```

As an alternative way, you can provide a fee recipient address:

```
sedge cli -y --fee-recipient <your_address>
```