#!/usr/bin/env bash

set -eo pipefail

SWAGGER_DIR=./client/docs

# find all proto directories
proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

# loop through proto directories
for dir in $proto_dirs; do

  # generate swagger files for query service proto files
  query_file=$(find "${dir}" -maxdepth 2 -name 'query.proto')
  if [[ ! -z "$query_file" ]]; then
    buf alpha protoc  \
      -I "proto" \
      -I "third_party/proto" \
      "$query_file" \
      --swagger_out=${SWAGGER_DIR} \
      --swagger_opt=logtostderr=true \
      --swagger_opt=fqn_for_swagger_name=true \
      --swagger_opt=simple_operation_ids=true
  fi
done
