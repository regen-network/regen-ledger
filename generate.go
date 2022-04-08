package regen_ledger

//go:generate mockgen -source=x/data/expected_keepers.go -package mocks -destination x/data/mocks/expected_keepers.go
//go:generate mockgen -source=x/ecocredit/expected_keepers.go -package mocks -destination x/ecocredit/mocks/expected_keepers.go
//go:generate mockgen -source=x/ecocredit/server/basket/keeper.go -package mocks -destination x/ecocredit/server/basket/mocks/keeper.go
//go:generate mockgen -source=x/ecocredit/server/core/keeper.go -package mocks -destination x/ecocredit/server/core/mocks/keeper.go
