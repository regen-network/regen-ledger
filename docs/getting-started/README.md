# Quick Start

This document provides instructions for running a single node network on your local machine and then submitting your first few transactions to that network using the command line. Running a single node network is a great way to get familiar with Regen Ledger and its functionality.

<!-- TODO: update example to use ecocredit module once updated to v1.1 -->

## Prerequisites

In order to install the `regen` binary, you'll need the following: 

- Git `>=2`
- Make `>=4`
- Go `>=1.15`

For more information, see [Prerequisites](./prerequisites). 

## Install Regen

The `regen` binary serves as the node client and the application client. In other words, the `regen` binary can be used to both run a node and interact with it.

Clone the `regen-ledger` repository:
```
git clone https://github.com/regen-network/regen-ledger
```

Change to the `regen-ledger` directory:
```
cd regen-ledger
```

Check out the latest stable version:
```
git checkout v1.0.0
```

Install the `regen` binary:
```
make install
```

Check to make sure the install was successful:
```
regen version
```

You should see `v1.0.0` printed to the console. Now that you have successfully installed the `regen` binary, the next step will be to add a couple test accounts.

## Create Accounts

In this section, you will create two test accounts. You will name the first account `validator` and the second account `delegator`. You will create both accounts using the `test` backend, meaning both accounts will not be securely stored and should not be used in a production environment. When using the `test` backend, accounts are stored in the home directory (more on this in the next section). 

Create `validator` account:
```
regen keys add validator --keyring-backend test
```

Create `delegator` account:
```
regen keys add delegator --keyring-backend test
```

After running each command, information about each account will be printed to the console. The next step will be to initialize the node.

## Initialize Node

Initializing the node will create the `config` and `data` directories within the home directory. The `config` directory is where configuration files for the node are stored and the `data` directory is where the data for the blockchain is stored. The default home directory is `~/.regen`.

Initialize the node:
```
regen init node --chain-id test
```

In this case, `node` is the name (or "moniker") of the node and `test` is the chain ID. Feel free to change these values but make sure to use the same value for `chain-id` in the following steps.

## Update Genesis

When the node was initialized, a `genesis.json` file was created within the `config` directory. In this section, you will be adding two genesis accounts (accounts with an initial token balance) and a genesis transaction (a transaction that registers the validator account in the validator set).

Update native staking token to `uregen`:
```
sed -i "s/stake/uregen/g" ~/.regen/config/genesis.json
```

Add `validator` account to `genesis.json`:
```
regen add-genesis-account validator 5000000000uregen --keyring-backend test
```

Add `delegator` account to `genesis.json`:
```
regen add-genesis-account delegator 2000000000uregen --keyring-backend test
```

Create genesis transaction:
```
regen gentx validator 1000000uregen --chain-id test --keyring-backend test
```

Add genesis transaction to `genesis.json`:
```
regen collect-gentxs
```

Now that you have updated the `genesis.json` file, you are ready to start the node. Starting a node with a new genesis file will create a new blockchain.

## Start Node

Well, what are you waiting for?

Start the node:
```
regen start
```

You should see logs printed in your terminal with information about services starting up followed by blocks being produced and committed to your local blockchain.

## Test Commands

Now that you have a single node network running, you can open a new terminal window and interact with the node using the same `regen` binary. Let's delegate some `uregen` tokens to the validator and then collect the rewards.

Get the validator address for the `validator` account:
```
regen keys show validator --bech val --keyring-backend test
```

Using the validator address, delegate some `uregen` tokens:
```
regen tx staking delegate [validator_address] 10000000uregen --from delegator --keyring-backend test --chain-id test
```

In order to query all delegations, you'll need the address for the `delegator` account:
```
regen keys show delegator --keyring-backend test
```

Using the address, query all delegations for the `delegator` account:
```
regen q staking delegations [delegator_address]
```

Query the rewards using the delegator address and the validator address:
```
regen q distribution rewards [delegator_address] [validator_address]
```

Withdraw the rewards:
```
regen tx distribution withdraw-all-rewards --from delegator --keyring-backend test --chain-id test
```

Check the account balance:
```
regen q bank balances [delegator_address]
```

You have successfully delegated `uregen` tokens to the `validator` account from the `delegator` account and then collected the rewards.