package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v4/marketplace/types/v1"
)

func (k Keeper) GovSendFromFeePool(ctx context.Context, msg *types.MsgGovSendFromFeePool) (*types.MsgGovSendFromFeePoolResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return nil, err
	}

	if !authority.Equals(k.authority) {
		return nil, sdkerrors.ErrUnauthorized
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(ctx), k.feePoolName, recipient, msg.Coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgGovSendFromFeePoolResponse{}, nil
}
