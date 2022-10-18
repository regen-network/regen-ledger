package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgBuyDirectSuite struct {
	t         gocuke.TestingT
	msg       *MsgBuyDirect
	err       error
	signBytes string
}

func TestMsgBuyDirect(t *testing.T) {
	gocuke.NewRunner(t, &msgBuyDirectSuite{}).Path("./features/msg_buy_direct.feature").Run()
}

func (s *msgBuyDirectSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgBuyDirectSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgBuyDirect{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgBuyDirectSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgBuyDirectSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgBuyDirectSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgBuyDirectSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgBuyDirectSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
