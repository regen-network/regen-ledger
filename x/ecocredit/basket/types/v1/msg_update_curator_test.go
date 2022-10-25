package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateCuratorSuite struct {
	t         gocuke.TestingT
	msg       *MsgUpdateCurator
	err       error
	signBytes string
}

func TestMsgUpdateCurator(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateCuratorSuite{}).Path("./features/msg_update_curator.feature").Run()
}

func (s *msgUpdateCuratorSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateCuratorSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateCurator{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateCuratorSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateCuratorSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateCuratorSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateCuratorSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgUpdateCuratorSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
