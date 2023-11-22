package data

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var _ legacytx.LegacyMsg = &MsgAttest{}

// ValidateBasic does a sanity check on the provided data.
func (m *MsgAttest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Attestor); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	if len(m.ContentHashes) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("content hashes cannot be empty")
	}

	for _, hash := range m.ContentHashes {
		if hash == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("content hash cannot be empty")
		}
		err := hash.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// GetSigners returns the expected signers for MsgAttest.
func (m *MsgAttest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Attestor)
	return []sdk.AccAddress{addr}
}

// Route implements the LegacyMsg interface.
func (m MsgAttest) Route() string { return sdk.MsgTypeURL(&m) }

// Type implements the LegacyMsg interface.
func (m MsgAttest) Type() string { return sdk.MsgTypeURL(&m) }

// GetSignBytes implements the LegacyMsg interface.
func (m MsgAttest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
