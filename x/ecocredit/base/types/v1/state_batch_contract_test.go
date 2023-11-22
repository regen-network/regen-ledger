package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type batchContract struct {
	t             gocuke.TestingT
	batchContract *BatchContract
	err           error
}

func TestBatchContract(t *testing.T) {
	gocuke.NewRunner(t, &batchContract{}).Path("./features/state_batch_contract.feature").Run()
}

func (s *batchContract) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *batchContract) TheBatchContract(a gocuke.DocString) {
	s.batchContract = &BatchContract{}
	err := jsonpb.UnmarshalString(a.Content, s.batchContract)
	require.NoError(s.t, err)
}

func (s *batchContract) TheBatchContractIsValidated() {
	s.err = s.batchContract.Validate()
}

func (s *batchContract) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *batchContract) ExpectNoError() {
	require.NoError(s.t, s.err)
}
