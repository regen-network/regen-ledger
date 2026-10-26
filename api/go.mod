module github.com/regen-network/regen-ledger/api/v2

go 1.26.5

require (
	cosmossdk.io/api v1.1.0
	github.com/cosmos/cosmos-proto v1.0.0-beta.5
	github.com/cosmos/gogoproto v1.7.2 // NOTE: v1.4.11+ is not compatible with sdk v0.47
	google.golang.org/genproto/googleapis/api v0.0.0-20260511170946-3700d4141b60
	google.golang.org/grpc v1.79.3
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/go-cmp v0.7.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260511170946-3700d4141b60 // indirect
)

require (
	cosmossdk.io/core v1.1.0
	cosmossdk.io/errors v1.1.0
	github.com/cosmos/cosmos-db v1.1.3
	github.com/golang/mock v1.6.0
	github.com/iancoleman/strcase v0.3.0
	github.com/regen-network/gocuke v1.1.1
	github.com/stretchr/testify v1.11.1
	gotest.tools/v3 v3.5.1
	pgregory.net/rapid v1.1.0
)

require (
	github.com/DataDog/zstd v1.5.7 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/cockroachdb/errors v1.12.0 // indirect
	github.com/cockroachdb/fifo v0.0.0-20240816210425-c5d0cb0b6fc0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20241215232642-bb51bb14a506 // indirect
	github.com/cockroachdb/pebble v1.1.5 // indirect
	github.com/cockroachdb/redact v1.1.6 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20250429170803-42689b6311bb // indirect
	github.com/cucumber/gherkin/go/v27 v27.0.0 // indirect
	github.com/cucumber/messages/go/v22 v22.0.0 // indirect
	github.com/cucumber/tag-expressions/go/v6 v6.1.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/getsentry/sentry-go v0.42.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/klauspost/compress v1.18.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/linxGnu/grocksdb v1.8.15 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/onsi/gomega v1.20.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.23.2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.67.5 // indirect
	github.com/prometheus/procfs v0.19.2 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	golang.org/x/exp v0.0.0-20260112195511-716be5621a96 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// stick with compatible version or x/exp in v0.47.x line for gogoproto v1.4.10
replace golang.org/x/exp => golang.org/x/exp v0.0.0-20230711153332-06a737ee72cb
