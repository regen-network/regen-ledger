//nolint:revive,stylecheck
package v1

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgBridgeReceive struct {
	t         gocuke.TestingT
	msg       *MsgBridgeReceive
	err       error
	signBytes string
}

func TestMsgBridgeReceive(t *testing.T) {
	gocuke.NewRunner(t, &msgBridgeReceive{}).Path("./features/msg_bridge_receive.feature").Run()
}

func (s *msgBridgeReceive) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgBridgeReceive) TheMessage(a gocuke.DocString) {
	s.msg = &MsgBridgeReceive{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgBridgeReceive) AProjectReferenceIdWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Project = &MsgBridgeReceive_Project{}
	s.msg.Project.ReferenceId = strings.Repeat("x", int(length))
}

func (s *msgBridgeReceive) ProjectMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Project.Metadata = strings.Repeat("x", int(length))
}

func (s *msgBridgeReceive) BatchMetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Batch.Metadata = strings.Repeat("x", int(length))
}

func (s *msgBridgeReceive) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgBridgeReceive) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgBridgeReceive) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgBridgeReceive) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgBridgeReceive) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
