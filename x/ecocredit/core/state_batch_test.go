package core

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type batch struct {
	t     gocuke.TestingT
	batch *Batch
	err   error
}

func TestBatch(t *testing.T) {
	gocuke.NewRunner(t, &batch{}).Path("./features/state_batch.feature").Run()
}

func (s *batch) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *batch) TheBatch(a gocuke.DocString) {
	s.batch = &Batch{}
	err := jsonpb.UnmarshalString(a.Content, s.batch)
	require.NoError(s.t, err)
}

func (s *batch) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.batch.Metadata = strings.Repeat("x", int(length))
}

func (s *batch) TheBatchIsValidated() {
	s.err = s.batch.Validate()
}

func (s *batch) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *batch) ExpectNoError() {
	require.NoError(s.t, s.err)
}
