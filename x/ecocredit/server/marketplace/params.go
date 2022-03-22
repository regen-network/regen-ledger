package marketplace

import (
	"context"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AllowAskDenom(ctx context.Context, req *marketplace.MsgAllowAskDenom) (*marketplace.MsgAllowAskDenomResponse, error) {
	govAcc, err := sdk.AccAddressFromBech32(req.RootAddress)
	if err != nil {
		return nil, err
	}
	if err = ecocredit.AssertGovernance(govAcc, k.accountKeeper); err != nil {
		return nil, err
	}

	for _, denom := range req.RemoveDenoms {
		if err = k.stateStore.AllowedDenomTable().Delete(ctx, &api.AllowedDenom{BankDenom: denom}); err != nil {
			return nil, err
		}
	}

	for _, add := range req.AddDenoms {
		if err = sdk.ValidateDenom(add.Denom); err != nil {
			return nil, err
		}
		// TODO: validate display denom?
		if err = k.stateStore.AllowedDenomTable().Insert(ctx, &api.AllowedDenom{
			BankDenom:    add.Denom,
			DisplayDenom: add.DisplayDenom,
			Exponent:     add.Exponent,
		}); err != nil {
			return nil, err
		}
	}

	return &marketplace.MsgAllowAskDenomResponse{}, nil
}
