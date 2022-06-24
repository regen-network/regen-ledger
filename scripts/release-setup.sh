#!/usr/bin/env bash

# this script should only be called within github actions running goreleaser

sudo apt-get update
sudo apt-get install clang
sudo apt-get install gcc-multilib g++-multilib
#sudo apt-get install gcc-mingw-w64-x86-64 g++-mingw-w64-x86-64

mkdir /home/runner/work/osxcross
git clone https://github.com/likhita-809/osxcross-target.git /home/runner/work/osxcross/target
sudo ln -s /home/runner/work/osxcross/target/target/lib/libcrypto.so.1.0.0 /usr/lib/x86_64-linux-gnu/libcrypto.so.1.0.0

curl -o /home/runner/work/libwasmvm_muslc.a -L https://github.com/CosmWasm/wasmvm/releases/download/v1.0.0-beta5/libwasmvm_muslc.a
sha256sum /home/runner/work/libwasmvm_muslc.a | grep d16a2cab22c75dbe8af32265b9346c6266070bdcf9ed5aa9b7b39a7e32e25fe0
