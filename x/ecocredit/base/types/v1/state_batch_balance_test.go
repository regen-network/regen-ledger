package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type batchBalance struct {
	t            gocuke.TestingT
	batchBalance *BatchBalance
	err          error
}

func TestBatchBalance(t *testing.T) {
	gocuke.NewRunner(t, &batchBalance{}).Path("./features/state_batch_balance.feature").Run()
}

func (s *batchBalance) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *batchBalance) TheBatchBalance(a gocuke.DocString) {
	s.batchBalance = &BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, s.batchBalance)
	require.NoError(s.t, err)
}

func (s *batchBalance) TheBatchBalanceIsValidated() {
	s.err = s.batchBalance.Validate()
}

func (s *batchBalance) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *batchBalance) ExpectNoError() {
	require.NoError(s.t, s.err)
}
