#!/usr/bin/env bash
set -x
ln -s ./prepare-upgrade.sh $1/config/prepare-upgrade
ln -s ./do-upgrade.sh $1/config/do-upgrade
