package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type allowedDenom struct {
	t            gocuke.TestingT
	allowedDenom *AllowedDenom
	err          error
}

func TestAllowedDenom(t *testing.T) {
	gocuke.NewRunner(t, &allowedDenom{}).Path("./features/state_allowed_denom.feature").Run()
}

func (s *allowedDenom) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *allowedDenom) TheAllowedDenom(a gocuke.DocString) {
	s.allowedDenom = &AllowedDenom{}
	err := jsonpb.UnmarshalString(a.Content, s.allowedDenom)
	require.NoError(s.t, err)
}

func (s *allowedDenom) TheAllowedDenomIsValidated() {
	s.err = s.allowedDenom.Validate()
}

func (s *allowedDenom) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *allowedDenom) ExpectNoError() {
	require.NoError(s.t, s.err)
}
