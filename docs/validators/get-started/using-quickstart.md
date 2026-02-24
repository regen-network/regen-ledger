# Using Quickstart

We recommend validators become familiar with the manual setup process, but we do have a handy quickstart script available for spinning up a full node from scratch.

The quickstart script automates the following steps:

1. Installs system dependencies
2. Downloads and installs the `regen` binary
3. Initializes the node configuration and genesis file
4. Configures persistent peers
5. Starts the node

This covers the steps in [Initial Setup](README.md), [Install Regen](install-regen.md), and [Initialize Node](initialize-node.md).

*For Regen Mainnet:*

```bash
bash <(curl -s https://raw.githubusercontent.com/regen-network/mainnet/blob/main/scripts/mainnet-val-setup.sh)
```

After the script completes, the node will begin syncing with the network. You can then proceed to [Create a Validator](create-a-validator.md) once the node is fully synced.
