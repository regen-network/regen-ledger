package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// AddCreditType adds a new credit type to the network
func (k Keeper) AddCreditType(ctx sdk.Context, ctp *core.CreditTypeProposal) error {
	if ctp == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("nil proposal")
	}
	if err := ctp.ValidateBasic(); err != nil {
		return err
	}
	ct := ctp.CreditType
	if err := k.stateStore.CreditTypeTable().Insert(sdk.WrapSDKContext(ctx), &api.CreditType{
		Abbreviation: ct.Abbreviation,
		Name:         ct.Name,
		Unit:         ct.Unit,
		Precision:    ct.Precision,
	}); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("could not insert credit type with abbreviation %s: %s", ct.Abbreviation, err.Error())
	}
	return nil
}
