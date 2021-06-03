#!/bin/sh

set -e

SUBMODULES=$(find . -type f -name 'go.mod' -print0 | xargs -0 -n1 dirname | sort)
PWD=$(pwd)
echo "mode: atomic" > coverage.txt

for m in ${SUBMODULES[@]}; do
    cd $PWD/$m
    PKGS=$(go list ./...)
    for pkg in ${PKGS[@]}; do
        go test -v -timeout 30m -race -coverprofile=profile.out -covermode=atomic -tags='ledger test_ledger_mock' "$pkg"
        if [ -f profile.out ]; then
            tail -n +2 profile.out >> coverage.txt;
            rm profile.out
        fi
    done
done