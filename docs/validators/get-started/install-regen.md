# Install Regen

The following instructions are for building the `regen` binary from source, also available at [Install Regen](../../ledger/get-started) alongside general information about the `regen` binary.

A significant difference here is that the genesis binary is used for each chain because a validator node needs to start from genesis (unless [Using State Sync](using-state-sync.md)).

The following instructions also assume that you have already completed [Initial Setup](README).

### Installation

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