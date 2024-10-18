# Makefile for managing Regen DevNet Node operations

.PHONY: init config config-core config-node start all clean reset

# Paths and variables
BASE_PATH := /mnt/nvme
HOME_DIR := $(BASE_PATH)/.regen
MONIKER ?= regen-devnet-node

# Execute scripts for initializing, configuring, and starting the node
init:
	@echo "Initializing Regen node with moniker $(MONIKER)..."
	@./initchain.bash $(MONIKER)

config-node:
	@echo "Configuring Regen node with persistent peers: $(REGEN_P2P_PERSISTENT_PEERS)..."
	@REGEN_P2P_PERSISTENT_PEERS=$(REGEN_P2P_PERSISTENT_PEERS) ./setconfig-node.bash

start:
	@echo "Starting Regen node..."
	@./startnode.bash $(BASE_PATH)

# Combine all steps into a single make command
all: init config-node start

# Clean up data, logs, and temporary files
clean:
	@echo "Cleaning up temporary files..."
	@find . -name "*.log" -type f -delete
	@find . -name "*.tmp" -type f -delete

# Reset everything to start from scratch
reset:
	@echo "Removing all configurations and data to start from scratch..."
	@rm -rf $(HOME_DIR)
