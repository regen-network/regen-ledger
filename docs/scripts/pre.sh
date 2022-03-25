#!/usr/bin/env bash

# module specification documentation

mkdir -p modules

for D in ../x/*; do
  if [ -d "${D}" ]; then
    rm -rf "modules/$(echo $D | awk -F/ '{print $NF}')"
    mkdir -p "modules/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/spec/* "$_"
  fi
done

cp MODULES.md modules/README.md

# regen app command-line documentation

rm -rf commands
mkdir -p commands
go run ../scripts/generate-cli-docs.go

cp COMMANDS.md commands/README.md
