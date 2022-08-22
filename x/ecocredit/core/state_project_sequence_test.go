package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type projectSequence struct {
	t               gocuke.TestingT
	projectSequence *ProjectSequence
	err             error
}

func TestProjectSequence(t *testing.T) {
	gocuke.NewRunner(t, &projectSequence{}).Path("./features/state_project_sequence.feature").Run()
}

func (s *projectSequence) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *projectSequence) TheProjectSequence(a gocuke.DocString) {
	s.projectSequence = &ProjectSequence{}
	err := jsonpb.UnmarshalString(a.Content, s.projectSequence)
	require.NoError(s.t, err)
}

func (s *projectSequence) TheProjectSequenceIsValidated() {
	s.err = s.projectSequence.Validate()
}

func (s *projectSequence) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *projectSequence) ExpectNoError() {
	require.NoError(s.t, s.err)
}
