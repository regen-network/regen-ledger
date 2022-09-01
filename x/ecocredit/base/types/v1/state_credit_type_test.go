package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type creditType struct {
	t          gocuke.TestingT
	creditType *CreditType
	err        error
}

func TestCreditType(t *testing.T) {
	gocuke.NewRunner(t, &creditType{}).Path("./features/state_credit_type.feature").Run()
}

func (s *creditType) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *creditType) TheCreditType(a gocuke.DocString) {
	s.creditType = &CreditType{}
	err := jsonpb.UnmarshalString(a.Content, s.creditType)
	require.NoError(s.t, err)
}

func (s *creditType) NameWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.creditType.Name = strings.Repeat("x", int(length))
}

func (s *creditType) TheCreditTypeIsValidated() {
	s.err = s.creditType.Validate()
}

func (s *creditType) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *creditType) ExpectNoError() {
	require.NoError(s.t, s.err)
}
