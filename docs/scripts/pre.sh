#!/usr/bin/env bash

# general specification documentation

rm -r specs

mkdir specs

cp README_specs.md specs/README.md

cp -r ../rfcs specs/rfcs

# module specification documentation

mkdir -p modules

cp README_modules.md modules/README.md

for D in ../x/*; do
  if [ -d "${D}" ]; then
    rm -rf "modules/$(echo $D | awk -F/ '{print $NF}')"
    mkdir -p "modules/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/spec/* "$_"
  fi
done

# regen app command-line documentation

mkdir -p commands

cp README_commands.md commands/README.md

go run ../scripts/generate_cli_docs.go

# module feature documentation

(cd .. ; ./scripts/generate_feature_docs.sh)
