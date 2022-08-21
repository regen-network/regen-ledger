package basket

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateBasketFeesSuite struct {
	t   gocuke.TestingT
	msg *MsgUpdateBasketFees
	err error
}

func TestMsgUpdateBasketFeesSuite(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateBasketFeesSuite{}).Path("./features/msg_update_basket_fee.feature").Run()
}

func (s *msgUpdateBasketFeesSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateBasketFeesSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateBasketFees{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateBasketFeesSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateBasketFeesSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateBasketFeesSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
