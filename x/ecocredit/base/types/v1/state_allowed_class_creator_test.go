package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type allowedClassCreatorSuite struct {
	t                   gocuke.TestingT
	allowedClassCreator *AllowedClassCreator
	err                 error
}

func TestAllowedClassCreator(t *testing.T) {
	gocuke.NewRunner(t, &allowedClassCreatorSuite{}).Path(
		"./features/state_allowed_class_creator.feature",
	).Run()
}

func (s *allowedClassCreatorSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *allowedClassCreatorSuite) TheAllowedClassCreator(a gocuke.DocString) {
	s.allowedClassCreator = &AllowedClassCreator{}
	err := jsonpb.UnmarshalString(a.Content, s.allowedClassCreator)
	require.NoError(s.t, err)
}

func (s *allowedClassCreatorSuite) TheAllowedClassCreatorIsValidated() {
	s.err = s.allowedClassCreator.Validate()
}

func (s *allowedClassCreatorSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *allowedClassCreatorSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
