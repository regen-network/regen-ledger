#!/usr/bin/env bash

set -eo pipefail

go install github.com/raviqqe/gherkin2markdown@latest

echo "Generating data feature docs..."

data_server_dir="docs/modules/data/features/server"
data_types_dir="docs/modules/data/features/types"

mkdir -p $data_server_dir
mkdir -p $data_types_dir

for file in $(find ./x/data -path -prune -o -name '*.feature'); do

  name=$(basename "${file%.*}")

  if [[ $file == *"server"* ]]; then
    gherkin2markdown "$file" > "$data_server_dir/$name.md"
  else
    gherkin2markdown "$file" > "$data_types_dir/$name.md"
  fi
done

echo "Generating ecocredit feature docs..."

eco_server_dir="docs/modules/ecocredit/features/server"
eco_types_dir="docs/modules/ecocredit/features/types"

mkdir -p $eco_server_dir
mkdir -p $eco_types_dir

for file in $(find ./x/ecocredit -path -prune -o -name '*.feature'); do

  name=$(basename "${file%.*}")

  if [[ $file == *"server"* ]]; then
    gherkin2markdown "$file" > "$eco_server_dir/$name.md"
  else
    gherkin2markdown "$file" > "$eco_types_dir/$name.md"
  fi

done
