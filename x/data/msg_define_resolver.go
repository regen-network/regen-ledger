package data

import (
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateBasic does a sanity check on the provided data.
func (m *MsgDefineResolver) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Definer); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrap(err.Error())
	}

	if _, err := url.ParseRequestURI(m.ResolverUrl); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrap("invalid resolver url")
	}

	return nil
}
