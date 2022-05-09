package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// AllowDenom is a gov handler method that adds a denom to the list of approved denoms that may be used in the
// marketplace.
func (k Keeper) AllowDenom(ctx sdk.Context, p *marketplace.AllowDenomProposal) error {
	if p == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("nil proposal")
	}
	if err := p.ValidateBasic(); err != nil {
		return err
	}
	denom := p.Denom
	if err := k.stateStore.AllowedDenomTable().Insert(sdk.WrapSDKContext(ctx), &api.AllowedDenom{
		BankDenom:    denom.BankDenom,
		DisplayDenom: denom.DisplayDenom,
		Exponent:     denom.Exponent,
	}); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("could not add denom %s: %s", denom.BankDenom, err.Error())
	}
	return ctx.EventManager().EmitTypedEvent(&marketplace.EventAllowDenom{Denom: denom.BankDenom})
}
