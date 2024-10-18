#!/bin/bash
# Starts the Regen node.

set -uxe

# Set minimum gas prices
MINIMUM_GAS_PRICES="0.025uregen"
export HOME_DIR=/mnt/nvme/.regen

# Start the node with the specified minimum gas prices
regen start --home "$HOME_DIR" --minimum-gas-prices="$MINIMUM_GAS_PRICES"
