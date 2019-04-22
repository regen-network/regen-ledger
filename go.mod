module github.com/regen-network/regen-ledger

require (
	github.com/DATA-DOG/godog v0.7.10
	github.com/ZondaX/hid-go v0.4.0 // indirect
	github.com/ZondaX/ledger-go v0.4.0 // indirect
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/campoy/unique v0.0.0-20180121183637-88950e537e7e
	github.com/cosmos/cosmos-sdk v0.32.0
	github.com/ethereum/go-ethereum v1.8.22 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/leanovate/gopter v0.0.0-20190000000000-6e7780f59df75750618bf30eeafcb1a88e457fcb
	github.com/lib/pq v1.0.0
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.0.3
	github.com/stretchr/testify v1.2.2
	github.com/syndtr/goleveldb v0.0.0-20181128100959-b001fa50d6b2 // indirect
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/tendermint v0.31.5
	github.com/twpayne/go-geom v1.0.4
	github.com/zondax/ledger-cosmos-go v0.9.2 // indirect
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd // indirect
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2 // indirect
	google.golang.org/genproto v0.0.0-20190201180003-4b09977fb922 // indirect
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.0.0-0.20190329021654-6d0b10989a738540d655b556884aab55b99b3b7d
