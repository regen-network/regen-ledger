package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRemoveClassCreator struct {
	t   gocuke.TestingT
	msg *MsgRemoveClassCreator
	err error
}

func TestMsgRemoveClassCreators(t *testing.T) {
	gocuke.NewRunner(t, &msgRemoveClassCreator{}).Path("./features/msg_remove_class_creator.feature").Run()
}

func (s *msgRemoveClassCreator) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRemoveClassCreator) TheMessage(a gocuke.DocString) {
	s.msg = &MsgRemoveClassCreator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRemoveClassCreator) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRemoveClassCreator) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgRemoveClassCreator) ExpectNoError() {
	require.NoError(s.t, s.err)
}
