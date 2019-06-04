## Running a full node with NixOS

Node configurations for [NixOS](https://nixos.org) are provided in this repository.

[./module.nix](module.nix) contains a NixOS module for running a node.

The steps for setting this up on a running NixOS machine are roughly as follows:
1. Clone the git repository into a local folder on the machine
2. Import `module.nix` in `/etc/nixos/configuration`:
```
  imports =
    [
      ./hardware-configuration.nix
      /path-to-regen-ledger-git-repo/module.nix
    ];
```
3. Enable the `xrn` programs in `/etc/nixos/configuration` and run `nixos-rebuild switch):
```
  programs.xrn.enable = true;
```
4. Run `xrncli init --home /var/xrnd`
5. Configure node configuration in `/var/xrnd/config` (`genesis.json`, `config.toml`, etc.)
6. If you want to enable automatic upgrades, run `link-upgrade-scripts.sh` from this directory with the path to
your `xrnd` home directory as the first argument and set `services.xrnd.repoPath` to the path to your
local regen-ledger git repository.
7. Enable the xrnd service in `/etc/nixos/configuration` and run `nixos-rebuild switch`
```
  services.xrnd.enable = true;
  services.xrnd.moniker = "my-node-moniker";
```
