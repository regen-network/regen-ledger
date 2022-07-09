package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgRetire struct {
	t   gocuke.TestingT
	msg *MsgRetire
	err error
}

func TestMsgRetire(t *testing.T) {
	gocuke.NewRunner(t, &msgRetire{}).Path("./features/msg_retire.feature").Run()
}

func (s *msgRetire) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
}

func (s *msgRetire) TheMessage(a gocuke.DocString) {
	s.msg = &MsgRetire{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRetire) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRetire) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgRetire) ExpectNoError() {
	require.NoError(s.t, s.err)
}
