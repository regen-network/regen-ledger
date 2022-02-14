package basket

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
)

func (k Keeper) Create(ctx context.Context, create *baskettypes.MsgCreate) (*baskettypes.MsgCreateResponse, error) {
	// k.ecocreditKeeper.GetCreateBasketFee()
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sender, _ := sdk.AccAddressFromBech32(create.Curator)
	err := k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, sender, ecocredit.ModuleName, sdk.NewCoins(sdk.NewInt64Coin("foo", 20)))
	if err != nil {
		return nil, err
	}
	err = k.stateStore.BasketStore().Insert(ctx, &v1.Basket{
		BasketDenom:       "FooBar",
		DisableAutoRetire: false,
		CreditTypeName:    "carbon",
		DateCriteria:      nil,
		Exponent:          6,
	})
	return nil, err
}
