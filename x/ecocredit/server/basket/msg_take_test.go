package basket_test

import (
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/regen-network/regen-ledger/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type takeSuite struct {
	*baseSuite
	alice               sdk.AccAddress
	aliceTokenBalance   sdk.Coin
	basketTokenSupply   sdk.Coin
	classId             string
	creditTypeAbbrev    string
	creditTypePrecision uint32
	batchDenom          string
	basketDenom         string
	tokenAmount         string
	jurisdiction        string
	res                 *basket.MsgTakeResponse
	err                 error
}

func TestTake(t *testing.T) {
	gocuke.NewRunner(t, &takeSuite{}).Path("./features/msg_take.feature").Run()
}

func (s *takeSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.aliceTokenBalance = sdk.Coin{
		Denom:  "eco.uC.NCT",
		Amount: sdk.NewInt(100),
	}
	s.basketTokenSupply = sdk.Coin{
		Denom:  "eco.uC.NCT",
		Amount: sdk.NewInt(100),
	}
	s.classId = "C01"
	s.creditTypeAbbrev = "C"
	s.creditTypePrecision = 6
	s.batchDenom = "C01-001-20200101-20210101-001"
	s.basketDenom = "eco.uC.NCT"
	s.tokenAmount = "100"
	s.jurisdiction = "US-WA"

	// TODO: remove after updating to cosmos-sdk v0.46 #857
	sdk.SetCoinDenomRegex(func() string {
		return types.CoinDenomRegex
	})
}

func (s *takeSuite) ACreditType() {
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *takeSuite) ACreditTypeWithAbbreviation(a string) {
	s.creditTypeAbbrev = a

	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *takeSuite) ACreditTypeWithAbbreviationAndPrecision(a string, b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	s.creditTypeAbbrev = a
	s.creditTypePrecision = uint32(precision)

	err = s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *takeSuite) ABasket() {
	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	// add balance with credit amount = token amount
	s.addBasketClassAndBalance(basketId, s.tokenAmount)
}

func (s *takeSuite) ABasketWithDenom(a string) {
	s.basketDenom = a

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	// add balance with credit amount = token amount
	s.addBasketClassAndBalance(basketId, s.tokenAmount)
}

func (s *takeSuite) ABasketWithDisableAutoRetire(a string) {
	disableAutoRetire, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:       s.basketDenom,
		CreditTypeAbbrev:  s.creditTypeAbbrev,
		DisableAutoRetire: disableAutoRetire,
	})
	require.NoError(s.t, err)

	// add balance with credit amount = token amount
	s.addBasketClassAndBalance(basketId, s.tokenAmount)
}

func (s *takeSuite) ABasketWithCreditBalance(a string) {
	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	s.addBasketClassAndBalance(basketId, a)
}

func (s *takeSuite) ABasketWithCreditTypeAndDisableAutoRetire(a string, b string) {
	s.creditTypeAbbrev = a

	disableAutoRetire, err := strconv.ParseBool(b)
	require.NoError(s.t, err)

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:       s.basketDenom,
		CreditTypeAbbrev:  s.creditTypeAbbrev,
		DisableAutoRetire: disableAutoRetire,
	})
	require.NoError(s.t, err)

	// add balance with credit amount = token amount
	s.addBasketClassAndBalance(basketId, s.tokenAmount)
}

func (s *takeSuite) ABasketWithCreditTypeAndCreditBalance(a string, b string) {
	s.creditTypeAbbrev = a

	basketId, err := s.stateStore.BasketTable().InsertReturningID(s.ctx, &api.Basket{
		BasketDenom:      s.basketDenom,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)

	s.addBasketClassAndBalance(basketId, b)
}

func (s *takeSuite) AliceOwnsBasketTokens() {
	amount, ok := sdk.NewIntFromString(s.tokenAmount)
	require.True(s.t, ok)

	s.aliceTokenBalance = sdk.NewCoin(s.basketDenom, amount)
}

func (s *takeSuite) AliceOwnsBasketTokenAmount(a string) {
	amount, ok := sdk.NewIntFromString(a)
	require.True(s.t, ok)

	s.aliceTokenBalance = sdk.NewCoin(s.basketDenom, amount)
}

func (s *takeSuite) AliceOwnsTokensWithDenom(a string) {
	amount, ok := sdk.NewIntFromString(s.tokenAmount)
	require.True(s.t, ok)

	s.aliceTokenBalance = sdk.NewCoin(a, amount)
}

func (s *takeSuite) BasketTokenSupplyAmount(a string) {
	amount, err := strconv.ParseInt(a, 10, 32)
	require.NoError(s.t, err)

	coin := sdk.NewInt64Coin(s.basketDenom, amount)

	s.basketTokenSupply = coin
}

func (s *takeSuite) AliceAttemptsToTakeCreditsWithBasketDenom(a string) {
	s.basketDenom = a

	s.takeExpectCalls()

	s.res, s.err = s.k.Take(s.ctx, &basket.MsgTake{
		Owner:                  s.alice.String(),
		BasketDenom:            s.basketDenom,
		Amount:                 s.aliceTokenBalance.Amount.String(),
		RetirementJurisdiction: s.jurisdiction,
		RetireOnTake:           true, // satisfy default auto-retire
	})
}

func (s *takeSuite) AliceAttemptsToTakeCreditsWithBasketTokenAmount(a string) {
	s.tokenAmount = a

	s.takeExpectCalls()

	s.res, s.err = s.k.Take(s.ctx, &basket.MsgTake{
		Owner:                  s.alice.String(),
		BasketDenom:            s.basketDenom,
		Amount:                 s.tokenAmount,
		RetirementJurisdiction: s.jurisdiction,
		RetireOnTake:           true, // satisfy default auto-retire
	})
}

func (s *takeSuite) AliceAttemptsToTakeCreditsWithBasketTokenAmountAndRetireOnTake(a string, b string) {
	s.tokenAmount = a

	retireOnTake, err := strconv.ParseBool(b)
	require.NoError(s.t, err)

	s.takeExpectCalls()

	s.res, s.err = s.k.Take(s.ctx, &basket.MsgTake{
		Owner:                  s.alice.String(),
		BasketDenom:            s.basketDenom,
		Amount:                 s.tokenAmount,
		RetirementJurisdiction: s.jurisdiction,
		RetireOnTake:           retireOnTake,
	})
}

func (s *takeSuite) AliceAttemptsToTakeCreditsWithRetireOnTake(a string) {
	retireOnTake, err := strconv.ParseBool(a)
	require.NoError(s.t, err)

	s.takeExpectCalls()

	s.res, s.err = s.k.Take(s.ctx, &basket.MsgTake{
		Owner:                  s.alice.String(),
		BasketDenom:            s.basketDenom,
		Amount:                 s.aliceTokenBalance.Amount.String(),
		RetirementJurisdiction: s.jurisdiction,
		RetireOnTake:           retireOnTake,
	})
}

func (s *takeSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *takeSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *takeSuite) ExpectAliceTradableCreditBalanceAmount(a string) {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, a, balance.TradableAmount)
}

