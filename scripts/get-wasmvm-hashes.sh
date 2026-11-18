#!/usr/bin/env bash

set -o pipefail  # Remove strict error handling that might cause early exits

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
WASMVM_VERSION="v1.5.0"
BASE_URL="https://github.com/CosmWasm/wasmvm/releases/download/${WASMVM_VERSION}"
TEMP_DIR="$(mktemp -d)"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Platform configurations - we fetch ALL platforms regardless of current system
declare -A PLATFORMS=(
    ["x86_64-linux"]="libwasmvm_muslc.x86_64.a"
    ["aarch64-linux"]="libwasmvm_muslc.aarch64.a"
    ["aarch64-darwin"]="libwasmvmstatic_darwin.a"
)

echo -e "${BLUE}🔍 Fetching wasmvm library hashes for ALL target platforms${NC}"
echo -e "${BLUE}Version: ${WASMVM_VERSION}${NC}"
echo -e "${BLUE}Platforms: x86_64-linux, aarch64-linux, aarch64-darwin${NC}"
echo -e "${YELLOW}Note: Downloads all platforms regardless of current system${NC}"
echo

# Function to download and hash a file
download_and_hash() {
    local platform="$1"
    local filename="$2"
    local url="${BASE_URL}/${filename}"
    local filepath="${TEMP_DIR}/${filename}"
    
    echo -e "${YELLOW}📥 Downloading ${filename} for ${platform}...${NC}"
    
    if curl -L -f -o "$filepath" "$url" 2>/dev/null; then
        local hash=$(sha256sum "$filepath" | cut -d' ' -f1)
        echo -e "${GREEN}✅ ${platform}: ${hash}${NC}"
        
        # Store result for later processing
        echo "${platform}|${url}|${hash}" >> "${TEMP_DIR}/platform-data.txt"
        
        return 0
    else
        echo -e "${RED}❌ Failed to download ${filename} from ${url}${NC}"
        echo -e "${RED}    This might be due to network issues or the file not existing${NC}"
        return 1
    fi
}

# Initialize data file
echo -n > "${TEMP_DIR}/platform-data.txt"

# Download and hash each platform
success_count=0
total_count=${#PLATFORMS[@]}
failed_platforms=()

echo -e "${BLUE}Starting downloads for ${total_count} platforms...${NC}"
echo

for platform in "${!PLATFORMS[@]}"; do
    filename="${PLATFORMS[$platform]}"
    
    download_and_hash "$platform" "$filename"
    download_result=$?
    
    if [ $download_result -eq 0 ]; then
        ((success_count++))
        echo -e "${GREEN}✓ ${platform} completed successfully${NC}"
    else
        failed_platforms+=("$platform")
        echo -e "${RED}✗ ${platform} failed${NC}"
    fi
    echo
done

# Generate properly formatted Nix expression from collected data
generate_nix_expression() {
    local data_file="$1"
    local output_file="$2"
    
    cat > "$output_file" << EOF
# wasmvm library hashes for ${WASMVM_VERSION}
{ pkgs, system }:

{
  libwasmvm = pkgs.fetchurl (
EOF

    # Process each platform
    local first=true
    while IFS='|' read -r platform url hash; do
        if [ "$first" = true ]; then
            echo "    if system == \"$platform\" then {" >> "$output_file"
            first=false
        else
            echo "    } else if system == \"$platform\" then {" >> "$output_file"
        fi
        echo "      url = \"$url\";" >> "$output_file"
        echo "      sha256 = \"$hash\";" >> "$output_file"
    done < "$data_file"
    
    cat >> "$output_file" << 'EOF'
    } else throw "Unsupported system: ${system}"
  );
}
EOF
}

echo -e "${BLUE}📋 Summary${NC}"
echo -e "Successfully processed: ${GREEN}${success_count}/${total_count}${NC} libraries"

if [ ${#failed_platforms[@]} -gt 0 ]; then
    echo -e "${RED}Failed platforms: ${failed_platforms[*]}${NC}"
fi
echo

if [ $success_count -gt 0 ]; then
    if [ $success_count -eq $total_count ]; then
        echo -e "${GREEN}🎉 All hashes calculated successfully!${NC}"
    else
        echo -e "${YELLOW}⚠️  Partial success: ${success_count}/${total_count} hashes calculated${NC}"
    fi
    echo
    
    # Generate the Nix expression
    nix_output="${TEMP_DIR}/nix-hashes.txt"
    generate_nix_expression "${TEMP_DIR}/platform-data.txt" "$nix_output"
    
    echo -e "${BLUE}📄 Generated Nix flake snippet:${NC}"
    echo "----------------------------------------"
    cat "$nix_output"
    echo "----------------------------------------"
    echo
    
    # Create nix directory if it doesn't exist
    nix_dir="${SCRIPT_DIR}/../nix"
    mkdir -p "$nix_dir"
    
    # Save to nix directory
    output_file="${nix_dir}/wasmvm-hashes.nix"
    cp "$nix_output" "$output_file"
    echo -e "${GREEN}💾 Nix expression saved to: ${output_file}${NC}"
    
else
    echo -e "${RED}❌ No downloads succeeded. Check your internet connection and try again.${NC}"
    exit 1
fi

# Final status
if [ $success_count -eq $total_count ]; then
    echo -e "${GREEN}🎯 All platforms completed successfully!${NC}"
    exit 0
elif [ $success_count -gt 0 ]; then
    echo -e "${YELLOW}⚠️  Partial completion. You may need to manually fetch the missing platforms.${NC}"
    exit 0
else
    exit 1
fi

# Cleanup
rm -rf "$TEMP_DIR"
echo -e "${BLUE}🧹 Cleaned up temporary files${NC}" 