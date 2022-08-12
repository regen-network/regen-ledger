#!/usr/bin/env bash

set -e

SUBMODULES=$(find . -type f -name 'go.mod' -print0 | xargs -0 -n1 dirname | sort)
CURDIR=$(pwd)
echo "mode: atomic" > coverage.txt

for m in ${SUBMODULES[@]}; do
    cd $CURDIR/$m
    PKGS=$(go list ./...)
    for pkg in ${PKGS[@]}; do
        go test -v -timeout 30m -coverpkg=all -coverprofile=profile.out -covermode=atomic -tags="norace ledger test_ledger_mock" "$pkg"
        if [ -f profile.out ]; then
            tail -n +2 profile.out >> $CURDIR/coverage.txt;
            rm profile.out
        fi
    done
done

# filter out DONTCOVER, pb.go, pb.gw.go
cd $CURDIR
excludelist=" $(find ./ -type f -name '*.pb.go')"
excludelist+=" $(find ./ -type f -name '*.pb.gw.go')"
excludelist+=" $(find ./ -type f -name '*.pulsar.go')"    
excludelist+=" $(find ./ -type f -name '*.cosmos_orm.go')"
excludelist+="$(find ./ -type f -name '*.go' | xargs grep -l 'DONTCOVER')"
for filename in ${excludelist}; do
    filename=$(echo $filename | sed 's/^./github.com\/regen-network\/regen-ledger/g')
    echo "Excluding ${filename} from coverage report..."
    sed -i.bak "/$(echo $filename | sed 's/\//\\\//g')/d" coverage.txt
done
