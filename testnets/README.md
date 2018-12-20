# Testnets

## Testnet Status

### `xrn-1`

The initial Regen Ledger testnet `xrn-1` has been deployed.

In this testnet, validator nodes currently have ports 26656, 26657 and 1317 open for testing purposes. In the future,
the testnet will be setup with more security hardening via sentry and seed nodes.

The validator node URL's are as follows:

* [xrn-us-east-1.regen.network](http://xrn-us-east-1.regen.network:26657)
* [xrn-us-west-1.regen.network](http://xrn-us-west-1.regen.network:26657)
* [xrn-eu-central-1.regen.network](http://xrn-eu-central-1.regen.network:26657)

`xrncli` can be configured to connect to the testnet as follows:

```sh
xrncli init --chain-id xrn-1 --node tcp://xrn-us-east-1.regen.network:26657
```

## Running a full node in the cloud

Node configurations for [NixOS](https://nixos.org) are provided in this repository.

[../module.nix](../module.nix) contains a NixOS module for running a node.