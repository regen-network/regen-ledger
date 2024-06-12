package v1

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type msgCreateUnregisteredProject struct {
	gocuke.TestingT
	msg *MsgCreateUnregisteredProject
	err error
}

func TestMsgCreateUnregisteredProject(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateUnregisteredProject{}).Path("./features/msg_create_unregistered_project.feature").Run()
}

func (s *msgCreateUnregisteredProject) Before() {
	s.msg = &MsgCreateUnregisteredProject{}
}

func (s *msgCreateUnregisteredProject) Admin(a string) {
	s.msg.Admin = a
}

func (s *msgCreateUnregisteredProject) Jurisdiction(a string) {
	s.msg.Jurisdiction = a
}

func (s *msgCreateUnregisteredProject) Metadata(a string) {
	s.msg.Metadata = a
}

func (s *msgCreateUnregisteredProject) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateUnregisteredProject) ExpectErrorContains(a string) {
	if a != "" {
		require.ErrorContains(s, s.err, a)
	} else {
		require.NoError(s, s.err)
	}
}

func (s *msgCreateUnregisteredProject) ExpectGetsignersReturns(a string) {
	signers := s.msg.GetSigners()
	require.Len(s, signers, 1)
	require.Equal(s, a, s.msg.GetSigners()[0].String())
}
