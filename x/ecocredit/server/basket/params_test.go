package basket_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams_UpdateBasketFee(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	fee := sdk.NewInt64Coin("foo", 50)

	err := s.stateStore.BasketFeeTable().Insert(s.ctx, &api.BasketFee{
		Denom:  fee.Denom,
		Amount: fee.Amount.String(),
	})
	assert.NilError(t, err)

	addFee := sdk.NewInt64Coin("bar", 20)

	govAddr := sdk.AccAddress("foo")
	s.accountKeeper.EXPECT().GetModuleAddress(gomock.Any()).Return(govAddr).Times(3)

	_, err = s.k.UpdateBasketFee(s.ctx, &basket.MsgUpdateBasketFeeRequest{
		RootAddress: govAddr.String(),
		AddFees:     []*sdk.Coin{&addFee},
		RemoveFees:  []string{fee.Denom},
	})
	assert.NilError(t, err)

	res, err := s.stateStore.BasketFeeTable().Get(s.ctx, addFee.Denom)
	assert.NilError(t, err)
	assert.Equal(t, res.Amount, addFee.Amount.String())

	_, err = s.stateStore.BasketFeeTable().Get(s.ctx, fee.Denom)
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	// no duplicates
	_, err = s.k.UpdateBasketFee(s.ctx, &basket.MsgUpdateBasketFeeRequest{
		RootAddress: govAddr.String(),
		AddFees:     []*sdk.Coin{&addFee},
	})
	assert.ErrorContains(t, err, ormerrors.PrimaryKeyConstraintViolation.Error())

	_, err = s.k.UpdateBasketFee(s.ctx, &basket.MsgUpdateBasketFeeRequest{RootAddress: sdk.AccAddress("not_governance").String()})
	assert.ErrorContains(t, err, "params can only be updated via governance")
}
