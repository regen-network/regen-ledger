#!/usr/bin/env bash
set -x
ln -s ./scripts/prepare-upgrade.sh $1/config/prepare-upgrade
ln -s ./scripts/do-upgrade.sh $1/config/do-upgrade
