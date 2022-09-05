package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateBasketFeeSuite struct {
	t   gocuke.TestingT
	msg *MsgUpdateBasketFee
	err error
}

func TestMsgUpdateBasketFeeSuite(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateBasketFeeSuite{}).Path("./features/msg_update_basket_fee.feature").Run()
}

func (s *msgUpdateBasketFeeSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateBasketFeeSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateBasketFee{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateBasketFeeSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateBasketFeeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateBasketFeeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
