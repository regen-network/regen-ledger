package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgUpdateClassAdmin struct {
	t   gocuke.TestingT
	msg *MsgUpdateClassAdmin
	err error
}

func TestMsgUpdateClassAdmin(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateClassAdmin{}).Path("./features/msg_update_class_admin.feature").Run()
}

func (s *msgUpdateClassAdmin) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
}

func (s *msgUpdateClassAdmin) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateClassAdmin{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateClassAdmin) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateClassAdmin) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateClassAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}
