package core

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// AddCreditType adds a new credit type to the network.
func (k Keeper) AddCreditType(ctx context.Context, req *core.MsgAddCreditType) (*core.MsgAddCreditTypeResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	for _, cType := range req.CreditType {
		found, err := k.stateStore.CreditTypeTable().HasByName(ctx, cType.Name)
		if err != nil {
			return nil, err
		}

		if found {
			return nil, sdkerrors.ErrConflict.Wrapf("credit type with %s name already exists", cType.Name)
		}

		if err := k.stateStore.CreditTypeTable().Insert(ctx, &ecocreditv1.CreditType{
			Abbreviation: cType.Abbreviation,
			Name:         cType.Name,
			Unit:         cType.Unit,
			Precision:    cType.Precision,
		}); err != nil {
			return nil, err
		}

		sdkCtx := sdk.UnwrapSDKContext(ctx)
		sdkCtx.GasMeter().ConsumeGas(ecocredit.GasCostPerIteration, "ecocredit/core/AddCreditType credit type iteration")
	}

	return &core.MsgAddCreditTypeResponse{}, nil
}
