# Validators

Regen Ledger is built on top of the [CometBFT](https://docs.cometbft.com/) (formerly Tendermint) Proof-of-Stake (PoS) consensus algorithm and requires validators to secure the network.

Validators run a full node and participate in consensus by broadcasting votes that contain cryptographic signatures signed by the validator's private key. Validators commit new blocks in the blockchain and receive revenue in exchange for their work.

Validators must also participate in governance by voting on proposals. Validators are weighted according to their total stake.

## Getting Started

Follow these guides to set up and run a validator node:

- [Initial Setup](get-started/README.md) - System requirements and initial configuration
- [Install Regen](get-started/install-regen.md) - Install the `regen` binary
- [Initialize Node](get-started/initialize-node.md) - Configure and connect to a network
- [Create a Validator](get-started/create-a-validator.md) - Register your node as a validator
- [Using Quickstart](get-started/using-quickstart.md) - Automated setup script
- [Using State Sync](get-started/using-state-sync.md) - Fast node synchronization
- [Using Cosmovisor](get-started/using-cosmovisor.md) - Automated binary upgrades

## Upgrade Guides

For chain upgrade instructions, see the [Upgrade Guides](upgrades/README.md).