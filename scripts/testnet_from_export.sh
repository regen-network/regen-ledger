#!/usr/bin/env bash
# shellcheck disable=SC2094

set -e

# home is directory path of node home
home=$HOME/.regen

# export is the filepath of the state export
export=$HOME/Downloads/state_export.json

# set script input options
while getopts ":h:e:" option; do
  case $option in
    h)
      home=$OPTARG;;
    e)
      export=$OPTARG;;
    \?)
      echo "Error: invalid option"
      exit 1
  esac
done

# check home directory and confirm removal if exists
if [ -d "$home" ]; then
  read -r -p "WARNING: This script will remove $home. Would you like to continue? [y/N] " confirm
  case "$confirm" in
    [yY][eE][sS]|[yY])
      rm -rf "$home"
      ;;
    *)
      exit 0
      ;;
  esac
fi

# node_genesis is the node genesis file
node_genesis=$home/config/genesis.json

# tmp_dir is the temporary directory used for
# exported state to be merged into node_genesis
tmp_dir=./tmp

# tmp_genesis is the exported state from a single
# node network started by the new validator
tmp_genesis=$tmp_dir/genesis.json

# chain_id is the chain id
chain_id=$(jq .chain_id "$export")

# bond_denom is the staking token denom
bond_denom=$(jq .app_state.staking.params.bond_denom "$export")

# amount is the amount of tokens to stake (must be
# more than 2/3 of the total amount staked)
amount="1000000000000000"

# tokens is the token amount and denom to stake
tokens=$(echo "${amount}${bond_denom}" | tr -d '"')

# tokens2 is token amount x2 providing a starting balance with amount
tokens2=$(echo "$(( "$amount" + "$amount" ))${bond_denom}" | tr -d '"')

# gas_price is the minimum gas price for the node
gas_price=$(echo "0${bond_denom}" | tr -d '"')

# build binary
make build

# add test key using keyring backend
./build/regen keys add test --home "$home" --keyring-backend test

# create genesis values for validator node
./build/regen init test --chain-id "$chain_id" --home "$home"
./build/regen add-genesis-account test "$tokens2" --home "$home" --keyring-backend test
./build/regen gentx test "$tokens" --chain-id "$chain_id" --home "$home" --keyring-backend test
./build/regen collect-gentxs --home "$home"

# update default denom to bond denom
sed -i "s|\"stake\"|$bond_denom|g" "$node_genesis"

# start node and deliver genesis transaction
./build/regen start --home "$home" --halt-height 1 --minimum-gas-prices "$gas_price" && wait $!

# create temporary directory
mkdir -p "$tmp_dir"

# copy single node network state to temporary genesis file
./build/regen export --home "$home" > "$tmp_genesis"

# keys are the genesis file key names to update
keys=(
  "app_state.auth.accounts"
  "app_state.bank.balances"
  "app_state.bank.supply"
  "app_state.distribution.delegator_starting_infos"
  "app_state.distribution.outstanding_rewards"
  "app_state.distribution.validator_current_rewards"
  "app_state.distribution.validator_historical_rewards"
  "app_state.slashing.signing_infos"
  "app_state.staking.delegations"
  "app_state.staking.last_total_power"
  "app_state.staking.last_validator_powers"
  "app_state.staking.params.max_validators"
  "app_state.staking.validators"
  "validators"
)

# copy single node network state to json files
for i in "${!keys[@]}"; do

  # simple var
  k=${keys[$i]}

  # create json file for each state object (some are too large
  # to pass directly to jq so we create a file for each)
  cat <<< $(jq ".$k" "$tmp_genesis") > "$tmp_dir/$k.json"

done

# overwrite node genesis file with state export
cp "$export" "$node_genesis"

