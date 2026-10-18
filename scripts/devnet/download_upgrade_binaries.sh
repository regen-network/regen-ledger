#!/bin/bash
set -e

# Config
BUILD_DIR="./upgrade-binaries"
GITHUB_BASE="https://github.com/regen-network/regen-ledger/releases/download"

# Version/Asset info
V5_VERSION="v5.1.4"
V5_ASSET="regen-ledger_linux_amd64.zip"
V5_TARGET="regen-v5"

V6_VERSION="v6.0.0-rc4"
V6_ASSET="regen-ledger-6.0.0-rc4-linux-amd64.tar.gz"
V6_TARGET="regen-v6"

# Colors
GREEN='\033[0;32m'
NC='\033[0m'
INFO="${GREEN}ℹ️${NC}"
SUCCESS="${GREEN}✅${NC}"

log() { echo -e "$1 $2"; }

mkdir -p "$BUILD_DIR"

download_and_extract_zip() {
  local VERSION=$1
  local ZIP_NAME=$2
  local TARGET_NAME=$3
  local TMP_DIR
  TMP_DIR=$(mktemp -d)
  local URL="${GITHUB_BASE}/${VERSION}/${ZIP_NAME}"

  log "$INFO" "Downloading ${VERSION} from ${URL}..."
  curl -sL "$URL" -o "${TMP_DIR}/${ZIP_NAME}"

  log "$INFO" "Extracting ZIP..."
  unzip -q "${TMP_DIR}/${ZIP_NAME}" -d "$TMP_DIR"

  local BIN_PATH
  BIN_PATH=$(find "$TMP_DIR" -type f -name regen | head -n 1)

  if [ -z "$BIN_PATH" ]; then
    log "$INFO" "❌ regen binary not found in ZIP for ${VERSION}"
    exit 1
  fi

  mv "$BIN_PATH" "${BUILD_DIR}/${TARGET_NAME}"
  chmod +x "${BUILD_DIR}/${TARGET_NAME}"
  rm -rf "$TMP_DIR"

  log "$SUCCESS" "Saved ${TARGET_NAME} to ${BUILD_DIR}"
}

download_and_extract_tarball() {
  local VERSION=$1
  local TAR_NAME=$2
  local TARGET_NAME=$3
  local TMP_DIR
  TMP_DIR=$(mktemp -d)
  local URL="${GITHUB_BASE}/${VERSION}/${TAR_NAME}"

  log "$INFO" "Downloading ${VERSION} from ${URL}..."
  curl -sL "$URL" -o "${TMP_DIR}/${TAR_NAME}"

  log "$INFO" "Extracting TAR.GZ..."
  tar -xzf "${TMP_DIR}/${TAR_NAME}" -C "$TMP_DIR"

  log "$INFO" "Looking for regend binary..."
  local BIN_PATH
  BIN_PATH=$(find "$TMP_DIR" -type f -name regend | head -n 1)

  if [ -z "$BIN_PATH" ]; then
    log "$INFO" "❌ regend binary not found in TAR for ${VERSION}"
    exit 1
  fi

  mv "$BIN_PATH" "${BUILD_DIR}/${TARGET_NAME}"
  chmod +x "${BUILD_DIR}/${TARGET_NAME}"
  rm -rf "$TMP_DIR"

  log "$SUCCESS" "Saved ${TARGET_NAME} to ${BUILD_DIR}"
}

# Run
download_and_extract_zip "$V5_VERSION" "$V5_ASSET" "$V5_TARGET"
download_and_extract_tarball "$V6_VERSION" "$V6_ASSET" "$V6_TARGET"

log "$SUCCESS" "✅ All done. Binaries available in ${BUILD_DIR}:"
ls -lh "$BUILD_DIR"
