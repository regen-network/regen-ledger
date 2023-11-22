package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/capability/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
)

//go:generate mockgen -source=expected_keepers.go -package mocks -destination mocks/expected_keepers.go

type CapabilityKeeper interface {
	ClaimCapability(ctx sdk.Context, cap *types.Capability, name string) error
	GetCapability(ctx sdk.Context, name string) (*types.Capability, bool)
}

type ICAControllerKeeper interface {
	RegisterInterchainAccount(ctx sdk.Context, connectionID, owner, version string) error
	GetActiveChannelID(ctx sdk.Context, connectionID, portID string) (string, bool)
	SendTx(ctx sdk.Context, chanCap *types.Capability, connectionID, portID string, icaPacketData icatypes.InterchainAccountPacketData, timeoutTimestamp uint64) (uint64, error)
	GetInterchainAccountAddress(ctx sdk.Context, connectionID string, portID string) (string, bool)
}
