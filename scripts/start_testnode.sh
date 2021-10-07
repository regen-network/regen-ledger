#!/bin/bash

exit_with_error()
{
  echo "$1" 1>&2
  exit 1
}

KeyringBackend=test
Chain=test

# Get the options
while getopts ":kc:" option; do
   case $option in
      k) # Keyring backend
         KeyringBackend=$OPTARG;;
      c) # Enter a name
         Chain=$OPTARG;;
     \?) # Invalid option
         echo "Error: Invalid option"
         exit 1
   esac
done

# Make sure the path is set correctly
export PATH=~/go/bin:$PATH

echo "REGEN Version: `regen version`"

regen keys add validator --keyring-backend ${KeyringBackend} || exit_with_error "Error: Validator add failed"
regen keys add delegator --keyring-backend ${KeyringBackend} || exit_with_error "Error: Delegator add failed"
regen init node --chain-id ${Chain} || exit_with_error "Error: Could not init node"

# Change the staking token to uregen
# Note: sed works differently on different platforms
echo "Updating your staking token to uregen in the genesis file..."
OS=`uname`
if [[ $OS == "Linux"* ]]; then
    echo "Your OS is a Linux variant..."
    sed -i "s/stake/uregen/g" ~/.regen/config/genesis.json || exit_with_error "Error: Could not update staking token"
elif [[ $OS == "Darwin"* ]]; then
    echo "Your OS is Mac OS/darwin..."
    sed -i "" "s/stake/uregen/g" ~/.regen/config/genesis.json || exit_with_error "Error: Could not update staking token"
else
    # Dunno
    echo "Your OS is not supported"
    exit 1
fi

echo "Adding validator to genesis.json..."
regen add-genesis-account validator 5000000000uregen --keyring-backend ${KeyringBackend} || exit_with_error "Error: Could not add validator to genesis"

echo "Adding delegator to genesis.json..."
regen add-genesis-account delegator 2000000000uregen --keyring-backend ${KeyringBackend} || exit_with_error "Error: Could not add delegator to genesis"
echo "Creating genesis transaction..."
regen gentx validator 1000000uregen --chain-id ${Chain} --keyring-backend ${KeyringBackend} || exit_with_error "Error: Genesis transaction failed"

echo "Adding genesis transaction to genesis.json..."
regen collect-gentxs || exit_with_error "Error: Could not add transaction to genesis"

echo "If there were no errors above, you can now type 'regen start' to start your node"
