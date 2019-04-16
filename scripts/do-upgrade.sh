set -x
XRNDHOME=$1

if [ -f "$XRNDHOME/data/upgrade-info" ]; then
  mv $XRNDHOME/data/upgrade-info $XRNDHOME/data/upgrade-info.bak
  export UPGRADE_COMMIT=$(jq -r '.commit' < $XRNDHOME/data/upgrade-info.bak)
  if  [ $UPGRADE_COMMIT != "null" ]; then
    echo "Doing upgrade for $XRNDHOME to $UPGRADE_COMMIT"
    git fetch
    git clean -f
    git checkout -f $UPGRADE_COMMIT
    nixos-rebuild --upgrade switch
  fi
fi
