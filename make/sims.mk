#!/usr/bin/make -f

###############################################################################
###                               Simulation                                ###
###############################################################################

SIM_TEST_DIR = ./app/simulation

JOBS ?= 4
SEED ?= 1
SEEDS ?= # https://github.com/cosmos/tools/blob/master/cmd/runsim/main.go#L32-L40
PERIOD ?= 5
NUM_BLOCKS ?= 100
BLOCK_SIZE ?= 200
GENESIS ?= ${HOME}/.regen/config/genesis.json

runsim:
	go install github.com/cosmos/tools/cmd/runsim@latest

sim-app:
	@echo "Running app simulation..."
	@echo "Seed=$(SEED) Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) BlockSize=$(BLOCK_SIZE)"
	@go test $(SIM_TEST_DIR) -run TestApp$$ -v -timeout 24h \
 		-Enabled=true \
 		-Commit=true \
		-Seed=$(SEED) \
 		-Period=$(PERIOD) \
		-NumBlocks=$(NUM_BLOCKS) \
		-BlockSize=$(BLOCK_SIZE)

sim-app-genesis:
	@echo "Running app simulation with custom genesis..."
	@echo "Seed=$(SEED) Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) BlockSize=$(BLOCK_SIZE) Genesis=$(GENESIS)"
	@go test $(SIM_TEST_DIR) -run TestApp$$ -v -timeout 24h \
		-Enabled=true \
		-Commit=true \
		-Seed=$(SEED) \
 		-Period=$(PERIOD) \
		-NumBlocks=$(NUM_BLOCKS) \
		-BlockSize=$(BLOCK_SIZE) \
		-Genesis=$(GENESIS)

sim-app-multi-seed: runsim
	@echo "Running app simulation with multiple seeds..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS)"
	runsim -Jobs=$(JOBS) -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail -Seeds $(SEEDS) \
		$(NUM_BLOCKS) $(PERIOD) TestApp

sim-app-multi-seed-genesis: runsim
	@echo "Running app simulation with multiple seeds and custom genesis..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) Genesis=$(GENESIS)"
	runsim -Jobs=$(JOBS) -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail -Seeds $(SEEDS) -Genesis=$(GENESIS) \
		$(NUM_BLOCKS) $(PERIOD) TestApp

sim-determinism:
	@echo "Running app state determinism simulation..."
	@echo "Seed=$(SEED) Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) BlockSize=$(BLOCK_SIZE)"
	@go test $(SIM_TEST_DIR) -run TestAppDeterminism -v -timeout 24h \
 		-Enabled=true \
		-Commit=true \
		-Seed=$(SEED) \
		-Period=$(PERIOD) \
		-NumBlocks=$(NUM_BLOCKS) \
		-BlockSize=$(BLOCK_SIZE)

sim-determinism-multi-seed: runsim
	@echo "Running app determinism simulation..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS)"
	runsim -Jobs=$(JOBS) -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail -Seeds $(SEEDS) \
		$(NUM_BLOCKS) $(PERIOD) TestAppDeterminism

sim-import-export:
	@echo "Running app state determinism simulation..."
	@echo "Seed=$(SEED) Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) BlockSize=$(BLOCK_SIZE)"
	@go test $(SIM_TEST_DIR) -run TestAppImportExport -v -timeout 24h \
 		-Enabled=true \
		-Commit=true \
		-Seed=$(SEED) \
		-Period=$(PERIOD) \
		-NumBlocks=$(NUM_BLOCKS) \
		-BlockSize=$(BLOCK_SIZE)

sim-import-export-multi-seed: runsim
	@echo "Running import export simulation..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS)"
	runsim -Jobs=$(JOBS) -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail -Seeds $(SEEDS) \
		$(NUM_BLOCKS) $(PERIOD) TestAppImportExport

sim-after-import:
	@echo "Running app state determinism simulation..."
	@echo "Seed=$(SEED) Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS) BlockSize=$(BLOCK_SIZE)"
	@go test $(SIM_TEST_DIR) -run TestAppAfterImport -v -timeout 24h \
 		-Enabled=true \
		-Commit=true \
		-Seed=$(SEED) \
		-Period=$(PERIOD) \
		-NumBlocks=$(NUM_BLOCKS) \
		-BlockSize=$(BLOCK_SIZE)

sim-after-import-multi-seed: runsim
	@echo "Running app after import simulation..."
	@echo "Period=$(PERIOD) NumBlocks=$(NUM_BLOCKS)"
	runsim -Jobs=$(JOBS) -SimAppPkg=$(SIM_TEST_DIR) -ExitOnFail -Seeds $(SEEDS) \
		$(NUM_BLOCKS) $(PERIOD) TestAppAfterImport

.PHONY: runsim sim-app sim-app-genesis sim-app-multi-seed sim-app-multi-seed-genesis \
	sim-determinism sim-determinism-multi-seed sim-import-export sim-import-export-multi-seed \
	sim-after-import sim-after-import-multi-seed
