# Nix Development Guide

This guide covers how to use Nix for regen-ledger development, providing reproducible builds and development environments across different platforms.

## Overview

Regen Ledger supports Nix flakes for:
- ✅ **Reproducible builds** 
- ✅ **Cross-platform development** (Linux x86_64/aarch64, macOS Apple Silicon)
- ✅ **Isolated development environments** with all required tools
- ✅ **Static linking** with wasmvm v1.5.0
- ✅ **No Docker required** for local development

## Prerequisites

### Install Nix

**macOS/Linux:**
```bash
curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install
```

**Enable flakes** (if using older Nix):
```bash
echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf
```

### Verify Installation
```bash
nix --version
# Should show: nix (Nix) 2.18+ 
```

## Quick Start

### 1. Clone and Build
```bash
git clone https://github.com/regen-network/regen-ledger
cd regen-ledger

# Build binary to result/bin/regen
nix build .#regen-ledger

# Or build directly to build/ directory
nix build .#regen-ledger --out-link build

# Test the binary
./result/bin/regen version
# or
./build/bin/regen version
```

### 2. Development Environment
```bash
# Enter development shell with all tools
nix develop

# Now you have access to:
# - Go toolchain (1.24+)
# - Protocol Buffers (protoc, protoc-gen-go)
# - Build tools (make, pkg-config)
# - CosmWasm Rust toolchain
# - Development utilities (jq, curl, git)
```

### 3. Run Without Installing
```bash
# Run regen-ledger directly from Nix store
nix run .#regen-ledger -- version
nix run .#regen-ledger -- --help
```

## Development Workflow

### Building

| Command | Output Location | Description |
|---------|----------------|-------------|
| `nix build .#regen-ledger` | `result/bin/regen` | Standard Nix build (symlink) |
| `nix build .#regen-ledger --out-link build` | `build/bin/regen` | Build to custom directory |
| `make build` | `build/regen` | Traditional Go build |

### When Go Dependencies Change

If you modify `go.mod` or run `go mod tidy`, update the Nix vendorHash:

```bash
# After go mod tidy or dependency changes
./scripts/update-vendor-hash.sh

# Then build normally
nix build .#regen-ledger --out-link build
```

### Development Shell

```bash
# Enter development environment
nix develop

# Available tools:
go version          # Go 1.24+
protoc --version    # Protocol Buffers
rustc --version     # Rust for CosmWasm
make --version      # Build system

# Use traditional make commands
make build          # Build with Go
make test           # Run tests
make proto-gen      # Generate protobuf files
```

### CosmWasm Verification

```bash
# Verify CosmWasm integration
./result/bin/regen query wasm --help
./result/bin/regen query wasm libwasmvm-version
# Should output: 1.5.0
```

## Build Configuration

### Platform Support
- ✅ **macOS Apple Silicon** (aarch64-darwin)
- ✅ **Linux x86_64** (x86_64-linux) 
- ✅ **Linux ARM64** (aarch64-linux)

### Build Tags
- **macOS**: `netgo,static_wasm` 
- **Linux**: `netgo,muslc,osusergo`

*Note: Ledger support currently disabled to avoid hidapi compilation issues*

### Static Linking
Nix builds use static wasmvm libraries:
- **macOS**: `libwasmvmstatic_darwin.a`
- **Linux**: `libwasmvm_muslc.{arch}.a`

This matches the production `.goreleaser.yml` configuration.

## Flake Structure

```nix
# Available packages
nix build .#regen-ledger    # Main binary package

# Available development shells  
nix develop                 # Full development environment

# Available apps
nix run .#regen-ledger      # Run regen-ledger directly
```

## Comparing Build Methods

### Nix vs Make

| Aspect | Nix Build | Make Build |
|--------|-----------|------------|
| **Dependencies** | Automatically managed | Manual installation required |
| **Reproducibility** | ✅ Pinned versions | ❌ System-dependent |
| **Cross-platform** | ✅ Same everywhere | ❌ Platform differences |
| **CosmWasm** | ✅ Static linking | ✅ Dynamic linking |
| **Build time** | ~2-3 minutes | ~1-2 minutes |
| **Binary size** | ~105MB | ~100MB |

