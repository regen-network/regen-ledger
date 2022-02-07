package fill_test

import (
	"bytes"
	"context"
	"testing"

	mathtestutil "github.com/regen-network/regen-ledger/types/testutil/math"

	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/regen-network/regen-ledger/x/ecocredit/fill"

	"github.com/golang/mock/gomock"

	"github.com/regen-network/regen-ledger/x/ecocredit/fill/mocks"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	marketplacev1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1beta1"
	"gotest.tools/v3/assert"

	"github.com/regen-network/regen-ledger/x/ecocredit/testutil"

	"github.com/rs/zerolog"
)

type suite struct {
	ctrl             *gomock.Controller
	transferMgr      *mocks.MockTransferManager
	db               ormdb.ModuleDB
	market           *marketplacev1beta1.Market
	marketplaceStore marketplacev1beta1.StateStore
	acct1            sdk.AccAddress
	acct2            sdk.AccAddress
	loggerBuf        *bytes.Buffer
	logger           zerolog.Logger
	ctx              context.Context
	fillMgr          fill.Manager
	marketId         uint64
}

func setup(t *testing.T) *suite {
	s := &suite{}
	s.ctrl = gomock.NewController(t)
	s.transferMgr = mocks.NewMockTransferManager(s.ctrl)
	var err error
	s.db, err = ormdb.NewModuleDB(testutil.TestModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)

	s.loggerBuf = &bytes.Buffer{}
	s.logger = zerolog.New(s.loggerBuf)

	s.fillMgr, err = fill.NewManager(s.db, s.transferMgr, s.logger)
	assert.NilError(t, err)
	s.ctx = ormtable.WrapContextDefault(ormtest.NewMemoryBackend())

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
		BidPrice:           "10",
		DisableAutoRetire:  false,
		DisablePartialFill: false,
		Expiration:         nil,
		Maker:              false,
	}

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

	s.transferMgr.EXPECT().SendCreditsTo(uint64(1), mathtestutil.MatchDecFromInt64(10), s.acct2, s.acct1, true)
	s.transferMgr.EXPECT().SendCoinsTo("foo", mathtestutil.MatchInt(100), s.acct1, s.acct2)

	state, err := s.fillMgr.Fill(s.ctx, s.market, buyOrder, sellOrder)
	assert.NilError(t, err)
	assert.Equal(t, fill.BothFilled, state)
}

func TestBuyFilled(t *testing.T) {
	s := setup(t)
	buyOrder := &marketplacev1beta1.BuyOrder{
		Buyer:              s.acct1,
		Quantity:           "5",
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
		Quantity:          "10",
		MarketId:          s.marketId,
		AskPrice:          "10",
		DisableAutoRetire: false,
		Expiration:        nil,
		Maker:             false,
	}
	assert.NilError(t, s.marketplaceStore.SellOrderStore().Insert(s.ctx, sellOrder))

	s.transferMgr.EXPECT().SendCreditsTo(uint64(1), mathtestutil.MatchDecFromInt64(5), s.acct2, s.acct1, true)
	s.transferMgr.EXPECT().SendCoinsTo("foo", mathtestutil.MatchInt(50), s.acct1, s.acct2)

	state, err := s.fillMgr.Fill(s.ctx, s.market, buyOrder, sellOrder)
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

	s.transferMgr.EXPECT().SendCreditsTo(uint64(1), mathtestutil.MatchDecFromInt64(5), s.acct2, s.acct1, true)
	s.transferMgr.EXPECT().SendCoinsTo("foo", mathtestutil.MatchInt(50), s.acct1, s.acct2)

	state, err := s.fillMgr.Fill(s.ctx, s.market, buyOrder, sellOrder)
	assert.NilError(t, err)
	assert.Equal(t, fill.SellFilled, state)
}
