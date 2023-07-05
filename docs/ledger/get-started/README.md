# Install Regen

The `regen` binary serves as the node client and the application client. In other words, the `regen` binary can be used to both run a node and interact with it.

If you have not already, check out [Command-Line Interface](../interfaces.md#command-line-interface) for an introduction.

## Build From Source

The following instructions are for building and installing the `regen` binary from its source code. This is the recommended way for most users. If you do not have Git, Make, and Go installed, and you would prefer not to install them, you can install the `regen` binary using a [pre-built package](#pre-built-package).

In the following examples we use the latest available version of Regen Ledger (`v5.1.2`), which is the same version used by node operators running [Regen Mainnet](regen-mainnet.md) and [Regen Redwood](redwood-testnet.md).

### Prerequisites

In order to build the `regen` binary from source, you'll need the following: 

- [Git](https://git-scm.com) `>=2` .
- [Make](https://www.gnu.org/software/make/) `>=4`
- [Go](https://golang.org/) `>=1.19`

*Note: If you are new to go and installing it for the first time, make sure you include the following `PATH` export in your bash profile: `export PATH=$(go env GOPATH)/bin:$PATH`.*

### Source Code

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
git checkout v5.1.2
```

### Build Only

The following command only builds the `regen` binary. After running the command, the built binary will be available in a `build` directory within the current directory.

Build the `regen` binary:

```bash
make build
```

Check to make sure the build was successful:

```bash
./build/regen version
```

You should see the following:

```bash
v5.1.2
```

Are you not seeing the above?

- Check the version of your source code using `git status` and try rerunning `git checkout` to see if there are any errors that might point to the problem.

### Build and Install

The following command builds and installs the `regen` binary using `go install` under the hood. After running the command, a separate built binary (separate from the binary in your `build` directory) will be available in your go `bin` directory and therefore available globally.

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
v5.1.2
```

Are you not seeing the above?

- Check the version of your source code using `git status` and try rerunning `git checkout` to see if there are any errors that might point to the problem.
- You may have previously installed the `regen` binary and placed it within `/usr/local/bin`. To use the version in your go `bin` directory, remove the `regen` binary from this location.

## Pre-Built Package

The following instructions are for installing the `regen` binary using a pre-built package. Packages for different operating systems are provided as release assets included with each release.

### For Mac OS

Download the zip file:

```bash
curl -LO https://github.com/regen-network/regen-ledger/releases/download/v5.1.0/regen-ledger_5.1.0_darwin_amd64.zip
```

Verify the checksum:

```bash
sha256sum regen-ledger_5.1.0_darwin_amd64.zip
```

You should see the following:

```bash
ab0d9e1a87681e2e3775c3b5e19dd8cd8af6111eb6719354ac5a6bad7e30e743  regen-ledger_5.1.0_darwin_amd64.zip
```

Unzip the zip file:

```bash
unzip regen-ledger_5.1.0_darwin_amd64.zip
```

Move the binary to your local bin directory:

```bash
sudo mv regen-ledger_5.1.0_darwin_amd64/regen /usr/local/bin
```

Open a new terminal window and check if the installation was successful:

```bash
regen version
```

You should see the following:

```bash
v5.1.0
```

### For Linux Distributions

Download the zip file:

```bash
curl -LO https://github.com/regen-network/regen-ledger/releases/download/v5.1.2/regen-ledger_5.1.2_linux_amd64.zip
```

Verify the checksum:

```bash
sha256sum regen-ledger_5.1.2_linux_amd64.zip
```

You should see the following:

```bash
c1419d8b3fcfefa8ad8b34c402269963a745f513c77ccda2ef96d575af70dd19  regen-ledger_5.1.2_linux_amd64.zip
```

Unzip the zip file:

```bash
unzip regen-ledger_5.1.2_linux_amd64.zip
```

Move the binary to your local bin directory:

```bash
sudo mv regen-ledger_5.1.2_linux_amd64/regen /usr/local/bin
```

Check if the installation was successful:

```bash
regen version
```

You should see the following:

```bash
v5.1.2
```

### Additional Packages

Additional packages and checksums are available under "Assets" on the [Release Page](https://github.com/regen-network/regen-ledger/releases/tag/v5.1.2).
