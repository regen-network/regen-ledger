//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type sealBatch struct {
	*baseSuite
	alice            sdk.AccAddress
	bob              sdk.AccAddress
	creditTypeAbbrev string
	classKey         uint64
	res              *types.MsgSealBatchResponse
	err              error
	projectKey       uint64
}

func TestSealBatch(t *testing.T) {
	gocuke.NewRunner(t, &sealBatch{}).Path("./features/msg_seal_batch.feature").Run()
}

func (s *sealBatch) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
}

func (s *sealBatch) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
}

func (s *sealBatch) ACreditClassWithIdAndIssuerAlice(a string) {
	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               a,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: cKey,
		Issuer:   s.alice,
	})
	require.NoError(s.t, err)

	s.classKey = cKey
}

func (s *sealBatch) AProjectWithId(a string) {
	var err error
	s.projectKey, err = s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id: a,
	})
	require.NoError(s.t, err)
}

func (s *sealBatch) ACreditBatchWithDenomAndIssuerAlice(a string) {
	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: s.projectKey,
		Issuer:     s.alice,
		Denom:      a,
		Open:       true,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        bKey,
		TradableAmount:  "0",
		RetiredAmount:   "0",
		CancelledAmount: "0",
	})
	require.NoError(s.t, err)
}

func (s *sealBatch) AliceAttemptsToSealBatchWithDenom(a string) {
	s.res, s.err = s.k.SealBatch(s.ctx, &types.MsgSealBatch{
		Issuer:     s.alice.String(),
		BatchDenom: a,
	})
}

func (s *sealBatch) BobAttemptsToSealBatchWithDenom(a string) {
	s.res, s.err = s.k.SealBatch(s.ctx, &types.MsgSealBatch{
		Issuer:     s.bob.String(),
		BatchDenom: a,
	})
}

func (s *sealBatch) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *sealBatch) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *sealBatch) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventSealBatch
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}
