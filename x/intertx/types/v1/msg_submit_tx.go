package v1

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var (
	_ legacytx.LegacyMsg = &MsgSubmitTx{}
)

// ValidateBasic does a sanity check on the provided data.
func (m MsgSubmitTx) ValidateBasic() error {
	if m.Owner == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("owner cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Owner); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("owner: %s", err.Error())
	}
	if m.ConnectionId == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("connection_id cannot be empty")
	}
	if m.Msg == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("msg cannot be empty")
	}
	return nil
}

// GetSigners returns the expected signers.
func (m MsgSubmitTx) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Owner)
	return []sdk.AccAddress{addr}
}

// GetSignBytes implements the LegacyMsg interface.
func (m MsgSubmitTx) GetSignBytes() []byte {
	return ModuleCdc.MustMarshalJSON(&m)
}

// Route implements the LegacyMsg interface.
func (m MsgSubmitTx) Route() string {
	return sdk.MsgTypeURL(&m)
}

// Type implements the LegacyMsg interface.
func (m MsgSubmitTx) Type() string {
	return sdk.MsgTypeURL(&m)
}

// NewMsgSubmitTx creates a new MsgSubmitTx instance
func NewMsgSubmitTx(owner string, connectionID string, msg sdk.Msg) *MsgSubmitTx {

	anyMsg, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		panic(err)
	}

	return &MsgSubmitTx{
		Owner:        owner,
		ConnectionId: connectionID,
		Msg:          anyMsg,
	}
}
