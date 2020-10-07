module github.com/regen-network/regen-ledger

go 1.14

require (
	bou.ke/monkey v1.0.1 // indirect
	github.com/DATA-DOG/godog v0.7.10
	github.com/ZondaX/hid-go v0.4.0 // indirect
	github.com/ZondaX/ledger-go v0.4.0 // indirect
	github.com/btcsuite/btcutil v1.0.2
	github.com/campoy/unique v0.0.0-20180121183637-88950e537e7e
	github.com/cosmos/cosmos-sdk v0.34.4-0.20201005082218-6c1c2cce0461
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d // indirect
	github.com/enigmampc/btcutil v1.0.3-0.20200723161021-e2fb6adb2a25
	github.com/ethereum/go-ethereum v1.8.22 // indirect
	github.com/golang/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/leanovate/gopter v0.2.5-0.20190326081808-6e7780f59df7
	github.com/lib/pq v1.0.0
	github.com/otiai10/copy v1.2.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.0-rc4.0.20201005104435-7e27e9b85222
	github.com/twpayne/go-geom v1.0.4
	github.com/zondax/ledger-cosmos-go v0.9.2 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

replace github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.34.4-0.20201005082218-6c1c2cce0461
