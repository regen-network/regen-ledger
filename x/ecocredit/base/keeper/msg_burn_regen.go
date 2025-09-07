package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/ecocredit/v4"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/base/types/v1"
)

func (k Keeper) BurnRegen(ctx context.Context, regen *types.MsgBurnRegen) (*types.MsgBurnRegenResponse, error) {
	if err := regen.ValidateBasic(); err != nil {
		return nil, err
	}

	from, err := sdk.AccAddressFromBech32(regen.Burner)
	if err != nil {
		return nil, err
	}

	amount, ok := math.NewIntFromString(regen.Amount)
	if !ok {
		return nil, fmt.Errorf("invalid amount: %s", regen.Amount)
	}
	if !amount.IsPositive() {
		return nil, fmt.Errorf("amount must be positive: %s", regen.Amount)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	coins := sdk.NewCoins(sdk.NewCoin("uregen", amount))

	err = k.bankKeeper.SendCoinsFromAccountToModule(sdkCtx, from, ecocredit.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(sdkCtx, ecocredit.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&types.EventBurnRegen{
		Burner: regen.Burner,
		Amount: regen.Amount,
		Reason: regen.Reason,
	})

	return &types.MsgBurnRegenResponse{}, err
}
