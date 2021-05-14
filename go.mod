module github.com/regen-network/regen-ledger

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/cosmos/ibc-go v1.0.0-alpha2
	github.com/gorilla/mux v1.8.0
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/regen-ledger/types v0.0.0-00010101000000-000000000000
	github.com/regen-network/regen-ledger/x/data v0.0.0-00010101000000-000000000000
	github.com/regen-network/regen-ledger/x/ecocredit v0.0.0-00010101000000-000000000000
	github.com/regen-network/regen-ledger/x/group v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	google.golang.org/genproto v0.0.0-20210423144448-3a41ef94ed2b // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/regen-network/regen-ledger/types => ./types

replace github.com/regen-network/regen-ledger/orm => ./orm

replace github.com/regen-network/regen-ledger/x/data => ./x/data

replace github.com/regen-network/regen-ledger/x/ecocredit => ./x/ecocredit

replace github.com/regen-network/regen-ledger/x/group => ./x/group
