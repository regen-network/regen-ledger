package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (m AllowedDenom) Validate() error {
	if err := sdk.ValidateDenom(m.BankDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid bank denom: %s", err.Error())
	}
	if err := sdk.ValidateDenom(m.DisplayDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid display_denom: %s", err.Error())
	}
	if _, err := core.ExponentToPrefix(m.Exponent); err != nil {
		return err
	}
	return nil
}
