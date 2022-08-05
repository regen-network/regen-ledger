package core

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// AddCreditType adds a new credit type to the network.
func (k Keeper) AddCreditType(ctx context.Context, req *core.MsgAddCreditType) (*core.MsgAddCreditTypeResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	found, err := k.stateStore.CreditTypeTable().HasByName(ctx, req.CreditType.Name)
	if err != nil {
		return nil, err
	}

	if found {
		return nil, sdkerrors.ErrConflict.Wrapf("credit type with %s name already exists", req.CreditType.Name)
	}

	if err := k.stateStore.CreditTypeTable().Insert(ctx, &ecocreditv1.CreditType{
		Abbreviation: req.CreditType.Abbreviation,
		Name:         req.CreditType.Name,
		Unit:         req.CreditType.Unit,
		Precision:    req.CreditType.Precision,
	}); err != nil {
		return nil, err
	}

	return &core.MsgAddCreditTypeResponse{}, nil
}
