#!/usr/bin/env bash

set -eo pipefail

# NOTE: fork of github.com/raviqqe/gherkin2markdown with Rule support
go install github.com/ryanchristo/gherkin2markdown@latest

echo "Generating data feature docs..."

data_dir="docs/modules/data/features"
data_server_dir="$data_dir/server"
data_types_dir="$data_dir/types"

mkdir -p $data_server_dir
mkdir -p $data_types_dir

data_readme="# Features\n\n"
data_readme+="## Types\n\n"

for file in $(find ./x/data/features -path -prune -o -name '*.feature' | sort -t '\0' -n); do
  name=$(basename "${file%.*}")
  data_readme+="- [$name](./types/$name.html)\n"
  "$GOBIN/gherkin2markdown" "$file" > "$data_types_dir/$name.md"
done

data_readme+="\n## Server\n\n"

for file in $(find ./x/data/server -path -prune -o -name '*.feature' | sort -t '\0' -n); do
  name=$(basename "${file%.*}")
  data_readme+="- [$name](./server/$name.html)\n"
  "$GOBIN/gherkin2markdown" "$file" > "$data_server_dir/$name.md"
done

echo -e "$data_readme" > "$data_dir/README.md"

echo "Generating ecocredit feature docs..."

eco_dir="docs/modules/ecocredit/features"
eco_server_dir="$eco_dir/server"
eco_types_dir="$eco_dir/types"

mkdir -p $eco_server_dir
mkdir -p $eco_types_dir

eco_readme="# Features\n\n"
eco_readme+="## Types\n\n"
eco_readme+="### Base\n\n"

for file in $(find ./x/ecocredit/base/types/v1/features -path -prune -o -name '*.feature' | sort -t '\0' -n); do
  name=$(basename "${file%.*}")
  eco_readme+="- [$name](./types/$name.html)\n"
  "$GOBIN/gherkin2markdown" "$file" > "$eco_types_dir/$name.md"
done

eco_readme+="### Basket\n\n"

for file in $(find ./x/ecocredit/basket/types/v1/features -path -prune -o -name '*.feature' | sort -t '\0' -n); do
  name=$(basename "${file%.*}")
  eco_readme+="- [$name](./types/$name.html)\n"
  "$GOBIN/gherkin2markdown" "$file" > "$eco_types_dir/$name.md"
done

eco_readme+="### Marketplace\n\n"

for file in $(find ./x/ecocredit/marketplace/types/v1/features -path -prune -o -name '*.feature' | sort -t '\0' -n); do
  name=$(basename "${file%.*}")
  eco_readme+="- [$name](./types/$name.html)\n"
  "$GOBIN/gherkin2markdown" "$file" > "$eco_types_dir/$name.md"
done

eco_readme+="## Server\n\n"
eco_readme+="### Base\n\n"

for file in $(find ./x/ecocredit/base/keeper/features -path -prune -o -name '*.feature' | sort -t '\0' -n); do
  name=$(basename "${file%.*}")
  eco_readme+="- [$name](./server/$name.html)\n"
  "$GOBIN/gherkin2markdown" "$file" > "$eco_server_dir/$name.md"
done

eco_readme+="### Basket\n\n"

for file in $(find ./x/ecocredit/basket/keeper/features -path -prune -o -name '*.feature' | sort -t '\0' -n); do
  name=$(basename "${file%.*}")
  eco_readme+="- [$name](./server/$name.html)\n"
  "$GOBIN/gherkin2markdown" "$file" > "$eco_server_dir/$name.md"
done

eco_readme+="### Marketplace\n\n"

for file in $(find ./x/ecocredit/marketplace/keeper/features -path -prune -o -name '*.feature' | sort -t '\0' -n); do
  name=$(basename "${file%.*}")
  eco_readme+="- [$name](./server/$name.html)\n"
  "$GOBIN/gherkin2markdown" "$file" > "$eco_server_dir/$name.md"
done

echo -e "$eco_readme" > "$eco_dir/README.md"