func (s *takeSuite) ExpectAliceRetiredCreditBalanceAmount(a string) {
	batch, err := s.coreStore.BatchTable().GetByDenom(s.ctx, s.batchDenom)
	require.NoError(s.t, err)

	balance, err := s.coreStore.BatchBalanceTable().Get(s.ctx, s.alice, batch.Key)
	require.NoError(s.t, err)

	require.Equal(s.t, a, balance.RetiredAmount)
}

func (s *takeSuite) ExpectAliceBasketTokenBalanceAmount(a string) {
	amount, err := strconv.ParseInt(a, 10, 32)
	require.NoError(s.t, err)

	tokenBalance := sdk.NewInt64Coin(s.basketDenom, amount)

	require.Equal(s.t, tokenBalance, s.aliceTokenBalance)
}

func (s *takeSuite) ExpectBasketCreditBalanceAmount(a string) {
	basket, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, s.basketDenom)
	require.NoError(s.t, err)

	balance, err := s.stateStore.BasketBalanceTable().Get(s.ctx, basket.Id, s.batchDenom)
	require.NoError(s.t, err)

	require.Equal(s.t, a, balance.Balance)
}

func (s *takeSuite) ExpectBasketTokenSupplyAmount(a string) {
	amount, err := strconv.ParseInt(a, 10, 32)
	require.NoError(s.t, err)

	tokenSupply := sdk.NewInt64Coin(s.basketDenom, amount)

	require.Equal(s.t, tokenSupply, s.basketTokenSupply)
}

func (s *takeSuite) ExpectTheResponse(a gocuke.DocString) {
	res := &basket.MsgTakeResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *takeSuite) addBasketClassAndBalance(basketId uint64, creditAmount string) {
	err := s.stateStore.BasketClassTable().Insert(s.ctx, &api.BasketClass{
		BasketId: basketId,
		ClassId:  s.classId,
	})
	require.NoError(s.t, err)

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

	err = s.coreStore.BatchSupplyTable().Insert(s.ctx, &coreapi.BatchSupply{
		BatchKey:       batchKey,
		TradableAmount: creditAmount,
	})
	require.NoError(s.t, err)

	err = s.stateStore.BasketBalanceTable().Insert(s.ctx, &api.BasketBalance{
		BasketId:   basketId,
		BatchDenom: s.batchDenom,
		Balance:    creditAmount,
	})
	require.NoError(s.t, err)
}

func (s *takeSuite) takeExpectCalls() {
	amount, ok := sdk.NewIntFromString(s.tokenAmount)
	require.True(s.t, ok)

	sendCoin := sdk.NewCoin(s.basketDenom, amount)
	sendCoins := sdk.NewCoins(sendCoin)

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.alice, s.basketDenom).
		Return(s.aliceTokenBalance).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromAccountToModule(s.sdkCtx, s.alice, basket.BasketSubModuleName, sendCoins).
		Do(func(sdk.Context, sdk.AccAddress, string, sdk.Coins) {
			// simulate token balance update unavailable with mocks
			s.aliceTokenBalance = s.aliceTokenBalance.Sub(sendCoin)
		}).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		BurnCoins(s.sdkCtx, basket.BasketSubModuleName, sendCoins).
		Do(func(sdk.Context, string, sdk.Coins) {
			// simulate token supply update unavailable with mocks
			s.basketTokenSupply = s.basketTokenSupply.Sub(sendCoin)
		}).
		Return(nil).
		AnyTimes() // not expected on failed attempt
}
