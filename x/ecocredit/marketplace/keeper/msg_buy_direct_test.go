//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/testutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

type buyDirectSuite struct {
	*baseSuite
	alice             sdk.AccAddress
	aliceBankBalance  sdk.Coin
	bob               sdk.AccAddress
	bobBankBalance    sdk.Coin
	creditTypeAbbrev  string
	classID           string
	batchDenom        string
	sellOrderID       uint64
	disableAutoRetire bool
	quantity          string
	askPrice          sdk.Coin
	bidPrice          sdk.Coin
	res               *types.MsgBuyDirectResponse
	err               error
}

func TestBuyDirect(t *testing.T) {
	gocuke.NewRunner(t, &buyDirectSuite{}).Path("./features/msg_buy_direct.feature").Run()
}

func (s *buyDirectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t, 2)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
	s.aliceBankBalance = sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}
	s.bobBankBalance = sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}
	s.creditTypeAbbrev = "C"
	s.classID = testClassID
	s.batchDenom = testBatchDenom
	s.sellOrderID = 1
	s.quantity = "10"
	s.askPrice = sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(10),
	}
	s.bidPrice = sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(10),
	}
}

func (s *buyDirectSuite) ACreditType() {
	err := s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) ACreditTypeWithPrecision(a string) {
	precision, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	err = s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) AliceHasBankBalance(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.aliceBankBalance = coin
}

func (s *buyDirectSuite) BobHasTheBankBalance(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.bobBankBalance = coin
}

func (s *buyDirectSuite) BobHasABankBalanceWithDenom(a string) {
	s.bobBankBalance = sdk.NewCoin(a, s.bobBankBalance.Amount)
}

func (s *buyDirectSuite) BobHasABankBalanceWithAmount(a string) {
	amount, ok := sdk.NewIntFromString(a)
	require.True(s.t, ok)

	s.bobBankBalance = sdk.NewCoin(s.bidPrice.Denom, amount)
}

