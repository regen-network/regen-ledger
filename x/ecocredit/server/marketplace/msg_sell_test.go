package marketplace

import (
	"strconv"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/regen-network/regen-ledger/types"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

type sellSuite struct {
	*baseSuite
	alice             sdk.AccAddress
	aliceBatchBalance string
	creditTypeAbbrev  string
	batchDenom        string
	askPrice          *sdk.Coin
	quantity          string
	expiration        *time.Time
	res               *marketplace.MsgSellResponse
	err               error
}

func TestSell(t *testing.T) {
	gocuke.NewRunner(t, &sellSuite{}).Path("./features/msg_sell.feature").Run()
}

func (s *sellSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addr
	s.aliceBatchBalance = "200"
	s.creditTypeAbbrev = "C"
	s.batchDenom = "C01-001-20200101-20210101-001"
	s.askPrice = &sdk.Coin{
		Denom:  "regen",
		Amount: sdk.NewInt(100),
	}
	s.quantity = "100"

	expiration, err := types.ParseDate("expiration", "2030-01-01")
	require.NoError(s.t, err)

	s.expiration = &expiration
}

func (s *sellSuite) ABlockTimeWithTimestamp(a string) {
	blockTime, err := types.ParseDate("block time", a)
	require.NoError(s.t, err)

	s.sdkCtx = s.sdkCtx.WithBlockTime(blockTime)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)
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

func (s *sellSuite) AMarketWithCreditTypeAndBankDenom(a string, b string) {
	s.marketStore.MarketTable().Insert(s.ctx, &api.Market{
		CreditType: a,
		BankDenom:  b,
	})
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
		Tradable: s.aliceBatchBalance,
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
		Tradable: s.quantity,
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
	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom: a,
				Quantity:   s.quantity,
				AskPrice:   s.askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithBatchDenomAndAskDenom(a string, b string) {
	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
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
	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   a,
				AskPrice:   s.askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateTwoSellOrdersEachWithCreditQuantity(a string) {
	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
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

	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
			{
				BatchDenom: s.batchDenom,
				Quantity:   s.quantity,
				AskPrice:   &askPrice,
			},
		},
	})
}

func (s *sellSuite) AliceAttemptsToCreateASellOrderWithAskDenom(a string) {
	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
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
	expiration, err := types.ParseDate("expiration", a)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
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
	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
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
	s.res, s.err = s.k.Sell(s.ctx, &marketplace.MsgSell{
		Owner: s.alice.String(),
		Orders: []*marketplace.MsgSell_Order{
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

func (s *sellSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *sellSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *sellSuite) ExpectAliceTradableCreditBalance(a string) {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, balance.Tradable, a)
}

func (s *sellSuite) ExpectAliceEscrowedCreditBalance(a string) {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, balance.Escrowed, a)
}

func (s *sellSuite) ExpectMarketWithIdAndDenom(a string, b string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	market, err := s.marketStore.MarketTable().Get(s.ctx, id)
	require.NoError(s.t, err)

	require.Equal(s.t, market.Id, id)
	require.Equal(s.t, market.CreditType, s.creditTypeAbbrev) // TODO: credit_type_abbrev
	require.Equal(s.t, market.BankDenom, b)
	require.Equal(s.t, market.PrecisionModifier, uint32(0)) // always zero
}

func (s *sellSuite) ExpectNoMarketWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	_, err = s.marketStore.MarketTable().Get(s.ctx, id)
	require.ErrorContains(s.t, err, ormerrors.NotFound.Error())
}

func (s *sellSuite) ExpectSellOrderWithId(a string) {
	id, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	order, err := s.marketStore.SellOrderTable().Get(s.ctx, id)
	require.NoError(s.t, err)

	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	market, err := s.marketStore.MarketTable().GetByCreditTypeBankDenom(s.ctx, s.creditTypeAbbrev, s.askPrice.Denom)
	require.NoError(s.t, err)

	expiration := order.Expiration.AsTime()

	require.Equal(s.t, order.Id, id)
	require.Equal(s.t, order.Seller, s.alice.Bytes())
	require.Equal(s.t, order.AskPrice, s.askPrice.Amount.String()) // TODO: ask_amount
	require.Equal(s.t, &expiration, s.expiration)
	require.Equal(s.t, order.BatchId, batch.Key) // TODO: batch_key
	require.Equal(s.t, order.Quantity, s.quantity)
	require.Equal(s.t, order.DisableAutoRetire, true) // verify non-default is set
	require.Equal(s.t, order.MarketId, market.Id)
	require.Equal(s.t, order.Maker, true) // always true
}

func (s *sellSuite) ExpectTheResponse(a gocuke.DocString) {
	res := &marketplace.MsgSellResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}
