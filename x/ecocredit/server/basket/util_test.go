package basket_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"
)

func TestGetBasketBalances(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	batchDenom1, classId1 := "C01-001-0000000-0000000-001", "C01"
	batchDenom2, classId2 := "C02-001-0000000-0000000-002", "C02"
	userStartingBalance, amtToDeposit := math.NewDecFromInt64(10), math.NewDecFromInt64(3)
	insertClass(t, s, "C01", "C")
	insertClass(t, s, "C02", "C")

	insertBasket(t, s, "foo", "basket", "C", &api.DateCriteria{YearsInThePast: 3}, []string{classId1})
	insertBasket(t, s, "bar", "basket1", "C", &api.DateCriteria{YearsInThePast: 3}, []string{classId1, classId2})
	initBatch(t, s, 1, batchDenom1, timestamppb.Now())
	initBatch(t, s, 2, batchDenom2, timestamppb.Now())

	insertBatchBalance(t, s, s.addrs[0], 1, userStartingBalance.String())
	insertBatchBalance(t, s, s.addrs[0], 2, userStartingBalance.String())
	insertBatchBalance(t, s, sdk.AccAddress("abcde"), 2, userStartingBalance.String())

	s.bankKeeper.EXPECT().MintCoins(gmAny, gmAny, gmAny).Return(nil).Times(3)
	s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(gmAny, gmAny, gmAny, gmAny).Return(nil).Times(3)

	_, err := s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addrs[0].String(),
		BasketDenom: "foo",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom1, Amount: amtToDeposit.String()},
		},
	})
	assert.NilError(t, err)

	bIdToBalance, err := s.k.GetBasketBalanceMap(s.ctx)
	require.NoError(t, err)
	require.Len(t, bIdToBalance, 1)
	require.Equal(t, bIdToBalance[1], amtToDeposit)

	_, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addrs[0].String(),
		BasketDenom: "bar",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom1, Amount: amtToDeposit.String()},
		},
	})
	assert.NilError(t, err)

	_, err = s.k.Put(s.ctx, &basket.MsgPut{
		Owner:       s.addrs[0].String(),
		BasketDenom: "bar",
		Credits: []*basket.BasketCredit{
			{BatchDenom: batchDenom2, Amount: amtToDeposit.String()},
		},
	})
	assert.NilError(t, err)

	bIdToBalance, err = s.k.GetBasketBalanceMap(s.ctx)
	require.NoError(t, err)
	require.Len(t, bIdToBalance, 2)

	expBatch1Amount, err := amtToDeposit.Add(amtToDeposit)
	require.NoError(t, err)

	require.Equal(t, bIdToBalance[1], expBatch1Amount)
	require.Equal(t, bIdToBalance[2], amtToDeposit)
}

func initBatch(t *testing.T, s *baseSuite, pid uint64, denom string, startDate *timestamppb.Timestamp) {
	assert.NilError(t, s.coreStore.BatchTable().Insert(s.ctx, &ecoApi.Batch{
		ProjectKey: pid,
		Denom:      denom,
		StartDate:  startDate,
		EndDate:    nil,
	}))
}

func insertBatchBalance(t *testing.T, s *baseSuite, user sdk.AccAddress, batchKey uint64, amount string) {
	assert.NilError(t, s.coreStore.BatchBalanceTable().Insert(s.ctx, &ecoApi.BatchBalance{
		BatchKey:       batchKey,
		Address:        user,
		TradableAmount: amount,
		RetiredAmount:  "",
		EscrowedAmount: "",
	}))
}

func insertClass(t *testing.T, s *baseSuite, name, creditTypeAbb string) {
	assert.NilError(t, s.coreStore.ClassTable().Insert(s.ctx, &ecoApi.Class{
		Id:               name,
		Admin:            s.addrs[0],
		Metadata:         "",
		CreditTypeAbbrev: creditTypeAbb,
	}))
}

func insertBasket(t *testing.T, s *baseSuite, denom, name, ctAbbrev string, criteria *api.DateCriteria, classes []string) {
	id, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:       denom,
		Name:              name,
		DisableAutoRetire: false,
		CreditTypeAbbrev:  ctAbbrev,
		DateCriteria:      criteria,
		Exponent:          6,
		Curator:           s.addrs[0],
	})
	assert.NilError(t, err)

	for _, class := range classes {
		assert.NilError(t, s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
			BasketId: id,
			ClassId:  class,
		}))
	}
}
