#!/bin/sh

set -e
rm -rf completions
mkdir completions

for sh in bash zsh fish; do
	go run cmd/regen/main.go completion "$sh" >"completions/regen.$sh"
done
