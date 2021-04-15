module github.com/regen-network/regen-ledger/orm

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/regen-network/regen-ledger/testutil v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tm-db v0.6.4
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/regen-network/regen-ledger/testutil => ../testutil

replace github.com/regen-network/regen-ledger => ../
