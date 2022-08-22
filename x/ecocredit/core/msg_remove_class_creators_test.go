package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveClassCreators struct {
	t   gocuke.TestingT
	msg *MsgRemoveClassCreators
	err error
}

func TestMsgRemoveClassCreators(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveClassCreators{}).Path("./features/msg_remove_class_creators.feature").Run()
}

func (s *msgRemoveClassCreators) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveClassCreators) TheMessage(a gocuke.DocString) {
	s.msg = &MsgRemoveClassCreators{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveClassCreators) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveClassCreators) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgRemoveClassCreators) ExpectNoError() {
	require.NoError(s.t, s.err)
}
