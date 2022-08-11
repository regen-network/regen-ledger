package marketplace

import (
	"context"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	marketplacev1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

// RemoveAllowedDenom removes denom from the allowed denoms.
func (k Keeper) RemoveAllowedDenom(ctx context.Context, req *marketplace.MsgRemoveAllowedDenom) (*marketplace.MsgRemoveAllowedDenomResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	if err := k.stateStore.AllowedDenomTable().Delete(ctx, &marketplacev1.AllowedDenom{
		BankDenom: req.Denom,
	}); err != nil {
		return nil, err
	}

	return &marketplace.MsgRemoveAllowedDenomResponse{}, nil
}
