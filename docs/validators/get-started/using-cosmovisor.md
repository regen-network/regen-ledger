# Using Cosmovisor

[Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/main/cosmovisor#cosmovisor) is a process manager for running application binaries. Using Cosmovisor is not required but recommended for node operators that would like to automate the upgrade process.

## Installation

::: tip Cosmovisor Versions
The following instructions use the latest version (`v1.4`), which now supports installation with `go install`.

For more information about each version, see the release notes:
- [Cosmovisor v1.0.0](https://github.com/cosmos/cosmos-sdk/releases/tag/cosmovisor%2Fv1.0.0)
- [Cosmovisor v1.1.0](https://github.com/cosmos/cosmos-sdk/releases/tag/cosmovisor%2Fv1.1.0)
- [Cosmovisor v1.2.0](https://github.com/cosmos/cosmos-sdk/releases/tag/cosmovisor%2Fv1.2.0)
- [Cosmovisor v1.3.0](https://github.com/cosmos/cosmos-sdk/releases/tag/cosmovisor%2Fv1.3.0)
- [Cosmovisor v1.4.0](https://github.com/cosmos/cosmos-sdk/releases/tag/tools%2Fcosmovisor%2Fv1.4.0)
:::

Use go install to install cosmovisor directly without building from source

```bash
go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.4.0
```

Check the version:

```bash
cosmovisor version
```

You should see the following (the errors following the version are expected if environment variables are not yet set, which will be set in the systemd file in the next section):

```bash
cosmovisor version: 1.4.0
```

### Building Cosmovisor from source
Clone the `cosmos-sdk` repository (if not already cloned):

```bash
git clone https://github.com/cosmos/cosmos-sdk
```

Change into the `cosmos-sdk` repository:

```bash
cd cosmos-sdk
```

Fetch the latest tags (if already cloned):

```bash
git fetch --all
```

Check out the tagged release:

```bash
git checkout tools/cosmovisor/v1.4.0
```

Install the cosmovisor binary:

```bash
go install cosmovisor
```

Check the version:

```bash
cosmovisor version
```

You should see the following (the errors following the version are expected if environment variables are not yet set, which will be set in the systemd file in the next section):

```bash
cosmovisor version: 1.4.0
```

## Cosmovisor Service

Create a `cosmovisor.service` systemd service file and make sure the environment variables are set to the desired values (the following example includes the default `cosmovisor` configuration settings with the exception of `DAEMON_NAME` and `DAEMON_HOME`):

::: tip Unsafe Skip Backups
The following recommended settings include `UNSAFE_SKIP_BACKUP=false` as a precaution but setting this to `true` will make the upgrade go much faster. Ideally backups are created ahead of time in order to limit the time it takes to bring validators back online.
:::

```bash
echo "[Unit]
Description=Cosmovisor daemon
After=network-online.target
[Service]
Environment="DAEMON_NAME=regen"
Environment="DAEMON_HOME=${HOME}/.regen"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment="DAEMON_POLL_INTERVAL=300ms"
Environment="DAEMON_DATA_BACKUP_DIR=${HOME}/.regen"
Environment="UNSAFE_SKIP_BACKUP=false"
Environment="DAEMON_PREUPGRADE_MAX_RETRIES=0"
User=${USER}
ExecStart=${GOBIN}/cosmovisor run start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```

For more information about the configuration options used in the example above, see [Command Line Arguments And Environment Variables](https://github.com/cosmos/cosmos-sdk/tree/main/cosmovisor#command-line-arguments-and-environment-variables).

Move the file to the systemd directory:

```bash
sudo mv cosmovisor.service /lib/systemd/system/cosmovisor.service
```

## Initialize Cosmovisor

Run the initialization command (if you built the `regen` binary from source, the path will be different, so make sure you provide the path to the `regen` binary that will be used as the starting binary):

```bash
DAEMON_HOME=~/.regen DAEMON_NAME=regen cosmovisor init $HOME/go/bin/regen
```

## Starting Cosmovisor

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

## Configuration Updates

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
