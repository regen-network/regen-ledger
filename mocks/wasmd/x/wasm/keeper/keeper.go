package keeper

import (
	"github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Option struct{}

type Keeper struct{}

func NewKeeper(
	codec.Codec,
	sdk.StoreKey,
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
