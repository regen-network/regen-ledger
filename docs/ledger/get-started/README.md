# Install Regen

The `regen` binary serves as the node client and the application client. In other words, the `regen` binary can be used to both run a node and interact with it.

The `regen` binary can be installed using a pre-built package or by building and installing the binary from source. We recommend basic users install `regen` using the pre-built package for convenience, which does not require additional dependencies such as Git, Make, and Go.

:::tip Experimental App Configuration
Users wanting to interact with [Hambach Testnet](live-networks.md#hambach-testnet) will need to build and install from source using the `EXPERIMENTAL` option. See [Building From Source](#building-from-source) for more information.
:::

## Pre-Built Package

### For Mac OS

Download the zip file:

```bash
curl -LO https://github.com/regen-network/regen-ledger/releases/download/v4.0.0/regen-ledger_4.0.0_darwin_amd64.zip
```

Verify the checksum:

```bash
sha256sum regen-ledger_4.0.0_darwin_amd64.zip
```

You should see the following:

```bash
19a3e2107d56ef727961f8204acd272fe416794e794697199be6ff11399f9930  regen-ledger_4.0.0_darwin_amd64.zip
```

Unzip the zip file:

```bash
unzip regen-ledger_4.0.0_darwin_amd64.zip
```

Move the binary to your local bin directory:

```bash
sudo mv regen-ledger_4.0.0_darwin_amd64/regen /usr/local/bin
```

Open a new terminal window and check if the installation was successful:

```bash
regen version
```

You should see the following:

```bash
v4.0.0
```

### For Linux Distributions

Download the zip file:

```bash
curl -LO https://github.com/regen-network/regen-ledger/releases/download/v4.0.0/regen-ledger_4.0.0_linux_amd64.zip
```

Verify the checksum:

```bash
sha256sum regen-ledger_4.0.0_linux_amd64.zip
```

You should see the following:

```bash
ca8e6020f2024f4cdb7722f917650a6334ad1f3068a8d14b0dba226bdd5532f0  regen-ledger_4.0.0_linux_amd64.zip
```

Unzip the zip file:

```bash
unzip regen-ledger_4.0.0_linux_amd64.zip
```

Move the binary to your local bin directory:

```bash
sudo mv regen-ledger_4.0.0_linux_amd64/regen /usr/local/bin
```

Check if the installation was successful:

```bash
regen version
```

You should see the following:

```bash
v4.0.0
```

### Other Packages

Additional packages and checksums are available under "Assets" on the [Release Page](https://github.com/regen-network/regen-ledger/releases/tag/v4.0.0).

## Building From Source

The following installation instructions include two options, one for installing the `regen` binary with the "stable app configuration" and another with the "experimental app configuration".

If you are looking to interact with features on [Hambach Testnet](live-networks.md#hambach-testnet) not available on [Regen Mainnet](live-networks.md#regen-mainnet) and [Redwood Testnet](live-networks.md#redwood-testnet), then you will want to use the "experimental app configuration", which means you will need to add the `EXPERIMENTAL` option to the `make install` command.

### Prerequisites

In order to build the `regen` binary from source, you'll need the following: 

- [Git](https://git-scm.com) `>=2` .
- [Make](https://www.gnu.org/software/make/) `>=4`
- [Go](https://golang.org/) `>=1.19`

### Go Environment

The `regen` binary is installed in `$(go env GOPATH)/bin`. Make sure `$(go env GOPATH)/bin` is in your `PATH` if not already there (e.g. `export PATH=$(go env GOPATH)/bin:$PATH`).

### Installation

Clone the `regen-ledger` repository:

```bash
git clone https://github.com/regen-network/regen-ledger
```

Change to the `regen-ledger` directory:

```bash
cd regen-ledger
```

Check out the latest stable version:

```bash
git checkout v4.0.0
```

Build and install the `regen` binary:

*For the stable app configuration (used on Regen Mainnet and Redwood Testnet):*

```bash
make install
```

*For the experimental app configuration (used on Hambach Testnet):*

```bash
EXPERIMENTAL=true make install
```

Check to make sure the installation was successful:

```bash
regen version
```

You should see the following:

```bash
v4.0.0
```
