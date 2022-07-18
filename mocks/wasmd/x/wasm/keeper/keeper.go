package keeper

import (
	"github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Option struct{}

type Keeper struct{}

func NewKeeper(
	codec.Codec,
	storetypes.StoreKey,
	paramtypes.Subspace,
	types.AccountKeeper,
	types.BankKeeper,
	types.StakingKeeper,
	types.DistributionKeeper,
	types.ChannelKeeper,
	types.PortKeeper,
	types.CapabilityKeeper,
	types.ICS20TransferPortSource,
	MessageRouter,
	GRPCQueryRouter,
	string,
	types.WasmConfig,
	string,
	...Option,
) Keeper {
	return Keeper{}
}
