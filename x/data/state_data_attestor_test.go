package data

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type dataAttestor struct {
	t            gocuke.TestingT
	dataAttestor *DataAttestor
	err          error
}

func TestDataAttestor(t *testing.T) {
	gocuke.NewRunner(t, &dataAttestor{}).Path("./features/state_data_attestor.feature").Run()
}

func (s *dataAttestor) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *dataAttestor) TheDataAttestor(a gocuke.DocString) {
	s.dataAttestor = &DataAttestor{}
	err := jsonpb.UnmarshalString(a.Content, s.dataAttestor)
	require.NoError(s.t, err)
}

func (s *dataAttestor) TheDataAttestorIsValidated() {
	s.err = s.dataAttestor.Validate()
}

func (s *dataAttestor) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *dataAttestor) ExpectNoError() {
	require.NoError(s.t, s.err)
}
