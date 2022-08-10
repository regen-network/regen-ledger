#!/usr/bin/env bash

# this script should only be called within github actions running goreleaser

sudo apt-get update
sudo apt-get install clang
sudo apt-get install gcc-multilib g++-multilib
sudo apt-get install gcc-mingw-w64-x86-64 g++-mingw-w64-x86-64

mkdir /home/runner/work/osxcross
git clone https://github.com/likhita-809/osxcross-target.git /home/runner/work/osxcross/target
sudo ln -s /home/runner/work/osxcross/target/target/lib/libcrypto.so.1.0.0 /usr/lib/x86_64-linux-gnu/libcrypto.so.1.0.0
