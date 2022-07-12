#!/usr/bin/env bash

echo "Cleaning up go.mod from experimental build..."

search="// replace github.com/CosmWasm/wasmd => ./mocks/wasmd"
replace="replace github.com/CosmWasm/wasmd => ./mocks/wasmd"

# using -i.bak makes this compatible with both GNU and BSD/Mac
sed -i.bak "s|$search|$replace|" go.mod

go mod tidy

rm go.mod.bak