# add single node network state to node genesis file
for i in "${!keys[@]}"; do

  # simple var
  k=${keys[$i]}

  if [ "$k" == "app_state.bank.balances" ]; then

    # 1) add balance for validator account

    # append balances from single node network state
    cat <<< $(jq --argfile v "$tmp_dir/$k.json" '.'"$k"' += $v' "$node_genesis") > "$node_genesis"

    # 2) update balance for "bonded tokens pool"

    # get account address for "bonded tokens pool"
    a1=$(jq "[.app_state.auth.accounts[]|select(.name==\"bonded_tokens_pool\")][0].base_account.address" "$node_genesis")

    # get balance amount for "bonded tokens pool"
    b1=$(jq "[.${k}[]|select(.address==$a1)][0].coins[]|select(.denom==$bond_denom).amount" "$node_genesis" | tr -d '"')

    # add new validator stake amount to balance amount
    coins='[{"amount": "'$(( "$b1" + "$amount" ))'", "denom": '$bond_denom'}]'

    # create json object for account balance with deleted bond denom balance
    json=$(jq '[.'"$k"'[]|select(.address=='"$a1"')][0]|del(.coins[]|select(.denom=='"$bond_denom"'))' "$node_genesis")

    # update account balance with updated bond denom balance
    json=$(jq '.coins = '"$coins"'' <<< "$json")

    # delete old account balance
    cat <<< $(jq 'del(.'"$k"'[]|select(.address=='"$a1"'))' "$node_genesis") > "$node_genesis"

    # add updated account balance
    cat <<< $(jq --argjson v "[$json]" '.'"$k"' += $v' "$node_genesis") > "$node_genesis"

    # 3) update balance for "fee collector"

    # get account address for "fee collector"
    a2=$(jq "[.app_state.auth.accounts[]|select(.name==\"fee_collector\")][0].base_account.address" "$node_genesis")

    # TODO: "fee collector" does not have a balance in state exports used when
    # testing therefore simply adding the balance in step (1) is sufficient but
    # this will need to be updated if "fee collector" has an existing balance

  elif [ "$k" == "app_state.bank.supply" ]; then

    # get supply amount for bond denom from single node network state
    a1=$(jq "[.${k}[]|select(.denom==$bond_denom)][0].amount" "$tmp_genesis" | tr -d '"')

    # get supply amount for bond denom from node genesis
    a2=$(jq "[.${k}[]|select(.denom==$bond_denom)][0].amount" "$node_genesis" | tr -d '"')

    # set new supply for bond denom
    ts='[{"amount": "'$(( "$a1" + "$a2" ))'", "denom": '$bond_denom'}]'

    # delete supply for bond denom
    cat <<< $(jq 'del(.'"$k"'[]|select(.denom=='"$bond_denom"'))' "$node_genesis") > "$node_genesis"

    # add updated supply for bond denom
    cat <<< $(jq --argjson v "$ts" '.'"$k"' += $v' "$node_genesis") > "$node_genesis"

  elif [ "$k" == "app_state.staking.last_total_power" ]; then

    # get last total power from single node network state
    p1=$(jq ".$k" "$tmp_genesis" | tr -d '"')

    # get last total power from state export
    p2=$(jq ".$k" "$node_genesis" | tr -d '"')

    # set new last total power
    tp=$(( "$p1" + "$p2" ))

    # update last total power
    cat <<< $(jq --arg v "$tp" '.'"$k"' = $v' "$node_genesis") > "$node_genesis"

  elif [ "$k" == "app_state.staking.params.max_validators" ]; then

    # get max validators from state export
    mv=$(jq ".$k" "$node_genesis" | tr -d '"')

    # update max validators to include new validator
    cat <<< $(jq --arg v "$(( "$mv" + 1 ))" '.'"$k"' = $v' "$node_genesis") > "$node_genesis"

  else

    # append single node network state values
    cat <<< $(jq --argfile v "$tmp_dir/$k.json" '.'"$k"' += $v' "$node_genesis") > "$node_genesis"

  fi
done

# remove previous state (but keep validator state)
rm -rf "$home/data/application.db/"
rm -rf "$home/data/blockstore.db/"
rm -rf "$home/data/cs.wal/"
rm -rf "$home/data/evidence.db/"
rm -rf "$home/data/snapshots/"
rm -rf "$home/data/state.db/"
rm -rf "$home/data/tx_index.db/"

# remove temporary directory
rm -rf $tmp_dir

# reduce voting period to 20 seconds
cat <<< $(jq '.app_state.gov.voting_params.voting_period = "20s"' "$node_genesis") > "$node_genesis"

# start node
./build/regen start --home "$home" --fast_sync=false --minimum-gas-prices "$gas_price"
