package data

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type dataAnchor struct {
	t          gocuke.TestingT
	dataAnchor *DataAnchor
	err        error
}

func TestDataAnchor(t *testing.T) {
	gocuke.NewRunner(t, &dataAnchor{}).Path("./features/state_data_anchor.feature").Run()
}

func (s *dataAnchor) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *dataAnchor) TheDataAnchor(a gocuke.DocString) {
	s.dataAnchor = &DataAnchor{}
	err := jsonpb.UnmarshalString(a.Content, s.dataAnchor)
	require.NoError(s.t, err)
}

func (s *dataAnchor) TheDataAnchorIsValidated() {
	s.err = s.dataAnchor.Validate()
}

func (s *dataAnchor) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *dataAnchor) ExpectNoError() {
	require.NoError(s.t, s.err)
}