### When to use each:

**Use Nix when:**
- Setting up new development environment
- Need reproducible builds
- Cross-platform development
- CI/CD pipelines
- Want isolated dependencies

**Use Make when:**
- Quick iterative development
- Already have Go toolchain installed
- Familiar with traditional Go development

## Troubleshooting

### Common Issues

**1. "experimental-features" error:**
```bash
# Enable flakes
echo "experimental-features = nix-command flakes" >> ~/.config/nix/nix.conf
```

**2. Build fails with "vendorHash mismatch":**
```bash
# Use the automatic script to update vendorHash
./scripts/update-vendor-hash.sh

# Or manually update flake.nix with the hash from error message
```

**3. "cannot create symlink 'result'":**
```bash
# Remove existing result directory/file
rm -rf result
nix build .#regen-ledger
```

**4. macOS: "command line tools" error:**
```bash
# Install Xcode command line tools
xcode-select --install
```

### Getting Help

**Check build logs:**
```bash
nix log .#regen-ledger
```

**Debug development shell:**
```bash
nix develop --verbose
```

**Verify flake:**
```bash
nix flake check
nix flake show
```

## Advanced Usage

### Cross-Platform Builds

```bash
# Build for Linux from macOS
nix build .#regen-ledger --system x86_64-linux
nix build .#regen-ledger --system aarch64-linux

# Note: Requires remote builders or emulation
```

### Custom Binary Location

```bash
# Build to custom location
nix build .#regen-ledger --out-link custom-build-name
./custom-build-name/bin/regen version
```

### Update Dependencies

```bash
# Update all flake inputs
nix flake update

# Update specific input
nix flake lock --update-input nixpkgs

# Commit the updated flake.lock
git add flake.lock
git commit -m "nix: update flake inputs"
```

### Update Go Dependencies (vendorHash)

When Go dependencies change (after `go mod tidy` or updating `go.mod`/`go.sum`), you need to regenerate the `vendorHash`:

```bash
# Automatically update vendorHash when Go dependencies change
./scripts/update-vendor-hash.sh
```

**When you need this:**
- After running `go mod tidy`
- When adding/removing Go dependencies
- When Nix build fails with "vendorHash mismatch" error
- After updating `go.mod` or `go.sum` files

**What the script does:**
1. Sets `vendorHash = ""` in `flake.nix`
2. Attempts Nix build (fails with correct hash)
3. Extracts the correct hash from error output
4. Updates `flake.nix` with correct vendorHash
5. Verifies build works with new hash

**Manual process (if script fails):**
```bash
# 1. Set vendorHash to empty string in flake.nix
vendorHash = "";

# 2. Try to build and copy the hash from error message
nix build .#regen-ledger
# Error will show: got: sha256-XXXXXXXXXX

# 3. Update flake.nix with the correct hash
vendorHash = "sha256-XXXXXXXXXX";
```

## Performance Tips

1. **Use binary cache:** Nix automatically uses cache.nixos.org
2. **Keep builds:** Don't delete `result/` symlinks to avoid rebuilds
3. **Parallel builds:** Nix builds dependencies in parallel automatically
4. **Garbage collection:** Run `nix-collect-garbage -d` occasionally

## Integration with IDEs

### VS Code
```bash
# Use in VS Code workspace
nix develop --command code .
```

### Vim/Neovim
```bash
# Use in editor
nix develop --command nvim
```

### Any Editor
```bash
# Enter development shell first
nix develop
# Then start your editor with all tools available
```

## Contributing

When modifying Nix configuration:

1. **Test on multiple platforms** if possible
2. **Update documentation** if changing commands
3. **Verify CosmWasm functionality** with test queries
4. **Check build reproducibility** with clean builds
5. **Update flake.lock** when adding new dependencies

For help with Nix configuration, see:
- [Nix Flakes Documentation](https://nixos.wiki/wiki/Flakes)
- [buildGoModule Reference](https://nixos.org/manual/nixpkgs/stable/#buildgomodule)
