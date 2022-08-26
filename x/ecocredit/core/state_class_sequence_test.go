package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type classSequence struct {
	t             gocuke.TestingT
	classSequence *ClassSequence
	err           error
}

func TestClassSequence(t *testing.T) {
	gocuke.NewRunner(t, &classSequence{}).Path("./features/state_class_sequence.feature").Run()
}

func (s *classSequence) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *classSequence) TheClassSequence(a gocuke.DocString) {
	s.classSequence = &ClassSequence{}
	err := jsonpb.UnmarshalString(a.Content, s.classSequence)
	require.NoError(s.t, err)
}

func (s *classSequence) TheClassSequenceIsValidated() {
	s.err = s.classSequence.Validate()
}

func (s *classSequence) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *classSequence) ExpectNoError() {
	require.NoError(s.t, s.err)
}
