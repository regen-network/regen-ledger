# Regen Ledger Sandbox

This container (and bootstrapping scripts) is intended to aid with end-to-end testing of client libraries, applications, manual local network tests, and other scenarios where you want to be able to quickly bootstrap a single node regen network (targeting any version or commit hash of regen-ledger), and populate the network with some basic data for your e2e testing or exploratory purposes.


Building locally:
```
cd regen-ledger
docker build . -f images/regen-sandbox/Dockerfile -t regen-sandbox
```

Running the container does the following:
```
cd regen-ledger/images/regen-sandbox
export REGEN_MNEMONIC="YOUR TESTING MNEMONIC"

(cd images/regen-sandbox && docker run -v $(pwd):/regen --env REGEN_MNEMONIC regen-sandbox:latest)
```

The above command will simply start up a new chain from genesis. If a `./.regen` home directory is detected, it will not initiatize a new chain, but simply run `regen start` with the existing home directory.

You can additionally provide a comma separated list of setup scripts to run as an argument at the end of the `docker run` command, like so:

```
docker run -v $(pwd):/regen --env REGEN_MNEMONIC regen-sandbox:latest ecocredit,data
```

This will look in `./setup` for two files, `./setup/ecocredit.sh` and `./setup/data.sh` and run them in order. Stopping the docker process and running with these arguments a second time will attempt to run the scripts again and potentially popupate a second round of data.

