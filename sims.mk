#!/usr/bin/make -f

########################################
### Simulations

SIMAPP = github.com/regen-network/regen-ledger/app

sim-xrn-nondeterminism:
	@echo "Running nondeterminism test..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -SimulationEnabled=true -v -timeout 10m

sim-xrn-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.xrnd/config/genesis.json will be used."
	@go test -mod=readonly github.com/regen-network/regen-ledger/app -run TestFullGaiaSimulation -SimulationGenesis=${HOME}/.xrnd/config/genesis.json \
		-SimulationEnabled=true -SimulationNumBlocks=100 -SimulationBlockSize=200 -SimulationCommit=true -SimulationSeed=99 -SimulationPeriod=5 -v -timeout 24h

sim-xrn-fast:
	@echo "Running quick Gaia simulation. This may take several minutes..."
	@go test -mod=readonly github.com/regen-network/regen-ledger/app -run TestFullGaiaSimulation -SimulationEnabled=true -SimulationNumBlocks=100 -SimulationBlockSize=200 -SimulationCommit=true -SimulationSeed=99 -SimulationPeriod=5 -v -timeout 24h

sim-xrn-import-export: runsim
	@echo "Running Gaia import/export simulation. This may take several minutes..."
	$(GOPATH)/bin/runsim 25 5 TestGaiaImportExport

sim-xrn-simulation-after-import: runsim
	@echo "Running Gaia simulation-after-import. This may take several minutes..."
	$(GOPATH)/bin/runsim 25 5 TestGaiaSimulationAfterImport

sim-xrn-custom-genesis-multi-seed: runsim
	@echo "Running multi-seed custom genesis simulation..."
	@echo "By default, ${HOME}/.xrnd/config/genesis.json will be used."
	$(GOPATH)/bin/runsim -g ${HOME}/.xrnd/config/genesis.json 400 5 TestFullGaiaSimulation

sim-xrn-multi-seed: runsim
	@echo "Running multi-seed Gaia simulation. This may take awhile!"
	$(GOPATH)/bin/runsim 400 5 TestFullGaiaSimulation

sim-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@go test -mod=readonly github.com/regen-network/regen-ledger/app -benchmem -bench=BenchmarkInvariants -run=^$ \
	-SimulationEnabled=true -SimulationNumBlocks=1000 -SimulationBlockSize=200 \
	-SimulationCommit=true -SimulationSeed=57 -v -timeout 24h

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_COMMIT ?= true
sim-xrn-benchmark:
	@echo "Running Gaia benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ github.com/regen-network/regen-ledger/app -bench ^BenchmarkFullGaiaSimulation$$  \
		-SimulationEnabled=true -SimulationNumBlocks=$(SIM_NUM_BLOCKS) -SimulationBlockSize=$(SIM_BLOCK_SIZE) -SimulationCommit=$(SIM_COMMIT) -timeout 24h

sim-xrn-profile:
	@echo "Running Gaia benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ github.com/regen-network/regen-ledger/app -bench ^BenchmarkFullGaiaSimulation$$ \
		-SimulationEnabled=true -SimulationNumBlocks=$(SIM_NUM_BLOCKS) -SimulationBlockSize=$(SIM_BLOCK_SIZE) -SimulationCommit=$(SIM_COMMIT) -timeout 24h -cpuprofile cpu.out -memprofile mem.out


.PHONY: runsim sim-xrn-nondeterminism sim-xrn-custom-genesis-fast sim-xrn-fast sim-xrn-import-export \
	sim-xrn-simulation-after-import sim-xrn-custom-genesis-multi-seed sim-xrn-multi-seed \
	sim-benchmark-invariants sim-xrn-benchmark sim-xrn-profile
