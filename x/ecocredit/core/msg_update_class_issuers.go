package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ legacytx.LegacyMsg = &MsgUpdateClassIssuers{}

func (m MsgUpdateClassIssuers) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassIssuers) Type() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassIssuers) GetSignBytes() []byte {
	return sdk.MustSortJSON(ecocredit.ModuleCdc.MustMarshalJSON(&m))
}

func (m *MsgUpdateClassIssuers) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("admin: %s", err)
	}

	if err := ValidateClassID(m.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("class id: %s", err)
	}

	if len(m.AddIssuers) == 0 && len(m.RemoveIssuers) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("must specify at least one of add_issuers or remove_issuers")
	}

	duplicateAddMap := make(map[string]bool)
	for i, issuer := range m.AddIssuers {
		addIssuerIndex := fmt.Sprintf("add_issuers[%d]", i)

		if _, err := sdk.AccAddressFromBech32(issuer); err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("%s: %s", addIssuerIndex, err)
		}

		if _, ok := duplicateAddMap[issuer]; ok {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: duplicate issuer", addIssuerIndex)
		}

		duplicateAddMap[issuer] = true
	}

	duplicateRemoveMap := make(map[string]bool)
	for i, issuer := range m.RemoveIssuers {
		removeIssuerIndex := fmt.Sprintf("remove_issuers[%d]", i)

		if _, err := sdk.AccAddressFromBech32(issuer); err != nil {
			return sdkerrors.ErrInvalidAddress.Wrapf("%s: %s", removeIssuerIndex, err)
		}

		if _, ok := duplicateRemoveMap[issuer]; ok {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: duplicate issuer", removeIssuerIndex)
		}

		duplicateRemoveMap[issuer] = true
	}

	return nil
}

func (m *MsgUpdateClassIssuers) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Admin)
	return []sdk.AccAddress{addr}
}
