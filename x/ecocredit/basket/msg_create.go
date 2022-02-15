package basket

import (
	"fmt"
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var (
	_ legacytx.LegacyMsg = &MsgCreate{}
)

const nameMinLen = 3
const nameMaxLen = 8
const descrMaxLen = 200
const exponentMax = 32
const creditTypeAbbrMaxLen = 3
const prefixMaxLen = 1

var errBadReq = sdkerrors.ErrInvalidRequest

// first character must be alphabetic, the rest can be alphanumeric. we reduce by one to account for the first character
// being forced to alphabetic.
var reName = regexp.MustCompile(fmt.Sprintf("^[[:alpha:]][[:alnum:]]{%d,%d}$", nameMinLen-1, nameMaxLen-1))

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgCreate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("malformed curator address " + err.Error())
	}
	if !reName.MatchString(m.Name) {
		return errBadReq.Wrapf("name must start with an alphabetic character, and be between %d and %d alphanumeric characters long", nameMinLen, nameMaxLen)
	}
	if m.Exponent > exponentMax {
		return errBadReq.Wrapf("exponent must not be bigger than %d", exponentMax)
	}
	if m.CreditTypeAbbrev == "" {
		return errBadReq.Wrap("credit_type_abbrev must be defined")
	}
	if len(m.CreditTypeAbbrev) > creditTypeAbbrMaxLen {
		return errBadReq.Wrapf("credit_type_abbrev must not be longer than %d", creditTypeAbbrMaxLen)
	}
	if err := validateDateCriteria(m.DateCriteria); err != nil {
		return err
	}
	if len(m.Description) > descrMaxLen {
		return errBadReq.Wrapf("description can't be longer than %d characters", descrMaxLen)
	}
	if len(m.AllowedClasses) == 0 {
		return errBadReq.Wrap("allowed_classes is required")
	}
	for i := range m.AllowedClasses {
		if m.AllowedClasses[i] == "" {
			return errBadReq.Wrapf("allowed_classes[%d] must be defined", i)
		}
	}
	if len(m.Prefix) != prefixMaxLen {
		return sdkerrors.ErrInvalidRequest.Wrapf("prefix cannot be longer than %d character(s)", prefixMaxLen)
	}
	return m.Fee.Validate()
}

// ValidateMsgCreate additional validation with access to the state data.
// minFee must be sorted.
func ValidateMsgCreate(m *MsgCreate, minFee sdk.Coins) error {
	if !m.Fee.IsAllGTE(minFee) {
		return sdkerrors.ErrInsufficientFee.Wrapf("minimum fee %s, got %s", minFee, m.Fee)
	}

	return nil
}

// GetSigners returns the expected signers for MsgCreate.
func (m MsgCreate) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Curator)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements LegacyMsg.
func (m MsgCreate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// Route implements LegacyMsg.
func (m MsgCreate) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements LegacyMsg.
func (m MsgCreate) Type() string { return sdk.MsgTypeURL(&m) }

func validateDateCriteria(d *DateCriteria) error {
	if d == nil {
		return nil
	}
	if x := d.GetMinStartDate(); x != nil {
		if x.Seconds < -2208992400 { // batch older than 1900-01-01 is an obvious error
			return errBadReq.Wrap("date_criteria.min_start_date must be after 1900-01-01")
		}
	} else if x := d.GetStartDateWindow(); x != nil {
		if x.Seconds < 24*3600 {
			return errBadReq.Wrap("date_criteria.start_date_window must be at least 1 day")
		}
	} else {
		return errBadReq.Wrapf("unsupported date_criteria value %t", d)
	}
	return nil
}
