#!/usr/bin/env bash
set -x
ln -s ./nixos/prepare-upgrade.sh $1/config/prepare-upgrade
ln -s ./nixos/do-upgrade.sh $1/config/do-upgrade
