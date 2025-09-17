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

	authorityBz, err := k.ac.StringToBytes(msg.Authority)
	if err != nil {
		return nil, err
	}
	authorityAddr := sdk.AccAddress(authorityBz)

	if !authorityAddr.Equals(k.authority) {
		return nil, sdkerrors.ErrUnauthorized
	}

	recipientBz, err := k.ac.StringToBytes(msg.Recipient)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(ctx), k.feePoolName, recipientBz, msg.Coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgGovSendFromFeePoolResponse{}, nil
}
