//nolint:revive,stylecheck
package data

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type dataID struct {
	t      gocuke.TestingT
	dataID *DataID
	err    error
}

func TestDataID(t *testing.T) {
	gocuke.NewRunner(t, &dataID{}).Path("./features/state_data_id.feature").Run()
}

func (s *dataID) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *dataID) TheDataId(a gocuke.DocString) {
	s.dataID = &DataID{}
	err := jsonpb.UnmarshalString(a.Content, s.dataID)
	require.NoError(s.t, err)
}

func (s *dataID) TheDataIdIsValidated() {
	s.err = s.dataID.Validate()
}

func (s *dataID) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *dataID) ExpectNoError() {
	require.NoError(s.t, s.err)
}
