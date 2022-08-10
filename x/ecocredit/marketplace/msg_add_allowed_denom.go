package marketplace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit"

	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgAddAllowedDenom{}

func (m MsgAddAllowedDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.Wrapf(err, "invalid authority address")
	}

	msg := AllowedDenom{
		BankDenom:    m.BankDenom,
		DisplayDenom: m.DisplayDenom,
		Exponent:     m.Exponent,
	}

	return msg.Validate()
}

func (m MsgAddAllowedDenom) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgAddAllowedDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgAddAllowedDenom) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgAddAllowedDenom) Type() string { return sdk.MsgTypeURL(&m) }
