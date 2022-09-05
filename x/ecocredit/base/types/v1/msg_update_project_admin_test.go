package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateProjectAdmin struct {
	t   gocuke.TestingT
	msg *MsgUpdateProjectAdmin
	err error
}

func TestMsgUpdateProjectAdmin(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateProjectAdmin{}).Path("./features/msg_update_project_admin.feature").Run()
}

func (s *msgUpdateProjectAdmin) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateProjectAdmin) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateProjectAdmin{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateProjectAdmin) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateProjectAdmin) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateProjectAdmin) ExpectNoError() {
	require.NoError(s.t, s.err)
}
