package core

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgCreateProject struct {
	t   gocuke.TestingT
	msg *MsgCreateProject
	err error
}

func TestMsgCreateProject(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateProject{}).Path("./features/msg_create_project.feature").Run()
}

func (s *msgCreateProject) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
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
