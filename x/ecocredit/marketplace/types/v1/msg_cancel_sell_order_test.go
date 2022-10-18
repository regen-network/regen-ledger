package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCancelSuite struct {
	t         gocuke.TestingT
	msg       *MsgCancelSellOrder
	err       error
	signBytes string
}

func TestMsgCancelSellOrder(t *testing.T) {
	gocuke.NewRunner(t, &msgCancelSuite{}).Path("./features/msg_cancel_sell_order.feature").Run()
}

func (s *msgCancelSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCancelSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCancelSellOrder{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCancelSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCancelSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCancelSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCancelSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgCancelSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
