#!/usr/bin/env bash

set -eo pipefail


protoc_install_regen_orm() {
  cd api/orm
  go install ./cmd/protoc-gen-go-regen-orm #2>/dev/null
}

protoc_install_regen_orm

cd ../..

echo "Generating gogo proto code"
cd proto
proto_dirs=$(find ./regen -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    if grep go_package $file &> /dev/null ; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

cd ..

# move proto files to the right places
cp -r github.com/regen-network/regen-ledger/* ./
rm -rf github.com

./scripts/protocgen-pulsar.sh
