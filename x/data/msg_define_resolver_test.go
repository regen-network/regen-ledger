package data

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgDefineResolverSuite struct {
	t         gocuke.TestingT
	msg       *MsgDefineResolver
	err       error
	signBytes string
}

func TestMsgDefineResolver(t *testing.T) {
	runner := gocuke.NewRunner(t, &msgDefineResolverSuite{}).Path("./features/msg_define_resolver.feature")
	runner.Step(`^the\s+message\s+"((?:[^\"]|\")*)"`, (*msgDefineResolverSuite).TheMessage)
	runner.Run()
}

func (s *msgDefineResolverSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgDefineResolverSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgDefineResolver{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgDefineResolverSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgDefineResolverSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgDefineResolverSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgDefineResolverSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgDefineResolverSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
