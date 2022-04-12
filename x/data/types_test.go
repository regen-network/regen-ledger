package data

import (
	"encoding/json"
	"testing"

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
	runner.Step(`the content hash "((?:[^\"]|\")*)"`, (*contentHash).TheContentHash)
	runner.Run()
}

func (s *contentHash) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *contentHash) TheContentHash(a gocuke.DocString) {
	err := json.Unmarshal([]byte(a.Content), &s.ch)
	require.NoError(s.t, err)
}

func (s *contentHash) TheContentHashIsValidated() {
	s.err = s.ch.Validate()
}

func (s *contentHash) ExpectTheError(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.EqualError(s.t, s.err, a)
	}
}
