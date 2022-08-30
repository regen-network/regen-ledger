//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/types/testutil"
	types "github.com/regen-network/regen-ledger/x/ecocredit/marketplace/types/v1"
)

type updateSellOrdersSuite struct {
	*baseSuite
	alice               sdk.AccAddress
	bob                 sdk.AccAddress
	aliceTradableAmount string
	aliceEscrowedAmount string
	creditTypeAbbrev    string
	classID             string
	batchDenom          string
	sellOrderID         uint64
	askPrice            *sdk.Coin
	quantity            string
	disableAutoRetire   bool
	expiration          *time.Time
	res                 *types.MsgUpdateSellOrdersResponse
	err                 error
}

func TestUpdateSellOrders(t *testing.T) {
	gocuke.NewRunner(t, &updateSellOrdersSuite{}).Path("./features/msg_update_sell_orders.feature").Run()
}

func (s *updateSellOrdersSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t, 2)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
	s.aliceTradableAmount = "200"
	s.creditTypeAbbrev = "C"
	s.classID = testClassID
	s.batchDenom = testBatchDenom
	s.sellOrderID = 1
	s.askPrice = &sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}
	s.quantity = "100"
}

func (s *updateSellOrdersSuite) ABlockTimeWithTimestamp(a string) {
	blockTime, err := regentypes.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
}

func (s *updateSellOrdersSuite) ACreditType() {
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *updateSellOrdersSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: a,
	})
	require.NoError(s.t, err)
}

func (s *updateSellOrdersSuite) ACreditTypeWithPrecision(a string) {
	precision, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	err = s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)
}

func (s *updateSellOrdersSuite) AnAllowedDenom() {
	err := s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom: s.askPrice.Denom,
	})
	require.NoError(s.t, err)
}

func (s *updateSellOrdersSuite) AnAllowedDenomWithBankDenom(a string) {
	err := s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom:    a,
		DisplayDenom: a,
	})
	require.NoError(s.t, err)
}

func (s *updateSellOrdersSuite) AliceCreatedASellOrder() {
	s.sellOrderSetup(1)
}

func (s *updateSellOrdersSuite) AliceCreatedASellOrderWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.sellOrderID = id

	s.sellOrderSetup(1)
}

func (s *updateSellOrdersSuite) AliceCreatedASellOrderWithQuantity(a string) {
	s.quantity = a

	s.sellOrderSetup(1)
}

func (s *updateSellOrdersSuite) AliceCreatedTwoSellOrdersEachWithQuantity(a string) {
	s.quantity = a

	s.sellOrderSetup(2)
}

func (s *updateSellOrdersSuite) AliceCreatedASellOrderWithAskDenom(a string) {
	s.askPrice = &sdk.Coin{
		Denom:  a,
		Amount: s.askPrice.Amount,
	}

	s.sellOrderSetup(1)
}

func (s *updateSellOrdersSuite) AliceCreatedASellOrderWithBatchDenomAndAskDenom(a string, b string) {
	s.batchDenom = a
	s.askPrice = &sdk.Coin{
		Denom:  b,
		Amount: s.askPrice.Amount,
	}

	s.sellOrderSetup(1)
}

func (s *updateSellOrdersSuite) AliceCreatedASellOrderWithTheProperties(a gocuke.DocString) {
	order := &types.MsgSell_Order{}
	err := jsonpb.UnmarshalString(a.Content, order)
	require.NoError(s.t, err)

	s.batchDenom = order.BatchDenom
	s.quantity = order.Quantity
	s.askPrice = order.AskPrice
	s.disableAutoRetire = order.DisableAutoRetire
	s.expiration = order.Expiration

	s.sellOrderSetup(1)
}

func (s *updateSellOrdersSuite) AliceCreatedTwoSellOrdersEachWithTheProperties(a gocuke.DocString) {
	order := &types.MsgSell_Order{}
	err := jsonpb.UnmarshalString(a.Content, order)
	require.NoError(s.t, err)

	s.batchDenom = order.BatchDenom
	s.quantity = order.Quantity
	s.askPrice = order.AskPrice
	s.disableAutoRetire = order.DisableAutoRetire
	s.expiration = order.Expiration

	s.sellOrderSetup(2)
}