func (s *buyDirectSuite) AliceHasTheBatchBalance(a gocuke.DocString) {
	balance := &baseapi.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	batch, err := s.baseStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance.BatchKey = batch.Key
	balance.Address = s.alice

	// Save because the balance already exists from createSellOrders
	err = s.baseStore.BatchBalanceTable().Save(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) BobHasTheBatchBalance(a gocuke.DocString) {
	balance := &baseapi.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	batch, err := s.baseStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance.BatchKey = batch.Key
	balance.Address = s.bob

	err = s.baseStore.BatchBalanceTable().Insert(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) TheBatchSupply(a gocuke.DocString) {
	balance := &baseapi.BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	batch, err := s.baseStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance.BatchKey = batch.Key

	// Save because the supply already exists from createSellOrders
	err = s.baseStore.BatchSupplyTable().Save(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.sellOrderID = id

	s.createSellOrders(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithQuantity(a string) {
	s.quantity = a

	s.createSellOrders(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithAskDenom(a string) {
	s.askPrice = sdk.NewCoin(a, s.askPrice.Amount)

	s.createSellOrders(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithAskAmount(a string) {
	askAmount, ok := sdk.NewIntFromString(a)
	require.True(s.t, ok)

	s.askPrice = sdk.NewCoin(s.askPrice.Denom, askAmount)

	s.createSellOrders(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithDisableAutoRetire(a string) {
	disableAutoRetire, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	s.disableAutoRetire = disableAutoRetire

	s.createSellOrders(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithQuantityAndAskAmount(a string, b string) {
	askAmount, ok := sdk.NewIntFromString(b)
	require.True(s.t, ok)

	s.quantity = a
	s.askPrice = sdk.NewCoin(s.askPrice.Denom, askAmount)

	s.createSellOrders(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithQuantityAndAskPrice(a string, b string) {
	askPrice, err := sdk.ParseCoinNormalized(b)
	require.NoError(s.t, err)

	s.quantity = a
	s.askPrice = askPrice

	s.createSellOrders(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithQuantityAndDisableAutoRetire(a string, b string) {
	disableAutoRetire, err := strconv.ParseBool(b)
	require.NoError(s.t, err)

	s.quantity = a
	s.disableAutoRetire = disableAutoRetire

	s.createSellOrders(1)
}

func (s *buyDirectSuite) AliceCreatedTwoSellOrdersEachWithQuantityAndAskAmount(a string, b string) {
	askAmount, ok := sdk.NewIntFromString(b)
	require.True(s.t, ok)

	s.quantity = a
	s.askPrice = sdk.NewCoin(s.askPrice.Denom, askAmount)

	s.createSellOrders(2)
}

func (s *buyDirectSuite) AliceAttemptsToBuyCreditsWithSellOrderId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.alice.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId: id,
				Quantity:    s.quantity,
				BidPrice:    &s.bidPrice,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithSellOrderId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId: id,
				Quantity:    s.quantity,
				BidPrice:    &s.bidPrice,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithBidDenom(a string) {
	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId: s.sellOrderID,
				Quantity:    s.quantity,
				BidPrice: &sdk.Coin{
					Denom:  a,
					Amount: s.bidPrice.Amount,
				},
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithDisableAutoRetire(a string) {
	disableAutoRetire, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId:       s.sellOrderID,
				Quantity:          s.quantity,
				BidPrice:          &s.bidPrice,
				DisableAutoRetire: disableAutoRetire,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithQuantity(a string) {
	s.quantity = a

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId: s.sellOrderID,
				Quantity:    a,
				BidPrice:    &s.bidPrice,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithQuantityAndBidAmount(a string, b string) {
	bidAmount, ok := sdk.NewIntFromString(b)
	require.True(s.t, ok)

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId: s.sellOrderID,
				Quantity:    a,
				BidPrice: &sdk.Coin{
					Denom:  s.bidPrice.Denom,
					Amount: bidAmount,
				},
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithQuantityAndBidPrice(a string, b string) {
	bidPrice, err := sdk.ParseCoinNormalized(b)
	require.NoError(s.t, err)

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId: s.sellOrderID,
				Quantity:    a,
				BidPrice:    &bidPrice,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithQuantityAndDisableAutoRetire(a string, b string) {
	disableAutoRetire, err := strconv.ParseBool(b)
	require.NoError(s.t, err)

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId:       s.sellOrderID,
				Quantity:          a,
				BidPrice:          &s.bidPrice,
				DisableAutoRetire: disableAutoRetire,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsInTwoOrdersEachWithQuantityAndBidAmount(a string, b string) {
	s.quantity = a

	bidAmount, ok := sdk.NewIntFromString(b)
	require.True(s.t, ok)

	s.multipleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId: 1,
				Quantity:    a,
				BidPrice: &sdk.Coin{
					Denom:  s.bidPrice.Denom,
					Amount: bidAmount,
				},
			},
			{
				SellOrderId: 2,
				Quantity:    a,
				BidPrice: &sdk.Coin{
					Denom:  s.bidPrice.Denom,
					Amount: bidAmount,
				},
			},
		},
	})
}

func (s *buyDirectSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *buyDirectSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *buyDirectSuite) ExpectSellOrderWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	order, err := s.marketStore.SellOrderTable().Get(s.ctx, id)
	require.NoError(s.t, err)

	require.Equal(s.t, order.Id, id)
}

func (s *buyDirectSuite) ExpectSellOrderWithQuantity(a string) {
	order, err := s.marketStore.SellOrderTable().Get(s.ctx, s.sellOrderID)
	require.NoError(s.t, err)

	require.Equal(s.t, order.Quantity, a)
}

func (s *buyDirectSuite) ExpectNoSellOrderWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	_, err = s.marketStore.SellOrderTable().Get(s.ctx, id)
	require.EqualError(s.t, err, ormerrors.NotFound.Error())
}

func (s *buyDirectSuite) ExpectAliceBankBalance(a string) {
	bankBalance, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	require.Equal(s.t, bankBalance.Denom, s.aliceBankBalance.Denom)
	require.Equal(s.t, bankBalance.Amount.String(), s.aliceBankBalance.Amount.String())
}

func (s *buyDirectSuite) ExpectBobBankBalance(a string) {
	bankBalance, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	require.Equal(s.t, bankBalance.Denom, s.bobBankBalance.Denom)
	require.Equal(s.t, bankBalance.Amount.String(), s.bobBankBalance.Amount.String())
}

func (s *buyDirectSuite) ExpectAliceBatchBalance(a gocuke.DocString) {
	expected := &baseapi.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	batch, err := s.baseStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.baseStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.EscrowedAmount, balance.EscrowedAmount)
}

func (s *buyDirectSuite) ExpectBobBatchBalance(a gocuke.DocString) {
	expected := &baseapi.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	batch, err := s.baseStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.baseStore.BatchBalanceTable().Get(s.ctx, s.bob, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
	require.Equal(s.t, expected.EscrowedAmount, balance.EscrowedAmount)
}

func (s *buyDirectSuite) ExpectBatchSupply(a gocuke.DocString) {
	expected := &baseapi.BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	batch, err := s.baseStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.baseStore.BatchSupplyTable().Get(s.ctx, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
}

func (s *buyDirectSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventBuyDirect
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

// count is the number of sell orders created
func (s *buyDirectSuite) createSellOrders(count int) {
	totalQuantity := s.quantity

	if count > 1 {
		c := math.NewDecFromInt64(int64(count))
		q, err := math.NewDecFromString(s.quantity)
		require.NoError(s.t, err)
		t, err := c.Mul(q)
		require.NoError(s.t, err)
		totalQuantity = t.String()
	}

	err := s.baseStore.ClassTable().Insert(s.ctx, &baseapi.Class{
		Id:               s.classID,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	batchKey, err := s.baseStore.BatchTable().InsertReturningID(s.ctx, &baseapi.Batch{
		Denom: s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.baseStore.BatchBalanceTable().Insert(s.ctx, &baseapi.BatchBalance{
		BatchKey:       batchKey,
		Address:        s.alice,
		EscrowedAmount: totalQuantity,
	})
	require.NoError(s.t, err)

	err = s.baseStore.BatchSupplyTable().Insert(s.ctx, &baseapi.BatchSupply{
		BatchKey:       batchKey,
		TradableAmount: totalQuantity,
	})
	require.NoError(s.t, err)

	marketKey, err := s.marketStore.MarketTable().InsertReturningID(s.ctx, &api.Market{
		CreditTypeAbbrev: s.creditTypeAbbrev,
		BankDenom:        s.askPrice.Denom,
	})
	require.NoError(s.t, err)

	order := &api.SellOrder{
		Seller:            s.alice,
		BatchKey:          batchKey,
		Quantity:          s.quantity,
		MarketId:          marketKey,
		AskAmount:         s.askPrice.Amount.String(),
		DisableAutoRetire: s.disableAutoRetire,
	}

	sellOrderID, err := s.marketStore.SellOrderTable().InsertReturningID(s.ctx, order)
	require.NoError(s.t, err)
	require.Equal(s.t, sellOrderID, s.sellOrderID)

	for i := 1; i < count; i++ {
		order.Id = 0 // reset sell order id
		err = s.marketStore.SellOrderTable().Insert(s.ctx, order)
		require.NoError(s.t, err)
	}
}

func (s *buyDirectSuite) singleBuyOrderExpectCalls() {
	askTotal := s.calculateAskTotal(s.quantity, s.askPrice.Amount.String())
	sendCoin := sdk.NewCoin(s.askPrice.Denom, askTotal)
	sendCoins := sdk.NewCoins(sendCoin)

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.bob, s.bidPrice.Denom).
		Return(s.bobBankBalance).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoins(s.sdkCtx, s.bob, s.alice, sendCoins).
		Do(func(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) {
			s.bobBankBalance = s.bobBankBalance.Sub(sendCoin)
			s.aliceBankBalance = s.aliceBankBalance.Add(sendCoin)
		}).
		AnyTimes() // not expected on failed attempt
}

func (s *buyDirectSuite) multipleBuyOrderExpectCalls() {
	askTotal := s.calculateAskTotal(s.quantity, s.askPrice.Amount.String())
	sendCoin := sdk.NewCoin(s.askPrice.Denom, askTotal)
	sendCoins := sdk.NewCoins(sendCoin)

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.bob, s.bidPrice.Denom).
		Return(s.bobBankBalance).
		Times(1)

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.bob, s.bidPrice.Denom).
		Return(s.bobBankBalance.Sub(sendCoin)).
		Times(1)

	s.bankKeeper.EXPECT().
		SendCoins(s.sdkCtx, s.bob, s.alice, sendCoins).
		Do(func(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) {
			s.bobBankBalance = s.bobBankBalance.Sub(sendCoin)
			s.aliceBankBalance = s.aliceBankBalance.Add(sendCoin)
		}).
		AnyTimes() // not expected on failed attempt
}

func (s *buyDirectSuite) calculateAskTotal(quantity string, amount string) sdkmath.Int {
	q, err := math.NewDecFromString(quantity)
	require.NoError(s.t, err)

	a, err := math.NewDecFromString(amount)
	require.NoError(s.t, err)

	t, err := a.Mul(q)
	require.NoError(s.t, err)

	return t.SdkIntTrim()
}
