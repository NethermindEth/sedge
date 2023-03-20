---
sidebar_position: 5
id: run-validator
---

# Run a validator on Mainnet with one click

This guide shows you how to setup and run a validator on Mainnet using random execution, consensus, and validator clients.

First make sure you have Sedge installed and in your PATH following the [installation guide](quickstart/install-guide.mdx).

:::tip

If you don't have Sedge in your PATH, just open your Terminal on the folder which Sedge's executable / binary is and run `./sedge` instead of only `sedge`.

:::

Setup your keys running the following command from your Terminal:

```
sedge keys
```

Run the following command from your Terminal to generate your setup on Mainnet (default network):

```
sedge generate full-node
```

As an alternative way, you can provide a fee recipient address:

```
sedge generate full-node --fee-recipient <your_address>
```

After that, you just need to run your setup with the following command:

```
sedge run
```

:::note

Your validator will not start right away, it will wait until the consensus get synced and 1 more epoch, to avoid slashing.

:::