package testdata

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/modules/incubator/group"
	"github.com/regen-network/regen-ledger/orm"
)

func (m *MyAppProposal) GetBase() group.ProposalBase {
	return m.Base
}
func (m *MyAppProposal) SetBase(b group.ProposalBase) {
	m.Base = b
}

func (m *MyAppProposal) GetMsgs() []sdk.Msg {
	return MyAppMsgs(m.Msgs).AsSDKMsgs()
}

func (m *MyAppProposal) SetMsgs(new []sdk.Msg) error {
	m.Msgs = make([]MyAppMsg, len(new))
	for i := range new {
		if new[i] == nil {
			return errors.Wrap(group.ErrInvalid, "msg must not be nil")
		}
		n := MyAppMsg{}
		if err := n.SetMsg(new[i]); err != nil {
			return err
		}
		m.Msgs[i] = n
	}
	return nil
}

var _ orm.Validateable = MyAppProposal{}

func (m MyAppProposal) ValidateBasic() error {
	if err:=m.Base.ValidateBasic(); err!=nil{
		return errors.Wrap(err, "base")
	}
	for i, msg := range m.Msgs{
		if err:=msg.GetMsg().ValidateBasic(); err!=nil{
			return errors.Wrapf(err, "message %i", i)
		}
	}
	return nil
}

type MyAppMsgs []MyAppMsg

// AsSDKMsgs type conversion to sdk.Msg.
// Can return nil values in slice but should not.
func (m MyAppMsgs) AsSDKMsgs() []sdk.Msg {
	r := make([]sdk.Msg, len(m))
	for i := range m {
		r[i] = m[i].GetMsg()
	}
	return r
}
