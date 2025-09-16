package v1

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4/base"
)

var _ sdk.Msg = &MsgUpdateClassIssuers{}

func (m MsgUpdateClassIssuers) Route() string { return sdk.MsgTypeURL(&m) }

func (m MsgUpdateClassIssuers) Type() string { return sdk.MsgTypeURL(&m) }

func (m *MsgUpdateClassIssuers) ValidateBasic() error {
	if err := base.ValidateClassID(m.ClassId); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("class id: %s", err)
	}

	if len(m.AddIssuers) == 0 && len(m.RemoveIssuers) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("must specify at least one of add_issuers or remove_issuers")
	}

	duplicateAddMap := make(map[string]bool)
	for i, issuer := range m.AddIssuers {
		addIssuerIndex := fmt.Sprintf("add_issuers[%d]", i)

		if _, ok := duplicateAddMap[issuer]; ok {
			return sdkerrors.ErrInvalidRequest.Wrapf("%s: duplicate issuer", addIssuerIndex)
		}

		duplicateAddMap[issuer] = true
	}

	duplicateRemoveMap := make(map[string]bool)
	for i, issuer := range m.RemoveIssuers {
		removeIssuerIndex := fmt.Sprintf("remove_issuers[%d]", i)
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
