package v1

import (
	"fmt"

	proto "github.com/cosmos/gogoproto/proto"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var (
	_ legacytx.LegacyMsg = &MsgSubmitTx{}

	_ codectypes.UnpackInterfacesMessage = MsgSubmitTx{}
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

// LegacyMsg.Type implementations
func (m MsgSubmitTx) Route() string { return "" }
func (m MsgSubmitTx) Type() string  { return sdk.MsgTypeURL(&m) }
func (m *MsgSubmitTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

// NewMsgSubmitTx creates a new MsgSubmitTx instance
func NewMsgSubmitTx(owner string, connectionID string, msg sdk.Msg) *MsgSubmitTx {

	anyMsg, err := PackTxMsgAny(msg)
	if err != nil {
		panic(err)
	}

	return &MsgSubmitTx{
		Owner:        owner,
		ConnectionId: connectionID,
		Msg:          anyMsg,
	}
}

// PackTxMsgAny marshals the sdk.Msg payload to a protobuf Any type
func PackTxMsgAny(sdkMsg sdk.Msg) (*codectypes.Any, error) {
	msg, ok := sdkMsg.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("can't proto marshal %T", sdkMsg)
	}

	return codectypes.NewAnyWithValue(msg)
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (m MsgSubmitTx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var sdkMsg sdk.Msg

	return unpacker.UnpackAny(m.Msg, &sdkMsg)
}
