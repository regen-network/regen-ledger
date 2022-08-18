package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs a basic validation of project
func (p Project) Validate() error {
	if _, err := sdk.AccAddressFromBech32(sdk.AccAddress(p.Admin).String()); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	if p.ClassKey == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("class key cannot be zero")
	}

	if err := ValidateJurisdiction(p.Jurisdiction); err != nil {
		return err
	}

	if len(p.Metadata) > MaxMetadataLength {
		return ecocredit.ErrMaxLimit.Wrap("project metadata")
	}

	if err := ValidateProjectId(p.Id); err != nil {
		return err
	}

	return nil
}
