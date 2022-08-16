package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// AddAllowedDenom adds a denom to the list of approved denoms that may be used in the
// marketplace.
func (k Keeper) AddAllowedDenom(ctx context.Context, req *marketplace.MsgAddAllowedDenom) (*marketplace.MsgAddAllowedDenomResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	if err := k.stateStore.AllowedDenomTable().Insert(ctx, &api.AllowedDenom{
		BankDenom:    req.BankDenom,
		DisplayDenom: req.DisplayDenom,
		Exponent:     req.Exponent,
	}); err != nil {
		if ormerrors.PrimaryKeyConstraintViolation.Is(err) {
			return nil, sdkerrors.ErrConflict.Wrapf("bank denom %s already exists", req.BankDenom)
		} else if ormerrors.UniqueKeyViolation.Is(err) {
			return nil, sdkerrors.ErrConflict.Wrapf("display denom %s already exists", req.DisplayDenom)
		}

		return nil, sdkerrors.ErrInvalidRequest.Wrapf("could not add denom: %s", err.Error())
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(&marketplace.EventAllowDenom{Denom: req.BankDenom})

	return &marketplace.MsgAddAllowedDenomResponse{}, nil
}
