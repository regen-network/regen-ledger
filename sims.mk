#!/usr/bin/make -f

########################################
### Simulations

sim-regen-nondeterminism:
	@echo "Running nondeterminism test..."
	@go test -mod=readonly $(APP_DIR) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -v -timeout 24h

sim-regen-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.regen/config/genesis.json will be used."
	@go test -mod=readonly $(APP_DIR) -run TestFullAppSimulation -Genesis=${HOME}/.regen/config/genesis.json \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

sim-regen-fast:
	@echo "Running quick Gaia simulation. This may take several minutes..."
	@go test -mod=readonly $(APP_DIR) -run TestFullAppSimulation -Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

sim-regen-import-export: runsim
	@echo "Running Regen import/export simulation. This may take several minutes..."
	$(GOPATH)/bin/runsim -Jobs=4 -SimAppPkg=$(APP_DIR) -ExitOnFail 25 5 TestImportExport

sim-regen-simulation-after-import: runsim
	@echo "Running application simulation-after-import. This may take several minutes..."
	$(GOPATH)/bin/runsim -Jobs=4 -SimAppPkg=$(APP_DIR) -ExitOnFail 50 5 TestAppSimulationAfterImport

sim-regen-custom-genesis-multi-seed: runsim
	@echo "Running multi-seed custom genesis simulation..."
	@echo "By default, ${HOME}/.regen/config/genesis.json will be used."
	$(GOPATH)/bin/runsim -Genesis=${HOME}/.regen/config/genesis.json -SimAppPkg=$(APP_DIR) -ExitOnFail 400 5 TestFullAppSimulation

sim-regen-multi-seed: runsim
	@echo "Running multi-seed application simulation. This may take awhile!"
	$(GOPATH)/bin/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 500 50 TestFullAppSimulation

sim-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@go test -mod=readonly $(APP_DIR) -benchmem -bench=BenchmarkInvariants -run=^$ \
	-Enabled=true -NumBlocks=1000 -BlockSize=200 \
	-Commit=true -Seed=57 -v -timeout 24h

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_COMMIT ?= true

sim-regen-benchmark:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(APP_DIR) -bench ^BenchmarkFullAppSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h

sim-regen-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(APP_DIR) -bench ^BenchmarkFullAppSimulation$$ \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -cpuprofile cpu.out -memprofile mem.out


.PHONY: runsim sim-regen-nondeterminism sim-regen-custom-genesis-fast sim-regen-fast sim-regen-import-export \
	sim-regen-simulation-after-import sim-regen-custom-genesis-multi-seed sim-regen-multi-seed \
	sim-benchmark-invariants sim-regen-benchmark sim-regen-profile
