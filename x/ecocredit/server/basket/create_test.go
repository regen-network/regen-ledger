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
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestFeeToLow(t *testing.T) {
	t.Parallel()

	s := setupBase(t)
	minFee := sdk.NewCoins(sdk.NewCoin("regen", sdk.NewInt(100)))

	s.ecocreditKeeper.EXPECT().GetCreateBasketFee(gomock.Any()).Return(minFee)
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Fee: nil,
	})
	assert.ErrorContains(t, err, "insufficient fee")

	s.ecocreditKeeper.EXPECT().GetCreateBasketFee(gomock.Any()).Return(minFee)
	_, err = s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Fee: sdk.NewCoins(sdk.NewCoin("regen", sdk.NewInt(20))),
	})
	assert.ErrorContains(t, err, "insufficient fee")
}

func TestBadCreditType(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	s.ecocreditKeeper.EXPECT().GetCreateBasketFee(gomock.Any()).Return(nil) // nil fee
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any())
	s.ecocreditKeeper.EXPECT().CreditTypes(gomock.Any(), gomock.Any()).Return(
		&ecocredit.QueryCreditTypesResponse{CreditTypes: []*ecocredit.CreditType{
			{Abbreviation: "B"}, {Abbreviation: "C"},
		}}, nil,
	)
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "F",
	})
	assert.ErrorContains(t, err, `credit type abbreviation "F" doesn't exist`)

	s.ecocreditKeeper.EXPECT().GetCreateBasketFee(gomock.Any()).Return(nil) // nil fee
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any())
	s.ecocreditKeeper.EXPECT().CreditTypes(gomock.Any(), gomock.Any()).Return(
		&ecocredit.QueryCreditTypesResponse{CreditTypes: []*ecocredit.CreditType{
			{Abbreviation: "B", Precision: 3}, {Abbreviation: "C", Precision: 6},
		}}, nil,
	)
	_, err = s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         3,
	})
	assert.ErrorContains(t, err, "exponent")
}

func TestDuplicateDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	mc := baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "foo",
	}
	denom, _, err := basket.BasketDenom(mc.Name, mc.CreditTypeAbbrev, mc.Exponent)
	assert.NilError(t, err)
	assert.NilError(t, s.stateStore.BasketTable().Insert(s.ctx,
		&api.Basket{BasketDenom: denom},
	))

	s.ecocreditKeeper.EXPECT().GetCreateBasketFee(gomock.Any()).Return(nil) // nil fee
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any())
	s.ecocreditKeeper.EXPECT().CreditTypes(gomock.Any(), gomock.Any()).Return(
		&ecocredit.QueryCreditTypesResponse{CreditTypes: []*ecocredit.CreditType{
			{Abbreviation: "B", Precision: 3}, {Abbreviation: "C", Precision: 6},
		}}, nil,
	)
	_, err = s.k.Create(s.ctx, &mc)
	assert.ErrorContains(t, err, "unique")
}

func TestMissingClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	s.ecocreditKeeper.EXPECT().GetCreateBasketFee(gomock.Any()).Return(nil) // nil fee
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any())
	s.ecocreditKeeper.EXPECT().CreditTypes(gomock.Any(), gomock.Any()).Return(
		&ecocredit.QueryCreditTypesResponse{CreditTypes: []*ecocredit.CreditType{
			{Abbreviation: "B", Precision: 3}, {Abbreviation: "C", Precision: 6},
		}}, nil,
	)
	s.ecocreditKeeper.EXPECT().HasClassInfo(gomock.Any(), gomock.Any()).Return(false)
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "foo",
		AllowedClasses:   []string{"bar"},
	})
	assert.ErrorContains(t, err, "doesn't exist")
}

func TestGoodBasket(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	s.ecocreditKeeper.EXPECT().GetCreateBasketFee(gomock.Any()).Return(nil) // nil fee
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any())
	s.ecocreditKeeper.EXPECT().CreditTypes(gomock.Any(), gomock.Any()).Return(
		&ecocredit.QueryCreditTypesResponse{CreditTypes: []*ecocredit.CreditType{
			{Abbreviation: "B", Precision: 3}, {Abbreviation: "C", Precision: 6},
		}}, nil,
	)
	s.ecocreditKeeper.EXPECT().HasClassInfo(gomock.Any(), "bar").Return(true)
	seconds := time.Hour * 24 * 356 * 5
	dateCriteria := &baskettypes.DateCriteria{
		StartDateWindow: gogotypes.DurationProto(seconds),
	}
	s.bankKeeper.EXPECT().SetDenomMetaData(gomock.Any(),
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

	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		Description:      "hi",
		Name:             "foo",
		CreditTypeAbbrev: "C",
		Exponent:         6,
		AllowedClasses:   []string{"bar"},
		DateCriteria:     dateCriteria,
	})
	assert.NilError(t, err)

	basket, err := s.stateStore.BasketTable().GetByBasketDenom(s.ctx, "eco.uC.foo")
	assert.NilError(t, err)
	assert.Equal(t, s.addr.String(), basket.Curator)
	assert.Equal(t, "eco.uC.foo", basket.BasketDenom)
	assert.Equal(t, uint32(6), basket.Exponent)
	assert.Equal(t, "C", basket.CreditTypeAbbrev)
	assert.Equal(t, fmt.Sprintf("seconds:%.0f", seconds.Seconds()), basket.DateCriteria.StartDateWindow.String())
}
