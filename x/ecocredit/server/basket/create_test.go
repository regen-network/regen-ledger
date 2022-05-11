package basket_test

import (
	"fmt"
	"testing"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/utils"
)

func TestCreate_InvalidFees(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	utils.ExpectParamGet(&basketFees, s.paramsKeeper, core.KeyBasketCreationFee, 2)

	// no fee
	_, err := s.k.Create(s.ctx, &basket.MsgCreate{
		Fee: nil,
	})
	assert.ErrorContains(t, err, "insufficient fee")

	// fee too low
	_, err = s.k.Create(s.ctx, &basket.MsgCreate{
		Fee: sdk.NewCoins(sdk.NewCoin(validFee.Denom, validFee.Amount.Sub(sdk.NewInt(1)))),
	})
	assert.ErrorContains(t, err, "insufficient fee")
}

func TestCreate_InvalidCreditType(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	utils.ExpectParamGet(&basketFees, s.paramsKeeper, core.KeyBasketCreationFee, 2)
	s.distKeeper.EXPECT().FundCommunityPool(gmAny, gmAny, gmAny).Times(2)

	// non-existent credit type should fail
	_, err := s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.addrs[0].String(),
		CreditTypeAbbrev: "F",
		Fee:              basketFees,
	})
	assert.ErrorContains(t, err, `could not get credit type with abbreviation F: not found`)

	// exponent < precision should fail
	_, err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.addrs[0].String(),
		CreditTypeAbbrev: "C",
		Exponent:         2,
		Fee:              basketFees,
	})
	assert.ErrorContains(t, err, "exponent")
}

func TestCreate_DuplicateDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	mc := basket.MsgCreate{
		Curator:          s.addrs[0].String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "foo",
		Fee:              basketFees,
	}
	denom, _, err := basket.BasketDenom(mc.Name, mc.CreditTypeAbbrev, mc.Exponent)
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.BasketTable().Insert(s.ctx,
		&api.Basket{BasketDenom: denom},
	))

	utils.ExpectParamGet(&basketFees, s.paramsKeeper, core.KeyBasketCreationFee, 1)
	s.distKeeper.EXPECT().FundCommunityPool(gmAny, gmAny, gmAny)

	_, err = s.k.Create(s.ctx, &mc)
	assert.ErrorContains(t, err, "unique")
}

func TestCreate_InvalidClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	assert.NilError(t, s.coreStore.ClassTable().Insert(s.ctx, &ecoApi.Class{
		Id:               "bar",
		CreditTypeAbbrev: "BIO",
	}))

	utils.ExpectParamGet(&basketFees, s.paramsKeeper, core.KeyBasketCreationFee, 2)
	s.distKeeper.EXPECT().FundCommunityPool(gmAny, gmAny, gmAny).Return(nil).Times(2)

	// class doesn't exist
	_, err := s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.addrs[0].String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "foo",
		AllowedClasses:   []string{"foo"},
		Fee:              basketFees,
	})
	assert.ErrorContains(t, err, "could not get credit class")

	// mismatch credit type and class's credit type
	_, err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.addrs[0].String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "bar",
		AllowedClasses:   []string{"bar"},
		Fee:              basketFees,
	})
	assert.ErrorContains(t, err, "basket specified credit type C, but class bar is of type BIO")
}

func TestCreate_ValidBasket(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	s.distKeeper.EXPECT().FundCommunityPool(gmAny, gmAny, gmAny)
	seconds := time.Hour * 24 * 356 * 5
	dateCriteria := &basket.DateCriteria{
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
	assert.NilError(t, s.coreStore.ClassTable().Insert(s.ctx, &ecoApi.Class{
		Id:               "bar",
		Admin:            nil,
		Metadata:         "",
		CreditTypeAbbrev: "C",
	}))
	utils.ExpectParamGet(&basketFees, s.paramsKeeper, core.KeyBasketCreationFee, 1)

	_, err := s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.addrs[0].String(),
		Description:      "hi",
		Name:             "foo",
		CreditTypeAbbrev: "C",
		Exponent:         6,
		AllowedClasses:   []string{"bar"},
		DateCriteria:     dateCriteria,
		Fee:              basketFees,
	})
	assert.NilError(t, err)

	bskt, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, "eco.uC.foo")
	assert.NilError(t, err)
	assert.Equal(t, s.addrs[0].String(), sdk.AccAddress(bskt.Curator).String())
	assert.Equal(t, "eco.uC.foo", bskt.BasketDenom)
	assert.Equal(t, uint32(6), bskt.Exponent)
	assert.Equal(t, "C", bskt.CreditTypeAbbrev)
	assert.Equal(t, fmt.Sprintf("seconds:%.0f", seconds.Seconds()), bskt.DateCriteria.StartDateWindow.String())
}
