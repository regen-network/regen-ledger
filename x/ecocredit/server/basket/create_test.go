package basket_test

import (
	"fmt"
	"testing"
	"time"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func TestFeeToLow(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	minFee := sdk.NewCoins(sdk.NewCoin("regen", sdk.NewInt(100)))

	// no fee specified should fail
	s.ecocreditKeeper.EXPECT().Params(gomock.Any(), gomock.Any()).Return(&core.QueryParamsResponse{Params: &core.Params{BasketCreationFee: minFee}})
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Fee: nil,
	})
	assert.ErrorContains(t, err, "insufficient fee")

	// fee too low
	s.ecocreditKeeper.EXPECT().Params(gomock.Any(), gomock.Any()).Return(&core.QueryParamsResponse{Params: &core.Params{
		BasketCreationFee: minFee,
	}})
	_, err = s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Fee: sdk.NewCoins(sdk.NewCoin("regen", sdk.NewInt(20))),
	})
	assert.ErrorContains(t, err, "insufficient fee")
}

func TestBadCreditType(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	basketFee := sdk.Coin{Denom: "foo", Amount: sdk.NewInt(10)}

	s.ecocreditKeeper.EXPECT().Params(gomock.Any(), gomock.Any()).Return(&core.QueryParamsResponse{Params: &core.Params{
		BasketCreationFee: sdk.Coins{basketFee},
		CreditTypes:       []*core.CreditType{{Abbreviation: "A", Precision: 6}},
	}}).Times(2)
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any()).Times(2)

	// non-existent credit type should fail
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "F",
	})
	assert.ErrorContains(t, err, `credit type abbreviation "F" doesn't exist`)

	// exponent < precision should fail
	_, err = s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         2,
	})
	assert.ErrorContains(t, err, "exponent")
}

func TestDuplicateDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
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

	s.ecocreditKeeper.EXPECT().Params(gomock.Any(), gomock.Any()).Return(&core.QueryParamsResponse{
		Params: &core.Params{BasketCreationFee: sdk.Coins{fee},
			CreditTypes: []*core.CreditType{{Precision: 6, Abbreviation: "C"}},
		},
	}).Times(1)
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any())

	_, err = s.k.Create(s.ctx, &mc)
	assert.ErrorContains(t, err, "unique")
}

func TestBadClass(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	mockAny := gomock.Any()
	basketFee := sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdk.NewInt(10)}}
	s.ecocreditKeeper.EXPECT().Params(mockAny, mockAny).Return(&core.QueryParamsResponse{Params: &core.Params{
		BasketCreationFee: basketFee,
		CreditTypes:       []*core.CreditType{{Abbreviation: "C", Precision: 6}},
	},
	})
	s.ecocreditKeeper.EXPECT().ClassInfo(mockAny, mockAny).Return(nil, fmt.Errorf("not found"))
	s.distKeeper.EXPECT().FundCommunityPool(mockAny, mockAny, mockAny)

	// class doesn't exist
	_, err := s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "foo",
		AllowedClasses:   []string{"bar"},
	})
	assert.ErrorContains(t, err, "could not get credit class")

	// mismatch credit type and class's credit type
	s.ecocreditKeeper.EXPECT().ClassInfo(mockAny, mockAny).Return(&core.QueryClassInfoResponse{Info: &core.ClassInfo{CreditType: "B"}}, nil)
	_, err = s.k.Create(s.ctx, &baskettypes.MsgCreate{
		Curator:          s.addr.String(),
		CreditTypeAbbrev: "C",
		Exponent:         6,
		Name:             "foo",
		AllowedClasses:   []string{"bar"},
	})
	assert.ErrorContains(t, err, "basket specified credit type C but class bar is of type B")
}

func TestGoodBasket(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	fee := sdk.Coins{sdk.Coin{Denom: "foo", Amount: sdk.NewInt(10)}}
	s.distKeeper.EXPECT().FundCommunityPool(gomock.Any(), gomock.Any(), gomock.Any())
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
	s.ecocreditKeeper.EXPECT().Params(gomock.Any(), gomock.Any()).Return(&core.QueryParamsResponse{Params: &core.Params{
		BasketCreationFee: fee,
		CreditTypes:       []*core.CreditType{{Abbreviation: "C", Precision: 6}},
	}})
	s.ecocreditKeeper.EXPECT().ClassInfo(gomock.Any(), gomock.Any()).Return(&core.QueryClassInfoResponse{Info: &core.ClassInfo{
		CreditType: "C",
	}}, nil)

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
