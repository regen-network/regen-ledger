package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateDateCriteriaSuite struct {
	t         gocuke.TestingT
	msg       *MsgUpdateDateCriteria
	err       error
	signBytes string
}

func TestMsgUpdateDateCriteria(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateDateCriteriaSuite{}).Path("./features/msg_update_date_criteria.feature").Run()
}

func (s *msgUpdateDateCriteriaSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateDateCriteriaSuite) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateDateCriteria{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateDateCriteriaSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateDateCriteriaSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateDateCriteriaSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgUpdateDateCriteriaSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgUpdateDateCriteriaSuite) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
