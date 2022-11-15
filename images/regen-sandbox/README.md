# Regen Ledger Sandbox

This container (and corresponding setup scripts) is intended to aid with end-to-end testing of client libraries, applications, manual local network tests, and other scenarios where you want to be able to quickly bootstrap a single node regen network (targeting any version or commit hash of regen-ledger), and populate the network with some basic data for your end-to-end testing or exploratory purposes.


Build locally:
```sh
docker build . -f images/regen-sandbox/Dockerfile -t regen-sandbox
```

Optionally export your own testing mnemonic:
```sh
export REGEN_MNEMONIC="YOUR TESTING MNEMONIC"
```

*Note: If `REGEN_MNEMONIC` is not set, a testing mnemonic will be auto-generated and printed to STDOUT.*

Run the container:

```sh
(cd images/regen-sandbox && docker run -v $(pwd):/regen --env REGEN_MNEMONIC regen-sandbox:latest)
```

The above command will start up a new chain from genesis, with 5 accounts [`addr1`, ...`addr5`]. All accounts are generated from the same mnemonic with incremental `--account` indices (using HD derivation), and seeded with 10000 REGEN tokens. `addr1` is set as the single validator in this network and has an additional 40000 REGEN tokens, all of which are self-delegated.

If a `./.regen` home directory is detected, it will not initiatize a new chain, but simply run `regen start` with the existing home directory. If you want to always create a fresh home directory, you can override this behavior by adding a `-o` or `--overwrite-home-dir` flag to the end of your `docker run` command.

You can additionally provide a comma separated list of setup scripts to run as an argument at the end of the `docker run` command, like so:

```
(cd images/regen-sandbox && docker run -v $(pwd):/regen --env REGEN_MNEMONIC regen-sandbox:latest ecocredit,data)
```

This will look in `./setup` for two files, `./setup/ecocredit.sh` and `./setup/data.sh` and run them in order. Stopping the docker process and running with these arguments a second time will attempt to run the scripts again and potentially popupate a second round of data.

Available scripts:
- [setup/data.sh](./setup/data.sh) for data module transactions
- [setup/ecocredit.sh](./setup/ecocredit.sh) for ecocredit & basket transactions
- [setup/bridge.sh](./setup/bridge.sh) for bridging of ecocredits (on and off of regen network)