func (s *updateSellOrdersSuite) AliceHasABatchBalanceWithTradableAmountAndEscrowedAmount(a string, b string) {
	s.aliceTradableAmount = a
	s.aliceEscrowedAmount = b

	s.aliceBatchBalance()
}

func (s *updateSellOrdersSuite) AliceAttemptsToUpdateTheSellOrder() {
	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.alice.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId: s.sellOrderID,
				NewQuantity: s.quantity,
				NewAskPrice: s.askPrice,
			},
		},
	})
}

func (s *updateSellOrdersSuite) BobAttemptsToUpdateTheSellOrder() {
	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.bob.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId: s.sellOrderID,
				NewQuantity: s.quantity,
				NewAskPrice: s.askPrice,
			},
		},
	})
}

func (s *updateSellOrdersSuite) AliceAttemptsToUpdateTheSellOrderWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.alice.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId: id,
				NewQuantity: s.quantity,
				NewAskPrice: s.askPrice,
			},
		},
	})
}

func (s *updateSellOrdersSuite) AliceAttemptsToUpdateTheSellOrderWithQuantity(a string) {
	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.alice.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId: s.sellOrderID,
				NewQuantity: a,
				NewAskPrice: s.askPrice,
			},
		},
	})
}

func (s *updateSellOrdersSuite) AliceAttemptsToUpdateTheTwoSellOrdersEachWithQuantity(a string) {
	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.alice.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId: 1,
				NewQuantity: a,
				NewAskPrice: s.askPrice,
			},
			{
				SellOrderId: 2,
				NewQuantity: a,
				NewAskPrice: s.askPrice,
			},
		},
	})
}

func (s *updateSellOrdersSuite) AliceAttemptsToUpdateTheSellOrderWithAskDenom(a string) {
	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.alice.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId: s.sellOrderID,
				NewQuantity: s.quantity,
				NewAskPrice: &sdk.Coin{
					Denom:  a,
					Amount: s.askPrice.Amount,
				},
			},
		},
	})
}

func (s *updateSellOrdersSuite) AliceAttemptsToUpdateTheSellOrderWithExpiration(a string) {
	expiration, err := regentypes.ParseDate("expiration", a)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.alice.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId:   s.sellOrderID,
				NewQuantity:   s.quantity,
				NewAskPrice:   s.askPrice,
				NewExpiration: &expiration,
			},
		},
	})
}

func (s *updateSellOrdersSuite) AliceAttemptsToUpdateTheSellOrderWithTheProperties(a gocuke.DocString) {
	update := &types.MsgUpdateSellOrders_Update{}
	err := jsonpb.UnmarshalString(a.Content, update)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.alice.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId:       s.sellOrderID,
				NewQuantity:       update.NewQuantity,
				NewAskPrice:       update.NewAskPrice,
				DisableAutoRetire: update.DisableAutoRetire,
				NewExpiration:     update.NewExpiration,
			},
		},
	})
}

func (s *updateSellOrdersSuite) AliceAttemptsToUpdateTheTwoSellOrdersEachWithTheProperties(a gocuke.DocString) {
	update := &types.MsgUpdateSellOrders_Update{}
	err := jsonpb.UnmarshalString(a.Content, update)
	require.NoError(s.t, err)

	s.res, s.err = s.k.UpdateSellOrders(s.ctx, &types.MsgUpdateSellOrders{
		Seller: s.alice.String(),
		Updates: []*types.MsgUpdateSellOrders_Update{
			{
				SellOrderId:       1,
				NewQuantity:       update.NewQuantity,
				NewAskPrice:       update.NewAskPrice,
				DisableAutoRetire: update.DisableAutoRetire,
				NewExpiration:     update.NewExpiration,
			},
			{
				SellOrderId:       2,
				NewQuantity:       update.NewQuantity,
				NewAskPrice:       update.NewAskPrice,
				DisableAutoRetire: update.DisableAutoRetire,
				NewExpiration:     update.NewExpiration,
			},
		},
	})
}

