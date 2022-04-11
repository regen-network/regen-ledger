package core

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgMintBatchCredits{}

// Route implements the LegacyMsg interface.
func (m MsgMintBatchCredits) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgMintBatchCredits) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgMintBatchCredits) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgMintBatchCredits) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Issuer)
	if err != nil {
		return sdkerrors.Wrap(err, "malformed issuer address")
	}
	if err := ValidateDenom(m.BatchDenom); err != nil {
		return err
	}
	if len(m.Note) > 512 {
		return errBadReq.Wrap("note must not be longer than 512 characters")
	}
	if err = validateBatchIssuances(m.Issuance); err != nil {
		return err
	}
	return validateOriginTx(m.OriginTx, true)
}

// GetSigners returns the expected signers for MsgMintBatchCredits.
func (m *MsgMintBatchCredits) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Issuer)
	return []sdk.AccAddress{addr}
}

func validateBatchIssuances(iss []*BatchIssuance) error {
	if len(iss) == 0 {
		errBadReq.Wrap("issuance list must not be empty")
	}
	for idx, i := range iss {
		if i == nil {
			return errBadReq.Wrapf("issuance[%d] must be defined", idx)
		}
		_, err := sdk.AccAddressFromBech32(i.Recipient)
		if err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("issuance[%d].recipient", idx)
		}
		if i.TradableAmount != "" {
			if _, err := math.NewNonNegativeDecFromString(i.TradableAmount); err != nil {
				return errBadReq.Wrapf("issuance[%d].tradable_amount; %v", idx, err)
			}
		}

		if i.RetiredAmount != "" {
			retiredAmount, err := math.NewNonNegativeDecFromString(i.RetiredAmount)
			if err != nil {
				return errBadReq.Wrapf("issuance[%d].retired_amount; %v", idx, err)
			}

			if !retiredAmount.IsZero() {
				if err = ValidateLocation(i.RetirementLocation); err != nil {
					return errBadReq.Wrapf("issuance[%d].retirement_location; %v", idx, err)
				}
			}
		}
	}
	return nil
}

var reOriginTxNote = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9 _\-]{1,64}$`)

func validateOriginTx(o *OriginTx, required bool) error {
	if o == nil {
		if required {
			return errBadReq.Wrap("origin_tx is required")
		}
		return nil
	}
	if !reOriginTxNote.MatchString(o.Typ) {
		return errBadReq.Wrap("origin_tx.typ must be 2-64 long, valid characters: alpha-numberic, space, '-' or '_'")
	}
	if !reOriginTxNote.MatchString(o.Id) {
		return errBadReq.Wrap("origin_tx.id must be 2-64 long, valid characters: alpha-numberic, space, '-' or '_'")
	}
	return nil
}
