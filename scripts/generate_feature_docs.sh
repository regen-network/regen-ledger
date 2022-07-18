#!/usr/bin/env bash

set -eo pipefail

go install github.com/raviqqe/gherkin2markdown@latest

echo "Generating data feature docs..."

data_dir="docs/modules/data/features"
data_server_dir="$data_dir/server"
data_types_dir="$data_dir/types"

mkdir -p $data_server_dir
mkdir -p $data_types_dir

data_readme="# Data Features\n\n"

data_server_count=0
data_types_count=0

for file in $(find ./x/data -path -prune -o -name '*.feature'); do

  name=$(basename "${file%.*}")

  if [[ $file == *"server"* ]]; then
    if [[ $data_server_count == 0 ]]; then
      data_readme+="\n## Server\n\n"
    fi
    data_server_count+=1
    data_readme+="- [$name](./server/$name.html)\n"
    "$GOBIN/gherkin2markdown" "$file" > "$data_server_dir/$name.md"
  else
    if [[ $data_types_count == 0 ]]; then
      data_readme+="## Types\n\n"
    fi
    data_types_count+=1
    data_readme+="- [$name](./types/$name.html)\n"
    "$GOBIN/gherkin2markdown" "$file" > "$data_types_dir/$name.md"
  fi
done

echo -e "$data_readme" > "$data_dir/README.md"

echo "Generating ecocredit feature docs..."

eco_dir="docs/modules/ecocredit/features"
eco_server_dir="$eco_dir/server"
eco_types_dir="$eco_dir/types"

mkdir -p $eco_server_dir
mkdir -p $eco_types_dir

eco_readme="# Ecocredit Features\n\n"

eco_server_count=0
eco_types_count=0

for file in $(find ./x/ecocredit -path -prune -o -name '*.feature'); do
  name=$(basename "${file%.*}")
  if [[ $file == *"server"* ]]; then
    if [[ $eco_server_count == 0 ]]; then
      eco_readme+="\n## Server\n\n"
    fi
    eco_server_count+=1
    eco_readme+="- [$name](./server/$name.html)\n"
    "$GOBIN/gherkin2markdown" "$file" > "$eco_server_dir/$name.md"
  else
    if [[ $eco_types_count == 0 ]]; then
      eco_readme+="## Types\n\n"
    fi
    eco_types_count+=1
    eco_readme+="- [$name](./types/$name.html)\n"
    "$GOBIN/gherkin2markdown" "$file" > "$eco_types_dir/$name.md"
  fi
done

echo -e "$eco_readme" > "$eco_dir/README.md"
