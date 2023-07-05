# Install Regen

The following instructions are for building and installing the `regen` binary. In these instructions, we use the same version that was used to start both Regen Mainnet and Redwood Testnet. An alternative to syncing a node from genesis is [Using State Sync](using-state-sync.md) with the latest version.

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

Check out the genesis version:

```bash
git checkout v1.0.0
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
v1.0.0
```