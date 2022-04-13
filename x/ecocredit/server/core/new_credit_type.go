package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// NewCreditType adds a new credit type to the network
func (k Keeper) NewCreditType(ctx sdk.Context, ctp *core.CreditTypeProposal) error {
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
		return fmt.Errorf("could not insert credit type with abbreviation %s: %w", ct.Abbreviation, err)
	}
	return nil
}
