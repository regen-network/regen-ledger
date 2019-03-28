module github.com/regen-network/regen-ledger

require (
	github.com/DATA-DOG/godog v0.7.10
	github.com/ZondaX/hid-go v0.4.0 // indirect
	github.com/ZondaX/ledger-go v0.4.0 // indirect
	github.com/btcsuite/btcd v0.0.0-20190115013929-ed77733ec07d // indirect
	//github.com/cosmos/cosmos-sdk v0.0.0-00010101000000-000000000000
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/leanovate/gopter v0.2.4
	github.com/lib/pq v1.0.0
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/pkg/errors v0.8.1 //
  github.com/rcrowley/go-metrics v0.0.0-20180503174638-e2704e165165 // indirect
	github.com/spf13/afero v1.2.1 // indirect
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
  github.com/regen-network/cosmos-sdk v0.0.0-0.20190328142727-7fc01b12c61a
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.0.0-0.20190328142727-7fc01b12c61a
