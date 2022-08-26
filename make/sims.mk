#!/usr/bin/make -f

###############################################################################
###                               Simulation                                ###
###############################################################################

SIM_TEST_DIR = ./app/simulation

SEED ?= 1
PERIOD ?= 5
NUM_BLOCKS ?= 50
BLOCK_SIZE ?= 200
GENESIS ?= ${HOME}/.regen/config/genesis.json

runsim:
	go install github.com/cosmos/tools/cmd/runsim@latest

sim-app:
	@echo "Running full app simulation..."
	@echo "Seed=$(SEED) Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) BlockSize=$(BLOCK_SIZE)"
	@go test $(SIM_TEST_DIR) -run TestFullAppSimulation -v -timeout 24h \
 		-Enabled=true \
 		-Commit=true \
		-Seed=$(SEED) \
 		-Period=$(PERIOD) \
		-NumBlocks=$(NUM_BLOCKS) \
		-BlockSize=$(BLOCK_SIZE)

sim-app-genesis:
	@echo "Running full app simulation with custom genesis..."
	@echo "Seed=$(SEED) Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) BlockSize=$(BLOCK_SIZE) Genesis=$(GENESIS)"
	@go test $(SIM_TEST_DIR) -run TestFullAppSimulation -v -timeout 24h \
		-Enabled=true \
		-Commit=true \
		-Seed=$(SEED) \
 		-Period=$(PERIOD) \
		-NumBlocks=$(NUM_BLOCKS) \
		-BlockSize=$(BLOCK_SIZE) \
		-Genesis=$(GENESIS)

sim-app-multi-seed: runsim
	@echo "Running full app simulation with multiple seeds..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS)"
	runsim -Jobs=4 -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail \
		$(NUM_BLOCKS) $(PERIOD) TestFullAppSimulation

sim-app-multi-seed-genesis: runsim
	@echo "Running full app simulation with multiple seeds and custom genesis..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) Genesis=$(GENESIS)"
	runsim -Jobs=4 -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail -Genesis=$(GENESIS) \
		$(NUM_BLOCKS) $(PERIOD) TestFullAppSimulation

sim-determinism:
	@echo "Running app state determinism simulation..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) BlockSize=$(BLOCK_SIZE)"
	@go test $(SIM_TEST_DIR) -run TestAppStateDeterminism -v -timeout 24h \
 		-Enabled=true \
		-Commit=true \
		-Period=$(PERIOD) \
		-NumBlocks=$(NUM_BLOCKS) \
		-BlockSize=$(BLOCK_SIZE)

sim-import-export: runsim
	@echo "Running import export simulation..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS)"
	runsim -Jobs=4 -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail \
		$(NUM_BLOCKS) $(PERIOD) TestAppImportExport

sim-after-import: runsim
	@echo "Running app after import simulation..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS)"
	runsim -Jobs=4 -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail \
		$(NUM_BLOCKS) $(PERIOD) TestAppSimulationAfterImport

.PHONY: runsim \
	sim-app sim-app-genesis sim-app-multi-seed sim-app-multi-seed-genesis \
	sim-determinism sim-import-export sim-after-import
