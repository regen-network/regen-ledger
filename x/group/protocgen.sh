#!/usr/bin/env bash

set -eo pipefail

protoc \
-I=. \
-I=$(go list -f "{{ .Dir }}" -m github.com/gogo/protobuf) \
-I=$(go list -f "{{ .Dir }}" -m github.com/cosmos/cosmos-sdk) \
--gocosmos_out=\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=interfacetype+grpc,paths=source_relative:. \
types.proto

# generate testdata types
protoc \
-I=.. \
-I=. \
-I=$(go list -f "{{ .Dir }}" -m github.com/cosmos/cosmos-sdk) \
-I=$(go list -f "{{ .Dir }}" -m github.com/gogo/protobuf) \
--gocosmos_out=\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=interfacetype+grpc,paths=source_relative:. \
testdata/types.proto
