package basket_test

import (
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type createSuite struct {
	*baseSuite
	alice             sdk.AccAddress
	bob               sdk.AccAddress
	basketFeeParam    sdk.Coins
	aliceTokenBalance sdk.Coin
	bobTokenBalance   sdk.Coin
	basketName        string
	basketExponent    uint32
	classId           string
	creditTypeAbbrev  string
	tradableCredits   string
	res               *basket.MsgCreateResponse
	err               error
}

func TestCreate(t *testing.T) {
	gocuke.NewRunner(t, &createSuite{}).Path("./features/msg_create.feature").Run()
}

func (s *createSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.bob = s.addrs[1]
	s.basketName = "NCT"
	s.basketExponent = 6
	s.classId = "C01"
	s.creditTypeAbbrev = "C"
}

func (s *createSuite) TheBasketFeeParam(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.basketFeeParam = sdk.NewCoins(coin)
}

func (s *createSuite) TheCreditTypeWithAbbreviation(a string) {
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: a,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) TheCreditTypeWithAbbreviationAndPrecision(a string, b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	err = s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: a,
		Precision:    uint32(precision),
	})
	require.NoError(s.t, err)
}

func (s *createSuite) TheCreditClassWithId(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(a)

	err := s.coreStore.ClassTable().Insert(s.ctx, &coreapi.Class{
		Id:               a,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) TheBasketWithName(a string) {
	err := s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		Name: a,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) TheBasketWithDenom(a string) {
	err := s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		BasketDenom: a,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) AliceHasATokenBalance(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.aliceTokenBalance = coin
}

func (s *createSuite) AliceAttemptsToCreateABasketWithFee(a string) {
	var coins sdk.Coins

	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	basketFee := sdk.NewCoins(coin)

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.basketFeeParam
		}).
		Times(1)

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.alice, coin.Denom).
		Return(s.aliceTokenBalance).
		AnyTimes() // not expected on failed attempt

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.basketFeeParam, s.alice).
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
			*coins = s.basketFeeParam
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.basketFeeParam, s.alice).
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
			*coins = s.basketFeeParam
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.basketFeeParam, s.alice).
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

func (s *createSuite) AliceAttemptsToCreateABasketWithAllowedClass(a string) {
	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.basketFeeParam
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.basketFeeParam, s.alice).
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
			*coins = s.basketFeeParam
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.basketFeeParam, s.alice).
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

func (s *createSuite) AliceAttemptsToCreateABasketWithNameAndExponent(a string, b string) {
	exponent, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	s.basketExponent = uint32(exponent)

	var coins sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketCreationFee, &coins).
		Do(func(ctx sdk.Context, key []byte, coins *sdk.Coins) {
			*coins = s.basketFeeParam
		}).
		Times(1)

	s.distKeeper.EXPECT().
		FundCommunityPool(s.sdkCtx, s.basketFeeParam, s.alice).
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

func (s *createSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createSuite) ExpectTheResponse(a gocuke.DocString) {
	res := &basket.MsgCreateResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *createSuite) getDenomMetadata() banktypes.Metadata {
	denom, displayDenom, err := basket.FormatBasketDenom(s.basketName, s.creditTypeAbbrev, s.basketExponent)
	require.NoError(s.t, err)

	denomUnits := []*banktypes.DenomUnit{{
		Denom:    displayDenom,
		Exponent: s.basketExponent,
	}}

	if denom != displayDenom {
		denomUnits = append(denomUnits, &banktypes.DenomUnit{
			Denom:    denom,
			Exponent: 0,
		})
	}

	return banktypes.Metadata{
		Name:       s.basketName,
		Display:    displayDenom,
		Base:       denom,
		Symbol:     s.basketName,
		DenomUnits: denomUnits,
	}
}
