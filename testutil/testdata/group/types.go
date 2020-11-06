package group

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/x/group"
)

func (m *MyAppProposal) GetBase() group.ProposalBase {
	return m.Base
}

func (m *MyAppProposal) SetBase(b group.ProposalBase) {
	m.Base = b
}

func (m *MyAppProposal) GetMsgs() []sdk.Msg {
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

func (m *MyAppProposal) SetMsgs(new []sdk.Msg) error {
	m.Msgs = make([]*types.Any, len(new))
	for i := range new {
		if new[i] == nil {
			return errors.Wrap(group.ErrInvalid, "msg must not be nil")
		}
		any, err := types.NewAnyWithValue(new[i])
		if err != nil {
			return err
		}
		m.Msgs[i] = any
	}
	return nil
}

func (m MyAppProposal) ValidateBasic() error {
	if err := m.Base.ValidateBasic(); err != nil {
		return errors.Wrap(err, "base")
	}
	msgs := m.GetMsgs()
	for i, msg := range msgs {
		if err := msg.ValidateBasic(); err != nil {
			return errors.Wrapf(err, "message %i", i)
		}
	}
	return nil
}
