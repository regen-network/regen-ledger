module github.com/regen-network/regen-ledger

go 1.15

require (
	github.com/CosmWasm/wasmd v0.14.0
	github.com/btcsuite/btcutil v1.0.2
	github.com/cockroachdb/apd/v2 v2.0.2
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/enigmampc/btcutil v1.0.3-0.20200723161021-e2fb6adb2a25
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/ipfs/go-cid v0.0.7
	github.com/lib/pq v1.8.0 // indirect
	github.com/multiformats/go-multihash v0.0.14 // indirect
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/zerolog v1.20.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.3
	github.com/tendermint/tm-db v0.6.3
	google.golang.org/genproto v0.0.0-20210203152818-3206188e46ba // indirect
	google.golang.org/grpc v1.35.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
