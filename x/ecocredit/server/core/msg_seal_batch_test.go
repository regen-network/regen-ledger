package core

import (
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type sealBatch struct {
	*baseSuite
	alice            sdk.AccAddress
	creditTypeAbbrev string
	classId          string
	classKey         uint64
	projectId        string
	projectKey       uint64
	batchDenom       string
	batchKey         uint64
	res              *core.MsgSealBatchResponse
	err              error
}

func TestSealBatch(t *testing.T) {
	gocuke.NewRunner(t, &sealBatch{}).Path("./features/msg_seal_batch.feature").Run()
}

func (s *sealBatch) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
	s.projectId = "C01-001"
	s.batchDenom = "C01-001-20200101-20210101-001"
}

func (s *sealBatch) ACreditTypeWithAbbreviationAndPrecision(a, b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	err = s.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
}
