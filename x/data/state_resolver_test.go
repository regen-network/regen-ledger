package data

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type resolver struct {
	t        gocuke.TestingT
	resolver *Resolver
	err      error
}

func TestResolver(t *testing.T) {
	gocuke.NewRunner(t, &resolver{}).Path("./features/state_resolver.feature").Run()
}

func (s *resolver) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *resolver) TheResolver(a gocuke.DocString) {
	s.resolver = &Resolver{}
	err := jsonpb.UnmarshalString(a.Content, s.resolver)
	require.NoError(s.t, err)
}

func (s *resolver) TheResolverIsValidated() {
	s.err = s.resolver.Validate()
}

func (s *resolver) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *resolver) ExpectNoError() {
	require.NoError(s.t, s.err)
}
