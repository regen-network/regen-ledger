package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

func (k Keeper) AddAskDenom(ctx sdk.Context, p *marketplace.AskDenomProposal) error {
	if p == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("nil proposal")
	}
	if err := p.ValidateBasic(); err != nil {
		return err
	}
	askDenom := p.AllowedDenom
	if err := k.stateStore.AllowedDenomTable().Insert(sdk.WrapSDKContext(ctx), &api.AllowedDenom{
		BankDenom:    askDenom.BankDenom,
		DisplayDenom: askDenom.DisplayDenom,
		Exponent:     askDenom.Exponent,
	}); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("could not add denom %s: %s", askDenom.BankDenom, err.Error())
	}
	return nil
}
