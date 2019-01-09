# Testnets

## Testnet Status

### `xrn-test-2`

Deployed at `2018-12-19T20:40:06.463846Z`, live as of this writing.

In this testnet, validator nodes currently have ports 26656, 26657 and 1317 open for testing purposes. In the future,
the testnet will be setup with more security hardening via sentry and seed nodes.

The validator node URL's are as follows:

* [xrn-us-east-1.regen.network](http://xrn-us-east-1.regen.network:26657)
* [xrn-us-west-1.regen.network](http://xrn-us-west-1.regen.network:26657)
* [xrn-eu-central-1.regen.network](http://xrn-eu-central-1.regen.network:26657)

`xrncli` can be configured to connect to the testnet as follows:

```sh
xrncli init --chain-id xrn-test-2 --node tcp://xrn-us-east-1.regen.network:26657
```

### `xrn-1`

The initial Regen Ledger testnet `xrn-1` was deployed on 2018-12-19.

## Running a full node in the cloud

Node configurations for [NixOS](https://nixos.org) are provided in this repository.

[../module.nix](../module.nix) contains a NixOS module for running a node.