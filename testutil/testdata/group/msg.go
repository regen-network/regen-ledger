package group

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.MsgRequest = &MsgProposeRequest{}

// GetSigners returns the addresses that must sign over msg.GetSignBytes()
func (m MsgProposeRequest) GetSigners() []sdk.AccAddress {
	return m.Base.Proposers
}

// ValidateBasic does a sanity check on the provided data
func (m MsgProposeRequest) ValidateBasic() error {
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
func (m MsgProposeRequest) GetMsgs() []sdk.Msg {
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
