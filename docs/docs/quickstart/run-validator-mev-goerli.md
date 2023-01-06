---
sidebar_position: 6
id: run-teku-mev-boost
---

# Run a validator with mev-boost on Goerli

This guide shows you how to setup and run a validator using [Teku](https://github.com/Consensys/teku/) as consensus layer, with a random execution client, and mev-boost.

First make sure you have Sedge installed and in your PATH following the [installation guide](install-guide.mdx).

:::tip

If you don't have Sedge in your PATH, just open your Terminal on the folder which Sedge's executable / binary is and run `./sedge` instead of only `sedge`.

:::

Run the following command from your Terminal to setup a Teku consensus and validator nodes on Goerli with a random execution client:

```
sedge cli --network goerli -c teku
```

The `--network` flag allow you to choose the target network for the setup. To check out supported networks run `sedge networks`. Default network is mainnet.

The `-c/-v` flag is to select the desired consensus/validator client for the setup. If you only use one of those flags, then the same client pair will be used for consensus and validator nodes.

There is also a `-e` flag to select the execution client. The default behavior is to choose a randomized client, that's why if we skip the `-e` flag this time, a randomized execution client will be used.

mev-boost is a default setting as long as Sedge supports mev-boost for the selected client and network. If you don't want to use mev-boost in this case, then add the `--no-mev-boost` flag to the command. Check out the project's [README](https://github.com/NethermindEth/sedge) for more information on Sedge's mev-boost support.
