#!/usr/bin/env bash
set -x
if [ $1 != "null" ]; then
  export UPGRADE_COMMIT=$(echo $1 | jq -r '.info.commit')
  if  [ $UPGRADE_COMMIT != "null" ]; then
      echo "Doing upgrade to $UPGRADE_COMMIT"
      cd $REGEN_LEDGER_REPO
      git fetch
      git clean -f
      git checkout -f $UPGRADE_COMMIT
      nixos-rebuild --upgrade build
  fi
fi
