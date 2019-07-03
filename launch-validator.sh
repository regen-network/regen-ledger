#!/bin/bash
wget https://raw.githubusercontent.com/jim380/node_tooling/master/Cosmos/CLI/go_install.sh
chmod +x go_install.sh
./go_install.sh -v 1.12.5
echo $GOPATH
mkdir -p $GOPATH/src/github.com/regen
cd $GOPATH/src/github.com/regen
git clone https://github.com/regen-network/regen-ledger
cd regen-ledger
make install
xrnd init --chain-id=regen-test-1001 swid
# $ xrncli keys add <your_wallet_name>




