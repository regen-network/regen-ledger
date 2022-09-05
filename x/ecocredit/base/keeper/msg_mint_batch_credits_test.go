//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

type mintBatchCredits struct {
	*baseSuite
	alice            sdk.AccAddress
	bob              sdk.AccAddress
	creditTypeAbbrev string
	classKey         uint64
	projectKey       uint64
	batchDenom       string
	originTx         *types.OriginTx
	tradableAmount   string
	res              *types.MsgMintBatchCreditsResponse
	err              error
}

func TestMintBatchCredits(t *testing.T) {
	gocuke.NewRunner(t, &mintBatchCredits{}).Path("./features/msg_mint_batch_credits.feature").Run()
}

func (s *mintBatchCredits) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
	s.originTx = &types.OriginTx{
		Id:     "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
		Source: "polygon",
	}
}

func (s *mintBatchCredits) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
}

func (s *mintBatchCredits) ACreditClassWithIdAndIssuerAlice(a string) {
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

func (s *mintBatchCredits) AProjectWithId(a string) {
	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       a,
		ClassKey: s.classKey,
	})
	require.NoError(s.t, err)

	s.projectKey = pKey
}

func (s *mintBatchCredits) ACreditBatchWithDenomAndIssuerAlice(a string) {
	projectID := base.GetProjectIDFromBatchDenom(a)

	project, err := s.k.stateStore.ProjectTable().GetById(s.ctx, projectID)
	require.NoError(s.t, err)

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		ProjectKey: project.Key,
		Issuer:     s.alice,
		Denom:      a,
		Open:       true, // always true unless specified
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        bKey,
		TradableAmount:  "0",
		RetiredAmount:   "0",
		CancelledAmount: "0",
	})
	require.NoError(s.t, err)

	s.batchDenom = a
}

func (s *mintBatchCredits) ACreditBatchWithDenomOpenAndIssuerAlice(a, b string) {
	open, err := strconv.ParseBool(b)
	require.NoError(s.t, err)

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		Issuer:     s.alice,
		Denom:      a,
		ProjectKey: s.projectKey,
		Open:       open,
	})
	require.NoError(s.t, err)

	seq := s.getBatchSequence(a)

	err = s.k.stateStore.BatchSequenceTable().Insert(s.ctx, &api.BatchSequence{
		ProjectKey:   s.projectKey,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey:        bKey,
		TradableAmount:  "0",
		RetiredAmount:   "0",
		CancelledAmount: "0",
	})
	require.NoError(s.t, err)

	s.batchDenom = a
}

func (s *mintBatchCredits) AnOriginTxIndex(a gocuke.DocString) {
	originTxIndex := &api.OriginTxIndex{}
	err := jsonpb.UnmarshalString(a.Content, originTxIndex)
	require.NoError(s.t, err)

	err = s.k.stateStore.OriginTxIndexTable().Insert(s.ctx, originTxIndex)
	require.NoError(s.t, err)
}

func (s *mintBatchCredits) AliceAttemptsToMintCreditsWithBatchDenom(a string) {
	s.res, s.err = s.k.MintBatchCredits(s.ctx, &types.MsgMintBatchCredits{
		Issuer:     s.alice.String(),
		BatchDenom: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.bob.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		OriginTx: s.originTx,
	})
}

func (s *mintBatchCredits) BobAttemptsToMintCreditsWithBatchDenom(a string) {
	s.res, s.err = s.k.MintBatchCredits(s.ctx, &types.MsgMintBatchCredits{
		Issuer:     s.bob.String(),
		BatchDenom: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.alice.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		OriginTx: s.originTx,
	})
}

func (s *mintBatchCredits) AliceAttemptsToMintCreditsWithBatchDenomAndOriginTx(a string, b gocuke.DocString) {
	originTx := &types.OriginTx{}
	err := jsonpb.UnmarshalString(b.Content, originTx)
	require.NoError(s.t, err)

	s.res, s.err = s.k.MintBatchCredits(s.ctx, &types.MsgMintBatchCredits{
		Issuer:     s.alice.String(),
		BatchDenom: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.bob.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		OriginTx: originTx,
	})
}

func (s *mintBatchCredits) AliceAttemptsToMintCreditsWithBatchDenomAndTradableAmount(a, b string) {
	s.res, s.err = s.k.MintBatchCredits(s.ctx, &types.MsgMintBatchCredits{
		Issuer:     s.alice.String(),
		BatchDenom: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.bob.String(),
				TradableAmount: b,
			},
		},
		OriginTx: s.originTx,
	})
}

func (s *mintBatchCredits) AliceAttemptsToMintCreditsWithBatchDenomRecipientBobAndTradableAmount(a, b string) {
	s.res, s.err = s.k.MintBatchCredits(s.ctx, &types.MsgMintBatchCredits{
		Issuer:     s.alice.String(),
		BatchDenom: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      s.bob.String(),
				TradableAmount: b,
			},
		},
		OriginTx: s.originTx,
	})
}

func (s *mintBatchCredits) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *mintBatchCredits) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *mintBatchCredits) ExpectBobBatchBalance(a gocuke.DocString) {
	var expected api.BatchBalance
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.bob, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.EscrowedAmount, balance.EscrowedAmount)
}

func (s *mintBatchCredits) ExpectBatchSupply(a gocuke.DocString) {
	var expected api.BatchSupply
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchSupplyTable().Get(s.ctx, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.CancelledAmount, balance.CancelledAmount)
}

func (s *mintBatchCredits) AliceAttemptsToMintCreditsWithBatchDenomWithRetiredAmountFromTo(a, b, c, d string) {
	s.res, s.err = s.k.MintBatchCredits(s.ctx, &types.MsgMintBatchCredits{
		Issuer:     s.alice.String(),
		BatchDenom: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:              d,
				RetiredAmount:          b,
				RetirementJurisdiction: c,
			},
		},
		OriginTx: s.originTx,
	})
	require.NoError(s.t, s.err)
}

func (s *mintBatchCredits) ExpectEventRetireWithProperties(a gocuke.DocString) {
	var event types.EventRetire
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *mintBatchCredits) ExpectEventMintWithProperties(a gocuke.DocString) {
	var event types.EventMint
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *mintBatchCredits) ExpectEventMintBatchCreditsWithProperties(a gocuke.DocString) {
	var event types.EventMintBatchCredits
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *mintBatchCredits) AnOriginTxWithProperties(a gocuke.DocString) {
	var ot types.OriginTx
	err := json.Unmarshal([]byte(a.Content), &ot)
	require.NoError(s.t, err)
	s.originTx = &ot
}

func (s *mintBatchCredits) ExpectEventTransferWithProperties(a gocuke.DocString) {
	var event types.EventTransfer
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)
	event.Sender = s.k.moduleAddress.String()

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *mintBatchCredits) AliceAttemptsToMintCreditsWithBatchDenomWithTradableAmountTo(a string, b string, c string) {
	s.res, s.err = s.k.MintBatchCredits(s.ctx, &types.MsgMintBatchCredits{
		Issuer:     s.alice.String(),
		BatchDenom: a,
		Issuance: []*types.BatchIssuance{
			{
				Recipient:      c,
				TradableAmount: b,
			},
		},
		OriginTx: s.originTx,
	})
	require.NoError(s.t, s.err)
}

func (s *mintBatchCredits) EcocreditModuleAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.k.moduleAddress = addr
}

func (s *mintBatchCredits) getBatchSequence(batchDenom string) uint64 {
	str := strings.Split(batchDenom, "-")
	seq, err := strconv.ParseUint(str[4], 10, 32)
	require.NoError(s.t, err)
	return seq
}
