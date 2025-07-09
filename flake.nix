{
  description = "Regen Ledger - Cosmos SDK blockchain with CosmWasm smart contracts for ecological claims";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux" 
      "aarch64-darwin"
    ] (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        
        # VERSIONS for regen-ledger and wasmvm
        version = "v6.0.0";
        wasmvmVersion = "v1.5.0";
        
        # Import wasmvm library hashes from generated file
        wasmvmHashes = import ./nix/wasmvm-hashes.nix { inherit pkgs system; };
        libwasmvm = wasmvmHashes.libwasmvm;

        # Build regen-ledger with CosmWasm support
        regen-ledger = pkgs.buildGoModule rec {
          pname = "regen-ledger";
          inherit version;

          src = ./.;

          # Go module hash
          vendorHash = "sha256-4Xwrk/q7hSjZdmMjmObKZNeKQMwPle8fOP/mydjeAbk=";

          # Go version requirement from go.mod
          nativeBuildInputs = with pkgs; [
            go  
            pkg-config
          ];

          buildInputs = with pkgs; [
            # Required for CosmWasm/wasmvm CGO compilation
          ] ++ pkgs.lib.optionals pkgs.stdenv.isLinux [
            musl.dev 
            pkgsStatic.stdenv.cc
          ];

          # Critical: Set up wasmvm library for linking based on .goreleaser.yml
          preBuild = ''
            # Create lib directory and copy appropriate wasmvm library
            mkdir -p lib
            ${if system == "aarch64-darwin" then ''
              cp ${libwasmvm} lib/libwasmvmstatic_darwin.a
              export CGO_ENABLED=1
              export CGO_LDFLAGS="-L$PWD/lib"
            '' else ''
              mkdir -p /tmp/usr/lib/${if system == "x86_64-linux" then "x86_64-linux-gnu" else "aarch64-linux-gnu"}
              cp ${libwasmvm} /tmp/usr/lib/${if system == "x86_64-linux" then "x86_64-linux-gnu" else "aarch64-linux-gnu"}/libwasmvm_muslc.a
              export CGO_ENABLED=1
            ''}
          '';

          # Build tags for CosmWasm support - disable ledger initially to avoid hidapi issues
          tags = if system == "aarch64-darwin" then 
            [ "netgo" "static_wasm" ]
          else
            [ "netgo" "muslc" "osusergo" ];

          # Build the main binary
          subPackages = [ "./cmd/regen" ];

          # Add ldflags based on Makefile for version information
          ldflags = [
            "-X github.com/cosmos/cosmos-sdk/version.Name=regen"
            "-X github.com/cosmos/cosmos-sdk/version.AppName=regen-ledger"
            "-X github.com/cosmos/cosmos-sdk/version.Version=${version}"
            "-X github.com/cosmos/cosmos-sdk/version.Commit=nix-build"
            "-X github.com/cosmos/cosmos-sdk/version.BuildTags=${pkgs.lib.concatStringsSep "," (if system == "aarch64-darwin" then [ "netgo" "static_wasm" ] else [ "netgo" "muslc" "osusergo" ])}"
            "-w" "-s"  # Strip debug info
          ] ++ pkgs.lib.optionals pkgs.stdenv.isLinux [
            "-linkmode=external"
            "-extldflags=-static"
          ] ++ pkgs.lib.optionals pkgs.stdenv.isDarwin [
            "-linkmode=external"
          ];

         

          # Disable tests that require network or specific hardware
          doCheck = false;

          meta = with pkgs.lib; {
            description = "Regen Ledger - Cosmos SDK blockchain for ecological claims with CosmWasm smart contracts";
            homepage = "https://github.com/regen-network/regen-ledger";
            license = licenses.asl20;
            maintainers = [  ];
            platforms = [ "x86_64-linux" "aarch64-linux" "aarch64-darwin" ];
          };
        };

      in {
        packages = {
          default = regen-ledger;
          regen-ledger = regen-ledger;
        };

        # Development shell with all tools
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            # Go toolchain
            go  # Use current stable Go version (meets 1.21+ requirement)
            gopls
            gotools
            go-tools

            # Build tools
            gnumake
            pkg-config
            
            # Protocol Buffers
            protobuf
            protoc-gen-go
            protoc-gen-go-grpc
            
            # CosmWasm development
            rustc
            cargo
            
            # Development utilities
            jq
            git
            curl
            
            # Platform-specific libraries
          ] ++ pkgs.lib.optionals pkgs.stdenv.isLinux [
            glibc.dev
          ];

          shellHook = ''
            echo "🌱 Regen Ledger Development Environment"
            echo "Platform: ${system}"
            echo "Go version: $(go version)"
            echo "CosmWasm wasmvm: ${wasmvmVersion}"
            echo ""
            echo "Available commands:"
            echo "  nix build .#regen-ledger    - Build regen-ledger binary"
            echo "  nix develop                 - Enter development shell"
            echo "  nix run .#regen-ledger      - Run regen-ledger"
            echo "  make build                  - Build using Makefile"
            echo "  make test                   - Run test suite"
            echo "  make proto-gen              - Generate protobuf files"
            echo ""
            
            # Set up wasmvm library for development
            export CGO_ENABLED=1
            ${if system == "aarch64-darwin" then ''
              export DYLD_LIBRARY_PATH="${pkgs.lib.makeLibraryPath []}:''${DYLD_LIBRARY_PATH:-}"
            '' else ''
              export LD_LIBRARY_PATH="${pkgs.lib.makeLibraryPath []}:''${LD_LIBRARY_PATH:-}"
            ''}
          '';
        };

        # Apps for running binaries
        apps = {
          default = flake-utils.lib.mkApp {
            drv = regen-ledger;
            exePath = "/build/regen";
          };
          
          regen-ledger = flake-utils.lib.mkApp {
            drv = regen-ledger;
            exePath = "/build/regen";
          };
        };
      });
}