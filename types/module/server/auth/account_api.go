package auth

import (
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/module/server"
)

type accountAPI struct {
	authkeeper.AccountKeeper
}

var _ server.AccountAPI = accountAPI{}

func (a accountAPI) CreateNewAccount(moduleKey server.ModuleKey, ctx types.Context) error {
	panic("implement me")
}

func (a accountAPI) EnsureAccountExists(moduleKey server.ModuleKey, ctx types.Context) error {
	panic("implement me")
}

func (a accountAPI) AccountExists(moduleKey server.ModuleKey, ctx types.Context) bool {
	panic("implement me")
}
