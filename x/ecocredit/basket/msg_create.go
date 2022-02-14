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
const exponentMax = 32
const creditAbbrMaxLen = 32

var errBadReq = sdkerrors.ErrInvalidRequest
var reName = regexp.MustCompile(fmt.Sprintf("^[[:alnum:]]{%d,%d}$", nameMinLen, nameMaxLen))

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgCreate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("malformed curator address " + err.Error())
	}
	if !reName.MatchString(m.Name) {
		return errBadReq.Wrapf("name must be between %d and %d alpha-numeric characters long", nameMinLen, nameMaxLen)
	}
	if m.Exponent > exponentMax {
		return errBadReq.Wrapf("exponent must not be bigger than %d", exponentMax)
	}
	if m.CreditTypeAbbrev == "" {
		return errBadReq.Wrap("credit_type_name must be defined")
	}
	if len(m.CreditTypeAbbrev) > creditAbbrMaxLen {
		return errBadReq.Wrapf("credit_type_name must not be longer than %d", creditAbbrMaxLen)
	}
	if err := validateDateCriteria(m.DateCriteria); err != nil {
		return err
	}
	if len(m.AllowedClasses) == 0 {
		return errBadReq.Wrap("allowed_classes is required")
	}
	for i := range m.AllowedClasses {
		if m.AllowedClasses[i] == "" {
			return errBadReq.Wrapf("allowed_classes[%d] must be defined", i)
		}
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
