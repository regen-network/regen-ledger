regen keys add validator --keyring-backend test 
regen keys add key1 --keyring-backend test 
regen init myvalidator --chain-id testnet  --default-denom uregen
regen genesis add-genesis-account validator 1000000000uregen  --keyring-backend test
regen genesis add-genesis-account key1 1000000000uregen  --keyring-backend test
regen genesis gentx validator 1000000000uregen --chain-id testnet  --keyring-backend test
regen genesis collect-gentxs
regen start