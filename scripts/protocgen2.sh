# this script is for generating protobuf files for the new google.golang.org/protobuf API

set -eo pipefail

protoc_gen_install() {
  go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar@latest #2>/dev/null
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest #2>/dev/null
  go install github.com/cosmos/cosmos-sdk/orm/cmd/protoc-gen-go-cosmos-orm@latest #2>/dev/null
}

protoc_gen_install

echo "Generating API module"
(cd proto; buf generate --template buf.gen.pulsar.yaml)
