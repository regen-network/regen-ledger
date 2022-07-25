# Install Cosmovisor

[Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor) is a process manager for running application binaries. Using Cosmovisor is not required but recommended for node operators that would like to automate the upgrade process.

## Installation

::: tip Cosmovisor Versions
It is not possible to install Cosmovisor v1.1 using `go install` and, when Cosmovisor v1.0 is installed using `go install`, the `version` command does not print the version. Until a new version is available, the following includes installation instructions for Cosmovisor v1.0 using `go install` (with a checksum) and Cosmovisor v1.1 building from source.
:::

### Using Go Install

Install the binary:

```bash
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0
```

Verify the checksum of the binary:

```bash
sha256sum $HOME/go/bin/cosmovisor
```

You should see the following:

```bash
33b850556c55ee17934373598b1284a39eb64975af79eeda55872446990b9254
```

### Building From Source

Clone the `cosmos-sdk` repository:

```bash
git clone https://github.com/cosmos/cosmos-sdk
```

Change into the `cosmos-sdk` repository:

```bash
cd cosmos-sdk
```

Check out the tagged release:

```bash
git checkout cosmovisor/v1.1.0
```

Build the cosmovisor binary:

```bash
make cosmovisor
```

Copy the built binary to your `GOBIN` directory:

```bash
cp ./cosmovisor/cosmovisor $HOME/go/bin
```

Verify the checksum of the binary:

```bash
sha256sum $HOME/go/bin/cosmovisor
```

You should see the following:

```bash
3f27c7feb1c093ddf9c919857b1cdf1e3a0654504c6acef68116deb325388361
```

## Configuration

### Initial Setup

Create a `cosmovisor.service` systemd service file and make sure the environment variables are set to the appropriate values (the following example includes the recommended settings):

```bash
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=regen"
Environment="DAEMON_HOME=${HOME}/.regen"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_DATA_BACKUP_DIR=${HOME}/.regen/backups"
Environment="UNSAFE_SKIP_BACKUP=false"
User=${USER}
ExecStart=${GOBIN}/cosmovisor start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```

`cosmovisor` can be configured to automatically download upgrade binaries. It is recommended that validators do not use the auto-download option and that the upgrade binary is prepared manually. If you would like to enable the auto-download option, update the following environment variable in the systemd configuration file:

```bash
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=true"
```

`cosmovisor` will automatically create a backup of the data directory at the time of the upgrade and before the migration. If you would like to disable the auto-backup, update the following environment variable in the systemd configuration file:

```bash
Environment="UNSAFE_SKIP_BACKUP=true"
```

Move the file to the systemd directory:

```bash
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
```

Create a directory for data backups:

```bash
mkdir -p ${HOME}/.regen/backups
```

Create a directory for the genesis binary:

```bash
mkdir -p $HOME/.regen/cosmovisor/genesis/bin
```

Copy the genesis binary which will be used when starting `cosmovisor` (assuming that you have already built the binary from source following the instructions in [Install Regen](./install-regen.md)):

```bash
cp path/to/binary/regen $HOME/.regen/cosmovisor/genesis/bin
```

Start `cosmovisor` to make sure everything is configured correctly:

```bash
sudo systemctl start cosmovisor
```

Check the status of the `cosmovisor` service:

```bash
sudo systemctl status cosmovisor
```

Enable `cosmovisor` to start automatically when the machine reboots:

```bash
sudo systemctl enable cosmovisor.service
```

### Updating Setup

When you make changes to the configuration, be sure to stop and start the `cosmovisor` service so that you are using the latest changes.

```bash
sudo systemctl stop cosmovisor
sudo systemctl daemon-reload
sudo systemctl start cosmovisor
```

Check the status of the `cosmovisor` service:

```bash
sudo systemctl status cosmovisor
```
