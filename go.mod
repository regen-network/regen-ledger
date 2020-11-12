module github.com/regen-network/regen-ledger

go 1.14

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/aokoli/goutils v1.1.0 // indirect
	github.com/btcsuite/btcutil v1.0.2
	github.com/cockroachdb/apd/v2 v2.0.2
	github.com/cosmos/cosmos-sdk v0.40.0-rc2
	github.com/enigmampc/btcutil v1.0.3-0.20200723161021-e2fb6adb2a25
	github.com/envoyproxy/protoc-gen-validate v0.4.1 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.15.2
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/ipfs/go-cid v0.0.7
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/multiformats/go-multihash v0.0.14
	github.com/mwitkow/go-proto-validators v0.3.2 // indirect
	github.com/pseudomuto/protoc-gen-doc v1.3.2 // indirect
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc5
	github.com/tendermint/tm-db v0.6.2
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	google.golang.org/genproto v0.0.0-20201111145450-ac7456db90a6 // indirect
	google.golang.org/grpc v1.33.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
