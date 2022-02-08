package fill_test

import (
	"context"
	"testing"

	"github.com/regen-network/regen-ledger/x/ecocredit/server/fill"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/fill/mocks"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/testutil"

	"github.com/cosmos/cosmos-sdk/orm/testing/ormmocks"

	mathtestutil "github.com/regen-network/regen-ledger/types/testutil/math"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/golang/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	"gotest.tools/v3/assert"
)

type suite struct {
	ctrl             *gomock.Controller
	transferMgr      *mocks.MockTransferManager
	db               ormdb.ModuleDB
	market           *marketplacev1beta1.Market
	marketplaceStore marketplacev1beta1.StateStore
	acct1            sdk.AccAddress
	acct2            sdk.AccAddress
	ctx              context.Context
	fillMgr          fill.Manager
	marketId         uint64
	backend          ormtable.Backend
}

func setup(t *testing.T) *suite {
	s := &suite{}
	s.ctrl = gomock.NewController(t)
	s.transferMgr = mocks.NewMockTransferManager(s.ctrl)
	var err error
	s.db, err = ormdb.NewModuleDB(testutil.TestModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)

	s.fillMgr, err = fill.NewManager(s.db, s.transferMgr)
	assert.NilError(t, err)

	s.backend = ormtest.NewMemoryBackend()
	s.ctx = ormtable.WrapContextDefault(s.backend)

	s.marketplaceStore, err = marketplacev1beta1.NewStateStore(s.db)

	s.market = &marketplacev1beta1.Market{
		CreditType:        "C",
		BankDenom:         "foo",
		PrecisionModifier: 3,
	}
	s.marketId, err = s.marketplaceStore.MarketStore().InsertReturningID(s.ctx, s.market)
	assert.NilError(t, err)

	s.acct1 = sdk.AccAddress{0, 1, 2, 3, 4, 5}
	s.acct2 = sdk.AccAddress{5, 4, 3, 2, 1, 0}

	return s
}

func TestBothFilled(t *testing.T) {
	s := setup(t)
	buyOrder := &marketplacev1beta1.BuyOrder{
		Buyer:              s.acct1,
		Quantity:           "10",
		MarketId:           s.marketId,
		BidPrice:           "11",
		DisableAutoRetire:  false,
		DisablePartialFill: false,
		Expiration:         nil,
		Maker:              false,
	}
	assert.NilError(t, s.marketplaceStore.BuyOrderStore().Insert(s.ctx, buyOrder))

	sellOrder := &marketplacev1beta1.SellOrder{
		Seller:            s.acct2,
		BatchId:           1,
		Quantity:          "10",
		MarketId:          s.marketId,
		AskPrice:          "10",
		DisableAutoRetire: false,
		Expiration:        nil,
		Maker:             false,
	}
	assert.NilError(t, s.marketplaceStore.SellOrderStore().Insert(s.ctx, sellOrder))

	ormHooks := ormmocks.NewMockHooks(s.ctrl)
	ctx := ormtable.WrapContextDefault(s.backend.WithHooks(ormHooks))

	s.transferMgr.EXPECT().SendCreditsTo(uint64(1), mathtestutil.MatchDecFromInt64(10), s.acct2, s.acct1, true)
	s.transferMgr.EXPECT().SendCoinsTo("foo", mathtestutil.MatchInt(110), s.acct1, s.acct2)
	ormHooks.EXPECT().OnDelete(ormmocks.Eq(buyOrder))
	ormHooks.EXPECT().OnDelete(ormmocks.Eq(sellOrder))

	state, err := s.fillMgr.Fill(ctx, s.market, buyOrder, sellOrder)
	assert.NilError(t, err)
	assert.Equal(t, fill.BothFilled, state)
}

