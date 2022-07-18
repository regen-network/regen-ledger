#!/usr/bin/env bash

set -eo pipefail

export PATH=~/go/bin:$PATH

go install github.com/raviqqe/gherkin2markdown@latest

echo "Generating data feature docs..."

data_path="modules/data/features"
data_server_dir="docs/$data_path/server"
data_types_dir="docs/$data_path/types"

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
    data_readme+="- [$name]($data_path/server/$name.html)\n"
    gherkin2markdown "$file" > "$data_server_dir/$name.md"
  else
    if [[ $data_types_count == 0 ]]; then
      data_readme+="## Types\n\n"
    fi
    data_types_count+=1
    data_readme+="- [$name]($data_path/types/$name.html)\n"
    gherkin2markdown "$file" > "$data_types_dir/$name.md"
  fi
done

echo -e "$data_readme" > "docs/$data_path/README.md"

echo "Generating ecocredit feature docs..."

eco_path="modules/ecocredit/features"
eco_server_dir="docs/$eco_path/server"
eco_types_dir="docs/$eco_path/types"

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
    eco_readme+="- [$name]($eco_path/server/$name.html)\n"
    gherkin2markdown "$file" > "$eco_server_dir/$name.md"
  else
    if [[ $eco_types_count == 0 ]]; then
      eco_readme+="## Types\n\n"
    fi
    eco_types_count+=1
    eco_readme+="- [$name]($eco_path/types/$name.html)\n"
    gherkin2markdown "$file" > "$eco_types_dir/$name.md"
  fi
done

echo -e "$eco_readme" > "docs/$eco_path/README.md"
