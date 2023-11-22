package data

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type dataResolver struct {
	t            gocuke.TestingT
	dataResolver *DataResolver
	err          error
}

func TestDataResolver(t *testing.T) {
	gocuke.NewRunner(t, &dataResolver{}).Path("./features/state_data_resolver.feature").Run()
}

func (s *dataResolver) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *dataResolver) TheDataResolver(a gocuke.DocString) {
	s.dataResolver = &DataResolver{}
	err := jsonpb.UnmarshalString(a.Content, s.dataResolver)
	require.NoError(s.t, err)
}

func (s *dataResolver) TheDataResolverIsValidated() {
	s.err = s.dataResolver.Validate()
}

func (s *dataResolver) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *dataResolver) ExpectNoError() {
	require.NoError(s.t, s.err)
}
