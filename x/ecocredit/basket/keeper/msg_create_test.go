//nolint:revive,stylecheck
package keeper

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	"github.com/regen-network/regen-ledger/types/v2/testutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/base"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/basket"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

type createSuite struct {
	*baseSuite
	alice               sdk.AccAddress
	aliceBalance        sdk.Coin
	basketName          string
	creditTypeAbbrev    string
	creditTypePrecision uint32
	res                 *types.MsgCreateResponse
	err                 error
	basketFee           sdk.Coin
}

func TestCreate(t *testing.T) {
	gocuke.NewRunner(t, &createSuite{}).Path("./features/msg_create.feature").Run()
}

func (s *createSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.basketName = "NCT"
	s.creditTypeAbbrev = "C"
	s.creditTypePrecision = 6
}

func (s *createSuite) RequiredBasketFee(a string) {
	basketFee, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	basketFeeProto := regentypes.CoinToProtoCoin(basketFee)

	err = s.stateStore.BasketFeeTable().Save(s.ctx, &api.BasketFee{
		Fee: basketFeeProto,
	})
	require.NoError(s.t, err)

	s.basketFee = basketFee
}

func (s *createSuite) ACreditType() {
	err := s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: a,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithPrecision(b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	s.creditTypePrecision = uint32(precision)

	err = s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithAbbreviationAndPrecision(a string, b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	s.creditTypePrecision = uint32(precision)

	err = s.baseStore.CreditTypeTable().Insert(s.ctx, &baseapi.CreditType{
		Abbreviation: a,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) AlicesAddress(a string) {
	addr, err := sdk.AccAddressFromBech32(a)
	require.NoError(s.t, err)
	s.alice = addr
}

func (s *createSuite) ACreditClassWithId(a string) {
	creditTypeAbbrev := base.GetCreditTypeAbbrevFromClassID(a)

	err := s.baseStore.ClassTable().Insert(s.ctx, &baseapi.Class{
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

	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		Fee:              basketFee,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithNoFee() {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithCreditType(a string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: a,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithCreditTypeAndAllowedClass(a string, b string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: a,
		AllowedClasses:   []string{b},
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithAllowedClass(a string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		AllowedClasses:   []string{a},
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithName(a string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             a,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithNameAndCreditType(a string, b string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             a,
		CreditTypeAbbrev: b,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithMinimumStartDate(a string) {
	s.createExpectCalls()

	ts, err := regentypes.ParseDate(a, a)
	require.NoError(s.t, err)

	minStartDate, err := gogotypes.TimestampProto(ts)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		DateCriteria: &types.DateCriteria{
			MinStartDate: minStartDate,
		},
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithStartDateWindow(a string) {
	s.createExpectCalls()

	dur, err := time.ParseDuration(a)
	require.NoError(s.t, err)

	startDateWindow := gogotypes.DurationProto(dur)

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		DateCriteria: &types.DateCriteria{
			StartDateWindow: startDateWindow,
		},
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithYearsInThePast(a string) {
	s.createExpectCalls()

	yearsInThePast, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	s.res, s.err = s.k.Create(s.ctx, &types.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		DateCriteria: &types.DateCriteria{
			YearsInThePast: uint32(yearsInThePast),
		},
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
	res := &types.MsgCreateResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *createSuite) ExpectMinimumStartDate(a string) {
	ts, err := regentypes.ParseDate(a, a)
	require.NoError(s.t, err)

	minStartDate, err := gogotypes.TimestampProto(ts)
	require.NoError(s.t, err)

	dc, err := s.k.Basket(s.ctx, &types.QueryBasketRequest{
		BasketDenom: s.res.BasketDenom,
	})
	require.NoError(s.t, err)

	require.Equal(s.t, minStartDate, dc.BasketInfo.DateCriteria.MinStartDate)
}

func (s *createSuite) ExpectStartDateWindow(a string) {
	dur, err := time.ParseDuration(a)
	require.NoError(s.t, err)

	startDateWindow := gogotypes.DurationProto(dur)

	dc, err := s.k.Basket(s.ctx, &types.QueryBasketRequest{
		BasketDenom: s.res.BasketDenom,
	})
	require.NoError(s.t, err)
	require.Equal(s.t, startDateWindow, dc.BasketInfo.DateCriteria.StartDateWindow)
}

func (s *createSuite) ExpectYearsInThePast(a string) {
	yearsInThePast, err := strconv.ParseUint(a, 10, 32)
	require.NoError(s.t, err)

	dc, err := s.k.Basket(s.ctx, &types.QueryBasketRequest{
		BasketDenom: s.res.BasketDenom,
	})
	require.NoError(s.t, err)

	require.Equal(s.t, uint32(yearsInThePast), dc.BasketInfo.DateCriteria.YearsInThePast)
}

func (s *createSuite) ExpectEventWithProperties(a gocuke.DocString) {
	var event types.EventCreate
	err := json.Unmarshal([]byte(a.Content), &event)
	require.NoError(s.t, err)

	sdkEvent, found := testutil.GetEvent(&event, s.sdkCtx.EventManager().Events())
	require.True(s.t, found)

	err = testutil.MatchEvent(&event, sdkEvent)
	require.NoError(s.t, err)
}

func (s *createSuite) createExpectCalls() {

	var expectedFee sdk.Coin
	var expectedFees sdk.Coins

	if !s.basketFee.IsNil() {
		expectedFee = s.basketFee
		expectedFees = sdk.Coins{expectedFee}
	}

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.alice, expectedFee.Denom).
		Return(s.aliceBalance).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromAccountToModule(s.sdkCtx, s.alice, basket.BasketSubModuleName, expectedFees).
		Do(func(sdk.Context, sdk.AccAddress, string, sdk.Coins) {
			if !s.basketFee.IsNil() {
				// simulate token balance update unavailable with mocks
				s.aliceBalance = s.aliceBalance.Sub(expectedFee)
			}
		}).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		BurnCoins(s.sdkCtx, basket.BasketSubModuleName, expectedFees).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt
}

func (s *createSuite) getDenomMetadata() bank.Metadata {
	denom, displayDenom, err := basket.FormatBasketDenom(s.basketName, s.creditTypeAbbrev, s.creditTypePrecision)
	require.NoError(s.t, err)

	denomUnits := make([]*bank.DenomUnit, 0)
	if denom != displayDenom {
		denomUnits = append(denomUnits, &bank.DenomUnit{
			Denom: denom,
		})
	}
	denomUnits = append(denomUnits, &bank.DenomUnit{
		Denom:    displayDenom,
		Exponent: s.creditTypePrecision,
	})

	return bank.Metadata{
		Name:       s.basketName,
		Display:    displayDenom,
		Base:       denom,
		Symbol:     s.basketName,
		DenomUnits: denomUnits,
	}
}
