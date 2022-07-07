package core

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
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

func (s *msgCreateProject) TheMessageIsValidated() {
	s.checkAndSetMockValues()

	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateProject) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCreateProject) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *msgCreateProject) checkAndSetMockValues() {
	if strings.Contains(s.msg.Metadata, "[mock-string-257]") {
		s.msg.Metadata = strings.Repeat("x", 257)
	}
	if strings.Contains(s.msg.ReferenceId, "[mock-string-33]") {
		s.msg.ReferenceId = strings.Repeat("x", 33)
	}
}
