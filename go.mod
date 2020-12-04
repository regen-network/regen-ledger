module github.com/regen-network/regen-ledger

go 1.15

require (
	github.com/CosmWasm/wasmd v0.12.0-alpha1.0.20201105090402-41cf73dae976
	github.com/btcsuite/btcutil v1.0.2
	github.com/cockroachdb/apd/v2 v2.0.2
	github.com/cosmos/cosmos-sdk v0.40.0-rc2
	github.com/enigmampc/btcutil v1.0.3-0.20200723161021-e2fb6adb2a25
	github.com/gogo/protobuf v1.3.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.15.2
	github.com/ipfs/go-cid v0.0.7
	github.com/lib/pq v1.8.0 // indirect
	github.com/multiformats/go-multihash v0.0.14
	github.com/rakyll/statik v0.1.7
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc5
	github.com/tendermint/tm-db v0.6.2
	golang.org/x/crypto v0.0.0-20201203163018-be400aefbc4c // indirect
	google.golang.org/genproto v0.0.0-20201204160425-06b3db808446 // indirect
	google.golang.org/grpc v1.33.0
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
