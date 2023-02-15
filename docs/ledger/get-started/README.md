# Install Regen

The `regen` binary serves as the node client and the application client. In other words, the `regen` binary can be used to both run a node and interact with it.

The `regen` binary can be installed using a pre-built package or by building and installing the binary from source. We recommend basic users install `regen` using the pre-built package for convenience, which does not require additional dependencies such as Git, Make, and Go.

## Pre-Built Package

### For Mac OS

Download the zip file:

```bash
curl -LO https://github.com/regen-network/regen-ledger/releases/download/v5.0.0/regen-ledger_5.0.0_darwin_amd64.zip
```

Verify the checksum:

```bash
sha256sum regen-ledger_5.0.0_darwin_amd64.zip
```

You should see the following:

```bash
26d07f258d489650f0dba059f6d3979f7550bf59514c0ee8f8912cca71bff1c6  regen-ledger_5.0.0_darwin_amd64.zip
```

Unzip the zip file:

```bash
unzip regen-ledger_5.0.0_darwin_amd64.zip
```

Move the binary to your local bin directory:

```bash
sudo mv regen-ledger_5.0.0_darwin_amd64/regen /usr/local/bin
```

Open a new terminal window and check if the installation was successful:

```bash
regen version
```

You should see the following:

```bash
v5.0.0
```

### For Linux Distributions

Download the zip file:

```bash
curl -LO https://github.com/regen-network/regen-ledger/releases/download/v5.0.0/regen-ledger_5.0.0_linux_amd64.zip
```

Verify the checksum:

```bash
sha256sum regen-ledger_5.0.0_linux_amd64.zip
```

You should see the following:

```bash
edbf5beaa769f971cf6b6f3c5e45cc3f38ac7d5b9dd005cddc821ccc37771155  regen-ledger_5.0.0_linux_amd64.zip
```

Unzip the zip file:

```bash
unzip regen-ledger_5.0.0_linux_amd64.zip
```

Move the binary to your local bin directory:

```bash
sudo mv regen-ledger_5.0.0_linux_amd64/regen /usr/local/bin
```

Check if the installation was successful:

```bash
regen version
```

You should see the following:

```bash
v5.0.0
```

### Other Packages

Additional packages and checksums are available under "Assets" on the [Release Page](https://github.com/regen-network/regen-ledger/releases/tag/v5.0.0).

## Building From Source

The following installation instructions are for installing the `regen` binary.

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
git checkout v5.0.0
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
v5.0.0
```
