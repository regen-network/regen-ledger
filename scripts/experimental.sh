#!/usr/bin/env bash

echo "Updating go.mod for experimental build..."

search="replace github.com/CosmWasm/wasmd => ./mocks/wasmd"
replace="// replace github.com/CosmWasm/wasmd => ./mocks/wasmd"

# using -i.bak makes this compatible with both GNU and BSD/Mac
sed -i.bak "s|$search|$replace|" go.mod

go mod tidy
