#!/usr/bin/env bash

./scripts/post.sh

# specs directory

cp -r ../specs specs

# modules directory

mkdir -p modules

cp README_modules.md modules/README.md

for D in ../x/*; do
  if [ -d "${D}" ]; then
    rm -rf "modules/$(echo $D | awk -F/ '{print $NF}')"
    mkdir -p "modules/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/spec/* "$_"
  fi
done

(cd .. ; ./scripts/generate_feature_docs.sh)

# commands directory

mkdir -p commands

cp README_commands.md commands/README.md

go run ../scripts/generate_cli_docs.go
