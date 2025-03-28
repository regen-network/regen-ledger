package v1

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/jsonpb"
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

func (s *msgBuyDirectSuite) RetirementReasonWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Orders[0].RetirementReason = strings.Repeat("x", int(length))
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

func (s *msgBuyDirectSuite) ExpectErrorContains(a string) {
	if a != "" {
		require.ErrorContains(s.t, s.err, a)
	} else {
		require.NoError(s.t, s.err)
	}
}

func (s *msgBuyDirectSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgBuyDirectSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
