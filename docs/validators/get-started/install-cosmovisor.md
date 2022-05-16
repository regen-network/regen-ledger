# Install Cosmovisor

[Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor) is a process manager for running application binaries. Using Cosmovisor is not required but recommended for node operators that would like to automate the upgrade process.

Install the binary:

```bash
go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v1.0
```

Check to ensure the installation was successful:

```bash
cosmovisor version
```
