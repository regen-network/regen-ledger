package core

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type credits struct {
	t        gocuke.TestingT
	issuance *Credits
	err      error
}

func TestCredits(t *testing.T) {
	gocuke.NewRunner(t, &credits{}).Path("./features/types_credits.feature").Run()
}

func (s *credits) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *credits) TheMessage(a gocuke.DocString) {
	s.issuance = &Credits{}
	err := jsonpb.UnmarshalString(a.Content, s.issuance)
	require.NoError(s.t, err)
}

func (s *credits) TheMessageIsValidated() {
	s.err = s.issuance.Validate()
}

func (s *credits) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *credits) ExpectNoError() {
	require.NoError(s.t, s.err)
}
