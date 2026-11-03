# Install Regen

The following instructions are for building and installing the latest version of the `regen` binary running on Regen Mainnet.

The latest version can only start from recent state, so use it with [Using State Sync](using-state-sync.md), which is the recommended approach for new validators.

::: warning Syncing From Genesis
A node that syncs from genesis must start with the version Regen Mainnet launched with (`v1.0.0`) and switch binaries at each upgrade height. See [Upgrade Overview](../upgrades/README.md) for the upgrade guides, in order.
:::

## Prerequisites

- [Initial Setup](README)

## Installation

Clone the `regen-ledger` repository:

```bash
git clone https://github.com/regen-network/regen-ledger
```

Change to the `regen-ledger` directory:

```bash
cd regen-ledger
```

Check out the latest version:

```bash
git checkout v7.3.1
```

Build and install the `regen` binary:


```bash
make install
```

Check to make sure the installation was successful:

```bash
regen version
```

You should see the following:

```bash
v7.3.1
```