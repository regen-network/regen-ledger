package core

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type mintBatchCredits struct {
	*baseSuite
	alice            sdk.AccAddress
	creditTypeAbbrev string
	classId          string
	classKey         uint64
	projectId        string
	projectKey       uint64
	batchDenom       string
	originTx         *core.OriginTx
	tradableAmount   string
	res              *core.MsgMintBatchCreditsResponse
	err              error
}

func TestMintBatchCredits(t *testing.T) {
	gocuke.NewRunner(t, &mintBatchCredits{}).Path("./features/msg_mint_batch_credits.feature").Run()
}

func (s *mintBatchCredits) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
	s.projectId = "C01-001"
	s.batchDenom = "C01-001-20200101-20210101-001"
	s.originTx = &core.OriginTx{
		Id:     "0x0",
		Source: "polygon",
	}
}

func (s *mintBatchCredits) ACreditType() {
	// TODO: Save for now but credit type should not exist prior to unit test #893
	err := s.k.stateStore.CreditTypeTable().Save(s.ctx, &api.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Name:         s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *mintBatchCredits) ACreditClassWithIssuerAlice() {
	cKey, err := s.k.stateStore.ClassTable().InsertReturningID(s.ctx, &api.Class{
		Id:               s.classId,
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

func (s *mintBatchCredits) AProject() {
	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       s.projectId,
		ClassKey: s.classKey,
	})
	require.NoError(s.t, err)

	s.projectKey = pKey
}

func (s *mintBatchCredits) ACreditBatchWithOpenAndIssuerAlice(a string) {
	open, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	bKey, err := s.k.stateStore.BatchTable().InsertReturningID(s.ctx, &api.Batch{
		Issuer:     s.alice,
		Denom:      s.batchDenom,
		ProjectKey: s.projectKey,
		Open:       open,
	})
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchSupplyTable().Insert(s.ctx, &api.BatchSupply{
		BatchKey: bKey,
	})
	require.NoError(s.t, err)

	seq := s.getBatchSequence(s.batchDenom)

	// Save because batch sequence may already exist
	err = s.k.stateStore.BatchSequenceTable().Save(s.ctx, &api.BatchSequence{
		ProjectKey:   s.projectKey,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)
}

func (s *mintBatchCredits) AnOriginTxIndex(a gocuke.DocString) {
	originTxIndex := &api.OriginTxIndex{}
	err := jsonpb.UnmarshalString(a.Content, originTxIndex)
	require.NoError(s.t, err)

	err = s.k.stateStore.OriginTxIndexTable().Insert(s.ctx, originTxIndex)
	require.NoError(s.t, err)
}

func (s *mintBatchCredits) AliceAttemptsToMintCredits() {
	s.res, s.err = s.k.MintBatchCredits(s.ctx, &core.MsgMintBatchCredits{
		Issuer:     s.alice.String(),
		BatchDenom: s.batchDenom,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.alice.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		OriginTx: s.originTx,
	})
}

func (s *mintBatchCredits) AliceAttemptsToMintCreditsWithOriginTx(a gocuke.DocString) {
	originTx := &core.OriginTx{}
	err := jsonpb.UnmarshalString(a.Content, originTx)
	require.NoError(s.t, err)

	s.res, s.err = s.k.MintBatchCredits(s.ctx, &core.MsgMintBatchCredits{
		Issuer:     s.alice.String(),
		BatchDenom: s.batchDenom,
		Issuance: []*core.BatchIssuance{
			{
				Recipient:      s.alice.String(),
				TradableAmount: s.tradableAmount,
			},
		},
		OriginTx: originTx,
	})
}

func (s *mintBatchCredits) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *mintBatchCredits) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *mintBatchCredits) getBatchSequence(batchDenom string) uint64 {
	str := strings.Split(batchDenom, "-")
	seq, err := strconv.ParseUint(str[4], 10, 32)
	require.NoError(s.t, err)
	return seq
}
