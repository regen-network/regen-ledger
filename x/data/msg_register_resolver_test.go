package data

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgRegisterResolverSuite struct {
	t         gocuke.TestingT
	msg       *MsgRegisterResolver
	err       error
	signBytes string
}

func TestMsgRegisterResolver(t *testing.T) {
	runner := gocuke.NewRunner(t, &msgRegisterResolverSuite{}).Path("./features/msg_register_resolver.feature")
	runner.Step(`^the\s+message\s+"((?:[^\"]|\")*)"`, (*msgRegisterResolverSuite).TheMessage)
	runner.Run()
}

func (s *msgRegisterResolverSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgRegisterResolverSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgRegisterResolver{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgRegisterResolverSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgRegisterResolverSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgRegisterResolverSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgRegisterResolverSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgRegisterResolverSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
