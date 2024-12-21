module github.com/regen-network/regen-ledger/api/v2

go 1.21

require (
	cosmossdk.io/api v0.3.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.5
	github.com/cosmos/cosmos-sdk/orm v1.0.0-alpha.12.0.20240514101554-56648741cbd6
	github.com/cosmos/gogoproto v1.4.10 // NOTE: v1.4.11+ is not compatible with sdk v0.47
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240116215550-a9fa1716bcac // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
)

require (
	cosmossdk.io/errors v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cometbft/cometbft-db v0.7.0 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c // indirect
)

// stick with compatible version or x/exp in v0.47.x line for gogoproto v1.4.10
replace golang.org/x/exp => golang.org/x/exp v0.0.0-20230711153332-06a737ee72cb
