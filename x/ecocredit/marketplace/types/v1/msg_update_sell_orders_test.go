package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateSellOrdersSuite struct {
	t         gocuke.TestingT
	msg       *MsgUpdateSellOrders
	err       error
	signBytes string
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

func (s *msgUpdateSellOrdersSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgUpdateSellOrdersSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