func TestBuyFilled(t *testing.T) {
	s := setup(t)
	buyOrder := &marketplacev1beta1.BuyOrder{
		Buyer:              s.acct1,
		Quantity:           "5.5",
		MarketId:           s.marketId,
		BidPrice:           "108",
		DisableAutoRetire:  false,
		DisablePartialFill: false,
		Expiration:         nil,
		Maker:              false,
	}
	assert.NilError(t, s.marketplaceStore.BuyOrderStore().Insert(s.ctx, buyOrder))

	sellOrder := &marketplacev1beta1.SellOrder{
		Seller:            s.acct2,
		BatchId:           1,
		Quantity:          "10",
		MarketId:          s.marketId,
		AskPrice:          "105",
		DisableAutoRetire: false,
		Expiration:        nil,
		Maker:             false,
	}
	assert.NilError(t, s.marketplaceStore.SellOrderStore().Insert(s.ctx, sellOrder))

	ormHooks := ormmocks.NewMockHooks(s.ctrl)
	ctx := ormtable.WrapContextDefault(s.backend.WithHooks(ormHooks))

	s.transferMgr.EXPECT().SendCreditsTo(uint64(1), mathtestutil.MatchDecFromString("5.5"), s.acct2, s.acct1, true)
	s.transferMgr.EXPECT().SendCoinsTo("foo", mathtestutil.MatchInt(594), s.acct1, s.acct2)
	ormHooks.EXPECT().OnUpdate(gomock.Any(), ormmocks.Eq(sellOrder))
	ormHooks.EXPECT().OnDelete(ormmocks.Eq(buyOrder))

	state, err := s.fillMgr.Fill(ctx, s.market, buyOrder, sellOrder)
	assert.NilError(t, err)
	assert.Equal(t, fill.BuyFilled, state)
}

func TestSellFilled(t *testing.T) {
	s := setup(t)
	buyOrder := &marketplacev1beta1.BuyOrder{
		Buyer:              s.acct1,
		Quantity:           "10",
		MarketId:           s.marketId,
		BidPrice:           "10",
		DisableAutoRetire:  false,
		DisablePartialFill: false,
		Expiration:         nil,
		Maker:              false,
	}
	assert.NilError(t, s.marketplaceStore.BuyOrderStore().Insert(s.ctx, buyOrder))

	sellOrder := &marketplacev1beta1.SellOrder{
		Seller:            s.acct2,
		BatchId:           1,
		Quantity:          "5",
		MarketId:          s.marketId,
		AskPrice:          "10",
		DisableAutoRetire: false,
		Expiration:        nil,
		Maker:             false,
	}
	assert.NilError(t, s.marketplaceStore.SellOrderStore().Insert(s.ctx, sellOrder))

	ormHooks := ormmocks.NewMockHooks(s.ctrl)
	ctx := ormtable.WrapContextDefault(s.backend.WithHooks(ormHooks))

	s.transferMgr.EXPECT().SendCreditsTo(uint64(1), mathtestutil.MatchDecFromInt64(5), s.acct2, s.acct1, true)
	s.transferMgr.EXPECT().SendCoinsTo("foo", mathtestutil.MatchInt(50), s.acct1, s.acct2)
	ormHooks.EXPECT().OnUpdate(gomock.Any(), ormmocks.Eq(buyOrder))
	ormHooks.EXPECT().OnDelete(ormmocks.Eq(sellOrder))

	state, err := s.fillMgr.Fill(ctx, s.market, buyOrder, sellOrder)
	assert.NilError(t, err)
	assert.Equal(t, fill.SellFilled, state)
}

func TestBadAutoRetireMatch(t *testing.T) {
	s := setup(t)
	buyOrder := &marketplacev1beta1.BuyOrder{
		DisableAutoRetire: true,
	}

	sellOrder := &marketplacev1beta1.SellOrder{
		DisableAutoRetire: false,
	}

	_, err := s.fillMgr.Fill(s.ctx, s.market, buyOrder, sellOrder)
	assert.ErrorContains(t, err, "unexpected")
}
