module github.com/regen-network/regen-ledger/x/data

go 1.15

require (
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ipfs/go-cid v0.0.7
	github.com/regen-network/regen-ledger/types v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc v1.38.0
)

replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.44.0-regen-2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/regen-network/regen-ledger/types => ../../types
