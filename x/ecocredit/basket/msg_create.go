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
	nameMinLen           = 3
	nameMaxLen           = 8
	descrMaxLen          = 256
	creditTypeAbbrMaxLen = 3
)

var (
	_ legacytx.LegacyMsg = &MsgCreate{}

	// first character must be alphabetic, the rest can be alphanumeric. We reduce length
	// constraints by one to account for the first character being forced to alphabetic.
	reName    = regexp.MustCompile(fmt.Sprintf("^[[:alpha:]][[:alnum:]]{%d,%d}$", nameMinLen-1, nameMaxLen-1))
	errBadReq = sdkerrors.ErrInvalidRequest
)

// ValidateBasic does a stateless sanity check on the provided data.
func (m MsgCreate) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Curator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap("malformed curator address " + err.Error())
	}
	if !reName.MatchString(m.Name) {
		return errBadReq.Wrapf("name must start with an alphabetic character, and be between %d and %d alphanumeric characters long", nameMinLen, nameMaxLen)
	}
	if _, err := core.ExponentToPrefix(m.Exponent); err != nil {
		return err
	}
	if err := core.ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return err
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
	minStartDate := d.GetMinStartDate()
	startDateWindow := d.GetStartDateWindow()
	yearsInThePast := d.GetYearsInThePast()
	if (minStartDate != nil && startDateWindow != nil) || (startDateWindow != nil && yearsInThePast != 0) || (minStartDate != nil && yearsInThePast != 0) {
		return errBadReq.Wrap("only one of date_criteria.min_start_date, date_criteria.start_date_window, or date_criteria.years_in_the_past must be set")
	}
	if minStartDate != nil {
		if minStartDate.Seconds < -2208992400 { // batch older than 1900-01-01 is an obvious error
			return errBadReq.Wrap("date_criteria.min_start_date must be after 1900-01-01")
		}
	} else if startDateWindow != nil {
		if startDateWindow.Seconds < 24*3600 {
			return errBadReq.Wrap("date_criteria.start_date_window must be at least 1 day")
		}
	}
	return nil
}

// BasketDenom formats denom and display denom:
// * denom: eco.<m.Exponent><m.CreditTypeAbbrev>.<m.Name>
// * display denom: eco.<m.CreditTypeAbbrev>.<m.Name>
// Returns error if MsgCrete.Exponent is not supported
func BasketDenom(name, creditTypeAbbrev string, exponent uint32) (string, string, error) {
	const basketDenomPrefix = "eco."
	denomPrefix, err := core.ExponentToPrefix(exponent)
	if err != nil {
		return "", "", err
	}

	denomTail := creditTypeAbbrev + "." + name
	displayDenomName := basketDenomPrefix + denomTail    //
	denom := basketDenomPrefix + denomPrefix + denomTail // eco.<credit-class>.<name>
	return denom, displayDenomName, nil
}
