package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs a basic validation of credit class
func (c Class) Validate() error {
	if len(c.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("credit class metadata")
	}

	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(c.Admin).String()); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if len(c.Id) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class id cannot be empty")
	}

	if err := ValidateClassId(c.Id); err != nil {
		return err
	}

	if len(c.CreditTypeAbbrev) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("must specify a credit type abbreviation")
	}

	return nil
}
