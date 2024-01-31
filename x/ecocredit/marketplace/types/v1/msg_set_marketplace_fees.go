package v1

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/types/v2/math"
)

var _ legacytx.LegacyMsg = &MsgSetMarketplaceFees{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgSetMarketplaceFees) ValidateBasic() error {
	_, err := math.NewPositiveDecFromString(m.Fees.BuyerPercentageFee)
	if err != nil {
		return err
	}

	_, err = math.NewPositiveDecFromString(m.Fees.SellerPercentageFee)
	if err != nil {
		return err
	}

	_, err = types.AccAddressFromBech32(m.Authority)
	return err
}

// GetSigners implements the LegacyMsg interface.
func (m *MsgSetMarketplaceFees) GetSigners() []types.AccAddress {
	return []types.AccAddress{types.MustAccAddressFromBech32(m.Authority)}
}

// Route implements the LegacyMsg interface.
func (m *MsgSetMarketplaceFees) Route() string { return types.MsgTypeURL(m) }

// Type implements the LegacyMsg interface.
func (m *MsgSetMarketplaceFees) Type() string { return types.MsgTypeURL(m) }

// GetSignBytes implements the LegacyMsg interface.
func (m *MsgSetMarketplaceFees) GetSignBytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}
