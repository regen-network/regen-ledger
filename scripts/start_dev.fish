
set coins "10000000000stake,100000000000samoleans"
set coinsV "5000000000stake"
set REGEN_KEYRING_BACKEND test
set CHAINID regen-local
set TMCFG ~/.regen/config/config.toml
set APPCFG ~/.regen/config/app.toml

rm -rf ~/.regen
./build/regen keys add validator --keyring-backend test

./build/regen init dev-val --chain-id $CHAINID
./build/regen genesis add-genesis-account validator --keyring-backend test $coins
./build/regen genesis gentx validator $coinsV --chain-id $CHAINID --keyring-backend test
./build/regen genesis collect-gentxs
./build/regen genesis validate-genesis

perl -i -pe 's|timeout_commit = ".*?"|timeout_commit = "2s"|g' $TMCFG
perl -i -pe 's|minimum-gas-prices = ""|minimum-gas-prices = "0stake"|g' $APPCFG

./build/regen start --api.enable true --grpc.address="0.0.0.0:9090"


exit
