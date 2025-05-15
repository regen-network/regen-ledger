## Testing Regen Upgrades

To use this guide remove the `--upgrade-height 10` flag from the `entrypoint.sh` script.


To test a regen upgrade to v6 proposal a with a new binary follow
these steps:

### 1. Start devnet
```shell

chmod +x download_upgrade_binaries.sh
chmod +x setup.sh
./download_upgrade_binaries.sh
./setup.sh
```
The current version of the  `download_upgrade_binaries.sh` download only
versions `v6.0` and `v5.1.4`. New versions need to be added to the script.

### 2. Open a shell into regen-node1
```shell
docker exec -it regen-node1 bash
```
### 3. Recover the key
```shell
echo "$REGEN_NODE1_VALIDATOR_MNEMONIC" > /root/mnemonic.txt
regen keys add my_validator --recover --keyring-backend=test --home=/root/.regen < /root/mnemonic.txt

````
### 4. Create the upgrade proposal JSON
```shell
AUTHORITY=$(regen query auth module-accounts \
  --node tcp://localhost:26001 \
  --home /root/.regen \
  --output json | jq -r '.accounts[] | select(.name=="gov") | .base_account.address')

jq -n --arg authority "$AUTHORITY" '
{
  "messages": [
    {
      "@type": "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
      "authority": $authority,
      "plan": {
        "name": "v6.0",
        "height": 50,
        "info": ""
      }
    }
  ],
  "deposit": "10000000uregen",
  "title": "Upgrade to v6.0",
  "summary": "This proposal upgrades the chain to version 6 using Cosmovisor."
}
' > upgrade-info.json

```

### 4. Submit the upgrade proposal
```shell
regen tx gov submit-proposal upgrade-info.json \
  --from=my_validator \
  --chain-id=regen-devnet \
  --keyring-backend=test \
  --home=/root/.regen \
  --node tcp://localhost:26001 \
  --gas auto \
  --gas-adjustment 1.5 \
  --gas-prices 0.025uregen \
  --yes
```
### Send Deposit for proposal
```shell
regen tx gov deposit 1 90000000uregen \
  --from=my_validator \
  --chain-id=regen-devnet \
  --keyring-backend=test \
  --home=/root/.regen \
  --node tcp://localhost:26001 \
  --gas auto \
  --gas-adjustment 1.5 \
  --gas-prices 0.025uregen \
  --yes
```

### 5. Vote “Yes” from all validators

```shell
docker exec -it regen-node1 bash
```
### Node 1
```shell
regen tx gov vote 1 yes \
  --from=my_validator \
  --chain-id=regen-devnet \
  --keyring-backend=test \
  --home=/root/.regen \
  --node tcp://localhost:26001 \
  --gas auto \
  --gas-adjustment 1.5 \
  --gas-prices 0.025uregen \
  --yes


```

### Node 2
```shell
docker exec -it regen-node2 bash
```
```shell
regen tx gov vote 1 yes \
  --from=my_validator \
  --chain-id=regen-devnet \
  --keyring-backend=test \
  --home=/root/.regen \
  --node tcp://localhost:26004 \
  --gas auto \
  --gas-adjustment 1.5 \
  --gas-prices 0.025uregen \
  --yes

```

### Node 3
```shell
docker exec -it regen-node3 bash
```
```shell
regen tx gov vote 1 yes \
  --from=my_validator \
  --chain-id=regen-devnet \
  --keyring-backend=test \
  --home=/root/.regen \
  --node tcp://localhost:26007 \
  --gas auto \
  --gas-adjustment 1.5 \
  --gas-prices 0.025uregen \
  --yes

```

## Check proposal status
Voting period should end after 60s

```shell
docker exec -it regen-node1 bash
```
```shell
regen query gov proposal 1 \
  --node tcp://localhost:26001 \
  --output json | jq '.status'

```
you should see:
```text
"PROPOSAL_STATUS_PASSED"
```

## Restart Nodes
Wait for block 50 and you should see a chain halt, and then a restarts by cosmovisor:
