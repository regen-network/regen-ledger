//nolint:revive,stylecheck
package keeper

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
)

type bridgeReceiveSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	bob              sdk.AccAddress
	creditTypeAbbrev string
	classID          string
	classKey         uint64
	projectKey       uint64
	referenceID      string
	metadata         string
	jurisdiction     string
	batchKey         uint64
	tradableAmount   string
	startDate        *time.Time
	endDate          *time.Time
	originTx         *types.OriginTx
	res              *types.MsgBridgeReceiveResponse
	err              error
}

func TestBridgeReceive(t *testing.T) {
	gocuke.NewRunner(t, &bridgeReceiveSuite{}).Path("./features/msg_bridge_receive.feature").Run()
}

func (s *bridgeReceiveSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
	s.classID = testClassID
	s.jurisdiction = "US-WA"
	s.referenceID = "VCS-001"
	s.metadata = "regen:113gdjFKcVCt13Za6vN7TtbgMM6LMSjRnu89BMCxeuHdkJ1hWUmy.rdf"
	s.tradableAmount = "10"

	startDate, err := regentypes.ParseDate("start date", "2020-01-01")
	require.NoError(s.t, err)

	endDate, err := regentypes.ParseDate("end date", "2021-01-01")
	require.NoError(s.t, err)

	s.startDate = &startDate
	s.endDate = &endDate

	s.originTx = &types.OriginTx{
		Id:       "0x7a70692a348e8688f54ab2bdfe87d925d8cc88932520492a11eaa02dc128243e",
		Source:   "polygon",
		Contract: "0x0E65079a29d7793ab5CA500c2d88e60EE99Ba606",
	}
}

func (s *bridgeReceiveSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.k.stateStore.CreditTypeTable().Insert(s.ctx, &api.CreditType{
		Abbreviation: a,
		Name:         a,
	})
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
}

func (s *bridgeReceiveSuite) ACreditClassWithIdAndIssuerAlice(a string) {
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

func (s *bridgeReceiveSuite) AProjectWithId(a string) {
	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:       a,
		ClassKey: s.classKey,
	})
	require.NoError(s.t, err)

	seq := s.getProjectSequence(a)

	err = s.k.stateStore.ProjectSequenceTable().Insert(s.ctx, &api.ProjectSequence{
		ClassKey:     s.classKey,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)

	s.projectKey = pKey
}

func (s *bridgeReceiveSuite) AProjectWithIdAndReferenceId(a, b string) {
	pKey, err := s.k.stateStore.ProjectTable().InsertReturningID(s.ctx, &api.Project{
		Id:          a,
		ClassKey:    s.classKey,
		ReferenceId: b,
	})
	require.NoError(s.t, err)

	seq := s.getProjectSequence(a)

	err = s.k.stateStore.ProjectSequenceTable().Insert(s.ctx, &api.ProjectSequence{
		ClassKey:     s.classKey,
		NextSequence: seq + 1,
	})
	require.NoError(s.t, err)

	s.projectKey = pKey
}

func (s *bridgeReceiveSuite) ACreditBatchWithDenomAndIssuerAlice(a string) {
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

	seq := s.getBatchSequence(a)

	err = s.k.stateStore.BatchSequenceTable().Insert(s.ctx, &api.BatchSequence{
		ProjectKey:   project.Key,
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

	s.batchKey = bKey
}

func (s *bridgeReceiveSuite) TheBatchContract(a gocuke.DocString) {
	var batchContract api.BatchContract
	err := jsonpb.UnmarshalString(a.Content, &batchContract)
	require.NoError(s.t, err)

	err = s.k.stateStore.BatchContractTable().Insert(s.ctx, &batchContract)
	require.NoError(s.t, err)
}

func (s *bridgeReceiveSuite) AllowedBridgeChain(a string) {
	_, err := s.k.AddAllowedBridgeChain(s.ctx, &types.MsgAddAllowedBridgeChain{
		Authority: s.authority.String(),
		ChainName: a,
	})
	require.NoError(s.t, err)
}

func (s *bridgeReceiveSuite) AliceAttemptsToBridgeCreditsWithClassId(a string) {
	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.alice.String(),
		ClassId: a,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceID,
			Jurisdiction: s.jurisdiction,
			Metadata:     s.metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    s.tradableAmount,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: s.originTx,
	})
}

func (s *bridgeReceiveSuite) BobAttemptsToBridgeCreditsWithClassId(a string) {
	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.bob.String(),
		ClassId: a,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceID,
			Jurisdiction: s.jurisdiction,
			Metadata:     s.metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.alice.String(),
			Amount:    s.tradableAmount,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: s.originTx,
	})
}

func (s *bridgeReceiveSuite) AliceAttemptsToBridgeCreditsWithContract(a string) {
	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.alice.String(),
		ClassId: s.classID,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceID,
			Jurisdiction: s.jurisdiction,
			Metadata:     s.metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    s.tradableAmount,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: &types.OriginTx{
			Id:       s.originTx.Id,
			Source:   s.originTx.Source,
			Contract: a,
		},
	})
}

func (s *bridgeReceiveSuite) BobAttemptsToBridgeCreditsWithContract(a string) {
	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.bob.String(),
		ClassId: s.classID,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceID,
			Jurisdiction: s.jurisdiction,
			Metadata:     s.metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    s.tradableAmount,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: &types.OriginTx{
			Id:       s.originTx.Id,
			Source:   s.originTx.Source,
			Contract: a,
		},
	})
}

