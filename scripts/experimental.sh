#!/usr/bin/env bash

search="replace github.com/CosmWasm/wasmd => ./mock/wasmd"
replace="// replace directive removed as a result of experimental.sh"

sed -i "s|$search|$replace|" go.mod

go mod tidy
