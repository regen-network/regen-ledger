//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"strconv"
	"testing"

<<<<<<< HEAD
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/gogoproto/jsonpb"
||||||| 45044a35
	sdkmath "cosmossdk.io/math"
	"github.com/gogo/protobuf/jsonpb"
=======
	"github.com/cockroachdb/apd/v3"
	"github.com/gogo/protobuf/jsonpb"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
>>>>>>> origin/main
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/v2/math"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

type buyDirectSuite struct {
	*baseSuite
	alice             sdk.AccAddress
	bob               sdk.AccAddress
	balances          map[string]sdk.Coins
	moduleBalances    map[string]sdk.Coins
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
	maxFee            *sdk.Coin
}

func TestBuyDirect(t *testing.T) {
	gocuke.NewRunner(t, &buyDirectSuite{}).Path("./features/msg_buy_direct.feature").Run()
}

func (s *buyDirectSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t, 2)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
	s.balances = map[string]sdk.Coins{}
	s.balances[s.alice.String()] = sdk.Coins{sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}}
	s.balances[s.bob.String()] = sdk.Coins{sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}}
	s.moduleBalances = map[string]sdk.Coins{}
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

	s.buyOrderExpectCalls()
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

	s.balances[s.alice.String()] = sdk.Coins{coin}
}

func (s *buyDirectSuite) BobHasBankBalance(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.balances[s.bob.String()] = sdk.Coins{coin}
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

func (s *buyDirectSuite) AlicesAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.alice = addr
}

func (s *buyDirectSuite) BobsAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.bob = addr
}
func (s *buyDirectSuite) BobSetsAMaxFeeOf(a string) {
	maxFee, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)
	s.maxFee = &maxFee
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

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithSellOrderIdAndRetirementReason(a, b string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.buyOrderExpectCalls()

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId:      id,
				Quantity:         s.quantity,
				BidPrice:         &s.bidPrice,
				RetirementReason: b,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithBidDenom(a string) {
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

	s.res, s.err = s.k.BuyDirect(s.ctx, &types.MsgBuyDirect{
		Buyer: s.bob.String(),
		Orders: []*types.MsgBuyDirect_Order{
			{
				SellOrderId:  s.sellOrderID,
				Quantity:     a,
				BidPrice:     &bidPrice,
				MaxFeeAmount: s.maxFee,
			},
		},
	})
}

func (s *buyDirectSuite) BobAttemptsToBuyCreditsWithQuantityAndDisableAutoRetire(a string, b string) {
	disableAutoRetire, err := strconv.ParseBool(b)
	require.NoError(s.t, err)

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
func (s *buyDirectSuite) BuyerFeesAreAndSellerFeesAre(buyerFee *apd.Decimal, sellerFee *apd.Decimal) {
	err := s.k.stateStore.FeeParamsTable().Save(s.ctx, &api.FeeParams{
		BuyerPercentageFee:  buyerFee.String(),
		SellerPercentageFee: sellerFee.String(),
	})
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *buyDirectSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *buyDirectSuite) ExpectErrorContains(a string) {
	if a == "" {
		require.NoError(s.t, s.err)
	} else {
		require.ErrorContains(s.t, s.err, a)
	}
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
	s.expectBalance(s.alice, a)
}

func (s *buyDirectSuite) ExpectBobBankBalance(a string) {
	s.expectBalance(s.bob, a)
}

func (s *buyDirectSuite) expectBalance(address sdk.Address, a string) {
	expected, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)

	actual := s.balances[address.String()]

	if !actual.IsEqual(expected) {
		s.t.Fatalf("expected: %s, actual: %s", a, actual.String())
	}
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

func (s *buyDirectSuite) ExpectEventBuyDirectWithProperties(a gocuke.DocString) {
	var event types.EventBuyDirect
	err := jsonpb.UnmarshalString(a.Content, &event)
	require.NoError(s.t, err)

	s.expectEvent(&event)
}

func (s *buyDirectSuite) expectEvent(expected gogoproto.Message) {
	sdkEvent, found := testutil.GetEvent(expected, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	foundEvt, err := sdk.ParseTypedEvent(abci.Event(sdkEvent))
	require.NoError(s.t, err)

	msgType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(gogoproto.MessageName(expected)))
	require.NoError(s.t, err)
	evt := msgType.New().Interface()
	evt2 := msgType.New().Interface()

	require.NoError(s.t, gogoToProtoReflect(foundEvt, evt))
	require.NoError(s.t, gogoToProtoReflect(expected, evt2))

	if diff := cmp.Diff(evt, evt2, protocmp.Transform()); diff != "" {
		s.t.Fatalf("unexpected event diff: %s", diff)
	}
}

func (s *buyDirectSuite) ExpectEventTransferWithProperties(a gocuke.DocString) {
	var event basetypes.EventTransfer
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) ExpectEventRetireWithProperties(a gocuke.DocString) {
	var event basetypes.EventRetire
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *buyDirectSuite) ExpectFeePoolBalance(a string) {
	expected, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)

	actual := s.moduleBalances[marketplace.FeePoolName]

	if !actual.IsEqual(expected) {
		s.t.Fatalf("expected: %s, actual: %s", a, actual.String())
	}
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

func (s *buyDirectSuite) buyOrderExpectCalls() {
	s.bankKeeper.EXPECT().
		GetBalance(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
			return sdk.Coin{Denom: denom, Amount: s.balances[addr.String()].AmountOf(denom)}
		}).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoins(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ sdk.Context, from, to sdk.AccAddress, amt sdk.Coins) {
			s.balances[from.String()] = s.balances[from.String()].Sub(amt...)
			s.balances[to.String()] = s.balances[to.String()].Add(amt...)
		}).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ sdk.Context, from sdk.AccAddress, mod string, amt sdk.Coins) {
			s.balances[from.String()] = s.balances[from.String()].Sub(amt...)
			s.moduleBalances[mod] = s.moduleBalances[mod].Add(amt...)
		}).
		AnyTimes()

	s.bankKeeper.EXPECT().
		SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ sdk.Context, mod string, to sdk.AccAddress, amt sdk.Coins) {
			s.moduleBalances[mod] = s.moduleBalances[mod].Sub(amt...)
			s.balances[to.String()] = s.balances[to.String()].Add(amt...)
		}).
		AnyTimes()

	s.bankKeeper.EXPECT().
		BurnCoins(gomock.Any(), gomock.Any(), gomock.Any()).
		Do(func(_ sdk.Context, mod string, amt sdk.Coins) {
			s.moduleBalances[mod] = s.moduleBalances[mod].Sub(amt...)
		}).AnyTimes()
}