func (s *updateSellOrdersSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *updateSellOrdersSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *updateSellOrdersSuite) ExpectAliceTradableCreditBalance(a string) {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, a, balance.TradableAmount)
}

func (s *updateSellOrdersSuite) ExpectAliceEscrowedCreditBalance(a string) {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, a, balance.EscrowedAmount)
}

func (s *updateSellOrdersSuite) ExpectMarketWithIdAndDenom(a string, b string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	market, err := s.marketStore.MarketTable().Get(s.ctx, id)
	require.NoError(s.t, err)

	require.Equal(s.t, market.Id, id)
	require.Equal(s.t, market.CreditTypeAbbrev, s.creditTypeAbbrev)
	require.Equal(s.t, market.BankDenom, b)
	require.Equal(s.t, market.PrecisionModifier, uint32(0)) // always zero
}

func (s *updateSellOrdersSuite) ExpectNoMarketWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	_, err = s.marketStore.MarketTable().Get(s.ctx, id)
	require.ErrorContains(s.t, err, ormerrors.NotFound.Error())
}

func (s *updateSellOrdersSuite) ExpectSellOrderWithSellerAliceAndTheProperties(a gocuke.DocString) {
	expected := &types.SellOrder{}
	err := jsonpb.UnmarshalString(a.Content, expected)
	require.NoError(s.t, err)

	order, err := s.marketStore.SellOrderTable().Get(s.ctx, expected.Id)
	require.NoError(s.t, err)

	require.Equal(s.t, expected.Id, order.Id)
	require.Equal(s.t, s.alice.Bytes(), order.Seller)
	require.Equal(s.t, expected.AskAmount, order.AskAmount)
	require.Equal(s.t, expected.Expiration.Seconds, order.Expiration.Seconds)
	require.Equal(s.t, expected.Expiration.Nanos, order.Expiration.Nanos)
	require.Equal(s.t, expected.BatchKey, order.BatchKey)
	require.Equal(s.t, expected.Quantity, order.Quantity)
	require.Equal(s.t, expected.DisableAutoRetire, order.DisableAutoRetire)
	require.Equal(s.t, expected.MarketId, order.MarketId)
	require.Equal(s.t, expected.Maker, order.Maker)
}

func (s *updateSellOrdersSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventUpdateSellOrder
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *updateSellOrdersSuite) sellOrderSetup(count int) {
	totalQuantity := s.quantity

	if count > 1 {
		c := math.NewDecFromInt64(int64(count))
		q, err := math.NewDecFromString(s.quantity)
		require.NoError(s.t, err)
		t, err := c.Mul(q)
		require.NoError(s.t, err)
		totalQuantity = t.String()
	}

	err := s.coreStore.ClassTable().Insert(s.ctx, &coreapi.Class{
		Id:               s.classID,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	batchKey, err := s.coreStore.BatchTable().InsertReturningID(s.ctx, &coreapi.Batch{
		Denom: s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchBalanceTable().Insert(s.ctx, &coreapi.BatchBalance{
		BatchKey:       batchKey,
		Address:        s.alice,
		EscrowedAmount: totalQuantity,
	})
	require.NoError(s.t, err)

	err = s.coreStore.BatchSupplyTable().Insert(s.ctx, &coreapi.BatchSupply{
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

	if s.expiration != nil {
		order.Expiration = timestamppb.New(*s.expiration)
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

func (s *updateSellOrdersSuite) aliceBatchBalance() {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	if err == ormerrors.NotFound {
		classKey, err := s.coreStore.ClassTable().InsertReturningID(s.ctx, &coreapi.Class{
			Id:               s.classID,
			CreditTypeAbbrev: s.creditTypeAbbrev,
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

		batch.Key = batchKey
	} else if err != nil {
		require.NoError(s.t, err)
	}

	// Save because batch balance may already exist from sell order setup
	err = s.coreStore.BatchBalanceTable().Save(s.ctx, &coreapi.BatchBalance{
		BatchKey:       batch.Key,
		Address:        s.alice,
		TradableAmount: s.aliceTradableAmount,
		EscrowedAmount: s.aliceEscrowedAmount,
	})
	require.NoError(s.t, err)
}
