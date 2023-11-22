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

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/marketplace/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

type sellSuite struct {
	*baseSuite
	alice               sdk.AccAddress
	aliceTradableAmount string
	creditTypeAbbrev    string
	classID             string
	batchDenom          string
	askPrice            *sdk.Coin
	quantity            string
	expiration          *time.Time
	res                 *types.MsgSellResponse
	err                 error
}

func TestSell(t *testing.T) {
	gocuke.NewRunner(t, &sellSuite{}).Path("./features/msg_sell.feature").Run()
}

func (s *sellSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t, 1)
	s.alice = s.addrs[0]
	s.aliceTradableAmount = "200"
	s.creditTypeAbbrev = "C"
	s.classID = testClassID
	s.batchDenom = testBatchDenom
	s.askPrice = &sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}
	s.quantity = "100"

	expiration, err := regentypes.ParseDate("expiration", "2030-01-01")
	require.NoError(s.t, err)

	s.expiration = &expiration
}

func (s *sellSuite) ABlockTimeWithTimestamp(a string) {
	blockTime, err := regentypes.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
}

func (s *sellSuite) ACreditType() {
	err := s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: a,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) ACreditTypeWithPrecision(a string) {
	precision, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	err = s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AnAllowedDenom() {
	err := s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom: s.askPrice.Denom,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AnAllowedDenomWithBankDenom(a string) {
	err := s.marketStore.AllowedDenomTable().Insert(s.ctx, &api.AllowedDenom{
		BankDenom: a,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) ACreditBatchWithBatchDenom(a string) {
	classID := base.GetClassIDFromBatchDenom(a)
	creditTypeAbbrev := base.GetCreditTypeAbbrevFromClassID(classID)

	classKey, err := s.baseStore.ClassTable().InsertReturningID(s.ctx, &baseapi.Class{
		Id:               classID,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.baseStore.ProjectTable().InsertReturningID(s.ctx, &baseapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	err = s.baseStore.BatchTable().Insert(s.ctx, &baseapi.Batch{
		ProjectKey: projectKey,
		Denom:      a,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AMarketWithCreditTypeAndBankDenom(a string, b string) {
	err := s.marketStore.MarketTable().Insert(s.ctx, &api.Market{
		CreditTypeAbbrev: a,
		BankDenom:        b,
	})
	require.NoError(s.t, err)
}

func (s *sellSuite) AliceHasATradableBatchBalance() {
	s.aliceTradableBatchBalance()
}

func (s *sellSuite) AliceHasATradableBatchBalanceWithDenom(a string) {
	s.batchDenom = a
	s.classID = base.GetClassIDFromBatchDenom(s.batchDenom)
	s.creditTypeAbbrev = base.GetCreditTypeAbbrevFromClassID(s.classID)

	s.aliceTradableBatchBalance()
}

func (s *sellSuite) AliceHasATradableBatchBalanceWithAmount(a string) {
	s.aliceTradableAmount = a

	s.aliceTradableBatchBalance()
}

func (s *sellSuite) AliceHasATradableBatchBalanceWithDenomAndAmount(a string, b string) {
	s.batchDenom = a
	s.classID = base.GetClassIDFromBatchDenom(s.batchDenom)
	s.creditTypeAbbrev = base.GetCreditTypeAbbrevFromClassID(s.classID)
	s.aliceTradableAmount = b

	s.aliceTradableBatchBalance()
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithBatchDenom(a string) {
	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom: a,
				Quantity:   s.quantity,
				AskPrice:   s.askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithBatchDenomAndAskDenom(a string, b string) {
	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom: a,
				Quantity:   s.quantity,
				AskPrice: &sdk.Coin{
					Denom:  b,
					Amount: s.askPrice.Amount,
				},
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithCreditQuantity(a string) {
	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   a,
				AskPrice:   s.askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithTheProperties(a gocuke.DocString) {
	order := &types.MsgSell_Order{}
	err := jsonpb.UnmarshalString(a.Content, order)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{order},
	})
}

func (s *sellSuite) AliceAttemptsToCreateTwoSellOrdersEachWithCreditQuantity(a string) {
	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   a,
				AskPrice:   s.askPrice,
			},
			{
				BatchDenom: s.batchDenom,
				Quantity:   a,
				AskPrice:   s.askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithAskPrice(a string) {
	askPrice, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   s.quantity,
				AskPrice:   &askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithAskDenom(a string) {
	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   s.quantity,
				AskPrice: &sdk.Coin{
					Denom:  a,
					Amount: s.askPrice.Amount,
				},
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithExpiration(a string) {
	expiration, err := regentypes.ParseDate("expiration", a)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   s.quantity,
				AskPrice:   s.askPrice,
				Expiration: &expiration,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrder() {
	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom:        s.batchDenom,
				Quantity:          s.quantity,
				AskPrice:          s.askPrice,
				DisableAutoRetire: true, // verify non-default is set
				Expiration:        s.expiration,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateTwoSellOrders() {
	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{
			{
				BatchDenom:        s.batchDenom,
				Quantity:          s.quantity,
				AskPrice:          s.askPrice,
				DisableAutoRetire: true, // verify non-default is set
				Expiration:        s.expiration,
			},
			{
				BatchDenom:        s.batchDenom,
				Quantity:          s.quantity,
				AskPrice:          s.askPrice,
				DisableAutoRetire: true, // verify non-default is set
				Expiration:        s.expiration,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateTwoSellOrdersEachWithTheProperties(a gocuke.DocString) {
	order := &types.MsgSell_Order{}
	err := jsonpb.UnmarshalString(a.Content, order)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Sell(s.ctx, &types.MsgSell{
		Seller: s.alice.String(),
		Orders: []*types.MsgSell_Order{order, order},
	})
}

func (s *sellSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *sellSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *sellSuite) ExpectAliceTradableCreditBalance(a string) {
	batch, err := s.baseStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.baseStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, balance.TradableAmount, a)
}

func (s *sellSuite) ExpectAliceEscrowedCreditBalance(a string) {
	batch, err := s.baseStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.baseStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, balance.EscrowedAmount, a)
}

func (s *sellSuite) ExpectMarketWithIdAndDenom(a string, b string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	market, err := s.marketStore.MarketTable().Get(s.ctx, id)
	require.NoError(s.t, err)

	require.Equal(s.t, market.Id, id)
	require.Equal(s.t, market.CreditTypeAbbrev, s.creditTypeAbbrev)
	require.Equal(s.t, market.BankDenom, b)
	require.Equal(s.t, market.PrecisionModifier, uint32(0)) // always zero
}

func (s *sellSuite) ExpectNoMarketWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	_, err = s.marketStore.MarketTable().Get(s.ctx, id)
	require.ErrorContains(s.t, err, ormerrors.NotFound.Error())
}

func (s *sellSuite) ExpectSellOrderWithSellerAliceAndTheProperties(a gocuke.DocString) {
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

func (s *sellSuite) ExpectTheResponse(a gocuke.DocString) {
	res := &types.MsgSellResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *sellSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventSell
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *sellSuite) aliceTradableBatchBalance() {
	classKey, err := s.baseStore.ClassTable().InsertReturningID(s.ctx, &baseapi.Class{
		Id:               s.classID,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	projectKey, err := s.baseStore.ProjectTable().InsertReturningID(s.ctx, &baseapi.Project{
		ClassKey: classKey,
	})
	require.NoError(s.t, err)

	batchKey, err := s.baseStore.BatchTable().InsertReturningID(s.ctx, &baseapi.Batch{
		ProjectKey: projectKey,
		Denom:      s.batchDenom,
	})
	require.NoError(s.t, err)

	err = s.baseStore.BatchBalanceTable().Insert(s.ctx, &baseapi.BatchBalance{
		BatchKey:       batchKey,
		Address:        s.alice,
		TradableAmount: s.aliceTradableAmount,
	})
	require.NoError(s.t, err)
}
