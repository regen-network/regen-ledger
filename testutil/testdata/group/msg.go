package group

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/jsonpb"
)

const (
	msgTypeMyMsgA = "always_succeed"
	msgTypeMyMsgB = "always_fail"
	msgTypeMyMsgC = "set_value"
	msgTypeMyMsgD = "inc_counter"
	msgTypeMyMsgE = "conditional"
	msgTypeMyMsgF = "authenticate"
)

var _ sdk.Msg = &MsgPropose{}

func (m MsgPropose) Route() string { return ModuleName }

func (m MsgPropose) Type() string { return msgTypeMyMsgA }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgPropose) GetSigners() []sdk.AccAddress {
	return m.Base.Proposers
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgPropose) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgPropose) ValidateBasic() error {
	if err := m.Base.ValidateBasic(); err != nil {
		return err
	}
	for i, any := range m.Msgs {
		msg, ok := any.GetCachedValue().(sdk.Msg)
		if !ok {
			return errors.Wrapf(errors.ErrUnpackAny, "cannot unpack Any into sdk.Msg %T", any)
		}
		if err := msg.ValidateBasic(); err != nil {
			return errors.Wrapf(err, "msg %d", i)
		}
	}
	return nil
}

// GetMsgs unpacks m.Msgs Any's into sdk.Msg's
func (m MsgPropose) GetMsgs() []sdk.Msg {
	msgs := make([]sdk.Msg, len(m.Msgs))
	for i, any := range m.Msgs {
		msg, ok := any.GetCachedValue().(sdk.Msg)
		if !ok {
			return nil
		}
		msgs[i] = msg
	}
	return msgs
}

var _ sdk.Msg = &MsgAlwaysSucceed{}

func (m MsgAlwaysSucceed) Route() string { return ModuleName }

func (m MsgAlwaysSucceed) Type() string { return msgTypeMyMsgA }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgAlwaysSucceed) GetSigners() []sdk.AccAddress {
	return nil // nothing to do
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgAlwaysSucceed) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(errors.Wrap(err, "get sign bytes"))
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgAlwaysSucceed) ValidateBasic() error {
	return nil
}

var _ sdk.Msg = &MsgAlwaysFail{}

func (m MsgAlwaysFail) Route() string { return ModuleName }

func (m MsgAlwaysFail) Type() string { return msgTypeMyMsgB }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgAlwaysFail) GetSigners() []sdk.AccAddress {
	return nil
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgAlwaysFail) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgAlwaysFail) ValidateBasic() error {
	return nil
}

var _ sdk.Msg = &MsgSetValue{}

func (m MsgSetValue) Route() string { return ModuleName }

func (m MsgSetValue) Type() string { return msgTypeMyMsgC }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgSetValue) GetSigners() []sdk.AccAddress {
	return nil
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgSetValue) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgSetValue) ValidateBasic() error {
	return nil
}

var _ sdk.Msg = &MsgIncCounter{}

func (m MsgIncCounter) Route() string { return ModuleName }

func (m MsgIncCounter) Type() string { return msgTypeMyMsgD }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgIncCounter) GetSigners() []sdk.AccAddress {
	return nil
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgIncCounter) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgIncCounter) ValidateBasic() error {
	return nil
}

var _ sdk.Msg = &MsgConditional{}

func (m MsgConditional) Route() string { return ModuleName }

func (m MsgConditional) Type() string { return msgTypeMyMsgE }

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgConditional) GetSigners() []sdk.AccAddress {
	return nil
}

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgConditional) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgConditional) ValidateBasic() error {
	return nil
}

var _ sdk.Msg = &MsgAuthenticate{}

func (m MsgAuthenticate) Route() string { return ModuleName }

func (m MsgAuthenticate) Type() string { return msgTypeMyMsgF }

// GetSignBytes returns the bytes for the message signer to sign on
func (m MsgAuthenticate) GetSignBytes() []byte {
	var buf bytes.Buffer
	enc := jsonpb.Marshaler{}
	if err := enc.Marshal(&buf, &m); err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(buf.Bytes())
}

// ValidateBasic does a sanity check on the provided data
func (m MsgAuthenticate) ValidateBasic() error {
	return nil
}
