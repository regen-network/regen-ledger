package marketplace

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateSellOrdersSuite struct {
	t   gocuke.TestingT
	msg *MsgUpdateSellOrders
	err error
}

func TestMsgUpdateSellOrders(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateSellOrdersSuite{}).Path("./features/msg_update_sell_orders.feature").Run()
}

func (s *msgUpdateSellOrdersSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateSellOrdersSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateSellOrders{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateSellOrdersSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateSellOrdersSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateSellOrdersSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
