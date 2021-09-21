module github.com/regen-network/regen-ledger/x/group

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/regen-network/regen-ledger/orm v0.0.0-00010101000000-000000000000
	github.com/regen-network/regen-ledger/types v0.0.0-00010101000000-000000000000
	github.com/regen-network/regen-ledger/x/data v0.0.0-20210602121340-fa967f821a6e
	github.com/regen-network/regen-ledger/x/ecocredit v0.0.0-20210602121340-fa967f821a6e
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.12
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.44.0-regen-1

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/regen-network/regen-ledger/orm => ../../orm

replace github.com/regen-network/regen-ledger/types => ../../types

replace github.com/regen-network/regen-ledger/x/ecocredit => ../ecocredit
