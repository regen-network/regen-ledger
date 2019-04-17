#!/usr/bin/env bash
set -x
cd $REGEN_LEDGER_REPO
./scripts/prepare-upgrade.sh
nixos-rebuild --upgrade switch
