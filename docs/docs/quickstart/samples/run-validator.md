---
sidebar_position: 1
id: run-validator
---

# Run a validator on Mainnet

This guide shows you how to setup and run a validator on Mainnet using random execution, consensus, and validator clients.

First make sure you have Sedge installed and in your PATH following the [installation guide](quickstart/install-guide.mdx).

:::tip

If you don't have Sedge in your PATH, just open your terminal on the folder which Sedge's executable / binary is and run `./sedge` instead of only `sedge`.

:::

Run the following command from your terminal to generate your setup on Mainnet (default network):

```
sedge generate full-node
```

As an alternative way, you can provide a fee recipient address:

```
sedge generate full-node --fee-recipient <your_address>
```

Set up your keys running the following command from your terminal:

```
sedge keys
```

Import the keys that you just generate in the command above using the following command:

```
sedge import-key
```

After that, you just need to run your setup with the following command:

```
sedge run
```

:::note

Your validator will not start right away, it will wait until the consensus get synced and 1 more epoch, to avoid slashing.

:::