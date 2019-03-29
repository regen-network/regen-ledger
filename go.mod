module github.com/regen-network/regen-ledger

require (
	github.com/DATA-DOG/godog v0.7.10
	github.com/ZondaX/hid-go v0.4.0 // indirect
	github.com/ZondaX/ledger-go v0.4.0 // indirect
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/campoy/unique v0.0.0-20180121183637-88950e537e7e
	github.com/cosmos/cosmos-sdk v0.32.0
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/lint v0.0.0-20180702182130-06c8688daad7 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/leanovate/gopter v0.2.4
	github.com/lib/pq v1.0.0
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pkg/errors v0.8.0 //
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.0.3
	github.com/stretchr/testify v1.2.2
	github.com/syndtr/goleveldb v0.0.0-20181128100959-b001fa50d6b2 // indirect
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/tendermint v0.31.0-dev0
	github.com/twpayne/go-geom v1.0.4
	github.com/zondax/ledger-cosmos-go v0.9.9 // indirect
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd // indirect
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2 // indirect
	google.golang.org/genproto v0.0.0-20190201180003-4b09977fb922 // indirect
//github.com/regen-network/cosmos-sdk v0.0.0-0.20190328142727-7fc01b12c61a
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.0.0-0.20190329021654-e7d1e85e4888
