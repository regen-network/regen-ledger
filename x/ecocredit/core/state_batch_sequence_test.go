package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type batchSequence struct {
	t             gocuke.TestingT
	batchSequence *BatchSequence
	err           error
}

func TestBatchSequence(t *testing.T) {
	gocuke.NewRunner(t, &batchSequence{}).Path("./features/state_batch_sequence.feature").Run()
}

func (s *batchSequence) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *batchSequence) TheBatchSequence(a gocuke.DocString) {
	s.batchSequence = &BatchSequence{}
	err := jsonpb.UnmarshalString(a.Content, s.batchSequence)
	require.NoError(s.t, err)
}

func (s *batchSequence) TheBatchSequenceIsValidated() {
	s.err = s.batchSequence.Validate()
}

func (s *batchSequence) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *batchSequence) ExpectNoError() {
	require.NoError(s.t, s.err)
}
