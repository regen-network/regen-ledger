package group

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/group/server"
	"github.com/regen-network/regen-ledger/x/group/types"
)

// ExportGenesis returns a GenesisState for a given context and Keeper.
func ExportGenesis(ctx sdk.Context, k *server.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}
