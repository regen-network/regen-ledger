package server

import (
	"encoding/json"

	"github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	marketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	basekeeper "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/keeper"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
	basketkeeper "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/keeper"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/basket/types/v1"
	marketkeeper "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/keeper"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
)

var _ Keeper = serverImpl{}

// Keeper defines a set of methods the ecocredit module exposes.
type Keeper interface {
	PruneOrders(ctx sdk.Context) error
	RegisterInvariants(sdk.InvariantRegistry)
	InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) ([]types.ValidatorUpdate, error)
	ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) (json.RawMessage, error)
	QueryServers() (basetypes.QueryServer, baskettypes.QueryServer, markettypes.QueryServer)
	GetStateStores() (api.StateStore, basketapi.StateStore, marketapi.StateStore)

	GetBaseKeeper() basekeeper.Keeper
	GetBasketKeeper() basketkeeper.Keeper
	GetMarketKeeper() marketkeeper.Keeper
}
