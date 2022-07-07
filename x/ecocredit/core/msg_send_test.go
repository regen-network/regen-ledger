package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgSend struct {
	t   gocuke.TestingT
	msg *MsgSend
	err error
}

func TestMsgSend(t *testing.T) {
	gocuke.NewRunner(t, &msgSend{}).Path("./features/msg_send.feature").Run()
}

func (s *msgSend) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
}

func (s *msgSend) TheMessage(a gocuke.DocString) {
	s.msg = &MsgSend{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgSend) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgSend) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgSend) ExpectNoError() {
	require.NoError(s.t, s.err)
}
