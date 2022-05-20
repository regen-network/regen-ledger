package marketplace

import (
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type sellSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	creditTypeAbbrev string
	batchDenom       string
	askDenom         string
	askAmount        string
	creditQuantity   string
	res              *marketplace.MsgSellResponse
	err              error
}

func TestSell(t *testing.T) {
	gocuke.NewRunner(t, &sellSuite{}).Path("./features/msg_sell.feature").Run()
}

func (s *sellSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.creditTypeAbbrev = "C"
	s.batchDenom = "C01-001-20200101-20210101-001"
	s.askDenom = "regen"
	s.askAmount = "100"
	s.creditQuantity = "100"
}

func (s *sellSuite) AnAllowedDenom() {
	err := s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom: s.askDenom,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AnAllowedDenomWithDenom(a string) {
	err := s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom: a,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) ACreditBatchWithBatchDenom(a string) {
	classId := core.GetClassIdFromBatchDenom(a)
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(classId)

	classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               classId,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchTable().Insert(s.ctx, &coreapi.Batch{
		ProjectKey: projectKey,
		Denom:      a,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AliceOwnsCredits() {
	classId := core.GetClassIdFromBatchDenom(s.batchDenom)
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(classId)

	classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               classId,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		ProjectKey: projectKey,
		Denom:      s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Tradable: s.creditQuantity,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AliceOwnsCreditsWithBatchDenom(a string) {
	classId := core.GetClassIdFromBatchDenom(a)
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(classId)

	classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               classId,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		ProjectKey: projectKey,
		Denom:      a,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Tradable: s.creditQuantity,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AliceOwnsCreditQuantity(a string) {
	classId := core.GetClassIdFromBatchDenom(s.batchDenom)
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(classId)

	classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
		Id:               classId,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.coreStore.ProjectTable().InsertReturningID(s.ctx, &coreapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		ProjectKey: projectKey,
		Denom:      s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Tradable: a,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithBatchDenom(a string) {
	amount, ok := sdk.NewIntFromString(s.askAmount)
	require.True(s.t, ok)

	askPrice := sdk.NewCoin(s.askDenom, amount)

	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom: a,
				Quantity:   s.creditQuantity,
				AskPrice:   &askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithCreditQuantity(a string) {
	amount, ok := sdk.NewIntFromString(s.askAmount)
	require.True(s.t, ok)

	askPrice := sdk.NewCoin(s.askDenom, amount)

	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   a,
				AskPrice:   &askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithAskPrice(a string) {
	askPrice, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   s.creditQuantity,
				AskPrice:   &askPrice,
			},
		},
	})
}

func (s *sellSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *sellSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}
