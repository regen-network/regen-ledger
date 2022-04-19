package basket_test

import (
	"fmt"
	"testing"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestFeeToLow(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	minFee := sdk.NewCoins(sdk.NewCoin("regen", sdk.NewInt(100)))

	// no fee specified should fail
	gmAny := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.BasketFee = minFee
	}).Times(2)

	// no fee
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Fee: nil,
	})
	assert.ErrorContains(t, err, "insufficient fee")

	// fee too low
	_, err = s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Fee: sdk.NewCoins(sdk.NewCoin("regen", sdk.NewInt(20))),
	})
	assert.ErrorContains(t, err, "insufficient fee")
}

func TestInvalidCreditType(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	basketFee := sdk.Coin{Denom: "foo", Amount: sdk.NewInt(10)}
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.BasketFee = sdk.Coins{basketFee}
		p.CreditTypes = []*core.CreditType{{Abbreviation: "C", Precision: 6}}
	}).Times(2)
	s.distKeeper.EXPECT().FundCommunityPool(gmAny, gmAny, gmAny).Times(2)

	// non-existent credit type should fail
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "F",
		Fee:              sdk.Coins{basketFee},
	})
	assert.ErrorContains(t, err, `credit type abbreviation "F" doesn't exist`)

	// exponent < precision should fail
	_, err = s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         2,
		Fee:              sdk.Coins{basketFee},
	})
	assert.ErrorContains(t, err, "exponent")
}

func TestDuplicateDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	fee := sdk.Coin{Denom: "foo", Amount: sdk.NewInt(10)}
	mc := baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "foo",
		Fee:              sdk.Coins{fee},
	}
	denom, _, err := basket.BasketDenom(mc.Name, mc.CreditTypeAbbrev, mc.Exponent)
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.BasketTable().Insert(s.ctx,
		&api.Basket{BasketDenom: denom},
	))

	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.BasketFee = sdk.Coins{fee}
		p.CreditTypes = []*core.CreditType{{Precision: 6, Abbreviation: "C"}}
	}).Times(1)
	s.distKeeper.EXPECT().FundCommunityPool(gmAny, gmAny, gmAny)

	_, err = s.k.Create(s.ctx, &mc)
	assert.ErrorContains(t, err, "unique")
}

func TestInvalidClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	assert.NilError(t, s.coreStore.ClassInfoTable().Insert(s.ctx, &ecoApi.ClassInfo{
		Id:               "bar",
		CreditTypeAbbrev: "BIO",
	}))
	mockAny := gomock.Any()
	basketFee := sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdk.NewInt(10)}}
	s.paramsKeeper.EXPECT().GetParamSet(mockAny, mockAny).Do(func(any interface{}, p *core.Params) {
		p.BasketFee = basketFee
		p.CreditTypes = []*core.CreditType{{Abbreviation: "C", Precision: 6}}
	}).Times(2)
	s.distKeeper.EXPECT().FundCommunityPool(mockAny, mockAny, mockAny).Return(nil).Times(2)

	// class doesn't exist
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "foo",
		AllowedClasses:   []string{"foo"},
		Fee:              basketFee,
	})
	assert.ErrorContains(t, err, "could not get credit class")

	// mismatch credit type and class's credit type
	_, err = s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "bar",
		AllowedClasses:   []string{"bar"},
		Fee:              basketFee,
	})
	assert.ErrorContains(t, err, "basket specified credit type C, but class bar is of type BIO")
}

func TestValidBasket(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	gmAny := gomock.Any()
	fee := sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdk.NewInt(10)}}
	s.distKeeper.EXPECT().FundCommunityPool(gmAny, gmAny, gmAny)
	seconds := time.Hour * 24 * 356 * 5
	dateCriteria := &baskettypes.DateCriteria{
		StartDateWindow: gogotypes.DurationProto(seconds),
	}
	s.bankKeeper.EXPECT().SetDenomMetaData(gmAny,
		banktypes.Metadata{
			Name:        "foo",
			Display:     "eco.C.foo",
			Base:        "eco.uC.foo",
			Symbol:      "foo",
			Description: "hi",
			DenomUnits: []*banktypes.DenomUnit{{
				Denom:    "eco.C.foo",
				Exponent: 6,
			}, {
				Denom:    "eco.uC.foo",
				Exponent: 0,
			}},
		},
	)
	assert.NilError(t, s.coreStore.ClassInfoTable().Insert(s.ctx, &ecoApi.ClassInfo{
		Id:               "bar",
		Admin:            nil,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	}))
	s.paramsKeeper.EXPECT().GetParamSet(gmAny, gmAny).Do(func(any interface{}, p *core.Params) {
		p.BasketFee = fee
		p.CreditTypes = []*core.CreditType{{Abbreviation: "C", Precision: 6}}
	}).Times(1)

	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		Description:      "hi",
		Name:             "foo",
		CreditTypeAbbrev: "C",
		Exponent:         6,
		AllowedClasses:   []string{"bar"},
		DateCriteria:     dateCriteria,
		Fee:              fee,
	})
	assert.NilError(t, err)

	bskt, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, "eco.uC.foo")
	assert.NilError(t, err)
	assert.Equal(t, s.addr.String(), bskt.Curator)
	assert.Equal(t, "eco.uC.foo", bskt.BasketDenom)
	assert.Equal(t, uint32(6), bskt.Exponent)
	assert.Equal(t, "C", bskt.CreditTypeAbbrev)
	assert.Equal(t, fmt.Sprintf("seconds:%.0f", seconds.Seconds()), bskt.DateCriteria.StartDateWindow.String())
}
