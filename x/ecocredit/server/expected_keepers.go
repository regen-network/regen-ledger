package server

import (
	"encoding/json"

	"github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	marketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

var _ Keeper = serverImpl{}

// Keeper defines a set of methods the ecocredit module exposes.
type Keeper interface {
	PruneOrders(ctx sdk.Context) error
	RegisterInvariants(sdk.InvariantRegistry)
	InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) ([]types.ValidatorUpdate, error)
	ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) (json.RawMessage, error)
	QueryServers() (core.QueryServer, baskettypes.QueryServer, markettypes.QueryServer)
	GetStateStores() (api.StateStore, basketapi.StateStore, marketapi.StateStore)
}
