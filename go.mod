module github.com/regen-network/regen-ledger

go 1.15

require (
	github.com/CosmWasm/wasmd v0.15.0
	github.com/btcsuite/btcutil v1.0.2
	github.com/cockroachdb/apd/v2 v2.0.2
	github.com/cosmos/cosmos-sdk v0.42.0-rc0
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
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	google.golang.org/genproto v0.0.0-20210406143921-e86de6bf7a46 // indirect
	google.golang.org/grpc v1.36.1
	gopkg.in/yaml.v2 v2.4.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/cosmos/cosmos-sdk => github.com/regen-network/cosmos-sdk v0.42.4-regen-1
