package marketplace

import (
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type buyDirectSuite struct {
	*baseSuite
	alice             sdk.AccAddress
	aliceBankBalance  sdk.Coin
	bob               sdk.AccAddress
	bobBankBalance    sdk.Coin
	creditTypeAbbrev  string
	classId           string
	batchDenom        string
	sellOrderId       uint64
	disableAutoRetire bool
	quantity          string
	askPrice          sdk.Coin
	bidPrice          sdk.Coin
	res               *marketplace.MsgBuyDirectResponse
	err               error
}

func TestBuyDirect(t *testing.T) {
	gocuke.NewRunner(t, &buyDirectSuite{}).Path("./features/msg_buy_direct.feature").Run()
}

func (s *buyDirectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.bob = s.addr2
	s.aliceBankBalance = sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}
	s.bobBankBalance = sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}
	s.creditTypeAbbrev = "C"
	s.classId = "C01"
	s.batchDenom = "C01-001-20200101-20210101-001"
	s.sellOrderId = 1
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
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) ACreditTypeWithPrecision(a string) {
	precision, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	err = s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
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
	balance := &coreapi.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance.BatchKey = batch.Key
	balance.Address = s.alice

	// Update because the balance already exists from sellOrderSetup
	err = s.coreStore.BatchBalanceTable().Update(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) BobHasTheBatchBalance(a gocuke.DocString) {
	balance := &coreapi.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance.BatchKey = batch.Key
	balance.Address = s.bob

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) TheBatchSupply(a gocuke.DocString) {
	balance := &coreapi.BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, balance)
	require.NoError(s.t, err)

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance.BatchKey = batch.Key

	// Update because the supply already exists from sellOrderSetup
	err = s.coreStore.BatchSupplyTable().Update(s.ctx, balance)
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.sellOrderId = id // required for sell order setup

	s.sellOrderSetup(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithQuantity(a string) {
	s.quantity = a // required for sell order setup

	s.sellOrderSetup(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithAskDenom(a string) {
	s.askPrice = sdk.NewCoin(a, s.askPrice.Amount) // required for sell order setup

	s.sellOrderSetup(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithAskAmount(a string) {
	askAmount, ok := sdk.NewIntFromString(a)
	require.True(s.t, ok)

	s.askPrice = sdk.NewCoin(s.askPrice.Denom, askAmount) // required for sell order setup

	s.sellOrderSetup(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithDisableAutoRetire(a string) {
	disableAutoRetire, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	s.disableAutoRetire = disableAutoRetire // required for sell order setup

	s.sellOrderSetup(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithQuantityAndAskAmount(a string, b string) {
	askAmount, ok := sdk.NewIntFromString(b)
	require.True(s.t, ok)

	s.quantity = a                                        // required for sell order setup
	s.askPrice = sdk.NewCoin(s.askPrice.Denom, askAmount) // required for sell order setup

	s.sellOrderSetup(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithQuantityAndAskPrice(a string, b string) {
	askPrice, err := sdk.ParseCoinNormalized(b)
	require.NoError(s.t, err)

	s.quantity = a        // required for sell order setup
	s.askPrice = askPrice // required for sell order setup

	s.sellOrderSetup(1)
}

func (s *buyDirectSuite) AliceCreatedASellOrderWithQuantityAndDisableAutoRetire(a string, b string) {
	disableAutoRetire, err := strconv.ParseBool(b)
	require.NoError(s.t, err)

	s.quantity = a                          // required for sell order setup
	s.disableAutoRetire = disableAutoRetire // required for sell order setup

	s.sellOrderSetup(1)
}

func (s *buyDirectSuite) AliceCreatedTwoSellOrdersEachWithQuantityAndAskAmount(a string, b string) {
	askAmount, ok := sdk.NewIntFromString(b)
	require.True(s.t, ok)

	s.quantity = a                                        // required for sell order setup
	s.askPrice = sdk.NewCoin(s.askPrice.Denom, askAmount) // required for sell order setup

	s.sellOrderSetup(2)
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithSellOrderId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
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

	s.res, s.err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
			{
				SellOrderId: s.sellOrderId,
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

	s.res, s.err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
			{
				SellOrderId:       s.sellOrderId,
				Quantity:          s.quantity,
				BidPrice:          &s.bidPrice,
				DisableAutoRetire: disableAutoRetire,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithQuantity(a string) {
	s.quantity = a // required for buy order expect calls

	s.singleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
			{
				SellOrderId: s.sellOrderId,
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

	s.res, s.err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
			{
				SellOrderId: s.sellOrderId,
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

	s.res, s.err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
			{
				SellOrderId: s.sellOrderId,
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

	s.res, s.err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
			{
				SellOrderId:       s.sellOrderId,
				Quantity:          a,
				BidPrice:          &s.bidPrice,
				DisableAutoRetire: disableAutoRetire,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsInTwoOrdersEachWithQuantityAndBidAmount(a string, b string) {
	s.quantity = a // required for buy order expect calls

	bidAmount, ok := sdk.NewIntFromString(b)
	require.True(s.t, ok)

	s.multipleBuyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &marketplace.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*marketplace.MsgBuyDirect_Order{
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
	order, err := s.marketStore.SellOrderTable().Get(s.ctx, s.sellOrderId)
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
	expected := &coreapi.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Retired, balance.Retired)
	require.Equal(s.t, expected.Tradable, balance.Tradable)
	require.Equal(s.t, expected.Escrowed, balance.Escrowed)
}

func (s *buyDirectSuite) ExpectBobBatchBalance(a gocuke.DocString) {
	expected := &coreapi.BatchBalance{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.bob, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Retired, balance.Retired)
	require.Equal(s.t, expected.Tradable, balance.Tradable)
	require.Equal(s.t, expected.Escrowed, balance.Escrowed)
}

func (s *buyDirectSuite) ExpectBatchSupply(a gocuke.DocString) {
	expected := &coreapi.BatchSupply{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchSupplyTable().Get(s.ctx, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.RetiredAmount, balance.RetiredAmount)
	require.Equal(s.t, expected.TradableAmount, balance.TradableAmount)
}

func (s *buyDirectSuite) sellOrderSetup(count int) {
	err := s.coreStore.ClassTable().Insert(s.ctx, &coreapi.Class{
		Id:               s.classId,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		Denom: s.batchDenom,
	})
	require.NoError(s.t, err)

	quantity := s.quantity

	if count == 2 {
		c, err := math.NewDecFromString("2")
		require.NoError(s.t, err)
		q, err := math.NewDecFromString(s.quantity)
		require.NoError(s.t, err)
		t, err := c.Mul(q)
		require.NoError(s.t, err)
		quantity = t.String()
	}

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey: batchKey,
		Address:  s.alice,
		Escrowed: quantity,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchSupplyTable().Insert(s.ctx, &coreapi.BatchSupply{
		BatchKey:       batchKey,
		TradableAmount: quantity, // TODO: tradable #1123
	})
	require.NoError(s.t, err)

	marketKey, err := s.marketStore.MarketTable().InsertReturningID(s.ctx, &api.Market{
		CreditType: s.creditTypeAbbrev, // TODO: credit_type_abbrev #1123
		BankDenom:  s.askPrice.Denom,
	})
	require.NoError(s.t, err)

	sellOrderId, err := s.marketStore.SellOrderTable().InsertReturningID(s.ctx, &api.SellOrder{
		Seller:            s.alice,
		BatchId:           batchKey, // TODO: batch_key #1123
		Quantity:          quantity,
		MarketId:          marketKey,                  // TODO: market_key #1124
		AskPrice:          s.askPrice.Amount.String(), // TODO: ask_amount #1123
		DisableAutoRetire: s.disableAutoRetire,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, sellOrderId, s.sellOrderId)

	if count == 2 {
		err = s.marketStore.SellOrderTable().Insert(s.ctx, &api.SellOrder{
			Seller:            s.alice,
			BatchId:           batchKey,
			Quantity:          quantity,
			MarketId:          marketKey,
			AskPrice:          s.askPrice.Amount.String(),
			DisableAutoRetire: s.disableAutoRetire,
		})
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

func (s *buyDirectSuite) calculateAskTotal(quantity string, amount string) sdk.Int {
	q, err := math.NewDecFromString(quantity)
	require.NoError(s.t, err)

	a, err := math.NewDecFromString(amount)
	require.NoError(s.t, err)

	t, err := a.Mul(q)
	require.NoError(s.t, err)

	return t.SdkIntTrim()
}
