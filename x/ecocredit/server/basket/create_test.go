package basket_test

import (
	"fmt"
	"testing"
	"time"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	gogotypes "github.com/gogo/protobuf/types"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"

	"github.com/regen-network/regen-ledger/x/ecocredit"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"gotest.tools/v3/assert"
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
		Curator:        s.addr.String(),
		CreditTypeName: "F",
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
		Curator:        s.addr.String(),
		CreditTypeName: "C",
		Exponent:       3,
	})
	assert.ErrorContains(t, err, "exponent")
}

func TestDuplicateDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	assert.NilError(t, s.stateStore.BasketStore().Insert(s.ctx,
		&basketv1.Basket{BasketDenom: "foo"},
	))

	s.ecocreditKeeper.EXPECT().GetCreateBasketFee(gomock.Any()).Return(nil) // nil fee
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any())
	s.ecocreditKeeper.EXPECT().CreditTypes(gomock.Any(), gomock.Any()).Return(
		&ecocredit.QueryCreditTypesResponse{CreditTypes: []*ecocredit.CreditType{
			{Abbreviation: "B", Precision: 3}, {Abbreviation: "C", Precision: 6},
		}}, nil,
	)
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:        s.addr.String(),
		CreditTypeName: "C",
		Exponent:       6,
		Name:           "foo",
	})
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
		Curator:        s.addr.String(),
		CreditTypeName: "C",
		Exponent:       6,
		Name:           "foo",
		AllowedClasses: []string{"bar"},
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
	dateCriteria := &baskettypes.DateCriteria_StartDateWindow{
		StartDateWindow: gogotypes.DurationProto(seconds),
	}
	s.bankKeeper.EXPECT().SetDenomMetaData(gomock.Any(),
		banktypes.Metadata{
			Name:    "foo",
			Display: "foo",
			Base:    "ufoo",
			Symbol:  "foo",
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    "foo",
					Exponent: 6,
				},
			},
		},
	)

	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:        s.addr.String(),
		CreditTypeName: "C",
		Exponent:       6,
		Name:           "ufoo",
		DisplayName:    "foo",
		AllowedClasses: []string{"bar"},
		DateCriteria:   &baskettypes.DateCriteria{Sum: dateCriteria},
	})
	assert.NilError(t, err)

	basket, err := s.stateStore.BasketStore().GetByBasketDenom(s.ctx, "ufoo")
	assert.NilError(t, err)
	assert.Equal(t, "ufoo", basket.BasketDenom)
	assert.Equal(t, uint32(6), basket.Exponent)
	assert.Equal(t, "C", basket.CreditTypeName)
	assert.Equal(t, fmt.Sprintf("seconds:%.0f", seconds.Seconds()), basket.DateCriteria.Sum.(*basketv1.DateCriteria_StartDateWindow).StartDateWindow.String())
}
