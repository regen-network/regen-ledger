module github.com/regen-network/regen-ledger/types/testutil

go 1.15

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/regen-network/regen-ledger/orm => ../../orm

replace github.com/regen-network/regen-ledger => ../../

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.10
	google.golang.org/grpc v1.37.0
)
