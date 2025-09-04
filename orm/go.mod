module github.com/regen-network/regen-ledger/orm

go 1.23.0

toolchain go1.23.2

require (
	cosmossdk.io/api v0.8.0
	cosmossdk.io/core v0.11.1
	cosmossdk.io/errors v1.0.1
	github.com/cosmos/cosmos-db v1.1.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.5
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.7.0
	github.com/iancoleman/strcase v0.2.0
	github.com/regen-network/gocuke v0.6.2
	github.com/regen-network/regen-ledger/api/v2 v2.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.10.0
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
	gotest.tools/v3 v3.5.1
	pgregory.net/rapid v1.1.0
)

require (
	buf.build/gen/go/cometbft/cometbft/protocolbuffers/go v1.36.1-20241120201313-68e42a58b301.1 // indirect
	buf.build/gen/go/cosmos/gogo-proto/protocolbuffers/go v1.36.1-20240130113600-88ef6483f90f.1 // indirect
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/alecthomas/participle/v2 v2.0.0-alpha7 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cockroachdb/apd/v3 v3.1.0 // indirect
	github.com/cockroachdb/errors v1.11.3 // indirect
	github.com/cockroachdb/fifo v0.0.0-20240606204812-0bbfbd93a7ce // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v1.1.2 // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/cosmos/gogoproto v1.7.0 // indirect
	github.com/cucumber/common/gherkin/go/v22 v22.0.0 // indirect
	github.com/cucumber/common/messages/go/v17 v17.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/linxGnu/grocksdb v1.8.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250422160041-2d3770c4ea7f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	cosmossdk.io/core => cosmossdk.io/core v0.11.1
	github.com/regen-network/regen-ledger/api/v2 => ./../api
)
