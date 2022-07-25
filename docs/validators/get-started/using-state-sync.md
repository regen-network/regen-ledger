# Using State Sync

[Regen Mainnet](../../ledger/get-started/live-networks.md#regen-mainnet) and [Redwood Testnet](../../ledger/get-started/live-networks.md#redwood-testnet) also support [state sync](https://docs.cosmos.network/v0.44/architecture/adr-040-storage-and-smt-state-commitments.html#snapshots-for-storage-sync-and-state-versioning) which allows node operators to quickly spin up a node without downloading the existing chain data. It should be noted that not many nodes should be spun up on the network using this method as these nodes will be unable to propogate the historical data to other nodes.

Export your node moniker as an environment variable:

```bash
export MONIKER=<your-node-moniker>
```

Download and execute the state sync script:

*For Regen Mainnet:*

```bash 
curl -s -L https://raw.githubusercontent.com/regen-network/regen-ledger/master/scripts/statesync.bash | bash -s $MONIKER
```

*For Redwood Testnet:*

```bash 
curl -s -L https://raw.githubusercontent.com/regen-network/testnets/master/scripts/testnet-statesync.bash | bash -s $MONIKER
```
