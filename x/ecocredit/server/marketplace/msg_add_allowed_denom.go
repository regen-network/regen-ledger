package marketplace

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) AddAllowedDenom(ctx context.Context, req *marketplace.MsgAddAllowedDenom) (*marketplace.MsgAddAllowedDenomResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "expected %s got %s", k.authority, req.Authority)
	}

	if err := k.stateStore.AllowedDenomTable().Insert(ctx, &api.AllowedDenom{
		BankDenom:    req.BankDenom,
		DisplayDenom: req.DisplayDenom,
		Exponent:     req.Exponent,
	}); err != nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not add denom %s: %s", req.BankDenom, err.Error())
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(&marketplace.EventAllowDenom{Denom: req.BankDenom})

	return &marketplace.MsgAddAllowedDenomResponse{}, nil
}
