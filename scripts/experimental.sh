#!/usr/bin/env bash

echo "Updating go.mod and go.sum to support experimental build..."

search="replace github.com/CosmWasm/wasmd => ./mock/wasmd"
replace="// replace directive removed as a result of experimental.sh"

sed -i "s|$search|$replace|" go.mod

go mod tidy
