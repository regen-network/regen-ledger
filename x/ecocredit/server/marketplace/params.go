package marketplace

import (
	"context"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	v1 "github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AllowAskDenom(ctx context.Context, req *v1.MsgAllowAskDenom) (*v1.MsgAllowAskDenomResponse, error) {
	govAcc, err := sdk.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = server.AssertGovernance(govAcc, k.accountKeeper); err != nil {
		return nil, err
	}

	for _, add := range req.AddDenoms {
		if err = sdk.ValidateDenom(add.Denom); err != nil {
			return nil, err
		}
		// TODO: validate display denom?
		if err = k.stateStore.AllowedDenomTable().Insert(ctx, &marketplacev1.AllowedDenom{
			BankDenom:    add.Denom,
			DisplayDenom: add.DisplayDenom,
			Exponent:     add.Exponent,
		}); err != nil {
			return nil, err
		}
	}

	for _, denom := range req.RemoveDenoms {
		if err = k.stateStore.AllowedDenomTable().Delete(ctx, &marketplacev1.AllowedDenom{BankDenom: denom}); err != nil {
			return nil, err
		}
	}

	return &v1.MsgAllowAskDenomResponse{}, nil
}
