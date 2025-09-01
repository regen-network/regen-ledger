module github.com/regen-network/regen-ledger/orm

go 1.23

require (
	cosmossdk.io/api v0.8.0
	cosmossdk.io/core v0.6.1
	cosmossdk.io/depinject v1.0.0-alpha.3
	cosmossdk.io/errors v1.0.1
	github.com/cosmos/cosmos-db v1.0.0-rc.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.5
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/iancoleman/strcase v0.2.0
	github.com/regen-network/gocuke v0.6.2
	github.com/regen-network/regen-ledger/api/v2 v2.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb
	google.golang.org/grpc v1.68.1
	google.golang.org/protobuf v1.36.1
	gotest.tools/v3 v3.5.1
	pgregory.net/rapid v1.1.0
)

require (
	buf.build/gen/go/cometbft/cometbft/protocolbuffers/go v1.36.1-20241120201313-68e42a58b301.1 // indirect
	buf.build/gen/go/cosmos/gogo-proto/protocolbuffers/go v1.36.1-20240130113600-88ef6483f90f.1 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/alecthomas/participle/v2 v2.0.0-alpha7 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cockroachdb/apd/v3 v3.1.0 // indirect
	github.com/cockroachdb/errors v1.9.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v0.0.0-20230412222916-60cfeb46143b // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/cosmos/gogoproto v1.7.0 // indirect
	github.com/cucumber/common/gherkin/go/v22 v22.0.0 // indirect
	github.com/cucumber/common/messages/go/v17 v17.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.20.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/linxGnu/grocksdb v1.7.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.15.1 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.43.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace github.com/regen-network/regen-ledger/api/v2 => ./../api
