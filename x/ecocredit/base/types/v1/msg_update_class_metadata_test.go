package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgUpdateClassMetadata struct {
	t   gocuke.TestingT
	msg *MsgUpdateClassMetadata
	err error
}

func TestMsgUpdateClassMetadata(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateClassMetadata{}).Path("./features/msg_update_class_metadata.feature").Run()
}

func (s *msgUpdateClassMetadata) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgUpdateClassMetadata) TheMessage(a gocuke.DocString) {
	s.msg = &MsgUpdateClassMetadata{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgUpdateClassMetadata) NewMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.NewMetadata = strings.Repeat("x", int(length))
}

func (s *msgUpdateClassMetadata) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateClassMetadata) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgUpdateClassMetadata) ExpectNoError() {
	require.NoError(s.t, s.err)
}
