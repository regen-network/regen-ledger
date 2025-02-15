//nolint:revive,stylecheck
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

type msgCreateProject struct {
	t         gocuke.TestingT
	msg       *MsgCreateProject
	err       error
	signBytes string
}

func TestMsgCreateProject(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateProject{}).Path("./features/msg_create_project.feature").Run()
}

func (s *msgCreateProject) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *msgCreateProject) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCreateProject{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateProject) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgCreateProject) AReferenceIdWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.ReferenceId = strings.Repeat("x", int(length))
}

func (s *msgCreateProject) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateProject) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCreateProject) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateProject) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *msgCreateProject) ExpectTheSignBytes(expected gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(expected.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