func (s *bridgeReceiveSuite) AliceAttemptsToBridgeCreditsWithClassIdAndProjectReferenceId(a, b string) {
	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.alice.String(),
		ClassId: a,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  b,
			Jurisdiction: s.jurisdiction,
			Metadata:     s.metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    s.tradableAmount,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: s.originTx,
	})
}

func (s *bridgeReceiveSuite) AliceAttemptsToBridgeCreditsWithProjectProperties(a gocuke.DocString) {
	var project types.MsgBridgeReceive_Project
	err := jsonpb.UnmarshalString(a.Content, &project)
	require.NoError(s.t, err)

	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.alice.String(),
		ClassId: s.classID,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  project.ReferenceId,
			Jurisdiction: project.Jurisdiction,
			Metadata:     project.Metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    s.tradableAmount,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: s.originTx,
	})

	require.NoError(s.t, s.err)
}

func (s *bridgeReceiveSuite) AliceAttemptsToBridgeCreditsWithBatchProperties(a gocuke.DocString) {
	var batch types.MsgBridgeReceive_Batch
	err := jsonpb.UnmarshalString(a.Content, &batch)
	require.NoError(s.t, err)

	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.alice.String(),
		ClassId: s.classID,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceID,
			Jurisdiction: s.jurisdiction,
			Metadata:     s.metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    s.tradableAmount,
			StartDate: batch.StartDate,
			EndDate:   batch.EndDate,
			Metadata:  batch.Metadata,
		},
		OriginTx: s.originTx,
	})

	require.NoError(s.t, s.err)
}

func (s *bridgeReceiveSuite) AliceAttemptsToBridgeCreditsToBobWithTradableAmount(a string) {
	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.alice.String(),
		ClassId: s.classID,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceID,
			Jurisdiction: s.jurisdiction,
			Metadata:     s.metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    a,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: s.originTx,
	})

	require.NoError(s.t, s.err)
}

func (s *bridgeReceiveSuite) AliceAttemptsToBridgeCreditsWithOrigintxSource(a string) {
	originTx := s.originTx
	originTx.Source = a
	s.res, s.err = s.k.BridgeReceive(s.ctx, &types.MsgBridgeReceive{
		Issuer:  s.alice.String(),
		ClassId: s.classID,
		Project: &types.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceID,
			Jurisdiction: s.jurisdiction,
			Metadata:     s.metadata,
		},
		Batch: &types.MsgBridgeReceive_Batch{
			Recipient: s.bob.String(),
			Amount:    s.tradableAmount,
			StartDate: s.startDate,
			EndDate:   s.endDate,
			Metadata:  s.metadata,
		},
		OriginTx: originTx,
	})
}

func (s *bridgeReceiveSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *bridgeReceiveSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *bridgeReceiveSuite) ExpectErrorContains(a string) {
	require.ErrorContains(s.t, s.err, a)
}

func (s *bridgeReceiveSuite) ExpectTotalCreditBatches(a string) {
	expTotal, err := strconv.ParseUint(a, 10, 64)
	require.NoError(s.t, err)

	it, err := s.k.stateStore.BatchTable().List(s.ctx, api.BatchPrimaryKey{})
	require.NoError(s.t, err)

	var total uint64
	for it.Next() {
		total++
	}
	it.Close()

	require.Equal(s.t, expTotal, total)
}

func (s *bridgeReceiveSuite) ExpectTotalProjects(a string) {
	expTotal, err := strconv.ParseUint(a, 10, 64)
	require.NoError(s.t, err)

	it, err := s.k.stateStore.ProjectTable().List(s.ctx, api.ProjectPrimaryKey{})
	require.NoError(s.t, err)

	var total uint64
	for it.Next() {
		total++
	}
	it.Close()

	require.Equal(s.t, expTotal, total)
}

func (s *bridgeReceiveSuite) ExpectProjectProperties(a gocuke.DocString) {
	var expected types.Project
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.ProjectTable().GetById(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.ReferenceId, batch.ReferenceId)
	require.Equal(s.t, expected.Metadata, batch.Metadata)
	require.Equal(s.t, expected.Jurisdiction, batch.Jurisdiction)
}

func (s *bridgeReceiveSuite) ExpectBatchProperties(a gocuke.DocString) {
	var expected types.Batch
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, expected.Denom)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Metadata, batch.Metadata)
	require.Equal(s.t, expected.StartDate.Seconds, batch.StartDate.Seconds)
	require.Equal(s.t, expected.EndDate.Seconds, batch.EndDate.Seconds)
}

func (s *bridgeReceiveSuite) ExpectBobBatchBalance(a gocuke.DocString) {
	var expected api.BatchBalance
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, s.res.BatchDenom)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchBalanceTable().Get(s.ctx, s.bob, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.EscrowedAmount, balance.EscrowedAmount)
}

func (s *bridgeReceiveSuite) ExpectBatchSupply(a gocuke.DocString) {
	var expected api.BatchSupply
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	batch, err := s.stateStore.BatchTable().GetByDenom(s.ctx, s.res.BatchDenom)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BatchSupplyTable().Get(s.ctx, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.CancelledAmount, balance.CancelledAmount)
}

func (s *bridgeReceiveSuite) getProjectSequence(projectID string) uint64 {
	str := strings.Split(projectID, "-")
	seq, err := strconv.ParseUint(str[1], 10, 32)
	require.NoError(s.t, err)
	return seq
}

func (s *bridgeReceiveSuite) getBatchSequence(batchDenom string) uint64 {
	str := strings.Split(batchDenom, "-")
	seq, err := strconv.ParseUint(str[4], 10, 32)
	require.NoError(s.t, err)
	return seq
}
