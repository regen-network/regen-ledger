package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type batchSupply struct {
	t           gocuke.TestingT
	batchSupply *BatchSupply
	err         error
}

func TestBatchSupply(t *testing.T) {
	gocuke.NewRunner(t, &batchSupply{}).Path("./features/state_batch_supply.feature").Run()
}

func (s *batchSupply) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *batchSupply) TheBatchSupply(a gocuke.DocString) {
	s.batchSupply = &BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, s.batchSupply)
	require.NoError(s.t, err)
}

func (s *batchSupply) TheBatchSupplyIsValidated() {
	s.err = s.batchSupply.Validate()
}

func (s *batchSupply) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *batchSupply) ExpectNoError() {
	require.NoError(s.t, s.err)
}
