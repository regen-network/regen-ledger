package v1

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreateOrUpdateApplication struct {
	gocuke.TestingT
	msg *MsgCreateOrUpdateApplication
	err error
}

func TestMsgCreateOrUpdateApplication(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateOrUpdateApplication{}).Path("./features/msg_create_or_update_application.feature").Run()
}

func (s *msgCreateOrUpdateApplication) Before() {
	s.msg = &MsgCreateOrUpdateApplication{}
}

func (s *msgCreateOrUpdateApplication) ProjectAdmin(a string) {
	s.msg.ProjectAdmin = a
}

func (s *msgCreateOrUpdateApplication) ProjectId(a string) {
	s.msg.ProjectId = a
}

func (s *msgCreateOrUpdateApplication) ClassId(a string) {
	s.msg.ClassId = a
}

func (s *msgCreateOrUpdateApplication) Metadata(a string) {
	s.msg.Metadata = a
}

func (s *msgCreateOrUpdateApplication) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateOrUpdateApplication) ExpectErrorContains(a string) {
	if a != "" {
		require.ErrorContains(s, s.err, a)
	} else {
		require.NoError(s, s.err)
	}
}

func (s *msgCreateOrUpdateApplication) ExpectGetsignersReturns(a string) {
	signers := s.msg.GetSigners()
	require.Len(s, signers, 1)
	require.Equal(s, a, s.msg.GetSigners()[0].String())
}
