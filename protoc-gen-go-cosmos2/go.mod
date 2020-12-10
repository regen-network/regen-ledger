module github.com/regen-network/regen-ledger/protoc-gen-go-cosmos2

go 1.15

require (
	google.golang.org/protobuf v1.23.0
	github.com/regen-network/regen-ledger/types/proto v0.0.0
)

replace github.com/regen-network/regen-ledger/types/proto => ../types/proto
