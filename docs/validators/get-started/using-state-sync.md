# Using State Sync

[Regen Mainnet](../../ledger/get-started/regen-mainnet.md) supports [State Sync](https://docs.cosmos.network/v0.44/architecture/adr-040-storage-and-smt-state-commitments.html#snapshots-for-storage-sync-and-state-versioning), which allows node operators to quickly spin up a node without downloading the existing chain data.

State sync downloads a recent application state snapshot instead of replaying all blocks from genesis. This is significantly faster than syncing from genesis (minutes vs. hours/days) and is the recommended approach for new nodes that do not need full historical block data.

Although convenient, only a limited number of nodes should be spun up on the network using this method as these nodes will be unable to propagate historical data to other nodes. If you need to serve historical queries or run an archive node, you should sync from genesis instead.

Export a node moniker for the script to use:

```bash
export MONIKER=<your-node-moniker>
```

Download and execute the state sync script:

*For Regen Mainnet:*

```bash 
curl -s -L https://raw.githubusercontent.com/regen-network/regen-ledger/main/scripts/statesync.bash | bash -s $MONIKER
```
