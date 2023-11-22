package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type classIssuer struct {
	t           gocuke.TestingT
	classIssuer *ClassIssuer
	err         error
}

func TestClassIssuer(t *testing.T) {
	gocuke.NewRunner(t, &classIssuer{}).Path("./features/state_class_issuer.feature").Run()
}

func (s *classIssuer) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *classIssuer) TheClassIssuer(a gocuke.DocString) {
	s.classIssuer = &ClassIssuer{}
	err := jsonpb.UnmarshalString(a.Content, s.classIssuer)
	require.NoError(s.t, err)
}

func (s *classIssuer) TheClassIssuerIsValidated() {
	s.err = s.classIssuer.Validate()
}

func (s *classIssuer) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *classIssuer) ExpectNoError() {
	require.NoError(s.t, s.err)
}
