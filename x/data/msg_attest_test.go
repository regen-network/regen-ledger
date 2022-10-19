package data

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgAttestSuite struct {
	t         gocuke.TestingT
	msg       *MsgAttest
	err       error
	signBytes string
}

func TestMsgAttest(t *testing.T) {
	runner := gocuke.NewRunner(t, &msgAttestSuite{}).Path("./features/msg_attest.feature")
	runner.Step(`^the\s+message\s+"((?:[^\"]|\")*)"`, (*msgAttestSuite).TheMessage)
	runner.Run()
}

func (s *msgAttestSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgAttestSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgAttest{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgAttestSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgAttestSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgAttestSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgAttestSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgAttestSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
