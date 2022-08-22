package core

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type originTxIndex struct {
	t             gocuke.TestingT
	originTxIndex *OriginTxIndex
	err           error
}

func TestOriginTxIndex(t *testing.T) {
	gocuke.NewRunner(t, &originTxIndex{}).Path("./features/state_origin_tx_index.feature").Run()
}

func (s *originTxIndex) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *originTxIndex) TheOriginTxIndex(a gocuke.DocString) {
	s.originTxIndex = &OriginTxIndex{}
	err := jsonpb.UnmarshalString(a.Content, s.originTxIndex)
	require.NoError(s.t, err)
}

func (s *originTxIndex) IdWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.originTxIndex.Id = strings.Repeat("x", int(length))
}

func (s *originTxIndex) SourceWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.originTxIndex.Source = strings.Repeat("x", int(length))
}

func (s *originTxIndex) TheOriginTxIndexIsValidated() {
	s.err = s.originTxIndex.Validate()
}

func (s *originTxIndex) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *originTxIndex) ExpectNoError() {
	require.NoError(s.t, s.err)
}
