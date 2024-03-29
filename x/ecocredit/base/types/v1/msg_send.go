package v1

import (
	"fmt"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
)

var _ legacytx.LegacyMsg = &MsgSend{}

// Route implements the LegacyMsg interface.
func (m MsgSend) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgSend) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSend) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("sender: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("recipient: %s", err)
	}

	if m.Sender == m.Recipient {
		return sdkerrors.ErrInvalidRequest.Wrap("sender and recipient cannot be the same")
	}

	if len(m.Credits) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("credits cannot be empty")
	}

	for i, credits := range m.Credits {
		creditIndex := fmt.Sprintf("credits[%d]", i)

		if err := base.ValidateBatchDenom(credits.BatchDenom); err != nil {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: batch denom: %s", creditIndex, err)
		}

		if credits.TradableAmount == "" && credits.RetiredAmount == "" {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: tradable amount or retired amount required", creditIndex)
		}

		if _, err := math.NewNonNegativeDecFromString(credits.TradableAmount); err != nil {
			return errors.Wrapf(err, "%s", creditIndex)
		}

		retiredAmount, err := math.NewNonNegativeDecFromString(credits.RetiredAmount)
		if err != nil {
			return errors.Wrapf(err, "%s", creditIndex)
		}

		if !retiredAmount.IsZero() {
			if err = base.ValidateJurisdiction(credits.RetirementJurisdiction); err != nil {
				return sdkerrors.ErrInvalidRequest.Wrapf("%s: retirement jurisdiction: %s", creditIndex, err)
			}

			if len(credits.RetirementReason) > base.MaxNoteLength {
				return ecocredit.ErrMaxLimit.Wrapf("%s: retirement reason: max length %d", creditIndex, base.MaxNoteLength)
			}
		}
	}

	return nil
}

// GetSigners returns the expected signers for MsgSend.
func (m *MsgSend) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}
