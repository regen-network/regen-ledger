#!/usr/bin/env bash

# module specification documentation

mkdir -p modules

cp MODULES.md modules/README.md

for D in ../x/*; do
  if [ -d "${D}" ]; then
    rm -rf "modules/$(echo $D | awk -F/ '{print $NF}')"
    mkdir -p "modules/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/spec/* "$_"
  fi
done

# regen app command-line documentation

mkdir -p commands

cp COMMANDS.md commands/README.md

go run ../scripts/generate-cli-docs.go
