package basket

import (
	"fmt"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

const (
	nameMinLen  = 3
	nameMaxLen  = 8
	descrMaxLen = 256
)

var (
	_ legacytx.LegacyMsg = &MsgCreate{}

	// first character must be alphabetic, the rest can be alphanumeric. We reduce length
	// constraints by one to account for the first character being forced to alphabetic.
	reName = regexp.MustCompile(fmt.Sprintf("^[[:alpha:]][[:alnum:]]{%d,%d}$", nameMinLen-1, nameMaxLen-1))
)

// Route implements LegacyMsg.
func (m MsgCreate) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgCreate) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements LegacyMsg.
func (m MsgCreate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgCreate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("malformed curator address: " + err.Error())
	}

	if len(m.Name) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrapf("name cannot be empty")
	}

	if !reName.MatchString(m.Name) {
		return sdkerrors.ErrInvalidRequest.Wrapf("name must start with an alphabetic character, and be between %d and %d alphanumeric characters long", nameMinLen, nameMaxLen)
	}

	if len(m.Description) > descrMaxLen {
		return sdkerrors.ErrInvalidRequest.Wrapf("description length cannot be greater than %d characters", descrMaxLen)
	}

	if _, err := core.ExponentToPrefix(m.Exponent); err != nil {
		return err
	}

	if len(m.CreditTypeAbbrev) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrapf("credit type abbreviation cannot be empty")
	}

	if err := core.ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return err
	}

	if len(m.AllowedClasses) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("allowed classes cannot be empty")
	}

	for i := range m.AllowedClasses {
		if m.AllowedClasses[i] == "" {
			return sdkerrors.ErrInvalidRequest.Wrapf("allowed_classes[%d] cannot be empty", i)
		}
		if err := core.ValidateClassId(m.AllowedClasses[i]); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("allowed_classes[%d] is not a valid class ID: %s", i, err)
		}
	}

	if err := m.DateCriteria.Validate(); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("invalid date criteria: %s", err)
	}

	return m.Fee.Validate()
}

// GetSigners returns the expected signers for MsgCreate.
func (m MsgCreate) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}
