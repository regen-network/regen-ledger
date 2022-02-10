module github.com/regen-network/regen-ledger/v2

go 1.17

require (
	github.com/cosmos/cosmos-sdk v0.44.3
	github.com/cosmos/ibc-go/v2 v2.0.0
	github.com/gorilla/mux v1.8.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/regen-ledger/types v1.0.0
	github.com/regen-network/regen-ledger/x/data v0.0.0-20210602121340-fa967f821a6e
	github.com/regen-network/regen-ledger/x/ecocredit v1.0.0
	github.com/regen-network/regen-ledger/x/group v1.0.0-beta1
	github.com/rs/zerolog v1.23.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
	github.com/tendermint/tm-db v0.6.6
	golang.org/x/crypto v0.0.0-20220210151621-f4118a5b28e2 // indirect
	google.golang.org/genproto v0.0.0-20220210181026-6fee9acbd336 // indirect
)

require golang.org/x/text v0.3.7 // indirect

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.44.2-regen-1

replace github.com/regen-network/regen-ledger/types => ./types

replace github.com/regen-network/regen-ledger/orm => ./orm

replace github.com/regen-network/regen-ledger/x/data => ./x/data

replace github.com/regen-network/regen-ledger/x/ecocredit => ./x/ecocredit

replace github.com/regen-network/regen-ledger/x/group => ./x/group
