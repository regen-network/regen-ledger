package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgSellSuite struct {
	t   gocuke.TestingT
	msg *MsgSell
	err error
}

func TestMsgSell(t *testing.T) {
	gocuke.NewRunner(t, &msgSellSuite{}).Path("./features/msg_sell.feature").Run()
}

func (s *msgSellSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgSellSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgSell{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgSellSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgSellSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgSellSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
