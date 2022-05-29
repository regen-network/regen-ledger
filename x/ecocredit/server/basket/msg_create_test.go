package basket_test

import (
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type createSuite struct {
	*baseSuite
	alice            sdk.AccAddress
	aliceBalance     sdk.Coin
	minBasketFee     sdk.Coins
	basketName       string
	basketExponent   uint32
	creditTypeAbbrev string
	res              *basket.MsgCreateResponse
	err              error
}

func TestCreate(t *testing.T) {
	gocuke.NewRunner(t, &createSuite{}).Path("./features/msg_create.feature").Run()
}

func (s *createSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.basketName = "NCT"
	s.basketExponent = 6
	s.creditTypeAbbrev = "C"
}

func (s *createSuite) AMinimumBasketFee(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.minBasketFee = sdk.NewCoins(coin)
}

func (s *createSuite) ACreditType() {
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: a,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithPrecision(b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	err = s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithAbbreviationAndPrecision(a string, b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	err = s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: a,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditClassWithId(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(a)

	err := s.coreStore.ClassTable().Insert(s.ctx, &coreapi.Class{
		Id:               a,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ABasketWithName(a string) {
	err := s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		Name: a,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) AliceHasATokenBalance(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.aliceBalance = coin
}

func (s *createSuite) AliceAttemptsToCreateABasketWithFee(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	basketFee := sdk.NewCoins(coin)

	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.minBasketFee
		}).
		Times(1)

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.alice, coin.Denom).
		Return(s.aliceBalance).
		AnyTimes() // not expected on failed attempt

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.minBasketFee, s.alice).
		Do(func(sdk.Context, sdk.Coins, sdk.AccAddress) {
			if s.minBasketFee != nil {
				// simulate token balance update unavailable with mocks
				s.aliceBalance = s.aliceBalance.Sub(s.minBasketFee[0])
			}
		}).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		Exponent:         s.basketExponent,
		Fee:              basketFee,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithNoFee() {
	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.minBasketFee
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.minBasketFee, s.alice).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		Exponent:         s.basketExponent,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithCreditType(a string) {
	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.minBasketFee
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.minBasketFee, s.alice).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		Exponent:         s.basketExponent,
		CreditTypeAbbrev: a,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithCreditTypeAndAllowedClass(a string, b string) {
	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.minBasketFee
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.minBasketFee, s.alice).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		Exponent:         s.basketExponent,
		CreditTypeAbbrev: a,
		AllowedClasses:   []string{b},
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithAllowedClass(a string) {
	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.minBasketFee
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.minBasketFee, s.alice).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		Exponent:         s.basketExponent,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		AllowedClasses:   []string{a},
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithName(a string) {
	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.minBasketFee
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.minBasketFee, s.alice).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             a,
		Exponent:         s.basketExponent,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithExponent(a string) {
	exponent, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	// set exponent for denom metadata
	s.basketExponent = uint32(exponent)

	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.minBasketFee
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.minBasketFee, s.alice).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		Exponent:         uint32(exponent),
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithNameAndExponent(a string, b string) {
	exponent, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	// set exponent for denom metadata
	s.basketExponent = uint32(exponent)

	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.minBasketFee
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.minBasketFee, s.alice).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             a,
		Exponent:         uint32(exponent),
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createSuite) ExpectAliceTokenBalance(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	require.Equal(s.t, coin, s.aliceBalance)
}

func (s *createSuite) ExpectTheResponse(a gocuke.DocString) {
	res := &basket.MsgCreateResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *createSuite) getDenomMetadata() bank.Metadata {
	denom, displayDenom, err := basket.FormatBasketDenom(s.basketName, s.creditTypeAbbrev, s.basketExponent)
	require.NoError(s.t, err)

	denomUnits := []*bank.DenomUnit{{
		Denom:    displayDenom,
		Exponent: s.basketExponent,
	}}

	if denom != displayDenom {
		denomUnits = append(denomUnits, &bank.DenomUnit{
			Denom: denom,
		})
	}

	return bank.Metadata{
		Name:       s.basketName,
		Display:    displayDenom,
		Base:       denom,
		Symbol:     s.basketName,
		DenomUnits: denomUnits,
	}
}
