#!/usr/bin/env bash
set -euo pipefail

# Script to update vendorHash in flake.nix when go.mod/go.sum changes
# This is needed when Go dependencies are updated

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$ROOT_DIR"

echo "🔄 Updating vendorHash in flake.nix..."

# First, set vendorHash to an empty string to get the correct hash
echo "🔧 Setting vendorHash to empty string to get correct hash..."

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/vendorHash = ".*"/vendorHash = ""/' flake.nix
    sed -i '' 's/vendorHash = null/vendorHash = ""/' flake.nix
else
    sed -i 's/vendorHash = ".*"/vendorHash = ""/' flake.nix
    sed -i 's/vendorHash = null/vendorHash = ""/' flake.nix
fi

echo "🔨 Building to get the correct vendorHash..."

# Capture the build output to extract the expected hash
BUILD_OUTPUT=$(nix build .#regen-ledger 2>&1 || true)

# Extract the correct hash from the error message
CORRECT_HASH=$(echo "$BUILD_OUTPUT" | grep -o "got: *sha256-[A-Za-z0-9+/=]*" | sed 's/got: *//' | head -1)

if [[ -z "$CORRECT_HASH" ]]; then
    echo "❌ Could not extract vendorHash from build output"
    echo "Build output:"
    echo "$BUILD_OUTPUT"
    exit 1
fi

echo "✅ Found correct vendorHash: $CORRECT_HASH"

# Update flake.nix with the correct hash
echo "🔧 Updating flake.nix with correct vendorHash..."

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' "s|vendorHash = \"\"|vendorHash = \"$CORRECT_HASH\"|" flake.nix
else
    sed -i "s|vendorHash = \"\"|vendorHash = \"$CORRECT_HASH\"|" flake.nix
fi

echo "✅ Updated vendorHash in flake.nix"

# Verify the build works now
echo "🔍 Verifying build with new vendorHash..."
if nix build .#regen-ledger --no-link; then
    echo "✅ Build successful with new vendorHash!"
else
    echo "❌ Build failed with new vendorHash"
    exit 1
fi

echo ""
echo "🎉 vendorHash update complete!"
echo "📋 New vendorHash: $CORRECT_HASH"
echo ""
echo "🚀 You can now build with:"
echo "  nix build .#regen-ledger --out-link build" 