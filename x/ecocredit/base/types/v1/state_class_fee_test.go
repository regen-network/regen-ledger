package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type classFeeSuite struct {
	t        gocuke.TestingT
	classFee *ClassFee
	err      error
}

func TestClassFee(t *testing.T) {
	gocuke.NewRunner(t, &classFeeSuite{}).Path("./features/state_class_fee.feature").Run()
}

func (s *classFeeSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *classFeeSuite) TheClassFee(a gocuke.DocString) {
	s.classFee = &ClassFee{}
	err := jsonpb.UnmarshalString(a.Content, s.classFee)
	require.NoError(s.t, err)
}

func (s *classFeeSuite) TheClassFeeIsValidated() {
	s.err = s.classFee.Validate()
}

func (s *classFeeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *classFeeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
