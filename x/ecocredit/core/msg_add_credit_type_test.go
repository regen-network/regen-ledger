package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAddCreditType struct {
	t   gocuke.TestingT
	msg *MsgAddCreditType
	err error
}

func TestMsgAddCreditType(t *testing.T) {
	gocuke.NewRunner(t, &msgAddCreditType{}).Path("./features/msg_add_credit_type.feature").Run()
}

func (s *msgAddCreditType) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAddCreditType) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAddCreditType{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAddCreditType) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAddCreditType) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAddCreditType) ExpectNoError() {
	require.NoError(s.t, s.err)
}
