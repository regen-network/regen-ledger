package data

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type contentHash struct {
	t   gocuke.TestingT
	ch  *ContentHash
	err error
}

func TestTypes(t *testing.T) {
	runner := gocuke.NewRunner(t, &contentHash{}).Path("./features/types_content_hash.feature")
	runner.Step(`^the\s+content\s+hash\s+"((?:[^\"]|\")*)"`, (*contentHash).TheContentHash)
	runner.Run()
}

func (s *contentHash) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *contentHash) TheContentHash(a gocuke.DocString) {
	s.ch = &ContentHash{}
	err := jsonpb.UnmarshalString(a.Content, s.ch)
	require.NoError(s.t, err)
}

func (s *contentHash) TheContentHashIsValidated() {
	s.err = s.ch.Validate()
}

func (s *contentHash) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *contentHash) ExpectNoError() {
	require.NoError(s.t, s.err)
}
